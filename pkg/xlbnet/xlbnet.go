package xlbnet

import (
	"errors"
	"fmt"
	apiserver "github.com/flomesh-io/fsmxlb/api"
	prometheus "github.com/flomesh-io/fsmxlb/api/prometheus"
	nlp "github.com/flomesh-io/fsmxlb/api/xlbnlp"
	cmn "github.com/flomesh-io/fsmxlb/pkg/common"
	opts "github.com/flomesh-io/fsmxlb/pkg/options"
	tk "github.com/flomesh-io/fsmxlb/pkg/xlblib"
	"io"
	"net"
	"net/http"
	_ "net/http/pprof"
	"net/rpc"
	"os"
	"os/signal"
	"runtime/debug"
	"runtime/pprof"
	"strings"
	"sync"
	"syscall"
	"time"
)

// string constant representing root security zone
const (
	RootZone = "root"
)

// constants
const (
	XlbnetTiVal    = 10
	GoBGPInitTiVal = 5
	KAInitTiVal    = 35
)

// utility variables
const (
	MkfsScript     = "/usr/local/sbin/mkxlb_bpffs"
	BpfFsCheckFile = "/opt/fsmxlb/dp/bpf/intf_map"
)

type xlbNetH struct {
	dpEbpf *DpEbpfH
	dp     *DpH
	zn     *ZoneH
	zr     *Zone
	mtx    sync.RWMutex
	ticker *time.Ticker
	tDone  chan bool
	sigCh  chan os.Signal
	wg     sync.WaitGroup
	bgp    *GoBgpH
	sumDis bool
	pProbe bool
	has    *CIStateH
	logger *tk.Logger
	ready  bool
	self   int
	rssEn  bool
	pFile  *os.File
}

// XlbXsyncMain - State Sync subsystem init
func XlbXsyncMain() {
	if opts.Opts.ClusterNodes == "none" {
		return
	}

	// Stack trace logger
	defer func() {
		if e := recover(); e != nil {
			if mh.logger != nil {
				tk.LogIt(tk.LogCritical, "%s: %s", e, debug.Stack())
			}
		}
	}()

	for {
		rpcObj := new(XSync)
		err := rpc.Register(rpcObj)
		if err != nil {
			panic("Failed to register rpc")
		}

		rpc.HandleHTTP()

		http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
			io.WriteString(res, "fsmxlb-xsync\n")
		})

		listener := fmt.Sprintf(":%d", XSyncPort)
		err = http.ListenAndServe(listener, nil)
		if err != nil {
			panic("Failed to rpc-listen")
		}
	}
}

// NodeWalker - an implementation of node walker interface
func (mh *xlbNetH) NodeWalker(b string) {
	tk.LogIt(tk.LogDebug, "%s\n", b)
}

// ParamSet - Set Xlbnet Params
func (mh *xlbNetH) ParamSet(param cmn.ParamMod) (int, error) {
	logLevel := LogString2Level(param.LogLevel)

	if mh.logger != nil {
		mh.logger.LogItSetLevel(logLevel)
	}

	DpEbpfSetLogLevel(logLevel)

	return 0, nil
}

// ParamGet - Get Xlbnet Params
func (mh *xlbNetH) ParamGet(param *cmn.ParamMod) (int, error) {
	logLevel := "n/a"
	switch mh.logger.CurrLogLevel {
	case tk.LogDebug:
		logLevel = "debug"
	case tk.LogInfo:
		logLevel = "info"
	case tk.LogError:
		logLevel = "error"
	case tk.LogNotice:
		logLevel = "notice"
	case tk.LogWarning:
		logLevel = "warning"
	case tk.LogAlert:
		logLevel = "alert"
	case tk.LogCritical:
		logLevel = "critical"
	case tk.LogEmerg:
		logLevel = "emergency"
	default:
		param.LogLevel = logLevel
		return -1, errors.New("unknown log level")
	}

	param.LogLevel = logLevel
	return 0, nil
}

// xlbNetTicker - this ticker routine runs every XLBNET_TIVAL seconds
func xlbNetTicker() {
	for {
		select {
		case <-mh.tDone:
			return
		case sig := <-mh.sigCh:
			if sig == syscall.SIGCHLD {
				var ws syscall.WaitStatus
				var ru syscall.Rusage
				wpid := 1
				try := 0
				for wpid >= 0 && try < 100 {
					wpid, _ = syscall.Wait4(-1, &ws, syscall.WNOHANG, &ru)
					try++
				}
			} else if sig == syscall.SIGHUP {
				tk.LogIt(tk.LogCritical, "SIGHUP received\n")
				pprof.StopCPUProfile()
			} else if sig == syscall.SIGINT || sig == syscall.SIGTERM {
				tk.LogIt(tk.LogCritical, "Shutdown on sig %v\n", sig)
				mh.dpEbpf.DpEbpfUnInit()
				apiserver.ApiServerShutOk()
			}
		case t := <-mh.ticker.C:
			tk.LogIt(-1, "Tick at %v\n", t)
			// Do any housekeeping activities for security zones
			mh.zn.ZoneTicker()
			mh.has.CITicker()
		}
	}
}

var mh xlbNetH

func xlbNetInit() {
	spawnKa, kaMode := KAString2Mode(opts.Opts.Ka)
	clusterMode := false
	if opts.Opts.ClusterNodes != "none" {
		clusterMode = true
	}

	// Initialize logger and specify the log file
	logfile := fmt.Sprintf("%s%s.log", "/var/log/fsmxlb", os.Getenv("HOSTNAME"))
	logLevel := LogString2Level(opts.Opts.LogLevel)
	mh.logger = tk.LogItInit(logfile, logLevel, true)

	// Stack trace logger
	defer func() {
		if e := recover(); e != nil {
			tk.LogIt(tk.LogCritical, "%s: %s", e, debug.Stack())
		}
	}()

	// It is important to make sure fsmxlb's eBPF filesystem
	// is in place and mounted to make sure maps are pinned properly
	if FileExists(BpfFsCheckFile) == false {
		if FileExists(MkfsScript) {
			RunCommand(MkfsScript, true)
		}
	}

	mh.self = opts.Opts.ClusterSelf
	mh.rssEn = opts.Opts.RssEnable
	mh.sumDis = opts.Opts.CSumDisable
	mh.pProbe = opts.Opts.PassiveEPProbe
	mh.sigCh = make(chan os.Signal, 5)
	signal.Notify(mh.sigCh, os.Interrupt, syscall.SIGCHLD, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	// Check if profiling is enabled
	if opts.Opts.CPUProfile != "none" {
		var err error
		mh.pFile, err = os.Create(opts.Opts.CPUProfile)
		if err != nil {
			tk.LogIt(tk.LogNotice, "profile file create failed\n")
			return
		}
		err = pprof.StartCPUProfile(mh.pFile)
		if err != nil {
			tk.LogIt(tk.LogNotice, "CPU profiler start failed\n")
			return
		}
	}

	// Initialize the ebpf datapath subsystem
	mh.dpEbpf = DpEbpfInit(clusterMode, mh.self, mh.rssEn, -1)
	mh.dp = DpBrokerInit(mh.dpEbpf)

	// Initialize the security zone subsystem
	mh.zn = ZoneInit()

	// Add a root zone by default
	mh.zn.ZoneAdd(RootZone)
	mh.zr, _ = mh.zn.Zonefind(RootZone)
	if mh.zr == nil {
		tk.LogIt(tk.LogError, "root zone not found\n")
		return
	}

	// Initialize the clustering subsystem
	mh.has = CIInit(spawnKa, kaMode)
	if clusterMode {
		// Add cluster nodes if specified
		cNodes := strings.Split(opts.Opts.ClusterNodes, ",")
		for _, cNode := range cNodes {
			addr := net.ParseIP(cNode)
			if addr == nil {
				continue
			}
			mh.has.ClusterNodeAdd(cmn.CluserNodeMod{Addr: addr})
		}
	}

	// Initialize goBgp client
	if opts.Opts.Bgp {
		mh.bgp = GoBgpInit()
	}

	// Initialize and spawn the api server subsystem
	if opts.Opts.NoApi == false {
		apiserver.RegisterAPIHooks(NetAPIInit())
		go apiserver.RunAPIServer()
		apiserver.WaitAPIServerReady()
	}

	// Initialize the nlp subsystem
	if opts.Opts.NoNlp == false {
		nlp.NlpRegister(NetAPIInit())
		nlp.NlpInit()
	}

	// Initialize the Prometheus subsystem
	if opts.Opts.NoPrometheus == false {
		prometheus.PrometheusRegister(NetAPIInit())
		prometheus.Init()
	}

	// Spawn CI maintenance application
	mh.has.CISpawn()

	// Initialize the xlbnet global ticker(s)
	mh.tDone = make(chan bool)
	mh.ticker = time.NewTicker(XlbnetTiVal * time.Second)
	mh.wg.Add(1)
	go xlbNetTicker()

	mh.ready = true
}

// xlbNetRun - This routine will not return
func xlbNetRun() {
	mh.wg.Wait()
}

// Main -  main routine of xlbnet
func Main() {
	xlbNetInit()
	xlbNetRun()
}

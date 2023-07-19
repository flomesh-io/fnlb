package xlbnet

import (
	"errors"
	"fmt"
	"net"
	"time"

	cmn "github.com/cybwan/fsmxlb/pkg/common"
	tk "github.com/cybwan/fsmxlb/pkg/xlblib"
)

// error codes
const (
	L2ErrBase = iota - 3000
	L2SameFdbErr
	L2OifErr
	L2NoFdbErr
	L2VxattrErr
)

// constants
const (
	FdbGts = 10
)

// FdbKey - key to find a fwd entry
type FdbKey struct {
	MacAddr  [6]byte
	BridgeID int
}

// FdbAttr - extra attribs for a fwd entry
type FdbAttr struct {
	Oif     string
	Dst     net.IP
	FdbType int
}

// FdbTunAttr - attribs for a tun fwd entry
type FdbTunAttr struct {
	rt *Rt
	nh *Neigh
	ep *NeighTunEp
}

// FdbStat - statistics for fwd entry
type FdbStat struct {
	Packets uint64
	Bytes   uint64
}

// FdbEnt - a forwarding database entry
type FdbEnt struct {
	FdbKey   FdbKey
	FdbAttr  FdbAttr
	FdbTun   FdbTunAttr
	Port     *Port
	itime    time.Time
	stime    time.Time
	unReach  bool
	inActive bool
	Sync     DpStatusT
}

// L2H - context container
type L2H struct {
	FdbMap map[FdbKey]*FdbEnt
	Zone   *Zone
}

// L2Init - Initialize the layer2 subsystem
func L2Init(z *Zone) *L2H {
	var nL2 = new(L2H)
	nL2.FdbMap = make(map[FdbKey]*FdbEnt)
	nL2.Zone = z
	z.Ports.PortNotifierRegister(nL2)

	return nL2
}

func l2FdbAttrEqual(a1 *FdbAttr, a2 *FdbAttr) bool {
	if a1.FdbType == a2.FdbType &&
		a1.Oif == a2.Oif &&
		a1.Dst.Equal(a2.Dst) {
		return true
	}
	return false
}

func l2FdbAttrCopy(dst *FdbAttr, src *FdbAttr) {
	dst.FdbType = src.FdbType
	dst.Oif = src.Oif
	dst.Dst = src.Dst
}

// L2FdbResolveNh - For TunFDB, try to associate with appropriate neighbor
func (f *FdbEnt) L2FdbResolveNh() (bool, int, error) {
	p := f.Port
	attr := f.FdbAttr
	unRch := false

	if p == nil {
		return true, L2VxattrErr, errors.New("fdb port error")
	}

	zone, _ := mh.zn.Zonefind(p.Zone)
	if zone == nil {
		return true, L2VxattrErr, errors.New("fdb zone error")
	}

	if p.SInfo.PortType&cmn.PortVxlanBr == cmn.PortVxlanBr {
		if attr.FdbType != cmn.FdbTun {
			return true, L2VxattrErr, errors.New("fdb attr error")
		}

		if attr.Dst.To4() == nil {
			return true, L2VxattrErr, errors.New("fdb v6 dst unsupported")
		}

		tk.LogIt(tk.LogDebug, "fdb tun rt lookup %s\n", attr.Dst.String())
		// Check if the end-point is reachable
		err, pDstNet, tDat := zone.Rt.Trie4.FindTrie(attr.Dst.String())
		if err == 0 && pDstNet != nil {
			switch rtn := tDat.(type) {
			case *Neigh:
				if rtn == nil {
					return true, -1, errors.New("no neigh found")
				}
			default:
				return true, -1, errors.New("no neigh found")
			}
			if nh, ok := tDat.(*Neigh); ok && !nh.Inactive {
				rt := zone.Rt.RtFind(*pDstNet, zone.Name)
				if rt == nil {
					unRch = true
					tk.LogIt(tk.LogDebug, "fdb tun rtlookup %s no-rt\n", attr.Dst.String())
				} else {
					ret, tep := zone.Nh.NeighAddTunEP(nh, attr.Dst, nil, p.HInfo.TunID, DpTunVxlan, true)
					if ret == 0 {
						rt.RtDepObjs = append(rt.RtDepObjs, f)
						f.FdbTun.rt = rt
						f.FdbTun.nh = nh
						f.FdbTun.ep = tep
						unRch = false
					} else {
						unRch = true
					}
				}
			}
		} else {
			unRch = true
			tk.LogIt(tk.LogDebug, "fdb tun rtlookup %s no trie-ent\n", attr.Dst.String())
		}
	}
	if unRch {
		tk.LogIt(tk.LogDebug, "fdb tun rtlookup %s unreachable\n", attr.Dst.String())
	}
	return unRch, 0, nil
}

// L2FdbFind - Find a fwd entry given the key
func (l2 *L2H) L2FdbFind(key FdbKey) *FdbEnt {
	fdb, found := l2.FdbMap[key]

	if found == true {
		return fdb
	}

	return nil
}

// L2FdbAdd - Add a l2 forwarding entry
func (l2 *L2H) L2FdbAdd(key FdbKey, attr FdbAttr) (int, error) {

	p := l2.Zone.Ports.PortFindByName(attr.Oif)
	if p == nil || !p.SInfo.PortActive {
		tk.LogIt(tk.LogDebug, "fdb port not found %s\n", attr.Oif)
		return L2OifErr, errors.New("no such port")
	}

	fdb, found := l2.FdbMap[key]

	if found == true {
		// Check if it is a modify
		if l2FdbAttrEqual(&attr, &fdb.FdbAttr) {
			tk.LogIt(tk.LogDebug, "fdb ent exists, %v\n", key)
			return L2SameFdbErr, errors.New("same fdb")
		}
		// Handle modify by deleting and reinstalling
		l2.L2FdbDel(key)
	}

	// Need to double check vlan associations are valid ??
	nfdb := new(FdbEnt)
	nfdb.FdbKey = key
	l2FdbAttrCopy(&nfdb.FdbAttr, &attr)
	nfdb.Port = p
	nfdb.itime = time.Now()
	nfdb.stime = time.Now()

	if p.SInfo.PortType&cmn.PortVxlanBr == cmn.PortVxlanBr {
		unRch, ret, err := nfdb.L2FdbResolveNh()
		if err != nil {
			tk.LogIt(tk.LogDebug, "tun-fdb ent resolve error, %v", key)
			return ret, err
		}
		nfdb.unReach = unRch
	}

	l2.FdbMap[nfdb.FdbKey] = nfdb

	nfdb.DP(DpCreate)

	tk.LogIt(tk.LogDebug, "added fdb ent, %v\n", key)

	return 0, nil
}

// L2FdbDel - Delete a l2 forwarding entry
func (l2 *L2H) L2FdbDel(key FdbKey) (int, error) {

	fdb, found := l2.FdbMap[key]
	if found == false {
		tk.LogIt(tk.LogDebug, "fdb ent not found, %v\n", key)
		return L2NoFdbErr, errors.New("no such fdb")
	}

	if fdb.Port.SInfo.PortType == cmn.PortVxlanBr {
		// Remove route dependencies if any
		n := 0
		if fdb.FdbTun.rt != nil {
			rt := fdb.FdbTun.rt
			for _, obj := range rt.RtDepObjs {
				if f, ok := obj.(*FdbEnt); ok {
					if f == fdb {
						rt.RtDepObjs = rt.rtRemoveDepObj(n)
						break
					}
				}
				n++
			}
		}

		fdb.FdbTun.rt = nil
		if fdb.FdbTun.nh != nil {
			fdb.FdbTun.nh.Resolved = false
			fdb.FdbTun.nh = nil
		}
		fdb.FdbTun.ep = nil
	}

	fdb.DP(DpRemove)

	fdb.inActive = true

	delete(l2.FdbMap, fdb.FdbKey)

	tk.LogIt(tk.LogDebug, "deleted fdb ent, %v\n", key)

	return 0, nil
}

// FdbTicker - Ticker routine for a fwd entry
func (l2 *L2H) FdbTicker(f *FdbEnt) {
	if time.Now().Sub(f.stime) > FdbGts {
		// This scans for inconsistencies in a fdb
		// 1. Do garbage cleaning if underlying oif or vlan is not valid anymore
		// 2. If FDB is a TunFDB, we need to make sure NH is reachable
		if f.Port.SInfo.PortActive == false {
			l2.L2FdbDel(f.FdbKey)
		} else if f.unReach == true {
			tk.LogIt(tk.LogDebug, "unrch scan - %v\n", f)
			unRch, _, _ := f.L2FdbResolveNh()
			if f.unReach != unRch {
				f.unReach = unRch
				f.DP(DpCreate)
			}
		}
		f.stime = time.Now()
	}
}

// FdbsTicker - Ticker for Fdbs
func (l2 *L2H) FdbsTicker() {
	n := 1
	for _, e := range l2.FdbMap {
		l2.FdbTicker(e)
		n++
	}
	return
}

// PortNotifier - Implementation of PortEventIntf interface
func (l2 *L2H) PortNotifier(name string, osID int, evType PortEvent) {
	if evType&PortEvDown|PortEvDelete|PortEvLowerDown != 0 {
		for _, f := range l2.FdbMap {
			if f.FdbAttr.Oif == name {
				l2.L2FdbDel(f.FdbKey)
			}
		}
	}
	return
}

func fdb2String(f *FdbEnt, it IterIntf, n *int) {
	var s string
	s = fmt.Sprintf("FdbEnt%-3d : ether %02x:%02x:%02x:%02x:%02x:%02x,br %d :: Oif %s\n",
		*n, f.FdbKey.MacAddr[0], f.FdbKey.MacAddr[1], f.FdbKey.MacAddr[2],
		f.FdbKey.MacAddr[3], f.FdbKey.MacAddr[4], f.FdbKey.MacAddr[5],
		f.FdbKey.BridgeID, f.FdbAttr.Oif)
	it.NodeWalker(s)
}

// Fdbs2String - Format all fwd entries to string
func (l2 *L2H) Fdbs2String(it IterIntf) error {
	n := 1
	for _, e := range l2.FdbMap {
		fdb2String(e, it, &n)
		n++
	}
	return nil
}

// L2DestructAll - Destructor for all layer2 fwd entries
func (l2 *L2H) L2DestructAll() {
	for _, f := range l2.FdbMap {
		l2.L2FdbDel(f.FdbKey)
	}
	return
}

// DP - Sync state of L2 entities to data-path
func (f *FdbEnt) DP(work DpWorkT) int {

	if work == DpCreate && f.unReach == true {
		return 0
	}

	l2Wq := new(L2AddrDpWorkQ)
	l2Wq.Work = work
	l2Wq.Status = &f.Sync
	if f.Port.SInfo.PortType&cmn.PortVxlanBr == cmn.PortVxlanBr {
		l2Wq.Tun = DpTunVxlan
	}

	if f.FdbTun.nh != nil {
		l2Wq.NhNum = int(f.FdbTun.nh.Mark)
	}

	for i := 0; i < 6; i++ {
		l2Wq.L2Addr[i] = uint8(f.FdbKey.MacAddr[i])
	}
	l2Wq.PortNum = f.Port.PortNo
	l2Wq.BD = f.Port.L2.Vid
	if f.Port.L2.IsPvid {
		l2Wq.Tagged = 0
	} else {
		l2Wq.Tagged = 1
		l2Wq.PortNum = f.Port.SInfo.PortReal.PortNo
	}
	mh.dp.ToDpCh <- l2Wq

	if l2Wq.Tun == DpTunVxlan {
		rmWq := new(RouterMacDpWorkQ)
		rmWq.Work = work
		rmWq.Status = nil

		if f.Port.SInfo.PortReal == nil ||
			f.FdbTun.ep == nil {
			return -1
		}

		up := f.Port.SInfo.PortReal

		for i := 0; i < 6; i++ {
			rmWq.L2Addr[i] = uint8(f.FdbKey.MacAddr[i])
		}
		rmWq.PortNum = up.PortNo
		rmWq.TunID = f.Port.HInfo.TunID
		rmWq.TunType = DpTunVxlan
		rmWq.BD = f.Port.L2.Vid
		rmWq.NhNum = int(f.FdbTun.ep.Mark)
		mh.dp.ToDpCh <- rmWq
	}

	return 0
}

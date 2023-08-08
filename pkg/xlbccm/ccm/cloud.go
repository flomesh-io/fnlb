package ccm

import (
	"context"
	"github.com/flomesh-io/fsmxlb/pkg/common"
	"github.com/flomesh-io/fsmxlb/pkg/xlbccm/client"
	"github.com/flomesh-io/fsmxlb/pkg/xlbccm/ippool"
	"io"
	"net/url"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog/v2"

	tk "github.com/flomesh-io/fsmxlb/pkg/xlblib"
)

type LBMode int32

const (
	LBModeDefault LBMode = iota
	LBModeOneArm
	LBModeFullNAT
	LBModeDSR
)

type XlbClient struct {
	providerName string

	Version        string
	APIServerURL   []*url.URL
	ExternalIPPool *ippool.IPPool
	SetBGP         bool
	SetLBMode      int32

	RESTClient *client.RESTClient
	k8sClient  kubernetes.Interface
}

// Initialize provides the cloud with a kubernetes client builder and may spawn goroutines
// to perform housekeeping or run custom controllers specific to the cloud provider.
// Any tasks started here should be cleaned up when the stop channel closes.
func (l *XlbClient) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
	l.RESTClient = client.CreateRESTClient()
	l.k8sClient = clientBuilder.ClientOrDie("fsm-ccm")

	for _, serverUrl := range l.APIServerURL {
		aliveCh := l.CreateLBHealthCheckChan(stop, serverUrl.String())
		go l.reinstallLBRules(stop, aliveCh)
	}

	// Get all loadbalancer service in all namespace
	/*
		svcList, err := l.k8sClient.CoreV1().Services("").List(context.TODO(), v1.ListOptions{})
		if err != nil {
			klog.Errorf("Failed to initialize when get k8s services. err :%s", err.Error())
			return
		}

		klog.Infof("XlbClient.Initialize: ")
		for _, svc := range svcList.Items {
			if svc.Spec.Type != "LoadBalancer" {
				continue
			}
			klog.Infof("type LoadBalancer service name: %s, namespace: %s", svc.Name, svc.Namespace)
		}
	*/
}

// LoadBalancer returns a balancer interface. Also returns true if the interface is supported, false otherwise.
func (l *XlbClient) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return l, true
}

// Instances returns an instances interface. Also returns true if the interface is supported, false otherwise.
func (l *XlbClient) Instances() (cloudprovider.Instances, bool) {
	return nil, false
}

// InstancesV2 is an implementation for instances and should only be implemented by external cloud providers.
// Implementing InstancesV2 is behaviorally identical to Instances but is optimized to significantly reduce
// API calls to the cloud provider when registering and syncing nodes. Implementation of this interface will
// disable calls to the Zones interface. Also returns true if the interface is supported, false otherwise.
func (l *XlbClient) InstancesV2() (cloudprovider.InstancesV2, bool) {
	return nil, false
}

// Zones returns a zones interface. Also returns true if the interface is supported, false otherwise.
// DEPRECATED: Zones is deprecated in favor of retrieving zone/region information from InstancesV2.
// This interface will not be called if InstancesV2 is enabled.
func (l *XlbClient) Zones() (cloudprovider.Zones, bool) {
	return nil, false
}

// Clusters returns a clusters interface.  Also returns true if the interface is supported, false otherwise.
func (l *XlbClient) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

// Routes returns a routes interface along with whether the interface is supported.
func (l *XlbClient) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

// ProviderName returns the cloud provider ID.
func (l *XlbClient) ProviderName() string {
	klog.V(5).Infof("XlbClient.ProviderName() returned %s", l.providerName)
	return l.providerName
}

// HasClusterID returns true if a ClusterID is required and set
func (l *XlbClient) HasClusterID() bool {
	klog.V(5).Info("XlbClient.HasClusterID() returned true")
	return true
}

// io.Reader contains the contents of the CCM config file. (--cloud-config cmdline options)
func init() {
	cloudprovider.RegisterCloudProvider(common.ProviderName, func(_ io.Reader) (cloudprovider.Interface, error) {
		o, err := ReadCCMEnvronment()
		if err != nil {
			klog.Errorf("fsm-ccm: failed to get environment")
			return nil, err
		}

		ipPool, err := ippool.NewIPPool(tk.IpAllocatorNew(), o.ExternalCIDR, !o.ExclIPAM)
		if err != nil {
			klog.Errorf("fsm-ccm: failed to create external IP Pool (CIDR: %s)", o.ExternalCIDR)
			return nil, err
		}

		return &XlbClient{
			providerName:   common.ProviderName,
			Version:        "v1",
			APIServerURL:   o.APIServerUrlList,
			ExternalIPPool: ipPool,
			SetBGP:         o.SetBGP,
			SetLBMode:      o.SetLBMode,
		}, nil
	})
}

func (l *XlbClient) CreateLBHealthCheckChan(stop <-chan struct{}, apiUrl string) chan string {
	aliveCh := make(chan string)
	isAlive := true

	go wait.Until(func() {
		if err := l.LBHealthCheck(apiUrl); err != nil {
			if isAlive {
				klog.Infof("CreateLBHealthCheckChan: fsmxlb(%s) is down. isAlive is changed to 'false'", apiUrl)
				isAlive = false
			}
		} else {
			if !isAlive {
				klog.Infof("CreateLBHealthCheckChan: fsmxlb(%s) is alive again. isAlive is set 'true'", apiUrl)
				isAlive = true
				aliveCh <- apiUrl
			}
		}
	}, time.Second*2, stop)

	return aliveCh
}

func (l *XlbClient) LBHealthCheck(apiUrl string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	resp, err := l.RESTClient.GET(ctx, apiUrl)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}

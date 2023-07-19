package ccm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

type LoadBalancerModel struct {
	Service   LoadBalancerService    `json:"serviceArguments"`
	Endpoints []LoadBalancerEndpoint `json:"endpoints"`
}

type LoadBalancerListModel struct {
	LoadBalancerList []LoadBalancerModel `json:"lbAttr"`
}

type LoadBalancerService struct {
	ExternalIP string   `json:"externalIP"`
	Port       int32    `json:"port"`
	Protocol   string   `json:"protocol"`
	Bgp        bool     `json:"bgp"`
	Sel        EpSelect `json:"sel"`
	Mode       int32    `json:"mode"`
}

type EpSelect uint

type LoadBalancerEndpoint struct {
	EndpointIP string `json:"endpointIP"`
	TargetPort int32  `json:"targetPort"`
	Weight     int8   `json:"weight"`
}

type SvcPair struct {
	IPString string
	Port     int32
	Protocol string
}

const (
	LoadBalancerResource = "config/loadbalancer"
	MaxWeight            = 10
)

func (l *XlbClient) GetLoadBalancerAPIUrlString(serverURL *url.URL, subResource []string) string {
	p := path.Join(l.providerName, l.Version, LoadBalancerResource)
	if len(subResource) > 0 {
		subPath := path.Join(subResource...)
		p = path.Join(p, subPath)
	}

	lbURL := url.URL{
		Scheme: serverURL.Scheme,
		Host:   serverURL.Host,
		Path:   p,
	}

	return lbURL.String()
}

func (l *XlbClient) UpdateQueryToUrl(urlStr string, query map[string]string) (string, error) {
	url, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	q := url.Query()
	for key, value := range query {
		q.Add(key, value)
	}

	url.RawQuery = q.Encode()
	return url.String(), nil
}

// Implementations must treat the *v1.Service parameter as read-only and not modify it.
// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
func (l *XlbClient) GetLoadBalancer(ctx context.Context, clusterName string, service *v1.Service) (*v1.LoadBalancerStatus, bool, error) {
	var resp *http.Response
	var err error

	subResource := []string{
		"all",
	}
	var fsmGetLoadBalancerURLs []string
	for _, u := range l.APIServerURL {
		fsmGetLoadBalancerURLs = append(fsmGetLoadBalancerURLs, l.GetLoadBalancerAPIUrlString(u, subResource))
	}

	ingresses := service.Status.LoadBalancer.Ingress
	for _, fsmGetLoadBalancerURL := range fsmGetLoadBalancerURLs {
		resp, err = l.RESTClient.GET(ctx, fsmGetLoadBalancerURL)
		if err != nil {
			continue
		}

		defer resp.Body.Close()
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			klog.Errorf("failed to read fsmxlb(%s) response body in LoadBalancer.GetLoadBalancer(). err: %s", fsmGetLoadBalancerURL, err.Error())
			continue
		}
		if resp.StatusCode != http.StatusOK {
			klog.Errorf("GetLoadBalancer: fsmxlb(%s) return response status: %d, msg: %v", resp.StatusCode, respBody)

			continue
		}

		lbListModel := LoadBalancerListModel{}
		if err := json.Unmarshal(respBody, &lbListModel); err != nil {
			klog.Errorf("failed to unmarshal response body in LoadBalancer.GetLoadBalancer()")
			return nil, false, err
		}

		for _, lbModel := range lbListModel.LoadBalancerList {
			for _, ingress := range ingresses {
				klog.Infof("  ### Service.LoadBalancer.Ingress: %s == fsmxlb.ExternalIP: %s", ingress.IP, lbModel.Service.ExternalIP)
				if lbModel.Service.ExternalIP == ingress.IP {
					status := &v1.LoadBalancerStatus{}
					status.Ingress = []v1.LoadBalancerIngress{{IP: lbModel.Service.ExternalIP}}
					status.Ingress[0].Ports = []v1.PortStatus{{Port: lbModel.Service.Port, Protocol: v1.Protocol(strings.ToUpper(lbModel.Service.Protocol))}}
					return status, true, nil
				}
			}
		}
	}

	klog.Infof("not found Load Balancer (Ingresses: %v)", ingresses)
	return nil, false, errors.New("Not found")
}

// GetLoadBalancerName returns the name of the load balancer. Implementations must treat the
// *v1.Service parameter as read-only and not modify it.
func (l *XlbClient) GetLoadBalancerName(ctx context.Context, clusterName string, service *v1.Service) string {
	klog.Infof("LoadBalancer.GetLoadBalancerName() returned v1.service.Name: %s", service.Name)
	return service.Name
}

// EnsureLoadBalancer creates a new load balancer 'name', or updates the existing one. Returns the status of the balancer
// Implementations must treat the *v1.Service and *v1.Node
// parameters as read-only and not modify them.
// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
func (l *XlbClient) EnsureLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) (*v1.LoadBalancerStatus, error) {
	klog.Infof("LoadBalancer.EnsureLoadBalancer() called. service: %s", service.Name)
	if l.SvcHasLBClass(*service) {
		klog.Infof("service %s has Spec.LoadBalancerClass %s. ignore.", service.Name, *service.Spec.LoadBalancerClass)
		return nil, nil
	}

	// Get endpoint IP/port pairs for this service
	epPairs := l.getEndpointsForLB(nodes, service)

	// validation check if service has ingress IP already
	ingSvcPairs, err := l.getIngressSvcPairs(service)
	if err != nil {
		return nil, err
	}

	status := &v1.LoadBalancerStatus{}

	// set defer for deallocating IP on error
	isFailed := false
	defer func() {
		if isFailed {
			klog.Infof("deallocateOnFailure defer function called")
			for _, sp := range ingSvcPairs {
				klog.Infof("ip %s is newIP so retrieve pool", sp.IPString)
				l.ExternalIPPool.ReturnIPAddr(sp.IPString, uint32(sp.Port), string(sp.Protocol))
			}
		}
	}()

	var fsmCreateLoadBalancerURLs []string
	for _, u := range l.APIServerURL {
		fsmCreateLoadBalancerURLs = append(fsmCreateLoadBalancerURLs, l.GetLoadBalancerAPIUrlString(u, nil))
	}
	for _, ingSvcPair := range ingSvcPairs {
		var errChList []chan error

		for _, fsmCreateLoadBalancerURL := range fsmCreateLoadBalancerURLs {
			ch := make(chan error)

			go func(urlStr string, ch chan error) {
				ch <- l.addLoadBalancerRule(ctx, urlStr, ingSvcPair, service, epPairs)
			}(fsmCreateLoadBalancerURL, ch)

			errChList = append(errChList, ch)
		}

		isError := true
		for _, errCh := range errChList {
			err := <-errCh
			if err == nil {
				isError = false
			}
		}
		if isError {
			isFailed = isError
			return nil, fmt.Errorf("failed to add fsmxlb loadBalancer")
		}

		retIngress := v1.LoadBalancerIngress{IP: ingSvcPair.IPString}
		retIngress.Ports = append(retIngress.Ports, v1.PortStatus{Port: ingSvcPair.Port, Protocol: v1.Protocol(strings.ToUpper(ingSvcPair.Protocol))})
		status.Ingress = append(status.Ingress, retIngress)
	}

	return status, nil
}

// UpdateLoadBalancer updates hosts under the specified load balancer.
// Implementations must treat the *v1.Service and *v1.Node
// parameters as read-only and not modify them.
// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
func (l *XlbClient) UpdateLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) error {
	klog.Infof("LoadBalancer.UpdateLoadBalancer() called by service %s", service.Name)
	return nil
}

// EnsureLoadBalancerDeleted deletes the specified load balancer if it
// exists, returning nil if the load balancer specified either didn't exist or
// was successfully deleted.
// This construction is useful because many cloud providers' load balancers
// have multiple underlying components, meaning a Get could say that the LB
// doesn't exist even if some part of it is still laying around.
// Implementations must treat the *v1.Service parameter as read-only and not modify it.
// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
func (l *XlbClient) EnsureLoadBalancerDeleted(ctx context.Context, clusterName string, service *v1.Service) error {
	if l.SvcHasLBClass(*service) {
		klog.Infof("service %s has Spec.LoadBalancerClass %s. ignore.", service.Name, *service.Spec.LoadBalancerClass)
		return nil
	}

	ingresses := service.Status.LoadBalancer.Ingress
	ports := service.Spec.Ports
	for _, ingress := range ingresses {
		for _, port := range ports {
			// TODO: need to change argument?
			var fsmDeleteLoadBalancerURLs []string
			for _, u := range l.APIServerURL {
				fsmDeleteLoadBalancerURLs = append(fsmDeleteLoadBalancerURLs, l.GetLoadBalancerAPIUrlString(u, []string{
					"externalipaddress", ingress.IP,
					"port", strconv.Itoa(int(port.Port)),
					"protocol", strings.ToLower(string(port.Protocol)),
				}))
			}

			var errChList []chan error
			for _, fsmDeleteLoadBalancerURL := range fsmDeleteLoadBalancerURLs {
				fsmDeleteLoadBalancerURLAndQuery, err := l.UpdateQueryToUrl(fsmDeleteLoadBalancerURL, map[string]string{
					"bgp": strconv.FormatBool(l.SetBGP),
				})
				if err != nil {
					klog.Errorf("URL (%s) is incorrect. err: %s", fsmDeleteLoadBalancerURL, err.Error())
				}
				klog.Infof("EnsureLoadBalancerDeleted(): fsmDeleteLoadBalancerURLAndQuery: %s", fsmDeleteLoadBalancerURLAndQuery)

				ch := make(chan error)
				go func(urlStr string, ch chan error) {
					resp, err := l.RESTClient.DELETE(ctx, urlStr)
					// TODO:
					if err != nil {
						klog.Errorf("failed to call fsmxlb API. err: %s", err.Error())
						ch <- err
						return
					}
					defer resp.Body.Close()
					if resp.StatusCode != http.StatusOK {
						respBody, _ := io.ReadAll(resp.Body)
						klog.Errorf("failed to delete Load Balancer (Ingress: %s). fsmxlb %s return response code %d. message: %v", ingress.IP, urlStr, resp.StatusCode, respBody)
						ch <- fmt.Errorf("fsmxlb %s return response code %d. message: %v", urlStr, resp.StatusCode, respBody)
						return
					}
					ch <- nil
				}(fsmDeleteLoadBalancerURLAndQuery, ch)

				errChList = append(errChList, ch)
			}

			isError := true
			for _, errCh := range errChList {
				err := <-errCh
				if err == nil {
					isError = false
					break
				}
			}
			if isError {
				return fmt.Errorf("failed to delete fsmxlb LoadBalancer")
			}
			l.ExternalIPPool.ReturnIPAddr(ingress.IP, uint32(port.Port), string(port.Protocol))
		}
	}
	return nil
}

func (l *XlbClient) getLBIngressSvcPairs(service *v1.Service) []SvcPair {
	var spairs []SvcPair
	for _, ingress := range service.Status.LoadBalancer.Ingress {
		for _, port := range service.Spec.Ports {
			sp := SvcPair{ingress.IP, port.Port, strings.ToLower(string(port.Protocol))}
			spairs = append(spairs, sp)
		}
	}

	return spairs
}

func (l *XlbClient) makeLoadBalancerModel(externalIP string, port v1.ServicePort, epPairs []SvcPair) LoadBalancerModel {
	fsmEndpointModelList := []LoadBalancerEndpoint{}

	if len(epPairs) > 0 {
		endpointWeight := int8(MaxWeight / len(epPairs))
		remainderWeight := int8(MaxWeight % len(epPairs))

		for _, endpoint := range epPairs {
			weight := endpointWeight
			if remainderWeight > 0 {
				weight++
				remainderWeight--
			}

			fsmEndpointModelList = append(fsmEndpointModelList, LoadBalancerEndpoint{
				EndpointIP: endpoint.IPString,
				TargetPort: port.NodePort,
				Weight:     weight,
			})
		}
	}

	return LoadBalancerModel{
		Service: LoadBalancerService{
			ExternalIP: externalIP,
			Port:       port.Port,
			Protocol:   strings.ToLower(string(port.Protocol)),
			Bgp:        l.SetBGP,
			Mode:       l.SetLBMode,
		},
		Endpoints: fsmEndpointModelList,
	}
}

func (l *XlbClient) addLoadBalancerRule(ctx context.Context, lbUrl string, sPair SvcPair, service *v1.Service, epPairs []SvcPair) error {
	for _, port := range service.Spec.Ports {
		lbModel := l.makeLoadBalancerModel(sPair.IPString, port, epPairs)
		body, err := json.Marshal(lbModel)
		if err != nil {
			klog.Errorf("failed to EnsureLoadBalancer(). err: %s", err.Error())
			return err
		}

		resp, err := l.RESTClient.POST(ctx, lbUrl, body)
		if err != nil {
			klog.Errorf("failed to addLoadBalancerRule() call to fsmxlb(%s) API. err: %s", lbUrl, err.Error())
			return err
		}

		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			klog.Errorf("failed to addLoadBalancerRule(): fsmxlb API(%s) response status %d.", lbUrl, resp.StatusCode)
			return fmt.Errorf("fsmxlb(%s) response status %d", resp.Request.URL.String(), resp.StatusCode)
		}
	}

	return nil
}

func (l *XlbClient) getEndpointsForLB(nodes []*v1.Node, service *v1.Service) []SvcPair {
	var epPairs []SvcPair
	for _, node := range nodes {
		addr, err := l.getNodeAddress(*node)
		if err != nil {
			klog.Errorf(err.Error())
			continue
		}
		for _, port := range service.Spec.Ports {
			ep := SvcPair{addr, port.NodePort, strings.ToLower(string(port.Protocol))}
			epPairs = append(epPairs, ep)
		}
	}

	return epPairs
}

func (l *XlbClient) getNodeAddress(node v1.Node) (string, error) {
	addrs := node.Status.Addresses
	if len(addrs) == 0 {
		return "", errors.New("no address found for host")
	}

	for _, addr := range addrs {
		if addr.Type == v1.NodeInternalIP {
			return addr.Address, nil
		}
	}

	return addrs[0].Address, nil
}

func (l *XlbClient) reinstallLBRules(stopCh <-chan struct{}, aliveCh <-chan string) {
loop:
	for {
		select {
		case <-stopCh:
			break loop
		case aliveUrl := <-aliveCh:
			isSuccess := false
			for retry := 0; retry < 5; retry++ {
				klog.Infof("try reinstall LB rule...")
				if err := l.tryReinstallLBRules(aliveUrl); err == nil {
					isSuccess = true
					break
				} else {
					time.Sleep(1 * time.Second)
				}
			}
			if !isSuccess {
				klog.Exit("restart fsm-ccm")
			}
		}
	}
}

func (l *XlbClient) tryReinstallLBRules(apiUrlStr string) error {
	klog.Infof("fsmxlb alive again so reinstall all LB rules")

	services, err := l.k8sClient.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		klog.Errorf("failed to get k8s service list when reinstall LB. err: %v", err)
		return err
	}
	nodes, err := l.k8sClient.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{LabelSelector: "!node.kubernetes.io/exclude-from-external-load-balancers"})
	if err != nil {
		klog.Errorf("failed to get k8s nodes list when reinstall LB. err: %v", err)
		return err
	}

	apiUrl, _ := url.Parse(apiUrlStr)
	lbUrl := l.GetLoadBalancerAPIUrlString(apiUrl, nil)
	for _, svc := range services.Items {
		if svc.Spec.Type != v1.ServiceTypeLoadBalancer {
			continue
		}
		if l.SvcHasLBClass(svc) {
			continue
		}
		ingSvcPairs := l.getLBIngressSvcPairs(&svc)
		var epPairs []SvcPair
		for _, node := range nodes.Items {
			nodeIP, err := l.getNodeAddress(node)
			if err != nil {
				klog.Errorf("reinstallLBRules: failed to get nodeAddress of %s.", node.Name)
				continue
			}

			for _, port := range svc.Spec.Ports {
				ep := SvcPair{nodeIP, port.NodePort, strings.ToLower(string(port.Protocol))}
				epPairs = append(epPairs, ep)
			}
		}

		for _, ingSvcPair := range ingSvcPairs {
			if err := l.addLoadBalancerRule(context.TODO(), lbUrl, ingSvcPair, &svc, epPairs); err != nil {
				return err
			}
		}
	}

	return nil
}

// getIngressSvcPairs check validation if service have ingress IP already.
// If service have no ingress IP, assign new IP in IP pool
func (l *XlbClient) getIngressSvcPairs(service *v1.Service) ([]SvcPair, error) {
	var sPairs []SvcPair
	inSPairs := l.getLBIngressSvcPairs(service)
	if len(inSPairs) >= 1 {
		for _, inSPair := range inSPairs {
			ident := inSPair.Port
			klog.Infof("ingress service exists")

			if l.ExternalIPPool.CheckAndReserveIP(inSPair.IPString, uint32(ident), inSPair.Protocol) {
				sp := SvcPair{inSPair.IPString, ident, inSPair.Protocol}
				sPairs = append(sPairs, sp)
			} else {
				newIP := l.ExternalIPPool.GetNewIPAddr(uint32(ident), inSPair.Protocol)
				if newIP == nil {
					klog.Errorf("failed to generate external IP. IP Pool is full")
					return nil, errors.New("failed to generate external IP. IP Pool is full")
				}

				sp := SvcPair{newIP.String(), ident, inSPair.Protocol}
				sPairs = append(sPairs, sp)
			}
		}
	} else {
		for _, port := range service.Spec.Ports {
			newIP := l.ExternalIPPool.GetNewIPAddr(uint32(port.Port), string(port.Protocol))
			if newIP == nil {
				klog.Errorf("failed to generate external IP. IP Pool is full")
				return nil, errors.New("failed to generate external IP. IP Pool is full")
			}
			sp := SvcPair{newIP.String(), port.Port, strings.ToLower(string(port.Protocol))}
			sPairs = append(sPairs, sp)
		}
	}

	return sPairs, nil
}

func (l *XlbClient) SvcHasLBClass(service v1.Service) bool {
	return service.Spec.LoadBalancerClass != nil
}

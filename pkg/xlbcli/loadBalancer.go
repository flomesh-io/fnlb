package xlbcli

import (
	"fmt"
	"sort"
)

type LoadBalancer struct {
	CommonAPI
}

type EpSelect uint
type LbMode int32
type LbRuleModGet struct {
	LbRules []LoadBalancerModel `json:"lbAttr"`
}

type LoadBalancerModel struct {
	Service      LoadBalancerService    `json:"serviceArguments" yaml:"serviceArguments"`
	SecondaryIPs []LoadBalancerSecIp    `json:"secondaryIPs" yaml:"secondaryIPs"`
	Endpoints    []LoadBalancerEndpoint `json:"endpoints" yaml:"endpoints"`
}

type LoadBalancerService struct {
	ExternalIP string   `json:"externalIP" yaml:"externalIP"`
	Port       uint16   `json:"port"           yaml:"port" `
	Protocol   string   `json:"protocol"       yaml:"protocol"`
	Sel        EpSelect `json:"sel"            yaml:"sel"`
	Mode       LbMode   `json:"mode"           yaml:"mode"`
	BGP        bool     `json:"BGP"            yaml:"BGP"`
	Monitor    bool     `json:"Monitor"        yaml:"Monitor"`
	Timeout    uint32   `json:"inactiveTimeOut" yaml:"inactiveTimeOut"`
	Block      uint16   `json:"block"          yaml:"block"`
	Managed    bool     `json:"managed,omitempty" yaml:"managed"`
}

type LoadBalancerEndpoint struct {
	EndpointIP string `json:"endpointIP" yaml:"endpointIP"`
	TargetPort uint16 `json:"targetPort" yaml:"targetPort"`
	Weight     uint8  `json:"weight"     yaml:"weight"`
	State      string `json:"state"      yaml:"state"`
}

type LoadBalancerSecIp struct {
	SecondaryIP string `json:"secondaryIP" yaml:"secondaryIP"`
}

type ConfigurationLBFile struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
	Spec       LoadBalancerModel `yaml:"spec"`
}

func (service LoadBalancerService) Key() string {
	return fmt.Sprintf("%s|%05d|%s", service.ExternalIP, service.Port, service.Protocol)
}

func (lbresp LbRuleModGet) Sort() {
	sort.Slice(lbresp.LbRules, func(i, j int) bool {
		return lbresp.LbRules[i].Service.Key() < lbresp.LbRules[j].Service.Key()
	})
}

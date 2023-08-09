package xlbcli

import (
	"github.com/flomesh-io/fnlb/pkg/common"
	"net/http"
	"time"
)

const (
	xlbApiVersion              = "v1"
	xlbLoadBalancerResource    = "config/loadbalancer"
	xlbLoadBalancerResourceAll = "config/loadbalancer/all"
	xlbConntrackResource       = "config/conntrack/all"
	xlbPortResource            = "config/port/all"
	xlbSessionResource         = "config/session"
	xlbSessionUlClResource     = "config/sessionulcl"
	xlbPolicyResource          = "config/policy"
	xlbRouteResource           = "config/route"
	xlbIPv4AddressResource     = "config/ipv4address"
	xlbNeighborResource        = "config/neighbor"
	xlbFDBResource             = "config/fdb"
	xlbVlanResource            = "config/vlan"
	xlbVxlanResource           = "config/tunnel/vxlan"
	xlbMirrorResource          = "config/mirror"
	xlbFirewallResource        = "config/firewall"
	xlbEndPointResource        = "config/endpoint"
	xlbParamResource           = "config/params"
	xlbStatusResource          = "status"
)

type XlbClient struct {
	restClient RESTClient
}

func NewXlbClient(o *RESTOptions) *XlbClient {
	return &XlbClient{
		restClient: RESTClient{
			Options: *o,
			Client: &http.Client{
				Timeout: time.Second * time.Duration(o.Timeout),
			},
		},
	}
}

func (l *XlbClient) LoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbLoadBalancerResource,
			},
		},
	}
}

func (l *XlbClient) LoadBalancerAll() *LoadBalancer {
	return &LoadBalancer{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbLoadBalancerResourceAll,
			},
		},
	}
}

func (l *XlbClient) Conntrack() *Conntrack {
	return &Conntrack{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbConntrackResource,
			},
		},
	}
}

func (l *XlbClient) Port() *Port {
	return &Port{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbPortResource,
			},
		},
	}
}

func (l *XlbClient) Session() *Session {
	return &Session{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbSessionResource,
			},
		},
	}
}

func (l *XlbClient) SessionUlCL() *SessionUlCl {
	return &SessionUlCl{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbSessionUlClResource,
			},
		},
	}
}

func (l *XlbClient) Policy() *Policy {
	return &Policy{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbPolicyResource,
			},
		},
	}
}

func (l *XlbClient) Route() *Route {
	return &Route{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbRouteResource,
			},
		},
	}
}

func (l *XlbClient) IPv4Address() *IPv4Address {
	return &IPv4Address{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbIPv4AddressResource,
			},
		},
	}
}

func (l *XlbClient) Neighbor() *Neighbor {
	return &Neighbor{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbNeighborResource,
			},
		},
	}
}

func (l *XlbClient) FDB() *FDB {
	return &FDB{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbFDBResource,
			},
		},
	}
}

func (l *XlbClient) Vlan() *Vlan {
	return &Vlan{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbVlanResource,
			},
		},
	}
}

func (l *XlbClient) Vxlan() *Vxlan {
	return &Vxlan{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbVxlanResource,
			},
		},
	}
}

func (l *XlbClient) Status() *Status {
	return &Status{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbStatusResource,
			},
		},
	}
}

func (l *XlbClient) Firewall() *Firewall {
	return &Firewall{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbFirewallResource,
			},
		},
	}
}

func (l *XlbClient) Mirror() *Mirror {
	return &Mirror{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbMirrorResource,
			},
		},
	}
}

func (l *XlbClient) EndPoint() *Firewall {
	return &Firewall{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbEndPointResource,
			},
		},
	}
}

func (l *XlbClient) Param() *Param {
	return &Param{
		CommonAPI: CommonAPI{
			restClient: &l.restClient,
			requestInfo: RequestInfo{
				provider:   common.ProviderName,
				apiVersion: xlbApiVersion,
				resource:   xlbParamResource,
			},
		},
	}
}

package xlbcli

import "sort"

type Route struct {
	CommonAPI
}

type RouteModGet struct {
	RouteAttr []Routev4Get `json:"routeAttr"`
}

// RouteGetEntryStatistic - Info about an route statistic
type RouteGetEntryStatistic struct {
	// Statistic of the ingress port bytes.
	Bytes int `json:"bytes"`
	// Statistic of the egress port bytes.
	Packets int `json:"packets"`
}

// Routev4Get - Info about an route
type Routev4Get struct {
	// Flags - flag type
	Flags string `json:"flags" yaml:"flags"`
	// Gw - gateway information if any
	Gw string `json:"gateway" yaml:"gateway"`
	// Dst - ip addr
	Dst string `json:"destinationIPNet" yaml:"destinationIPNet"`
	// index of the route
	HardwareMark int `json:"hardwareMark" yaml:"hardwareMark"`
	// statistic
	Statistic RouteGetEntryStatistic `json:"statistic" yaml:"statistic"`
}

type ConfigurationRouteFile struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
	Spec       Routev4Get `yaml:"spec"`
}

func (routeresp RouteModGet) Sort() {
	sort.Slice(routeresp.RouteAttr, func(i, j int) bool {
		return routeresp.RouteAttr[i].Dst < routeresp.RouteAttr[j].Dst
	})
}

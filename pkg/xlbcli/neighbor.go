package xlbcli

import (
	"fmt"
	"sort"
)

type Neighbor struct {
	CommonAPI
}

type NeighborModGet struct {
	NeighborAttr []NeighborMod `json:"neighborAttr"`
}

type NeighborMod struct {
	// Dev - name of the related device
	Dev string `json:"dev" yaml:"dev"`
	// IP - Actual IP address
	IP string `json:"ipAddress" yaml:"ipAddress"`
	// MacAddress - Hardware address
	MacAddress string `json:"macAddress" yaml:"macAddress"`
}
type ConfigurationNeighborFile struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
	Spec       NeighborMod `yaml:"spec"`
}

func (nei NeighborMod) Key() string {
	return fmt.Sprintf("%s|%s|%s", nei.IP, nei.Dev, nei.MacAddress)
}

func (Neighborsresp NeighborModGet) Sort() {
	sort.Slice(Neighborsresp.NeighborAttr, func(i, j int) bool {
		return Neighborsresp.NeighborAttr[i].Key() < Neighborsresp.NeighborAttr[j].Key()
	})
}

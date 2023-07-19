package xlbcli

import (
	"fmt"
	"sort"
)

type IPv4Address struct {
	CommonAPI
}

type Ipv4AddrModGet struct {
	IPv4Attr []Ipv4AddrGet `json:"ipAttr"`
}

// Ipv4AddrGet - Info about an ip addresses
type Ipv4AddrGet struct {
	// Dev - name of the related device
	Dev string `json:"dev"`
	// IP - Actual IP address
	IP []string `json:"ipAddress"`
	// Sync - sync state
	Sync DpStatusT `json:"sync"`
}

// Ipv4AddrGet - Info about an ip addresses
type Ipv4AddrMod struct {
	// Dev - name of the related device
	Dev string `json:"dev" yaml:"dev"`
	// IP - Actual IP address
	IP string `json:"ipAddress" yaml:"ipAddress"`
}
type ConfigurationIPv4File struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
	Spec       Ipv4AddrMod `yaml:"spec"`
}

func (ipaddr Ipv4AddrGet) Key() string {
	return fmt.Sprintf("%s|%s", ipaddr.Dev, ipaddr.IP)
}

func (IPv4Addressresp Ipv4AddrModGet) Sort() {
	sort.Slice(IPv4Addressresp.IPv4Attr, func(i, j int) bool {
		return IPv4Addressresp.IPv4Attr[i].Key() < IPv4Addressresp.IPv4Attr[j].Key()
	})
}

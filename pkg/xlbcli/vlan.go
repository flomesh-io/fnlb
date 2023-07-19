package xlbcli

import "sort"

type Vlan struct {
	CommonAPI
}
type VlanGet struct {
	Vlans []VlanDump `json:"vlanAttr"`
}

type VlanDump struct {
	Vid       int             `json:"vid"`
	Dev       string          `json:"dev"`
	Member    []VlanMemberMod `json:"member"`
	Statistic VlanStat        `json:"vlanStatistic"`
}

// vlanStat - statistics for vlan interface
type VlanStat struct {
	InBytes    uint64
	InPackets  uint64
	OutBytes   uint64
	OutPackets uint64
}

// VlanBridgerMod - Info about an Vlan bridge
type VlanBridgeMod struct {
	// Vid - Virtual LAN ID
	Vid int `json:"vid" yaml:"vid"`
}

// VlanMemberMod - Info about an Vlan bridge member
type VlanMemberMod struct {
	// Dev - name of the related device
	Dev string `json:"dev" yaml:"dev"`
	// Tagged - Tagging status of the device
	Tagged bool `json:"tagged" yaml:"tagged"`
}
type ConfigurationVlanFile struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
	Spec       VlanBridgeMod `yaml:"spec"`
}

type ConfigurationVlanMemberFile struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
	Spec       VlanMemberMod `yaml:"spec"`
}

func (Vlanresp VlanGet) Sort() {
	sort.Slice(Vlanresp.Vlans, func(i, j int) bool {
		return Vlanresp.Vlans[i].Vid < Vlanresp.Vlans[j].Vid
	})
}

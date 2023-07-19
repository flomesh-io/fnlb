package xlbcli

import "sort"

type Policy struct {
	CommonAPI
}

type PolObjType uint

type PolInformationGet struct {
	PolModInfo []PolMod `json:"polAttr"`
}

type PolInfo struct {
	PolType           int    `json:"type" yaml:"type"`
	ColorAware        bool   `json:"colorAware" yaml:"colorAware"`
	CommittedInfoRate uint64 `json:"committedInfoRate" yaml:"committedInfoRate"` // CIR in Mbps
	PeakInfoRate      uint64 `json:"peakInfoRate" yaml:"peakInfoRate"`           // PIR in Mbps
	CommittedBlkSize  uint64 `json:"committedBlkSize" yaml:"committedBlkSize"`   // CBS in bytes
	ExcessBlkSize     uint64 `json:"excessBlkSize" yaml:"excessBlkSize"`         // EBS in bytes
}

type PolObj struct {
	PolObjName string     `json:"polObjName" yaml:"polObjName"`
	AttachMent PolObjType `json:"attachment" yaml:"attachment"`
}

type PolMod struct {
	Ident  string  `json:"policyIdent" yaml:"policyIdent"`
	Info   PolInfo `json:"policyInfo" yaml:"policyInfo"`
	Target PolObj  `json:"targetObject" yaml:"targetObject"`
}
type ConfigurationPolicyFile struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
	Spec       PolMod `yaml:"spec"`
}

func (Polresp PolInformationGet) Sort() {
	sort.Slice(Polresp.PolModInfo, func(i, j int) bool {
		return Polresp.PolModInfo[i].Ident < Polresp.PolModInfo[j].Ident
	})
}

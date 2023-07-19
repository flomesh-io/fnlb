package xlbcli

import (
	"fmt"
	"sort"
)

type Firewall struct {
	CommonAPI
}

type FWInformationGet struct {
	FWInfo []FwRuleMod `json:"fwAttr"`
}

// FwRuleOpts - Information related to Firewall options
type FwOptArg struct {
	// Drop - Drop any matching rule
	Drop bool `json:"drop" yaml:"drop"`
	// Trap - Trap anything matching rule
	Trap bool `json:"trap" yaml:"trap"`
	// Redirect - Redirect any matching rule
	Rdr     bool   `json:"redirect" yaml:"redirect"`
	RdrPort string `json:"redirectPortName" yaml:"redirectPortName"`
	// Allow - Allow any matching rule
	Allow bool `json:"allow" yaml:"allow"`
	Mark  int  `json:"fwMark" yaml:"fwMark"`
	// Record - Record packets matching rule
	Record bool `json:"record" yaml:"record"`
}

// FwRuleArg - Information related to firewall rule
type FwRuleArg struct {
	// SrcIP - Source IP in CIDR notation
	SrcIP string `json:"sourceIP" yaml:"sourceIP"`
	// DstIP - Destination IP in CIDR notation
	DstIP string `json:"destinationIP" yaml:"destinationIP"`
	// SrcPortMin - Minimum source port range
	SrcPortMin uint16 `json:"minSourcePort" yaml:"minSourcePort"`
	// SrcPortMax - Maximum source port range
	SrcPortMax uint16 `json:"maxSourcePort" yaml:"maxSourcePort"`
	// DstPortMin - Minimum destination port range
	DstPortMin uint16 `json:"minDestinationPort" yaml:"minDestinationPort"`
	// SrcPortMax - Maximum source port range
	DstPortMax uint16 `json:"maxDestinationPort" yaml:"maxDestinationPort"`
	// Proto - the protocol
	Proto uint8 `json:"protocol" yaml:"protocol"`
	// InPort - the incoming port
	InPort string `json:"portName" yaml:"portName"`
	// Pref - User preference for ordering
	Pref uint16 `json:"preference" yaml:"preference"`
}

// FwRuleMod - Info related to a firewall entry
type FwRuleMod struct {
	// Serv - service argument of type FwRuleArg
	Rule FwRuleArg `json:"ruleArguments" yaml:"ruleArguments"`
	// Opts - firewall options
	Opts FwOptArg `json:"opts" yaml:"opts"`
}

type ConfigurationFWFile struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
	Spec       FwRuleMod `yaml:"spec"`
}

func (fw FwRuleArg) Key() string {
	return fmt.Sprintf("%s|%s|%05d|%05d|%05d|%05d|%d",
		fw.SrcIP, fw.DstIP, fw.SrcPortMin, fw.SrcPortMax,
		fw.DstPortMin, fw.DstPortMax, fw.Proto)
}

func (fwresp FWInformationGet) Sort() {
	sort.Slice(fwresp.FWInfo, func(i, j int) bool {
		return fwresp.FWInfo[i].Rule.Key() < fwresp.FWInfo[j].Rule.Key()
	})
}

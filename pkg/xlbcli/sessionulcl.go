package xlbcli

import (
	"net"
	"sort"
)

type SessionUlCl struct {
	CommonAPI
}

type UlclInformationGet struct {
	UlclInfo []SessionUlClMod `json:"ulclAttr"`
}

type SessionUlClMod struct {
	Ident string  `json:"ulclIdent" yaml:"ulclIdent"`
	Args  UlClArg `json:"ulclArgument" yaml:"ulclArgument"`
}

type UlClArg struct {
	Addr net.IP `json:"ulclIP" yaml:"ulclIP"`
	Qfi  uint8  `json:"qfi" yaml:"qfi"`
}

type ConfigurationSessionUlclFile struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
	Spec       SessionUlClMod `yaml:"spec"`
}

func (ulclresp UlclInformationGet) Sort() {
	sort.Slice(ulclresp.UlclInfo, func(i, j int) bool {
		return ulclresp.UlclInfo[i].Ident < ulclresp.UlclInfo[j].Ident
	})
}

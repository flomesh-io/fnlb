package xlbcli

import (
	"fmt"
	"net"
	"sort"
)

type Session struct {
	CommonAPI
}

type SessionInformationGet struct {
	SessionInfo []SessionMod `json:"sessionAttr"`
}

type SessionMod struct {
	Ident string  `json:"ident" yaml:"ident"`
	Ip    net.IP  `json:"sessionIP" yaml:"sessionIP"`
	AnTun SessTun `json:"accessNetworkTunnel" yaml:"accessNetworkTunnel"`
	CnTun SessTun `json:"coreNetworkTunnel" yaml:"coreNetworkTunnel"`
}

type SessTun struct {
	TeID uint32 `json:"teID" yaml:"teID"`
	Addr net.IP `json:"tunnelIP" yaml:"tunnelIP"`
}

type ConfigurationSessionFile struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
	Spec       SessionMod `yaml:"spec"`
}

func (s SessionMod) Validation() error {
	if s.AnTun.TeID == 0 || s.CnTun.TeID == 0 {
		return fmt.Errorf("TeID need to be not 0")
	}
	return nil
}

func (sessionresp SessionInformationGet) Sort() {
	sort.Slice(sessionresp.SessionInfo, func(i, j int) bool {
		return sessionresp.SessionInfo[i].Ident < sessionresp.SessionInfo[j].Ident
	})
}

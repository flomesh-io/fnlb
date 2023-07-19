package xlbcli

import (
	"fmt"
	"sort"
)

type Conntrack struct {
	CommonAPI
}

type CtInformationGet struct {
	CtInfo []ConntrackInformation `json:"ctAttr"`
}

type ConntrackInformation struct {
	Dip    string `json:"destinationIP"`
	Sip    string `json:"sourceIP"`
	Dport  uint16 `json:"destinationPort"`
	Sport  uint16 `json:"sourcePort"`
	Proto  string `json:"protocol"`
	CState string `json:"conntrackState"`
	CAct   string `json:"conntrackAct"`
	Pkts   uint64 `json:"packets"`
	Bytes  uint64 `json:"bytes"`
}

func (ct ConntrackInformation) Key() string {
	return fmt.Sprintf("%s|%s|%05d|%05d|%s", ct.Dip, ct.Sip, ct.Dport, ct.Sport, ct.Proto)
}

func (ctresp CtInformationGet) Sort() {
	sort.Slice(ctresp.CtInfo, func(i, j int) bool {
		return ctresp.CtInfo[i].Key() < ctresp.CtInfo[j].Key()
	})
}

package xlbcli

import "sort"

type Vxlan struct {
	CommonAPI
}
type VxlanGet struct {
	VxlanAttr []VxlanDump `json:"vxlanAttr"`
}

type VxlanDump struct {
	VxlanName string `json:"vxlanName"`
	// VxLanID - name of the endpoint device in the vxlan
	VxLanID int `json:"vxlanID"`
	// EndpointDev - name of the endpoint device in the vxlan
	EndpointDev string `json:"epIntf"`
	// PeerIP - Peer IP address in the vxlan config
	PeerIP []string `json:"peerIP"`
}

// VxlanBridgeMod - Info about an Vxlan bridge
type VxlanBridgeMod struct {
	// VxLanID - name of the endpoint device in the vxlan
	VxLanID int `json:"vxlanID" yaml:"vxlanID"`
	// EndpointDev - name of the endpoint device in the vxlan
	EndpointDev string `json:"epIntf" yaml:"epIntf"`
}

// VxlanPeerMod - Info about an Vlan bridge
type VxlanPeerMod struct {
	// PeerIP - Peer IP address in the vxlan config
	PeerIP string `json:"peerIP" yaml:"peerIP"`
}
type ConfigurationVxlanFile struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
	Spec       VxlanBridgeMod `yaml:"spec"`
}

type ConfigurationVxlanPeerFile struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
	Spec       VxlanPeerMod `yaml:"spec"`
}

func (vxlanresp VxlanGet) Sort() {
	sort.Slice(vxlanresp.VxlanAttr, func(i, j int) bool {
		return vxlanresp.VxlanAttr[i].VxLanID < vxlanresp.VxlanAttr[j].VxLanID
	})
}

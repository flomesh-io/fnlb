package xlbnlp

func GetFDBNoHook() ([]map[string]string, error) {
	panic("Unsupported!")
}

func AddFDBNoHook(macAddress, ifName string) int {
	panic("Unsupported!")
}

func DelFDBNoHook(macAddress, ifName string) int {
	panic("Unsupported!")
}

func AddAddrNoHook(address, ifName string) int {
	panic("Unsupported!")
}

func DelAddrNoHook(address, ifName string) int {
	panic("Unsupported!")
}

func AddNeighNoHook(address, ifName, macAddress string) int {
	panic("Unsupported!")
}

func DelNeighNoHook(address, ifName string) int {
	panic("Unsupported!")
}

func AddRouteNoHook(DestinationIPNet, gateway string) int {
	panic("Unsupported!")
}

func DelRouteNoHook(DestinationIPNet string) int {
	panic("Unsupported!")
}

func AddVLANNoHook(vlanid int) int {
	panic("Unsupported!")
}

func DelVLANNoHook(vlanid int) int {
	panic("Unsupported!")
}

func AddVLANMemberNoHook(vlanid int, intfName string, tagged bool) int {
	panic("Unsupported!")
}

func DelVLANMemberNoHook(vlanid int, intfName string, tagged bool) int {
	panic("Unsupported!")
}

func AddVxLANBridgeNoHook(vxlanid int, epIntfName string) int {
	panic("Unsupported!")
}

func DelVxLANNoHook(vxlanid int) int {
	panic("Unsupported!")
}

func GetVxLANPeerNoHook() (map[int][]string, error) {
	panic("Unsupported!")
}

func AddVxLANPeerNoHook(vxlanid int, PeerIP string) int {
	panic("Unsupported!")
}

func DelVxLANPeerNoHook(vxlanid int, PeerIP string) int {
	panic("Unsupported!")
}

func NlpInit() {
	panic("Unsupported!")
}

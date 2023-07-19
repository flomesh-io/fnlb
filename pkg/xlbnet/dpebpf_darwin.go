package xlbnet

import tk "github.com/cybwan/fsmxlb/pkg/xlblib"

// DpEbpfSetLogLevel - Set log level for ebpf subsystem
func DpEbpfSetLogLevel(logLevel tk.LogLevelT) {
	panic("Unsupported!")
}

// DpEbpfInit - initialize the ebpf dp subsystem
func DpEbpfInit(clusterEn bool, nodeNum int, rssEn bool, logLevel tk.LogLevelT) *DpEbpfH {
	panic("Unsupported!")
}

// DpEbpfUnInit - uninitialize the ebpf dp subsystem
func (e *DpEbpfH) DpEbpfUnInit() {
	panic("Unsupported!")
}

func (e *DpEbpfH) DpMirrAdd(*MirrDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpMirrDel(*MirrDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpPolAdd(*PolDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpPolDel(*PolDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpPortPropAdd(*PortDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpPortPropDel(*PortDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpL2AddrAdd(*L2AddrDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpL2AddrDel(*L2AddrDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpRouterMacAdd(*RouterMacDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpRouterMacDel(*RouterMacDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpNextHopAdd(*NextHopDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpNextHopDel(*NextHopDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpRouteAdd(*RouteDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpRouteDel(*RouteDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpNatLbRuleAdd(*NatDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpNatLbRuleDel(*NatDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpFwRuleAdd(w *FwDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpFwRuleDel(w *FwDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpStat(*StatDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpUlClAdd(w *UlClDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpUlClDel(w *UlClDpWorkQ) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpTableGet(w *TableDpWorkQ) (DpRetT, error) {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpCtAdd(w *DpCtInfo) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpCtDel(w *DpCtInfo) int {
	panic("Unsupported!")
}
func (e *DpEbpfH) DpCtGetAsync() {
	panic("Unsupported!")
}

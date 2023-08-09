package xlbnet

import (
	cmn "github.com/flomesh-io/fnlb/pkg/common"
	tk "github.com/flomesh-io/fnlb/pkg/xlblib"
)

// This file implements interface defined in cmn.NetHookInterface
// The implementation is thread-safe and can be called by multiple-clients at once

// NetAPIStruct - empty struct for anchoring client routines
type NetAPIStruct struct {
}

// NetAPIInit - Initialize a new instance of NetAPI
func NetAPIInit() *NetAPIStruct {
	na := new(NetAPIStruct)
	return na
}

// NetMirrorGet - Get a mirror in xlbnet
func (*NetAPIStruct) NetMirrorGet() ([]cmn.MirrGetMod, error) {
	// There is no locking requirement for this operation
	ret, _ := mh.zr.Mirrs.MirrGet()
	return ret, nil
}

// NetMirrorAdd - Add a mirror in xlbnet
func (*NetAPIStruct) NetMirrorAdd(mm *cmn.MirrMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Mirrs.MirrAdd(mm.Ident, mm.Info, mm.Target)
	return ret, err
}

// NetMirrorDel - Delete a mirror in xlbnet
func (*NetAPIStruct) NetMirrorDel(mm *cmn.MirrMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Mirrs.MirrDelete(mm.Ident)
	return ret, err
}

// NetPortGet - Get Port Information of xlbnet
func (*NetAPIStruct) NetPortGet() ([]cmn.PortDump, error) {
	ret, err := mh.zr.Ports.PortsToGet()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// NetPortAdd - Add a port in xlbnet
func (*NetAPIStruct) NetPortAdd(pm *cmn.PortMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Ports.PortAdd(pm.Dev, pm.LinkIndex, pm.Ptype, RootZone,
		PortHwInfo{pm.MacAddr, pm.Link, pm.State, pm.Mtu, pm.Master, pm.Real,
			uint32(pm.TunID), pm.TunSrc, pm.TunDst}, PortLayer2Info{false, 0})

	return ret, err
}

// NetPortDel - Delete port from xlbnet
func (*NetAPIStruct) NetPortDel(pm *cmn.PortMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Ports.PortDel(pm.Dev, pm.Ptype)
	return ret, err
}

// NetVlanGet - Get Vlan Information of xlbnet
func (*NetAPIStruct) NetVlanGet() ([]cmn.VlanGet, error) {
	ret, err := mh.zr.Vlans.VlanGet()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// NetVlanAdd - Add vlan info to xlbnet
func (*NetAPIStruct) NetVlanAdd(vm *cmn.VlanMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Vlans.VlanAdd(vm.Vid, vm.Dev, RootZone, vm.LinkIndex,
		PortHwInfo{vm.MacAddr, vm.Link, vm.State, vm.Mtu, "", "", vm.TunID, nil, nil})
	if ret == VlanExistsErr {
		ret = 0
	}

	return ret, err
}

// NetVlanDel - Delete vlan info from xlbnet
func (*NetAPIStruct) NetVlanDel(vm *cmn.VlanMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Vlans.VlanDelete(vm.Vid)
	return ret, err
}

// NetVlanPortAdd - Add a port to vlan in xlbnet
func (*NetAPIStruct) NetVlanPortAdd(vm *cmn.VlanPortMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Vlans.VlanPortAdd(vm.Vid, vm.Dev, vm.Tagged)
	return ret, err
}

// NetVlanPortDel - Delete a port from vlan in xlbnet
func (*NetAPIStruct) NetVlanPortDel(vm *cmn.VlanPortMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Vlans.VlanPortDelete(vm.Vid, vm.Dev, vm.Tagged)
	return ret, err
}

// NetIpv4AddrGet - Get an IPv4 Address info from xlbnet
func (*NetAPIStruct) NetAddrGet() ([]cmn.IpAddrGet, error) {
	// There is no locking requirement for this operation
	ret := mh.zr.L3.IfaGet()
	return ret, nil
}

// NetAddrAdd - Add an ipv4 address in xlbnet
func (*NetAPIStruct) NetAddrAdd(am *cmn.IpAddrMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.L3.IfaAdd(am.Dev, am.IP)
	return ret, err
}

// NetAddrDel - Delete an ipv4 address in xlbnet
func (*NetAPIStruct) NetAddrDel(am *cmn.IpAddrMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.L3.IfaDelete(am.Dev, am.IP)
	return ret, err
}

// NetNeighGet - Get a neighbor in xlbnet
func (*NetAPIStruct) NetNeighGet() ([]cmn.NeighMod, error) {
	ret, err := mh.zr.Nh.NeighGet()
	return ret, err
}

// NetNeighAdd - Add a neighbor in xlbnet
func (*NetAPIStruct) NetNeighAdd(nm *cmn.NeighMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Nh.NeighAdd(nm.IP, RootZone, NeighAttr{nm.LinkIndex, nm.State, nm.HardwareAddr})
	if err != nil {
		if ret != NeighExistsErr {
			return ret, err
		}
	}

	return 0, nil
}

// NetNeighDel - Delete a neighbor in xlbnet
func (*NetAPIStruct) NetNeighDel(nm *cmn.NeighMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Nh.NeighDelete(nm.IP, RootZone)
	return ret, err
}

// NetFdbAdd - Add a forwarding database entry in xlbnet
func (*NetAPIStruct) NetFdbAdd(fm *cmn.FdbMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()
	fdbKey := FdbKey{fm.MacAddr, fm.BridgeID}
	fdbAttr := FdbAttr{fm.Dev, fm.Dst, fm.Type}
	ret, err := mh.zr.L2.L2FdbAdd(fdbKey, fdbAttr)
	return ret, err
}

// NetFdbDel - Delete a forwarding database entry in xlbnet
func (*NetAPIStruct) NetFdbDel(fm *cmn.FdbMod) (int, error) {
	fdbKey := FdbKey{fm.MacAddr, fm.BridgeID}
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.L2.L2FdbDel(fdbKey)
	return ret, err
}

// NetRouteGet - Get Route info from xlbnet
func (*NetAPIStruct) NetRouteGet() ([]cmn.RouteGet, error) {
	// There is no locking requirement for this operation
	ret, _ := mh.zr.Rt.RouteGet()
	return ret, nil
}

// NetRouteAdd - Add a route in xlbnet
func (*NetAPIStruct) NetRouteAdd(rm *cmn.RouteMod) (int, error) {
	var ret int
	var err error

	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ra := RtAttr{rm.Protocol, rm.Flags, false, rm.LinkIndex}
	if rm.Gw != nil {
		na := []RtNhAttr{{rm.Gw, rm.LinkIndex}}
		ret, err = mh.zr.Rt.RtAdd(rm.Dst, RootZone, ra, na)
	} else {
		ret, err = mh.zr.Rt.RtAdd(rm.Dst, RootZone, ra, nil)
	}

	return ret, err
}

// NetRouteDel - Delete a route in xlbnet
func (*NetAPIStruct) NetRouteDel(rm *cmn.RouteMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Rt.RtDelete(rm.Dst, RootZone)
	return ret, err
}

// NetLbRuleAdd - Add a load-balancer rule in xlbnet
func (*NetAPIStruct) NetLbRuleAdd(lm *cmn.LbRuleMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()
	var ips []string
	ret, err := mh.zr.Rules.AddNatLbRule(lm.Serv, lm.SecIPs[:], lm.Eps[:])
	if err == nil && lm.Serv.Bgp {
		if mh.bgp != nil {
			ips = append(ips, lm.Serv.ServIP)
			for _, ip := range lm.SecIPs {
				ips = append(ips, ip.SecIP)
			}
			mh.bgp.AddBGPRule("default", ips)
		} else {
			tk.LogIt(tk.LogDebug, "fsmxlb BGP mode is disabled \n")
		}
	}
	return ret, err
}

// NetLbRuleDel - Delete a load-balancer rule in xlbnet
func (*NetAPIStruct) NetLbRuleDel(lm *cmn.LbRuleMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ips := mh.zr.Rules.GetNatLbRuleSecIPs(lm.Serv)
	ret, err := mh.zr.Rules.DeleteNatLbRule(lm.Serv)
	if lm.Serv.Bgp {
		if mh.bgp != nil {
			ips = append(ips, lm.Serv.ServIP)
			mh.bgp.DelBGPRule("default", ips)
		} else {
			tk.LogIt(tk.LogDebug, "fsmxlb BGP mode is disabled \n")
		}
	}
	return ret, err
}

// NetLbRuleGet - Get a load-balancer rule from xlbnet
func (*NetAPIStruct) NetLbRuleGet() ([]cmn.LbRuleMod, error) {
	ret, err := mh.zr.Rules.GetNatLbRule()
	return ret, err
}

// NetCtInfoGet - Get connection track info from xlbnet
func (*NetAPIStruct) NetCtInfoGet() ([]cmn.CtInfo, error) {
	// There is no locking requirement for this operation
	ret := mh.dp.DpMapGetCt4()
	return ret, nil
}

// NetSessionAdd - Add a 3gpp user-session info in xlbnet
func (*NetAPIStruct) NetSessionAdd(sm *cmn.SessionMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Sess.SessAdd(sm.Ident, sm.IP, sm.AnTun, sm.CnTun)
	return ret, err
}

// NetSessionDel - Delete a 3gpp user-session info in xlbnet
func (*NetAPIStruct) NetSessionDel(sm *cmn.SessionMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Sess.SessDelete(sm.Ident)
	return ret, err
}

// NetSessionUlClAdd - Add a 3gpp ulcl-filter info in xlbnet
func (*NetAPIStruct) NetSessionUlClAdd(sr *cmn.SessionUlClMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Sess.UlClAddCls(sr.Ident, sr.Args)
	return ret, err
}

// NetSessionUlClDel - Delete a 3gpp ulcl-filter info in xlbnet
func (*NetAPIStruct) NetSessionUlClDel(sr *cmn.SessionUlClMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Sess.UlClDeleteCls(sr.Ident, sr.Args)
	return ret, err
}

// NetSessionGet - Get 3gpp user-session info in xlbnet
func (*NetAPIStruct) NetSessionGet() ([]cmn.SessionMod, error) {
	// There is no locking requirement for this operation
	ret, err := mh.zr.Sess.SessGet()
	return ret, err
}

// NetSessionUlClGet - Get 3gpp ulcl filter info from xlbnet
func (*NetAPIStruct) NetSessionUlClGet() ([]cmn.SessionUlClMod, error) {
	// There is no locking requirement for this operation
	ret, err := mh.zr.Sess.SessUlclGet()
	return ret, err
}

// NetPolicerGet - Get a policer in xlbnet
func (*NetAPIStruct) NetPolicerGet() ([]cmn.PolMod, error) {
	// There is no locking requirement for this operation
	ret, err := mh.zr.Pols.PolGetAll()
	return ret, err
}

// NetPolicerAdd - Add a policer in xlbnet
func (*NetAPIStruct) NetPolicerAdd(pm *cmn.PolMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Pols.PolAdd(pm.Ident, pm.Info, pm.Target)
	return ret, err
}

// NetPolicerDel - Delete a policer in xlbnet
func (*NetAPIStruct) NetPolicerDel(pm *cmn.PolMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Pols.PolDelete(pm.Ident)
	return ret, err
}

// NetCIStateGet - Get current node cluster state
func (*NetAPIStruct) NetCIStateGet() ([]cmn.HASMod, error) {
	// There is no locking requirement for this operation
	ret, err := mh.has.CIStateGet()
	return ret, err
}

// NetCIStateMod - Modify cluster state
func (*NetAPIStruct) NetCIStateMod(hm *cmn.HASMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	_, err := mh.has.CIStateUpdate(*hm)
	if err != nil {
		return -1, err
	}

	return 0, nil
}

// NetFwRuleAdd - Add a firewall rule in xlbnet
func (*NetAPIStruct) NetFwRuleAdd(fm *cmn.FwRuleMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Rules.AddFwRule(fm.Rule, fm.Opts)
	return ret, err
}

// NetFwRuleDel - Delete a firewall rule in xlbnet
func (*NetAPIStruct) NetFwRuleDel(fm *cmn.FwRuleMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Rules.DeleteFwRule(fm.Rule)
	return ret, err
}

// NetFwRuleGet - Get a firewall rule from xlbnet
func (*NetAPIStruct) NetFwRuleGet() ([]cmn.FwRuleMod, error) {
	ret, err := mh.zr.Rules.GetFwRule()
	return ret, err
}

// NetEpHostAdd - Add a LB end-point in xlbnet
func (*NetAPIStruct) NetEpHostAdd(em *cmn.EndPointMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	epArgs := epHostOpts{inActTryThr: em.InActTries, probeType: em.ProbeType,
		probeReq: em.ProbeReq, probeResp: em.ProbeResp,
		probeDuration: em.ProbeDuration, probePort: em.ProbePort,
	}
	ret, err := mh.zr.Rules.AddEPHost(true, em.HostName, em.Name, epArgs)
	return ret, err
}

// NetEpHostDel - Delete a LB end-point in xlbnet
func (*NetAPIStruct) NetEpHostDel(em *cmn.EndPointMod) (int, error) {
	mh.mtx.Lock()
	defer mh.mtx.Unlock()

	ret, err := mh.zr.Rules.DeleteEPHost(true, em.Name, em.HostName, em.ProbeType, em.ProbePort)
	return ret, err
}

// NetEpHostGet - Get LB end-points from xlbnet
func (*NetAPIStruct) NetEpHostGet() ([]cmn.EndPointMod, error) {
	ret, err := mh.zr.Rules.GetEpHosts()
	return ret, err
}

// NetParamSet - Set operational params of xlbnet
func (*NetAPIStruct) NetParamSet(param cmn.ParamMod) (int, error) {
	ret, err := mh.ParamSet(param)
	return ret, err
}

// NetParamGet - Get operational params of xlbnet
func (*NetAPIStruct) NetParamGet(param *cmn.ParamMod) (int, error) {
	ret, err := mh.ParamGet(param)
	return ret, err
}

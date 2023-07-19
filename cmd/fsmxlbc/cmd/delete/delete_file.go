/*
 * Copyright (c) 2022 NetLOX Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package delete

import (
	"context"
	"fmt"
	"github.com/cybwan/fsmxlb/pkg/xlbcli"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"
)

func DeleteFileConfig(file string, restOptions *xlbcli.RESTOptions) error {
	// open file
	var comm xlbcli.TypeMeta
	var err error
	byteBuf, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Unmashal to yaml
	if err := yaml.Unmarshal(byteBuf, &comm); err != nil {
		fmt.Printf("Error: Failed to unmarshal File: (%s)\n", err.Error())
		return err
	}
	switch comm.Kind {
	case "Loadbalancer", "lb", "LB":
		err = LoadBalancerDeleteWithFile(restOptions, byteBuf)
	case "Endpoint", "ep", "endpoints":
		err = EndPointDeleteWithFile(restOptions, byteBuf)
	case "FDB", "fdb":
		err = FDBDeleteWithFile(restOptions, byteBuf)
	case "Firewall", "fw", "firewalls", "firewall":
		err = FirewallDeleteWithFile(restOptions, byteBuf)
	case "ipv4address", "ipv4", "ipaddress", "ip", "IP", "IPaddress":
		err = IPv4AddressDeleteWithFile(restOptions, byteBuf)
	case "mirror", "mirr", "mirrors", "Mirror":
		err = MirrorDeleteWithFile(restOptions, byteBuf)
	case "nei", "neigh", "Neighbor", "Neigh", "neighbor":
		err = NeighborDeleteWithFile(restOptions, byteBuf)
	case "Policy", "pol", "policys", "pols", "polices":
		err = PolicyDeleteWithFile(restOptions, byteBuf)
	case "Route", "route":
		err = RouteDeleteWithFile(restOptions, byteBuf)
	case "Session", "session", "sessions":
		err = SessionDeleteWithFile(restOptions, byteBuf)
	case "SessionULCL", "ulcl", "sessionulcls", "ulcls", "ULCL":
		err = SessionUlClDeleteWithFile(restOptions, byteBuf)
	case "VlanMember", "vlanMember", "vlan-member", "vlan_member", "vlanmember":
		err = VlanMemberDeleteWithFile(restOptions, byteBuf)
	case "Vlan", "vlan":
		err = VlanDeleteWithFile(restOptions, byteBuf)
	case "VxlanPeer", "vxlanpeer", "vxlan-peer", "vxlan_peer":
		err = VxlanPeerDeleteWithFile(restOptions, byteBuf)
	case "Vxlan", "vxlan":
		err = VxlanDeleteWithFile(restOptions, byteBuf)
	default:
		fmt.Printf("Not Supported\n")
	}
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func GetClientWithCtx(restOptions *xlbcli.RESTOptions) (*xlbcli.XlbClient, context.Context, context.CancelFunc) {
	client := xlbcli.NewXlbClient(restOptions)
	ctx := context.TODO()
	var cancel context.CancelFunc
	if restOptions.Timeout > 0 {
		ctx, cancel = context.WithTimeout(context.TODO(), time.Duration(restOptions.Timeout)*time.Second)
	}

	return client, ctx, cancel
}

func LoadBalancerDeleteWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationLBFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	client, ctx, cancel := GetClientWithCtx(restOptions)
	if restOptions.Timeout > 0 {
		defer cancel()
	}
	subResources := []string{
		"externalipaddress", c.Spec.Service.ExternalIP,
		"port", strconv.Itoa(int(c.Spec.Service.Port)),
		"protocol", c.Spec.Service.Protocol,
	}
	qmap := map[string]string{}
	qmap["bgp"] = fmt.Sprintf("%v", c.Spec.Service.BGP)
	qmap["block"] = fmt.Sprintf("%v", c.Spec.Service.Block)
	client.LoadBalancer().SubResources(subResources).Query(qmap).Delete(ctx)

	return nil
}

func NeighborDeleteWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationNeighborFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	client, ctx, cancel := GetClientWithCtx(restOptions)
	if restOptions.Timeout > 0 {
		defer cancel()
	}

	subResources := []string{
		c.Spec.IP, "dev", c.Spec.Dev,
	}
	_, err := client.Neighbor().SubResources(subResources).Delete(ctx)
	if err != nil {
		fmt.Printf("Error: Failed to delete Neighbor\n")
		return err
	}
	return nil
}

func MirrorDeleteWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationMirrorFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	client, ctx, cancel := GetClientWithCtx(restOptions)
	if restOptions.Timeout > 0 {
		defer cancel()
	}
	subResources := []string{"ident", c.Spec.Ident}
	_, err := client.Mirror().SubResources(subResources).Delete(ctx)
	if err != nil {
		fmt.Printf("Error: Failed to delete Mirror\n")
		return err
	}
	return nil
}
func IPv4AddressDeleteWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationIPv4File
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	client, ctx, cancel := GetClientWithCtx(restOptions)
	if restOptions.Timeout > 0 {
		defer cancel()
	}
	subResources := []string{
		c.Spec.IP, "dev", c.Spec.Dev,
	}
	_, err := client.IPv4Address().SubResources(subResources).Delete(ctx)
	if err != nil {
		fmt.Printf("Error: Failed to delete IPv4Address\n")
		return err
	}
	return nil
}

func FDBDeleteWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationFDBFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	client, ctx, cancel := GetClientWithCtx(restOptions)
	if restOptions.Timeout > 0 {
		defer cancel()
	}
	subResources := []string{
		c.Spec.MacAddress, "dev", c.Spec.Dev,
	}
	_, err := client.FDB().SubResources(subResources).Delete(ctx)
	if err != nil {
		fmt.Printf("Error: Failed to delete FDB\n")
		return err
	}
	return nil
}

func EndPointDeleteWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationEndPointFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	client, ctx, cancel := GetClientWithCtx(restOptions)
	if restOptions.Timeout > 0 {
		defer cancel()
	}

	subResources := []string{
		"epipaddress", c.HostName,
		"probetype", c.Spec.ProbeType,
		"probeport", strconv.Itoa(int(c.Spec.ProbePort)),
	}
	_, err := client.EndPoint().SubResources(subResources).Delete(ctx)
	if err != nil {
		fmt.Printf("Error: Failed to delete EndPoint\n")
		return err
	}
	return nil
}

func PolicyDeleteWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationPolicyFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	client, ctx, cancel := GetClientWithCtx(restOptions)
	if restOptions.Timeout > 0 {
		defer cancel()
	}
	subResources := []string{
		"ident", c.Spec.Ident,
	}
	_, err := client.Policy().SubResources(subResources).Delete(ctx)
	if err != nil {
		fmt.Printf("Error: Failed to delete Policy\n")
		return err
	}
	return nil
}

func MakefirewallDeleteRuleToQeury(FirewallRule xlbcli.FwRuleArg) map[string]string {
	query := map[string]string{}

	if FirewallRule.DstIP != "" {
		query["destinationIP"] = FirewallRule.DstIP
	}
	if FirewallRule.DstPortMin != 0 {
		query["minDestinationPort"] = strconv.Itoa(int(FirewallRule.DstPortMin))
	}
	if FirewallRule.DstPortMax != 0 {
		query["maxDestinationPort"] = strconv.Itoa(int(FirewallRule.DstPortMax))
	}
	if FirewallRule.InPort != "" {
		query["portName"] = FirewallRule.InPort
	}
	if FirewallRule.Pref != 0 {
		query["preference"] = strconv.Itoa(int(FirewallRule.Pref))
	}
	if FirewallRule.Proto != 0 {
		query["protocol"] = strconv.Itoa(int(FirewallRule.Proto))
	}
	if FirewallRule.SrcIP != "" {
		query["sourceIP"] = FirewallRule.SrcIP
	}
	if FirewallRule.SrcPortMax != 0 {
		query["maxSourcePort"] = strconv.Itoa(int(FirewallRule.SrcPortMax))
	}
	if FirewallRule.SrcPortMin != 0 {
		query["minSourcePort"] = strconv.Itoa(int(FirewallRule.SrcPortMin))
	}

	return query
}

func FirewallDeleteWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationFWFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	client, ctx, cancel := GetClientWithCtx(restOptions)
	if restOptions.Timeout > 0 {
		defer cancel()
	}
	qeury := MakefirewallDeleteRuleToQeury(c.Spec.Rule)
	_, err := client.Firewall().Query(qeury).Delete(ctx)
	if err != nil {
		fmt.Printf("Error: Failed to delete Firewall\n")
		return err
	}
	return nil
}

func RouteDeleteWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationRouteFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	client, ctx, cancel := GetClientWithCtx(restOptions)
	if restOptions.Timeout > 0 {
		defer cancel()
	}

	subResources := []string{
		"destinationIPNet", c.Spec.Dst,
	}
	_, err := client.Route().SubResources(subResources).Delete(ctx)
	if err != nil {
		fmt.Printf("Error: Failed to delete Route\n")
		return err
	}
	return nil
}

func SessionDeleteWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationSessionFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	client, ctx, cancel := GetClientWithCtx(restOptions)
	if restOptions.Timeout > 0 {
		defer cancel()
	}
	subResources := []string{
		"ident", c.Spec.Ident,
	}
	_, err := client.Session().SubResources(subResources).Delete(ctx)
	if err != nil {
		fmt.Printf("Error: Failed to delete Session\n")
		return err
	}
	return nil
}

func SessionUlClDeleteWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationSessionUlclFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	client, ctx, cancel := GetClientWithCtx(restOptions)
	if restOptions.Timeout > 0 {
		defer cancel()
	}
	subResources := []string{
		"ident", c.Spec.Ident, "ulclAddress", c.Spec.Args.Addr.String(),
	}
	_, err := client.SessionUlCL().SubResources(subResources).Delete(ctx)
	if err != nil {
		fmt.Printf("Error: Failed to delete SessionUlCl\n")
		return err
	}
	return nil
}

func VlanDeleteWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationVlanFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	client, ctx, cancel := GetClientWithCtx(restOptions)
	if restOptions.Timeout > 0 {
		defer cancel()
	}
	subResources := []string{
		strconv.Itoa(c.Spec.Vid),
	}
	_, err := client.Vlan().SubResources(subResources).Delete(ctx)
	if err != nil {
		fmt.Printf("Error: Failed to delete Vlan\n")
		return err
	}
	return nil
}

func VlanMemberDeleteWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationVlanMemberFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	client, ctx, cancel := GetClientWithCtx(restOptions)
	if restOptions.Timeout > 0 {
		defer cancel()
	}
	Tagged := fmt.Sprintf("%v", c.Spec.Tagged)
	subResources := []string{
		strconv.Itoa(c.ObjectMeta.VlanID), "member", c.Spec.Dev, "tagged", Tagged,
	}
	_, err := client.Vlan().SubResources(subResources).Delete(ctx)
	if err != nil {
		fmt.Printf("Error: Failed to delete Vlan\n")
		return err
	}
	return nil
}

func VxlanDeleteWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationVxlanFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	client, ctx, cancel := GetClientWithCtx(restOptions)
	if restOptions.Timeout > 0 {
		defer cancel()
	}

	subResources := []string{
		strconv.Itoa(c.Spec.VxLanID),
	}
	_, err := client.Vxlan().SubResources(subResources).Delete(ctx)
	if err != nil {
		fmt.Printf("Error: Failed to delete Vxlan\n")
		return err
	}
	return nil
}

func VxlanPeerDeleteWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationVxlanPeerFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	client, ctx, cancel := GetClientWithCtx(restOptions)
	if restOptions.Timeout > 0 {
		defer cancel()
	}

	subResources := []string{
		strconv.Itoa(c.ObjectMeta.VxlanID), "peer", c.Spec.PeerIP,
	}
	_, err := client.Vxlan().SubResources(subResources).Delete(ctx)
	if err != nil {
		fmt.Printf("Error: Failed to delete Vxlan\n")
		return err
	}
	return nil
}

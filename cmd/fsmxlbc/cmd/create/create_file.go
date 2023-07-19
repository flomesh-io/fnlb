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
package create

import (
	"fmt"
	"github.com/cybwan/fsmxlb/pkg/xlbcli"
	"gopkg.in/yaml.v2"
)

func EndPointCreateWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationEndPointFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	if _, err := EndPointAPICall(restOptions, c.Spec); err != nil {
		return err
	}
	return nil
}

func FDBCreateWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationFDBFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	if _, err := FDBAPICall(restOptions, c.Spec); err != nil {
		return err
	}
	return nil
}

func FirewallCreateWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationFWFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	if _, err := FirewallAPICall(restOptions, c.Spec); err != nil {
		return err
	}
	return nil
}

func IPv4AddressCreateWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationIPv4File
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	if _, err := IPv4AddressAPICall(restOptions, c.Spec); err != nil {
		return err
	}
	return nil
}

func LoadBalancerCreateWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationLBFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	if _, err := LoadbalancerAPICall(restOptions, c.Spec); err != nil {
		return err
	}
	return nil
}

func MirrorCreateWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationMirrorFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	if _, err := MirrorAPICall(restOptions, c.Spec); err != nil {
		return err
	}
	return nil
}
func NeighborsCreateWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationNeighborFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	if _, err := NeighborsAPICall(restOptions, c.Spec); err != nil {
		return err
	}
	return nil
}

func PolicyCreateWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationPolicyFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	if _, err := PolicyAPICall(restOptions, c.Spec); err != nil {
		return err
	}
	return nil
}

func RouteCreateWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationRouteFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	if _, err := RouteAPICall(restOptions, c.Spec); err != nil {
		return err
	}
	return nil
}

func SessionCreateWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationSessionFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	if err := c.Spec.Validation(); err != nil {
		return err
	}
	if _, err := SessionAPICall(restOptions, c.Spec); err != nil {
		return err
	}
	return nil
}

func SessionUlClCreateWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationSessionUlclFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	if _, err := SessionUlClAPICall(restOptions, c.Spec); err != nil {
		return err
	}
	return nil
}

func VlanMemberCreateWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationVlanMemberFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	// URL Maker
	url := fmt.Sprintf("/config/vlan/%d/member", c.ObjectMeta.VlanID)
	if _, err := VlanMemberAPICall(restOptions, c.Spec, url); err != nil {
		return err
	}
	return nil
}

func VlanBridgeCreateWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationVlanFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}

	if _, err := VlanBridgeAPICall(restOptions, c.Spec); err != nil {
		return err
	}
	return nil
}

func VxlanPeerCreateWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationVxlanPeerFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	// URL Maker
	url := fmt.Sprintf("/config/tunnel/vxlan/%d/peer", c.ObjectMeta.VxlanID)
	if _, err := VxlanPeerAPICall(restOptions, c.Spec, url); err != nil {
		return err
	}
	return nil
}

func VxlanBridgeCreateWithFile(restOptions *xlbcli.RESTOptions, byteBuf []byte) error {
	var c xlbcli.ConfigurationVxlanFile
	if err := yaml.Unmarshal(byteBuf, &c); err != nil {
		return err
	}
	if _, err := VxlanBridgeAPICall(restOptions, c.Spec); err != nil {
		return err
	}
	return nil
}

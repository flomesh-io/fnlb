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
	"fmt"
	"github.com/flomesh-io/fsmxlb/pkg/xlbcli"
	"github.com/spf13/cobra"
)

func DeleteCmd(restOptions *xlbcli.RESTOptions) *cobra.Command {

	var NormalConfigFile string
	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a Load balance features in the fsmxlb.",
		Long: `Delete a Load balance features in the fsmxlb. 
Delete - Service type external load-balancer, Vlan, Vxlan, Qos Policies,
	 Endpoint client,FDB, IPaddress, Neighbor, Route,Firewall, Mirror, Session, UlCl
		`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(NormalConfigFile) > 0 {
				if err := DeleteFileConfig(NormalConfigFile, restOptions); err != nil {
					fmt.Printf("Configuration failed - %s\n", NormalConfigFile)
				} else {
					fmt.Printf("Configuration applied - %s\n", NormalConfigFile)
				}
			}
			if len(NormalConfigFile) == 0 && len(args) == 0 {
				cmd.Help()
			}

		},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			fmt.Printf("Error: unknown command \"%v\"for \"fsmxlbc\" \nRun \"fsmxlbc --help\" for usage.\n", args)
			cmd.Help()
			return err
		},
	}
	deleteCmd.AddCommand(NewDeleteLoadBalancerCmd(restOptions))
	deleteCmd.AddCommand(NewDeleteSessionCmd(restOptions))
	deleteCmd.AddCommand(NewDeleteSessionUlClCmd(restOptions))
	deleteCmd.AddCommand(NewDeletePolicyCmd(restOptions))
	deleteCmd.AddCommand(NewDeleteRouteCmd(restOptions))
	deleteCmd.AddCommand(NewDeleteIPv4AddressCmd(restOptions))
	deleteCmd.AddCommand(NewDeleteNeighborsCmd(restOptions))
	deleteCmd.AddCommand(NewDeleteFDBCmd(restOptions))
	deleteCmd.AddCommand(NewDeleteVlanBridgeCmd(restOptions))
	deleteCmd.AddCommand(NewDeleteVlanMemberCmd(restOptions))
	deleteCmd.AddCommand(NewDeleteVxlanBridgeCmd(restOptions))
	deleteCmd.AddCommand(NewDeleteVxlanPeerCmd(restOptions))
	deleteCmd.AddCommand(NewDeleteMirrorCmd(restOptions))
	deleteCmd.AddCommand(NewDeleteFirewallCmd(restOptions))
	deleteCmd.AddCommand(NewDeleteEndPointCmd(restOptions))
	deleteCmd.Flags().StringVarP(&NormalConfigFile, "file", "f", "", "Config file to apply as like K8s")
	return deleteCmd
}

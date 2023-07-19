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
package cmd

import (
	"fmt"
	"github.com/cybwan/fsmxlb/pkg/version"
	"github.com/cybwan/fsmxlb/pkg/xlbcli"
	"os"

	"github.com/spf13/cobra"

	"github.com/cybwan/fsmxlb/cmd/fsmxlbc/cmd/create"
	"github.com/cybwan/fsmxlb/cmd/fsmxlbc/cmd/delete"
	"github.com/cybwan/fsmxlb/cmd/fsmxlbc/cmd/dump"
	"github.com/cybwan/fsmxlb/cmd/fsmxlbc/cmd/get"
	"github.com/cybwan/fsmxlb/cmd/fsmxlbc/cmd/set"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get a version",
	Long:  `It shows fsmxlbc version.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %s %s\n", version.Version, version.GitCommit, version.BuildDate)
	},
}

var CompletionCmd = &cobra.Command{
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate completion script",
	Long:                  "To load completions",
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "fsmxlbc",
		Short: "fsmxlbc is the command-line tool for fsmxlb.",
		Long: `fsmxlbc is the command-line tool for fsmxlb. It is equivalent of "kubectl" for fsmxlb. fsmxlbc provides the following (currently) :
	- Create/Delete/Get - Service type external load-balancer, Vlan, Vxlan, Qos Policies, Endpoint client,FDB, IPaddress, Neighbor, Route,Firewall, Mirror, Session, UlCl
	- Get Port(interface) dump used by fsmxlb or its docker
	- Get Connection track (TCP/UDP/ICMP/SCTP) information
fsmxlbc aim to provide all of the configuation for the fsmxlb.`,
	}
	restOptions := &xlbcli.RESTOptions{}
	saveOptions := &dump.SaveOptions{}
	applyOptions := &dump.ApplyOptions{}

	rootCmd.PersistentFlags().Int16VarP(&restOptions.Timeout, "timeout", "t", 10, "Set timeout")
	rootCmd.PersistentFlags().StringVarP(&restOptions.Protocol, "protocol", "", "http", "Set API server http/https")
	rootCmd.PersistentFlags().StringVarP(&restOptions.PrintOption, "output", "o", "", "Set output layer (ex.) wide, json)")
	rootCmd.PersistentFlags().StringVarP(&restOptions.ServerIP, "cli.erver", "s", "127.0.0.1", "Set API server IP address")
	rootCmd.PersistentFlags().Int16VarP(&restOptions.ServerPort, "port", "p", 11111, "Set API server port number")

	rootCmd.AddCommand(get.GetCmd(restOptions))
	rootCmd.AddCommand(create.CreateCmd(restOptions))
	rootCmd.AddCommand(delete.DeleteCmd(restOptions))
	rootCmd.AddCommand(set.SetParamCmd(restOptions))

	saveCmd := dump.SaveCmd(saveOptions, restOptions)
	applyCmd := dump.ApplyCmd(applyOptions, restOptions)

	saveCmd.Flags().BoolVarP(&saveOptions.SaveAllConfig, "all", "a", false, "Saves all fsmxlb configuration")
	saveCmd.Flags().BoolVarP(&saveOptions.SaveIpConfig, "ip", "i", false, "Saves IP configuration")
	saveCmd.Flags().BoolVarP(&saveOptions.SaveLBConfig, "lb", "l", false, "Saves Load Balancer rules configuration")
	saveCmd.Flags().BoolVarP(&saveOptions.SaveSessionConfig, "session", "", false, "Saves session configuration")
	saveCmd.Flags().BoolVarP(&saveOptions.SaveUlClConfig, "ulcl", "", false, "Saves ulcl configuration")
	saveCmd.Flags().BoolVarP(&saveOptions.SaveFWConfig, "firewall", "", false, "Saves firewall configuration")
	saveCmd.Flags().BoolVarP(&saveOptions.SaveEPConfig, "endpoint", "", false, "Saves endpoint configuration")

	saveCmd.MarkFlagsMutuallyExclusive("all", "ip", "lb", "session", "ulcl", "firewall", "endpoint")

	applyCmd.Flags().StringVarP(&applyOptions.IpConfigFile, "ip", "i", "", "IP config file to apply")
	applyCmd.Flags().StringVarP(&applyOptions.Intf, "per-intf", "", "", "Apply configuration only for specific interface")
	applyCmd.Flags().BoolVarP(&applyOptions.Route, "ipv4route", "r", false, "Apply route configuration only for specific interface")
	applyCmd.Flags().StringVarP(&applyOptions.ConfigPath, "config-path", "c", "/opt/fsmxlb/ipconfig/", "Configuration path only for applying per interface config")
	applyCmd.Flags().StringVarP(&applyOptions.LBConfigFile, "lb", "l", "", "Load Balancer config file to apply")
	applyCmd.Flags().StringVarP(&applyOptions.SessionConfigFile, "session", "", "", "Session config file to apply")
	applyCmd.Flags().StringVarP(&applyOptions.SessionUlClConfigFile, "ulcl", "", "", "Ulcl config file to apply")
	applyCmd.Flags().StringVarP(&applyOptions.FWConfigFile, "firewall", "", "", "Firewall config file to apply")
	applyCmd.Flags().StringVarP(&applyOptions.NormalConfigFile, "file", "f", "", "Config file to apply as like K8s")

	rootCmd.AddCommand(saveCmd)
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(CompletionCmd)
	rootCmd.AddCommand(VersionCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

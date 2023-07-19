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
	"context"
	"errors"
	"fmt"
	"github.com/cybwan/fsmxlb/pkg/xlbcli"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func NewCreateIPv4AddressCmd(restOptions *xlbcli.RESTOptions) *cobra.Command {
	var createIPv4AddressCmd = &cobra.Command{
		Use:   "ip <DeviceIPNet> <device>",
		Short: "Create a IPv4Address",
		Long: `Create a IPv4Address using fsmxlb. It is working as "ip addr add <DeviceIPNet> dev <device>"
ex) fsmxlbc create ip 192.168.0.1/24 eno7
`,
		Aliases: []string{"ipv4address", "ipv4", "ipaddress"},
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var IPv4AddressMod xlbcli.Ipv4AddrMod
			// Make IPv4AddressMod
			if err := ReadCreateIPv4AddressOptions(&IPv4AddressMod, args); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			resp, err := IPv4AddressAPICall(restOptions, IPv4AddressMod)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				PrintCreateResult(resp, *restOptions)
				return
			}

		},
	}

	return createIPv4AddressCmd
}

func ReadCreateIPv4AddressOptions(o *xlbcli.Ipv4AddrMod, args []string) error {
	if len(args) > 2 {
		return errors.New("create IPv4Address command get so many args")
	} else if len(args) <= 1 {
		return errors.New("create IPv4Address need <DeviceIPNet>  args")
	}

	if _, _, err := net.ParseCIDR(args[0]); err != nil {
		return fmt.Errorf("DeviceIPNet '%s' is invalid format", args[0])
	}
	o.IP = args[0]
	o.Dev = args[1]

	return nil
}

func IPv4AddressAPICall(restOptions *xlbcli.RESTOptions, IPv4AddressModel xlbcli.Ipv4AddrMod) (*http.Response, error) {
	client := xlbcli.NewXlbClient(restOptions)
	ctx := context.TODO()
	var cancel context.CancelFunc
	if restOptions.Timeout > 0 {
		ctx, cancel = context.WithTimeout(context.TODO(), time.Duration(restOptions.Timeout)*time.Second)
		defer cancel()
	}

	return client.IPv4Address().Create(ctx, IPv4AddressModel)
}

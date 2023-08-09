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
	"github.com/flomesh-io/fnlb/pkg/xlbcli"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func NewCreateRouteCmd(restOptions *xlbcli.RESTOptions) *cobra.Command {
	var createRouteCmd = &cobra.Command{
		Use:   "route <DestinationIPNet> <gateway>",
		Short: "Create a Route",
		Long: `Create a Route using fsmxlb. It is working as "ip route add <DestinationIPNet> via <gateway>"
	
ex) fsmxlbc create route 192.168.212.0/24 172.17.0.254
`,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var RouteMod xlbcli.Routev4Get
			// Make RouteMod
			if err := ReadCreateRouteOptions(&RouteMod, args); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			resp, err := RouteAPICall(restOptions, RouteMod)
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

	return createRouteCmd
}

func ReadCreateRouteOptions(o *xlbcli.Routev4Get, args []string) error {
	if len(args) > 3 {
		return errors.New("create Route command get so many args")
	} else if len(args) <= 1 {
		return errors.New("create Route need <DestinationIPNet> and <gateway> args")
	}

	if _, _, err := net.ParseCIDR(args[0]); err != nil {
		return fmt.Errorf("DestinationIPNet '%s' is invalid format", args[0])
	}
	o.Dst = args[0]

	if val := net.ParseIP(args[1]); val != nil {
		o.Gw = args[1]
	} else {
		return fmt.Errorf("gateway IP '%s' is invalid format", args[1])
	}
	return nil
}

func RouteAPICall(restOptions *xlbcli.RESTOptions, RouteModel xlbcli.Routev4Get) (*http.Response, error) {
	client := xlbcli.NewXlbClient(restOptions)
	ctx := context.TODO()
	var cancel context.CancelFunc
	if restOptions.Timeout > 0 {
		ctx, cancel = context.WithTimeout(context.TODO(), time.Duration(restOptions.Timeout)*time.Second)
		defer cancel()
	}

	return client.Route().Create(ctx, RouteModel)
}

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
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func NewCreateVxlanBridgeCmd(restOptions *xlbcli.RESTOptions) *cobra.Command {
	var createvxlanCmd = &cobra.Command{
		Use:   "vxlan <VxlanID> <EndpointDeviceName>",
		Short: "Create a vxlan",
		Long: `Create a vxlan using fsmxlb. 

ex) fsmxlbc create vxlan 100 eno7

`,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var vxlanMod xlbcli.VxlanBridgeMod
			// Make vxlanMod
			if err := ReadCreateVxlanBridgeOptions(&vxlanMod, args); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			VxLanID, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			vxlanMod.VxLanID = VxLanID
			resp, err := VxlanBridgeAPICall(restOptions, vxlanMod)
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

	return createvxlanCmd
}

func ReadCreateVxlanBridgeOptions(o *xlbcli.VxlanBridgeMod, args []string) error {
	if len(args) > 3 {
		return errors.New("create vxlan command get so many args")
	} else if len(args) <= 1 {
		return errors.New("create vxlan need <VxlanID> <EndpointDeviceName> args")
	}

	o.EndpointDev = args[1]

	return nil
}

func VxlanBridgeAPICall(restOptions *xlbcli.RESTOptions, vxlanModel xlbcli.VxlanBridgeMod) (*http.Response, error) {
	client := xlbcli.NewXlbClient(restOptions)
	ctx := context.TODO()
	var cancel context.CancelFunc
	if restOptions.Timeout > 0 {
		ctx, cancel = context.WithTimeout(context.TODO(), time.Duration(restOptions.Timeout)*time.Second)
		defer cancel()
	}

	return client.Vxlan().Create(ctx, vxlanModel)
}

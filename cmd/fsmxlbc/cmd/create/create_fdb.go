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
	"github.com/flomesh-io/fsmxlb/pkg/xlbcli"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func NewCreateFDBCmd(restOptions *xlbcli.RESTOptions) *cobra.Command {
	var createFDBCmd = &cobra.Command{
		Use:   "fdb <MacAddress> <DeviceName>",
		Short: "Create a FDB",
		Long: `Create a FDB using fsmxlb. It is working as "bridge fdb add <MacAddress> dev <device>"
ex) fsmxlbc create fdb aa:aa:aa:aa:bb:bb eno7	
`,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var FDBMod xlbcli.FDBMod
			// Make FDBMod
			if err := ReadCreateFDBOptions(&FDBMod, args); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			resp, err := FDBAPICall(restOptions, FDBMod)
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

	return createFDBCmd
}

func ReadCreateFDBOptions(o *xlbcli.FDBMod, args []string) error {
	if len(args) > 2 {
		return errors.New("create FDB command get so many args")
	} else if len(args) <= 1 {
		return errors.New("create FDB need <MacAddress>  args")
	}

	if _, err := net.ParseMAC(args[0]); err != nil {
		return fmt.Errorf("MacAddress '%s' is invalid format", args[0])
	}
	o.MacAddress = args[0]
	o.Dev = args[1]

	return nil
}

func FDBAPICall(restOptions *xlbcli.RESTOptions, FDBModel xlbcli.FDBMod) (*http.Response, error) {
	client := xlbcli.NewXlbClient(restOptions)
	ctx := context.TODO()
	var cancel context.CancelFunc
	if restOptions.Timeout > 0 {
		ctx, cancel = context.WithTimeout(context.TODO(), time.Duration(restOptions.Timeout)*time.Second)
		defer cancel()
	}

	return client.FDB().Create(ctx, FDBModel)
}

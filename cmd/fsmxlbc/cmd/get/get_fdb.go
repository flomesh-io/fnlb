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
package get

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cybwan/fsmxlb/pkg/xlbcli"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

func NewGetFDBCmd(restOptions *xlbcli.RESTOptions) *cobra.Command {
	var GetFDBCmd = &cobra.Command{
		Use:   "fdb",
		Short: "Get a fdb",
		Long:  `It shows fdb Information in the fsmxlb`,

		Run: func(cmd *cobra.Command, args []string) {
			client := xlbcli.NewXlbClient(restOptions)
			ctx := context.TODO()
			var cancel context.CancelFunc
			if restOptions.Timeout > 0 {
				ctx, cancel = context.WithTimeout(context.TODO(), time.Duration(restOptions.Timeout)*time.Second)
				defer cancel()
			}
			resp, err := client.FDB().SetUrl("/config/fdb/all").Get(ctx)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			if resp.StatusCode == http.StatusOK {
				PrintGetFDBResult(resp, *restOptions)
				return
			}

		},
	}

	return GetFDBCmd
}

func PrintGetFDBResult(resp *http.Response, o xlbcli.RESTOptions) {
	FDBresp := xlbcli.FDBModGet{}
	var data [][]string
	resultByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: Failed to read HTTP response: (%s)\n", err.Error())
		return
	}

	if err := json.Unmarshal(resultByte, &FDBresp); err != nil {
		fmt.Printf("Error: Failed to unmarshal HTTP response: (%s)\n", err.Error())
		return
	}

	// if json options enable, it print as a json format.
	if o.PrintOption == "json" {
		resultIndent, _ := json.MarshalIndent(FDBresp, "", "    ")
		fmt.Println(string(resultIndent))
		return
	}

	FDBresp.Sort()

	// Table Init
	table := TableInit()

	// Making fdb data
	for _, fdb := range FDBresp.FdbAttr {

		table.SetHeader(FDB_TITLE)
		data = append(data, []string{fdb.Dev, fdb.MacAddress})

	}
	// Rendering the fdb data to table
	TableShow(data, table)
}

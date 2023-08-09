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
	"errors"
	"fmt"
	"github.com/flomesh-io/fnlb/pkg/xlbcli"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func NewGetPolicyCmd(restOptions *xlbcli.RESTOptions) *cobra.Command {
	var GetPolCmd = &cobra.Command{
		Use:     "policy",
		Short:   "Get a Policy",
		Aliases: []string{"pol", "policys", "pols", "polices"},
		Long:    `It shows policy Informations`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd
			_ = args
			client := xlbcli.NewXlbClient(restOptions)
			ctx := context.TODO()
			var cancel context.CancelFunc
			if restOptions.Timeout > 0 {
				ctx, cancel = context.WithTimeout(context.TODO(), time.Duration(restOptions.Timeout)*time.Second)
				defer cancel()
			}
			resp, err := client.Policy().SetUrl("/config/policy/all").Get(ctx)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			if resp.StatusCode == http.StatusOK {
				PrintGetPolResult(resp, *restOptions)
				return
			}

		},
	}

	return GetPolCmd
}

func PrintGetPolResult(resp *http.Response, o xlbcli.RESTOptions) {
	Polresp := xlbcli.PolInformationGet{}
	var data [][]string
	resultByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: Failed to read HTTP response: (%s)\n", err.Error())
		return
	}

	if err := json.Unmarshal(resultByte, &Polresp); err != nil {
		fmt.Printf("Error: Failed to unmarshal HTTP response: (%s)\n", err.Error())
		return
	}

	// if json options enable, it print as a json format.
	if o.PrintOption == "json" {
		resultIndent, _ := json.MarshalIndent(Polresp, "", "    ")
		fmt.Println(string(resultIndent))
		return
	}

	Polresp.Sort()

	// Table Init
	table := TableInit()

	// Making load balance data
	for _, Pol := range Polresp.PolModInfo {
		if o.PrintOption == "wide" {
			table.SetHeader(POLICY_WIDE_TITLE)
			data = append(data, []string{Pol.Ident, fmt.Sprintf("%d", Pol.Info.PeakInfoRate), fmt.Sprintf("%d", Pol.Info.CommittedInfoRate),
				fmt.Sprintf("%d", Pol.Info.ExcessBlkSize), fmt.Sprintf("%d", Pol.Info.CommittedBlkSize),
				fmt.Sprintf("%d", Pol.Info.PolType), fmt.Sprintf("%t", Pol.Info.ColorAware),
				Pol.Target.PolObjName, fmt.Sprintf("%d", Pol.Target.AttachMent)})
		} else {
			table.SetHeader(POLICY_TITLE)
			data = append(data, []string{Pol.Ident, fmt.Sprintf("%d", Pol.Info.PeakInfoRate), fmt.Sprintf("%d", Pol.Info.CommittedInfoRate)})
		}
	}

	// Rendering the load balance data to table
	TableShow(data, table)
}

func PolicyAPICall(restOptions *xlbcli.RESTOptions) (*http.Response, error) {
	client := xlbcli.NewXlbClient(restOptions)
	ctx := context.TODO()
	var cancel context.CancelFunc
	if restOptions.Timeout > 0 {
		ctx, cancel = context.WithTimeout(context.TODO(), time.Duration(restOptions.Timeout)*time.Second)
		defer cancel()
	}
	resp, err := client.Policy().SetUrl("/config/policy/all").Get(ctx)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return nil, err
	}
	return resp, nil
}

func Poldump(restOptions *xlbcli.RESTOptions, path string) (string, error) {
	// File Open
	fileP := []string{"Polconfig_", ".txt"}
	t := time.Now()
	file := strings.Join(fileP, t.Local().Format("2006-01-02_15:04:05"))
	f, err := os.Create(file)
	if err != nil {
		fmt.Printf("Can't create dump file\n")
		os.Exit(1)
	}
	defer f.Close()

	// API Call
	resp, err := PolicyAPICall(restOptions)
	if err != nil {
		return "", err
	}
	resultByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: Failed to read HTTP response: (%s)\n", err.Error())
	}
	// Write
	f.Write(resultByte)

	cfile := path + "Polconfig.txt"
	if _, err := os.Stat(cfile); errors.Is(err, os.ErrNotExist) {
		if err != nil {
			fmt.Println("There is no saved config file")
		}
	} else {
		command := "mv " + cfile + " " + cfile + ".bk"
		cmd := exec.Command("bash", "-c", command)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("Can't backup ", cfile)
			return file, err
		}
	}
	command := "cp -R " + file + " " + cfile
	cmd := exec.Command("bash", "-c", command)
	fmt.Println(cmd)
	_, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	return file, nil
}

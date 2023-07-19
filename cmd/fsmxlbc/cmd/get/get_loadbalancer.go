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
	"github.com/cybwan/fsmxlb/pkg/xlbcli"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func NewGetLoadBalancerCmd(restOptions *xlbcli.RESTOptions) *cobra.Command {
	var GetLbCmd = &cobra.Command{
		Use:     "loadbalancer",
		Short:   "Get a LoadBalancer",
		Aliases: []string{"lb", "loadbalancers", "lbs"},
		Long:    `It shows Load balancer Information`,
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
			resp, err := client.LoadBalancerAll().Get(ctx)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			if resp.StatusCode == http.StatusOK {
				PrintGetLbResult(resp, *restOptions)
				return
			}

		},
	}

	return GetLbCmd
}

func NumToSelect(sel int) string {
	var ret string
	switch sel {
	case 0:
		ret = "rr"
	case 1:
		ret = "hash"
	case 2:
		ret = "priority"
	default:
		ret = "rr"
	}
	return ret
}

func NumToMode(mode int) string {
	var ret string
	switch mode {
	case 1:
		ret = "onearm"
	case 2:
		ret = "fullnat"
	case 3:
		ret = "dsr"
	default:
		ret = "default"
	}
	return ret
}

func BoolToMon(mon bool) string {
	var ret string
	switch mon {
	case false:
		ret = "Off"
	case true:
		ret = "On"
	}
	return ret
}

func PrintGetLbResult(resp *http.Response, o xlbcli.RESTOptions) {
	lbresp := xlbcli.LbRuleModGet{}
	var data [][]string
	var secIPs string
	resultByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: Failed to read HTTP response: (%s)\n", err.Error())
		return
	}

	if err := json.Unmarshal(resultByte, &lbresp); err != nil {
		fmt.Printf("Error: Failed to unmarshal HTTP response: (%s)\n", err.Error())
		return
	}

	// if json options enable, it print as a json format.
	if o.PrintOption == "json" {
		resultIndent, _ := json.MarshalIndent(lbresp, "", "    ")
		fmt.Println(string(resultIndent))
		return
	}

	lbresp.Sort()

	// Table Init
	table := TableInit()

	// Making load balance data
	for _, lbrule := range lbresp.LbRules {
		if o.PrintOption == "wide" {
			table.SetHeader(LOADBALANCER_WIDE_TITLE)
			secIPs = ""
			if len(lbrule.SecondaryIPs) > 0 {
				secIPs = lbrule.SecondaryIPs[0].SecondaryIP
				for i := 1; i < len(lbrule.SecondaryIPs); i++ {
					secIPs = secIPs + ", " + lbrule.SecondaryIPs[i].SecondaryIP
				}
			}

			if lbrule.Service.Monitor {
				for i, eps := range lbrule.Endpoints {
					if i == 0 {
						data = append(data, []string{lbrule.Service.ExternalIP, secIPs, fmt.Sprintf("%d", lbrule.Service.Port), lbrule.Service.Protocol, fmt.Sprintf("%d", lbrule.Service.Block), NumToSelect(int(lbrule.Service.Sel)), NumToMode(int(lbrule.Service.Mode)),
							eps.EndpointIP, fmt.Sprintf("%d", eps.TargetPort), fmt.Sprintf("%d", eps.Weight), eps.State})
					} else {
						data = append(data, []string{"", "", "", "", "", "", "", eps.EndpointIP, fmt.Sprintf("%d", eps.TargetPort), fmt.Sprintf("%d", eps.Weight), eps.State})
					}
				}
			} else {
				for i, eps := range lbrule.Endpoints {
					if i == 0 {
						data = append(data, []string{lbrule.Service.ExternalIP, secIPs, fmt.Sprintf("%d", lbrule.Service.Port), lbrule.Service.Protocol, fmt.Sprintf("%d", lbrule.Service.Block), NumToSelect(int(lbrule.Service.Sel)), NumToMode(int(lbrule.Service.Mode)),
							eps.EndpointIP, fmt.Sprintf("%d", eps.TargetPort), fmt.Sprintf("%d", eps.Weight), "-"})
					} else {
						data = append(data, []string{"", "", "", "", "", "", "", eps.EndpointIP, fmt.Sprintf("%d", eps.TargetPort), fmt.Sprintf("%d", eps.Weight), "-"})
					}
				}
			}
		} else {
			table.SetHeader(LOADBALANCER_TITLE)
			data = append(data, []string{lbrule.Service.ExternalIP, fmt.Sprintf("%d", lbrule.Service.Port), lbrule.Service.Protocol, fmt.Sprintf("%d", lbrule.Service.Block), NumToSelect(int(lbrule.Service.Sel)), NumToMode(int(lbrule.Service.Mode)), fmt.Sprintf("%d", len(lbrule.Endpoints)), BoolToMon(lbrule.Service.Monitor)})
		}
	}

	// Rendering the load balance data to table
	TableShow(data, table)
}

func LoadbalancerAPICall(restOptions *xlbcli.RESTOptions) (*http.Response, error) {
	client := xlbcli.NewXlbClient(restOptions)
	ctx := context.TODO()
	var cancel context.CancelFunc
	if restOptions.Timeout > 0 {
		ctx, cancel = context.WithTimeout(context.TODO(), time.Duration(restOptions.Timeout)*time.Second)
		defer cancel()
	}
	resp, err := client.LoadBalancerAll().Get(ctx)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return nil, err
	}
	return resp, nil
}

func Lbdump(restOptions *xlbcli.RESTOptions, path string) (string, error) {
	lbresp := xlbcli.LbRuleModGet{}
	dresp := xlbcli.LbRuleModGet{}
	// File Open
	fileP := []string{"lbconfig_", ".txt"}
	t := time.Now()
	file := strings.Join(fileP, t.Local().Format("2006-01-02_15:04:05"))
	f, err := os.Create(file)
	if err != nil {
		fmt.Printf("Can't create dump file\n")
		os.Exit(1)
	}
	defer f.Close()

	// API Call
	client := xlbcli.NewXlbClient(restOptions)
	ctx := context.TODO()
	var cancel context.CancelFunc
	if restOptions.Timeout > 0 {
		ctx, cancel = context.WithTimeout(context.TODO(), time.Duration(restOptions.Timeout)*time.Second)
		defer cancel()
	}
	resp, err := client.LoadBalancerAll().Get(ctx)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return "", err
	}

	resultByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: Failed to read HTTP response: (%s)\n", err.Error())
	}

	if err := json.Unmarshal(resultByte, &lbresp); err != nil {
		fmt.Printf("Error: Failed to unmarshal HTTP response: (%s)\n", err.Error())
		return "", err
	}

	for _, lbrule := range lbresp.LbRules {
		if !lbrule.Service.Managed {
			dresp.LbRules = append(dresp.LbRules, lbrule)
		}
	}

	dumpBytes, err := json.Marshal(dresp)
	if err != nil {
		fmt.Printf("Error: Failed to marshal dump LB rules (%s)\n", err.Error())
		return "", err
	}

	// Write
	_, err = f.Write(dumpBytes)
	if err != nil {
		fmt.Println("File write error")
	}
	cfile := path + "lbconfig.txt"
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
		fmt.Println("Failed copy file to", cfile)
		return file, err
	}
	return file, nil
}

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

func NewGetSessionCmd(restOptions *xlbcli.RESTOptions) *cobra.Command {
	var GetsessionCmd = &cobra.Command{
		Use:     "session",
		Short:   "Get a session",
		Aliases: []string{"session", "sessions"},
		Long:    `It shows Session Information`,
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
			resp, err := client.Session().SetUrl("/config/session/all").Get(ctx)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			if resp.StatusCode == http.StatusOK {
				PrintGetSessionResult(resp, *restOptions)
				return
			}

		},
	}

	return GetsessionCmd
}

func PrintGetSessionResult(resp *http.Response, o xlbcli.RESTOptions) {
	sessionresp := xlbcli.SessionInformationGet{}
	var data [][]string
	resultByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: Failed to read HTTP response: (%s)\n", err.Error())
		return
	}

	if err := json.Unmarshal(resultByte, &sessionresp); err != nil {
		fmt.Printf("Error: Failed to unmarshal HTTP response: (%s)\n", err.Error())
		return
	}

	// if json options enable, it print as a json format.
	if o.PrintOption == "json" {
		resultIndent, _ := json.MarshalIndent(sessionresp, "", "    ")
		fmt.Println(string(resultIndent))
		return
	}

	sessionresp.Sort()

	// Table Init
	table := TableInit()

	// Making load balance data
	for _, sessionrule := range sessionresp.SessionInfo {
		if o.PrintOption == "wide" {
			table.SetHeader(SESSION_WIDE_TITLE)
			data = append(data, []string{
				sessionrule.Ident,
				sessionrule.Ip.String(),
				fmt.Sprintf("TeID: %v TunnelIP: %s", sessionrule.AnTun.TeID, sessionrule.AnTun.Addr.String()),
				fmt.Sprintf("TeID: %v TunnelIP: %s", sessionrule.CnTun.TeID, sessionrule.CnTun.Addr.String()),
			})
		} else {
			table.SetHeader(SESSION_TITLE)
			data = append(data, []string{sessionrule.Ident, sessionrule.Ip.String()})
		}
	}

	// Rendering the load balance data to table
	TableShow(data, table)
}

func SessionAPICall(restOptions *xlbcli.RESTOptions) (*http.Response, error) {
	client := xlbcli.NewXlbClient(restOptions)
	ctx := context.TODO()
	var cancel context.CancelFunc
	if restOptions.Timeout > 0 {
		ctx, cancel = context.WithTimeout(context.TODO(), time.Duration(restOptions.Timeout)*time.Second)
		defer cancel()
	}
	resp, err := client.Session().SetUrl("/config/session/all").Get(ctx)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return nil, err
	}
	return resp, nil
}

func Sessiondump(restOptions *xlbcli.RESTOptions, path string) (string, error) {
	// File Open
	fileP := []string{"sessionconfig_", ".txt"}
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
	resp, err := client.Session().SetUrl("/config/session/all").Get(ctx)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return "", err
	}
	resultByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: Failed to read HTTP response: (%s)\n", err.Error())
	}
	// Write
	f.Write(resultByte)

	cfile := path + "sessionconfig.txt"
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

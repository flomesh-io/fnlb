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
	"context"
	"fmt"
	"github.com/cybwan/fsmxlb/pkg/xlbcli"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func DeleteMirrorValidation(args []string) error {
	if len(args) > 1 {
		fmt.Println("delete Mirror command get so many args")
	}
	return nil
}

func NewDeleteMirrorCmd(restOptions *xlbcli.RESTOptions) *cobra.Command {

	var deleteMirrorCmd = &cobra.Command{
		Use:     "mirror <MirrorIdent>",
		Short:   "Delete a Mirror",
		Long:    `Delete a Mirror using MirrorIdent in the fsmxlb.`,
		Aliases: []string{"mirror", "mirr", "mirrors"},
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := DeleteMirrorValidation(args); err != nil {
				fmt.Println("not valid <MirrorIdent>")
				return
			}
			MirrorIdent := args[0]

			client := xlbcli.NewXlbClient(restOptions)
			ctx := context.TODO()
			var cancel context.CancelFunc
			if restOptions.Timeout > 0 {
				ctx, cancel = context.WithTimeout(ctx, time.Duration(restOptions.Timeout)*time.Second)
				defer cancel()
			}
			subResources := []string{"ident", MirrorIdent}
			resp, err := client.Mirror().SubResources(subResources).Delete(ctx)
			if err != nil {
				fmt.Printf("Error: Failed to delete Mirror : %s", MirrorIdent)
				return
			}
			defer resp.Body.Close()
			fmt.Printf("Debug: response.StatusCode: %d\n", resp.StatusCode)
			if resp.StatusCode == http.StatusOK {
				PrintDeleteResult(resp, *restOptions)
				return
			}

		},
	}

	return deleteMirrorCmd
}

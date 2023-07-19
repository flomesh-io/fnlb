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
package handler

import (
	"fmt"
	"github.com/cybwan/fsmxlb/api/xlbnlp"
	"strings"

	"github.com/cybwan/fsmxlb/api/models"
	"github.com/cybwan/fsmxlb/api/restapi/operations"
	tk "github.com/cybwan/fsmxlb/pkg/xlblib"
	"github.com/go-openapi/runtime/middleware"
)

func ConfigPostRoute(params operations.PostConfigRouteParams) middleware.Responder {
	tk.LogIt(tk.LogDebug, "[API] Route  %s API called. url : %s\n", params.HTTPRequest.Method, params.HTTPRequest.URL)
	ret := xlbnlp.AddRouteNoHook(params.Attr.DestinationIPNet, params.Attr.Gateway)
	if ret != 0 {
		tk.LogIt(tk.LogDebug, "[API] Error occur : %v\n", ret)
		return &ResultResponse{Result: "fail"}
	}
	return &ResultResponse{Result: "Success"}
}

func ConfigDeleteRoute(params operations.DeleteConfigRouteDestinationIPNetIPAddressMaskParams) middleware.Responder {
	tk.LogIt(tk.LogDebug, "[API] Route  %s API called. url : %s\n", params.HTTPRequest.Method, params.HTTPRequest.URL)
	DstIP := fmt.Sprintf("%s/%d", params.IPAddress, params.Mask)
	ret := xlbnlp.DelRouteNoHook(DstIP)
	if ret != 0 {
		tk.LogIt(tk.LogDebug, "[API] Error occur : %v\n", ret)
		return &ResultResponse{Result: "fail"}
	}
	return &ResultResponse{Result: "Success"}
}

func ConfigGetRoute(params operations.GetConfigRouteAllParams) middleware.Responder {
	tk.LogIt(tk.LogDebug, "[API] Route  %s API called. url : %s\n", params.HTTPRequest.Method, params.HTTPRequest.URL)
	res, _ := ApiHooks.NetRouteGet()
	var result []*models.RouteGetEntry
	result = make([]*models.RouteGetEntry, 0)
	for _, route := range res {
		var tmpResult models.RouteGetEntry
		tmpResult.DestinationIPNet = route.Dst
		tmpResult.Flags = strings.TrimSpace(route.Flags)
		tmpResult.Gateway = route.Gw
		tmpResult.HardwareMark = int64(route.HardwareMark)
		tmpResult.Protocol = int64(route.Protocol)
		tmpResult.Sync = int64(route.Sync)

		tmpStats := new(models.RouteGetEntryStatistic)

		tmpBytes := int64(route.Statistic.Bytes)
		tmpStats.Bytes = &tmpBytes
		tmpPackets := int64(route.Statistic.Packets)
		tmpStats.Packets = &tmpPackets
		tmpResult.Statistic = tmpStats

		result = append(result, &tmpResult)
	}
	return operations.NewGetConfigRouteAllOK().WithPayload(&operations.GetConfigRouteAllOKBody{RouteAttr: result})
}

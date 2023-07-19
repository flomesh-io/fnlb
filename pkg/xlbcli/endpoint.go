package xlbcli

import (
	"fmt"
	"sort"
)

type EndPoint struct {
	CommonAPI
}

type EPInformationGet struct {
	EPInfo []EndPointGetEntry `json:"Attr"`
}

// EndPointState - Info related to a end-point state
type EndPointGetEntry struct {
	// HostName - hostname in CIDR
	HostName string `json:"hostName"`
	// Name - Endpoint Identifier
	Name string `json:"name"`
	// InActTries - No. of inactive probes to mark
	// an end-point inactive
	InActTries int `json:"inactiveReTries"`
	// ProbeType - Type of probe : "icmp","tcp", "udp", "sctp", "http"
	ProbeType string `json:"probeType"`
	// ProbeReq - Request string in case of http probe
	ProbeReq string `json:"probeReq"`
	// ProbeResp - Response string in case of http probe
	ProbeResp string `json:"probeResp"`
	// ProbeDuration - How frequently (in seconds) to check activity
	ProbeDuration uint32 `json:"probeDuration"`
	// ProbePort - Port to probe for connect type
	ProbePort uint16 `json:"probePort"`
	// MinDelay - Minimum delay in this end-point
	MinDelay string `json:"minDelay"`
	// AvgDelay - Average delay in this end-point
	AvgDelay string `json:"avgDelay"`
	// MaxDelay - Max delay in this end-point
	MaxDelay string `json:"maxDelay"`
	// CurrState - Current state of this end-point
	CurrState string `json:"currState"`
}

type EPConfig struct {
	EPInfo []EndPointMod `json:"Attr"`
}

// EndPointMod - Info related to a end-point config entry
type EndPointMod struct {
	// HostName - hostname in CIDR
	HostName string `json:"hostName" yaml:"hostName"`
	// Name - Endpoint Identifier
	Name string `json:"name" yaml:"name"`
	// InActTries - No. of inactive probes to mark
	// an end-point inactive
	InActTries int `json:"inactiveReTries" yaml:"inactiveReTries"`
	// ProbeType - Type of probe : "icmp","tcp", "udp", "sctp", "http"
	ProbeType string `json:"probeType" yaml:"probeType"`
	// ProbeReq - Request string in case of http probe
	ProbeReq string `json:"probeReq" yaml:"probeReq"`
	// ProbeResp - Response string in case of http probe
	ProbeResp string `json:"probeResp" yaml:"probeResp"`
	// ProbeDuration - How frequently (in seconds) to check activity
	ProbeDuration uint32 `json:"probeDuration" yaml:"probeDuration"`
	// ProbePort - Port to probe for connect type
	ProbePort uint16 `json:"probePort" yaml:"probePort"`
}

type ConfigurationEndPointFile struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
	Spec       EndPointMod `yaml:"spec"`
}

func (epResp EndPointGetEntry) Key() string {
	return fmt.Sprintf("%s|%05d|%s|%s", epResp.HostName, epResp.ProbePort, epResp.ProbeType, epResp.Name)
}

func (epResp EPInformationGet) Sort() {
	sort.Slice(epResp.EPInfo, func(i, j int) bool {
		return epResp.EPInfo[i].Key() < epResp.EPInfo[j].Key()
	})
}

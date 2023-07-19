package api

import (
	"github.com/cybwan/fsmxlb/api/restapi"
	"github.com/cybwan/fsmxlb/api/restapi/handler"
	"github.com/cybwan/fsmxlb/api/restapi/operations"
	cmn "github.com/cybwan/fsmxlb/pkg/common"
	"github.com/cybwan/fsmxlb/pkg/options"
	tk "github.com/cybwan/fsmxlb/pkg/xlblib"
	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"log"
	"os"
	"runtime/debug"
	"time"
)

var (
	ApiReady  bool
	ApiShutOk chan bool
)

// RegisterAPIHooks - routine to register interface for api
func RegisterAPIHooks(hooks cmn.NetHookInterface) {
	handler.ApiHooks = hooks
}

// WaitAPIServerReady - routine to wait till api server is up
func WaitAPIServerReady() {
	for {
		if ApiReady {
			time.Sleep(2 * time.Second)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func waitApiServerShutOk() {
	for {
		select {
		case <-ApiShutOk:
			return
		}
	}
}

// ApiServerShutOk - Notifier for server to be shutdown on signals
func ApiServerShutOk() {
	ApiShutOk <- true
}

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

// RunAPIServer - routine to start API server
func RunAPIServer() {

	// Stack trace logger
	defer func() {
		if e := recover(); e != nil {
			tk.LogIt(tk.LogCritical, "%s: %s", e, debug.Stack())
		}
	}()

	if ApiShutOk == nil {
		ApiShutOk = make(chan bool)
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewFsmxlbRestAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Fsmxlb Rest API"
	parser.LongDescription = "Fsmxlb REST API for Baremetal Scenarios"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()
	// API server host list
	server.Host = options.Opts.Host
	server.TLSHost = options.Opts.TLSHost
	server.TLSCertificateKey = options.Opts.TLSCertificateKey
	server.TLSCertificate = options.Opts.TLSCertificate
	server.Port = options.Opts.Port
	server.TLSPort = options.Opts.TLSPort
	api.ServerShutdown = func() {
		waitApiServerShutOk()
		os.Exit(0)
	}
	ApiReady = true

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
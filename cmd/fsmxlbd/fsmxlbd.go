package main

import (
	"fmt"
	opts "github.com/flomesh-io/fnlb/pkg/options"
	"github.com/flomesh-io/fnlb/pkg/version"
	ln "github.com/flomesh-io/fnlb/pkg/xlbnet"
	"github.com/jessevdk/go-flags"
	"os"
	"time"
)

func main() {
	fmt.Printf("Starting fsm-xlb %s %s %s\n", version.Version, version.GitCommit, version.BuildDate)

	// Parse command-line arguments
	_, err := flags.Parse(&opts.Opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if opts.Opts.Version {
		fmt.Printf("fsm-xlb %s %s %s\n", version.Version, version.GitCommit, version.BuildDate)
		os.Exit(0)
	}

	go ln.XlbXsyncMain()
	// Need some time for RPC Handler to be up
	time.Sleep(2 * time.Second)

	ln.Main()
}

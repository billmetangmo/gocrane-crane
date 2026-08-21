package main

import (
	"flag"
	"fmt"
	"os"

	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"

	"github.com/gocrane/crane/cmd/craned/app"
)

// craned main.
func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	klog.InitFlags(flag.CommandLine)

	ctx := signals.SetupSignalHandler()

	if err := app.NewManagerCommand(ctx).Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

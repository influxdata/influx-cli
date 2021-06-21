package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/influxdata/influx-cli/v2/pkg/signals"
	"github.com/urfave/cli/v2"
)

// Fields set via ldflags at build time.
var (
	version = "dev"
	commit  = "none"
	date    = ""
)

func init() {
	if len(date) == 0 {
		date = time.Now().UTC().Format(time.RFC3339)
	}

	cli.VersionFlag = nil
}

var app = cli.App{
	Name:                 "influx",
	Usage:                "Influx Client",
	UsageText:            "influx [command]",
	EnableBashCompletion: true,
	Commands: []*cli.Command{
		newVersionCmd(),
		newPingCmd(),
		newSetupCmd(),
		newWriteCmd(),
		newBucketCmd(),
		newCompletionCmd(),
		newBucketSchemaCmd(),
		newQueryCmd(),
		newConfigCmd(),
		newOrgCmd(),
		newDeleteCmd(),
		newUserCmd(),
		newTaskCommand(),
		newBackupCmd(),
		newRestoreCmd(),
		newTelegrafsCommand(),
		newDashboardsCommand(),
		newExportCmd(),
		newSecretCommand(),
	},
}

func main() {
	ctx := signals.WithStandardSignals(context.Background())
	if err := app.RunContext(ctx, os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

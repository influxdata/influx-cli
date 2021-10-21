package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
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
	BashComplete:         cli.DefaultAppComplete,
	Commands: []cli.Command{
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
		newV1SubCommand(),
		newAuthCommand(),
		newApplyCmd(),
		newStacksCmd(),
		newTemplateCmd(),
		newRemoteCmd(),
		newReplicationCmd(),
	},
	Before: withContext(),
	CommandNotFound: func(c *cli.Context, command string) {
		_, _ = fmt.Fprintf(
			os.Stderr,
			"Error: command %q not recognized. Run `%v --help` to see the list of commands\n",
			command,
			os.Args[0],
		)
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

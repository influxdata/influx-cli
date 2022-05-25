package main

import (
	"fmt"
	"os"
	"time"

	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
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

func cloudOnlyCommands() []cli.Command {
	// Include commands that are only intended to work on an InfluxDB Cloud host in this list. A specific error message
	// will be returned if these commands are run on an InfluxDB OSS host.
	cmds := []cli.Command{
		newBucketSchemaCmd(),
		newScriptsCmd(),
	}

	return middleware.AddMWToCmds(cmds, middleware.CloudOnly)
}

func ossOnlyCommands() []cli.Command {
	// Include commands that are only intended to work on an InfluxDB OSS host in this list. A specific error message will
	// be returned if these commands are run on an InfluxDB Cloud host.
	cmds := []cli.Command{
		newPingCmd(),
		newSetupCmd(),
		newBackupCmd(),
		newRestoreCmd(),
		newRemoteCmd(),
		newReplicationCmd(),
		newServerConfigCommand(),
	}

	return middleware.AddMWToCmds(cmds, middleware.OSSOnly)
}

func allCommands() []cli.Command {
	// Commands which should work with an InfluxDB Cloud or InfluxDB OSS host should be included in this list.
	commonCmds := []cli.Command{
		newVersionCmd(),
		newWriteCmd(),
		newBucketCmd(),
		newCompletionCmd(),
		newQueryCmd(),
		newConfigCmd(),
		newOrgCmd(),
		newDeleteCmd(),
		newUserCmd(),
		newTaskCommand(),
		newTelegrafsCommand(),
		newDashboardsCommand(),
		newExportCmd(),
		newSecretCommand(),
		newV1SubCommand(),
		newAuthCommand(),
		newApplyCmd(),
		newStacksCmd(),
		newTemplateCmd(),
	}
	specificCmds := append(cloudOnlyCommands(), ossOnlyCommands()...)

	return append(commonCmds, specificCmds...)
}

func newApp() cli.App {
	return cli.App{
		Name:                 "influx",
		Usage:                "Influx Client",
		UsageText:            "influx [command]",
		EnableBashCompletion: true,
		BashComplete:         cli.DefaultAppComplete,
		Commands:             allCommands(),
		Before:               middleware.WithBeforeFns(withContext(), middleware.NoArgs),
		ExitErrHandler:       middleware.HandleExit,
	}
}

func main() {
	app := newApp()
	if err := app.Run(os.Args); err != nil {
		// Errors will normally be handled by cli.HandleExitCoder via ExitErrHandler set on app. Any error not implementing
		// the cli.ExitCoder interface can be handled here.
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

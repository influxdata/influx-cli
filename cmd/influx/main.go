package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
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
		UsageText:            "influx [command]\n\nHINT: If you are looking for the InfluxQL shell from 1.x, run \"influx v1 shell\"",
		EnableBashCompletion: true,
		BashComplete:         cli.DefaultAppComplete,
		Commands:             allCommands(),
		Before:               middleware.WithBeforeFns(withContext(), middleware.NoArgs),
		ExitErrHandler:       middleware.HandleExit,
	}
}

// This creates a new slice and replaces `-t "-FOO-TOKEN"` with `-t=-FOO-TOKEN`
// This is necessary to do because the command line arg:
//  `-t "-FOO-TOKEN"`` will be parsed as two separate flags instead of a flag and token value.
func ReplaceTokenArg(args []string) []string {
	if len(args) == 0 {
		return []string{}
	}
	newArgs := make([]string, len(args))
	copy(newArgs, args)
	for i, arg := range newArgs[:len(newArgs)-1] {
		switch arg {
		case "--token", "-t":
			if strings.HasPrefix(newArgs[i+1], "-") {
				color.HiYellow("warning: %[1]s %[2]s interpreted as %[1]s=%[2]s, consider using %[1]s=%[2]s syntax when tokens start with a hyphen",
					newArgs[i], newArgs[i+1],
				)
			}
			newArgs[i] = strings.Join(newArgs[i:i+2], "=")
			newArgs = append(newArgs[:i+1], newArgs[i+2:]...)
		}
	}
	return newArgs
}

func main() {
	app := newApp()
	args := ReplaceTokenArg(os.Args)
	if err := app.Run(args); err != nil {
		// Errors will normally be handled by cli.HandleExitCoder via ExitErrHandler set on app. Any error not implementing
		// the cli.ExitCoder interface can be handled here.
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
	"os"
	"strings"
	"time"
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
		UsageText:            "influx [command]",
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
func replaceTokenArg(args []string) []string {
	newArgs := make([]string, len(args))
	copy(newArgs, args)
	for i, arg := range newArgs {
		switch arg {
		case "--token", "-t":
			// if last element, this will be invalid later
			if i == len(args)-1 {
				break
			}
			if strings.HasPrefix(newArgs[i+1], "-") {
				color.HiYellow("warning: %s %s interpreted as %s=%s, consider using %s=%s syntax when tokens start with a hyphen",
					newArgs[i], newArgs[i+1],
					newArgs[i], newArgs[i+1],
					newArgs[i], newArgs[i+1],
				)
			}
			newArgs[i] = strings.Join(newArgs[i:i+2], "=")
			// if there are 2+ elements after this
			if len(newArgs) > i+2 {
				newArgs = append(newArgs[:i+1], newArgs[i+2:]...)
			} else {
				newArgs = newArgs[:i+1]
			}
		}
	}
	return newArgs
}

func main() {
	app := newApp()
	args := replaceTokenArg(os.Args)
	if err := app.Run(args); err != nil {
		// Errors will normally be handled by cli.HandleExitCoder via ExitErrHandler set on app. Any error not implementing
		// the cli.ExitCoder interface can be handled here.
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

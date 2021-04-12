package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
	commit  = "none"
	date    = ""
)

func main() {
	if len(date) == 0 {
		date = time.Now().UTC().Format(time.RFC3339)
	}

	cli.VersionFlag = nil

	app := cli.App{
		Name:      "influx",
		Usage:     "Influx Client",
		UsageText: "influx [command]",
		Commands: []*cli.Command{
			{
				Name:  "version",
				Usage: "Print the influx CLI version",
				Action: func(*cli.Context) error {
					fmt.Printf("Influx CLI %s (git: %s) build_date: %s\n", version, commit, date)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}

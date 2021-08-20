package main

import (
	"fmt"

	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newVersionCmd() cli.Command {
	return cli.Command{
		Name:   "version",
		Usage:  "Print the influx CLI version",
		Before: middleware.NoArgs,
		Action: func(*cli.Context) error {
			fmt.Printf("Influx CLI %s (git: %s) build_date: %s\n", version, commit, date)
			return nil
		},
	}
}

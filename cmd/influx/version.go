package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func newVersionCmd() cli.Command {
	return cli.Command{
		Name:  "version",
		Usage: "Print the influx CLI version",
		Action: func(*cli.Context) error {
			fmt.Printf("Influx CLI %s (git: %s) build_date: %s\n", version, commit, date)
			return nil
		},
	}
}

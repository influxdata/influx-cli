package main

import "github.com/urfave/cli/v2"

func newV1DBRPCmd() *cli.Command {
	return &cli.Command{
		Name:        "dbrp",
		Usage:       "Commands to manage database and retention policy mappings for v1 APIs",
		Subcommands: []*cli.Command{
			// newV1DBRPListCmd(),
			// newV1DBRPCreateCmd(),
			// newV1DBRPDeleteCmd(),
			// newV1DBRPUpdateCmd(),
		},
	}
}

// commands will be implemented below, with business logic residing in their
// respective client package, like "v1_dbrp"

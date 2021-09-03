package main

import (
	"fmt"

	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newReplicationCmd() cli.Command {
	return cli.Command{
		Name:   "replication",
		Usage:  "Replication stream management commands",
		Hidden: true, // Remove this line when all subcommands are completed
		Subcommands: []cli.Command{
			newReplicationCreateCmd(),
			newReplicationDeleteCmd(),
			newReplicationListCmd(),
			newReplicationUpdateCmd(),
		},
	}
}

func newReplicationCreateCmd() cli.Command {
	return cli.Command{
		Name:   "create",
		Usage:  "Create a new replication stream",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags:  commonFlags(),
		Action: func(ctx *cli.Context) {
			fmt.Println("replication create command was called")
		},
	}
}

func newReplicationDeleteCmd() cli.Command {
	return cli.Command{
		Name:   "delete",
		Usage:  "Delete an existing replication stream",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags:  commonFlags(),
		Action: func(ctx *cli.Context) {
			fmt.Println("replication delete command was called")
		},
	}
}

func newReplicationListCmd() cli.Command {
	return cli.Command{
		Name:    "list",
		Usage:   "List all replication streams and corresponding metrics",
		Aliases: []string{"find", "ls"},
		Before:  middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags:   commonFlags(),
		Action: func(ctx *cli.Context) {
			fmt.Println("replication list command was called")
		},
	}
}

func newReplicationUpdateCmd() cli.Command {
	return cli.Command{
		Name:   "update",
		Usage:  "Update an existing replication stream",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags:  commonFlags(),
		Action: func(ctx *cli.Context) {
			fmt.Println("replication update command was called")
		},
	}
}

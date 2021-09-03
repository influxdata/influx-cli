package main

import (
	"fmt"

	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newRemoteCmd() cli.Command {
	return cli.Command{
		Name:   "remote",
		Usage:  "Remote connection management commands",
		Hidden: true, // Remove this line when all subcommands are completed
		Subcommands: []cli.Command{
			newRemoteCreateCmd(),
			newRemoteDeleteCmd(),
			newRemoteListCmd(),
			newRemoteUpdateCmd(),
		},
	}
}

func newRemoteCreateCmd() cli.Command {
	return cli.Command{
		Name:   "create",
		Usage:  "Create a new remote connection",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags:  commonFlags(),
		Action: func(ctx *cli.Context) {
			fmt.Println("remote create command was called")
		},
	}
}

func newRemoteDeleteCmd() cli.Command {
	return cli.Command{
		Name:   "delete",
		Usage:  "Delete an existing remote connection",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags:  commonFlags(),
		Action: func(ctx *cli.Context) {
			fmt.Println("remote delete command was called")
		},
	}
}

func newRemoteListCmd() cli.Command {
	return cli.Command{
		Name:    "list",
		Usage:   "List all remote connections",
		Aliases: []string{"find", "ls"},
		Before:  middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags:   commonFlags(),
		Action: func(ctx *cli.Context) {
			fmt.Println("remote list command was called")
		},
	}
}

func newRemoteUpdateCmd() cli.Command {
	return cli.Command{
		Name:   "update",
		Usage:  "Update an existing remote connection",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags:  commonFlags(),
		Action: func(ctx *cli.Context) {
			fmt.Println("remote update command was called")
		},
	}
}

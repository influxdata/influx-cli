package main

import (
	"fmt"

	"github.com/influxdata/influx-cli/v2/clients/remote"
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
	var params remote.CreateParams
	return cli.Command{
		Name:   "create",
		Usage:  "Create a new remote connection",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "Name for the new remote connection",
				Required:    true,
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "description, d",
				Usage:       "Description for the new remote connection",
				Destination: &params.Description,
			},
			&cli.StringFlag{
				Name:        "org-id",
				Usage:       "The ID of the organization",
				EnvVar:      "INFLUX_ORG_ID",
				Destination: &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "remote-url",
				Usage:       "The url for the remote database",
				Destination: &params.RemoteURL,
			},
			&cli.StringFlag{
				Name:        "remote-api-token",
				Usage:       "The API token for the remote database",
				Destination: &params.RemoteAPIToken,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := remote.Client{
				CLI:              getCLI(ctx),
				OrganizationsApi: getAPI(ctx).OrganizationsApi,
				RemoteConnectionsApi: getAPI(ctx).RemoteConnectionsApi,
			}
			return client.Create(getContext(ctx), &params)
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

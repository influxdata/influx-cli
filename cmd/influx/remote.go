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
				Usage:       "The ID of the local organization",
				EnvVar:      "INFLUX_ORG_ID",
				Destination: &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "org, o",
				Usage:       "The name of the organization",
				EnvVar:      "INFLUX_ORG",
				Destination: &params.OrgName,
			},
			&cli.StringFlag{
				Name:        "remote-url",
				Usage:       "The url for the remote database",
				Required:    true,
				Destination: &params.RemoteURL,
			},
			&cli.StringFlag{
				Name:        "remote-api-token",
				Usage:       "The API token for the remote database",
				Required:    true,
				Destination: &params.RemoteAPIToken,
			},
			&cli.StringFlag{
				Name:        "remote-org-id",
				Usage:       "The ID of the remote organization",
				Required:    true,
				Destination: &params.RemoteOrgID,
			},
			&cli.BoolFlag{
				Name:        "allow-insecure-tls",
				Usage:       "Allows insecure TLS",
				Destination: &params.AllowInsecureTLS,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := remote.Client{
				CLI:                  getCLI(ctx),
				RemoteConnectionsApi: getAPI(ctx).RemoteConnectionsApi,
				OrganizationsApi:     getAPI(ctx).OrganizationsApi,
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
	var params remote.ListParams
	return cli.Command{
		Name:    "list",
		Usage:   "List all remote connections",
		Aliases: []string{"find", "ls"},
		Before:  middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "Name filter for remote connections list",
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "org-id",
				Usage:       "Local org ID",
				EnvVar:      "INFLUX_ORG_ID",
				Destination: &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "org, o",
				Usage:       "Local org name",
				EnvVar:      "INFLUX_ORG",
				Destination: &params.OrgName,
			},
			&cli.StringFlag{
				Name:        "remote-url",
				Usage:       "Remote URL filter for remote connections list",
				Destination: &params.RemoteURL,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := remote.Client{
				CLI:                  getCLI(ctx),
				RemoteConnectionsApi: getAPI(ctx).RemoteConnectionsApi,
				OrganizationsApi:     getAPI(ctx).OrganizationsApi,
			}

			return client.List(getContext(ctx), &params)
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

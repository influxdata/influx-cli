package main

import (
	"github.com/influxdata/influx-cli/v2/clients/remote"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newRemoteCmd() cli.Command {
	return cli.Command{
		Name:   "remote",
		Usage:  "Remote connection management commands",
		Before: middleware.NoArgs,
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
			append(commonFlags(), getOrgFlags(&params.OrgParams)...),
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
			if err := checkOrgFlags(&params.OrgParams); err != nil {
				return err
			}
			api := getAPI(ctx)

			client := remote.Client{
				CLI:                  getCLI(ctx),
				RemoteConnectionsApi: api.RemoteConnectionsApi,
				OrganizationsApi:     api.OrganizationsApi,
			}

			return client.Create(getContext(ctx), &params)
		},
	}
}

func newRemoteDeleteCmd() cli.Command {
	var remoteID string
	return cli.Command{
		Name:   "delete",
		Usage:  "Delete an existing remote connection",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "id, i",
				Usage:       "ID of the remote connection to be deleted",
				Required:    true,
				Destination: &remoteID,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)

			client := remote.Client{
				CLI:                  getCLI(ctx),
				RemoteConnectionsApi: api.RemoteConnectionsApi,
			}

			return client.Delete(getContext(ctx), remoteID)
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
				Usage:       "Filter results to only connections with a specific name",
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
				Usage:       "Filter results to only connections for a specific remote URL",
				Destination: &params.RemoteURL,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)

			client := remote.Client{
				CLI:                  getCLI(ctx),
				RemoteConnectionsApi: api.RemoteConnectionsApi,
				OrganizationsApi:     api.OrganizationsApi,
			}

			return client.List(getContext(ctx), &params)
		},
	}
}

func newRemoteUpdateCmd() cli.Command {
	var params remote.UpdateParams
	return cli.Command{
		Name:   "update",
		Usage:  "Update an existing remote connection",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "id, i",
				Usage:       "Remote connection ID",
				Required:    true,
				Destination: &params.RemoteID,
			},
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "New name for the remote connection",
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "description, d",
				Usage:       "New description for the remote connection",
				Destination: &params.Description,
			},
			&cli.StringFlag{
				Name:        "remote-url",
				Usage:       "New url for the remote database",
				Destination: &params.RemoteURL,
			},
			&cli.StringFlag{
				Name:        "remote-api-token",
				Usage:       "New API token for the remote database",
				Destination: &params.RemoteAPIToken,
			},
			&cli.StringFlag{
				Name:        "remote-org-id",
				Usage:       "New ID of the remote organization",
				Destination: &params.RemoteOrgID,
			},
			&cli.BoolFlag{
				Name:        "allow-insecure-tls",
				Usage:       "Allows insecure TLS",
				Destination: &params.AllowInsecureTLS,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)

			client := remote.Client{
				CLI:                  getCLI(ctx),
				RemoteConnectionsApi: api.RemoteConnectionsApi,
			}

			params.TLSFlagIsSet = ctx.IsSet("allow-insecure-tls")

			return client.Update(getContext(ctx), &params)
		},
	}
}

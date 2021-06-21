package main

import (
	"github.com/influxdata/influx-cli/v2/clients/secret"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli/v2"
)

func newSecretCommand() *cli.Command {
	return &cli.Command{
		Name:  "secret",
		Usage: "Secret management commands",
		Subcommands: []*cli.Command{
			newDeleteSecretCmd(),
			newListSecretCmd(),
			newUpdateSecretCmd(),
		},
	}
}

func newDeleteSecretCmd() *cli.Command {
	var params secret.DeleteParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	flags = append(flags, &cli.StringFlag{
		Name:        "key",
		Usage:       "The secret key (required)",
		Required:    true,
		Aliases:     []string{"k"},
		Destination: &params.Key,
	})
	return &cli.Command{
		Name:   "delete",
		Usage:  "Delete secret",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := secret.Client{
				CLI:              getCLI(ctx),
				SecretsApi:       api.SecretsApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Delete(ctx.Context, &params)
		},
	}
}

func newListSecretCmd() *cli.Command {
	var params secret.ListParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	return &cli.Command{
		Name:    "list",
		Usage:   "List secrets",
		Aliases: []string{"find", "ls"},
		Flags:   flags,
		Before:  middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := secret.Client{
				CLI:              getCLI(ctx),
				SecretsApi:       api.SecretsApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.List(ctx.Context, &params)
		},
	}
}

func newUpdateSecretCmd() *cli.Command {
	var params secret.UpdateParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	flags = append(flags,
		&cli.StringFlag{
			Name:        "key",
			Usage:       "The secret key (required)",
			Required:    true,
			Aliases:     []string{"k"},
			Destination: &params.Key,
		},
		&cli.StringFlag{
			Name:        "value",
			Usage:       "Optional secret value for scripting convenience, using this might expose the secret to your local history",
			Aliases:     []string{"v"},
			Destination: &params.Value,
		},
	)
	return &cli.Command{
		Name:   "update",
		Usage:  "Update secret",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := secret.Client{
				CLI:              getCLI(ctx),
				SecretsApi:       api.SecretsApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Update(ctx.Context, &params)
		},
	}
}

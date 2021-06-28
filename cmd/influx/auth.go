package main

import (
	"github.com/influxdata/influx-cli/v2/clients/auth"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newAuthCommand() cli.Command {
	return cli.Command{
		Name:    "auth",
		Usage:   "Authorization management commands",
		Aliases: []string{"authorization"},
		Subcommands: []cli.Command{
			newCreateCommand(),
			newDeleteCommand(),
			newListCommand(),
			newSetActiveCommand(),
			newSetInactiveCommand(),
		},
	}
}

func newCreateCommand() cli.Command {
	var params auth.CreateParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	flags = append(flags,
		&cli.StringFlag{
			Name:        "user, u",
			Usage:       "The user name",
			Destination: &params.User,
		},
		&cli.StringFlag{
			Name:        "description, d",
			Usage:       "Token description",
			Destination: &params.Description,
		},

		&cli.BoolFlag{
			Name:        "write-user",
			Usage:       "Grants the permission to perform mutative actions against organization users",
			Destination: &params.WriteUserPermission,
		},
		&cli.BoolFlag{
			Name:        "read-user",
			Usage:       "Grants the permission to perform read actions against organization users",
			Destination: &params.ReadUserPermission,
		},
		&cli.BoolFlag{
			Name:        "write-buckets",
			Usage:       "Grants the permission to perform mutative actions against organization buckets",
			Destination: &params.WriteBucketsPermission,
		},
		&cli.BoolFlag{
			Name:        "read-buckets",
			Usage:       "Grants the permission to perform read actions against organization buckets",
			Destination: &params.ReadBucketsPermission,
		},
		&cli.StringSliceFlag{
			Name:  "write-bucket",
			Usage: "The bucket id",
		},
		&cli.StringSliceFlag{
			Name:  "read-bucket",
			Usage: "The bucket id",
		},
		&cli.BoolFlag{
			Name:        "write-tasks",
			Usage:       "Grants the permission to create tasks",
			Destination: &params.WriteTasksPermission,
		},
		&cli.BoolFlag{
			Name:        "read-tasks",
			Usage:       "Grants the permission to read tasks",
			Destination: &params.ReadTasksPermission,
		},
		&cli.BoolFlag{
			Name:        "write-telegrafs",
			Usage:       "Grants the permission to create telegraf configs",
			Destination: &params.WriteTelegrafsPermission,
		},
		&cli.BoolFlag{
			Name:        "read-telegrafs",
			Usage:       "Grants the permission to read telegraf configs",
			Destination: &params.ReadTelegrafsPermission,
		},
		&cli.BoolFlag{
			Name:        "write-orgs",
			Usage:       "Grants the permission to create organizations",
			Destination: &params.WriteOrganizationsPermission,
		},
		&cli.BoolFlag{
			Name:        "read-orgs",
			Usage:       "Grants the permission to read organizations",
			Destination: &params.ReadOrganizationsPermission,
		},
		&cli.BoolFlag{
			Name:        "write-dashboards",
			Usage:       "Grants the permission to create dashboards",
			Destination: &params.WriteDashboardsPermission,
		},
		&cli.BoolFlag{
			Name:        "read-dashboards",
			Usage:       "Grants the permission to read dashboards",
			Destination: &params.ReadDashboardsPermission,
		},
		&cli.BoolFlag{
			Name:        "write-checks",
			Usage:       "Grants the permission to create checks",
			Destination: &params.WriteCheckPermission,
		},
		&cli.BoolFlag{
			Name:        "read-checks",
			Usage:       "Grants the permission to read checks",
			Destination: &params.ReadCheckPermission,
		},
		&cli.BoolFlag{
			Name:        "write-notificationRules",
			Usage:       "Grants the permission to create notificationRules",
			Destination: &params.WriteNotificationRulePermission,
		},
		&cli.BoolFlag{
			Name:        "read-notificationRules",
			Usage:       "Grants the permission to read notificationRules",
			Destination: &params.ReadNotificationRulePermission,
		},
		&cli.BoolFlag{
			Name:        "write-notificationEndpoints",
			Usage:       "Grants the permission to create notificationEndpoints",
			Destination: &params.WriteNotificationEndpointPermission,
		},
		&cli.BoolFlag{
			Name:        "read-notificationEndpoints",
			Usage:       "Grants the permission to read notificationEndpoints",
			Destination: &params.ReadNotificationEndpointPermission,
		},
		&cli.BoolFlag{
			Name:        "write-dbrps",
			Usage:       "Grants the permission to create database retention policy mappings",
			Destination: &params.WriteDBRPPermission,
		},
		&cli.BoolFlag{
			Name:        "read-dbrps",
			Usage:       "Grants the permission to read database retention policy mappings",
			Destination: &params.ReadDBRPPermission,
		},
	)
	return cli.Command{
		Name:   "create",
		Usage:  "Create authorization",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			params.WriteBucketIds = ctx.StringSlice("write-bucket")
			params.ReadBucketIds = ctx.StringSlice("read-bucket")

			api := getAPI(ctx)
			client := auth.Client{
				CLI:               getCLI(ctx),
				AuthorizationsApi: api.AuthorizationsApi,
				UsersApi:          api.UsersApi,
				OrganizationsApi:  api.OrganizationsApi,
			}
			return client.Create(getContext(ctx), &params)
		},
	}
}

func newDeleteCommand() cli.Command {
	return cli.Command{
		Name:  "delete",
		Usage: "Delete authorization",
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:     "id, i",
				Usage:    "The authorization ID (required)",
				Required: true,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := auth.Client{
				CLI:               getCLI(ctx),
				AuthorizationsApi: api.AuthorizationsApi,
				UsersApi:          api.UsersApi,
				OrganizationsApi:  api.OrganizationsApi,
			}
			return client.Remove(getContext(ctx), ctx.String("id"))
		},
	}
}

func newListCommand() cli.Command {
	var params auth.ListParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	flags = append(flags,
		&cli.StringFlag{
			Name:        "id, i",
			Usage:       "The authorization ID",
			Destination: &params.Id,
		},
		&cli.StringFlag{
			Name:        "user, u",
			Usage:       "The user",
			Destination: &params.User,
		},
		&cli.StringFlag{
			Name:        "user-id",
			Usage:       "The user ID",
			Destination: &params.UserID,
		},
	)
	return cli.Command{
		Name:    "list",
		Usage:   "List authorizations",
		Aliases: []string{"find", "ls"},
		Flags:   flags,
		Before:  middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := auth.Client{
				CLI:               getCLI(ctx),
				AuthorizationsApi: api.AuthorizationsApi,
				UsersApi:          api.UsersApi,
				OrganizationsApi:  api.OrganizationsApi,
			}
			return client.List(getContext(ctx), &params)
		},
	}
}

func newSetActiveCommand() cli.Command {
	return cli.Command{
		Name:  "active",
		Usage: "Active authorization",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id, i",
				Usage:    "The authorization ID (required)",
				Required: true,
			},
		},
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := auth.Client{
				CLI:               getCLI(ctx),
				AuthorizationsApi: api.AuthorizationsApi,
				UsersApi:          api.UsersApi,
				OrganizationsApi:  api.OrganizationsApi,
			}
			return client.SetActive(getContext(ctx), ctx.String("id"), true)
		},
	}
}

func newSetInactiveCommand() cli.Command {
	return cli.Command{
		Name:  "inactive",
		Usage: "Inactive authorization",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id, i",
				Usage:    "The authorization ID (required)",
				Required: true,
			},
		},
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := auth.Client{
				CLI:               getCLI(ctx),
				AuthorizationsApi: api.AuthorizationsApi,
				UsersApi:          api.UsersApi,
				OrganizationsApi:  api.OrganizationsApi,
			}
			return client.SetActive(getContext(ctx), ctx.String("id"), false)
		},
	}
}

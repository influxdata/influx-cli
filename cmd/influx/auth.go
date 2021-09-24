package main

import (
	"fmt"

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

func helpText(perm string) struct{ readHelp, writeHelp string } {
	var helpOverrides = map[string]struct{ readHelp, writeHelp string }{
		"user":      {"perform read actions against organization users", "perform mutative actions against organization users"},
		"buckets":   {"perform mutative actions against organization buckets", "perform mutative actions against organization buckets"},
		"telegrafs": {"read telegraf configs", "create telegraf configs"},
		"orgs":      {"read organizations", "create organizations"},
		"dbrps":     {"read database retention policy mappings", "create database retention policy mappings"},
	}

	help := helpOverrides[perm]
	if help.readHelp == "" {
		help.readHelp = fmt.Sprintf("read %s", perm)
	}
	if help.writeHelp == "" {
		help.writeHelp = fmt.Sprintf("create or update %s", perm)
	}

	help.readHelp = "Grants the permission to " + help.readHelp
	help.writeHelp = "Grants the permission to " + help.writeHelp
	return help
}

func newCreateCommand() cli.Command {
	var params auth.CreateParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)

	// default: create and update foo / reed foo
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
		&cli.StringSliceFlag{
			Name:  "write-bucket",
			Usage: "The bucket id",
		},
		&cli.StringSliceFlag{
			Name:  "read-bucket",
			Usage: "The bucket id",
		},
		&cli.BoolFlag{
			Name:        "operator",
			Usage:       "Grants all permissions in all organizations",
			Destination: &params.OperatorPermission,
		},
		&cli.BoolFlag{
			Name:        "all-access",
			Usage:       "Grants all permissions in a single organization",
			Destination: &params.AllAccess,
		},
	)

	params.ResourcePermissions = auth.BuildResourcePermissions()
	for _, perm := range params.ResourcePermissions {
		help := helpText(perm.Name)
		ossVsCloud := ""
		if perm.IsCloud && !perm.IsOss {
			ossVsCloud = " (Cloud only)"
		}
		if !perm.IsCloud && perm.IsOss {
			ossVsCloud = " (OSS only)"
		}
		flags = append(flags,
			&cli.BoolFlag{
				Name:        "read-" + perm.Name,
				Usage:       help.readHelp + ossVsCloud,
				Destination: &perm.Read,
			},
			&cli.BoolFlag{
				Name:        "write-" + perm.Name,
				Usage:       help.writeHelp + ossVsCloud,
				Destination: &perm.Write,
			})
	}

	return cli.Command{
		Name:   "create",
		Usage:  "Create authorization",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			params.WriteBucketIds = ctx.StringSlice("write-bucket")
			params.ReadBucketIds = ctx.StringSlice("read-bucket")

			api := getAPI(ctx)
			client := auth.Client{
				CLI:               getCLI(ctx),
				AuthorizationsApi: api.AuthorizationsApi,
				UsersApi:          api.UsersApi,
				OrganizationsApi:  api.OrganizationsApi,
				ResourceListApi:   api.ResourceListApi,
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
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
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
		Before:  middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
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
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
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
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
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

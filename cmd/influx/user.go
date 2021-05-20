package main

import (
	"github.com/influxdata/influx-cli/v2/internal/cmd/user"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/urfave/cli/v2"
)

func newUserCmd() *cli.Command {
	return &cli.Command{
		Name:  "user",
		Usage: "User management commands",
		Subcommands: []*cli.Command{
			newUserCreateCmd(),
			newUserDeleteCmd(),
			newUserListCmd(),
			newUserUpdateCmd(),
			newUserSetPasswordCmd(),
		},
	}
}

func newUserCreateCmd() *cli.Command {
	var params user.CreateParams
	return &cli.Command{
		Name:  "create",
		Usage: "Create user",
		Flags: append(
			commonFlags(),
			&cli.GenericFlag{
				Name:    "org-id",
				Usage:   "The ID of the organization",
				EnvVars: []string{"INFLUX_ORG_ID"},
				Value:   &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "org",
				Usage:       "The name of the organization",
				Aliases:     []string{"o"},
				EnvVars:     []string{"INFLUX_ORG"},
				Destination: &params.OrgName,
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "The user name",
				Aliases:     []string{"n"},
				EnvVars:     []string{"INFLUX_NAME"},
				Required:    true,
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "password",
				Usage:       "The user password",
				Aliases:     []string{"p"},
				Destination: &params.Password,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := user.Client{
				CLI:              getCLI(ctx),
				UsersApi:         api.UsersApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Create(ctx.Context, &params)
		},
	}
}

func newUserDeleteCmd() *cli.Command {
	var id influxid.ID
	return &cli.Command{
		Name:  "delete",
		Usage: "Delete user",
		Flags: append(
			commonFlags(),
			&cli.GenericFlag{
				Name:     "id",
				Usage:    "The user ID",
				Aliases:  []string{"i"},
				Required: true,
				Value:    &id,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			client := user.Client{CLI: getCLI(ctx), UsersApi: getAPI(ctx).UsersApi}
			return client.Delete(ctx.Context, id)
		},
	}
}

func newUserListCmd() *cli.Command {
	var params user.ListParams
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"find", "ls"},
		Usage:   "List users",
		Flags: append(
			commonFlags(),
			&cli.GenericFlag{
				Name:    "id",
				Usage:   "The user ID",
				Aliases: []string{"i"},
				Value:   &params.Id,
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "The user name",
				Aliases:     []string{"n"},
				Destination: &params.Name,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			client := user.Client{CLI: getCLI(ctx), UsersApi: getAPI(ctx).UsersApi}
			return client.List(ctx.Context, &params)
		},
	}
}

func newUserUpdateCmd() *cli.Command {
	var params user.UpdateParmas
	return &cli.Command{
		Name: "update",
		Flags: append(
			commonFlags(),
			&cli.GenericFlag{
				Name:     "id",
				Usage:    "The user ID",
				Aliases:  []string{"i"},
				Required: true,
				Value:    &params.Id,
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "The user name",
				Aliases:     []string{"n"},
				Destination: &params.Name,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			client := user.Client{CLI: getCLI(ctx), UsersApi: getAPI(ctx).UsersApi}
			return client.Update(ctx.Context, &params)
		},
	}
}

func newUserSetPasswordCmd() *cli.Command {
	var params user.SetPasswordParams
	return &cli.Command{
		Name: "password",
		Flags: append(
			commonFlagsNoPrint(),
			&cli.GenericFlag{
				Name:    "id",
				Usage:   "The user ID",
				Aliases: []string{"i"},
				Value:   &params.Id,
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "The user name",
				Aliases:     []string{"n"},
				Destination: &params.Name,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			client := user.Client{CLI: getCLI(ctx), UsersApi: getAPI(ctx).UsersApi}
			return client.SetPassword(ctx.Context, &params)
		},
	}
}

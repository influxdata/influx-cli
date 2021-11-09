package main

import (
	"github.com/influxdata/influx-cli/v2/clients/v1_auth"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

// getAuthLookupFlags returns flags used for authorization, requiring either an ID or Username,
// and can specify to require one but not both.
func getAuthLookupFlags(params *v1_auth.AuthLookupParams) []cli.Flag {
	return []cli.Flag{
		&cli.GenericFlag{
			Name:  "id",
			Usage: "The ID of the authorization",
			Value: &params.ID,
		},
		&cli.StringFlag{
			Name:        "username",
			Usage:       "The username of the authorization",
			Destination: &params.Username,
		},
	}
}

func newV1AuthCommand() cli.Command {
	return cli.Command{
		Name:    "auth",
		Usage:   "Authorization management commands for v1 APIs",
		Aliases: []string{"authorization"},
		Before:  middleware.NoArgs,
		Subcommands: []cli.Command{
			newCreateV1AuthCmd(),
			newRemoveV1AuthCmd(),
			newListV1AuthCmd(),
			newSetActiveV1AuthCmd(),
			newSetInactiveV1AuthCmd(),
			newSetPswdV1AuthCmd(),
		},
	}
}

func newCreateV1AuthCmd() cli.Command {
	var params v1_auth.CreateParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	flags = append(flags,
		cli.StringFlag{
			Name:        "description, d",
			Usage:       "Token description",
			Destination: &params.Desc,
		},
		&cli.StringFlag{
			Name:        "username",
			Usage:       "The username to identify this token",
			Required:    true,
			Destination: &params.Username,
		},
		&cli.StringFlag{
			Name:        "password",
			Usage:       "The password to set on this token",
			Destination: &params.Password,
		},
		&cli.BoolFlag{
			Name:        "no-password",
			Usage:       "Don't prompt for a password. You must use v1 auth set-password command before using the token.",
			Destination: &params.NoPassword,
		},
		&cli.StringSliceFlag{
			Name:  "read-bucket",
			Usage: "The bucket id",
		},
		&cli.StringSliceFlag{
			Name:  "write-bucket",
			Usage: "The bucket id",
		},
	)
	return cli.Command{
		Name:   "create",
		Usage:  "Create authorization",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			params.ReadBucket = ctx.StringSlice("read-bucket")
			params.WriteBucket = ctx.StringSlice("write-bucket")
			api := getAPI(ctx)
			client := v1_auth.Client{
				CLI:                     getCLI(ctx),
				LegacyAuthorizationsApi: api.LegacyAuthorizationsApi,
				UsersApi:                api.UsersApi,
				OrganizationsApi:        api.OrganizationsApi,
			}
			return client.Create(getContext(ctx), &params)
		},
	}
}

func newRemoveV1AuthCmd() cli.Command {
	var params v1_auth.RemoveParams
	flags := append(commonFlags(), getAuthLookupFlags(&params.AuthLookupParams)...)
	return cli.Command{
		Name:   "delete",
		Usage:  "Delete authorization",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			if err := params.AuthLookupParams.Validate(); err != nil {
				return err
			}

			api := getAPI(ctx)
			client := v1_auth.Client{
				CLI:                     getCLI(ctx),
				LegacyAuthorizationsApi: api.LegacyAuthorizationsApi,
				UsersApi:                api.UsersApi,
				OrganizationsApi:        api.OrganizationsApi,
			}
			return client.Remove(getContext(ctx), &params)
		},
	}
}

func newListV1AuthCmd() cli.Command {
	var params v1_auth.ListParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	flags = append(flags, getAuthLookupFlags(&params.AuthLookupParams)...)
	flags = append(flags,
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
		Aliases: []string{"ls", "find"},
		Flags:   flags,
		Before:  middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := v1_auth.Client{
				CLI:                     getCLI(ctx),
				LegacyAuthorizationsApi: api.LegacyAuthorizationsApi,
				UsersApi:                api.UsersApi,
				OrganizationsApi:        api.OrganizationsApi,
			}
			return client.List(getContext(ctx), &params)
		},
	}
}

func newSetActiveV1AuthCmd() cli.Command {
	var params v1_auth.ActiveParams
	flags := append(commonFlags(), getAuthLookupFlags(&params.AuthLookupParams)...)
	return cli.Command{
		Name:   "set-active",
		Usage:  "Change the status of an authorization to active",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			if err := params.AuthLookupParams.Validate(); err != nil {
				return err
			}

			api := getAPI(ctx)
			client := v1_auth.Client{
				CLI:                     getCLI(ctx),
				LegacyAuthorizationsApi: api.LegacyAuthorizationsApi,
				UsersApi:                api.UsersApi,
				OrganizationsApi:        api.OrganizationsApi,
			}
			return client.SetActive(getContext(ctx), &params, true)
		},
	}
}

func newSetInactiveV1AuthCmd() cli.Command {
	var params v1_auth.ActiveParams
	flags := append(commonFlags(), getAuthLookupFlags(&params.AuthLookupParams)...)
	return cli.Command{
		Name:   "set-inactive",
		Usage:  "Change the status of an authorization to inactive",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			if err := params.AuthLookupParams.Validate(); err != nil {
				return err
			}

			api := getAPI(ctx)
			client := v1_auth.Client{
				CLI:                     getCLI(ctx),
				LegacyAuthorizationsApi: api.LegacyAuthorizationsApi,
				UsersApi:                api.UsersApi,
				OrganizationsApi:        api.OrganizationsApi,
			}
			return client.SetActive(getContext(ctx), &params, false)
		},
	}
}

func newSetPswdV1AuthCmd() cli.Command {
	var params v1_auth.SetPasswordParams
	flags := append(coreFlags(), commonTokenFlag())
	flags = append(flags, getAuthLookupFlags(&params.AuthLookupParams)...)
	flags = append(flags,
		&cli.StringFlag{
			Name:        "password",
			Usage:       "Password to set on the authorization",
			Destination: &params.Password,
		},
	)
	return cli.Command{
		Name:   "set-password",
		Usage:  "Set a password for an existing authorization",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := v1_auth.Client{
				CLI:                     getCLI(ctx),
				LegacyAuthorizationsApi: api.LegacyAuthorizationsApi,
			}
			return client.SetPassword(getContext(ctx), &params)
		},
	}
}

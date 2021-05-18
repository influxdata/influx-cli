package main

import (
	"github.com/influxdata/influx-cli/v2/internal/cmd/org"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli/v2"
)

func newOrgMembersCmd() *cli.Command {
	return &cli.Command{
		Name:  "members",
		Usage: "Organization membership commands",
		Subcommands: []*cli.Command{
			newOrgMembersAddCmd(),
			newOrgMembersListCmd(),
			newOrgMembersRemoveCmd(),
		},
	}
}

func newOrgMembersAddCmd() *cli.Command {
	var params org.AddMemberParams
	return &cli.Command{
		Name:   "add",
		Usage:  "Add organization member",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlagsNoPrint(),
			&cli.GenericFlag{
				Name:     "member",
				Usage:    "The member ID",
				Aliases:  []string{"m"},
				Required: true,
				Value:    &params.MemberId,
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "The organization name",
				Aliases:     []string{"n"},
				EnvVars:     []string{"INFLUX_ORG"},
				Destination: &params.OrgName,
			},
			&cli.GenericFlag{
				Name:    "id",
				Usage:   "The organization ID",
				Aliases: []string{"i"},
				EnvVars: []string{"INFLUX_ORG_ID"},
				Value:   &params.OrgID,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := org.Client{
				CLI:              getCLI(ctx),
				OrganizationsApi: getAPI(ctx).OrganizationsApi,
			}
			return client.AddMember(ctx.Context, &params)
		},
	}
}

func newOrgMembersListCmd() *cli.Command {
	var params org.ListMemberParams
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"find", "ls"},
		Usage:   "List organization members",
		Before:  middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "name",
				Usage:       "The organization name",
				Aliases:     []string{"n"},
				EnvVars:     []string{"INFLUX_ORG"},
				Destination: &params.OrgName,
			},
			&cli.GenericFlag{
				Name:    "id",
				Usage:   "The organization ID",
				Aliases: []string{"i"},
				EnvVars: []string{"INFLUX_ORG_ID"},
				Value:   &params.OrgID,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := org.Client{
				CLI:              getCLI(ctx),
				OrganizationsApi: getAPI(ctx).OrganizationsApi,
				UsersApi:         getAPI(ctx).UsersApi,
			}
			return client.ListMembers(ctx.Context, &params)
		},
	}
}

func newOrgMembersRemoveCmd() *cli.Command {
	var params org.RemoveMemberParams
	return &cli.Command{
		Name:   "remove",
		Usage:  "Remove organization member",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlagsNoPrint(),
			&cli.GenericFlag{
				Name:     "member",
				Usage:    "The member ID",
				Aliases:  []string{"m"},
				Required: true,
				Value:    &params.MemberId,
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "The organization name",
				Aliases:     []string{"n"},
				EnvVars:     []string{"INFLUX_ORG"},
				Destination: &params.OrgName,
			},
			&cli.GenericFlag{
				Name:    "id",
				Usage:   "The organization ID",
				Aliases: []string{"i"},
				EnvVars: []string{"INFLUX_ORG_ID"},
				Value:   &params.OrgID,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := org.Client{
				CLI:              getCLI(ctx),
				OrganizationsApi: getAPI(ctx).OrganizationsApi,
			}
			return client.RemoveMember(ctx.Context, &params)
		},
	}
}

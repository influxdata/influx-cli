package main

import (
	"github.com/influxdata/influx-cli/v2/clients/org"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newOrgMembersCmd() cli.Command {
	return cli.Command{
		Name:   "members",
		Usage:  "Organization membership commands",
		Before: middleware.NoArgs,
		Subcommands: []cli.Command{
			newOrgMembersAddCmd(),
			newOrgMembersListCmd(),
			newOrgMembersRemoveCmd(),
		},
	}
}

func newOrgMembersAddCmd() cli.Command {
	var params org.AddMemberParams
	return cli.Command{
		Name:   "add",
		Usage:  "Add organization member",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			commonFlagsNoPrint(),
			&cli.GenericFlag{
				Name:     "member, m",
				Usage:    "The member ID",
				Required: true,
				Value:    &params.MemberId,
			},
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "The organization name",
				EnvVar:      "INFLUX_ORG",
				Destination: &params.OrgName,
			},
			&cli.GenericFlag{
				Name:   "id, i",
				Usage:  "The organization ID",
				EnvVar: "INFLUX_ORG_ID",
				Value:  &params.OrgID,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := org.Client{
				CLI:              getCLI(ctx),
				OrganizationsApi: getAPI(ctx).OrganizationsApi,
			}
			return client.AddMember(getContext(ctx), &params)
		},
	}
}

func newOrgMembersListCmd() cli.Command {
	var params org.ListMemberParams
	return cli.Command{
		Name:    "list",
		Aliases: []string{"find", "ls"},
		Usage:   "List organization members",
		Before:  middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "The organization name",
				EnvVar:      "INFLUX_ORG",
				Destination: &params.OrgName,
			},
			&cli.GenericFlag{
				Name:   "id, i",
				Usage:  "The organization ID",
				EnvVar: "INFLUX_ORG_ID",
				Value:  &params.OrgID,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := org.Client{
				CLI:              getCLI(ctx),
				OrganizationsApi: getAPI(ctx).OrganizationsApi,
				UsersApi:         getAPI(ctx).UsersApi,
			}
			return client.ListMembers(getContext(ctx), &params)
		},
	}
}

func newOrgMembersRemoveCmd() cli.Command {
	var params org.RemoveMemberParams
	return cli.Command{
		Name:   "remove",
		Usage:  "Remove organization member",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			commonFlagsNoPrint(),
			&cli.GenericFlag{
				Name:     "member, m",
				Usage:    "The member ID",
				Required: true,
				Value:    &params.MemberId,
			},
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "The organization name",
				EnvVar:      "INFLUX_ORG",
				Destination: &params.OrgName,
			},
			&cli.GenericFlag{
				Name:   "id, i",
				Usage:  "The organization ID",
				EnvVar: "INFLUX_ORG_ID",
				Value:  &params.OrgID,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := org.Client{
				CLI:              getCLI(ctx),
				OrganizationsApi: getAPI(ctx).OrganizationsApi,
			}
			return client.RemoveMember(getContext(ctx), &params)
		},
	}
}

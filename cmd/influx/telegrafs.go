package main

import (
	"strings"

	"github.com/influxdata/influx-cli/v2/clients/telegrafs"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli/v2"
)

func newTelegrafsCommand() *cli.Command {
	var params telegrafs.TelegrafParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	flags = append(flags, &cli.StringFlag{
		Name:        "id",
		Usage:       "Telegraf configuration ID to retrieve",
		Aliases:     []string{"i"},
		Destination: &params.Id,
	})
	return &cli.Command{
		Name:  "telegrafs",
		Usage: "List Telegraf configuration(s). Subcommands manage Telegraf configurations.",
		Subcommands: []*cli.Command{
			newCreateTelegrafCmd(),
			newRemoveTelegrafCmd(),
			newUpdateTelegrafCmd(),
		},
		Flags: flags,
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := telegrafs.Client{
				CLI:          getCLI(ctx),
				TelegrafsApi: api.TelegrafsApi,
			}
			return client.Telegrafs(ctx.Context, &params)
		},
	}
}

func newCreateTelegrafCmd() *cli.Command {
	var params telegrafs.TelegrafParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "description",
			Usage:       "Description for Telegraf configuration",
			Aliases:     []string{"d"},
			Destination: &params.Desc,
		},
		&cli.StringFlag{
			Name:        "file",
			Usage:       "Path to Telegraf configuration",
			Aliases:     []string{"f"},
			Destination: &params.File,
		},
		&cli.StringFlag{
			Name:        "name",
			Usage:       "Name of Telegraf configuration",
			Aliases:     []string{"n"},
			Destination: &params.Name,
		},
	}...)
	return &cli.Command{
		Name:   "create",
		Usage:  "The telegrafs create command creates a new Telegraf configuration.",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := telegrafs.Client{
				CLI:          getCLI(ctx),
				TelegrafsApi: api.TelegrafsApi,
			}
			return client.Create(ctx.Context, &params)
		},
	}
}

func newRemoveTelegrafCmd() *cli.Command {
	var params telegrafs.TelegrafParams
	flags := append(commonFlags())
	flags = append(flags, &cli.StringSliceFlag{
		Name:    "id",
		Usage:   "Telegraf configuration ID(s) to remove",
		Aliases: []string{"i"},
	})
	return &cli.Command{
		Name:    "rm",
		Usage:   "The telegrafs rm command removes Telegraf configuration(s).",
		Aliases: []string{"remove"},
		Flags:   flags,
		Before:  middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {

			// The old CLI allowed specifying this either via repeated flags or
			// via a single flag w/ a comma-separated value.
			rawIds := ctx.StringSlice("id")
			var ids []string
			for _, p := range rawIds {
				ids = append(ids, strings.Split(p, ",")...)
			}
			params.Ids = ids

			api := getAPI(ctx)
			client := telegrafs.Client{
				CLI:          getCLI(ctx),
				TelegrafsApi: api.TelegrafsApi,
			}
			return client.Remove(ctx.Context, &params)
		},
	}
}

func newUpdateTelegrafCmd() *cli.Command {
	var params telegrafs.TelegrafParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "description",
			Usage:       "Description for Telegraf configuration",
			Aliases:     []string{"d"},
			Destination: &params.Desc,
		},
		&cli.StringFlag{
			Name:        "file",
			Usage:       "Path to Telegraf configuration",
			Aliases:     []string{"f"},
			Destination: &params.File,
		},
		&cli.StringFlag{
			Name:        "name",
			Usage:       "Name of Telegraf configuration",
			Aliases:     []string{"n"},
			Destination: &params.Name,
		},
		&cli.StringFlag{
			Name:        "id",
			Usage:       "Telegraf configuration ID to retrieve",
			Aliases:     []string{"i"},
			Destination: &params.Id,
		},
	}...)
	return &cli.Command{
		Name:   "update",
		Usage:  "Update a Telegraf configuration.",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := telegrafs.Client{
				CLI:          getCLI(ctx),
				TelegrafsApi: api.TelegrafsApi,
			}
			return client.Update(ctx.Context, &params)
		},
	}
}

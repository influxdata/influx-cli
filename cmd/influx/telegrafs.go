package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/influxdata/influx-cli/v2/clients/telegrafs"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli/v2"
)

func newTelegrafsCommand() *cli.Command {
	var params telegrafs.ListParams
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
		Description: `List Telegraf configuration(s). Subcommands manage Telegraf configurations.

Examples:
	# list all known Telegraf configurations
	influx telegrafs

	# list Telegraf configuration corresponding to specific ID
	influx telegrafs --id $ID

	# list Telegraf configuration corresponding to specific ID shorts
	influx telegrafs -i $ID
`,
		Subcommands: []*cli.Command{
			newCreateTelegrafCmd(),
			newRemoveTelegrafCmd(),
			newUpdateTelegrafCmd(),
		},
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := telegrafs.Client{
				CLI:              getCLI(ctx),
				TelegrafsApi:     api.TelegrafsApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.List(ctx.Context, &params)
		},
	}
}

func newCreateTelegrafCmd() *cli.Command {
	var params telegrafs.CreateParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	flags = append(flags,
		&cli.StringFlag{
			Name:        "description",
			Usage:       "Description for Telegraf configuration",
			Aliases:     []string{"d"},
			Destination: &params.Desc,
		},
		&cli.StringFlag{
			Name:    "file",
			Usage:   "Path to Telegraf configuration",
			Aliases: []string{"f"},
		},
		&cli.StringFlag{
			Name:        "name",
			Usage:       "Name of Telegraf configuration",
			Aliases:     []string{"n"},
			Destination: &params.Name,
		},
	)
	return &cli.Command{
		Name:  "create",
		Usage: "The telegrafs create command creates a new Telegraf configuration.",
		Description: `The telegrafs create command creates a new Telegraf configuration.

Examples:
	# create new Telegraf configuration
	influx telegrafs create --name $CFG_NAME --description $CFG_DESC --file $PATH_TO_TELE_CFG

	# create new Telegraf configuration using shorts
	influx telegrafs create -n $CFG_NAME -d $CFG_DESC -f $PATH_TO_TELE_CFG

	# create a new Telegraf config with a config provided via STDIN
	cat $CONFIG_FILE | influx telegrafs create -n $CFG_NAME -d $CFG_DESC
`,
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			conf, err := readConfig(ctx.String("file"))
			if err != nil {
				return err
			}
			params.Config = conf
			api := getAPI(ctx)
			client := telegrafs.Client{
				CLI:              getCLI(ctx),
				TelegrafsApi:     api.TelegrafsApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Create(ctx.Context, &params)
		},
	}
}

func newRemoveTelegrafCmd() *cli.Command {
	var params telegrafs.RemoveParams
	flags := commonFlags()
	flags = append(flags,
		&cli.StringSliceFlag{
			Name:    "id",
			Usage:   "Telegraf configuration ID(s) to remove",
			Aliases: []string{"i"},
		},
		&cli.StringFlag{Name: "org", Hidden: true},
		&cli.StringFlag{Name: "org-id", Hidden: true},
	)
	return &cli.Command{
		Name:  "rm",
		Usage: "The telegrafs rm command removes Telegraf configuration(s).",
		Description: `The telegrafs rm command removes Telegraf configuration(s).

Examples:
	# remove a single Telegraf configuration
	influx telegrafs rm --id $ID

	# remove multiple Telegraf configurations
	influx telegrafs rm --id $ID1 --id $ID2

	# remove using short flags
	influx telegrafs rm -i $ID1
`,
		Aliases: []string{"remove"},
		Flags:   flags,
		Before:  middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			rawIds := ctx.StringSlice("id")
			params.Ids = rawIds

			api := getAPI(ctx)
			client := telegrafs.Client{
				CLI:              getCLI(ctx),
				TelegrafsApi:     api.TelegrafsApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Remove(ctx.Context, &params)
		},
	}
}

func newUpdateTelegrafCmd() *cli.Command {
	var params telegrafs.UpdateParams
	flags := commonFlags()
	flags = append(flags,
		&cli.StringFlag{
			Name:        "description",
			Usage:       "Description for Telegraf configuration",
			Aliases:     []string{"d"},
			Destination: &params.Desc,
		},
		&cli.StringFlag{
			Name:    "file",
			Usage:   "Path to Telegraf configuration",
			Aliases: []string{"f"},
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
		&cli.StringFlag{Name: "org", Hidden: true},
		&cli.StringFlag{Name: "org-id", Hidden: true},
	)
	return &cli.Command{
		Name:  "update",
		Usage: "Update a Telegraf configuration.",
		Description: `The telegrafs update command updates a Telegraf configuration to match the
specified parameters. If a name or description is not provided, then are set
to an empty string.

Examples:
	# update new Telegraf configuration
	influx telegrafs update --id $ID --name $CFG_NAME --description $CFG_DESC --file $PATH_TO_TELE_CFG

	# update new Telegraf configuration using shorts
	influx telegrafs update -i $ID -n $CFG_NAME -d $CFG_DESC -f $PATH_TO_TELE_CFG

	# update a Telegraf config with a config provided via STDIN
	cat $CONFIG_FILE | influx telegrafs update -i $ID  -n $CFG_NAME -d $CFG_DESC
`,
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			conf, err := readConfig(ctx.String("file"))
			if err != nil {
				return err
			}
			params.Config = conf
			api := getAPI(ctx)
			client := telegrafs.Client{
				CLI:              getCLI(ctx),
				TelegrafsApi:     api.TelegrafsApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Update(ctx.Context, &params)
		},
	}
}

func readConfig(file string) (string, error) {
	if file != "" {
		bb, err := ioutil.ReadFile(file)
		if err != nil {
			return "", fmt.Errorf("failed to read telegraf config from %q: %w", file, err)
		}

		return string(bb), nil
	}

	bb, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return "", fmt.Errorf("failed to read telegraf config from stdin: %w", err)
	}
	return string(bb), nil
}

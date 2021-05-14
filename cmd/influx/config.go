package main

import (
	"fmt"
	"os"
	"path"

	"github.com/influxdata/influx-cli/v2/internal/cmd/config"
	iconfig "github.com/influxdata/influx-cli/v2/internal/config"
	"github.com/urfave/cli/v2"
)

var configPathAndPrintFlags = append([]cli.Flag{configPathFlag()}, printFlags()...)

func newConfigCmd() *cli.Command {
	return &cli.Command{
		Name:      "config",
		Usage:     "Config management commands",
		ArgsUsage: "[config name]",
		Description: `
Providing no argument to the config command will print the active configuration.
When an argument is provided, the active config will be switched to the config with
a name matching that of the argument provided.

Examples:
	# show active config
	influx config

	# set active config to previously active config
	influx config -

	# set active config
	influx config $CONFIG_NAME

The influx config command displays the active InfluxDB connection configuration and
manages multiple connection configurations stored, by default, in ~/.influxdbv2/configs.
Each connection includes a URL, token, associated organization, and active setting.
InfluxDB reads the token from the active connection configuration, so you don't have
to manually enter a token to log into InfluxDB.

For information about the config command, see
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/config/
`,
		Before: withCli(),
		Flags:  configPathAndPrintFlags,
		Action: func(ctx *cli.Context) error {
			prog := path.Base(os.Args[0])
			if ctx.NArg() > 1 {
				return fmt.Errorf("usage: %s config [config name]", prog)
			}
			client := config.Client{CLI: getCLI(ctx)}
			if ctx.NArg() == 1 {
				return client.SwitchActive(ctx.Args().Get(0))
			}
			return client.PrintActive()
		},
		Subcommands: []*cli.Command{
			newConfigCreateCmd(),
			newConfigDeleteCmd(),
			newConfigUpdateCmd(),
			newConfigListCmd(),
		},
	}
}

func newConfigCreateCmd() *cli.Command {
	var cfg iconfig.Config
	return &cli.Command{
		Name:  "create",
		Usage: "Create config",
		Description: `
The influx config create command creates a new InfluxDB connection configuration
and stores it in the configs file (by default, stored at ~/.influxdbv2/configs).

Examples:
	# create a config and set it active
	influx config create -a -n $CFG_NAME -u $HOST_URL -t $TOKEN -o $ORG_NAME

	# create a config and without setting it active
	influx config create -n $CFG_NAME -u $HOST_URL -t $TOKEN -o $ORG_NAME

For information about the config command, see
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/config/
and
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/config/create/
`,
		Before: withCli(),
		Flags: append(
			configPathAndPrintFlags,
			&cli.StringFlag{
				Name:        "config-name",
				Usage:       "Name for the new config",
				Aliases:     []string{"n"},
				Required:    true,
				Destination: &cfg.Name,
			},
			&cli.StringFlag{
				Name:        "host-url",
				Usage:       "Base URL of the InfluxDB server the new config should target",
				Aliases:     []string{"u"},
				Required:    true,
				Destination: &cfg.Host,
			},
			&cli.StringFlag{
				Name:        "token",
				Usage:       "Auth token to use when communicating with the InfluxDB server",
				Aliases:     []string{"t"},
				Required:    true,
				Destination: &cfg.Token,
			},
			&cli.StringFlag{
				Name:        "org",
				Usage:       "Default organization name to use in the new config",
				Aliases:     []string{"o"},
				Destination: &cfg.Org,
			},
			&cli.BoolFlag{
				Name:        "active",
				Usage:       "Set the new config as active",
				Aliases:     []string{"a"},
				Destination: &cfg.Active,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := config.Client{CLI: getCLI(ctx)}
			return client.Create(cfg)
		},
	}
}

func newConfigDeleteCmd() *cli.Command {
	return &cli.Command{
		Name:      "rm",
		Aliases:   []string{"delete", "remove"},
		Usage:     "Delete config",
		ArgsUsage: "[cfg_name]...",
		Description: `
The influx config delete command deletes an InfluxDB connection configuration from
the configs file (by default, stored at ~/.influxdbv2/configs).

Examples:
	# delete a config
	influx config rm $CFG_NAME

	# delete multiple configs
	influx config rm $CFG_NAME_1 $CFG_NAME_2

For information about the config command, see
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/config/
and
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/config/rm/
`,
		Before: withCli(),
		Flags:  append([]cli.Flag{configPathFlag()}, printFlags()...),
		Action: func(ctx *cli.Context) error {
			client := config.Client{CLI: getCLI(ctx)}
			return client.Delete(ctx.Args().Slice())
		},
	}
}

func newConfigUpdateCmd() *cli.Command {
	var cfg iconfig.Config
	return &cli.Command{
		Name:    "set",
		Aliases: []string{"update"},
		Usage:   "Update config",
		Description: `
The influx config set command updates information in an InfluxDB connection
configuration in the configs file (by default, stored at ~/.influxdbv2/configs).

Examples:
	# update a config and set active
	influx config set -a -n $CFG_NAME -u $HOST_URL -t $TOKEN -o $ORG_NAME

	# update a config and do not set to active
	influx config set -n $CFG_NAME -u $HOST_URL -t $TOKEN -o $ORG_NAME

For information about the config command, see
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/config/
and
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/config/set/
`,
		Before: withCli(),
		Flags: append(
			configPathAndPrintFlags,
			&cli.StringFlag{
				Name:        "config-name",
				Usage:       "Name of the config to update",
				Aliases:     []string{"n"},
				Required:    true,
				Destination: &cfg.Name,
			},
			&cli.StringFlag{
				Name:        "host-url",
				Usage:       "New URL to set on the config",
				Aliases:     []string{"u"},
				Destination: &cfg.Host,
			},
			&cli.StringFlag{
				Name:        "token",
				Usage:       "New auth token to set on the config",
				Aliases:     []string{"t"},
				Destination: &cfg.Token,
			},
			&cli.StringFlag{
				Name:        "org",
				Usage:       "New default organization to set on the config",
				Aliases:     []string{"o"},
				Destination: &cfg.Org,
			},
			&cli.BoolFlag{
				Name:        "active",
				Usage:       "Set the config as active",
				Aliases:     []string{"a"},
				Destination: &cfg.Active,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := config.Client{CLI: getCLI(ctx)}
			return client.Update(cfg)
		},
	}
}

func newConfigListCmd() *cli.Command {
	return &cli.Command{
		Name:    "ls",
		Aliases: []string{"list"},
		Usage:   "List configs",
		Description: `
The influx config ls command lists all InfluxDB connection configurations
in the configs file (by default, stored at ~/.influxdbv2/configs). Each
connection configuration includes a URL, authentication token, and active
setting. An asterisk (*) indicates the active configuration.

Examples:
	# list configs
	influx config ls

	# list configs with long alias
	influx config list

For information about the config command, see
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/config/
and
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/config/list/
`,
		Before: withCli(),
		Flags:  configPathAndPrintFlags,
		Action: func(ctx *cli.Context) error {
			client := config.Client{CLI: getCLI(ctx)}
			return client.List()
		},
	}
}

package main

import (
	"fmt"
	"os"
	"path"

	cmd "github.com/influxdata/influx-cli/v2/clients/config"
	"github.com/influxdata/influx-cli/v2/config"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

var configPathAndPrintFlags = append([]cli.Flag{configPathFlag()}, printFlags()...)

func newConfigCmd() cli.Command {
	return cli.Command{
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
			client := cmd.Client{CLI: getCLI(ctx)}
			if ctx.NArg() == 1 {
				return client.SwitchActive(ctx.Args().Get(0))
			}
			return client.PrintActive()
		},
		Subcommands: []cli.Command{
			newConfigCreateCmd(),
			newConfigDeleteCmd(),
			newConfigUpdateCmd(),
			newConfigListCmd(),
		},
	}
}

func newConfigCreateCmd() cli.Command {
	var cfg config.Config
	var userpass string
	return cli.Command{
		Name:  "create",
		Usage: "Create config",
		Description: `
The influx config create command creates a new InfluxDB connection configuration
and stores it in the configs file (by default, stored at ~/.influxdbv2/configs).

Authentication:
	Authentication can be provided by either an api token or username/password, but not both.
	When setting the username and password, the password is saved unencrypted in your local config file.
	Optionally, you can omit the password and only provide the username.
	You will then be prompted for the password each time.

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
		Before: middleware.WithBeforeFns(withCli(), middleware.NoArgs),
		Flags: append(
			configPathAndPrintFlags,
			&cli.StringFlag{
				Name:        "config-name, n",
				Usage:       "Name for the new config",
				Required:    true,
				Destination: &cfg.Name,
			},
			&cli.StringFlag{
				Name:        "host-url, u",
				Usage:       "Base URL of the InfluxDB server the new config should target",
				Required:    true,
				Destination: &cfg.Host,
			},
			&cli.StringFlag{
				Name:        "token, t",
				Usage:       "Auth token to use when communicating with the InfluxDB server",
				Destination: &cfg.Token,
			},
			&cli.StringFlag{
				Name:        "username-password, p",
				Usage:       "Username (and optionally password) to use for authentication. Only supported in OSS",
				Destination: &userpass,
			},
			&cli.StringFlag{
				Name:        "org, o",
				Usage:       "Default organization name to use in the new config",
				Destination: &cfg.Org,
			},
			&cli.BoolFlag{
				Name:        "active, a",
				Usage:       "Set the new config as active",
				Destination: &cfg.Active,
			},
		),
		Action: func(ctx *cli.Context) error {
			if cfg.Token != "" && userpass != "" {
				return fmt.Errorf("cannot specify `--token` and `--username-password` together, please choose one")
			}
			client := cmd.Client{CLI: getCLI(ctx)}
			if userpass != "" {
				return client.CreateWithUserPass(cfg, userpass)
			}
			return client.Create(cfg)
		},
	}
}

func newConfigDeleteCmd() cli.Command {
	return cli.Command{
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
			client := cmd.Client{CLI: getCLI(ctx)}
			return client.Delete(ctx.Args())
		},
	}
}

func newConfigUpdateCmd() cli.Command {
	var cfg config.Config
	var userpass string
	return cli.Command{
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
		Before: middleware.WithBeforeFns(withCli(), middleware.NoArgs),
		Flags: append(
			configPathAndPrintFlags,
			&cli.StringFlag{
				Name:        "config-name, n",
				Usage:       "Name of the config to update",
				Required:    true,
				Destination: &cfg.Name,
			},
			&cli.StringFlag{
				Name:        "host-url, u",
				Usage:       "New URL to set on the config",
				Destination: &cfg.Host,
			},
			&cli.StringFlag{
				Name:        "token, t",
				Usage:       "New auth token to set on the config",
				Destination: &cfg.Token,
			},
			&cli.StringFlag{
				Name:        "username-password, p",
				Usage:       "Username (and optionally password) to use for authentication. Only supported in OSS",
				Destination: &userpass,
			},
			&cli.StringFlag{
				Name:        "org, o",
				Usage:       "New default organization to set on the config",
				Destination: &cfg.Org,
			},
			&cli.BoolFlag{
				Name:        "active, a",
				Usage:       "Set the config as active",
				Destination: &cfg.Active,
			},
		),
		Action: func(ctx *cli.Context) error {
			if cfg.Token != "" && userpass != "" {
				return fmt.Errorf("cannot specify `--token` and `--username-password` together, please choose one")
			}
			client := cmd.Client{CLI: getCLI(ctx)}
			if userpass != "" {
				return client.UpdateWithUserPass(cfg, userpass)
			}
			return client.Update(cfg)
		},
	}
}

func newConfigListCmd() cli.Command {
	return cli.Command{
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
		Before: middleware.WithBeforeFns(withCli(), middleware.NoArgs),
		Flags:  configPathAndPrintFlags,
		Action: func(ctx *cli.Context) error {
			client := cmd.Client{CLI: getCLI(ctx)}
			return client.List()
		},
	}
}

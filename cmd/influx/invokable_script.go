package main

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/influxdata/influx-cli/v2/clients/invokable_script"
	"github.com/influxdata/influx-cli/v2/clients/query"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newInvokableScriptCmd() cli.Command {
	return cli.Command{
		Name:  "script",
		Usage: "Invokable script management commands",
		Subcommands: []cli.Command{
			newInvokableScriptListCommand(),
			newInvokableScriptCreateCommand(),
			newInvokableScriptGetCommand(),
			newInvokableScriptDeleteCommand(),
			newInvokableScriptUpdateCommand(),
			newInvokableScriptInvokeCommand(),
		},
	}
}

func newInvokableScriptListCommand() cli.Command {
	var params invokable_script.ListParams
	flags := commonFlags()

	return cli.Command{
		Name:    "list",
		Usage:   "List invokable scripts",
		Aliases: []string{"find", "ls"},
		Before:  middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			flags,
			&cli.StringFlag{
				Name:        "limit",
				Usage:       "Limit the number of results returned",
				Destination: &params.Limit,
			},
			&cli.StringFlag{
				Name:        "offset",
				Usage:       "Number of scripts to skip over in the list",
				Destination: &params.Offset,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := invokable_script.Client{
				CLI:                 getCLI(ctx),
				InvocableScriptsApi: api.InvocableScriptsApi.OnlyCloud(),
			}
			return client.List(getContext(ctx), &params)
		},
	}
}

func newInvokableScriptCreateCommand() cli.Command {
	var params invokable_script.CreateParams
	var scriptFilePath string

	return cli.Command{
		Name:   "create",
		Usage:  "Create a new invokable script",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "Name of the new invokable script",
				Destination: &params.Name,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "description, d",
				Usage:       "Description of the new invokable script",
				Destination: &params.Description,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "language, l",
				Usage:       "Language of the invokable script",
				Destination: &params.Language,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "script, s",
				Usage:       "Updated script to be executed",
				Destination: &params.Script,
			},
			&cli.StringFlag{
				Name:        "file, f",
				Usage:       "Path to the script file",
				TakesFile:   true,
				Destination: &scriptFilePath,
			},
		),
		Action: func(ctx *cli.Context) error {
			if scriptFilePath == "" && params.Script == "" {
				return errors.New("must provide a script using --file or --script")
			} else if scriptFilePath != "" && params.Script != "" {
				return errors.New("cannot use --file and --script simultaneously")
			}

			if scriptFilePath != "" {
				bytes, err := os.ReadFile(scriptFilePath)
				if err != nil {
					return err
				}
				params.Script = string(bytes)
			}

			api := getAPI(ctx)
			client := invokable_script.Client{
				CLI:                 getCLI(ctx),
				InvocableScriptsApi: api.InvocableScriptsApi.OnlyCloud(),
			}
			return client.Create(getContext(ctx), &params)
		},
	}
}

func newInvokableScriptGetCommand() cli.Command {
	var params invokable_script.GetParams

	return cli.Command{
		Name:   "get",
		Usage:  "Get the information about a single invokable script",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "id",
				Usage:       "ID of the invokable script to get",
				Destination: &params.ID,
				Required:    true,
			},
			&cli.BoolFlag{
				Name:        "print-script-only",
				Usage:       "Print only the formatted script value itself for the selected script",
				Destination: &params.PrintScriptOnly,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := invokable_script.Client{
				CLI:                 getCLI(ctx),
				InvocableScriptsApi: api.InvocableScriptsApi.OnlyCloud(),
			}
			return client.Get(getContext(ctx), &params)
		},
	}
}

func newInvokableScriptDeleteCommand() cli.Command {
	var params invokable_script.DeleteParams

	return cli.Command{
		Name:   "delete",
		Usage:  "Delete an invokable script by ID",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "id",
				Usage:       "ID of the invokable script to delete",
				Destination: &params.ID,
				Required:    true,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := invokable_script.Client{
				CLI:                 getCLI(ctx),
				InvocableScriptsApi: api.InvocableScriptsApi.OnlyCloud(),
			}
			return client.Delete(getContext(ctx), &params)
		},
	}
}

func newInvokableScriptUpdateCommand() cli.Command {
	var params invokable_script.UpdateParams
	var scriptFilePath string

	return cli.Command{
		Name:   "update",
		Usage:  "Update an invokable script by ID",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "id",
				Usage:       "ID of the invokable script to update",
				Destination: &params.ID,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "Updated name of the invokable script",
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "description, d",
				Usage:       "Updated description of the invokable script",
				Destination: &params.Description,
			},
			&cli.StringFlag{
				Name:        "script, s",
				Usage:       "Updated script to be executed",
				Destination: &params.Script,
			},
			&cli.StringFlag{
				Name:        "file, f",
				Usage:       "The path to the updated script file",
				TakesFile:   true,
				Destination: &scriptFilePath,
			},
		),
		Action: func(ctx *cli.Context) error {
			if scriptFilePath != "" && params.Script != "" {
				return errors.New("cannot use --file and --script simultaneously")
			}

			if scriptFilePath != "" {
				bytes, err := os.ReadFile(scriptFilePath)
				if err != nil {
					return err
				}
				params.Script = string(bytes)
			}

			api := getAPI(ctx)
			client := invokable_script.Client{
				CLI:                 getCLI(ctx),
				InvocableScriptsApi: api.InvocableScriptsApi.OnlyCloud(),
			}
			return client.Update(getContext(ctx), &params)
		},
	}
}

func newInvokableScriptInvokeCommand() cli.Command {
	var params invokable_script.InvokeParams
	var paramsFilePath, paramsString string

	return cli.Command{
		Name:   "invoke",
		Usage:  "Invoke a script by ID, with parameters provided as JSON.",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "id",
				Usage:       "ID of the script to invoke",
				Destination: &params.ID,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "parameters, p",
				Usage:       "JSON string of parameters",
				Destination: &paramsString,
			},
			&cli.StringFlag{
				Name:        "file, f",
				Usage:       "Path to the parameters file",
				TakesFile:   true,
				Destination: &paramsFilePath,
			},
		),
		Action: func(ctx *cli.Context) error {
			if paramsFilePath != "" && paramsString != "" {
				return errors.New("cannot use --file and --parameters simultaneously")
			}

			if paramsFilePath != "" {
				var err error
				params.ScriptParams, err = os.Open(paramsFilePath)
				if err != nil {
					return err
				}
			}

			if paramsString != "" {
				params.ScriptParams = io.NopCloser(strings.NewReader(paramsString))
			}

			api := getAPI(ctx)
			client := invokable_script.Client{
				CLI:                 getCLI(ctx),
				InvocableScriptsApi: api.InvocableScriptsApi.OnlyCloud(),
				// Line protocol results are not currently returned in annotated format, so the formatting printer is not an
				// option. Only the raw result printer will work.
				ResultPrinter: query.RawResultPrinter,
			}
			return client.Invoke(getContext(ctx), &params)
		},
	}
}

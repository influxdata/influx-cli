package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/influxdata/influx-cli/v2/clients/script"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newScriptsCmd() cli.Command {
	return cli.Command{
		Name:   "scripts",
		Usage:  "Scripts management commands",
		Before: middleware.NoArgs,
		Subcommands: []cli.Command{
			newScriptsListCmd(),
			newScriptsCreateCmd(),
			newScriptsDeleteCmd(),
			newScriptsRetrieveCmd(),
			newScriptsUpdateCmd(),
			newScriptsInvokeCmd(),
		},
	}
}

func newScriptsListCmd() cli.Command {
	var params script.ListParams
	flags := []cli.Flag{
		&cli.IntFlag{
			Name:        "limit, l",
			Usage:       "The number of scripts to return",
			Destination: &params.Limit,
		},
		&cli.IntFlag{
			Name:        "offset, o",
			Usage:       "The offset for pagination",
			Destination: &params.Offset,
		}}
	flags = append(flags, commonFlags()...)

	return cli.Command{
		Name:   "list",
		Usage:  "Lists scripts",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := script.Client{
				CLI:                 getCLI(ctx),
				InvokableScriptsApi: api.InvokableScriptsApi,
			}

			return client.List(getContext(ctx), &params)
		},
	}
}

func newScriptsCreateCmd() cli.Command {
	var params script.CreateParams
	var scriptFile string
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "description, d",
			Usage:       "The purpose or functionality of the script",
			Destination: &params.Description,
		},
		&cli.StringFlag{
			Name:        "language, l",
			Usage:       "What language the script is written in",
			Destination: &params.Language,
		},
		&cli.StringFlag{
			Name:        "name, n",
			Usage:       "Name of the script",
			Destination: &params.Name,
		},
		&cli.StringFlag{
			Name:        "script, s",
			Usage:       "Contents of the script to be executed",
			Destination: &params.Script,
		},
		&cli.StringFlag{
			Name:        "file, f",
			Usage:       "File name containing the contents of the script to be executed",
			Destination: &scriptFile,
		},
	}
	flags = append(flags, commonFlags()...)

	return cli.Command{
		Name:   "create",
		Usage:  "Creates a new script",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			if len(params.Script) > 0 && len(scriptFile) > 0 {
				return errors.New("cannot specify both a script string and a file")
			}

			if len(scriptFile) > 0 {
				data, err := os.ReadFile(scriptFile)
				if err != nil {
					return fmt.Errorf("failed to create script: %v", err)
				}
				params.Script = string(data)
			}

			api := getAPI(ctx)
			client := script.Client{
				CLI:                 getCLI(ctx),
				InvokableScriptsApi: api.InvokableScriptsApi,
			}

			return client.Create(getContext(ctx), &params)
		},
	}
}

func newScriptsDeleteCmd() cli.Command {
	var params script.DeleteParams
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "scriptID, i",
			Usage:       "Script identifier",
			Destination: &params.ScriptID,
		}}
	flags = append(flags, commonFlags()...)

	return cli.Command{
		Name:   "delete",
		Usage:  "Deletes a script",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := script.Client{
				CLI:                 getCLI(ctx),
				InvokableScriptsApi: api.InvokableScriptsApi,
			}

			return client.Delete(getContext(ctx), &params)
		},
	}
}

func newScriptsRetrieveCmd() cli.Command {
	var params script.RetrieveParams
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "scriptID, i",
			Usage:       "Script identifier",
			Destination: &params.ScriptID,
		}}
	flags = append(flags, commonFlags()...)

	return cli.Command{
		Name:   "retrieve",
		Usage:  "Retrieves a script",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := script.Client{
				CLI:                 getCLI(ctx),
				InvokableScriptsApi: api.InvokableScriptsApi,
			}

			return client.Retrieve(getContext(ctx), &params)
		},
	}
}

func newScriptsUpdateCmd() cli.Command {
	var params script.UpdateParams
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "scriptID, i",
			Usage:       "Script identifier",
			Destination: &params.ScriptID,
		},
		&cli.StringFlag{
			Name:        "description, d",
			Usage:       "New script description",
			Destination: &params.Description,
		},
		&cli.StringFlag{
			Name:        "name, n",
			Usage:       "New script name",
			Destination: &params.Name,
		},
		&cli.StringFlag{
			Name:        "script, s",
			Usage:       "New script contents",
			Destination: &params.Script,
		}}
	flags = append(flags, commonFlags()...)

	return cli.Command{
		Name:   "update",
		Usage:  "Updates a script",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := script.Client{
				CLI:                 getCLI(ctx),
				InvokableScriptsApi: api.InvokableScriptsApi,
			}

			return client.Update(getContext(ctx), &params)
		},
	}
}

func newScriptsInvokeCmd() cli.Command {
	var params script.InvokeParams
	var jsonParams string
	var jsonFile string
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "scriptID, i",
			Usage:       "Script identifier",
			Destination: &params.ScriptID,
		},
		&cli.StringFlag{
			Name:        "params, p",
			Usage:       "JSON string containing the parameters",
			Destination: &jsonParams,
		},
		&cli.StringFlag{
			Name:        "file, f",
			Usage:       "File name containing the script parameters, in JSON",
			Destination: &jsonFile,
		},
	}
	flags = append(flags, commonFlags()...)

	return cli.Command{
		Name:   "invoke",
		Usage:  "Invokes a script",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			if len(jsonParams) > 0 && len(jsonFile) > 0 {
				return errors.New("cannot specify both a parameter string and a file")
			}

			if len(jsonFile) > 0 {
				data, err := os.ReadFile(jsonFile)
				if err != nil {
					return fmt.Errorf("failed to invoke script: %v", err)
				}
				jsonParams = string(data)
			}

			api := getAPI(ctx)
			client := script.Client{
				CLI:                 getCLI(ctx),
				InvokableScriptsApi: api.InvokableScriptsApi,
			}

			params.Params = make(map[string]interface{})
			if len(jsonParams) > 0 {
				if err := json.NewDecoder(strings.NewReader(jsonParams)).Decode(&params.Params); err != nil {
					return err
				}
			}

			return client.Invoke(getContext(ctx), &params)
		},
	}
}

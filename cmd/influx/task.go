package main

import (
	"github.com/influxdata/influx-cli/v2/internal/cmd/task"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli/v2"
)

func newTaskCommand() *cli.Command {
	return &cli.Command{
		Name:  "task",
		Usage: "Task management commands",
		Subcommands: []*cli.Command{
			//newTaskLogCmd(),
			//newTaskRunCmd(),
			newTaskCreateCmd(),
			//newTaskDeleteCmd(),
			//newTaskFindCmd(),
			//newTaskUpdateCmd(),
			//newTaskRetryFailedCmd(),
		},
	}
}

func newTaskCreateCmd() *cli.Command {
	var params task.CreateParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	flags = append(flags, &cli.StringFlag{
		Name:      "file",
		Usage:     "Path to Flux script file",
		Aliases:   []string{"f"},
		TakesFile: true,
	})
	return &cli.Command{
		Name:      "create",
		Usage:     "Create a task with a Flux script provided via the first argument or a file or stdin",
		ArgsUsage: "[flux script or '-' for stdin]",
		Flags:     flags,
		Before:    middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := task.Client{
				CLI:      getCLI(ctx),
				TasksApi: api.TasksApi,
			}
			var err error
			params.FluxQuery, err = readQuery(ctx)
			if err != nil {
				return err
			}
			return client.Create(ctx.Context, &params)
		},
	}
}

package main

import (
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/task"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

const TaskMaxPageSize = 500

func newTaskCommand() cli.Command {
	return cli.Command{
		Name:   "task",
		Usage:  "Task management commands",
		Before: middleware.NoArgs,
		Subcommands: []cli.Command{
			newTaskLogCmd(),
			newTaskRunCmd(),
			newTaskCreateCmd(),
			newTaskDeleteCmd(),
			newTaskFindCmd(),
			newTaskUpdateCmd(),
			newTaskRetryFailedCmd(),
		},
	}
}

func newTaskCreateCmd() cli.Command {
	var params task.CreateParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	flags = append(flags, &cli.StringFlag{
		Name:      "file, f",
		Usage:     "Path to Flux script file",
		TakesFile: true,
	})
	return cli.Command{
		Name:      "create",
		Usage:     "Create a task with a Flux script provided via the first argument or a file or stdin.",
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
			params.FluxQuery, err = clients.ReadQuery(ctx.String("file"), ctx.Args())
			if err != nil {
				return err
			}
			return client.Create(getContext(ctx), &params)
		},
	}
}

func newTaskFindCmd() cli.Command {
	var params task.FindParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "id, i",
			Usage:       "task ID",
			Destination: &params.TaskID,
		},
		&cli.StringFlag{
			Name:        "user-id, n",
			Usage:       "task owner ID",
			Destination: &params.UserID,
		},
		&cli.IntFlag{
			Name:        "limit",
			Usage:       "the number of tasks to find",
			Destination: &params.Limit,
			Value:       TaskMaxPageSize,
		},
	}...)
	return cli.Command{
		Name:    "list",
		Usage:   "List tasks",
		Aliases: []string{"find", "ls"},
		Flags:   flags,
		Before:  middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := task.Client{
				CLI:      getCLI(ctx),
				TasksApi: api.TasksApi,
			}
			return client.Find(getContext(ctx), &params)
		},
	}
}

func newTaskRetryFailedCmd() cli.Command {
	var params task.RetryFailedParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "id, i",
			Usage:       "task ID",
			Destination: &params.TaskID,
		},
		&cli.StringFlag{
			Name:        "before",
			Usage:       "runs before this time",
			Destination: &params.RunFilter.Before,
		},
		&cli.StringFlag{
			Name:        "after",
			Usage:       "runs after this time",
			Destination: &params.RunFilter.After,
		},
		&cli.BoolFlag{
			Name:        "dry-run",
			Usage:       "print info about runs that would be retried",
			Destination: &params.DryRun,
		},
		&cli.IntFlag{
			Name:        "task-limit",
			Usage:       "max number of tasks to retry failed runs for",
			Destination: &params.TaskLimit,
			Value:       100,
		},
		&cli.IntFlag{
			Name:        "run-limit",
			Usage:       "max number of failed runs to retry per task",
			Destination: &params.RunFilter.Limit,
			Value:       100,
		},
	}...)
	return cli.Command{
		Name:    "retry-failed",
		Usage:   "Retry failed runs",
		Aliases: []string{"rtf"},
		Flags:   flags,
		Before:  middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := task.Client{
				CLI:      getCLI(ctx),
				TasksApi: api.TasksApi,
			}
			return client.RetryFailed(getContext(ctx), &params)
		},
	}
}

func newTaskUpdateCmd() cli.Command {
	var params task.UpdateParams
	flags := commonFlags()
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "id, i",
			Usage:       "task ID (required)",
			Destination: &params.TaskID,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "status",
			Usage:       "update tasks status",
			Destination: &params.Status,
		},
		&cli.StringFlag{
			Name:      "file, f",
			Usage:     "Path to Flux script file",
			TakesFile: true,
		},
	}...)
	return cli.Command{
		Name:      "update",
		Usage:     "Update task status or script. Provide a Flux script via the first argument or a file.",
		ArgsUsage: "[flux script or '-' for stdin]",
		Flags:     flags,
		Before:    middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := task.Client{
				CLI:      getCLI(ctx),
				TasksApi: api.TasksApi,
			}
			var err error
			if ctx.String("file") != "" || ctx.NArg() != 0 {
				params.FluxQuery, err = clients.ReadQuery(ctx.String("file"), ctx.Args())
				if err != nil {
					return err
				}
			}
			return client.Update(getContext(ctx), &params)
		},
	}
}

func newTaskDeleteCmd() cli.Command {
	var params task.DeleteParams
	flags := commonFlags()
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "id, i",
			Usage:       "task ID (required)",
			Destination: &params.TaskID,
			Required:    true,
		},
	}...)
	return cli.Command{
		Name:   "delete",
		Usage:  "Delete tasks",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := task.Client{
				CLI:      getCLI(ctx),
				TasksApi: api.TasksApi,
			}
			return client.Delete(getContext(ctx), &params)
		},
	}
}

func newTaskLogCmd() cli.Command {
	return cli.Command{
		Name:  "log",
		Usage: "Log related commands",
		Subcommands: []cli.Command{
			newTaskLogFindCmd(),
		},
	}
}

func newTaskLogFindCmd() cli.Command {
	var params task.LogFindParams
	flags := commonFlags()
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "task-id",
			Usage:       "task id (required)",
			Destination: &params.TaskID,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "run-id",
			Usage:       "run id",
			Destination: &params.RunID,
		},
	}...)
	return cli.Command{
		Name:    "list",
		Usage:   "List logs for a task",
		Aliases: []string{"find", "ls"},
		Flags:   flags,
		Before:  middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := task.Client{
				CLI:      getCLI(ctx),
				TasksApi: api.TasksApi,
			}
			return client.FindLogs(getContext(ctx), &params)
		},
	}
}

func newTaskRunCmd() cli.Command {
	return cli.Command{
		Name:   "run",
		Usage:  "Run related commands",
		Before: middleware.NoArgs,
		Subcommands: []cli.Command{
			newTaskRunFindCmd(),
			newTaskRunRetryCmd(),
		},
	}
}

func newTaskRunFindCmd() cli.Command {
	var params task.RunFindParams
	flags := commonFlags()
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "task-id",
			Usage:       "task ID (required)",
			Destination: &params.TaskID,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "run-id",
			Usage:       "run id",
			Destination: &params.RunID,
		},
		&cli.StringFlag{
			Name:        "before",
			Usage:       "runs before this time",
			Destination: &params.Filter.Before,
		},
		&cli.StringFlag{
			Name:        "after",
			Usage:       "runs after this time",
			Destination: &params.Filter.After,
		},
		&cli.IntFlag{
			Name:        "limit",
			Usage:       "limit the results",
			Destination: &params.Filter.Limit,
			Value:       100,
		},
	}...)
	return cli.Command{
		Name:    "list",
		Usage:   "List runs for a tasks",
		Aliases: []string{"find", "ls"},
		Flags:   flags,
		Before:  middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := task.Client{
				CLI:      getCLI(ctx),
				TasksApi: api.TasksApi,
			}
			return client.FindRuns(getContext(ctx), &params)
		},
	}
}

func newTaskRunRetryCmd() cli.Command {
	var params task.RunRetryParams
	flags := commonFlags()
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name:        "task-id, i",
			Usage:       "task ID (required)",
			Destination: &params.TaskID,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "run-id, r",
			Usage:       "run ID (required)",
			Destination: &params.RunID,
			Required:    true,
		},
	}...)
	return cli.Command{
		Name:   "retry",
		Usage:  "Retry a run",
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := task.Client{
				CLI:      getCLI(ctx),
				TasksApi: api.TasksApi,
			}
			return client.RetryRun(getContext(ctx), &params)
		},
	}
}

package main

import (
	"fmt"

	"github.com/influxdata/influx-cli/v2/clients/replication"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newReplicationCmd() cli.Command {
	return cli.Command{
		Name:   "replication",
		Usage:  "Replication stream management commands",
		Hidden: true, // Remove this line when all subcommands are completed
		Subcommands: []cli.Command{
			newReplicationCreateCmd(),
			newReplicationDeleteCmd(),
			newReplicationListCmd(),
			newReplicationUpdateCmd(),
		},
	}
}

func newReplicationCreateCmd() cli.Command {
	var params replication.CreateParams
	return cli.Command{
		Name:   "create",
		Usage:  "Create a new replication stream",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "Name for new replication stream",
				Required:    true,
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "description, d",
				Usage:       "Description for new replication stream",
				Destination: &params.Description,
			},
			&cli.StringFlag{
				Name:        "org-id",
				Usage:       "The ID of the local organization",
				EnvVar:      "INFLUX_ORG_ID",
				Destination: &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "org, o",
				Usage:       "The name of the local organization",
				EnvVar:      "INFLUX_ORG",
				Destination: &params.OrgName,
			},
			&cli.StringFlag{
				Name:        "remote-id",
				Usage:       "Remote connection ID new replication stream will be registered with",
				Required:    true,
				Destination: &params.RemoteID,
			},
			&cli.StringFlag{
				Name:        "local-bucket",
				Usage:       "Local bucket ID for new replication stream",
				Required:    true,
				Destination: &params.LocalBucketID,
			},
			&cli.StringFlag{
				Name:        "remote-bucket",
				Usage:       "Remote bucket ID for new replication stream",
				Required:    true,
				Destination: &params.RemoteBucketID,
			},
			&cli.Int64Flag{
				Name:        "max-queue",
				Usage:       "Max queue size in bytes",
				Value:       67108860,
				Destination: &params.MaxQueueSize,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)

			client := replication.Client{
				CLI:              getCLI(ctx),
				ReplicationsApi:  api.ReplicationsApi,
				OrganizationsApi: api.OrganizationsApi,
			}

			return client.Create(getContext(ctx), &params)
		},
	}
}

func newReplicationDeleteCmd() cli.Command {
	return cli.Command{
		Name:   "delete",
		Usage:  "Delete an existing replication stream",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags:  commonFlags(),
		Action: func(ctx *cli.Context) {
			fmt.Println("replication delete command was called")
		},
	}
}

func newReplicationListCmd() cli.Command {
	return cli.Command{
		Name:    "list",
		Usage:   "List all replication streams and corresponding metrics",
		Aliases: []string{"find", "ls"},
		Before:  middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags:   commonFlags(),
		Action: func(ctx *cli.Context) {
			fmt.Println("replication list command was called")
		},
	}
}

func newReplicationUpdateCmd() cli.Command {
	return cli.Command{
		Name:   "update",
		Usage:  "Update an existing replication stream",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags:  commonFlags(),
		Action: func(ctx *cli.Context) {
			fmt.Println("replication update command was called")
		},
	}
}

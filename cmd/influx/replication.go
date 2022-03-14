package main

import (
	"github.com/influxdata/influx-cli/v2/clients/replication"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newReplicationCmd() cli.Command {
	return cli.Command{
		Name:   "replication",
		Usage:  "Replication stream management commands",
		Before: middleware.NoArgs,
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
				Usage:       "Remote connection the new replication stream should send data to",
				Required:    true,
				Destination: &params.RemoteID,
			},
			&cli.StringFlag{
				Name:        "local-bucket",
				Usage:       "ID of local bucket data should be replicated from",
				Required:    true,
				Destination: &params.LocalBucketID,
			},
			&cli.StringFlag{
				Name:        "remote-bucket",
				Usage:       "ID of remote bucket data should be replicated to",
				Required:    true,
				Destination: &params.RemoteBucketID,
			},
			&cli.Int64Flag{
				Name:        "max-queue-bytes",
				Usage:       "Max queue size in bytes",
				Value:       67108860, // source: https://github.com/influxdata/openapi/blob/588064fe68e7dfeebd019695aa805832632cbfb6/src/oss/schemas/ReplicationCreationRequest.yml#L19
				Destination: &params.MaxQueueSize,
			},
			&cli.BoolFlag{
				Name:        "drop-non-retryable-data",
				Usage:       "Drop data when a non-retryable error is encountered instead of retrying",
				Destination: &params.DropNonRetryableData,
			},
			&cli.BoolFlag{
				Name:        "no-drop-non-retryable-data",
				Usage:       "Do not drop data when a non-retryable error is encountered",
				Destination: &params.NoDropNonRetryableData,
			},
			&cli.Int64Flag{
				Name:        "max-age",
				Usage:       "Specify a maximum age (in seconds) for replications data before it is dropped, or 0 for infinite",
				Value:       604800, // source: https://github.com/influxdata/openapi/blob/6ea7df4daa5735a063be3db60d0165b34b26c096/src/oss/schemas/ReplicationCreationRequest.yml#L27
				Destination: &params.MaxAge,
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
	var replicationID string
	return cli.Command{
		Name:   "delete",
		Usage:  "Delete an existing replication stream",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "id, i",
				Usage:       "ID of the replication stream to be deleted",
				Required:    true,
				Destination: &replicationID,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)

			client := replication.Client{
				CLI:             getCLI(ctx),
				ReplicationsApi: api.ReplicationsApi,
			}

			return client.Delete(getContext(ctx), replicationID)
		},
	}
}

func newReplicationListCmd() cli.Command {
	var params replication.ListParams
	return cli.Command{
		Name:    "list",
		Usage:   "List all replication streams and corresponding metrics",
		Aliases: []string{"find", "ls"},
		Before:  middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "Filter results to only replication streams with a specific name",
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "org-id",
				Usage:       "Local org ID",
				EnvVar:      "INFLUX_ORG_ID",
				Destination: &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "org, o",
				Usage:       "Local org name",
				EnvVar:      "INFLUX_ORG",
				Destination: &params.OrgName,
			},
			&cli.StringFlag{
				Name:        "remote-id",
				Usage:       "Filter results to only replication streams for a specific remote connection",
				Destination: &params.RemoteID,
			},
			&cli.StringFlag{
				Name:        "local-bucket",
				Usage:       "Filter results to only replication streams for a specific local bucket",
				Destination: &params.LocalBucketID,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)

			client := replication.Client{
				CLI:              getCLI(ctx),
				ReplicationsApi:  api.ReplicationsApi,
				OrganizationsApi: api.OrganizationsApi,
			}

			return client.List(getContext(ctx), &params)
		},
	}
}

func newReplicationUpdateCmd() cli.Command {
	var params replication.UpdateParams
	return cli.Command{
		Name:   "update",
		Usage:  "Update an existing replication stream",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "id, i",
				Usage:       "ID of the replication stream to be updated",
				Required:    true,
				Destination: &params.ReplicationID,
			},
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "New name for the replication stream",
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "description, d",
				Usage:       "New description for the replication stream",
				Destination: &params.Description,
			},
			&cli.StringFlag{
				Name:        "remote-id",
				Usage:       "New ID of remote connection the replication stream should send data to",
				Destination: &params.RemoteID,
			},
			&cli.StringFlag{
				Name:        "remote-bucket",
				Usage:       "New ID of remote bucket that data should be replicated to",
				Destination: &params.RemoteBucketID,
			},
			&cli.Int64Flag{
				Name:        "max-queue-bytes",
				Usage:       "New max queue size in bytes",
				Destination: &params.MaxQueueSize,
			},
			&cli.BoolFlag{
				Name:        "drop-non-retryable-data",
				Usage:       "Drop data when a non-retryable error is encountered instead of retrying",
				Destination: &params.DropNonRetryableData,
			},
			&cli.BoolFlag{
				Name:        "no-drop-non-retryable-data",
				Usage:       "Do not drop data when a non-retryable error is encountered",
				Destination: &params.NoDropNonRetryableData,
			},
			&cli.Int64Flag{
				Name:        "max-age",
				Usage:       "Specify a maximum age (in seconds) for replications data before it is dropped, or 0 for infinite",
				Destination: &params.MaxAge,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)

			client := replication.Client{
				CLI:             getCLI(ctx),
				ReplicationsApi: api.ReplicationsApi,
			}

			return client.Update(getContext(ctx), &params)
		},
	}
}

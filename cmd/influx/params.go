package main

import (
	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/urfave/cli/v2"
)

func getOrgBucketFlags(c *internal.OrgBucketParams) []cli.Flag {
	return []cli.Flag{
		&cli.GenericFlag{
			Name:    "bucket-id",
			Usage:   "The bucket ID, required if name isn't provided",
			Aliases: []string{"i"},
			Value:   &c.BucketID,
		},
		&cli.StringFlag{
			Name:        "bucket",
			Usage:       "The bucket name, org or org-id will be required by choosing this",
			Aliases:     []string{"n"},
			Destination: &c.BucketName,
		},
		&cli.GenericFlag{
			Name:    "org-id",
			Usage:   "The ID of the organization",
			EnvVars: []string{"INFLUX_ORG_ID"},
			Value:   &c.OrgID,
		},
		&cli.StringFlag{
			Name:        "org",
			Usage:       "The name of the organization",
			Aliases:     []string{"o"},
			EnvVars:     []string{"INFLUX_ORG"},
			Destination: &c.OrgName,
		},
	}
}

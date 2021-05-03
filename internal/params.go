package internal

import (
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
)

type OrgParams struct {
	OrgID   influxid.ID
	OrgName string
}

type BucketParams struct {
	BucketID   influxid.ID
	BucketName string
}

type OrgBucketParams struct {
	OrgParams
	BucketParams
}

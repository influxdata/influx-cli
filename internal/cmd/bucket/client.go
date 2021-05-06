package bucket

import (
	"errors"
	"time"

	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/internal/cmd"
)

const InfiniteRetention = 0

var (
	ErrMustSpecifyOrg             = errors.New("must specify org ID or org name")
	ErrMustSpecifyOrgDeleteByName = errors.New("must specify org ID or org name when deleting bucket by name")
	ErrMustSpecifyBucket          = errors.New("must specify bucket ID or bucket name")
)

type Client struct {
	cmd.CLI
	api.OrganizationsApi
	api.BucketsApi
}

type bucketPrintOptions struct {
	deleted bool
	bucket  *api.Bucket
	buckets []api.Bucket
}

func (c Client) printBuckets(options bucketPrintOptions) error {
	if c.PrintAsJSON {
		var v interface{}
		if options.bucket != nil {
			v = options.bucket
		} else {
			v = options.buckets
		}
		return c.PrintJSON(v)
	}

	headers := []string{"ID", "Name", "Retention", "Shard group duration", "Organization ID", "Schema Type"}
	if options.deleted {
		headers = append(headers, "Deleted")
	}

	if options.bucket != nil {
		options.buckets = append(options.buckets, *options.bucket)
	}

	var rows []map[string]interface{}
	for _, bkt := range options.buckets {
		var rpDuration time.Duration // zero value implies infinite retention policy
		var sgDuration time.Duration // zero value implies the server should pick a value

		if rules := bkt.GetRetentionRules(); len(rules) > 0 {
			rpDuration = time.Duration(rules[0].GetEverySeconds()) * time.Second
			sgDuration = time.Duration(rules[0].GetShardGroupDurationSeconds()) * time.Second
		}

		rp := rpDuration.String()
		if rpDuration == InfiniteRetention {
			rp = "infinite"
		}
		sgDur := sgDuration.String()
		// ShardGroupDuration will be zero if listing buckets from InfluxDB Cloud.
		// Show something more useful here in that case.
		if sgDuration == 0 {
			sgDur = "n/a"
		}

		schemaType := bkt.GetSchemaType()
		if schemaType == "" {
			// schemaType will be empty when querying OSS.
			schemaType = api.SCHEMATYPE_IMPLICIT
		}

		row := map[string]interface{}{
			"ID":                   bkt.GetId(),
			"Name":                 bkt.GetName(),
			"Retention":            rp,
			"Shard group duration": sgDur,
			"Organization ID":      bkt.GetOrgID(),
			"Schema Type":          schemaType,
		}
		if options.deleted {
			row["Deleted"] = true
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

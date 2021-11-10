package bucket

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type BucketsDeleteParams struct {
	clients.OrgBucketParams
}

func (c Client) Delete(ctx context.Context, params *BucketsDeleteParams) error {
	if params.BucketID == "" && params.BucketName == "" {
		return clients.ErrMustSpecifyBucket
	}

	var bucket api.Bucket
	var getReq api.ApiGetBucketsRequest
	if params.BucketID != "" {
		getReq = c.GetBuckets(ctx).Id(params.BucketID)
	} else {
		if params.OrgID == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
			return ErrMustSpecifyOrgDeleteByName
		}
		getReq = c.GetBuckets(ctx)
		getReq = getReq.Name(params.BucketName)
		if params.OrgID != "" {
			getReq = getReq.OrgID(params.OrgID)
		}
		if params.OrgName != "" {
			getReq = getReq.Org(params.OrgName)
		}
		if params.OrgID == "" && params.OrgName == "" {
			getReq = getReq.Org(c.ActiveConfig.Org)
		}
	}

	displayId := params.BucketID
	if displayId == "" {
		displayId = params.BucketName
	}

	resp, err := getReq.Execute()
	if err != nil {
		return fmt.Errorf("failed to find bucket %q: %w", displayId, err)
	}
	buckets := resp.GetBuckets()
	if len(buckets) == 0 {
		return fmt.Errorf("bucket %q not found", displayId)
	}
	bucket = buckets[0]

	if err := c.DeleteBucketsID(ctx, bucket.GetId()).Execute(); err != nil {
		return fmt.Errorf("failed to delete bucket %q: %w", displayId, err)
	}

	return c.printBuckets(bucketPrintOptions{bucket: &bucket, deleted: true})
}

package bucket

import (
	"context"
	"fmt"
	"time"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/pkg/duration"
)

type BucketsUpdateParams struct {
	ID                 string
	Name               string
	Description        string
	Retention          string
	ShardGroupDuration string
}

func (c Client) Update(ctx context.Context, params *BucketsUpdateParams) error {
	reqBody := api.PatchBucketRequest{}
	if params.Name != "" {
		reqBody.SetName(params.Name)
	}
	if params.Description != "" {
		reqBody.SetDescription(params.Description)
	}
	if params.Retention != "" || params.ShardGroupDuration != "" {
		patchRule := api.NewPatchRetentionRuleWithDefaults()
		if params.Retention != "" {
			rp, err := duration.RawDurationToTimeDuration(params.Retention)
			if err != nil {
				return err
			}
			patchRule.SetEverySeconds(int64(rp.Round(time.Second) / time.Second))
		}
		if params.ShardGroupDuration != "" {
			sgd, err := duration.RawDurationToTimeDuration(params.ShardGroupDuration)
			if err != nil {
				return err
			}
			patchRule.SetShardGroupDurationSeconds(int64(sgd.Round(time.Second) / time.Second))
		}
		reqBody.SetRetentionRules([]api.PatchRetentionRule{*patchRule})
	}

	bucket, err := c.PatchBucketsID(ctx, params.ID).PatchBucketRequest(reqBody).Execute()
	if err != nil {
		return fmt.Errorf("failed to update bucket %q: %w", params.ID, err)
	}

	return c.printBuckets(bucketPrintOptions{bucket: &bucket})
}

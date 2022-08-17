package bucket

import (
	"context"
	"fmt"
	"time"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/pkg/duration"
)

type BucketsCreateParams struct {
	clients.OrgParams
	Name               string
	Description        string
	Retention          string
	ShardGroupDuration string
	SchemaType         api.SchemaType
}

func (c Client) Create(ctx context.Context, params *BucketsCreateParams) error {
	orgId, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	if err != nil {
		return err
	}

	rp, err := duration.RawDurationToTimeDuration(params.Retention)
	if err != nil {
		return err
	}
	sgd, err := duration.RawDurationToTimeDuration(params.ShardGroupDuration)
	if err != nil {
		return err
	}

	var rr []api.RetentionRule
	reqBody := api.PostBucketRequest{
		OrgID:          orgId,
		Name:           params.Name,
		RetentionRules: &rr,
		SchemaType:     &params.SchemaType,
	}
	if params.Description != "" {
		reqBody.Description = &params.Description
	}
	// Only append a retention rule if the user wants to explicitly set
	// a parameter on the rule.
	//
	// This is for backwards-compatibility with older versions of the API,
	// which didn't support setting shard-group durations and used an empty
	// array of rules to represent infinite retention.
	if rp > 0 || sgd > 0 {
		rule := api.NewRetentionRuleWithDefaults()
		if rp > 0 {
			rule.SetEverySeconds(int64(rp.Round(time.Second) / time.Second))
		}
		if sgd > 0 {
			rule.SetShardGroupDurationSeconds(int64(sgd.Round(time.Second) / time.Second))
		}
		*reqBody.RetentionRules = append(*reqBody.RetentionRules, *rule)
	}

	bucket, err := c.PostBuckets(ctx).PostBucketRequest(reqBody).Execute()
	if err != nil {
		return fmt.Errorf("failed to create bucket: %w", err)
	}

	return c.printBuckets(bucketPrintOptions{bucket: &bucket})
}

package bucket

import (
	"context"
	"fmt"
	"time"

	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/internal/cmd"
	"github.com/influxdata/influx-cli/v2/internal/duration"
)

type BucketsCreateParams struct {
	OrgID              string
	OrgName            string
	Name               string
	Description        string
	Retention          string
	ShardGroupDuration string
	SchemaType         api.SchemaType
}

func (c Client) Create(ctx context.Context, params *BucketsCreateParams) error {
	if params.OrgID == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return cmd.ErrMustSpecifyOrg
	}

	rp, err := duration.RawDurationToTimeDuration(params.Retention)
	if err != nil {
		return err
	}
	sgd, err := duration.RawDurationToTimeDuration(params.ShardGroupDuration)
	if err != nil {
		return err
	}

	reqBody := api.PostBucketRequest{
		OrgID:          params.OrgID,
		Name:           params.Name,
		RetentionRules: []api.RetentionRule{},
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
		reqBody.RetentionRules = append(reqBody.RetentionRules, *rule)
	}
	if reqBody.OrgID == "" {
		name := params.OrgName
		if name == "" {
			name = c.ActiveConfig.Org
		}
		resp, err := c.GetOrgs(ctx).Org(name).Execute()
		if err != nil {
			return fmt.Errorf("failed to lookup ID of org %q: %w", name, err)
		}
		orgs := resp.GetOrgs()
		if len(orgs) == 0 {
			return fmt.Errorf("no organization found with name %q", name)
		}
		reqBody.OrgID = orgs[0].GetId()
	}

	bucket, err := c.PostBuckets(ctx).PostBucketRequest(reqBody).Execute()
	if err != nil {
		return fmt.Errorf("failed to create bucket: %w", err)
	}

	return c.printBuckets(bucketPrintOptions{bucket: &bucket})
}

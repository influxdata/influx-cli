package internal

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/internal/duration"
)

const InfiniteRetention = 0

type BucketsClients struct {
	OrgApi    api.OrganizationsApi
	BucketApi api.BucketsApi
}

type BucketsCreateParams struct {
	OrgID              string
	OrgName            string
	Name               string
	Description        string
	Retention          string
	ShardGroupDuration string
}

var (
	ErrMustSpecifyOrg             = errors.New("must specify org ID or org name")
	ErrMustSpecifyOrgDeleteByName = errors.New("must specify org ID or org name when deleting bucket by name")
	ErrMustSpecifyBucket          = errors.New("must specify bucket ID or bucket name")
)

func (c *CLI) BucketsCreate(ctx context.Context, clients *BucketsClients, params *BucketsCreateParams) error {
	if params.OrgID == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return ErrMustSpecifyOrg
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
		lookupReq := clients.OrgApi.GetOrgs(ctx).Org(name)
		resp, _, err := clients.OrgApi.GetOrgsExecute(lookupReq)
		if err != nil {
			return fmt.Errorf("failed to lookup ID of org %q: %w", name, err)
		}
		orgs := resp.GetOrgs()
		if len(orgs) == 0 {
			return fmt.Errorf("no organization found with name %q", name)
		}
		reqBody.OrgID = orgs[0].GetId()
	}

	req := clients.BucketApi.PostBuckets(ctx).PostBucketRequest(reqBody)
	bucket, _, err := clients.BucketApi.PostBucketsExecute(req)
	if err != nil {
		return fmt.Errorf("failed to create bucket: %w", err)
	}

	return c.printBuckets(bucketPrintOptions{bucket: &bucket})
}

type BucketsListParams struct {
	OrgID   string
	OrgName string
	Name    string
	ID      string
}

func (c *CLI) BucketsList(ctx context.Context, client api.BucketsApi, params *BucketsListParams) error {
	if params.OrgID == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return ErrMustSpecifyOrg
	}

	req := client.GetBuckets(ctx)
	if params.OrgID != "" {
		req = req.OrgID(params.OrgID)
	}
	if params.OrgName != "" {
		req = req.Org(params.OrgName)
	}
	if params.OrgID == "" && params.OrgName == "" {
		req = req.Org(c.ActiveConfig.Org)
	}
	if params.Name != "" {
		req = req.Name(params.Name)
	}
	if params.ID != "" {
		req = req.Id(params.ID)
	}

	buckets, _, err := client.GetBucketsExecute(req)
	if err != nil {
		return fmt.Errorf("failed to list buckets: %w", err)
	}

	printOpts := bucketPrintOptions{}
	if buckets.Buckets != nil {
		printOpts.buckets = *buckets.Buckets
	}
	return c.printBuckets(printOpts)
}

type BucketsUpdateParams struct {
	ID                 string
	Name               string
	Description        string
	Retention          string
	ShardGroupDuration string
}

func (c *CLI) BucketsUpdate(ctx context.Context, client api.BucketsApi, params *BucketsUpdateParams) error {
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

	req := client.PatchBucketsID(ctx, params.ID).PatchBucketRequest(reqBody)
	bucket, _, err := client.PatchBucketsIDExecute(req)
	if err != nil {
		return fmt.Errorf("failed to update bucket %q: %w", params.ID, err)
	}

	return c.printBuckets(bucketPrintOptions{bucket: &bucket})
}

type BucketsDeleteParams struct {
	ID      string
	Name    string
	OrgID   string
	OrgName string
}

func (c *CLI) BucketsDelete(ctx context.Context, client api.BucketsApi, params *BucketsDeleteParams) error {
	if params.ID == "" && params.Name == "" {
		return ErrMustSpecifyBucket
	}

	var bucket api.Bucket
	getReq := client.GetBuckets(ctx)
	if params.ID != "" {
		getReq = getReq.Id(params.ID)
	} else {
		if params.OrgID == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
			return ErrMustSpecifyOrgDeleteByName
		}
		getReq = getReq.Name(params.Name)
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

	displayId := params.ID
	if displayId == "" {
		displayId = params.Name
	}

	resp, _, err := client.GetBucketsExecute(getReq)
	if err != nil {
		return fmt.Errorf("failed to find bucket %q: %w", displayId, err)
	}
	buckets := resp.GetBuckets()
	if len(buckets) == 0 {
		return fmt.Errorf("bucket %q not found", displayId)
	}
	bucket = buckets[0]

	req := client.DeleteBucketsID(ctx, bucket.GetId())
	if _, err := client.DeleteBucketsIDExecute(req); err != nil {
		return fmt.Errorf("failed to delete bucket %q: %w", displayId, err)
	}

	return c.printBuckets(bucketPrintOptions{bucket: &bucket, deleted: true})
}

type bucketPrintOptions struct {
	deleted bool
	bucket  *api.Bucket
	buckets []api.Bucket
}

func (c *CLI) printBuckets(options bucketPrintOptions) error {
	if c.PrintAsJSON {
		var v interface{}
		if options.bucket != nil {
			v = options.bucket
		} else {
			v = options.buckets
		}
		return c.PrintJSON(v)
	}

	headers := []string{"ID", "Name", "Retention", "Shard group duration", "Organization ID"}
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

		row := map[string]interface{}{
			"ID":                   bkt.GetId(),
			"Name":                 bkt.GetName(),
			"Retention":            rp,
			"Shard group duration": sgDur,
			"Organization ID":      bkt.GetOrgID(),
		}
		if options.deleted {
			row["Deleted"] = true
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

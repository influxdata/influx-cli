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

func (c *CLI) BucketsCreate(ctx context.Context, clients *BucketsClients, params *BucketsCreateParams) error {
	if params.OrgID == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return errors.New("must specify org ID or org name")
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
		rule.SetEverySeconds(int64(rp.Round(time.Second) / time.Second))
		rule.SetShardGroupDurationSeconds(int64(sgd.Round(time.Second) / time.Second))
		reqBody.RetentionRules = append(reqBody.RetentionRules, *rule)
	}
	if reqBody.OrgID == "" {
		name := params.OrgName
		if name == "" {
			name = c.ActiveConfig.Org
		}
		lookupReq := clients.OrgApi.GetOrgs(ctx).Org(name)
		if c.TraceId != "" {
			lookupReq = lookupReq.ZapTraceSpan(c.TraceId)
		}
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
	if c.TraceId != "" {
		req = req.ZapTraceSpan(c.TraceId)
	}
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

func (c *CLI) BucketsList(ctx context.Context, clients *BucketsClients, params *BucketsListParams) error {
	if params.OrgID == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return errors.New("must specify org ID or org name")
	}

	req := clients.BucketApi.GetBuckets(ctx)
	if c.TraceId != "" {
		req = req.ZapTraceSpan(c.TraceId)
	}
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

	buckets, _, err := clients.BucketApi.GetBucketsExecute(req)
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

func (c *CLI) BucketsUpdate(ctx context.Context, clients *BucketsClients, params *BucketsUpdateParams) error {
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

	req := clients.BucketApi.PatchBucketsID(ctx, params.ID).PatchBucketRequest(reqBody)
	if c.TraceId != "" {
		req = req.ZapTraceSpan(c.TraceId)
	}

	bucket, _, err := clients.BucketApi.PatchBucketsIDExecute(req)
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

func (c *CLI) BucketsDelete(ctx context.Context, clients *BucketsClients, params *BucketsDeleteParams) error {
	id := params.ID
	if id == "" {
		if params.OrgID == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
			return errors.New("must specify org ID or org name when deleting bucket by name")
		}
		req := clients.BucketApi.GetBuckets(ctx).Name(params.Name)
		if c.TraceId != "" {
			req = req.ZapTraceSpan(c.TraceId)
		}
		if params.OrgID != "" {
			req = req.OrgID(params.OrgID)
		}
		if params.OrgName != "" {
			req = req.Org(params.OrgName)
		}
		if params.OrgID == "" && params.OrgName == "" {
			req = req.Org(c.ActiveConfig.Org)
		}

		resp, _, err := clients.BucketApi.GetBucketsExecute(req)
		if err != nil {
			return fmt.Errorf("failed to find bucket %q: %w", params.Name, err)
		}
		buckets := resp.GetBuckets()
		if len(buckets) == 0 {
			return fmt.Errorf("no bucket found with name %q", params.Name)
		}
		id = buckets[0].GetId()
	}

	getReq := clients.BucketApi.GetBucketsID(ctx, id)
	if c.TraceId != "" {
		getReq = getReq.ZapTraceSpan(c.TraceId)
	}
	bucket, _, err := clients.BucketApi.GetBucketsIDExecute(getReq)
	if err != nil {
		return fmt.Errorf("failed to find bucket with ID %q: %w", id, err)
	}

	req := clients.BucketApi.DeleteBucketsID(ctx, id)
	if c.TraceId != "" {
		req = req.ZapTraceSpan(c.TraceId)
	}
	if _, err := clients.BucketApi.DeleteBucketsIDExecute(req); err != nil {
		return fmt.Errorf("failed to delete bucket with ID %q: %w", id, err)
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
		var v interface{} = options.buckets
		if v == nil {
			v = options.bucket
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

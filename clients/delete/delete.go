package delete

import (
	"context"
	"fmt"
	"time"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	clients.CLI
	api.DeleteApi
}

type Params struct {
	clients.OrgBucketParams
	Predicate string
	Start     string
	Stop      string
}

func (c Client) Delete(ctx context.Context, params *Params) error {
	if params.OrgID == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return clients.ErrMustSpecifyOrg
	}
	if params.BucketID == "" && params.BucketName == "" {
		return clients.ErrMustSpecifyBucket
	}
	start, err := time.Parse(time.RFC3339Nano, params.Start)
	if err != nil {
		return fmt.Errorf("start time %q cannot be parsed as RFC3339Nano: %w", params.Start, err)
	}
	stop, err := time.Parse(time.RFC3339Nano, params.Stop)
	if err != nil {
		return fmt.Errorf("stop time %q cannot be parsed as RFC3339Nano: %w", params.Stop, err)
	}

	reqBody := api.NewDeletePredicateRequest(start, stop)
	if params.Predicate != "" {
		reqBody.SetPredicate(params.Predicate)
	}

	req := c.PostDelete(ctx).DeletePredicateRequest(*reqBody)
	if params.OrgID != "" {
		req = req.OrgID(params.OrgID)
	} else if params.OrgName != "" {
		req = req.Org(params.OrgName)
	} else {
		req = req.Org(c.ActiveConfig.Org)
	}
	if params.BucketID != "" {
		req = req.BucketID(params.BucketID)
	} else {
		req = req.Bucket(params.BucketName)
	}

	if err := req.Execute(); err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}

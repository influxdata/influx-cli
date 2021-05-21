package delete

import (
	"context"
	"errors"
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

var ErrMustSpecifyBucket = errors.New("must specify bucket ID or bucket name")

func (c Client) Delete(ctx context.Context, params *Params) error {
	if !params.OrgID.Valid() && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return clients.ErrMustSpecifyOrg
	}
	if !params.BucketID.Valid() && params.BucketName == "" {
		return ErrMustSpecifyBucket
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
	if params.OrgID.Valid() {
		req = req.OrgID(params.OrgID.String())
	} else if params.OrgName != "" {
		req = req.Org(params.OrgName)
	} else {
		req = req.Org(c.ActiveConfig.Org)
	}
	if params.BucketID.Valid() {
		req = req.BucketID(params.BucketID.String())
	} else {
		req = req.Bucket(params.BucketName)
	}

	if err := req.Execute(); err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}

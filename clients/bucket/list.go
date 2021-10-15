package bucket

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type BucketsListParams struct {
	OrgID    string
	OrgName  string
	Name     string
	ID       string
	Limit    int
	After    string
	Offset   int
	PageSize int
}

func (c Client) List(ctx context.Context, params *BucketsListParams) error {
	if params.OrgID == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return clients.ErrMustSpecifyOrg
	}

	req := c.GetBuckets(ctx)
	if params.OrgID != "" {
		req = req.OrgID(params.OrgID)
	}
	if params.OrgName != "" {
		req = req.Org(params.OrgName)
	}
	if params.OrgID == "" && params.OrgName == "" {
		req = req.Org(c.ActiveConfig.Org)
	}
	if params.Name != "" || params.ID != "" {
		return c.findOneBucket(params, req)
	}

	printOpts := bucketPrintOptions{}

	// Set the limit for the number of items to return per HTTP request.
	// NOTE this is not the same as the `--limit` option to the CLI, which sets
	// the max number of rows to pull across _all_ HTTP requests in total.
	limit := params.PageSize
	// Adjust if the total limit < the per-request limit.
	// This is convenient for users, since the per-request limit has a default
	// value that people might not want to override on every CLI call.
	if params.Limit != 0 && params.Limit < limit {
		limit = params.Limit
	}
	offset := params.Offset
	after := params.After

	// To guard against infinite loops from bugs in/incomplete implementations of server-side pagination,
	// we track the IDs we see in each request and bail out if we get a duplicate.
	seenIds := map[string]struct{}{}
	if after != "" {
		seenIds[after] = struct{}{}
	}

	for {
		req := req
		if limit != 0 {
			req = req.Limit(int32(limit))
		}
		if offset != 0 {
			req = req.Offset(int32(offset))
		}
		if after != "" {
			req = req.After(after)
		}

		res, err := req.Execute()
		if err != nil {
			return fmt.Errorf("failed to list buckets: %w", err)
		}
		var buckets []api.Bucket
		if res.Buckets != nil {
			buckets = *res.Buckets
		}
		bailOut := false
		for _, b := range buckets {
			id := *b.Id
			if _, alreadySeen := seenIds[id]; alreadySeen {
				bailOut = true
			} else {
				printOpts.buckets = append(printOpts.buckets, b)
				seenIds[id] = struct{}{}
			}
		}
		// If pagination appears to be broken OR if the server returned fewer results than we asked for,
		// break out of the loop because making additional requests won't give us any more information.
		if bailOut || len(buckets) == 0 || len(buckets) < params.PageSize {
			break
		}
		// If we've collected as many results as the user asked for, break out of the loop.
		if params.Limit != 0 && len(printOpts.buckets) == params.Limit {
			break
		}

		// Adjust the page-size for the next request (if needed) so we don't pull down more information
		// than the user requested.
		if params.Limit != 0 && len(printOpts.buckets)+limit > params.Limit {
			limit = params.Limit - len(printOpts.buckets)
		}
		// Use `after`-based pagination following the initial request.
		after = *buckets[len(buckets)-1].Id
		offset = 0
	}

	return c.printBuckets(printOpts)
}

// findOneBucket queries for a single bucket resoruce using the list API.
// Used to look up buckets by ID or name.
func (c Client) findOneBucket(params *BucketsListParams, req api.ApiGetBucketsRequest) error {
	var description string
	if params.ID != "" {
		req = req.Id(params.ID)
		description = " by ID"
	} else if params.Name != "" {
		req = req.Name(params.Name)
		description = " by name"
	}

	buckets, err := req.Execute()
	if err != nil {
		return fmt.Errorf("failed to find bucket%s: %w", description, err)
	}

	printOpts := bucketPrintOptions{}
	if buckets.Buckets != nil {
		printOpts.buckets = *buckets.Buckets
	}
	return c.printBuckets(printOpts)
}

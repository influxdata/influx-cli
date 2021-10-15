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

	// To guard against infinite loops from bugs in/incomplete implementations of server-side pagination,
	// we track the IDs we see in each request and bail out if we get a duplicate.
	seenIds := map[string]struct{}{}

	// Iteratively fetch pages of bucket metadata.
	//
	// NOTE: This algorithm used an `offset`-based pagination. The API also documents that an `after`-based
	// approach would work. Pagination via `after` would typically be much more efficient than the `offset`-
	// based approach used here, but we still use `offset` because:
	//   1. As of this writing (10/15/2021) Cloud doesn't support pagination via `after` when filtering by org.
	//      Versions of OSS prior to v2.1 had the same restriction. Both Cloud and OSS have supported `offset`-
	//      based pagination since before the OSS/Cloud split.
	//   2. Listing buckets-by-org uses a secondary index on orgID+bucketName in both OSS and Cloud. Since bucket
	//      ID isn't part of the index, implementing pagination via `after` requires scanning the subset of the
	//      index for the target org until the given ID is found. This is just as (in)efficient as the implementation
	//      of `offset`-based filtering.
	//
	// If you are thinking of copy-paste-adjusting this code for another command, CONSIDER USING `after` INSTEAD OF
	// `offset`! The conditions that make `offset` OK for buckets probably aren't present for other resource types.
	//
	// We should eventually convert this to use `after` for consistency across the CLI.
	for {
		req := req
		if limit != 0 {
			req = req.Limit(int32(limit))
		}
		if offset != 0 {
			req = req.Offset(int32(offset))
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
		// Bump the offset for the next request to pull the next page.
		offset = offset + len(buckets)
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

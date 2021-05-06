package bucket

import (
	"context"
	"fmt"
)

type BucketsListParams struct {
	OrgID   string
	OrgName string
	Name    string
	ID      string
}

func (c Client) List(ctx context.Context, params *BucketsListParams) error {
	if params.OrgID == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return ErrMustSpecifyOrg
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
	if params.Name != "" {
		req = req.Name(params.Name)
	}
	if params.ID != "" {
		req = req.Id(params.ID)
	}

	buckets, err := req.Execute()
	if err != nil {
		return fmt.Errorf("failed to list buckets: %w", err)
	}

	printOpts := bucketPrintOptions{}
	if buckets.Buckets != nil {
		printOpts.buckets = *buckets.Buckets
	}
	return c.printBuckets(printOpts)
}

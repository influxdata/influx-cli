package v1dbrps

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	clients.CLI
	api.DBRPsApi
}

type dbrpPrintOpts struct {
	dbrp  *api.DBRP
	dbrps []api.DBRP
}

type ListParams struct {
	clients.OrgParams
	ID       string
	BucketID string
	Default_ bool
	DB       string
	RP       string
}

func (c Client) List(ctx context.Context, params *ListParams) error {
	if !params.OrgID.Valid() && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return clients.ErrMustSpecifyOrg
	}

	req := c.GetDBRPs(ctx)
	if params.OrgID.Valid() {
		req = req.OrgID(params.OrgID.String())
	}
	if params.OrgName != "" {
		req = req.Org(params.OrgName)
	}
	if !params.OrgID.Valid() && params.OrgName == "" {
		req = req.Org(c.ActiveConfig.Org)
	}

	if params.ID != "" {
		req = req.Id(params.ID)
	}
	if params.BucketID != "" {
		req = req.BucketID(params.BucketID)
	}
	if params.RP != "" {
		req = req.Rp(params.RP)
	}
	if params.DB != "" {
		req = req.Db(params.DB)
	}

	// set this parameter if the --default flag was passed to true. this will
	// list only default DBRPs. Otherwise, don't set the parameter at all to list DBRPs
	// that are default and not default.
	if params.Default_ {
		req = req.Default_(params.Default_)
	}

	dbrps, err := req.Execute()
	if err != nil {
		return fmt.Errorf("failed to list dbrps: %w", err)
	}

	c.printDBRPs(dbrpPrintOpts{dbrps: dbrps.GetContent()})

	return nil
}

type CreateParams struct {
	clients.OrgParams
	BucketID string
	Default_ bool
	DB       string
	RP       string
}

func (c Client) Create(ctx context.Context, params *CreateParams) error {
	if !params.OrgID.Valid() && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return clients.ErrMustSpecifyOrg
	}

	reqBody := api.DBRPCreate{
		BucketID:        params.BucketID,
		Database:        params.DB,
		RetentionPolicy: &params.RP,
	}

	if params.OrgID.Valid() {
		reqBody.OrgID = params.OrgID.String()
	}
	if params.OrgName != "" {
		reqBody.Org = &params.OrgName
	}
	if !params.OrgID.Valid() && params.OrgName == "" {
		reqBody.Org = &c.ActiveConfig.Org
	}

	if params.OrgID.Valid() {
		reqBody.OrgID = params.OrgID.String()
	}
	if params.OrgName != "" {
		reqBody.Org = &params.OrgName
	}

	dbrp, err := c.PostDBRP(ctx).DBRPCreate(reqBody).Execute()
	if err != nil {
		return fmt.Errorf("failed to create dbrp for bucket %q: %w", params.BucketID, err)
	}

	c.printDBRPs(dbrpPrintOpts{dbrp: &dbrp})

	return nil
}

type UpdateParams struct {
	clients.OrgParams
	ID       string
	Default_ bool
	RP       string
}

func (c Client) Update(ctx context.Context, params *UpdateParams) error {
	if !params.OrgID.Valid() && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return clients.ErrMustSpecifyOrg
	}

	reqBody := api.DBRPUpdate{}

	if params.RP != "" {
		reqBody.RetentionPolicy = &params.RP
	}
	if params.Default_ {
		reqBody.Default = &params.Default_
	}

	req := c.PatchDBRPID(ctx, params.ID)
	if params.OrgID.Valid() {
		req = req.OrgID(params.OrgID.String())
	}
	if params.OrgName != "" {
		req = req.Org(params.OrgName)
	}
	if !params.OrgID.Valid() && params.OrgName == "" {
		req = req.Org(c.ActiveConfig.Org)
	}

	dbrp, err := req.DBRPUpdate(reqBody).Execute()
	if err != nil {
		return fmt.Errorf("failed to update DBRP mapping %q: %w", params.ID, err)
	}

	dbrpContent := dbrp.GetContent()
	c.printDBRPs(dbrpPrintOpts{dbrp: &dbrpContent})

	return nil
}

type DeleteParams struct {
	clients.OrgParams
	ID string
}

func (c Client) Delete(ctx context.Context, params *DeleteParams) error {
	if !params.OrgID.Valid() && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return clients.ErrMustSpecifyOrg
	}

	// we need to get the DBRP to verify that it exists, and to be able to print
	// the results from the deleted dbrp.
	getReq := c.GetDBRPsID(ctx, params.ID)
	deleteReq := c.DeleteDBRPID(ctx, params.ID)

	if params.OrgID.Valid() {
		getReq = getReq.OrgID(params.OrgID.String())
		deleteReq = deleteReq.OrgID(params.OrgID.String())
	}
	if params.OrgName != "" {
		getReq = getReq.Org(params.OrgName)
		deleteReq = deleteReq.Org(params.OrgName)
	}
	if !params.OrgID.Valid() && params.OrgName == "" {
		getReq = getReq.Org(c.ActiveConfig.Org)
		deleteReq = deleteReq.Org(c.ActiveConfig.Org)
	}

	dbrp, err := getReq.Execute()
	if err != nil {
		return fmt.Errorf("failed to find DBRP mapping %q: %w", params.ID, err)
	}

	if err := deleteReq.Execute(); err != nil {
		return fmt.Errorf("failed to delete DBRP mapping %q: %w", params.ID, err)
	}

	dbrpContent := dbrp.GetContent()
	c.printDBRPs(dbrpPrintOpts{dbrp: &dbrpContent})

	return nil
}

func (c Client) printDBRPs(opts dbrpPrintOpts) error {
	if c.PrintAsJSON {
		var v interface{} = opts.dbrps
		if opts.dbrp != nil {
			v = opts.dbrp
		}
		return c.PrintJSON(v)
	}

	headers := []string{
		"ID",
		"Database",
		"Bucket ID",
		"Retention Policy",
		"Default",
		"Organization ID",
	}

	if opts.dbrp != nil {
		opts.dbrps = append(opts.dbrps, *opts.dbrp)
	}

	var rows []map[string]interface{}
	for _, t := range opts.dbrps {
		row := map[string]interface{}{
			"ID":               *t.Id,
			"Database":         *t.Database,
			"Retention Policy": *t.RetentionPolicy,
			"Default":          *t.Default,
			"Organization ID":  *t.OrgID,
			"Bucket ID":        *t.BucketID,
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

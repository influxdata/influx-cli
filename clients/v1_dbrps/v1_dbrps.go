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
	api.OrganizationsApi
}

type dbrpPrintOpts struct {
	dbrp  *api.DBRP
	dbrps []api.DBRP
}

type ListParams struct {
	clients.OrgParams
	ID       string
	BucketID string
	Default  bool
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

	// Set this parameter if the --default flag was passed. This will list only
	// default DBRPs. Otherwise, don't set the parameter at all to list DBRPs that
	// are default and not default.
	if params.Default {
		req = req.Default_(params.Default) // Note: codegen sets this property as "Default_" instead of "Default"
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
	Default  bool
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
		RetentionPolicy: params.RP,
	}

	// For compatibility with the cloud API for creating a DBRP, an org ID must be
	// provided. The ID will be obtained based on the org name if an
	// org name is provided but no ID is.
	if params.OrgID.Valid() {
		reqBody.OrgID = api.PtrString(params.OrgID.String())
	} else {
		orgName := params.OrgName
		if orgName == "" {
			orgName = c.ActiveConfig.Org
		}
		res, err := c.GetOrgs(ctx).Org(orgName).Execute()
		if err != nil {
			return fmt.Errorf("failed to look up ID for org %q: %w", orgName, err)
		}
		if len(res.GetOrgs()) == 0 {
			return fmt.Errorf("no org found with name %q", orgName)
		}
		reqBody.OrgID = api.PtrString(res.GetOrgs()[0].GetId())
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
	ID      string
	Default bool
	RP      string
}

func (c Client) Update(ctx context.Context, params *UpdateParams) error {
	if !params.OrgID.Valid() && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return clients.ErrMustSpecifyOrg
	}

	reqBody := api.DBRPUpdate{}
	if params.RP != "" {
		reqBody.RetentionPolicy = &params.RP
	}
	if params.Default {
		reqBody.Default = &params.Default
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

	// Get the DBRP to verify that it exists, and to be able to print the results
	// of the delete command, which output the details of the deleted DBRP.
	getReq := c.GetDBRPsID(ctx, params.ID)
	deleteReq := c.DeleteDBRPID(ctx, params.ID)

	// The org name or ID must be set on requests for OSS because of how the OSS
	// authorization mechanism currently works.
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
			"ID":               t.Id,
			"Database":         t.Database,
			"Retention Policy": t.RetentionPolicy,
			"Default":          t.Default,
			"Organization ID":  t.OrgID,
			"Bucket ID":        t.BucketID,
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

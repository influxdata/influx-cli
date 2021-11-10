package org

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type CreateParams struct {
	Name        string
	Description string
}

func (c Client) Create(ctx context.Context, params *CreateParams) error {
	body := api.PostOrganizationRequest{Name: params.Name}
	if params.Description != "" {
		body.Description = &params.Description
	}
	res, err := c.PostOrgs(ctx).PostOrganizationRequest(body).Execute()
	if err != nil {
		return fmt.Errorf("failed to create org %q: %w", params.Name, err)
	}
	return c.printOrgs(printOrgOpts{org: &res})
}

func (c Client) Delete(ctx context.Context, id string) error {
	org, err := c.GetOrgsID(ctx, id).Execute()
	if err != nil {
		return fmt.Errorf("org %q not found: %w", id, err)

	}
	if err := c.DeleteOrgsID(ctx, id).Execute(); err != nil {
		return fmt.Errorf("failed to delete org %q: %w", id, err)
	}
	return c.printOrgs(printOrgOpts{org: &org, deleted: true})
}

type ListParams struct {
	clients.OrgParams
}

func (c Client) List(ctx context.Context, params *ListParams) error {
	req := c.GetOrgs(ctx)
	if params.OrgName != "" {
		req = req.Org(params.OrgName)
	}
	if params.OrgID != "" {
		req = req.OrgID(params.OrgID)
	}
	orgs, err := req.Execute()
	if err != nil {
		return fmt.Errorf("failed to list orgs: %w", err)
	}
	printOpts := printOrgOpts{}
	if orgs.Orgs != nil {
		printOpts.orgs = *orgs.Orgs
	}
	return c.printOrgs(printOpts)
}

type UpdateParams struct {
	clients.OrgParams
	Description string
}

func (c Client) Update(ctx context.Context, params *UpdateParams) error {
	body := api.PatchOrganizationRequest{}
	if params.OrgName != "" {
		body.Name = &params.OrgName
	}
	if params.Description != "" {
		body.Description = &params.Description
	}

	res, err := c.PatchOrgsID(ctx, params.OrgID).PatchOrganizationRequest(body).Execute()
	if err != nil {
		return fmt.Errorf("failed to update org %q: %w", params.OrgID, err)
	}
	return c.printOrgs(printOrgOpts{org: &res})
}

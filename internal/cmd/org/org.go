package org

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
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

func (c Client) Delete(ctx context.Context, id influxid.ID) error {
	org, err := c.GetOrgsID(ctx, id.String()).Execute()
	if err != nil {
		return fmt.Errorf("org %q not found", id)

	}
	if err := c.DeleteOrgsID(ctx, id.String()).Execute(); err != nil {
		return fmt.Errorf("failed to delete org %q: %w", id, err)
	}
	return c.printOrgs(printOrgOpts{org: &org, deleted: true})
}

type ListParams struct {
	Name string
	ID   influxid.ID
}

func (c Client) List(ctx context.Context, params *ListParams) error {
	req := c.GetOrgs(ctx)
	if params.Name != "" {
		req = req.Org(params.Name)
	}
	if params.ID.Valid() {
		req = req.OrgID(params.ID.String())
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
	ID          influxid.ID
	Name        string
	Description string
}

func (c Client) Update(ctx context.Context, params *UpdateParams) error {
	body := api.PatchOrganizationRequest{}
	if params.Name != "" {
		body.Name = &params.Name
	}
	if params.Description != "" {
		body.Description = &params.Description
	}

	res, err := c.PatchOrgsID(ctx, params.ID.String()).PatchOrganizationRequest(body).Execute()
	if err != nil {
		return fmt.Errorf("failed to update org %q: %w", params.ID, err)
	}
	return c.printOrgs(printOrgOpts{org: &res})
}

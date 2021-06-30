package stacks

import (
	"context"
	"fmt"
	"sort"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/pkg/template"
)

type Client struct {
	clients.CLI
	api.StacksApi
	api.OrganizationsApi
	api.TemplatesApi
}

type ListParams struct {
	OrgId   string
	OrgName string

	StackIds   []string
	StackNames []string
}

func (c Client) List(ctx context.Context, params *ListParams) error {
	if params.OrgId == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return clients.ErrMustSpecifyOrg
	}

	orgId := params.OrgId
	if orgId == "" {
		orgName := params.OrgName
		if orgName == "" {
			orgName = c.ActiveConfig.Org
		}
		res, err := c.GetOrgs(ctx).Org(orgName).Execute()
		if err != nil {
			return fmt.Errorf("failed to lookup org with name %q: %w", orgName, err)
		}
		if len(res.GetOrgs()) == 0 {
			return fmt.Errorf("no organization with name %q: %w", orgName, err)
		}
		orgId = res.GetOrgs()[0].GetId()
	}

	res, err := c.ListStacks(ctx).OrgID(orgId).Name(params.StackNames).StackID(params.StackIds).Execute()
	if err != nil {
		return fmt.Errorf("failed to list stacks: %w", err)
	}

	return c.printStacks(stackPrintOptions{stacks: &res})
}

type InitParams struct {
	OrgId   string
	OrgName string

	Name        string
	Description string
	URLs        []string
}

func (c Client) Init(ctx context.Context, params *InitParams) error {
	if params.OrgId == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return clients.ErrMustSpecifyOrg
	}

	orgId := params.OrgId
	if orgId == "" {
		orgName := params.OrgName
		if orgName == "" {
			orgName = c.ActiveConfig.Org
		}
		res, err := c.GetOrgs(ctx).Org(orgName).Execute()
		if err != nil {
			return fmt.Errorf("failed to lookup org with name %q: %w", orgName, err)
		}
		if len(res.GetOrgs()) == 0 {
			return fmt.Errorf("no organization with name %q: %w", orgName, err)
		}
		orgId = res.GetOrgs()[0].GetId()
	}

	req := api.StackPostRequest{
		OrgID: orgId,
		Name:  params.Name,
		Urls:  params.URLs,
	}
	if params.Description != "" {
		req.Description = &params.Description
	}

	stack, err := c.CreateStack(ctx).StackPostRequest(req).Execute()
	if err != nil {
		return fmt.Errorf("failed to create stack %q: %w", params.Name, err)
	}

	return c.printStacks(stackPrintOptions{stack: &stack})
}

type RemoveParams struct {
	OrgId   string
	OrgName string

	Ids   []string
	Force bool
}

func (c Client) Remove(ctx context.Context, params *RemoveParams) error {
	if params.OrgId == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return clients.ErrMustSpecifyOrg
	}

	orgId := params.OrgId
	if orgId == "" {
		orgName := params.OrgName
		if orgName == "" {
			orgName = c.ActiveConfig.Org
		}
		res, err := c.GetOrgs(ctx).Org(orgName).Execute()
		if err != nil {
			return fmt.Errorf("failed to lookup org with name %q: %w", orgName, err)
		}
		if len(res.GetOrgs()) == 0 {
			return fmt.Errorf("no organization with name %q: %w", orgName, err)
		}
		orgId = res.GetOrgs()[0].GetId()
	}

	stacks, err := c.ListStacks(ctx).OrgID(orgId).StackID(params.Ids).Execute()
	if err != nil {
		return fmt.Errorf("failed to look up stacks: %w", err)
	}

	for _, stack := range stacks.Stacks {
		if err := c.printStacks(stackPrintOptions{stack: &stack}); err != nil {
			return err
		}
		if !params.Force && !c.StdIO.GetConfirm(fmt.Sprintf("Confirm removal of the stack[%s] and all associated resources", stack.Id)) {
			continue
		}
		if err := c.DeleteStack(ctx, stack.Id).OrgID(orgId).Execute(); err != nil {
			return fmt.Errorf("failed to delete stack %q: %w", stack.Id, err)
		}
	}

	return nil
}

type AddedResource struct {
	Kind string
	Id   string
}

type UpdateParams struct {
	Id string

	Name           *string
	Description    *string
	URLs           []string
	AddedResources []AddedResource

	template.OutParams
}

func (c Client) Update(ctx context.Context, params *UpdateParams) error {
	req := api.StackPatchRequest{
		Name:                params.Name,
		Description:         params.Description,
		TemplateURLs:        params.URLs,
		AdditionalResources: make([]api.StackPatchRequestResource, len(params.AddedResources)),
	}
	for i, r := range params.AddedResources {
		req.AdditionalResources[i] = api.StackPatchRequestResource{
			ResourceID: r.Id,
			Kind:       r.Kind,
		}
	}

	stack, err := c.UpdateStack(ctx, params.Id).StackPatchRequest(req).Execute()
	if err != nil {
		return fmt.Errorf("failed to udpate stack %q: %w", params.Id, err)
	}
	if err := c.printStacks(stackPrintOptions{stack: &stack}); err != nil {
		return err
	}

	// Can skip exporting the updated template if no resources were added.
	if len(params.AddedResources) == 0 {
		return nil
	}

	if !c.StdIO.GetConfirm(`Your stack now differs from your template. Applying an outdated template will revert
these updates. Export a new template with these updates to prevent accidental changes?`) {
		return nil
	}

	exportReq := api.TemplateExport{StackID: &stack.Id}
	tmpl, err := c.ExportTemplate(ctx).TemplateExport(exportReq).Execute()
	if err != nil {
		return fmt.Errorf("failed to export stack %q: %w", stack.Id, err)
	}
	if err := params.OutParams.WriteTemplate(tmpl); err != nil {
		return fmt.Errorf("failed to write exported template: %w", err)
	}
	return nil
}

type stackPrintOptions struct {
	stack  *api.Stack
	stacks *api.Stacks
}

func (c Client) printStacks(options stackPrintOptions) error {
	if c.PrintAsJSON {
		var v interface{}
		if options.stack != nil {
			v = options.stack
		} else {
			v = options.stacks
		}
		return c.PrintJSON(v)
	}

	headers := []string{"ID", "OrgID", "Active", "Name", "Description", "Num Resources", "Sources", "URLs", "Created At", "Updated At"}
	var stacks []api.Stack
	if options.stacks != nil {
		stacks = options.stacks.Stacks
	}
	if options.stack != nil {
		stacks = append(stacks, *options.stack)
	}

	var rows []map[string]interface{}
	for _, s := range stacks {
		var latestEvent api.StackEvent
		if len(s.Events) > 0 {
			sort.Slice(s.Events, func(i, j int) bool {
				return s.Events[i].UpdatedAt.Before(s.Events[j].UpdatedAt)
			})
			latestEvent = s.Events[len(s.Events)-1]
		}
		var desc string
		if latestEvent.Description != nil {
			desc = *latestEvent.Description
		}
		row := map[string]interface{}{
			"ID":            s.Id,
			"OrgID":         s.OrgID,
			"Active":        latestEvent.EventType != "uninstall",
			"Name":          latestEvent.Name,
			"Description":   desc,
			"Num Resources": len(latestEvent.Resources),
			"Sources":       latestEvent.Sources,
			"URLs":          latestEvent.Urls,
			"Created At":    s.CreatedAt,
			"Updated At":    latestEvent.UpdatedAt,
		}
		rows = append(rows, row)
	}
	return c.PrintTable(headers, rows...)
}

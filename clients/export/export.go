package export

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	clients.CLI
	api.TemplatesApi
	api.OrganizationsApi
}

type Params struct {
	OutParams
	StackId string

	IdsPerType   map[string][]string
	NamesPerType map[string][]string
}

func (c Client) Export(ctx context.Context, params *Params) error {
	var exportReq api.TemplateExport

	if params.StackId != "" {
		exportReq.StackID = &params.StackId
	}

	for typ, ids := range params.IdsPerType {
		for _, id := range ids {
			id := id
			exportReq.Resources = append(exportReq.Resources, api.TemplateExportResources{
				Kind: typ,
				Id:   &id,
			})
		}
	}
	for typ, names := range params.NamesPerType {
		for _, name := range names {
			name := name
			exportReq.Resources = append(exportReq.Resources, api.TemplateExportResources{
				Kind: typ,
				Name: &name,
			})
		}
	}

	tmpl, err := c.ExportTemplate(ctx).TemplateExport(exportReq).Execute()
	if err != nil {
		return fmt.Errorf("failed to export template: %w", err)
	}
	if err := params.OutParams.writeTemplate(tmpl); err != nil {
		return fmt.Errorf("failed to write exported template: %w", err)
	}
	return nil
}

type AllParams struct {
	OutParams

	OrgId   string
	OrgName string

	LabelFilters []string
	KindFilters  []string
}

func (c Client) ExportAll(ctx context.Context, params *AllParams) error {
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
			return fmt.Errorf("failed to look up ID for org %q: %w", orgName, err)
		}
		if len(res.GetOrgs()) == 0 {
			return fmt.Errorf("no org found with name %q", orgName)
		}
		orgId = res.GetOrgs()[0].GetId()
	}

	orgExport := api.TemplateExportOrgIDs{OrgID: &orgId}
	if len(params.LabelFilters) > 0 || len(params.KindFilters) > 0 {
		orgExport.ResourceFilters = &api.TemplateExportResourceFilters{}
		if len(params.LabelFilters) > 0 {
			orgExport.ResourceFilters.ByLabel = &params.LabelFilters
		}
		if len(params.KindFilters) > 0 {
			orgExport.ResourceFilters.ByResourceKind = &params.KindFilters
		}
	}

	exportReq := api.TemplateExport{OrgIDs: &[]api.TemplateExportOrgIDs{orgExport}}
	tmpl, err := c.ExportTemplate(ctx).TemplateExport(exportReq).Execute()
	if err != nil {
		return fmt.Errorf("failed to export template: %w", err)
	}
	if err := params.OutParams.writeTemplate(tmpl); err != nil {
		return fmt.Errorf("failed to write exported template: %w", err)
	}
	return nil
}

type StackParams struct {
	OutParams
	StackId 	string
}

func (c Client) ExportStack(ctx context.Context, params *StackParams) error {
	if params.StackId == "" {
		return clients.ErrMustSpecifyStack
	}

	stackId := params.StackId

	exportReq := api.TemplateExport{StackID: &stackId}
	tmpl, err := c.ExportTemplate(ctx).TemplateExport(exportReq).Execute()
	if err != nil {
		return fmt.Errorf("failed to export template: %w", err)
	}
	if err := params.OutParams.writeTemplate(tmpl); err != nil {
		return fmt.Errorf("failed to write exported template: %w", err)
	}

	return nil
}
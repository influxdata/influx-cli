package export

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/pkg/template"
)

type Client struct {
	clients.CLI
	api.TemplatesApi
	api.OrganizationsApi
}

type Params struct {
	template.OutParams
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
	if err := params.OutParams.WriteTemplate(tmpl); err != nil {
		return fmt.Errorf("failed to write exported template: %w", err)
	}
	return nil
}

type AllParams struct {
	template.OutParams

	clients.OrgParams

	LabelFilters []string
	KindFilters  []string
}

func (c Client) ExportAll(ctx context.Context, params *AllParams) error {
	orgId, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	if err != nil {
		return err
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
	if err := params.OutParams.WriteTemplate(tmpl); err != nil {
		return fmt.Errorf("failed to write exported template: %w", err)
	}
	return nil
}

type StackParams struct {
	template.OutParams
	StackId string
}

func (c Client) ExportStack(ctx context.Context, params *StackParams) error {
	if params.StackId == "" {
		return fmt.Errorf("no stack id provided")
	}

	exportReq := api.TemplateExport{StackID: &params.StackId}
	tmpl, err := c.ExportTemplate(ctx).TemplateExport(exportReq).Execute()

	if err != nil {
		return fmt.Errorf("failed to export stack %q: %w", params.StackId, err)
	}
	if err := params.OutParams.WriteTemplate(tmpl); err != nil {
		return fmt.Errorf("failed to write exported template: %w", err)
	}

	return nil
}

package template

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

type SummarizeParams struct {
	clients.OrgParams

	Sources []template.Source

	RenderTableColors  bool
	RenderTableBorders bool
}

func (c Client) Summarize(ctx context.Context, params *SummarizeParams) error {
	orgID, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	if err != nil {
		return err
	}

	templates, err := template.ReadSources(ctx, params.Sources)
	if err != nil {
		return err
	}

	// Execute a dry-run to make the server summarize the template(s).
	req := api.TemplateApply{
		DryRun:    true,
		OrgID:     orgID,
		Templates: templates,
	}
	res, err := c.ApplyTemplate(ctx).TemplateApply(req).Execute()
	if err != nil {
		return fmt.Errorf("failed to summarize template: %w", err)
	}

	if c.PrintAsJSON {
		return c.PrintJSON(res.Summary)
	}
	return template.PrintSummary(res.Summary, c.StdIO, params.RenderTableColors, params.RenderTableBorders)
}

type ValidateParams struct {
	clients.OrgParams

	Sources []template.Source
}

func (c Client) Validate(ctx context.Context, params *ValidateParams) error {
	orgID, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	if err != nil {
		return err
	}

	templates, err := template.ReadSources(ctx, params.Sources)
	if err != nil {
		return err
	}

	// Execute a dry-run to make the server summarize the template(s).
	req := api.TemplateApply{
		DryRun:    true,
		OrgID:     orgID,
		Templates: templates,
	}
	_, err = c.ApplyTemplate(ctx).TemplateApply(req).Execute()
	if err == nil {
		return nil
	}

	if apiErr, ok := err.(api.GenericOpenAPIError); ok {
		if summary, ok := apiErr.Model().(*api.TemplateSummary); ok {
			return fmt.Errorf("template failed validation:\n\t%s", summary)
		}
	}
	return fmt.Errorf("failed to validate template: %w", err)
}

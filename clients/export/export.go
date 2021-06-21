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

	BucketIds      []string
	BucketNames    []string
	CheckIds       []string
	CheckNames     []string
	DashboardIds   []string
	DashboardNames []string
	EndpointIds    []string
	EndpointNames  []string
	LabelIds       []string
	LabelNames     []string
	RuleIds        []string
	RuleNames      []string
	TaskIds        []string
	TaskNames      []string
	TelegrafIds    []string
	TelegrafNames  []string
	VariableIds    []string
	VariableNames  []string
}

func (c Client) Export(ctx context.Context, params *Params) error {
	var exportReq api.TemplateExport

	if params.StackId != "" {
		exportReq.StackID = &params.StackId
	}

	filters := []struct {
		kind  api.TemplateKind
		ids   []string
		names []string
	}{
		{
			kind:  api.TEMPLATEKIND_BUCKET,
			ids:   params.BucketIds,
			names: params.BucketNames,
		},
		{
			kind:  api.TEMPLATEKIND_CHECK,
			ids:   params.CheckIds,
			names: params.CheckNames,
		},
		{
			kind:  api.TEMPLATEKIND_DASHBOARD,
			ids:   params.DashboardIds,
			names: params.DashboardNames,
		},
		{
			kind:  api.TEMPLATEKIND_LABEL,
			ids:   params.LabelIds,
			names: params.LabelNames,
		},
		{
			kind:  api.TEMPLATEKIND_NOTIFICATION_ENDPOINT,
			ids:   params.EndpointIds,
			names: params.EndpointNames,
		},
		{
			kind:  api.TEMPLATEKIND_NOTIFICATION_RULE,
			ids:   params.RuleIds,
			names: params.RuleNames,
		},
		{
			kind:  api.TEMPLATEKIND_TASK,
			ids:   params.TaskIds,
			names: params.TaskNames,
		},
		{
			kind:  api.TEMPLATEKIND_TELEGRAF,
			ids:   params.TelegrafIds,
			names: params.TelegrafNames,
		},
		{
			kind:  api.TEMPLATEKIND_VARIABLE,
			ids:   params.VariableIds,
			names: params.VariableNames,
		},
	}

	for _, filter := range filters {
		for _, id := range filter.ids {
			id := id
			exportReq.Resources = append(exportReq.Resources, api.TemplateExportResources{
				Kind: filter.kind,
				Id:   &id,
			})
		}
		for _, name := range filter.names {
			name := name
			exportReq.Resources = append(exportReq.Resources, api.TemplateExportResources{
				Kind: filter.kind,
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
	KindFilters  []ResourceType
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
			kinds := make([]api.TemplateKind, len(params.KindFilters))
			for i, kf := range params.KindFilters {
				switch kf {
				case TypeBucket:
					kinds[i] = api.TEMPLATEKIND_BUCKET
				case TypeCheck:
					kinds[i] = api.TEMPLATEKIND_CHECK
				case TypeDashboard:
					kinds[i] = api.TEMPLATEKIND_DASHBOARD
				case TypeLabel:
					kinds[i] = api.TEMPLATEKIND_LABEL
				case TypeNotificationEndpoint:
					kinds[i] = api.TEMPLATEKIND_NOTIFICATION_ENDPOINT
				case TypeNotificationRule:
					kinds[i] = api.TEMPLATEKIND_NOTIFICATION_RULE
				case TypeTask:
					kinds[i] = api.TEMPLATEKIND_TASK
				case TypeTelegraf:
					kinds[i] = api.TEMPLATEKIND_TELEGRAF
				case TypeVariable:
					kinds[i] = api.TEMPLATEKIND_VARIABLE
				default:
					return fmt.Errorf("unsupported resourceKind filter %q", kf)
				}
			}
			orgExport.ResourceFilters.ByResourceKind = &kinds
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

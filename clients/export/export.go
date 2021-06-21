package export

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"gopkg.in/yaml.v3"
)

type OutEncoding int

const (
	YamlEncoding OutEncoding = iota
	JsonEncoding
)

type Client struct {
	clients.CLI
	api.TemplatesApi
}

type Params struct {
	Out io.Writer
	OutEncoding
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
	if err := writeTemplate(params.Out, params.OutEncoding, tmpl); err != nil {
		return fmt.Errorf("failed to write exported template: %w", err)
	}
	return nil
}

func writeTemplate(out io.Writer, encoding OutEncoding, template []api.TemplateEntry) error {
	switch encoding {
	case JsonEncoding:
		enc := json.NewEncoder(out)
		enc.SetIndent("", "\t")
		return enc.Encode(template)
	case YamlEncoding:
		enc := yaml.NewEncoder(out)
		for _, entry := range template {
			if err := enc.Encode(entry); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("encoding %q is not recognized", encoding)
	}
	return nil
}

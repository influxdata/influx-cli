package apply

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/influxdata/influx-cli/v2/pkg/template"
)

type Client struct {
	clients.CLI
	api.TemplatesApi
	api.OrganizationsApi
}

type Params struct {
	clients.OrgParams

	StackId   string
	Sources   []template.Source
	Recursive bool

	Secrets map[string]string
	EnvVars map[string]string

	Filters []ResourceFilter

	Force              bool
	Quiet              bool
	RenderTableColors  bool
	RenderTableBorders bool
}

type ResourceFilter struct {
	Kind string
	Name *string
}

func (c Client) Apply(ctx context.Context, params *Params) error {
	orgID, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	if err != nil {
		return err
	}

	templates, err := template.ReadSources(ctx, params.Sources)
	if err != nil {
		return err
	}

	req := api.TemplateApply{
		DryRun:    true,
		OrgID:     orgID,
		Templates: templates,
		EnvRefs:   params.EnvVars,
		Secrets:   params.Secrets,
		Actions:   make([]api.TemplateApplyAction, len(params.Filters)),
	}
	if params.StackId != "" {
		req.StackID = &params.StackId
	}
	for i, f := range params.Filters {
		req.Actions[i].Action = api.TEMPLATEAPPLYACTIONKIND_SKIP_KIND
		if f.Name != nil {
			req.Actions[i].Action = api.TEMPLATEAPPLYACTIONKIND_SKIP_RESOURCE
		}
		req.Actions[i].Properties = api.TemplateApplyActionProperties{
			Kind:                 f.Kind,
			ResourceTemplateName: f.Name,
		}
	}

	// Initial dry-run to make the server summarize the template & report any missing env/secret values.
	res, err := c.ApplyTemplate(ctx).TemplateApply(req).Execute()
	if err != nil {
		return fmt.Errorf("failed to check template impact: %w", err)
	}

	if c.StdIO.IsInteractive() && (len(res.Summary.MissingEnvRefs) > 0 || len(res.Summary.MissingSecrets) > 0) {
		for _, e := range res.Summary.MissingEnvRefs {
			val, err := c.StdIO.GetStringInput(fmt.Sprintf("Please provide environment reference value for key %s", e), "")
			if err != nil {
				return err
			}
			req.EnvRefs[e] = val
		}
		for _, s := range res.Summary.MissingSecrets {
			val, err := c.StdIO.GetSecret(fmt.Sprintf("Please provide secret value for key %s (optional, press enter to skip)", s), 0)
			if err != nil {
				return err
			}
			if val != "" {
				req.Secrets[s] = val
			}
		}

		// 2nd dry-run to see the diff after resolving all env/secret values.
		res, err = c.ApplyTemplate(ctx).TemplateApply(req).Execute()
		if err != nil {
			return fmt.Errorf("failed to check template impact: %w", err)
		}
	}

	if !params.Quiet {
		if err := c.printDiff(res.Diff, params); err != nil {
			return err
		}
	}

	if !params.Force {
		if confirmed := c.StdIO.GetConfirm("Confirm application of the above resources"); !confirmed {
			return errors.New("aborted application of template")
		}
	}

	// Flip the dry-run flag and apply the template.
	req.DryRun = false
	res, err = c.ApplyTemplate(ctx).TemplateApply(req).Execute()
	if err != nil {
		return fmt.Errorf("failed to apply template: %w", err)
	}
	params.StackId = res.StackID

	if !params.Quiet {
		if err := c.printSummary(res.Summary, params); err != nil {
			return err
		}
	}

	return nil
}

func (c Client) printDiff(diff api.TemplateSummaryDiff, params *Params) error {
	if c.PrintAsJSON {
		return c.PrintJSON(diff)
	}

	newDiffPrinter := func(title string, headers []string) *template.DiffPrinter {
		return template.NewDiffPrinter(c.StdIO, params.RenderTableColors, params.RenderTableBorders).
			Title(title).
			SetHeaders(append([]string{"Metadata Name", "ID", "Resource Name"}, headers...)...)
	}

	if labels := diff.Labels; len(labels) > 0 {
		printer := newDiffPrinter("Labels", []string{"Color", "Description"})
		buildRow := func(metaName string, id string, lf api.TemplateSummaryDiffLabelFields) []string {
			var desc string
			if lf.Description != nil {
				desc = *lf.Description
			}
			return []string{metaName, id, lf.Name, lf.Color, desc}
		}
		for _, l := range labels {
			var oldRow, newRow []string
			hexId := influxid.Encode(l.Id)
			if l.Old != nil {
				oldRow = buildRow(l.TemplateMetaName, hexId, *l.Old)
			}
			if l.New != nil {
				newRow = buildRow(l.TemplateMetaName, hexId, *l.New)
			}
			printer.AppendDiff(oldRow, newRow, false)
		}
		printer.Render()
		_, _ = c.StdIO.Write([]byte("\n"))
	}

	if bkts := diff.Buckets; len(bkts) > 0 {
		printer := newDiffPrinter("Buckets", []string{"Retention Period", "Description", "Schema Type", "Num Measurements"})
		buildRow := func(metaName string, id string, bf api.TemplateSummaryDiffBucketFields) []string {
			var desc string
			if bf.Description != nil {
				desc = *bf.Description
			}
			var retention time.Duration
			if len(bf.RetentionRules) > 0 {
				retention = time.Duration(bf.RetentionRules[0].EverySeconds) * time.Second
			}
			schemaType := api.SCHEMATYPE_IMPLICIT
			if bf.SchemaType != nil {
				schemaType = *bf.SchemaType
			}
			return []string{metaName, id, bf.Name, retention.String(), desc, schemaType.String(), strconv.Itoa(len(bf.MeasurementSchemas))}
		}
		for _, b := range bkts {
			var oldRow, newRow []string
			hexId := influxid.Encode(b.Id)
			if b.Old != nil {
				oldRow = buildRow(b.TemplateMetaName, hexId, *b.Old)
			}
			if b.New != nil {
				newRow = buildRow(b.TemplateMetaName, hexId, *b.New)
			}
			printer.AppendDiff(oldRow, newRow, false)
		}
		printer.Render()
		_, _ = c.StdIO.Write([]byte("\n"))
	}

	if checks := diff.Checks; len(checks) > 0 {
		printer := newDiffPrinter("Checks", []string{"Description"})
		buildRow := func(metaName string, id string, cf api.TemplateSummaryDiffCheckFields) []string {
			var desc string
			if cf.Description != nil {
				desc = *cf.Description
			}
			return []string{metaName, id, cf.Name, desc}
		}
		for _, c := range checks {
			var oldRow, newRow []string
			hexId := influxid.Encode(c.Id)
			if c.Old != nil {
				oldRow = buildRow(c.TemplateMetaName, hexId, *c.Old)
			}
			if c.New != nil {
				newRow = buildRow(c.TemplateMetaName, hexId, *c.New)
			}
			printer.AppendDiff(oldRow, newRow, true)
		}
		printer.Render()
		_, _ = c.StdIO.Write([]byte("\n"))
	}

	if dashboards := diff.Dashboards; len(dashboards) > 0 {
		printer := newDiffPrinter("Dashboards", []string{"Description", "Num Charts"})
		buildRow := func(metaName string, id string, df api.TemplateSummaryDiffDashboardFields) []string {
			var desc string
			if df.Description != nil {
				desc = *df.Description
			}
			return []string{metaName, id, df.Name, desc, strconv.Itoa(len(df.Charts))}
		}
		for _, d := range dashboards {
			var oldRow, newRow []string
			hexId := influxid.Encode(d.Id)
			if d.Old != nil {
				oldRow = buildRow(d.TemplateMetaName, hexId, *d.Old)
			}
			if d.New != nil {
				newRow = buildRow(d.TemplateMetaName, hexId, *d.New)
			}
			printer.AppendDiff(oldRow, newRow, true)
		}
		printer.Render()
		_, _ = c.StdIO.Write([]byte("\n"))
	}

	if endpoints := diff.NotificationEndpoints; len(endpoints) > 0 {
		printer := newDiffPrinter("Notification Endpoints", nil)
		buildRow := func(metaName string, id string, nef api.TemplateSummaryDiffNotificationEndpointFields) []string {
			return []string{metaName, id, nef.Name}
		}
		for _, e := range endpoints {
			var oldRow, newRow []string
			hexId := influxid.Encode(e.Id)
			if e.Old != nil {
				oldRow = buildRow(e.TemplateMetaName, hexId, *e.Old)
			}
			if e.New != nil {
				newRow = buildRow(e.TemplateMetaName, hexId, *e.New)
			}
			printer.AppendDiff(oldRow, newRow, true)
		}
		printer.Render()
		_, _ = c.StdIO.Write([]byte("\n"))
	}

	if rules := diff.NotificationRules; len(rules) > 0 {
		printer := newDiffPrinter("Notification Rules", []string{"Description", "Every", "Offset", "Endpoint Name", "Endpoint ID", "Endpoint Type"})
		buildRow := func(metaName string, id string, nrf api.TemplateSummaryDiffNotificationRuleFields) []string {
			var desc string
			if nrf.Description != nil {
				desc = *nrf.Description
			}
			eid := influxid.Encode(nrf.EndpointID)
			return []string{metaName, id, nrf.Name, desc, nrf.Every, nrf.Offset, nrf.EndpointName, eid, nrf.EndpointType}
		}
		for _, r := range rules {
			var oldRow, newRow []string
			hexId := influxid.Encode(r.Id)
			if r.Old != nil {
				oldRow = buildRow(r.TemplateMetaName, hexId, *r.Old)
			}
			if r.New != nil {
				newRow = buildRow(r.TemplateMetaName, hexId, *r.New)
			}
			printer.AppendDiff(oldRow, newRow, true)
		}
		printer.Render()
		_, _ = c.StdIO.Write([]byte("\n"))
	}

	if teles := diff.TelegrafConfigs; len(teles) > 0 {
		printer := newDiffPrinter("Telegraf Configurations", []string{"Description"})
		buildRow := func(metaName string, id string, tc api.TemplateSummaryTelegrafConfig) []string {
			var desc string
			if tc.Description != nil {
				desc = *tc.Description
			}
			return []string{metaName, id, tc.Name, desc}
		}
		for _, t := range teles {
			var oldRow, newRow []string
			hexId := influxid.Encode(t.Id)
			if t.Old != nil {
				oldRow = buildRow(t.TemplateMetaName, hexId, *t.Old)
			}
			if t.New != nil {
				newRow = buildRow(t.TemplateMetaName, hexId, *t.New)
			}
			printer.AppendDiff(oldRow, newRow, true)
		}
		printer.Render()
		_, _ = c.StdIO.Write([]byte("\n"))
	}

	if tasks := diff.Tasks; len(tasks) > 0 {
		printer := newDiffPrinter("Tasks", []string{"Description", "Cycle"})
		buildRow := func(metaName string, id string, tf api.TemplateSummaryDiffTaskFields) []string {
			var desc string
			if tf.Description != nil {
				desc = *tf.Description
			}
			var timing string
			if tf.Cron != nil {
				timing = *tf.Cron
			} else {
				offset := time.Duration(0).String()
				if tf.Offset != nil {
					offset = *tf.Offset
				}
				// If `cron` isn't set, `every` must be set
				timing = fmt.Sprintf("every: %s offset: %s", *tf.Every, offset)
			}
			return []string{metaName, id, tf.Name, desc, timing}
		}
		for _, t := range tasks {
			var oldRow, newRow []string
			hexId := influxid.Encode(t.Id)
			if t.Old != nil {
				oldRow = buildRow(t.TemplateMetaName, hexId, *t.Old)
			}
			if t.New != nil {
				newRow = buildRow(t.TemplateMetaName, hexId, *t.New)
			}
			printer.AppendDiff(oldRow, newRow, true)
		}
		printer.Render()
		_, _ = c.StdIO.Write([]byte("\n"))
	}

	if vars := diff.Variables; len(vars) > 0 {
		printer := newDiffPrinter("Variables", []string{"Description", "Arg Type", "Arg Values"})
		buildRow := func(metaName string, id string, vf api.TemplateSummaryDiffVariableFields) []string {
			var desc, argType string
			if vf.Description != nil {
				desc = *vf.Description
			}
			if vf.Args != nil {
				argType = vf.Args.Type
			}
			return []string{metaName, id, vf.Name, desc, argType, vf.Args.Render()}
		}
		for _, v := range vars {
			var oldRow, newRow []string
			hexId := influxid.Encode(v.Id)
			if v.Old != nil {
				oldRow = buildRow(v.TemplateMetaName, hexId, *v.Old)
			}
			if v.New != nil {
				newRow = buildRow(v.TemplateMetaName, hexId, *v.New)
			}
			printer.AppendDiff(oldRow, newRow, false)
		}
		printer.Render()
		_, _ = c.StdIO.Write([]byte("\n"))
	}

	if mappings := diff.LabelMappings; len(mappings) > 0 {
		printer := template.NewDiffPrinter(c.StdIO, params.RenderTableColors, params.RenderTableBorders).
			Title("Label Associations").
			SetHeaders("Resource Type", "Resource Meta Name", "Resource Name", "Resource ID", "Label Package Name", "Label Name", "Label ID")

		for _, m := range mappings {
			resId := influxid.Encode(m.ResourceID)
			labelId := influxid.Encode(m.LabelID)
			row := []string{m.ResourceType, m.ResourceName, resId, m.LabelTemplateMetaName, m.LabelName, labelId}
			switch m.StateStatus {
			case "new":
				printer.AppendDiff(nil, row, false)
			case "remove":
				printer.AppendDiff(row, nil, false)
			default:
				printer.AppendDiff(row, row, false)
			}
		}
		printer.Render()
	}

	return nil
}

func (c Client) printSummary(summary api.TemplateSummaryResources, params *Params) error {
	if c.PrintAsJSON {
		return c.PrintJSON(struct {
			StackID string `json:"stackID"`
			Summary api.TemplateSummaryResources
		}{
			StackID: params.StackId,
			Summary: summary,
		})
	}

	if err := template.PrintSummary(summary, c.StdIO, params.RenderTableColors, params.RenderTableBorders); err != nil {
		return err
	}
	if params.StackId != "" {
		_, _ = c.StdIO.Write([]byte(fmt.Sprintf("Stack ID: %s\n", params.StackId)))
	}

	return nil
}

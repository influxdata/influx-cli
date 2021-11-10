package template

import (
	"fmt"
	"io"
	"time"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
)

// PrintSummary renders high-level info about a template as a table for display on the console.
//
// NOTE: The implementation here is very "static" in that it's hard-coded to look for specific
// resource-kinds and fields within those kinds. If the API changes to add more kinds / more fields,
// this function won't automatically pick them up & print them. It'd be nice to rework this to be
// less opinionated / more resilient to extension in the future...
func PrintSummary(summary api.TemplateSummaryResources, out io.Writer, useColor bool, useBorders bool) error {
	newPrinter := func(title string, headers []string) *TablePrinter {
		return NewTablePrinter(out, useColor, useBorders).
			Title(title).
			SetHeaders(append([]string{"Package Name", "ID", "Resource Name"}, headers...)...)
	}

	if labels := summary.Labels; len(labels) > 0 {
		printer := newPrinter("LABELS", []string{"Description", "Color"})
		for _, l := range labels {
			id := influxid.Encode(l.Id)
			var desc string
			if l.Properties.Description != nil {
				desc = *l.Properties.Description
			}
			printer.Append([]string{l.TemplateMetaName, id, l.Name, desc, l.Properties.Color})
		}
		printer.Render()
		_, _ = out.Write([]byte("\n"))
	}

	if buckets := summary.Buckets; len(buckets) > 0 {
		printer := newPrinter("BUCKETS", []string{"Retention", "Description", "Schema Type"})
		for _, b := range buckets {
			id := influxid.Encode(b.Id)
			var desc string
			if b.Description != nil {
				desc = *b.Description
			}
			rp := "inf"
			if b.RetentionPeriod != 0 {
				rp = time.Duration(b.RetentionPeriod).String()
			}
			schemaType := api.SCHEMATYPE_IMPLICIT
			if b.SchemaType != nil {
				schemaType = *b.SchemaType
			}
			printer.Append([]string{b.TemplateMetaName, id, b.Name, rp, desc, schemaType.String()})
		}
		printer.Render()
		_, _ = out.Write([]byte("\n"))
	}

	if checks := summary.Checks; len(checks) > 0 {
		printer := newPrinter("CHECKS", []string{"Description"})
		for _, c := range checks {
			id := influxid.Encode(c.Id)
			var desc string
			if c.Description != nil {
				desc = *c.Description
			}
			printer.Append([]string{c.TemplateMetaName, id, c.Name, desc})
		}
		printer.Render()
		_, _ = out.Write([]byte("\n"))
	}

	if dashboards := summary.Dashboards; len(dashboards) > 0 {
		printer := newPrinter("DASHBOARDS", []string{"Description"})
		for _, d := range dashboards {
			id := influxid.Encode(d.Id)
			var desc string
			if d.Description != nil {
				desc = *d.Description
			}
			printer.Append([]string{d.TemplateMetaName, id, d.Name, desc})
		}
		printer.Render()
		_, _ = out.Write([]byte("\n"))
	}

	if endpoints := summary.NotificationEndpoints; len(endpoints) > 0 {
		printer := newPrinter("NOTIFICATION ENDPOINTS", []string{"Description", "Status"})
		for _, e := range endpoints {
			id := influxid.Encode(e.Id)
			var desc string
			if e.Description != nil {
				desc = *e.Description
			}
			printer.Append([]string{e.TemplateMetaName, id, e.Name, desc, e.Status})
		}
		printer.Render()
		_, _ = out.Write([]byte("\n"))
	}

	if rules := summary.NotificationRules; len(rules) > 0 {
		printer := newPrinter("NOTIFICATION RULES", []string{"Description", "Every", "Offset", "Endpoint Name", "Endpoint ID", "Endpoint Type"})
		for _, r := range rules {
			id := influxid.Encode(r.Id)
			eid := influxid.Encode(r.EndpointID)
			var desc string
			if r.Description != nil {
				desc = *r.Description
			}
			printer.Append([]string{r.TemplateMetaName, id, r.Name, desc, r.Every, r.Offset, r.EndpointTemplateMetaName, eid, r.EndpointType})
		}
		printer.Render()
		_, _ = out.Write([]byte("\n"))
	}

	if tasks := summary.Tasks; len(tasks) > 0 {
		printer := newPrinter("TASKS", []string{"Description", "Cycle"})
		for _, t := range tasks {
			id := influxid.Encode(t.Id)
			var desc string
			if t.Description != nil {
				desc = *t.Description
			}
			var timing string
			if t.Cron != nil {
				timing = *t.Cron
			} else {
				offset := time.Duration(0).String()
				if t.Offset != nil {
					offset = *t.Offset
				}
				// If `cron` isn't set, `every` must be set
				timing = fmt.Sprintf("every: %s offset: %s", *t.Every, offset)
			}
			printer.Append([]string{t.TemplateMetaName, id, t.Name, desc, timing})
		}
		printer.Render()
		_, _ = out.Write([]byte("\n"))
	}

	if teles := summary.TelegrafConfigs; len(teles) > 0 {
		printer := newPrinter("TELEGRAF CONFIGS", []string{"Description"})
		for _, t := range teles {
			var desc string
			if t.TelegrafConfig.Description != nil {
				desc = *t.TelegrafConfig.Description
			}
			printer.Append([]string{t.TemplateMetaName, t.TelegrafConfig.Id, t.TelegrafConfig.Name, desc})
		}
		printer.Render()
		_, _ = out.Write([]byte("\n"))
	}

	if vars := summary.Variables; len(vars) > 0 {
		printer := newPrinter("VARIABLES", []string{"Description", "Arg Type", "Arg Values"})
		for _, v := range vars {
			id := influxid.Encode(v.Id)
			var desc string
			if v.Description != nil {
				desc = *v.Description
			}
			var argType string
			if v.Arguments != nil {
				argType = v.Arguments.Type
			}
			printer.Append([]string{v.TemplateMetaName, id, v.Name, desc, argType, v.Arguments.Render()})
		}
		printer.Render()
		_, _ = out.Write([]byte("\n"))
	}

	if mappings := summary.LabelMappings; len(mappings) > 0 {
		printer := NewTablePrinter(out, useColor, useBorders).
			Title("LABEL ASSOCIATIONS").
			SetHeaders("Resource Type", "Resource Name", "Resource ID", "Label Name", "Label ID")
		for _, m := range mappings {
			rid := influxid.Encode(m.ResourceID)
			lid := influxid.Encode(m.LabelID)
			printer.Append([]string{m.ResourceType, m.ResourceName, rid, m.LabelName, lid})
		}
		printer.Render()
		_, _ = out.Write([]byte("\n"))
	}

	if secrets := summary.MissingSecrets; len(secrets) > 0 {
		printer := NewTablePrinter(out, useColor, useBorders).
			Title("MISSING SECRETS").
			SetHeaders("Secret Key")
		for _, sk := range secrets {
			printer.Append([]string{sk})
		}
		printer.Render()
		_, _ = out.Write([]byte("\n"))
	}

	return nil
}

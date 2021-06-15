package dashboards

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	clients.CLI
	api.DashboardsApi
}

type Params struct {
	clients.OrgParams
	Ids []string
}

func (c Client) List(ctx context.Context, params *Params) error {
	if !params.OrgID.Valid() && params.OrgName == "" && c.ActiveConfig.Org == "" && len(params.Ids) == 0 {
		return fmt.Errorf("at least one of org, org-id, or id must be provided")
	}

	const limit = 100
	req := c.GetDashboards(ctx)
	req = req.Limit(limit)
	if params.OrgID.Valid() {
		req = req.OrgID(params.OrgID.String())
	}
	if params.OrgName != "" {
		req = req.Org(params.OrgName)
	}
	if !params.OrgID.Valid() && params.OrgName == "" {
		req = req.Org(c.ActiveConfig.Org)
	}
	dashboards, err := req.Id(params.Ids).Execute()
	if err != nil {
		return fmt.Errorf("failed to find dashboards: %w", err)
	}

	return c.printDashboards(dashboards)
}

func (c Client) printDashboards(dashboards api.Dashboards) error {
	if c.PrintAsJSON {
		return c.PrintJSON(dashboards)
	}

	headers := []string{"ID", "OrgID", "Name", "Description", "Num Cells"}
	var rows []map[string]interface{}
	for _, u := range dashboards.GetDashboards() {
		row := map[string]interface{}{
			"ID":          u.GetId(),
			"OrgID":       u.GetOrgID(),
			"Name":        u.GetName(),
			"Description": u.GetDescription(),
			"Num Cells":   len(u.GetCells()),
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

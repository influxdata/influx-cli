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
	api.OrganizationsApi
}

type Params struct {
	clients.OrgParams
	Ids []string
}

func (c Client) List(ctx context.Context, params *Params) error {
	if !params.OrgID.Valid() && params.OrgName == "" && len(params.Ids) == 0 {
		return fmt.Errorf("at least one of org, org-id, or id must be provided")
	}

	var req api.ApiGetOrgsRequest
	if params.OrgID.Valid() {
		req = req.OrgID(params.OrgID.String())
	}
	if params.OrgName != "" {
		req = req.Org(params.OrgName)
	}
	if !params.OrgID.Valid() && params.OrgName == "" {
		req = req.Org(c.ActiveConfig.Org)
	}
	orgReq, err := req.Execute()
	if err != nil {
		return fmt.Errorf("failed to lookup org with ID %q or name %q: %w", params.OrgID, params.OrgName, err)
	}
	orgIDs := orgReq.GetOrgs()
	if len(orgIDs) == 0 {
		return fmt.Errorf("no organization found with name %q: %w", params.OrgName, err)
	}

	orgID := orgIDs[0].GetId()
	const limit = 100
	dashReq := c.GetDashboards(ctx)
	dashReq = dashReq.Limit(limit)
	dashReq = dashReq.OrgID(orgID).Id(params.Ids)
	dashboards, err := dashReq.Execute()
	if err != nil {
		return fmt.Errorf("failed to find dashboards with OrgID %q and IDs %q: %w", orgID, params.Ids, err)
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

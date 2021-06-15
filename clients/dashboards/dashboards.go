package dashboards

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
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
	orgID, err := c.getOrgID(ctx, params.OrgID, params.OrgName)
	if err != nil {
		return err
	}
	if orgID == "" && len(params.Ids) == 0 {
		return fmt.Errorf("at least one of org, org-id, or id must be provided")
	}

	const limit = 100
	req := c.GetDashboards(ctx)
	req = req.Limit(limit)
	req = req.OrgID(orgID).Id(params.Ids)
	dashboards, err := req.Execute()
	if err != nil {
		return fmt.Errorf("failed to find dashboards with OrgID %q and IDs %q: %w", orgID, params.Ids, err)
	}

	return c.printDashboards(dashboards)
}

func (c Client) getOrgID(ctx context.Context, orgID influxid.ID, orgName string) (string, error) {
	if orgID.Valid() {
		return orgID.String(), nil
	} else {
		if orgName == "" {
			orgName = c.ActiveConfig.Org
		}
		resp, err := c.GetOrgs(ctx).Org(orgName).Execute()
		if err != nil {
			return "", fmt.Errorf("failed to lookup ID of org %q: %w", orgName, err)
		}
		orgs := resp.GetOrgs()
		if len(orgs) == 0 {
			return "", fmt.Errorf("no organization found with name %q", orgName)
		}
		return orgs[0].GetId(), nil
	}
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

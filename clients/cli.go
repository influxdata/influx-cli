package clients

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/config"
	"github.com/influxdata/influx-cli/v2/internal/tabwriter"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/influxdata/influx-cli/v2/pkg/stdio"
)

// CLI is a container for common functionality used to execute commands.
type CLI struct {
	StdIO stdio.StdIO

	HideTableHeaders bool
	PrintAsJSON      bool

	ActiveConfig  config.Config
	ConfigService config.Service
}

func (c *CLI) PrintJSON(v interface{}) error {
	enc := json.NewEncoder(c.StdIO)
	enc.SetIndent("", "\t")
	return enc.Encode(v)
}

func (c *CLI) PrintTable(headers []string, rows ...map[string]interface{}) error {
	w := tabwriter.NewTabWriter(c.StdIO, c.HideTableHeaders)
	defer w.Flush()
	if err := w.WriteHeaders(headers...); err != nil {
		return err
	}
	for _, r := range rows {
		if err := w.Write(r); err != nil {
			return err
		}
	}
	return nil
}

func (c *CLI) GetOrgId(ctx context.Context, paramOrgId, paramOrgName string, orgApi api.OrganizationsApi) (string, error) {
	if paramOrgId != "" {
		return paramOrgId, nil
	}
	orgName := paramOrgName
	if orgName == "" {
		orgName = c.ActiveConfig.Org
	}
	if orgName == "" {
		return "", ErrMustSpecifyOrg
	}
	res, err := orgApi.GetOrgs(ctx).Org(orgName).Execute()
	if err != nil {
		return "", fmt.Errorf("failed to lookup org with name %q: %w", orgName, err)
	}
	if len(res.GetOrgs()) == 0 {
		return "", fmt.Errorf("no organization with name %q", orgName)
	}
	return res.GetOrgs()[0].GetId(), nil
}

func (c *CLI) GetOrgIdI(ctx context.Context, paramOrgId influxid.ID, paramOrgName string, orgApi api.OrganizationsApi) (string, error) {
	orgId := ""
	if paramOrgId.Valid() {
		orgId = paramOrgId.String()
	}
	return c.GetOrgId(ctx, orgId, paramOrgName, orgApi)
}

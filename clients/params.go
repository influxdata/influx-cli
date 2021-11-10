package clients

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/config"
)

type OrgParams struct {
	OrgID   string
	OrgName string
}

func (p OrgParams) GetOrgID(ctx context.Context, activeConfig config.Config, orgApi api.OrganizationsApi) (string, error) {
	if p.OrgID != "" {
		return p.OrgID, nil
	}
	orgName := p.OrgName
	if orgName == "" {
		orgName = activeConfig.Org
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

type BucketParams struct {
	BucketID   string
	BucketName string
}

type OrgBucketParams struct {
	OrgParams
	BucketParams
}

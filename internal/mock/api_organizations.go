package mock

import (
	"context"

	"github.com/influxdata/influx-cli/v2/internal/api"
)

var _ api.OrganizationsApi = (*OrganizationsApi)(nil)

type OrganizationsApi struct {
	GetOrgsExecuteFn  func(api.ApiGetOrgsRequest) (api.Organizations, error)
	PostOrgsExecuteFn func(api.ApiPostOrgsRequest) (api.Organization, error)
}

func (o *OrganizationsApi) GetOrgs(context.Context) api.ApiGetOrgsRequest {
	return api.ApiGetOrgsRequest{ApiService: o}
}

func (o *OrganizationsApi) GetOrgsExecute(r api.ApiGetOrgsRequest) (api.Organizations, error) {
	return o.GetOrgsExecuteFn(r)
}

func (o *OrganizationsApi) PostOrgs(context.Context) api.ApiPostOrgsRequest {
	return api.ApiPostOrgsRequest{ApiService: o}
}

func (o *OrganizationsApi) PostOrgsExecute(r api.ApiPostOrgsRequest) (api.Organization, error) {
	return o.PostOrgsExecuteFn(r)
}

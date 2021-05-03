package mock

import (
	"context"
	"net/http"

	"github.com/influxdata/influx-cli/v2/internal/api"
)

var _ api.OrganizationsApi = (*OrganizationsApi)(nil)

type OrganizationsApi struct {
	GetOrgsExecuteFn  func(api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error)
	PostOrgsExecuteFn func(api.ApiPostOrgsRequest) (api.Organization, *http.Response, error)
}

func (o *OrganizationsApi) GetOrgs(context.Context) api.ApiGetOrgsRequest {
	return api.ApiGetOrgsRequest{ApiService: o}
}

func (o *OrganizationsApi) GetOrgsExecute(r api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
	return o.GetOrgsExecuteFn(r)
}

func (o *OrganizationsApi) PostOrgs(context.Context) api.ApiPostOrgsRequest {
	return api.ApiPostOrgsRequest{ApiService: o}
}

func (o *OrganizationsApi) PostOrgsExecute(r api.ApiPostOrgsRequest) (api.Organization, *http.Response, error) {
	return o.PostOrgsExecuteFn(r)
}

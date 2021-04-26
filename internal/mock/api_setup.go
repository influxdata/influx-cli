package mock

import (
	"context"
	"net/http"

	"github.com/influxdata/influx-cli/v2/internal/api"
)

var _ api.SetupApi = (*SetupApi)(nil)

type SetupApi struct {
	GetSetupExecuteFn  func(api.ApiGetSetupRequest) (api.InlineResponse200, *http.Response, error)
	PostSetupExecuteFn func(api.ApiPostSetupRequest) (api.OnboardingResponse, *http.Response, error)
}

func (s *SetupApi) GetSetup(context.Context) api.ApiGetSetupRequest {
	return api.ApiGetSetupRequest{
		ApiService: s,
	}
}
func (s *SetupApi) GetSetupExecute(req api.ApiGetSetupRequest) (api.InlineResponse200, *http.Response, error) {
	return s.GetSetupExecuteFn(req)
}
func (s *SetupApi) PostSetup(context.Context) api.ApiPostSetupRequest {
	return api.ApiPostSetupRequest{
		ApiService: s,
	}
}
func (s *SetupApi) PostSetupExecute(req api.ApiPostSetupRequest) (api.OnboardingResponse, *http.Response, error) {
	return s.PostSetupExecuteFn(req)
}

package mock

import (
	"context"

	"github.com/influxdata/influx-cli/v2/internal/api"
)

var _ api.SetupApi = (*SetupApi)(nil)

type SetupApi struct {
	GetSetupExecuteFn  func(api.ApiGetSetupRequest) (api.InlineResponse200, error)
	PostSetupExecuteFn func(api.ApiPostSetupRequest) (api.OnboardingResponse, error)
}

func (s *SetupApi) GetSetup(context.Context) api.ApiGetSetupRequest {
	return api.ApiGetSetupRequest{
		ApiService: s,
	}
}
func (s *SetupApi) GetSetupExecute(req api.ApiGetSetupRequest) (api.InlineResponse200, error) {
	return s.GetSetupExecuteFn(req)
}
func (s *SetupApi) PostSetup(context.Context) api.ApiPostSetupRequest {
	return api.ApiPostSetupRequest{
		ApiService: s,
	}
}
func (s *SetupApi) PostSetupExecute(req api.ApiPostSetupRequest) (api.OnboardingResponse, error) {
	return s.PostSetupExecuteFn(req)
}

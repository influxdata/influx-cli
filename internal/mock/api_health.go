package mock

import (
	"context"
	"net/http"

	"github.com/influxdata/influx-cli/v2/internal/api"
)

var _ api.HealthApi = (*HealthApi)(nil)

type HealthApi struct {
	GetHealthExecuteFn func(api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error)
}

func (p *HealthApi) GetHealth(context.Context) api.ApiGetHealthRequest {
	return api.ApiGetHealthRequest{
		ApiService: p,
	}
}

func (p *HealthApi) GetHealthExecute(req api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error) {
	return p.GetHealthExecuteFn(req)
}

package mock

import (
	"context"
	"net/http"

	"github.com/influxdata/influx-cli/v2/internal/api"
)

var _ api.WriteApi = (*WriteApi)(nil)

type WriteApi struct {
	PostWriteExecuteFn func(api.ApiPostWriteRequest) (*http.Response, error)
}

func (w *WriteApi) PostWrite(context.Context) api.ApiPostWriteRequest {
	return api.ApiPostWriteRequest{
		ApiService: w,
	}
}
func (w *WriteApi) PostWriteExecute(req api.ApiPostWriteRequest) (*http.Response, error) {
	return w.PostWriteExecuteFn(req)
}

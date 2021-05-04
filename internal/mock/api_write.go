package mock

import (
	"context"

	"github.com/influxdata/influx-cli/v2/internal/api"
)

var _ api.WriteApi = (*WriteApi)(nil)

type WriteApi struct {
	PostWriteExecuteFn func(api.ApiPostWriteRequest) error
}

func (w *WriteApi) PostWrite(context.Context) api.ApiPostWriteRequest {
	return api.ApiPostWriteRequest{
		ApiService: w,
	}
}
func (w *WriteApi) PostWriteExecute(req api.ApiPostWriteRequest) error {
	return w.PostWriteExecuteFn(req)
}

package tracing

import (
	"context"

	"github.com/influxdata/influx-cli/v2/api"
)

// Context is a convenience wrapper for tracking a global tracing
// ID through CLI commands.
type Context interface {
	context.Context
	TraceId() *api.TraceSpan
}

// WrapContext combines a standard context with an optional trace ID.
func WrapContext(ctx context.Context, traceId string) Context {
	tc := tracingContext{Context: ctx}
	if traceId != "" {
		tspan := api.TraceSpan(traceId)
		tc.traceId = &tspan
	}
	return &tc
}

type tracingContext struct {
	context.Context
	traceId *api.TraceSpan
}

func (tc *tracingContext) TraceId() *api.TraceSpan {
	return tc.traceId
}

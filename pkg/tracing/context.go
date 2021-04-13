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

// WithTraceId combines a context with an optional trace ID.
func WithTraceId(ctx context.Context, traceId *api.TraceSpan) Context {
	return &traceCtx{ctx, traceId}
}

type traceCtx struct {
	context.Context
	traceId *api.TraceSpan
}

func (tc *traceCtx) TraceId() *api.TraceSpan {
	return tc.traceId
}

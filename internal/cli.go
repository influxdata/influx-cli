package internal

import (
	"io"

	"github.com/influxdata/influx-cli/v2/internal/api"
)

type CLI struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer

	TraceId *api.TraceSpan
}

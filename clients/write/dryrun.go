package write

import (
	"context"
	"io"

	"github.com/influxdata/influx-cli/v2/clients"
)

type DryRunClient struct {
	clients.CLI
	LineReader
}

func (c DryRunClient) WriteDryRun(ctx context.Context) error {
	r, closer, err := c.LineReader.Open(ctx)
	if closer != nil {
		defer closer.Close()
	}
	if err != nil {
		return err
	}

	if _, err := io.Copy(c.StdIO, r); err != nil {
		return err
	}

	return nil
}

package write

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/internal/api"
)

type LineReader interface {
	Open(ctx context.Context) (io.Reader, io.Closer, error)
}

type RateLimiter interface {
	Throttle(ctx context.Context, in io.Reader) io.Reader
}

type BatchWriter interface {
	WriteBatches(ctx context.Context, r io.Reader, writeFn func(batch []byte) error) error
}

type Client struct {
	internal.CLI
	api.WriteApi
	LineReader
	RateLimiter
	BatchWriter
}

type Params struct {
	BucketID   string
	BucketName string
	OrgID      string
	OrgName    string
	Precision  api.WritePrecision
}

var ErrWriteCanceled = errors.New("write canceled")

func (c Client) Write(ctx context.Context, params *Params) error {
	if params.OrgID == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return errors.New("must specify org ID or org name")
	}
	if params.BucketID == "" && params.BucketName == "" {
		return errors.New("must specify bucket ID or bucket name")
	}

	r, closer, err := c.LineReader.Open(ctx)
	if closer != nil {
		defer closer.Close()
	}
	if err != nil {
		return err
	}

	writeBatch := func(batch []byte) error {
		req := c.PostWrite(ctx).Body(batch).ContentEncoding("gzip").Precision(params.Precision)
		if params.BucketID != "" {
			req = req.Bucket(params.BucketID)
		} else {
			req = req.Bucket(params.BucketName)
		}
		if params.OrgID != "" {
			req = req.Org(params.OrgID)
		} else if params.OrgName != "" {
			req = req.Org(params.OrgName)
		} else {
			req = req.Org(c.ActiveConfig.Org)
		}

		if err := req.Execute(); err != nil {
			return err
		}

		return nil
	}

	if err := c.BatchWriter.WriteBatches(ctx, c.RateLimiter.Throttle(ctx, r), writeBatch); err == context.Canceled {
		return ErrWriteCanceled
	} else if err != nil {
		return fmt.Errorf("failed to write data: %w", err)
	}

	return nil
}

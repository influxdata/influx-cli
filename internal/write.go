package internal

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/influxdata/influx-cli/v2/internal/api"
)

type LineReader interface {
	Open(ctx context.Context) (io.Reader, io.Closer, error)
}

type Throttler interface {
	Throttle(ctx context.Context, in io.Reader) io.Reader
}

type Batcher interface {
	WriteBatches(ctx context.Context, r io.Reader, writeFn func(batch []byte) error) error
}

type WriteClients struct {
	Reader    LineReader
	Throttler Throttler
	Writer    Batcher
	Client    api.WriteApi
}

type WriteParams struct {
	BucketID   string
	BucketName string
	OrgID      string
	OrgName    string
	Precision  api.WritePrecision
}

var ErrWriteCanceled = errors.New("write canceled")

func (c *CLI) Write(ctx context.Context, clients *WriteClients, params *WriteParams) error {
	if params.OrgID == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return errors.New("must specify org ID or org name")
	}
	if params.BucketID == "" && params.BucketName == "" {
		return errors.New("must specify bucket ID or bucket name")
	}

	r, closer, err := clients.Reader.Open(ctx)
	if closer != nil {
		defer closer.Close()
	}
	if err != nil {
		return err
	}

	writeBatch := func(batch []byte) error {
		req := clients.Client.PostWrite(ctx).Body(batch).ContentEncoding("gzip").Precision(params.Precision)
		if c.TraceId != "" {
			req = req.ZapTraceSpan(c.TraceId)
		}
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

		if _, err := clients.Client.PostWriteExecute(req); err != nil {
			return err
		}

		return nil
	}

	if err := clients.Writer.WriteBatches(ctx, clients.Throttler.Throttle(ctx, r), writeBatch); err == context.Canceled {
		return ErrWriteCanceled
	} else if err != nil {
		return fmt.Errorf("failed to write data: %w", err)
	}

	return nil
}

func (c *CLI) WriteDryRun(ctx context.Context, reader LineReader) error {
	r, closer, err := reader.Open(ctx)
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

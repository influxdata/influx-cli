package batcher_test

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/influxdata/influx-cli/v2/internal/batcher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScanLines(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []string
		wantErr bool
	}{
		{
			name:  "3 lines produced including their newlines",
			input: "m1,t1=v1 f1=1\nm2,t2=v2 f2=2\nm3,t3=v3 f3=3",
			want:  []string{"m1,t1=v1 f1=1\n", "m2,t2=v2 f2=2\n", "m3,t3=v3 f3=3"},
		},
		{
			name:  "single line without newline",
			input: "m1,t1=v1 f1=1",
			want:  []string{"m1,t1=v1 f1=1"},
		},
		{
			name:  "single line with newline",
			input: "m1,t1=v1 f1=1\n",
			want:  []string{"m1,t1=v1 f1=1\n"},
		},
		{
			name:  "no lines",
			input: "",
			want:  []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.input))
			scanner.Split(batcher.ScanLines)
			got := []string{}
			for scanner.Scan() {
				got = append(got, scanner.Text())
			}
			err := scanner.Err()

			if (err != nil) != tt.wantErr {
				t.Errorf("ScanLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

// errorReader mocks io.Reader but returns an error.
type errorReader struct{}

func (r *errorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("error")
}

func TestBatcher_WriteTo(t *testing.T) {
	createReader := func(data string) func() io.Reader {
		if data == "error" {
			return func() io.Reader {
				return &errorReader{}
			}
		}
		return func() io.Reader {
			return strings.NewReader(data)
		}
	}

	type fields struct {
		MaxFlushBytes    int
		MaxFlushInterval time.Duration
	}
	type args struct {
		r          func() io.Reader
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        string
		wantFlushes int
		wantErr     bool
	}{
		{
			name: "a line of line protocol is sent to the service",
			fields: fields{
				MaxFlushBytes: 1,
			},
			args: args{
				r: createReader("m1,t1=v1 f1=1"),
			},
			want:        "m1,t1=v1 f1=1",
			wantFlushes: 1,
		},
		{
			name: "multiple lines cause multiple flushes",
			fields: fields{
				MaxFlushBytes: len([]byte("m1,t1=v1 f1=1\n")),
			},
			args: args{
				r: createReader("m1,t1=v1 f1=1\nm2,t2=v2 f2=2\nm3,t3=v3 f3=3"),
			},
			want:        "m3,t3=v3 f3=3",
			wantFlushes: 3,
		},
		{
			name:   "errors during read return error",
			fields: fields{},
			args: args{
				r: createReader("error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &batcher.BufferBatcher{
				MaxFlushBytes:    tt.fields.MaxFlushBytes,
				MaxFlushInterval: tt.fields.MaxFlushInterval,
			}

			// mocking the batcher service here to either return an error
			// or get back all the bytes from the reader.
			var (
				got        string
				gotFlushes int
			)
			err := b.WriteBatches(context.Background(), tt.args.r(), func(batch []byte) error {
				got = string(batch)
				gotFlushes++
				return nil
			})
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.wantFlushes, gotFlushes)
			require.Equal(t, tt.want, got)
		})
		// test the same data, but now with WriteBatches function
		t.Run("WriteTo_"+tt.name, func(t *testing.T) {
			b := &batcher.BufferBatcher{
				MaxFlushBytes:    tt.fields.MaxFlushBytes,
				MaxFlushInterval: tt.fields.MaxFlushInterval,
			}

			// mocking the batcher service here to either return an error
			// or get back all the bytes from the reader.
			var (
				got        string
				gotFlushes int
			)
			err := b.WriteBatches(context.Background(), tt.args.r(), func(batch []byte) error {
				got = string(batch)
				gotFlushes++
				return nil
			})
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.wantFlushes, gotFlushes)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBatcher_WriteTimeout(t *testing.T) {
	b := &batcher.BufferBatcher{}

	// this mimics a reader like stdin that may never return data.
	r, _ := io.Pipe()

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	var got string
	require.Equal(t, context.DeadlineExceeded, b.WriteBatches(ctx, r, func(batch []byte) error {
		got = string(batch)
		return nil
	}))
	require.Empty(t, got, "BufferBatcher.Write() with timeout received data")
}

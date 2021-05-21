package write

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// errorReader mocks io.Reader but returns an error.
type errorReader struct{}

func (r *errorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("error")
}

func TestBatcher_read(t *testing.T) {
	type args struct {
		cancel bool
		r      io.Reader
		max    int
	}
	tests := []struct {
		name   string
		args   args
		want   []string
		expErr error
	}{
		{
			name: "reading two lines produces 2 lines",
			args: args{
				r: strings.NewReader("m1,t1=v1 f1=1\nm2,t2=v2 f2=2"),
			},
			want: []string{"m1,t1=v1 f1=1\n", "m2,t2=v2 f2=2"},
		},
		{
			name: "canceling returns no lines",
			args: args{
				cancel: true,
				r:      strings.NewReader("m1,t1=v1 f1=1"),
			},
			want:   nil,
			expErr: context.Canceled,
		},
		{
			name: "error from reader returns error",
			args: args{
				r: &errorReader{},
			},
			want:   nil,
			expErr: fmt.Errorf("error"),
		},
		{
			name: "error when input exceeds max line length",
			args: args{
				r:   strings.NewReader("m1,t1=v1 f1=1"),
				max: 5,
			},
			want:   nil,
			expErr: ErrLineTooLong,
		},
		{
			name: "lines greater than MaxScanTokenSize are allowed",
			args: args{
				r:   strings.NewReader(strings.Repeat("a", bufio.MaxScanTokenSize+1)),
				max: bufio.MaxScanTokenSize + 2,
			},
			want: []string{strings.Repeat("a", bufio.MaxScanTokenSize+1)},
		},
		{
			name: "lines greater than MaxScanTokenSize by default are not allowed",
			args: args{
				r: strings.NewReader(strings.Repeat("a", bufio.MaxScanTokenSize+1)),
			},
			want:   nil,
			expErr: ErrLineTooLong,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			var cancel context.CancelFunc
			if tt.args.cancel {
				ctx, cancel = context.WithCancel(ctx)
				cancel()
			}

			b := &BufferBatcher{MaxLineLength: tt.args.max}
			var got []string

			lines := make(chan []byte)
			errC := make(chan error, 1)

			go b.read(ctx, tt.args.r, lines, errC)

			if cancel == nil {
				for line := range lines {
					got = append(got, string(line))
				}
			}

			err := <-errC
			assert.Equal(t, err, tt.expErr)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestBatcher_write(t *testing.T) {
	type fields struct {
		MaxFlushBytes    int
		MaxFlushInterval time.Duration
	}
	type args struct {
		cancel     bool
		writeError bool
		line       string
		lines      chan []byte
		errC       chan error
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       string
		wantErr    bool
		wantNoCall bool
	}{
		{
			name: "sending a single line will send a line to the service",
			fields: fields{
				MaxFlushBytes: 1,
			},
			args: args{
				line:   "m1,t1=v1 f1=1",
				lines:  make(chan []byte),
				errC:   make(chan error),
			},
			want: "m1,t1=v1 f1=1",
		},
		{
			name: "sending a single line and waiting for a timeout will send a line to the service",
			fields: fields{
				MaxFlushInterval: time.Millisecond,
			},
			args: args{
				line:   "m1,t1=v1 f1=1",
				lines:  make(chan []byte),
				errC:   make(chan error),
			},
			want: "m1,t1=v1 f1=1",
		},
		{
			name: "batcher service returning error stops the batcher after timeout",
			fields: fields{
				MaxFlushInterval: time.Millisecond,
			},
			args: args{
				writeError: true,
				line:       "m1,t1=v1 f1=1",
				lines:      make(chan []byte),
				errC:       make(chan error),
			},
			wantErr: true,
		},
		{
			name: "canceling will batcher no data to service",
			fields: fields{
				MaxFlushBytes: 1,
			},
			args: args{
				cancel: true,
				line:   "m1,t1=v1 f1=1",
				lines:  make(chan []byte, 1),
				errC:   make(chan error, 1),
			},
			wantErr:    true,
			wantNoCall: true,
		},
		{
			name: "batcher service returning error stops the batcher",
			fields: fields{
				MaxFlushBytes: 1,
			},
			args: args{
				writeError: true,
				line:       "m1,t1=v1 f1=1",
				lines:      make(chan []byte),
				errC:       make(chan error),
			},
			wantErr: true,
		},
		{
			name: "blank line is not sent to service",
			fields: fields{
				MaxFlushBytes: 1,
			},
			args: args{
				line:   "\n",
				lines:  make(chan []byte),
				errC:   make(chan error),
			},
			wantNoCall: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			var cancel context.CancelFunc
			if tt.args.cancel {
				ctx, cancel = context.WithCancel(ctx)
			}

			// mocking the batcher service here to either return an error
			// or get back all the bytes from the reader.
			writeCalled := false
			b := &BufferBatcher{
				MaxFlushBytes:    tt.fields.MaxFlushBytes,
				MaxFlushInterval: tt.fields.MaxFlushInterval,
			}
			var got string
			writeFn := func(batch []byte) error {
				writeCalled = true
				if tt.wantErr {
					return errors.New("I broke")
				}
				got = string(batch)
				return nil
			}

			go b.write(ctx, writeFn, tt.args.lines, tt.args.errC)

			if cancel != nil {
				cancel()
				time.Sleep(500 * time.Millisecond)
			}

			tt.args.lines <- []byte(tt.args.line)
			// if the max flush interval is not zero, we are testing to see
			// if the data is flushed via the timer rather than forced by
			// closing the channel.
			if tt.fields.MaxFlushInterval != 0 {
				time.Sleep(tt.fields.MaxFlushInterval * 100)
			}
			close(tt.args.lines)

			err := <-tt.args.errC
			require.Equal(t, tt.wantErr, err != nil)

			require.Equal(t, tt.wantNoCall, !writeCalled)
			require.Equal(t, tt.want, got)
		})
	}
}

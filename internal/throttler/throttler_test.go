package throttler_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/influxdata/influx-cli/v2/internal/throttler"
	"github.com/stretchr/testify/require"
)

func TestThrottlerPassthrough(t *testing.T) {
	// Hard to test that rate-limiting actually works, so we just check
	// that no data is lost.
	in := "Hello world!"
	throttler, err := throttler.NewThrottler("1B/s")
	require.NoError(t, err)
	r := throttler.Throttle(context.Background(), strings.NewReader(in))

	out := bytes.Buffer{}
	_, err = out.ReadFrom(r)
	require.NoError(t, err)

	require.Equal(t, in, out.String())
}

func TestToBytesPerSecond(t *testing.T) {
	var tests = []struct {
		in    string
		out   float64
		error string
	}{
		{
			in:  "5 MB / 5 min",
			out: float64(5*1024*1024) / float64(5*60),
		},
		{
			in:  "17kBs",
			out: float64(17 * 1024),
		},
		{
			in:  "1B/m",
			out: float64(1) / float64(60),
		},
		{
			in:  "1B/2sec",
			out: float64(1) / float64(2),
		},
		{
			in:  "",
			out: 0,
		},
		{
			in:    "1B/munite",
			error: `invalid rate limit "1B/munite": it does not match format COUNT(B|kB|MB)/TIME(s|sec|m|min) with / and TIME being optional`,
		},
		{
			in:    ".B/s",
			error: `invalid rate limit ".B/s": '.' is not count of bytes:`,
		},
		{
			in:    "1B0s",
			error: `invalid rate limit "1B0s": positive time expected but 0 supplied`,
		},
		{
			in:    "1MB/42949672950s",
			error: `invalid rate limit "1MB/42949672950s": time is out of range`,
		},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			bytesPerSec, err := throttler.ToBytesPerSecond(test.in)
			if len(test.error) == 0 {
				require.Equal(t, test.out, bytesPerSec)
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
				// contains is used, since the error messages contains root cause that may evolve with go versions
				require.Contains(t, err.Error(), test.error)
			}
		})
	}
}

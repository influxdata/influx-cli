package csv2lp

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLineProtocolFilter(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{
			strings.Join([]string{
				"weather,location=us-midwest temperature=42 1465839830100400200",
				"awefw.,weather,location=us-east temperature=36 1465839830100400203",
				"weather,location=us-blah temperature=32 1465839830100400204",
				"weather,location=us-central temperature=31 1465839830100400205",
			}, "\n"),
			strings.Join([]string{
				"weather,location=us-midwest temperature=42 1465839830100400200",
				"weather,location=us-blah temperature=32 1465839830100400204",
				"weather,location=us-central temperature=31 1465839830100400205",
			}, "\n"),
		},
		{
			strings.Join([]string{
				"weather,location=us-midwest temperature=42 1465839830100400200",
				"weather,location=us-east temperature=36 1465839830100400203",
				"weather,location=us-blah temperature=32=33 1465839830100400204",
				"weather,,location=us-blah temperature=32 1465839830100400204",
				"weather,location=us-central temperature=31 1465839830100400205",
			}, "\n"),
			strings.Join([]string{
				"weather,location=us-midwest temperature=42 1465839830100400200",
				"weather,location=us-east temperature=36 1465839830100400203",
				"weather,location=us-central temperature=31 1465839830100400205",
			}, "\n"),
		},
		{
			strings.Join([]string{
				"weather,location=us-midwest temperature=42 1465839830100400200",
				"awefw.,weather,location=us-east temperature=36 1465839830100400203",
				"    weather,location=us-blah temperature=32 1465839830100400204",
				"weather,location=us-east temperature=36 1465839830100400203 13413413",
				"weather,location=us-central temperature=31 1465839830100400205",
				"# this is a comment",
			}, "\n"),
			strings.Join([]string{
				"weather,location=us-midwest temperature=42 1465839830100400200",
				"    weather,location=us-blah temperature=32 1465839830100400204",
				"weather,location=us-central temperature=31 1465839830100400205",
			}, "\n"),
		},
	}
	for _, tt := range tests {
		reader := LineProtocolFilter(strings.NewReader(tt.input))
		b, err := io.ReadAll(reader)
		if err != nil {
			t.Errorf("failed reading: %v", err)
			continue
		}
		require.Equal(t, strings.TrimSpace(string(b)), strings.TrimSpace(tt.expected))
	}
}

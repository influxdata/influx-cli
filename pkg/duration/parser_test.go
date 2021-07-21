package duration_test

import (
	"testing"
	"time"

	"github.com/influxdata/influx-cli/v2/pkg/duration"
	"github.com/stretchr/testify/require"
)

func Test_RawDurationToTimeDuration(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		expected  time.Duration
		expectErr bool
	}{
		{
			name:     "nanos",
			input:    "10ns",
			expected: 10,
		},
		{
			name:     "micros",
			input:    "12345us",
			expected: 12345 * time.Microsecond,
		},
		{
			name:     "millis",
			input:    "9876ms",
			expected: 9876 * time.Millisecond,
		},
		{
			name:     "seconds",
			input:    "300s",
			expected: 300 * time.Second,
		},
		{
			name:     "minutes",
			input:    "654m",
			expected: 654 * time.Minute,
		},
		{
			name:     "hours",
			input:    "127h",
			expected: 127 * time.Hour,
		},
		{
			name:     "days",
			input:    "29d",
			expected: 29 * duration.Day,
		},
		{
			name:     "weeks",
			input:    "396w",
			expected: 396 * duration.Week,
		},
		{
			name:     "weeks+hours+seconds+micros",
			input:    "1w2h3s4us",
			expected: duration.Week + 2*time.Hour + 3*time.Second + 4*time.Microsecond,
		},
		{
			name:     "days+minutes+millis+nanos",
			input:    "9d8m7ms6ns",
			expected: 9*duration.Day + 8*time.Minute + 7*time.Millisecond + 6,
		},
		{
			name:      "negative",
			input:     "-1d",
			expectErr: true,
		},
		{
			name:      "missing unit",
			input:     "123",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parsed, err := duration.RawDurationToTimeDuration(tc.input)
			if tc.expectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expected, parsed)
		})
	}
}

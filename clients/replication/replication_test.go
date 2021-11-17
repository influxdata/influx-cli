package replication

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDropNonRetryableDataBoolPtrFromFlags(t *testing.T) {
	tests := []struct {
		name                   string
		dropNonRetryableData   bool
		noDropNonRetryableData bool
		want                   *bool
		wantErr                error
	}{
		{
			name:                   "both true is an error",
			dropNonRetryableData:   true,
			noDropNonRetryableData: true,
			want:                   nil,
			wantErr:                errors.New("cannot specify both --drop-non-retryable-data and --no-drop-non-retryable-data at the same time"),
		},
		{
			name:                 "drop is true",
			dropNonRetryableData: true,
			want:                 boolPtr(true),
		},
		{
			name:                   "noDrop is true",
			noDropNonRetryableData: true,
			want:                   boolPtr(false),
		},
		{
			name: "both nil is nil",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dropNonRetryableDataBoolPtrFromFlags(tt.dropNonRetryableData, tt.noDropNonRetryableData)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantErr, err)
		})
	}
}

package testutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func MatchLines(t *testing.T, expectedLines []string, lines []string) {
	var nonEmptyLines []string
	for _, l := range lines {
		if l != "" {
			nonEmptyLines = append(nonEmptyLines, l)
		}
	}
	require.Equal(t, len(expectedLines), len(nonEmptyLines))
	for i, expected := range expectedLines {
		require.Regexp(t, expected, nonEmptyLines[i])
	}
}

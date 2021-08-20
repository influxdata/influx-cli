package github_test

import (
	"net/url"
	"testing"

	"github.com/influxdata/influx-cli/v2/pkg/github"
	"github.com/stretchr/testify/require"
)

func TestNormalize(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		in   url.URL
		exts []string
		out  url.URL
	}{
		{
			name: "github URL",
			in:   url.URL{Host: "github.com", Path: "/influxdata/influxdb/blob/master/flags.yml"},
			out:  url.URL{Host: "raw.githubusercontent.com", Path: "/influxdata/influxdb/master/flags.yml"},
		},
		{
			name: "github URL with extensions",
			in:   url.URL{Host: "github.com", Path: "/influxdata/community-templates/blob/master/github/github.yml"},
			exts: []string{".yaml", ".yml", ".jsonnet", ".json"},
			out:  url.URL{Host: "raw.githubusercontent.com", Path: "/influxdata/community-templates/master/github/github.yml"},
		},
		{
			name: "other URL",
			in:   url.URL{Host: "google.com", Path: "/fake.yml"},
			out:  url.URL{Host: "google.com", Path: "/fake.yml"},
		},
		{
			name: "github URL - wrong extension",
			in:   url.URL{Host: "github.com", Path: "/influxdata/influxdb/blob/master/flags.yml"},
			exts: []string{".json"},
			out:  url.URL{Host: "github.com", Path: "/influxdata/influxdb/blob/master/flags.yml"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			normalized := github.NormalizeURLToContent(&tc.in, tc.exts...)
			require.Equal(t, tc.out, *normalized)
		})
	}
}

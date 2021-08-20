package auth

import (
	"testing"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/stretchr/testify/require"
)

func Test_makePermResource(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		inType   string
		inId     string
		inOrgId  string
		expected api.PermissionResource
	}{
		{
			name:     "only type",
			inType:   "foo",
			expected: api.PermissionResource{Type: "foo"},
		},
		{
			name:     "type and ID",
			inType:   "bar",
			inId:     "12345",
			expected: api.PermissionResource{Type: "bar", Id: api.PtrString("12345")},
		},
		{
			name:     "type and org ID",
			inType:   "baz",
			inOrgId:  "45678",
			expected: api.PermissionResource{Type: "baz", OrgID: api.PtrString("45678")},
		},
		{
			name:     "type, ID, and org ID",
			inType:   "qux",
			inId:     "12345",
			inOrgId:  "45678",
			expected: api.PermissionResource{Type: "qux", Id: api.PtrString("12345"), OrgID: api.PtrString("45678")},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.expected, makePermResource(tc.inType, tc.inId, tc.inOrgId))
		})
	}
}

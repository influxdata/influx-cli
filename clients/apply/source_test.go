package apply_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients/apply"
	"github.com/stretchr/testify/require"
)

func TestSourcesFromPath(t *testing.T) {
	t.Parallel()

	type contents struct {
		name     string
		encoding apply.TemplateEncoding
		contents string
	}
	testCases := []struct {
		name       string
		setup      func(t *testing.T, rootDir string)
		inPath     func(rootDir string) string
		inEncoding apply.TemplateEncoding
		recursive  bool
		expected   func(rootDir string) []contents
	}{
		{
			name: "JSON file",
			setup: func(t *testing.T, rootDir string) {
				require.NoError(t, os.WriteFile(filepath.Join(rootDir, "foo.json"), []byte("foo"), os.ModePerm))
			},
			inPath: func(rootDir string) string {
				return filepath.Join(rootDir, "foo.json")
			},
			expected: func(string) []contents {
				return []contents{{
					name:     "foo.json",
					encoding: apply.TemplateEncodingJson,
					contents: "foo",
				}}
			},
		},
		{
			name: "YAML file",
			setup: func(t *testing.T, rootDir string) {
				require.NoError(t, os.WriteFile(filepath.Join(rootDir, "foo.yaml"), []byte("foo"), os.ModePerm))
			},
			inPath: func(rootDir string) string {
				return filepath.Join(rootDir, "foo.yaml")
			},
			expected: func(string) []contents {
				return []contents{{
					name:     "foo.yaml",
					encoding: apply.TemplateEncodingYaml,
					contents: "foo",
				}}
			},
		},
		{
			name: "YML file",
			setup: func(t *testing.T, rootDir string) {
				require.NoError(t, os.WriteFile(filepath.Join(rootDir, "foo.yml"), []byte("foo"), os.ModePerm))
			},
			inPath: func(rootDir string) string {
				return filepath.Join(rootDir, "foo.yml")
			},
			expected: func(string) []contents {
				return []contents{{
					name:     "foo.yml",
					encoding: apply.TemplateEncodingYaml,
					contents: "foo",
				}}
			},
		},
		{
			name: "JSONNET file",
			setup: func(t *testing.T, rootDir string) {
				require.NoError(t, os.WriteFile(filepath.Join(rootDir, "foo.jsonnet"), []byte("foo"), os.ModePerm))
			},
			inPath: func(rootDir string) string {
				return filepath.Join(rootDir, "foo.jsonnet")
			},
			expected: func(string) []contents {
				return []contents{{
					name:     "foo.jsonnet",
					encoding: apply.TemplateEncodingJsonnet,
					contents: "foo",
				}}
			},
		},
		{
			name: "explicit inEncoding",
			setup: func(t *testing.T, rootDir string) {
				require.NoError(t, os.WriteFile(filepath.Join(rootDir, "foo"), []byte("foo"), os.ModePerm))
			},
			inEncoding: apply.TemplateEncodingJson,
			inPath: func(rootDir string) string {
				return filepath.Join(rootDir, "foo")
			},
			expected: func(string) []contents {
				return []contents{{
					name:     "foo",
					encoding: apply.TemplateEncodingJson,
					contents: "foo",
				}}
			},
		},
		{
			name: "directory - non-recursive",
			setup: func(t *testing.T, rootDir string) {
				require.NoError(t, os.Mkdir(filepath.Join(rootDir, "bar"), os.ModePerm))
				require.NoError(t, os.WriteFile(filepath.Join(rootDir, "foo.json"), []byte("foo.json"), os.ModePerm))
				require.NoError(t, os.WriteFile(filepath.Join(rootDir, "foo.yml"), []byte("foo.yml"), os.ModePerm))
				require.NoError(t, os.WriteFile(filepath.Join(rootDir, "bar", "foo.jsonnet"), []byte("foo.jsonnet"), os.ModePerm))
				require.NoError(t, os.WriteFile(filepath.Join(rootDir, "bar", "foo.yaml"), []byte("foo.yaml"), os.ModePerm))
			},
			inPath: func(rootDir string) string {
				return rootDir
			},
			expected: func(rootDir string) []contents {
				return []contents{
					{
						name:     filepath.Join(rootDir, "foo.json"),
						contents: "foo.json",
						encoding: apply.TemplateEncodingJson,
					},
					{
						name:     filepath.Join(rootDir, "foo.yml"),
						contents: "foo.yml",
						encoding: apply.TemplateEncodingYaml,
					},
				}
			},
		},
		{
			name: "directory - recursive",
			setup: func(t *testing.T, rootDir string) {
				require.NoError(t, os.Mkdir(filepath.Join(rootDir, "bar"), os.ModePerm))
				require.NoError(t, os.WriteFile(filepath.Join(rootDir, "foo.json"), []byte("foo.json"), os.ModePerm))
				require.NoError(t, os.WriteFile(filepath.Join(rootDir, "foo.yml"), []byte("foo.yml"), os.ModePerm))
				require.NoError(t, os.WriteFile(filepath.Join(rootDir, "bar", "foo.jsonnet"), []byte("foo.jsonnet"), os.ModePerm))
				require.NoError(t, os.WriteFile(filepath.Join(rootDir, "bar", "foo.yaml"), []byte("foo.yaml"), os.ModePerm))
			},
			inPath: func(rootDir string) string {
				return rootDir
			},
			recursive: true,
			expected: func(rootDir string) []contents {
				return []contents{
					{
						name:     filepath.Join(rootDir, "foo.json"),
						contents: "foo.json",
						encoding: apply.TemplateEncodingJson,
					},
					{
						name:     filepath.Join(rootDir, "foo.yml"),
						contents: "foo.yml",
						encoding: apply.TemplateEncodingYaml,
					},
					{
						name:     filepath.Join(rootDir, "bar", "foo.yaml"),
						contents: "foo.yaml",
						encoding: apply.TemplateEncodingYaml,
					},
					{
						name:     filepath.Join(rootDir, "bar", "foo.jsonnet"),
						contents: "foo.jsonnet",
						encoding: apply.TemplateEncodingJsonnet,
					},
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tmp, err := os.MkdirTemp("", "")
			require.NoError(t, err)
			defer os.RemoveAll(tmp)
			tc.setup(t, tmp)

			sources, err := apply.SourcesFromPath(tc.inPath(tmp), tc.recursive, tc.inEncoding)
			require.NoError(t, err)
			expected := tc.expected(tmp)
			require.Len(t, sources, len(expected))

			sort.Slice(sources, func(i, j int) bool {
				return sources[i].Name < sources[j].Name
			})
			sort.Slice(expected, func(i, j int) bool {
				return expected[i].name < expected[j].name
			})

			for i := range expected {
				source := sources[i]
				contents := expected[i]

				require.Equal(t, contents.encoding, source.Encoding)
				sourceIn, err := source.Open(context.Background())
				require.NoError(t, err)
				bytes, err := io.ReadAll(sourceIn)
				sourceIn.Close()
				require.NoError(t, err)
				require.Equal(t, contents.contents, string(bytes))
			}
		})
	}
}

func TestSourceFromURL(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		filename       string
		inEncoding     apply.TemplateEncoding
		expectEncoding apply.TemplateEncoding
		resStatus      int
		resBody        string
		expectErr      bool
	}{
		{
			name:           "JSON file",
			filename:       "foo.json",
			expectEncoding: apply.TemplateEncodingJson,
			resStatus:      200,
			resBody:        "Foo bar",
		},
		{
			name:           "YAML file",
			filename:       "foo.yaml",
			expectEncoding: apply.TemplateEncodingYaml,
			resStatus:      200,
			resBody:        "Foo bar",
		},
		{
			name:           "YML file",
			filename:       "foo.yml",
			expectEncoding: apply.TemplateEncodingYaml,
			resStatus:      200,
			resBody:        "Foo bar",
		},
		{
			name:           "JSONNET file",
			filename:       "foo.jsonnet",
			expectEncoding: apply.TemplateEncodingJsonnet,
			resStatus:      200,
			resBody:        "Foo bar",
		},
		{
			name:           "explicit encoding",
			filename:       "foo",
			inEncoding:     apply.TemplateEncodingJson,
			expectEncoding: apply.TemplateEncodingJson,
			resStatus:      200,
			resBody:        "Foo bar",
		},
		{
			name:           "err response",
			filename:       "foo.json",
			expectEncoding: apply.TemplateEncodingJson,
			resStatus:      403,
			resBody:        "OH NO",
			expectErr:      true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				rw.WriteHeader(tc.resStatus)
				rw.Write([]byte(tc.resBody))
				return
			}))
			defer server.Close()

			u, err := url.Parse(server.URL)
			require.NoError(t, err)
			u.Path = tc.filename

			source := apply.SourceFromURL(u, tc.inEncoding)
			in, err := source.Open(context.Background())
			if tc.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), "bad response")
				require.Contains(t, err.Error(), tc.resBody)
				return
			}
			require.NoError(t, err)
			defer in.Close()
			bytes, err := io.ReadAll(in)
			require.NoError(t, err)
			require.Equal(t, tc.resBody, string(bytes))
		})
	}
}

func TestTemplateSource_Read(t *testing.T) {
	t.Parallel()

	yamlTemplate := `---
apiversion: influxdata.com/v2alpha1
kind: Bucket
meta: null
spec:
    name: test
    retentionRules:
      - type: expire
---
apiversion: influxdata.com/v2alpha1
kind: Bucket
meta: null
spec:
    name: test2
    retentionRules:
      - type: expire
`
	jsonTemplate := `[
	{
		"apiVersion": "influxdata.com/v2alpha1",
		"kind": "Bucket",
		"spec": {
			"name": "test",
			"retentionRules": [
				{
					"type": "expire"
				}
			]
		}
	},
	{
		"apiVersion": "influxdata.com/v2alpha1",
		"kind": "Bucket",
		"spec": {
			"name": "test2",
			"retentionRules": [
				{
					"type": "expire"
				}
			]
		}
	}
]
`
	jsonnetTemplate := `local Bucket(name) = {
	apiVersion: "influxdata.com/v2alpha1",
	kind: "Bucket",
	spec: {
		name: name,
		retentionRules: [{
			type: "expire",
		}],
	},
};
[Bucket("test"), Bucket("test2")]
`

	parsed := []api.TemplateEntry{
		{
			ApiVersion: api.PtrString("influxdata.com/v2alpha1"),
			Kind:       api.PtrString("Bucket"),
			Spec: &map[string]interface{}{
				"name": "test",
				"retentionRules": []interface{}{
					map[string]interface{}{
						"type":         "expire",
					},
				},
			},
		},
		{
			ApiVersion: api.PtrString("influxdata.com/v2alpha1"),
			Kind:       api.PtrString("Bucket"),
			Spec: &map[string]interface{}{
				"name": "test2",
				"retentionRules": []interface{}{
					map[string]interface{}{
						"type":         "expire",
					},
				},
			},
		},
	}

	testCases := []struct {
		name     string
		encoding apply.TemplateEncoding
		data     string
	}{
		{
			name:     "JSON",
			encoding: apply.TemplateEncodingJson,
			data:     jsonTemplate,
		},
		{
			name:     "YAML",
			encoding: apply.TemplateEncodingYaml,
			data:     yamlTemplate,
		},
		{
			name:     "JSONNET",
			encoding: apply.TemplateEncodingJsonnet,
			data:     jsonnetTemplate,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			source := apply.SourceFromReader(strings.NewReader(tc.data), tc.encoding)
			tmpls, err := source.Read(context.Background())
			require.NoError(t, err)

			require.Equal(t, parsed, tmpls)
		})
	}
}

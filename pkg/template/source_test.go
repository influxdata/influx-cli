package template_test

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
	"github.com/influxdata/influx-cli/v2/pkg/template"
	"github.com/stretchr/testify/require"
)

func TestSourcesFromPath(t *testing.T) {
	t.Parallel()

	type contents struct {
		name     string
		encoding template.Encoding
		contents string
	}
	testCases := []struct {
		name       string
		setup      func(t *testing.T, rootDir string)
		inPath     func(rootDir string) string
		inEncoding template.Encoding
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
					encoding: template.EncodingJson,
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
					encoding: template.EncodingYaml,
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
					encoding: template.EncodingYaml,
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
					encoding: template.EncodingJsonnet,
					contents: "foo",
				}}
			},
		},
		{
			name: "explicit inEncoding",
			setup: func(t *testing.T, rootDir string) {
				require.NoError(t, os.WriteFile(filepath.Join(rootDir, "foo"), []byte("foo"), os.ModePerm))
			},
			inEncoding: template.EncodingJson,
			inPath: func(rootDir string) string {
				return filepath.Join(rootDir, "foo")
			},
			expected: func(string) []contents {
				return []contents{{
					name:     "foo",
					encoding: template.EncodingJson,
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
						encoding: template.EncodingJson,
					},
					{
						name:     filepath.Join(rootDir, "foo.yml"),
						contents: "foo.yml",
						encoding: template.EncodingYaml,
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
						encoding: template.EncodingJson,
					},
					{
						name:     filepath.Join(rootDir, "foo.yml"),
						contents: "foo.yml",
						encoding: template.EncodingYaml,
					},
					{
						name:     filepath.Join(rootDir, "bar", "foo.yaml"),
						contents: "foo.yaml",
						encoding: template.EncodingYaml,
					},
					{
						name:     filepath.Join(rootDir, "bar", "foo.jsonnet"),
						contents: "foo.jsonnet",
						encoding: template.EncodingJsonnet,
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

			sources, err := template.SourcesFromPath(tc.inPath(tmp), tc.recursive, tc.inEncoding)
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
		inEncoding     template.Encoding
		expectEncoding template.Encoding
		resStatus      int
		resBody        string
		expectErr      bool
	}{
		{
			name:           "JSON file",
			filename:       "foo.json",
			expectEncoding: template.EncodingJson,
			resStatus:      200,
			resBody:        "Foo bar",
		},
		{
			name:           "YAML file",
			filename:       "foo.yaml",
			expectEncoding: template.EncodingYaml,
			resStatus:      200,
			resBody:        "Foo bar",
		},
		{
			name:           "YML file",
			filename:       "foo.yml",
			expectEncoding: template.EncodingYaml,
			resStatus:      200,
			resBody:        "Foo bar",
		},
		{
			name:           "JSONNET file",
			filename:       "foo.jsonnet",
			expectEncoding: template.EncodingJsonnet,
			resStatus:      200,
			resBody:        "Foo bar",
		},
		{
			name:           "explicit encoding",
			filename:       "foo",
			inEncoding:     template.EncodingJson,
			expectEncoding: template.EncodingJson,
			resStatus:      200,
			resBody:        "Foo bar",
		},
		{
			name:           "err response",
			filename:       "foo.json",
			expectEncoding: template.EncodingJson,
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
			}))
			defer server.Close()

			u, err := url.Parse(server.URL)
			require.NoError(t, err)
			u.Path = tc.filename

			source := template.SourceFromURL(u, tc.inEncoding)
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
apiVersion: influxdata.com/v2alpha1
kind: Bucket
meta: null
spec:
    name: test
    retentionRules:
      - type: expire
---
apiVersion: influxdata.com/v2alpha1
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
			ApiVersion: "influxdata.com/v2alpha1",
			Kind:       "Bucket",
			Spec: map[string]interface{}{
				"name": "test",
				"retentionRules": []interface{}{
					map[string]interface{}{
						"type": "expire",
					},
				},
			},
		},
		{
			ApiVersion: "influxdata.com/v2alpha1",
			Kind:       "Bucket",
			Spec: map[string]interface{}{
				"name": "test2",
				"retentionRules": []interface{}{
					map[string]interface{}{
						"type": "expire",
					},
				},
			},
		},
	}

	testCases := []struct {
		name     string
		encoding template.Encoding
		data     string
	}{
		{
			name:     "JSON",
			encoding: template.EncodingJson,
			data:     jsonTemplate,
		},
		{
			name:     "YAML",
			encoding: template.EncodingYaml,
			data:     yamlTemplate,
		},
		{
			name:     "JSONNET",
			encoding: template.EncodingJsonnet,
			data:     jsonnetTemplate,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			source := template.SourceFromReader(strings.NewReader(tc.data), tc.encoding)
			tmpl, err := source.Read(context.Background())
			require.NoError(t, err)
			expected := api.TemplateApplyTemplate{
				Sources:     []string{source.Name},
				Contents:    parsed,
			}
			require.Equal(t, expected, tmpl)
		})
	}
}

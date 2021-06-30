package template_test

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/pkg/template"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

var tmpls = []api.TemplateEntry{
	{
		ApiVersion: "api1",
		Kind:       "Foo",
		Metadata: api.TemplateEntryMetadata{
			Name: "foo",
		},
		Spec: map[string]interface{}{
			"hello":   "world",
			"1 + 1 =": "2",
		},
	},
	{
		ApiVersion: "api1",
		Kind:       "Bar",
		Metadata: api.TemplateEntryMetadata{
			Name: "bar",
		},
		Spec: map[string]interface{}{
			"success?": "true",
		},
	},
}

func TestOutParams(t *testing.T) {
	t.Parallel()

	t.Run("json to file", func(t *testing.T) {
		t.Parallel()

		tmp, err := os.MkdirTemp("", "")
		require.NoError(t, err)
		defer os.RemoveAll(tmp)

		out := filepath.Join(tmp, "test.json")
		params, closer, err := template.ParseOutParams(out, nil)
		require.NoError(t, err)
		require.NotNil(t, closer)
		defer closer()

		require.NoError(t, params.WriteTemplate(tmpls))
		contents, err := os.ReadFile(out)
		require.NoError(t, err)

		var written []api.TemplateEntry
		dec := json.NewDecoder(bytes.NewReader(contents))
		require.NoError(t, dec.Decode(&written))

		require.Equal(t, tmpls, written)
	})

	t.Run("yaml to file", func(t *testing.T) {
		t.Parallel()

		tmp, err := os.MkdirTemp("", "")
		require.NoError(t, err)
		defer os.RemoveAll(tmp)

		out := filepath.Join(tmp, "test.yaml")
		params, closer, err := template.ParseOutParams(out, nil)
		require.NoError(t, err)
		require.NotNil(t, closer)
		defer closer()

		require.NoError(t, params.WriteTemplate(tmpls))
		contents, err := os.ReadFile(out)
		require.NoError(t, err)

		var written []api.TemplateEntry
		dec := yaml.NewDecoder(bytes.NewReader(contents))
		for {
			var e api.TemplateEntry
			err := dec.Decode(&e)
			if err == io.EOF {
				break
			}
			require.NoError(t, err)
			written = append(written, e)
		}

		require.Equal(t, tmpls, written)
	})

	t.Run("yaml to buffer", func(t *testing.T) {
		t.Parallel()

		out := bytes.Buffer{}
		params, closer, err := template.ParseOutParams("", &out)
		require.NoError(t, err)
		require.Nil(t, closer)

		require.NoError(t, params.WriteTemplate(tmpls))

		var written []api.TemplateEntry
		dec := yaml.NewDecoder(&out)
		for {
			var e api.TemplateEntry
			err := dec.Decode(&e)
			if err == io.EOF {
				break
			}
			require.NoError(t, err)
			written = append(written, e)
		}

		require.Equal(t, tmpls, written)
	})
}

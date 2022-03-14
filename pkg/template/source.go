package template

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/pkg/github"
	"github.com/influxdata/influx-cli/v2/pkg/jsonnet"
	"gopkg.in/yaml.v3"
)

type Encoding int

const (
	EncodingUnknown Encoding = iota
	EncodingJson
	EncodingJsonnet
	EncodingYaml
)

func (e *Encoding) Set(v string) error {
	switch v {
	case "jsonnet":
		*e = EncodingJsonnet
	case "json":
		*e = EncodingJson
	case "yml", "yaml":
		*e = EncodingYaml
	default:
		return fmt.Errorf("unknown inEncoding %q", v)
	}
	return nil
}

func (e Encoding) String() string {
	switch e {
	case EncodingJsonnet:
		return "jsonnet"
	case EncodingJson:
		return "json"
	case EncodingYaml:
		return "yaml"
	case EncodingUnknown:
		fallthrough
	default:
		return "unknown"
	}
}

type Source struct {
	Name     string
	Encoding Encoding
	Open     func(context.Context) (io.ReadCloser, error)
}

func SourcesFromPath(path string, recur bool, encoding Encoding) ([]Source, error) {
	paths, err := findPaths(path, recur)
	if err != nil {
		return nil, fmt.Errorf("failed to find inputs at path %q: %w", path, err)
	}

	sources := make([]Source, len(paths))
	for i := range paths {
		path := paths[i] // Local var for the `Open` closure to capture.
		encoding := encoding
		if encoding == EncodingUnknown {
			switch filepath.Ext(path) {
			case ".jsonnet":
				encoding = EncodingJsonnet
			case ".json":
				encoding = EncodingJson
			case ".yml":
				fallthrough
			case ".yaml":
				encoding = EncodingYaml
			default:
			}
		}

		sources[i] = Source{
			Name:     path,
			Encoding: encoding,
			Open: func(context.Context) (io.ReadCloser, error) {
				return os.Open(path)
			},
		}
	}
	return sources, nil
}

func findPaths(path string, recur bool) ([]string, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return []string{path}, nil
	}

	dirFiles, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, df := range dirFiles {
		fullPath := filepath.Join(path, df.Name())
		if df.IsDir() && recur {
			subPaths, err := findPaths(fullPath, recur)
			if err != nil {
				return nil, err
			}
			paths = append(paths, subPaths...)
		} else if !df.IsDir() {
			paths = append(paths, fullPath)
		}
	}
	return paths, nil
}

func SourceFromURL(u *url.URL, encoding Encoding) Source {
	if encoding == EncodingUnknown {
		switch path.Ext(u.Path) {
		case ".jsonnet":
			encoding = EncodingJsonnet
		case ".json":
			encoding = EncodingJson
		case ".yml":
			fallthrough
		case ".yaml":
			encoding = EncodingYaml
		default:
		}
	}

	normalized := github.NormalizeURLToContent(u, ".yaml", ".yml", ".jsonnet", ".json").String()

	return Source{
		Name:     normalized,
		Encoding: encoding,
		Open: func(ctx context.Context) (io.ReadCloser, error) {
			req, err := http.NewRequestWithContext(ctx, "GET", normalized, nil)
			if err != nil {
				return nil, err
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				return nil, err
			}
			if res.StatusCode/100 != 2 {
				body, err := io.ReadAll(res.Body)
				res.Body.Close()
				if err != nil {
					return nil, err
				}
				return nil, fmt.Errorf("bad response: address=%s status_code=%d body=%q",
					normalized, res.StatusCode, strings.TrimSpace(string(body)))
			}
			return res.Body, nil
		},
	}
}

func SourceFromReader(r io.Reader, encoding Encoding) Source {
	return Source{
		Name:     "byte stream",
		Encoding: encoding,
		Open: func(context.Context) (io.ReadCloser, error) {
			return io.NopCloser(r), nil
		},
	}
}

func ReadSources(ctx context.Context, sources []Source) ([]api.TemplateApplyTemplate, error) {
	templates := make([]api.TemplateApplyTemplate, 0, len(sources))
	for _, source := range sources {
		tmpl, err := source.Read(ctx)
		if err != nil {
			return nil, err
		}
		// We always send the templates as JSON.
		tmpl.ContentType = "json"
		templates = append(templates, tmpl)
	}
	return templates, nil
}

func (s Source) Read(ctx context.Context) (api.TemplateApplyTemplate, error) {
	var entries []api.TemplateEntry
	if err := func() error {
		in, err := s.Open(ctx)
		if err != nil {
			return err
		}
		defer in.Close()

		switch s.Encoding {
		case EncodingJsonnet:
			err = jsonnet.NewDecoder(in).Decode(&entries)
		case EncodingJson:
			err = json.NewDecoder(in).Decode(&entries)
		case EncodingUnknown:
			fallthrough // Assume YAML if we can't make a better guess
		case EncodingYaml:
			dec := yaml.NewDecoder(in)
			for {
				var e api.TemplateEntry
				if err := dec.Decode(&e); err == io.EOF {
					break
				} else if err != nil {
					return err
				}
				entries = append(entries, e)
			}
		}
		return err
	}(); err != nil {
		return api.TemplateApplyTemplate{}, fmt.Errorf("failed to read template(s) from %q: %w", s.Name, err)
	}

	return api.TemplateApplyTemplate{
		Sources:  []string{s.Name + ".generated.json"},
		Contents: entries,
	}, nil
}

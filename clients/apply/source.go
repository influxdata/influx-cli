package apply

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

type TemplateEncoding int

const (
	TemplateEncodingUnknown TemplateEncoding = iota
	TemplateEncodingJson
	TemplateEncodingJsonnet
	TemplateEncodingYaml
)

func (e *TemplateEncoding) Set(v string) error {
	switch v {
	case "jsonnet":
		*e = TemplateEncodingJsonnet
	case "json":
		*e = TemplateEncodingJson
	case "yml", "yaml":
		*e = TemplateEncodingYaml
	default:
		return fmt.Errorf("unknown inEncoding %q", v)
	}
	return nil
}

func (e TemplateEncoding) String() string {
	switch e {
	case TemplateEncodingJsonnet:
		return "jsonnet"
	case TemplateEncodingJson:
		return "json"
	case TemplateEncodingYaml:
		return "yaml"
	case TemplateEncodingUnknown:
		fallthrough
	default:
		return "unknown"
	}
}

type TemplateSource struct {
	Name     string
	Encoding TemplateEncoding
	Open     func(context.Context) (io.ReadCloser, error)
}

func SourcesFromPath(path string, recur bool, encoding TemplateEncoding) ([]TemplateSource, error) {
	paths, err := findPaths(path, recur)
	if err != nil {
		return nil, fmt.Errorf("failed to find inputs at path %q: %w", path, err)
	}

	sources := make([]TemplateSource, len(paths))
	for i := range paths {
		path := paths[i] // Local var for the `Open` closure to capture.
		encoding := encoding
		if encoding == TemplateEncodingUnknown {
			ext := filepath.Ext(path)
			switch {
			case strings.HasPrefix(ext, ".jsonnet"):
				encoding = TemplateEncodingJsonnet
			case strings.HasPrefix(ext, ".json"):
				encoding = TemplateEncodingJson
			case strings.HasPrefix(ext, ".yml") || strings.HasPrefix(ext, ".yaml"):
				encoding = TemplateEncodingYaml
			default:
			}
		}

		sources[i] = TemplateSource{
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

func SourceFromURL(u *url.URL, encoding TemplateEncoding) TemplateSource {
	if encoding == TemplateEncodingUnknown {
		ext := path.Ext(u.Path)
		switch {
		case strings.HasPrefix(ext, ".jsonnet"):
			encoding = TemplateEncodingJsonnet
		case strings.HasPrefix(ext, ".json"):
			encoding = TemplateEncodingJson
		case strings.HasPrefix(ext, ".yml") || strings.HasPrefix(ext, ".yaml"):
			encoding = TemplateEncodingYaml
		default:
		}
	}

	normalized := github.NormalizeURLToContent(u, "yaml", "yml", "jsonnet", "json").String()

	return TemplateSource{
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

func SourceFromReader(r io.Reader, encoding TemplateEncoding) TemplateSource {
	return TemplateSource{
		Name:     "byte stream",
		Encoding: encoding,
		Open: func(context.Context) (io.ReadCloser, error) {
			return io.NopCloser(r), nil
		},
	}
}

func (s TemplateSource) Read(ctx context.Context) ([]api.TemplateEntry, error) {
	var entries []api.TemplateEntry
	if err := func() error {
		in, err := s.Open(ctx)
		if err != nil {
			return err
		}
		defer in.Close()

		switch s.Encoding {
		case TemplateEncodingJsonnet:
			err = jsonnet.NewDecoder(in).Decode(&entries)
		case TemplateEncodingJson:
			err = json.NewDecoder(in).Decode(&entries)
		case TemplateEncodingUnknown:
			fallthrough // Assume YAML if we can't make a better guess
		case TemplateEncodingYaml:
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
		return nil, fmt.Errorf("failed to read template(s) from %q: %w", s.Name, err)
	}

	return entries, nil
}

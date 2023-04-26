package config

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/influxdata/influx-cli/v2/api"
)

// store is the embedded store of the Config service.
type store interface {
	parsePreviousActive() (Config, error)
	ListConfigs() (Configs, error)
	writeConfigs(cfgs Configs) error
}

// localConfigsSVC can write and parse configs from a local path.
type localConfigsSVC struct {
	store
}

// NewLocalConfigService creates a new service that can write and parse configs
// to/from a path on local disk.
func NewLocalConfigService(path string) Service {
	return &localConfigsSVC{ioStore{Path: path}}
}

// CreateConfig create new config.
func (svc localConfigsSVC) CreateConfig(cfg Config) (Config, error) {
	if cfg.Name == "" {
		return Config{}, &api.Error{
			Code:    api.ERRORCODE_INVALID,
			Message: api.PtrString("config name is empty"),
		}
	}
	cfgs, err := svc.ListConfigs()
	if err != nil {
		return Config{}, err
	}
	if _, ok := cfgs[cfg.Name]; ok {
		return Config{}, &api.Error{
			Code:    api.ERRORCODE_CONFLICT,
			Message: api.PtrString(fmt.Sprintf("config %q already exists", cfg.Name)),
		}
	}
	cfgs[cfg.Name] = cfg
	if cfg.Active {
		if err := cfgs.switchActive(cfg.Name); err != nil {
			return Config{}, err
		}
	}

	return cfgs[cfg.Name], svc.writeConfigs(cfgs)
}

// DeleteConfig will delete a config.
func (svc localConfigsSVC) DeleteConfig(name string) (Config, error) {
	cfgs, err := svc.ListConfigs()
	if err != nil {
		return Config{}, err
	}

	p, ok := cfgs[name]
	if !ok {
		return Config{}, &api.Error{
			Code:    api.ERRORCODE_NOT_FOUND,
			Message: api.PtrString(fmt.Sprintf("config %q is not found", name)),
		}
	}
	delete(cfgs, name)

	if p.Active && len(cfgs) > 0 {
		for name, cfg := range cfgs {
			cfg.Active = true
			cfgs[name] = cfg
			break
		}
	}

	return p, svc.writeConfigs(cfgs)
}

func (svc localConfigsSVC) Active() (Config, error) {
	cfgs, err := svc.ListConfigs()
	if err != nil {
		return Config{}, err
	}
	return cfgs.active(), nil
}

// SwitchActive will active the config by name, if name is "-", active the previous one.
func (svc localConfigsSVC) SwitchActive(name string) (Config, error) {
	var up Config
	if name == "-" {
		p0, err := svc.parsePreviousActive()
		if err != nil {
			return Config{}, err
		}
		up.Name = p0.Name
	} else {
		up.Name = name
	}
	up.Active = true
	return svc.UpdateConfig(up)
}

// UpdateConfig will update the config.
func (svc localConfigsSVC) UpdateConfig(up Config) (Config, error) {
	cfgs, err := svc.ListConfigs()
	if err != nil {
		return Config{}, err
	}
	p0, ok := cfgs[up.Name]
	if !ok {
		return Config{}, &api.Error{
			Code:    api.ERRORCODE_NOT_FOUND,
			Message: api.PtrString(fmt.Sprintf("config %q is not found", up.Name)),
		}
	}
	if up.Token != "" {
		p0.Token = up.Token
		p0.Cookie = ""
	}
	if up.Cookie != "" {
		p0.Token = ""
		p0.Cookie = up.Cookie
	}
	if up.Host != "" {
		p0.Host = up.Host
	}
	if up.Org != "" {
		p0.Org = up.Org
	}

	cfgs[up.Name] = p0
	if up.Active {
		if err := cfgs.switchActive(up.Name); err != nil {
			return Config{}, err
		}
	}

	return cfgs[up.Name], svc.writeConfigs(cfgs)
}

type baseRW struct {
	r io.Reader
	w io.Writer
}

// parsePreviousActive return the previous active config from the reader
func (s baseRW) parsePreviousActive() (Config, error) {
	return s.parseActiveConfig(false)
}

// ListConfigs decodes configs from io readers
func (s baseRW) ListConfigs() (Configs, error) {
	cfgs := make(Configs)
	_, err := toml.NewDecoder(s.r).Decode(&cfgs)
	for n, cfg := range cfgs {
		cfg.Name = n
		cfgs[n] = cfg
	}
	return cfgs, err
}

func (s baseRW) writeConfigs(cfgs Configs) error {
	if err := blockBadName(cfgs); err != nil {
		return err
	}
	var b2 bytes.Buffer
	if err := toml.NewEncoder(s.w).Encode(cfgs); err != nil {
		return err
	}
	// a list cloud 2 clusters, commented out
	s.w.Write([]byte("# \n"))
	cfgs = map[string]Config{
		"us-central": {Host: "https://us-central1-1.gcp.cloud2.influxdata.com", Token: "XXX"},
		"us-west":    {Host: "https://us-west-2-1.aws.cloud2.influxdata.com", Token: "XXX"},
		"eu-central": {Host: "https://eu-central-1-1.aws.cloud2.influxdata.com", Token: "XXX"},
	}

	if err := toml.NewEncoder(&b2).Encode(cfgs); err != nil {
		return err
	}
	reader := bufio.NewReader(&b2)
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}
		s.w.Write([]byte("# " + string(line) + "\n"))
	}
	return nil
}

var badNames = map[string]bool{
	"-":      false,
	"list":   false,
	"update": false,
	"set":    false,
	"delete": false,
	"switch": false,
	"create": false,
}

func blockBadName(cfgs Configs) error {
	for n := range cfgs {
		if _, ok := badNames[n]; ok {
			return &api.Error{
				Code:    api.ERRORCODE_INVALID,
				Message: api.PtrString(fmt.Sprintf("%q is not a valid config name", n)),
			}
		}
	}
	return nil
}

func (s baseRW) parseActiveConfig(currentOrPrevious bool) (Config, error) {
	previousText := ""
	if !currentOrPrevious {
		previousText = "previous "
	}
	cfgs, err := s.ListConfigs()
	if err != nil {
		return DefaultConfig, err
	}
	var activated Config
	var hasActive bool
	for _, cfg := range cfgs {
		check := cfg.Active
		if !currentOrPrevious {
			check = cfg.PreviousActive
		}
		if check && !hasActive {
			activated = cfg
			hasActive = true
		} else if check {
			return DefaultConfig, &api.Error{
				Code:    api.ERRORCODE_CONFLICT,
				Message: api.PtrString(fmt.Sprintf("more than one %s activated configs found", previousText)),
			}
		}
	}
	if hasActive {
		return activated, nil
	}
	return DefaultConfig, &api.Error{
		Code:    api.ERRORCODE_NOT_FOUND,
		Message: api.PtrString(fmt.Sprintf("%s activated config is not found", previousText)),
	}
}

type ioStore struct {
	Path string
}

// ListConfigs from the local path.
func (s ioStore) ListConfigs() (Configs, error) {
	r, err := os.Open(s.Path)
	if err != nil {
		return make(Configs), nil
	}
	defer r.Close()
	return (baseRW{r: r}).ListConfigs()
}

// parsePreviousActive from the local path.
func (s ioStore) parsePreviousActive() (Config, error) {
	r, err := os.Open(s.Path)
	if err != nil {
		return Config{}, nil
	}
	defer r.Close()
	return (baseRW{r: r}).parsePreviousActive()
}

// writeConfigs to the path.
func (s ioStore) writeConfigs(cfgs Configs) error {
	if err := os.MkdirAll(filepath.Dir(s.Path), os.ModePerm); err != nil {
		return err
	}
	var b1 bytes.Buffer
	if err := (baseRW{w: &b1}).writeConfigs(cfgs); err != nil {
		return err
	}
	return os.WriteFile(s.Path, b1.Bytes(), 0600)
}

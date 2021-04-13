package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/influxdata/influx-cli/v2/api"
)

// Config store the crendentials of influxdb host and token.
type Config struct {
	Name string `toml:"-" json:"-"`
	Host string `toml:"url" json:"url"`
	// Token is base64 encoded sequence.
	Token          string `toml:"token" json:"token"`
	Org            string `toml:"org" json:"org"`
	Active         bool   `toml:"active,omitempty" json:"active,omitempty"`
	PreviousActive bool   `toml:"previous,omitempty" json:"previous,omitempty"`
}

// DefaultConfig is default config without token
var DefaultConfig = Config{
	Name:   "default",
	Host:   "http://localhost:8086",
	Active: true,
}

// DefaultPath computes the path where CLI configs will be stored if not overridden.
func DefaultPath() (string, error) {
	var dir string
	// By default, store meta and data files in current users home directory
	u, err := user.Current()
	if err == nil {
		dir = u.HomeDir
	} else if home := os.Getenv("HOME"); home != "" {
		dir = home
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		dir = wd
	}
	dir = filepath.Join(dir, ".influxdbv2", "config")

	return dir, nil
}

// Service is the service to list and write configs.
type Service interface {
	CreateConfig(Config) (Config, error)
	DeleteConfig(name string) (Config, error)
	UpdateConfig(Config) (Config, error)
	SwitchActive(name string) (Config, error)
	ListConfigs() (Configs, error)
}

// Configs is map of configs indexed by name.
type Configs map[string]Config

func GetConfigsOrDefault(path string) Configs {
	r, err := os.Open(path)
	if err != nil {
		return Configs{
			DefaultConfig.Name: DefaultConfig,
		}
	}
	defer r.Close()

	cfgs, err := NewLocalConfigService(path).ListConfigs()
	if err != nil {
		return Configs{
			DefaultConfig.Name: DefaultConfig,
		}
	}

	return cfgs
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
				Code:    api.ErrorCode_invalid,
				Message: fmt.Sprintf("%q is not a valid config name", n),
			}
		}
	}
	return nil
}

// Switch to another config.
func (cfgs Configs) Switch(name string) error {
	if _, ok := cfgs[name]; !ok {
		return &api.Error{
			Code:    api.ErrorCode_not_found,
			Message: fmt.Sprintf("config %q is not found", name),
		}
	}
	for k, v := range cfgs {
		v.PreviousActive = v.Active && (k != name)
		v.Active = k == name
		cfgs[k] = v
	}
	return nil
}

func (cfgs Configs) Active() Config {
	for _, cfg := range cfgs {
		if cfg.Active {
			return cfg
		}
	}
	if len(cfgs) > 0 {
		for _, cfg := range cfgs {
			return cfg
		}
	}
	return DefaultConfig
}

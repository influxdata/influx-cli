package config

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/internal/config"
)

var ErrInvalidHostUrlScheme = errors.New("a scheme of http or https must be provided for host url")

type Client struct {
	clients.CLI
}

func (c Client) SwitchActive(name string) error {
	cfg, err := c.ConfigService.SwitchActive(name)
	if err != nil {
		return err
	}
	return c.printConfigs(configPrintOpts{config: &cfg})
}

func (c Client) PrintActive() error {
	active, err := c.CLI.ConfigService.Active()
	if err != nil {
		return err
	}
	return c.printConfigs(configPrintOpts{config: &active})
}

func (c Client) Create(cfg config.Config) error {
	name := cfg.Name
	validated, err := validateHostUrl(cfg.Host)
	if err != nil {
		return fmt.Errorf("host URL %q is invalid: %w", cfg.Host, err)
	}
	cfg.Host = validated

	cfg, err = c.ConfigService.CreateConfig(cfg)
	if err != nil {
		return fmt.Errorf("failed to create config %q: %w", name, err)
	}
	return c.printConfigs(configPrintOpts{config: &cfg})
}

func (c Client) Delete(names []string) error {
	deleted := make(config.Configs)
	for _, name := range names {
		if name == "" {
			continue
		}
		cfg, err := c.ConfigService.DeleteConfig(name)
		if apiErr, ok := err.(*api.Error); ok && apiErr.Code == api.ERRORCODE_NOT_FOUND {
			continue
		} else if err != nil {
			return err
		}
		deleted[name] = cfg
	}
	return c.printConfigs(configPrintOpts{configs: deleted, deleted: true})
}

func (c Client) Update(cfg config.Config) error {
	name := cfg.Name
	if cfg.Host != "" {
		validated, err := validateHostUrl(cfg.Host)
		if err != nil {
			return fmt.Errorf("host URL %q is invalid: %w", cfg.Host, err)
		}
		cfg.Host = validated
	}

	cfg, err := c.ConfigService.UpdateConfig(cfg)
	if err != nil {
		return fmt.Errorf("failed to update config %q: %w", name, err)
	}
	return c.printConfigs(configPrintOpts{config: &cfg})
}

func (c Client) List() error {
	cfgs, err := c.ConfigService.ListConfigs()
	if err != nil {
		return err
	}
	return c.printConfigs(configPrintOpts{configs: cfgs})
}

type configPrintOpts struct {
	deleted bool
	config  *config.Config
	configs config.Configs
}

func (c Client) printConfigs(opts configPrintOpts) error {
	if c.PrintAsJSON {
		var v interface{}
		if opts.config != nil {
			v = opts.config
		} else {
			v = opts.configs
		}
		return c.PrintJSON(v)
	}

	headers := []string{"Active", "Name", "URL", "Org"}
	if opts.deleted {
		headers = append(headers, "Deleted")
	}

	if opts.config != nil {
		opts.configs = config.Configs{
			opts.config.Name: *opts.config,
		}
	}

	var rows []map[string]interface{}
	for _, c := range opts.configs {
		var active string
		if c.Active {
			active = "*"
		}
		row := map[string]interface{}{
			"Active": active,
			"Name":   c.Name,
			"URL":    c.Host,
			"Org":    c.Org,
		}
		if opts.deleted {
			row["Deleted"] = true
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

func validateHostUrl(hostUrl string) (string, error) {
	u, err := url.Parse(hostUrl)
	if err != nil {
		return "", err
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return "", ErrInvalidHostUrlScheme
	}
	return u.String(), nil
}

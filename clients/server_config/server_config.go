package server_config

import (
	"context"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"gopkg.in/yaml.v3"
)

type Client struct {
	clients.CLI
	api.ConfigApi
}

type ListParams struct {
	TOML bool
	YAML bool
}

func (c Client) List(ctx context.Context, params *ListParams) error {
	response, err := c.GetConfig(ctx).Execute()
	if err != nil {
		return fmt.Errorf("failed to retrieve config: %w", err)
	}

	config := response.GetConfig()

	// Any numeric values that were decoded as floats need to be converted to
	// integers. InfluxDB currently does not have any float config values, and
	// this prevents formatting issues with the output.
	for k, v := range config {
		if f, ok := v.(float64); ok {
			config[k] = int(f)
		}
	}

	if params.TOML {
		return toml.NewEncoder(c.StdIO).Encode(config)
	} else if params.YAML {
		return yaml.NewEncoder(c.StdIO).Encode(config)
	}

	return c.PrintJSON(config)
}

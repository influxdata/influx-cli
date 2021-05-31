package trino

import (
	v1 "github.com/influxdata/influx-cli/v2/api/v1"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	clients.CLI
	v1.QueryApi
}

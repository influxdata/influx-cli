package org

import (
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/internal/api"
)

type Client struct {
	clients.CLI
	api.OrganizationsApi
	api.UsersApi
}

type printOrgOpts struct {
	org     *api.Organization
	orgs    []api.Organization
	deleted bool
}

func (c Client) printOrgs(opts printOrgOpts) error {
	if c.PrintAsJSON {
		var v interface{}
		if opts.org != nil {
			v = opts.org
		} else {
			v = opts.orgs
		}
		return c.PrintJSON(v)
	}

	headers := []string{"ID", "Name"}
	if opts.deleted {
		headers = append(headers, "Deleted")
	}

	if opts.org != nil {
		opts.orgs = append(opts.orgs, *opts.org)
	}

	var rows []map[string]interface{}
	for _, o := range opts.orgs {
		row := map[string]interface{}{
			"ID":   o.GetId(),
			"Name": o.GetName(),
		}
		if opts.deleted {
			row["Deleted"] = true
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

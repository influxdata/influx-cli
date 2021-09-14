package remote

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	clients.CLI
	api.RemoteConnectionsApi
}

type CreateParams struct {
	Name             string
	Description      string
	OrgID            string
	RemoteURL        string
	RemoteAPIToken   string
	RemoteOrgID      string
	AllowInsecureTLS bool
}

func (c Client) Create(ctx context.Context, params *CreateParams) error {
	// set up a struct with required params
	body := api.RemoteConnectionCreationRequest{
		Name:             params.Name,
		OrgID:            params.OrgID,
		RemoteURL:        params.RemoteURL,
		RemoteAPIToken:   params.RemoteAPIToken,
		RemoteOrgID:      params.RemoteOrgID,
		AllowInsecureTLS: params.AllowInsecureTLS,
	}

	if params.Description != "" {
		body.Description = &params.Description
	}

	if body.OrgID == "" && c.ActiveConfig.Org == "" {
		return clients.ErrMustSpecifyOrg
	}

	if body.OrgID == "" {
		body.OrgID = c.ActiveConfig.Org
	}

	// send post request
	res, err := c.PostRemoteConnection(ctx).RemoteConnectionCreationRequest(body).Execute()
	if err != nil {
		return fmt.Errorf("failed to create remote connection %q: %w", params.Name, err)
	}
	// print confirmation of new connection
	return c.printRemote(printRemoteOpts{remote: &res})
}

type printRemoteOpts struct {
	remote  *api.RemoteConnection
	remotes []api.RemoteConnection
	deleted bool
}

func (c Client) printRemote(opts printRemoteOpts) error {
	if c.PrintAsJSON {
		var v interface{}
		if opts.remote != nil {
			v = opts.remote
		} else {
			v = opts.remotes
		}
		return c.PrintJSON(v)
	}

	headers := []string{"Remote ID", "Name"}
	if opts.deleted {
		headers = append(headers, "Deleted")
	}

	if opts.remote != nil {
		opts.remotes = append(opts.remotes, *opts.remote)
	}

	var rows []map[string]interface{}
	for _, r := range opts.remotes {
		row := map[string]interface{}{
			"Remote ID": r.GetId(),
			"Name":      r.GetName(),
		}
		if opts.deleted {
			row["Deleted"] = true
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

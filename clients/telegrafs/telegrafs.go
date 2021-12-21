package telegrafs

import (
	"context"
	"errors"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	clients.CLI
	api.TelegrafsApi
	api.OrganizationsApi
}

type telegrafPrintOpts struct {
	graf  *api.Telegraf
	grafs *api.Telegrafs
}

type ListParams struct {
	clients.OrgParams
	Id string
}

func (c Client) List(ctx context.Context, params *ListParams) error {
	if params.Id == "" && params.OrgID == "" && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return errors.New("at least one of org, org-id, or id must be provided")
	}

	if params.Id != "" {
		telegraf, err := c.GetTelegrafsID(ctx, params.Id).Execute()
		if err != nil {
			return fmt.Errorf("failed to get telegraf config with ID %q: %w", params.Id, err)
		}

		return c.printTelegrafs(telegrafPrintOpts{graf: &telegraf})
	}

	orgID, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	if err != nil {
		return err
	}

	telegrafs, err := c.GetTelegrafs(ctx).OrgID(orgID).Execute()
	if err != nil {
		return fmt.Errorf("failed to get telegraf configs from org %q: %w", orgID, err)
	}

	return c.printTelegrafs(telegrafPrintOpts{grafs: &telegrafs})
}

type CreateParams struct {
	clients.OrgParams
	Desc   string
	Config string
	Name   string
}

func (c Client) Create(ctx context.Context, params *CreateParams) error {
	orgID, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	if err != nil {
		return err
	}

	newTelegraf := api.TelegrafPluginRequest{
		Name:        &params.Name,
		Description: &params.Desc,
		Config:      &params.Config,
		OrgID:       &orgID,
	}

	graf, err := c.PostTelegrafs(ctx).TelegrafPluginRequest(newTelegraf).Execute()
	if err != nil {
		return fmt.Errorf("failed to create telegraf config %q: %w", params.Name, err)
	}

	return c.printTelegrafs(telegrafPrintOpts{graf: &graf})
}

type RemoveParams struct {
	Ids []string
}

func (c Client) Remove(ctx context.Context, params *RemoveParams) error {
	for _, rawID := range params.Ids {
		if err := c.DeleteTelegrafsID(ctx, rawID).Execute(); err != nil {
			return fmt.Errorf("failed to delete telegraf config with ID %q: %w", rawID, err)
		}
	}
	return nil
}

type UpdateParams struct {
	Desc   string
	Config string
	Name   string
	Id     string
}

func (c Client) Update(ctx context.Context, params *UpdateParams) error {
	oldTelegraf, err := c.GetTelegrafsID(ctx, params.Id).Execute()
	if err != nil {
		return fmt.Errorf("failed to find telegraf config with ID %q: %w", params.Id, err)
	}
	orgID := oldTelegraf.OrgID

	updateTelegraf := api.TelegrafRequest{
		Name:        &params.Name,
		Description: &params.Desc,
		Config:      &params.Config,
		OrgID:       orgID,
	}

	graf, err := c.PutTelegrafsID(ctx, params.Id).TelegrafRequest(updateTelegraf).Execute()
	if err != nil {
		return fmt.Errorf("failed to update telegraf config with ID %q: %w", params.Id, err)
	}

	return c.printTelegrafs(telegrafPrintOpts{graf: &graf})
}

func (c Client) printTelegrafs(opts telegrafPrintOpts) error {
	if c.PrintAsJSON {
		var v interface{}
		if opts.graf != nil {
			v = opts.graf
		} else {
			v = opts.grafs
		}
		return c.PrintJSON(v)
	}

	headers := []string{"ID", "OrgID", "Name", "Description"}

	telegrafs := opts.grafs.GetConfigurations()
	if opts.graf != nil {
		telegrafs = append(telegrafs, *opts.graf)
	}

	var rows []map[string]interface{}
	for _, u := range telegrafs {
		row := map[string]interface{}{
			"ID":          u.GetId(),
			"OrgID":       u.GetOrgID(),
			"Name":        u.GetName(),
			"Description": u.GetDescription(),
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

package telegrafs

import (
	"context"
	"errors"
	"io/ioutil"
	"os"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
)

type Client struct {
	clients.CLI
	api.TelegrafsApi
}

type TelegrafParams struct {
	clients.OrgParams

	Desc string
	File string
	Name string

	Id  string
	Ids []string
}

type telegrafPrintOpts struct {
	graf  *api.Telegraf
	grafs *api.Telegrafs
}

func (c Client) Telegrafs(ctx context.Context, params *TelegrafParams) error {
	if params.OrgID == 0 || params.Id == "" {
		return errors.New("at least one of org, org-id, or id must be provided")
	}

	if params.Id != "" {
		id, err := influxid.IDFromString(params.Id)
		if err != nil {
			return err
		}

		telegraf, err := c.GetTelegrafsID(ctx, id.String()).Execute()
		if err != nil {
			return err
		}

		return c.writeTelegrafConfigs(telegrafPrintOpts{graf: &telegraf})
	}

	telegrafs, err := c.GetTelegrafs(ctx).OrgID(params.OrgID.String()).Execute()
	if err != nil {
		return err
	}

	return c.writeTelegrafConfigs(telegrafPrintOpts{grafs: &telegrafs})
}

func (c Client) Create(ctx context.Context, params *TelegrafParams) error {
	cfg, err := c.readConfig(params.File)
	if err != nil {
		return err
	}

	orgID := params.OrgID.String()

	newTelegraf := api.TelegrafRequest{
		Name:        &params.Name,
		Description: &params.Desc,
		Config:      &cfg,
		OrgID:       &orgID,
	}

	graf, err := c.PostTelegrafs(ctx).TelegrafRequest(newTelegraf).Execute()
	if err != nil {
		return err
	}

	return c.writeTelegrafConfigs(telegrafPrintOpts{graf: &graf})
}

func (c Client) Remove(ctx context.Context, params *TelegrafParams) error {
	for _, rawID := range params.Ids {
		id, err := influxid.IDFromString(rawID)
		if err != nil {
			return err
		}

		if err = c.DeleteTelegrafsID(ctx, id.String()).Execute(); err != nil {
			return err
		}
	}
	return nil
}

func (c Client) Update(ctx context.Context, params *TelegrafParams) error {
	cfg, err := c.readConfig(params.File)
	if err != nil {
		return err
	}

	id, err := influxid.IDFromString(params.Id)
	if err != nil {
		return err
	}

	orgID := params.OrgID.String()
	updateTelegraf := api.TelegrafRequest{
		Name:        &params.Name,
		Description: &params.Desc,
		Config:      &cfg,
		OrgID:       &orgID,
	}

	graf, err := c.PutTelegrafsID(ctx, id.String()).TelegrafRequest(updateTelegraf).Execute()
	if err != nil {
		return err
	}

	return c.writeTelegrafConfigs(telegrafPrintOpts{graf: &graf})
}

func (c Client) writeTelegrafConfigs(opts telegrafPrintOpts) error {
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

func (c Client) readConfig(file string) (string, error) {
	if file != "" {
		bb, err := ioutil.ReadFile(file)
		if err != nil {
			return "", err
		}

		return string(bb), nil
	}

	bb, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}
	return string(bb), nil
}

package task

import (
	"context"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	clients.CLI
	api.TasksApi
	// AllowEmptyOrg will be useful for Kapacitor which doesn't use org / orgID
	AllowEmptyOrg bool
}

type CreateParams struct {
	clients.OrgParams
	FluxQuery string
}

func (c Client) getOrg(params *clients.OrgParams) (string, error) {
	if params.OrgID.Valid() {
		return params.OrgID.String(), nil
	}
	if params.OrgName != "" {
		return params.OrgName, nil
	}
	if c.ActiveConfig.Org != "" {
		return c.ActiveConfig.Org, nil
	}
	if c.AllowEmptyOrg {
		return "", nil
	}
	return "", clients.ErrMustSpecifyOrg
}

func (c Client) Create(ctx context.Context, params *CreateParams) error {
	org, err := c.getOrg(&params.OrgParams)
	if err != nil {
		return err
	}
	task, err := c.PostTasks(ctx).TaskCreateRequest(api.TaskCreateRequest{
		Org:  &org,
		Flux: params.FluxQuery,
	}).Execute()
	if err != nil {
		return err
	}
	return c.printTasks(taskPrintOpts{
		task: &task,
	})
}

type taskPrintOpts struct {
	task  *api.Task
	tasks []*api.Task
}

func (c Client) printTasks(printOpts taskPrintOpts) error {
	if c.PrintAsJSON {
		var v interface{} = printOpts.tasks
		if printOpts.task != nil {
			v = printOpts.task
		}
		return c.PrintJSON(v)
	}

	headers := []string{
		"ID",
		"Name",
		"Organization ID",
		"Organization",
		"Status",
		"Every",
		"Cron",
	}

	if printOpts.task != nil {
		printOpts.tasks = append(printOpts.tasks, printOpts.task)
	}

	var rows []map[string]interface{}
	for _, t := range printOpts.tasks {
		row := map[string]interface{}{
			"ID":              t.Id,
			"Name":            t.Name,
			"Organization ID": t.OrgID,
			"Organization":    t.Org,
			"Status":          t.Status,
			"Every":           t.Every,
			"Cron":            t.Cron,
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

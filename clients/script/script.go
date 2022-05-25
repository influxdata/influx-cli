package script

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/query"
)

type Client struct {
	clients.CLI
	api.InvocableScriptsApi
}

type printPayload struct {
	scripts []api.Script
}

func (c Client) printScripts(payload printPayload) error {
	if c.PrintAsJSON {
		return c.PrintJSON(payload.scripts)
	}

	headers := []string{
		"ID",
		"Name",
		"Description",
		"Organization ID",
		"Language",
	}

	var rows []map[string]interface{}
	for _, s := range payload.scripts {
		row := map[string]interface{}{
			"ID":              *s.Id,
			"Name":            s.Name,
			"Description":     *s.Description,
			"Organization ID": s.OrgID,
			"Language":        *s.Language,
		}

		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

type ListParams struct {
	Limit  int
	Offset int
}

func (c Client) List(ctx context.Context, params *ListParams) error {
	req := c.GetScripts(ctx)
	res, err := req.Execute()
	if err != nil {
		return fmt.Errorf("failed to list scripts: %q", err)
	}

	return c.printScripts(printPayload{
		scripts: *res.Scripts,
	})
}

type CreateParams struct {
	Description string
	Language    string
	Name        string
	Script      string
}

func (c Client) Create(ctx context.Context, params *CreateParams) error {
	req := api.ScriptCreateRequest{
		Name:        params.Name,
		Description: params.Description,
		Script:      params.Script,
		Language:    api.ScriptLanguage(params.Language),
	}

	script, err := c.PostScripts(ctx).ScriptCreateRequest(req).Execute()
	if err != nil {
		return fmt.Errorf("failed to create script: %q", err)
	}

	return c.printScripts(printPayload{
		scripts: []api.Script{script},
	})
}

type InvokeParams struct {
	ScriptID string
	Params   map[string]interface{}
}

func (c Client) Invoke(ctx context.Context, params *InvokeParams) error {
	req := api.ScriptInvocationParams{
		Params: &params.Params,
	}

	resp, err := c.PostScriptsIDInvoke(ctx, params.ScriptID).ScriptInvocationParams(req).Execute()
	if err != nil {
		return fmt.Errorf("failed to invoke script: %q", err)
	}
	defer resp.Body.Close()

	return query.RawResultPrinter.PrintQueryResults(resp.Body, c.StdIO)
}

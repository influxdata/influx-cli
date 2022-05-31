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

type DeleteParams struct {
	ScriptID string
}

func (c Client) Delete(ctx context.Context, params *DeleteParams) error {
	err := c.DeleteScriptsID(ctx, params.ScriptID).Execute()
	if err != nil {
		return fmt.Errorf("failed to delete script: %q", err)
	} else {
		return nil
	}
}

type RetrieveParams struct {
	ScriptID string
}

func (c Client) Retrieve(ctx context.Context, params *RetrieveParams) error {
	script, err := c.GetScriptsID(ctx, params.ScriptID).Execute()
	if err != nil {
		return fmt.Errorf("failed to retrieve script: %q", err)
	}

	return c.printScripts(printPayload{
		scripts: []api.Script{script},
	})
}

type UpdateParams struct {
	ScriptID    string
	Description string
	Name        string
	Script      string
}

func (c Client) Update(ctx context.Context, params *UpdateParams) error {
	// Retrieve the original since we might carry over some unchanged details.
	oldScript, err := c.GetScriptsID(ctx, params.ScriptID).Execute()
	if err != nil {
		return fmt.Errorf("failed to update script: %q", err)
	}

	if len(params.Description) == 0 {
		params.Description = *oldScript.Description
	}
	if len(params.Name) == 0 {
		params.Name = oldScript.Name
	}
	if len(params.Script) == 0 {
		params.Script = oldScript.Script
	}

	req := api.ScriptUpdateRequest{
		Name:        &params.Name,
		Description: &params.Description,
		Script:      &params.Script,
	}

	script, err := c.PatchScriptsID(ctx, params.ScriptID).ScriptUpdateRequest(req).Execute()
	if err != nil {
		return fmt.Errorf("failed to update script: %q", err)
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

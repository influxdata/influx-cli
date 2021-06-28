package invokable_script

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/query"
)

type Client struct {
	clients.CLI
	api.InvocableScriptsApi
	api.OrganizationsApi
	query.ResultPrinter
}

type scriptPrintOpts struct {
	script  *api.Script
	scripts []api.Script
	deleted bool
}

type ListParams struct {
	Limit  string
	Offset string
}

func (c Client) List(ctx context.Context, params *ListParams) error {
	req := c.GetScripts(ctx)

	if params.Limit != "" {
		l, err := strconv.Atoi(params.Limit)
		if err != nil {
			return fmt.Errorf("invalid value provided for limit: %w", err)
		}

		req = req.Limit(int32(l))
	}
	if params.Offset != "" {
		o, err := strconv.Atoi(params.Offset)
		if err != nil {
			return fmt.Errorf("invalid value provided for offset: %w", err)
		}

		req = req.Offset(int32(o))
	}

	res, err := req.Execute()
	if err != nil {
		return fmt.Errorf("failed to list scripts: %w", err)
	}

	return c.printScripts(scriptPrintOpts{scripts: *res.Scripts})
}

type CreateParams struct {
	Name        string
	Description string
	Script      string
	Language    string
}

func (c Client) Create(ctx context.Context, params *CreateParams) (err error) {
	reqBody := api.ScriptCreateRequest{
		Name:        params.Name,
		Script:      params.Script,
		Description: params.Description,
	}

	// Unmarshalling logic to handle the script type enum.
	var lang api.ScriptLanguage
	if err := lang.UnmarshalJSON([]byte(fmt.Sprintf("%q", params.Language))); err != nil {
		return err
	}
	reqBody.Language = lang

	script, err := c.PostScripts(ctx).ScriptCreateRequest(reqBody).Execute()
	if err != nil {
		return fmt.Errorf("failed to create invokable script: %w", err)
	}

	return c.printScripts(scriptPrintOpts{script: &script})
}

type GetParams struct {
	ID              string
	PrintScriptOnly bool
}

func (c Client) Get(ctx context.Context, params *GetParams) error {
	script, err := c.GetScriptsID(ctx, params.ID).Execute()
	if err != nil {
		return fmt.Errorf("failed to get invokable script %q: %w", params.ID, err)
	}

	if params.PrintScriptOnly {
		return c.printScriptOnly(script)
	}

	return c.printScripts(scriptPrintOpts{script: &script})
}

type DeleteParams struct {
	ID string
}

func (c Client) Delete(ctx context.Context, params *DeleteParams) error {
	script, err := c.GetScriptsID(ctx, params.ID).Execute()
	if err != nil {
		return fmt.Errorf("failed to delete invokable script %q: %w", params.ID, err)
	}

	err = c.DeleteScriptsID(ctx, params.ID).Execute()
	if err != nil {
		return fmt.Errorf("failed to delete invokable script %q: %w", params.ID, err)
	}

	return c.printScripts(scriptPrintOpts{script: &script, deleted: true})
}

type UpdateParams struct {
	ID          string
	Name        string
	Description string
	Script      string
}

func (c Client) Update(ctx context.Context, params *UpdateParams) error {
	reqBody := api.ScriptUpdateRequest{}

	if params.Name != "" {
		reqBody.Name = &params.Name
	}
	if params.Description != "" {
		reqBody.Description = &params.Description
	}
	if params.Script != "" {
		reqBody.Script = &params.Script
	}

	req := c.PatchScriptsID(ctx, params.ID)
	script, err := req.ScriptUpdateRequest(reqBody).Execute()
	if err != nil {
		return fmt.Errorf("failed to update invokable script %q: %w", params.ID, err)
	}

	return c.printScripts(scriptPrintOpts{script: &script})
}

type InvokeParams struct {
	ID           string
	ScriptParams io.ReadCloser
}

func (c Client) Invoke(ctx context.Context, params *InvokeParams) error {
	var reqBody api.ScriptInvocationParams
	sParams := make(map[string]interface{})
	if params.ScriptParams != nil {
		defer params.ScriptParams.Close()
		if err := json.NewDecoder(params.ScriptParams).Decode(&sParams); err != nil {
			return err
		}
	}

	// If the parameters JSON included the "params" key, drop it
	if _, ok := sParams["params"]; ok {
		v, ok := sParams["params"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("could not process parameters JSON")
		}
		sParams = v
	}

	reqBody.SetParams(sParams)
	req := c.PostScriptsIDInvoke(ctx, params.ID)
	resp, err := req.ScriptInvocationParams(reqBody).Execute()
	if err != nil {
		return fmt.Errorf("failed to invoke script %q: %w", params.ID, err)
	}

	return c.PrintQueryResults(io.NopCloser(strings.NewReader(resp)), c.StdIO)
}

// printScriptOnly prints the script string directly
func (c Client) printScriptOnly(s api.Script) error {
	_, err := io.Copy(c.StdIO, strings.NewReader(s.Script+"\n"))
	return err
}

func (c Client) printScripts(opts scriptPrintOpts) error {
	if opts.script != nil {
		opts.scripts = append(opts.scripts, *opts.script)
	}

	if c.PrintAsJSON {
		return c.PrintJSON(opts.scripts)
	}

	headers := []string{
		"ID",
		"Name",
		"Description",
		"Organization ID",
		"Language",
	}
	if opts.deleted {
		headers = append(headers, "Deleted")
	}

	var rows []map[string]interface{}
	for _, s := range opts.scripts {
		row := map[string]interface{}{
			"ID":              *s.Id,
			"Name":            s.Name,
			"Description":     *s.Description,
			"Organization ID": s.OrgID,
			"Language":        *s.Language,
		}
		if opts.deleted {
			row["Deleted"] = true
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

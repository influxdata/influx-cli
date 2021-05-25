package query

import (
	"context"
	"fmt"
	"io"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type ResultPrinter interface {
	PrintQueryResults(resultStream io.ReadCloser, out io.Writer) error
}

type rawResultPrinter struct{}

// RawResultPrinter streams query results directly to the output without
// any parsing or formatting.
var RawResultPrinter ResultPrinter = &rawResultPrinter{}

func (r *rawResultPrinter) PrintQueryResults(resultStream io.ReadCloser, out io.Writer) error {
	_, err := io.Copy(out, resultStream)
	return err
}

type Client struct {
	clients.CLI
	api.QueryApi
	ResultPrinter
}

type Params struct {
	clients.OrgParams
	Query     string
	Profilers []string
}

// BuildDefaultAST wraps a raw query string in the AST structure expected
// by the query API, injecting default values expected by the CLI formatter.
func BuildDefaultAST(query string) api.Query {
	return api.Query{
		Query: query,
		Type:  api.PtrString("flux"),
		Dialect: &api.Dialect{
			Annotations: &[]string{"group", "datatype", "default"},
			Delimiter:   api.PtrString(","),
			Header:      api.PtrBool(true),
		},
	}
}

// BuildExternAST constructs a Flux AST tree to import and set the profilers option.
//
// See the docs for more info: https://docs.influxdata.com/influxdb/cloud/reference/flux/stdlib/profiler/
func BuildExternAST(profilers []string) *api.Extern {
	// Construct AST statements to import and set the 'profilers' option.
	// NOTE: We've purposefully codegen'd a map[string]interface{} schema
	// for the field populated here because the API spec for our Flux AST
	// generates very hard-to-use models, and our attempts to change
	// that have all (so far) broken the codegen for the UI.
	//
	// We assume that this logic will be changed infrequently enough that
	// the lack of type-safety won't be a frequent pain point.

	// import "profiler"
	profilersImport := map[string]interface{}{
		"type": "ImportDeclaration",
		"path": map[string]interface{}{
			"type":  "StringLiteral",
			"value": "profiler",
		},
	}
	// "<profiler>" for each profiler
	profilerExprs := make([]interface{}, len(profilers))
	for i, profiler := range profilers {
		profilerExprs[i] = map[string]interface{}{
			"type":  "StringLiteral",
			"value": profiler,
		}
	}
	// ["<profiler>" for each profiler]
	profilersArrayExpr := map[string]interface{}{
		"type":     "ArrayExpression",
		"elements": profilerExprs,
	}
	// profiler.enabledProfilers
	profilersMemberExpr := map[string]interface{}{
		"type": "MemberExpression",
		"object": map[string]interface{}{
			"name": "profiler",
			"type": "Identifier",
		},
		"property": map[string]interface{}{
			"name": "enabledProfilers",
			"type": "Identifier",
		},
	}
	// profiler.enabledProfilers = ["<profiler>" for each profiler]
	profilersAssignmentExpr := map[string]interface{}{
		"type":   "MemberAssignment",
		"member": profilersMemberExpr,
		"init":   profilersArrayExpr,
	}
	// option profiler.enabledProfilers = ["<profiler>" for each profiler]
	profilersOptionExpr := map[string]interface{}{
		"type":       "OptionStatement",
		"assignment": profilersAssignmentExpr,
	}
	// import "profiler"
	// option profiler.enabledProfilers = ["<profiler>" for each profiler]
	profilersExternExpr := map[string]interface{}{
		"imports": []interface{}{profilersImport},
		"body":    []interface{}{profilersOptionExpr},
	}

	extern := api.NewExternWithDefaults()
	extern.AdditionalProperties = profilersExternExpr
	return extern
}

func (c Client) Query(ctx context.Context, params *Params) error {
	if !params.OrgID.Valid() && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return clients.ErrMustSpecifyOrg
	}

	query := BuildDefaultAST(params.Query)
	if len(params.Profilers) > 0 {
		query.Extern = BuildExternAST(params.Profilers)
	}

	req := c.PostQuery(ctx).Query(query).AcceptEncoding("gzip")
	if params.OrgID.Valid() {
		req = req.OrgID(params.OrgID.String())
	} else if params.OrgName != "" {
		req = req.Org(params.OrgName)
	} else {
		req = req.Org(c.ActiveConfig.Org)
	}

	resp, err := req.Execute()
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	defer resp.Close()

	return c.PrintQueryResults(resp, c.StdIO)
}

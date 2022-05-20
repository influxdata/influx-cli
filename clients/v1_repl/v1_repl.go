package v1repl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	clients.CLI
	PersistentQueryParams
	api.LegacyQueryApi
	api.PingApi
	api.OrganizationsApi
}

type PersistentQueryParams struct {
	clients.OrgParams
	Db     string // bucketID
	P      string // password OR token
	U      string
	Rp     string
	Epoch  string
	Format FormatType
	Pretty bool
}

func DefaultPersistentQueryParams() PersistentQueryParams {
	return PersistentQueryParams{
		Format: CsvFormat,
		Epoch:  "n",
	}
}

type FormatType string

var (
	CsvFormat  FormatType = "csv"
	JsonFormat FormatType = "json"
)

func (c *Client) Create(ctx context.Context) error {
	res, err := c.GetPing(ctx).ExecuteWithHttpInfo()
	if err != nil {
		color.Red("Unable to connect to InfluxDB")
		return err
	}
	build := res.Header.Get("X-Influxdb-Build")
	version := res.Header.Get("X-Influxdb-Version")
	color.Cyan("Connected to InfluxDB %s %s", build, version)
	p := prompt.New(c.executor,
		completer,
		prompt.OptionTitle("InfluxQL Shell"),
		prompt.OptionDescriptionTextColor(prompt.Cyan),
		prompt.OptionPrefixTextColor(prompt.Green),
	)
	p.Run()
	return nil
}

var AllInfluxQLKeywords []prompt.Suggest = []prompt.Suggest{
	{Text: "ALL"},
	{Text: "ALTER"},
	{Text: "ANY"},
	{Text: "AS"},
	{Text: "ASC"},
	{Text: "BEGIN"},
	{Text: "BY"},
	{Text: "CREATE"},
	{Text: "CONTINUOUS"},
	{Text: "DATABASE"},
	{Text: "DATABASES"},
	{Text: "DEFAULT"},
	{Text: "DELETE"},
	{Text: "DESC"},
	{Text: "DESTINATIONS"},
	{Text: "DIAGNOSTICS"},
	{Text: "DISTINCT"},
	{Text: "DROP"},
	{Text: "DURATION"},
	{Text: "END"},
	{Text: "EVERY"},
	{Text: "EXPLAIN"},
	{Text: "FIELD"},
	{Text: "FOR"},
	{Text: "FROM"},
	{Text: "GRANT"},
	{Text: "GRANTS"},
	{Text: "GROUP"},
	{Text: "GROUPS"},
	{Text: "IN"},
	{Text: "INF"},
	{Text: "INSERT"},
	{Text: "INTO"},
	{Text: "KEY"},
	{Text: "KEYS"},
	{Text: "KILL"},
	{Text: "LIMIT"},
	{Text: "SHOW"},
	{Text: "MEASUREMENT"},
	{Text: "MEASUREMENTS"},
	{Text: "NAME"},
	{Text: "OFFSET"},
	{Text: "ON"},
	{Text: "ORDER"},
	{Text: "PASSWORD"},
	{Text: "POLICY"},
	{Text: "POLICIES"},
	{Text: "PRIVILEGES"},
	{Text: "QUERIES"},
	{Text: "QUERY"},
	{Text: "READ"},
	{Text: "REPLICATION"},
	{Text: "RESAMPLE"},
	{Text: "RETENTION"},
	{Text: "REVOKE"},
	{Text: "SELECT"},
	{Text: "SERIES"},
	{Text: "SET"},
	{Text: "SHARD"},
	{Text: "SHARDS"},
	{Text: "SLIMIT"},
	{Text: "SOFFSET"},
	{Text: "STATS"},
	{Text: "SUBSCRIPTION"},
	{Text: "SUBSCRIPTIONS"},
	{Text: "TAG"},
	{Text: "TO"},
	{Text: "USER"},
	{Text: "USERS"},
	{Text: "VALUES"},
	{Text: "WHERE"},
	{Text: "WITH"},
	{Text: "WRITE"},
}

var ReplKeywords []prompt.Suggest = []prompt.Suggest{
	// {Text: "connect", Description: "Connect to another node"},
	// {Text: "auth", Description: "Prompt for username and password"},
	{Text: "pretty", Description: "Toggle pretty print for the json format"},
	{Text: "use", Description: "Set current database"},
	{Text: "precision", Description: "Specify the format of the timestamp"},
	// {Text: "history", Description: "Display shell history"},
	// {Text: "settings", Description: "Output the current shell settings"},
	// {Text: "clear", Description: "Clears settings such as database"},
	{Text: "exit", Description: "Exit the InfluxQL shell"},
	{Text: "quit", Description: "Exit the InfluxQL shell"},
	// {Text: "gopher", Description: "Display the Go Gopher"},
	{Text: "help", Description: "Display help options"},
	{Text: "format", Description: "Specify the data display format"},
}

func (c *Client) executor(cmd string) {
	if cmd == "" {
		return
	}
	cmdArgs := strings.Split(cmd, " ")
	switch cmdArgs[0] {
	case "quit", "exit":
		color.HiBlack("Goodbye!")
		os.Exit(0)
	// case "gopher":
	// 	c.gopher()
	// case "connect":
	// 	return c.Connect(cmd)
	// case "auth":
	// 	c.SetAuth(cmd)
	// case "help":
	// 	c.help()
	// case "history":
	//  c.History()
	case "format":
		c.SetFormat(cmdArgs)
	// case "precision":
	// 	c.SetPrecision(cmd)
	// case "consistency":
	// 	c.SetWriteConsistency(cmd)
	// case "settings":
	// 	c.Settings()
	case "pretty":
		c.TogglePretty()
	// case "use":
	// 	c.use(cmd)
	// case "node":
	// 	c.node(cmd)
	// case "insert":
	// 	return c.Insert(cmd)
	// case "clear":
	// 	c.clear(cmd)
	default:
		c.RunAndShowQuery(cmd)
	}
}

func (c *Client) RunAndShowQuery(query string) {
	response, err := c.Query(context.Background(), query)
	if err != nil {
		color.HiRed("Query failed.")
		color.Red("%v", err)
		return
	}
	displayMap := map[FormatType]func(string){
		CsvFormat:  c.OutputCsv,
		JsonFormat: c.OutputJson,
	}
	displayFunc := displayMap[c.Format]
	displayFunc(response)
}

func completer(d prompt.Document) []prompt.Suggest {
	return append(
		prompt.FilterHasPrefix(ReplKeywords, d.CurrentLine(), false),
		prompt.FilterHasPrefix(AllInfluxQLKeywords, d.GetWordBeforeCursor(), true)...,
	)
}

func (c *Client) Query(ctx context.Context, query string) (string, error) {
	var resContentType string
	switch c.Format {
	case CsvFormat:
		resContentType = "application/csv"
	case JsonFormat:
		resContentType = "application/json"
	default:
		return "", fmt.Errorf("unexpected format: %s", c.Format)
	}
	resBody, err := c.GetLegacyQuery(ctx).
		U(c.U).
		P(c.P).
		Db(c.Db).
		Q(query).
		Rp(c.Rp).
		Epoch(c.Epoch).
		Accept(resContentType).
		Execute()
	if err != nil {
		return "", err
	}
	return resBody, nil
}

func (c *Client) OutputCsv(csvBody string) {
	fmt.Println(csvBody)
}

func (c *Client) OutputJson(jsonBody string) {
	if !c.Pretty {
		fmt.Println(jsonBody)
	} else {
		var buf bytes.Buffer
		if err := json.Indent(&buf, []byte(jsonBody), "", "  "); err != nil {
			color.Red("Unable to prettify json response.")
			fmt.Println(jsonBody)
		} else {
			fmt.Println(buf.String())
		}
	}
}

// -- Command Helper Functions --

func (c *Client) SetFormat(args []string) {
	// args[0] is "format"
	if len(args) != 2 {
		color.Red("Expected a format, like csv or json")
		return
	}
	newFormat := FormatType(args[1])
	switch newFormat {
	case CsvFormat, JsonFormat:
		c.Format = newFormat
	default:
		color.HiRed("Unimplemented format %q, keeping %s format.", newFormat, c.Format)
	}
}

func (c *Client) TogglePretty() {
	c.Pretty = !c.Pretty
}

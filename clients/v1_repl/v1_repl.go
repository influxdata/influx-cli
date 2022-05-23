package v1repl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/influxdata/go-prompt"
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
	// Storage
	Databases         []string
	RetentionPolicies []string
	Measurements      []string
}

func DefaultPersistentQueryParams() PersistentQueryParams {
	return PersistentQueryParams{
		Format: CsvFormat,
		Epoch:  "n",
	}
}

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
		c.completer,
		prompt.OptionTitle("InfluxQL Shell"),
		prompt.OptionDescriptionTextColor(prompt.Cyan),
		prompt.OptionPrefixTextColor(prompt.Green),
		prompt.OptionCompletionWordSeparator(" ", "."),
	)
	c.Databases, _ = c.GetDatabases(ctx)
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
	// {Text: "precision", Description: "Specify the format of the timestamp"},
	// {Text: "history", Description: "Display shell history"},
	// {Text: "settings", Description: "Output the current shell settings"},
	// {Text: "clear", Description: "Clears settings such as database"},
	{Text: "exit", Description: "Exit the InfluxQL shell"},
	{Text: "quit", Description: "Exit the InfluxQL shell"},
	// {Text: "gopher", Description: "Display the Go Gopher"},
	// {Text: "help", Description: "Display help options"},
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
	case "use":
		c.use(cmdArgs)
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

type FormatType string

var (
	CsvFormat   FormatType = "csv"
	JsonFormat  FormatType = "json"
	TableFormat FormatType = "table"
)

func (c *Client) RunAndShowQuery(query string) {
	response, err := c.Query(context.Background(), query)
	if err != nil {
		color.HiRed("Query failed.")
		color.Red("%v", err)
		return
	}
	displayMap := map[FormatType]func(string){
		CsvFormat:   c.OutputCsv,
		JsonFormat:  c.OutputJson,
		TableFormat: c.OutputTable,
	}
	displayFunc := displayMap[c.Format]
	displayFunc(response)
}

var IdentifierRegex = `([0-9A-Za-z\"\-\_]+)`

func (c *Client) completer(d prompt.Document) []prompt.Suggest {
	currentLineUpper := strings.ToUpper(d.CurrentLine())
	var s []prompt.Suggest
	if strings.HasPrefix(d.CurrentLine(), "format ") {
		s = append(s, prompt.Suggest{Text: "table", Description: "Format Type"})
		s = append(s, prompt.Suggest{Text: "json", Description: "Format Type"})
		s = append(s, prompt.Suggest{Text: "csv", Description: "Format Type"})
		return prompt.FilterFuzzy(s, d.GetWordBeforeCursor(), true)
	} else if strings.HasPrefix(d.CurrentLine(), "use ") || strings.HasPrefix(d.CurrentLine(), "use \"") {
		for _, db := range c.Databases {
			s = append(s, prompt.Suggest{Text: "\"" + db + "\"", Description: "Table Name"})
		}
		return prompt.FilterFuzzy(s, d.GetWordBeforeCursor(), true)
	} else if strings.HasPrefix(currentLineUpper, "SELECT ") {
		if isMatch, _ := regexp.Match(`FROM `+IdentifierRegex+`?$`, []byte(currentLineUpper)); isMatch {
			if c.Db != "" && c.Rp != "" {
				for _, m := range c.Measurements {
					s = append(s, prompt.Suggest{Text: "\"" + m + "\"", Description: fmt.Sprintf("Measurement on \"%s\".\"%s\"", c.Db, c.Rp)})
				}
			}
			if c.Db != "" {
				for _, rp := range c.RetentionPolicies {
					s = append(s, prompt.Suggest{Text: "\"" + rp + "\"", Description: "Retention Policy on " + c.Db})
				}
			}
			for _, db := range c.Databases {
				s = append(s, prompt.Suggest{Text: "\"" + db + "\"", Description: "Table Name"})
			}
			return prompt.FilterFuzzy(s, d.GetWordBeforeCursorUntilAnySeparator(" ", "."), true)
		}
	}
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
	case JsonFormat, TableFormat:
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

func (c *Client) SetFormat(args []string) {
	// args[0] is "format"
	if len(args) != 2 {
		color.Red("Expected a format, like csv or json")
		return
	}
	newFormat := FormatType(args[1])
	switch newFormat {
	case CsvFormat, JsonFormat, TableFormat:
		c.Format = newFormat
	default:
		color.HiRed("Unimplemented format %q, keeping %s format.", newFormat, c.Format)
	}
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

func (c *Client) OutputTable(jsonBody string) {
	var responses api.InfluxqlJsonResponse
	if err := json.Unmarshal([]byte(jsonBody), &responses); err != nil {
		color.Red("Failed to parse JSON response")
	}
	for _, res := range responses.GetResults() {
		for _, series := range res.GetSeries() {
			color.Magenta("Table View (press q to exit interactive mode):")
			p := tea.NewProgram(NewModel(series))
			if err := p.Start(); err != nil {
				color.Red("Failed to display table")
			}
			fmt.Printf("\n")
		}
	}
}

func (c *Client) TogglePretty() {
	c.Pretty = !c.Pretty
}

func (c *Client) use(args []string) {
	if len(args) != 2 {
		color.Red("wrong number of args for \"use [DATABASE_NAME]\"")
		return
	}
	parsedDb, parsedRp, err := parseDatabaseAndRetentionPolicy([]byte(args[1]))
	if err != nil {
		color.Red("Unable to parse: %v", err)
		return
	}
	dbs, err := c.GetDatabases(context.Background())
	if err != nil {
		color.Red("Unable to check databases: %v", err)
		return
	}
	for _, db := range dbs {
		if parsedDb == db {
			exists := false
			prevDb := c.Db
			c.Db = parsedDb
			rps, _ := c.GetRetentionPolicies(context.Background())
			for _, rp := range rps {
				if parsedRp == rp || parsedRp == "" {
					if parsedRp == "" {
						c.Rp, _ = c.GetDefaultRetentionPolicy(context.Background())
					} else {
						c.Rp = parsedRp
					}
					c.RetentionPolicies = rps
					exists = true
					c.Measurements, _ = c.GetMeasurements(context.Background())
					break
				}
			}
			if !exists {
				color.Red("No such retention policy %q exists on %q", parsedRp, c.Db)
				color.HiBlack("Available retention policies on %q:", parsedDb)
				for _, rp := range rps {
					color.HiBlack("- %q", rp)
				}
				c.Db = prevDb
				return
			}
			c.Db = parsedDb
			c.Databases = dbs

			return
		}
	}
	color.Red("No such database %q exists", parsedDb)
	color.HiBlack("Available databases:")
	for _, db := range dbs {
		color.HiBlack("- %q", db)
	}

}

func (c *Client) GetRetentionPolicies(ctx context.Context) ([]string, error) {
	singleSeries, err := c.getDataSingleSeries(ctx,
		fmt.Sprintf("SHOW RETENTION POLICIES ON %q", c.Db))
	if err != nil {
		return []string{}, err
	}
	nameIndex := -1
	for i, colName := range singleSeries.GetColumns() {
		if colName == "name" {
			nameIndex = i
		}
	}
	if nameIndex == -1 {
		return []string{}, fmt.Errorf("expected a \"name\" column for retention policies")
	}
	var retentionPolicies []string
	for _, value := range singleSeries.GetValues() {
		if name, ok := value[nameIndex].(string); ok {
			retentionPolicies = append(retentionPolicies, name)
		} else {
			return []string{}, fmt.Errorf("expected \"name\" column to contain string value")
		}
	}
	return retentionPolicies, nil
}

func (c *Client) GetDefaultRetentionPolicy(ctx context.Context) (string, error) {
	singleSeries, err := c.getDataSingleSeries(ctx,
		fmt.Sprintf("SHOW RETENTION POLICIES ON %q", c.Db))
	if err != nil {
		return "", err
	}
	nameIndex := -1
	defaultIndex := -1
	for i, colName := range singleSeries.GetColumns() {
		if colName == "default" {
			defaultIndex = i
		} else if colName == "name" {
			nameIndex = i
		}
	}
	if nameIndex == -1 {
		return "", fmt.Errorf("expected a \"name\" column for retention policies")
	}
	if defaultIndex == -1 {
		return "", fmt.Errorf("expected a \"default\" column for retention policies")
	}
	for _, value := range singleSeries.GetValues() {
		isDefault := value[defaultIndex]
		if isDefault, ok := isDefault.(bool); ok {
			if isDefault {
				if name, ok := value[nameIndex].(string); ok {
					return name, nil
				} else {
					return "", fmt.Errorf("expected \"name\" column to contain string value")
				}
			}
		} else {
			return "", fmt.Errorf("expected \"default\" column to contain boolean value")
		}
	}
	return "", fmt.Errorf("no default retention policy")
}

func (c *Client) GetDatabases(ctx context.Context) ([]string, error) {
	singleSeries, err := c.getDataSingleSeries(ctx, "SHOW DATABASES")
	if err != nil {
		return []string{}, err
	}
	values := singleSeries.GetValues()
	if len(values) != 1 {
		return []string{}, fmt.Errorf("expected a single array in values array")
	}
	var databases []string
	for _, db := range values[0] {
		if db, ok := db.(string); ok {
			databases = append(databases, db)
		} else {
			return []string{}, fmt.Errorf("expected database names to be strings")
		}
	}
	return databases, nil
}

func (c *Client) GetMeasurements(ctx context.Context) ([]string, error) {
	singleSeries, err := c.getDataSingleSeries(ctx, "SHOW MEASUREMENTS")
	if err != nil {
		return []string{}, err
	}
	var measures []string
	for _, measureArr := range singleSeries.GetValues() {
		if len(measureArr) != 1 {
			return []string{}, fmt.Errorf("expected a single measurement name in each array in values array")
		}
		if measure, ok := measureArr[0].(string); ok {
			measures = append(measures, measure)
		} else {
			return []string{}, fmt.Errorf("expected measurement name to be a string")
		}
	}
	return measures, nil
}

func (c *Client) getDataSingleSeries(ctx context.Context, query string) (*api.InfluxqlJsonResponseSeries, error) {
	resBody, err := c.GetLegacyQuery(ctx).
		U(c.U).
		P(c.P).
		Db(c.Db).
		Q(query).
		Rp(c.Rp).
		Epoch(c.Epoch).
		Accept("application/json").
		Execute()
	if err != nil {
		return nil, err
	}
	var responses api.InfluxqlJsonResponse
	if err := json.Unmarshal([]byte(resBody), &responses); err != nil {
		return nil, err
	}
	results := responses.GetResults()
	if len(results) != 1 {
		return nil, fmt.Errorf("expected a single result from single query")
	}
	result := results[0]
	series := result.GetSeries()
	if len(series) != 1 {
		return nil, fmt.Errorf("expected a single series from single result")
	}
	return &series[0], nil
}

func parseDatabaseAndRetentionPolicy(stmt []byte) (string, string, error) {
	var db, rp []byte
	var quoted bool
	var seperatorCount int

	stmt = bytes.TrimSpace(stmt)

	for _, b := range stmt {
		if b == '"' {
			quoted = !quoted
			continue
		}
		if b == '.' && !quoted {
			seperatorCount++
			if seperatorCount > 1 {
				return "", "", fmt.Errorf("unable to parse database and retention policy from %s", string(stmt))
			}
			continue
		}
		if seperatorCount == 1 {
			rp = append(rp, b)
			continue
		}
		db = append(db, b)
	}
	return string(db), string(rp), nil
}

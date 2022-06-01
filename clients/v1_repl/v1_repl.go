package v1repl

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"text/tabwriter"

	tea "github.com/charmbracelet/bubbletea"
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
	api.WriteApi
	api.DBRPsApi
}

type PersistentQueryParams struct {
	clients.OrgParams
	Database        string
	Password        string
	Username        string
	RetentionPolicy string
	Precision       string
	Format          FormatType
	Pretty          bool
	historyFilePath string
	// Autocompletion Storage
	Databases         []string
	RetentionPolicies []string
	Measurements      []string
}

func (c *Client) clear(cmd string) {
	args := strings.Split(strings.TrimSuffix(strings.TrimSpace(cmd), ";"), " ")
	v := strings.ToLower(strings.Join(args[1:], " "))
	switch v {
	case "database", "db":
		c.Database = ""
		c.RetentionPolicies = []string{}
		fmt.Println("database context cleared")
		return
	case "retention policy", "rp":
		c.RetentionPolicy = ""
		fmt.Println("retention policy context cleared")
		return
	default:
		if len(args) > 1 {
			fmt.Printf("invalid command %q.\n", v)
		}
		fmt.Println(`Possible commands for 'clear' are:
    # Clear the database context
    clear database
    clear db

    # Clear the retention policy context
    clear retention policy
    clear rp
		`)
	}
}

func DefaultPersistentQueryParams() PersistentQueryParams {
	return PersistentQueryParams{
		Format:    CsvFormat,
		Precision: "ns",
	}
}

func (c *Client) readHistory() []string {
	// Attempt to load the history file.
	if c.historyFilePath != "" {
		if historyFile, err := os.Open(c.historyFilePath); err == nil {
			var history []string
			scanner := bufio.NewScanner(historyFile)
			for scanner.Scan() {
				history = append(history, scanner.Text())
			}
			historyFile.Close()
			// Limit to last 100 elements
			historyElems := 100
			if len(history) > historyElems {
				history = history[len(history)-historyElems:]
			}
			return history
		}
	}
	return []string{}
}

func (c *Client) writeCommandToHistory(cmd string) {
	if c.historyFilePath != "" {
		if historyFile, err := os.OpenFile(c.historyFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
			historyFile.WriteString(cmd + "\n")
			historyFile.Close()
		}
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

	// compute historyFilePath at REPL start
	// Only load/write history if HOME environment variable is set.
	var historyDir string
	if runtime.GOOS == "windows" {
		if userDir := os.Getenv("USERPROFILE"); userDir != "" {
			historyDir = userDir
		}
	}
	if homeDir := os.Getenv("HOME"); homeDir != "" {
		historyDir = homeDir
	}
	if historyDir != "" {
		c.historyFilePath = filepath.Join(historyDir, ".influx_history")
	}

	c.Databases, _ = c.GetDatabases(ctx)
	color.Cyan("InfluxQL Shell")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s", color.GreenString("> "))
	for scanner.Scan() {
		c.executor(scanner.Text())
		fmt.Printf("%s", color.GreenString("> "))
	}
	return nil
}

func (c *Client) gopher() {
	color.Cyan(Gopher)
}

// The logic for the main prompt that is run in the REPL loop
func (c *Client) executor(cmd string) {
	if cmd == "" {
		return
	}
	defer c.writeCommandToHistory(cmd)
	cmdArgs := strings.Split(cmd, " ")
	switch strings.ToLower(cmdArgs[0]) {
	case "quit", "exit":
		color.HiBlack("Goodbye!")
		os.Exit(0)
	case "gopher":
		c.gopher()
	case "node":
		color.Yellow("The 'node' command is enterprise only, not available in the influx 2.x CLI - were you looking for the 1.x InfluxDB CLI?")
	case "consistency":
		color.Yellow("The 'consistency' command is not available in the influx 2.x CLI - were you looking for the 1.x InfluxDB CLI?")
	case "help":
		c.help()
	case "history":
		color.HiBlack(strings.Join(c.readHistory(), "\n"))
	case "format":
		c.setFormat(cmdArgs)
	case "precision":
		c.setPrecision(cmdArgs)
	case "settings":
		c.settings()
	case "pretty":
		c.togglePretty()
	case "use":
		c.use(cmdArgs)
	case "insert":
		c.insert(cmd)
	case "clear":
		c.clear(cmd)
	default:
		c.runAndShowQuery(cmd)
	}
}

// Create a regex string for a named InfluxQL identifier, quoted or unquoted
func identRegex(name string) string {
	return `((?P<` + name + `>\w+)|(\"(?P<` + name + `_quote>.+)\"))`
}

// Get the value of a named InfluxQL identifier from a regex match map.
// Returns empty string if no match
func getIdentFromMatches(matches *map[string]string, name string) string {
	if val, ok := (*matches)[name]; ok && val != "" {
		return val
	} else if val, ok := (*matches)[name+"_quote"]; ok && val != "" {
		return val
	}
	return ""
}

// Create a regex match map from a regexp with named groups.
// Returns nil if no match.
func reSubMatchMap(r *regexp.Regexp, str string) *map[string]string {
	match := r.FindStringSubmatch(str)
	if match == nil {
		return nil
	}
	subMatchMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 {
			subMatchMap[name] = match[i]
		}
	}
	return &subMatchMap
}

var insertIntoRegex string = `^(?i)INSERT INTO ` + identRegex("db") + `(\.` + identRegex("rp") + `)? (?P<point>.+)$`
var insertRegex string = `^(?i)INSERT (?P<point>.+)$`

func (c Client) insert(cmd string) {
	var db string
	var rp string
	var point string
	insertRgx := regexp.MustCompile(insertRegex)
	insertIntoRgx := regexp.MustCompile(insertIntoRegex)
	insertMatches := reSubMatchMap(insertRgx, cmd)
	insertIntoMatches := reSubMatchMap(insertIntoRgx, cmd)
	if insertIntoMatches != nil {
		db = getIdentFromMatches(insertIntoMatches, "db")
		rp = getIdentFromMatches(insertIntoMatches, "rp")
		point = getIdentFromMatches(insertIntoMatches, "point")
		if db != "" && rp == "" {
			defaultRp, err := c.getDefaultRetentionPolicy(context.Background(), db)
			if err != nil {
				color.Red("Unable to get default retention policy for %q", db)
				return
			}
			rp = defaultRp
		}
	} else if !strings.HasPrefix(strings.ToUpper(cmd), "INSERT INTO") && insertMatches != nil {
		db = c.Database
		rp = c.RetentionPolicy
		point = getIdentFromMatches(insertMatches, "point")
		if db == "" {
			color.Red("Please run \"use <database>\" to run \"INSERT <point\"")
			return
		}
	} else {
		color.Red("Expected \"INSERT INTO <database>.<retention_policy> <point>\" OR \"INSERT <point>\".")
		return
	}

	buf := bytes.Buffer{}
	gzw := gzip.NewWriter(&buf)

	_, err := gzw.Write([]byte(point))
	gzw.Close()
	if err != nil {
		color.Red("Failed to gzip points")
		return
	}
	bucketID, err := c.getBucketIdForDBRP(db, rp)
	if err != nil {
		color.Red("Unable to match DBRP to BucketID: %v", err)
		return
	}
	writeReq := c.PostWrite(context.Background()).
		Bucket(bucketID).
		Precision(api.WritePrecision(c.Precision)).
		ContentEncoding("gzip").
		Body(buf.Bytes())
	if c.OrgID != "" {
		writeReq = writeReq.Org(c.OrgID)
	} else if c.OrgName != "" {
		writeReq = writeReq.Org(c.OrgName)
	} else {
		writeReq = writeReq.Org(c.ActiveConfig.Org)
	}
	if err := writeReq.Execute(); err != nil {
		color.Red("ERR: %v", err)
		if c.Database == "" {
			fmt.Println("Note: error may be due to not setting a database or retention policy.")
			fmt.Println(`Please set a database with the command "use <database>"`)
			return
		}
	}
}

// Reverse search for a bucket from db & rp
func (c Client) getBucketIdForDBRP(db string, rp string) (string, error) {
	dbrpsReq := c.GetDBRPs(context.Background())
	if c.OrgID != "" {
		dbrpsReq = dbrpsReq.OrgID(c.OrgID)
	}
	if c.OrgName != "" {
		dbrpsReq = dbrpsReq.Org(c.OrgName)
	}
	if c.OrgID == "" && c.OrgName == "" {
		dbrpsReq = dbrpsReq.Org(c.ActiveConfig.Org)
	}
	dbrps, err := dbrpsReq.Execute()
	if err != nil {
		return "", err
	}
	bucketID := ""
	for _, dbrp := range *dbrps.Content {
		if dbrp.Database == db && dbrp.RetentionPolicy == rp {
			bucketID = dbrp.BucketID
			break
		}
	}
	if bucketID == "" {
		return "", fmt.Errorf("unable to find bucket ID for %q.%q", db, rp)
	}
	return bucketID, nil
}

type FormatType string

var (
	CsvFormat   FormatType = "csv"
	JsonFormat  FormatType = "json"
	TableFormat FormatType = "table"
)

func (c *Client) runAndShowQuery(query string) {
	// TODO: guide users trying to use deprecated InfluxQL queries
	response, err := c.query(context.Background(), query)
	if err != nil {
		color.HiRed("Query failed.")
		color.Red("%v", err)
		return
	}
	displayMap := map[FormatType]func(string){
		CsvFormat:   c.outputCsv,
		JsonFormat:  c.outputJson,
		TableFormat: c.outputTable,
	}
	displayFunc := displayMap[c.Format]
	displayFunc(response)
}

func (c *Client) help() {
	fmt.Println(`Usage:
		pretty                toggles pretty print for the json format
        use <db_name>         sets current database
        format <format>       specifies the format of the server responses: json, csv, or column
        precision <format>    specifies the format of the timestamp: h, m, s, ms, u or ns
        history               displays command history
        settings              outputs the current settings for the shell
        clear                 clears settings such as database or retention policy.  run 'clear' for help
        exit/quit/ctrl+d      quits the influx shell

        show databases        show database names
        show series           show series information
        show measurements     show measurement information
        show tag keys         show tag key information
        show field keys       show field key information
		insert <point>        insert point into currently-used database

        A full list of influxql commands can be found at:
        https://docs.influxdata.com/influxdb/latest/query_language/spec/
		
	Keybindings:
		<CTRL+D>      exit 
		<CTRL+L>      clear screen
		<UP ARROW>    previous command
		<DOWN ARROW>  next command`)
}

func (c *Client) settings() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 1, 1, ' ', 0)
	fmt.Fprintln(w, "Setting\tValue")
	fmt.Fprintln(w, "--------\t--------")
	fmt.Fprintf(w, "Username\t%s\n", c.Username)
	fmt.Fprintf(w, "Database\t%s\n", c.Database)
	fmt.Fprintf(w, "RetentionPolicy\t%s\n", c.RetentionPolicy)
	fmt.Fprintf(w, "Pretty\t%v\n", c.Pretty)
	fmt.Fprintf(w, "Format\t%s\n", c.Format)
	fmt.Fprintf(w, "Precision\t%s\n", c.Precision)
	fmt.Fprintln(w)
	w.Flush()
}

func (c *Client) query(ctx context.Context, query string) (string, error) {
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
		U(c.Username).
		P(c.Password).
		Db(c.Database).
		Q(query).
		Rp(c.RetentionPolicy).
		Accept(resContentType).
		Execute()
	if err != nil {
		return "", err
	}
	return resBody, nil
}

func (c *Client) setFormat(args []string) {
	// args[0] is "format"
	if len(args) != 2 {
		color.Red("Expected a format [csv, json, table]")
		return
	}
	newFormat := FormatType(args[1])
	switch newFormat {
	case CsvFormat, JsonFormat, TableFormat:
		c.Format = newFormat
	default:
		color.HiRed("Unimplemented format %q, keeping %s format.", newFormat, c.Format)
		color.HiBlack("Choose a format from [csv, json, table]")
	}
}

func (c *Client) setPrecision(args []string) {
	// args[0] is "precision"
	if len(args) != 2 {
		color.Red("Expected a precision [ns, u, ms, s, m, or h]")
		return
	}
	precision := args[1]
	switch precision {
	case "ns", "u", "µ", "ms", "s", "m", "h":
		c.Precision = precision
	default:
		color.HiRed("Unimplemented precision %q, keeping %s precision.", precision, c.Precision)
		color.HiBlack("Choose a precision from [ns, u, ms, s, m, or h]")
	}
}

func (c *Client) outputCsv(csvBody string) {
	fmt.Println(csvBody)
}

func (c *Client) outputJson(jsonBody string) {
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

func (c *Client) outputTable(jsonBody string) {
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

func (c *Client) togglePretty() {
	c.Pretty = !c.Pretty
	color.HiBlack("Pretty: %v", c.Pretty)
}

func (c *Client) use(args []string) {
	if len(args) != 2 {
		color.Red("wrong number of args for \"use <database>\"")
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
	// discover if the parsedDb is a valid database
	for _, db := range dbs {
		if parsedDb == db {
			exists := false
			prevDb := c.Database
			c.Database = parsedDb
			rps, _ := c.getRetentionPolicies(context.Background())
			// discover if the parsedRp is a valid retention policy
			for _, rp := range rps {
				switch parsedRp {
				case "":
					c.RetentionPolicy, _ = c.getDefaultRetentionPolicy(context.Background(), c.Database)
				case rp:
					c.RetentionPolicy = parsedRp
				default:
					continue
				}
				c.RetentionPolicies = rps
				exists = true
				c.Measurements, _ = c.GetMeasurements(context.Background())
				break
			}
			if !exists {
				color.Red("No such retention policy %q exists on %q", parsedRp, c.Database)
				color.HiBlack("Available retention policies on %q:", parsedDb)
				for _, rp := range rps {
					color.HiBlack("- %q", rp)
				}
				c.Database = prevDb
				return
			}
			c.Database = parsedDb
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

// Get retention policies from the currently used database
func (c *Client) getRetentionPolicies(ctx context.Context) ([]string, error) {
	singleSeries, err := c.getDataSingleSeries(ctx,
		fmt.Sprintf("SHOW RETENTION POLICIES ON %q", c.Database))
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

// Get the default retention policy for a given database
func (c *Client) getDefaultRetentionPolicy(ctx context.Context, db string) (string, error) {
	singleSeries, err := c.getDataSingleSeries(ctx,
		fmt.Sprintf("SHOW RETENTION POLICIES ON %q", db))
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

// Get list of database names
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

// Get list of measurements for currently used database and retention policy
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

// Helper function to execute query & parse response, expecting a single series
func (c *Client) getDataSingleSeries(ctx context.Context, query string) (*api.InfluxqlJsonResponseSeries, error) {
	resBody, err := c.GetLegacyQuery(ctx).
		U(c.Username).
		P(c.Password).
		Db(c.Database).
		Q(query).
		Rp(c.RetentionPolicy).
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

// Parse database and retention policy from byte slice.
// Expects format like "db"."rp", db.rp, db, "db".
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

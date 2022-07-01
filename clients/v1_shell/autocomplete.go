package v1shell

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/influxdata/go-prompt"
)

type subsuggestFnType = func(string) (map[string]SuggestNode, string)

type SuggestNode struct {
	Description  string
	subsuggestFn subsuggestFnType
}

func (c *Client) suggestUse(remainder string) (map[string]SuggestNode, string) {
	s := map[string]SuggestNode{}
	for _, db := range c.Databases {
		s["\""+db+"\""] = SuggestNode{Description: "Table Name"}
	}
	return s, remainder
}

func (c *Client) suggestDropMeasurement(remainder string) (map[string]SuggestNode, string) {
	s := map[string]SuggestNode{}
	if c.Database != "" && c.RetentionPolicy != "" {
		for _, m := range c.Measurements {
			s["\""+m+"\""] = SuggestNode{Description: fmt.Sprintf("Measurement on \"%s\".\"%s\"", c.Database, c.RetentionPolicy)}
		}
	}
	return s, remainder
}

func (c *Client) suggestDelete(remainder string) (map[string]SuggestNode, string) {
	s := map[string]SuggestNode{}
	fromReg := regexp.MustCompile(`(?i)FROM(\s+)` + identRegex("from_clause") + `?$`)
	matches := reSubMatchMap(fromReg, remainder)
	if matches != nil {
		if c.Database != "" && c.RetentionPolicy != "" {
			for _, m := range c.Measurements {
				s["\""+m+"\""] = SuggestNode{Description: fmt.Sprintf("Measurement on \"%s\".\"%s\"", c.Database, c.RetentionPolicy)}
			}
		}
		return s, getIdentFromMatches(matches, "from_clause")
	}
	return s, remainder
}

func (c *Client) suggestSelect(remainder string) (map[string]SuggestNode, string) {
	s := map[string]SuggestNode{}
	fromReg := regexp.MustCompile(`(?i)FROM(\s+)` + identRegex("from_clause") + `?$`)
	matches := reSubMatchMap(fromReg, remainder)
	if matches != nil {
		if c.Database != "" && c.RetentionPolicy != "" {
			for _, m := range c.Measurements {
				s["\""+m+"\""] = SuggestNode{Description: fmt.Sprintf("Measurement on \"%s\".\"%s\"", c.Database, c.RetentionPolicy)}
			}
		}
		if c.Database != "" {
			for _, rp := range c.RetentionPolicies {
				s["\""+rp+"\""] = SuggestNode{Description: "Retention Policy on " + c.Database}
			}
		}
		for _, db := range c.Databases {
			s["\""+db+"\""] = SuggestNode{Description: "Table Name"}
		}
		return s, getIdentFromMatches(matches, "from_clause")
	}
	return s, remainder
}

func getSuggestions(remainder string, s map[string]SuggestNode) ([]prompt.Suggest, string) {
	if len(remainder) > 0 {
		for term, node := range s {
			if strings.HasPrefix(strings.ToLower(remainder), term) || strings.HasPrefix(strings.ToUpper(remainder), term) {
				// if the cursor is at the end of a just completed word without trailing space, don't give autosuggestions yet
				if len(term) >= len(remainder) {
					return []prompt.Suggest{}, ""
				}
				// remainder is everything after the just-matched term
				rem := strings.TrimLeft(remainder[len(term):], " ")
				// if sugsuggestFn is nil, this is a terminating suggestion (leaf node with no more following suggestions)
				if node.subsuggestFn == nil {
					return []prompt.Suggest{}, rem
				}
				sugs, rem := node.subsuggestFn(rem)
				return getSuggestions(rem, sugs)
			}
		}
	}
	// if no match was found, convert the s map into suggestions and filter by the remaining characters
	var sugs []prompt.Suggest
	for text, node := range s {
		sugs = append(sugs, prompt.Suggest{Text: text, Description: node.Description})
	}
	return prompt.FilterFuzzy(sugs, remainder, true), remainder
}

// recursively creates a SuggestNode for each space-delimited word for each subsuggestion
func newSugNodeFn(subsugs ...string) subsuggestFnType {
	s := make(map[string]SuggestNode)
	for _, sug := range subsugs {
		word, rest, found := strings.Cut(sug, " ")
		if !found {
			s[word] = SuggestNode{}
		} else {
			s[word] = SuggestNode{subsuggestFn: newSugNodeFn(rest)}
		}
	}
	return func(rem string) (map[string]SuggestNode, string) {
		return s, rem
	}
}

func (c *Client) completer(d prompt.Document) []prompt.Suggest {
	// the commented-out lines are unsupported in 2.x
	suggestions := map[string]SuggestNode{
		"use":        {Description: "Set current database", subsuggestFn: c.suggestUse},
		"pretty":     {Description: "Toggle pretty print for the json format"},
		"scientific": {Description: "Toggle scientific number format for the table format"},
		"precision": {Description: "Specify the format of the timestamp",
			subsuggestFn: newSugNodeFn(
				"rfc3339",
				"ns",
				"u",
				"ms",
				"s",
				"m",
				"h",
			)},
		"history":  {Description: "Display shell history"},
		"settings": {Description: "Output the current shell settings"},
		"clear": {Description: "Clears settings such as database",
			subsuggestFn: newSugNodeFn(
				"db",
				"database",
				"retention policy",
				"rp",
			)},
		"exit":   {Description: "Exit the InfluxQL shell"},
		"quit":   {Description: "Exit the InfluxQL shell"},
		"gopher": {Description: "Display the Go Gopher"},
		"help":   {Description: "Display help options"},
		"format": {Description: "Specify the data display format",
			subsuggestFn: newSugNodeFn(
				"column",
				"csv",
				"json",
				"table",
			)},
		"SELECT":      {subsuggestFn: c.suggestSelect},
		"INSERT":      {},
		"INSERT INTO": {},
		"DELETE":      {subsuggestFn: c.suggestDelete},
		"SHOW": {subsuggestFn: newSugNodeFn(
			// "CONTINUOUS QUERIES",
			"DATABASES",
			// "DIAGNOSTICS",
			"FIELD KEY CARDINALITY",
			"FIELD KEYS",
			// "GRANTS",
			// "MEASUREMENT CARDINALITY",
			"MEASUREMENT EXACT CARDINALITY",
			"MEASUREMENTS",
			// "QUERIES",
			// "RETENTION POLICIES",
			"SERIES",
			// "SERIES CARDINALITY",
			"SERIES EXACT CARDINALITY",
			// "SHARD GROUPS",
			// "SHARDS",
			// "STATS",
			// "SUBSCRIPTIONS",
			"TAG KEY CARDINALITY",
			"TAG KEY EXACT CARDINALITY",
			"TAG KEYS",
			"TAG VALUES",
			// "USERS",
		)},
		// "CREATE": {subsuggestFn: newSugNodeFn(
		// 	"CONTINUOUS QUERY",
		// 	"DATABASE",
		// 	"USER",
		// 	"RETENTION POLICY",
		// 	"SUBSCRIPTION",
		// )},
		"DROP": {subsuggestFn: func(rem string) (map[string]SuggestNode, string) {
			return map[string]SuggestNode{
				// "CONTINUOUS QUERY": {},
				// "DATABASE":         {},
				"MEASUREMENT": {subsuggestFn: c.suggestDropMeasurement},
				// "RETENTION POLICY": {},
				// "SERIES":           {},
				// "SHARD":            {},
				// "SUBSCRIPTION":     {},
				// "USER":             {},
			}, rem
		}},
		"EXPLAIN": {subsuggestFn: newSugNodeFn("ANALYZE")},
		// "GRANT":   {},
		// "REVOKE":  {},
		// "ALTER RETENTION POLICY": {},
		// "SET PASSOWRD FOR": {},
		// "KILL QUERY":       {},
	}
	line := d.CurrentLineBeforeCursor()
	currentSuggestions, _ := getSuggestions(line, suggestions)
	sort.Slice(currentSuggestions, func(i, j int) bool {
		return strings.ToLower(currentSuggestions[i].Text) < strings.ToLower(currentSuggestions[j].Text)
	})
	return currentSuggestions
}

package v1repl

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/influxdata/go-prompt"
)

type SuggestNode struct {
	Description  string
	subsuggestFn func(string) map[string]SuggestNode
}

func (c *Client) suggestUse(remainder string) map[string]SuggestNode {
	s := map[string]SuggestNode{}
	for _, db := range c.Databases {
		s["\""+db+"\""] = SuggestNode{Description: "Table Name"}
	}
	return s
}

func (c *Client) suggestSelect(remainder string) map[string]SuggestNode {
	s := map[string]SuggestNode{}
	if isMatch, _ := regexp.Match(`(?i)FROM `+identRegex("from_clause")+`?$`, []byte(remainder)); isMatch {
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
	}
	return s
}

func getSuggestions(remainder string, s map[string]SuggestNode) []prompt.Suggest {
	// if remainder == "" {
	// 	return []prompt.Suggest{}
	// }
	firstWord, rest, found := strings.Cut(remainder, " ")
	if found {
		if node, found := s[strings.ToLower(firstWord)]; found {
			if node.subsuggestFn == nil {
				return []prompt.Suggest{}
			}
			return getSuggestions(rest, node.subsuggestFn(rest))
		} else if node, found := s[strings.ToUpper(firstWord)]; found {
			if node.subsuggestFn == nil {
				return []prompt.Suggest{}
			}
			return getSuggestions(rest, node.subsuggestFn(rest))
		}
	}
	var sugs []prompt.Suggest
	for text, node := range s {
		sugs = append(sugs, prompt.Suggest{Text: text, Description: node.Description})
	}
	return sugs
}

func (c *Client) completer(d prompt.Document) []prompt.Suggest {
	// the commented-out lines are unsupported in 2.x
	suggestions := map[string]SuggestNode{
		"use":    {Description: "Set current database", subsuggestFn: c.suggestUse},
		"pretty": {Description: "Toggle pretty print for the json format"},
		"precision": {Description: "Specify the format of the timestamp", subsuggestFn: func(rem string) map[string]SuggestNode {
			return map[string]SuggestNode{
				"rfc3339": {},
				"ns":      {},
				"u":       {},
				"ms":      {},
				"s":       {},
				"m":       {},
				"h":       {},
			}
		}},
		"history":  {Description: "Display shell history"},
		"settings": {Description: "Output the current shell settings"},
		"clear": {Description: "Clears settings such as database", subsuggestFn: func(rem string) map[string]SuggestNode {
			return map[string]SuggestNode{
				"db":               {},
				"database":         {},
				"retention policy": {},
				"rp":               {},
			}
		}},
		"exit":   {Description: "Exit the InfluxQL shell"},
		"quit":   {Description: "Exit the InfluxQL shell"},
		"gopher": {Description: "Display the Go Gopher"},
		"help":   {Description: "Display help options"},
		"format": {Description: "Specify the data display format", subsuggestFn: func(rem string) map[string]SuggestNode {
			return map[string]SuggestNode{
				"column": {},
				"csv":    {},
				"json":   {},
			}
		}},
		"SELECT":      {subsuggestFn: c.suggestSelect},
		"INSERT":      {},
		"INSERT INTO": {},
		"DELETE":      {},
		"SHOW": {subsuggestFn: func(rem string) map[string]SuggestNode {
			return map[string]SuggestNode{
				// "CONTINUOUS QUERIES":            {},
				"DATABASES": {},
				// "DIAGNOSTICS":                   {},
				"FIELD KEY CARDINALITY": {},
				"FIELD KEYS":            {},
				// "GRANTS":                {},
				// "MEASUREMENT CARDINALITY":       {},
				"MEASUREMENT EXACT CARDINALITY": {},
				"MEASUREMENTS":                  {},
				// "QUERIES":                 {},
				// "RETENTION POLICIES":      {},
				"SERIES": {},
				// "SERIES CARDINALITY": {},
				"SERIES EXACT CARDINALITY": {},
				// "SHARD GROUPS":            {},
				// "SHARDS":                  {},
				// "STATS":                   {},
				// "SUBSCRIPTIONS":           {},
				"TAG KEY CARDINALITY":       {},
				"TAG KEY EXACT CARDINALITY": {},
				"TAG KEYS":                  {},
				"TAG VALUES":                {},
				"USERS":                     {},
			}
		}},
		"CREATE": {subsuggestFn: func(rem string) map[string]SuggestNode {
			return map[string]SuggestNode{
				"CONTINUOUS QUERY": {},
				// "DATABASE":         {},
				"USER":             {},
				"RETENTION POLICY": {},
				"SUBSCRIPTION":     {},
			}
		}},
		"DROP": {subsuggestFn: func(rem string) map[string]SuggestNode {
			return map[string]SuggestNode{
				"CONTINUOUS QUERY": {},
				"DATABASE":         {},
				"MEASUREMENT":      {},
				"RETENTION POLICY": {},
				"SERIES":           {},
				"SHARD":            {},
				"SUBSCRIPTION":     {},
				"USER":             {},
			}
		}},
		"EXPLAIN":         {},
		"EXPLAIN ANALYZE": {},
		// "GRANT":   {},
		// "REVOKE":  {},
		// "ALTER RETENTION POLICY": {},
		// "SET PASSOWRD FOR": {},
		// "KILL QUERY":       {},
	}
	line := d.CurrentLineBeforeCursor()
	currentSuggestions := getSuggestions(line, suggestions)
	sort.Slice(currentSuggestions, func(i, j int) bool {
		return strings.ToLower(currentSuggestions[i].Text) < strings.ToLower(currentSuggestions[j].Text)
	})
	return prompt.FilterFuzzy(currentSuggestions, d.GetWordBeforeCursor(), true)
}

package api

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func (args *TemplateSummaryVariableArgs) Render() string {
	if args == nil {
		return "<nil>"
	}
	switch args.Type {
	case "map":
		b, err := json.Marshal(args.Values)
		if err != nil {
			log.Printf("WARN: failed to parse map-variable args: expected JSON, got %v\n", args.Values)
			return "<parsing err>"
		}
		return string(b)
	case "constant":
		values, ok := args.Values.([]interface{})
		if !ok {
			log.Printf("WARN: failed to parse constant-variable args: expected array, got %v\n", args.Values)
			return "<parsing err>"
		}
		var out []string
		for _, v := range values {
			out = append(out, fmt.Sprintf("%q", v))
		}
		return fmt.Sprintf("[%s]", strings.Join(out, " "))
	case "query":
		values, ok := args.Values.(map[string]interface{})
		if !ok {
			log.Printf("WARN: failed to parse query-variable args: expected JSON object, got %v\n", args.Values)
			return "<parsing err>"
		}
		qVal, ok := values["query"]
		if !ok {
			log.Printf("WARN: failed to parse query-variable args: no 'query' key in %v\n", values)
			return "<parsing err>"
		}
		lVal, ok := values["language"]
		if !ok {
			log.Printf("WARN: failed to parse query-variable args: no 'language' key in %v\n", values)
			return "<parsing err>"
		}
		return fmt.Sprintf("language=%q query=%q", lVal, qVal)
	default:
	}
	return "<unknown variable type>"
}

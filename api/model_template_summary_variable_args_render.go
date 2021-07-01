package api

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (args *TemplateSummaryVariableArgs) String() string {
	if args == nil {
		return "<nil>"
	}
	switch args.Type {
	case "map":
		b, err := json.Marshal(args.Values)
		if err != nil {
			return "{}"
		}
		return string(b)
	case "constant":
		values, ok := args.Values.([]interface{})
		if !ok {
			return "[]"
		}
		var out []string
		for _, v := range values {
			out = append(out, fmt.Sprintf("%q", v))
		}
		return fmt.Sprintf("[%s]", strings.Join(out, " "))
	case "query":
		values, ok := args.Values.(map[string]interface{})
		if !ok {
			return ""
		}
		qVal, ok := values["query"]
		if !ok {
			return ""
		}
		lVal, ok := values["language"]
		if !ok {
			return ""
		}
		return fmt.Sprintf("language=%q query=%q", lVal, qVal)
	default:
	}
	return "unknown variable argument"
}

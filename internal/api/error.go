package api

import (
	"fmt"
	"strings"
)

// Extension to let our API error type be used as a "standard" error.
func (o *Error) Error() string {
	if o.Message != "" && o.Err != nil {
		var b strings.Builder
		b.WriteString(o.Message)
		b.WriteString(": ")
		b.WriteString(*o.Err)
		return b.String()
	} else if o.Message != "" {
		return o.Message
	} else if o.Err != nil {
		return *o.Err
	}
	return fmt.Sprintf("<%s>", o.Code)
}

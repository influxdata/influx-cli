package v1

import (
	"fmt"
	"strings"
)

// Extensions to let our API error types be used as "standard" errors.

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

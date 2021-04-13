package api

import (
	"fmt"
	"strings"
)

func (e *Error) Error() string {
	if e.Message != "" && e.Err != nil {
		var b strings.Builder
		b.WriteString(e.Message)
		b.WriteString(": ")
		b.WriteString(*e.Err)
		return b.String()
	} else if e.Message != "" {
		return e.Message
	} else if e.Err != nil {
		return *e.Err
	}
	return fmt.Sprintf("<%s>", e.Code)
}

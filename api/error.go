package api

import (
	"fmt"
	"strings"
)

type ApiError interface {
	error
	ErrorCode() ErrorCode
}

// Extensions to let our API error types be used as "standard" errors.

func (o *Error) Error() string {
	if o.Message != nil && o.Err != nil {
		var b strings.Builder
		b.WriteString(*o.Message)
		b.WriteString(": ")
		b.WriteString(*o.Err)
		return b.String()
	} else if o.Message != nil {
		return *o.Message
	} else if o.Err != nil {
		return *o.Err
	}
	return fmt.Sprintf("<%s>", o.Code)
}

func (o *Error) ErrorCode() ErrorCode {
	return o.Code
}

func (o *HealthCheck) Error() string {
	if o.Status == HEALTHCHECKSTATUS_PASS {
		// Make sure we aren't misusing HealthCheck responses.
		panic("successful healthcheck used as an error!")
	}
	var message string
	if o.Message != nil {
		message = *o.Message
	} else {
		message = fmt.Sprintf("check %s failed", o.Name)
	}
	return fmt.Sprintf("health check failed: %s", message)
}

func (o *HealthCheck) ErrorCode() ErrorCode {
	if o.Status == HEALTHCHECKSTATUS_PASS {
		// Make sure we aren't misusing HealthCheck responses.
		panic("successful healthcheck used as an error!")
	}

	return ERRORCODE_INTERNAL_ERROR
}

func (o *LineProtocolError) Error() string {
	if o.Message == nil {
		return ""
	}
	return *o.Message
}

func (o *LineProtocolError) ErrorCode() ErrorCode {
	switch o.Code {
	case LINEPROTOCOLERRORCODE_CONFLICT:
		return ERRORCODE_CONFLICT
	case LINEPROTOCOLERRORCODE_EMPTY_VALUE:
		return ERRORCODE_EMPTY_VALUE
	case LINEPROTOCOLERRORCODE_NOT_FOUND:
		return ERRORCODE_NOT_FOUND
	case LINEPROTOCOLERRORCODE_UNAVAILABLE:
		return ERRORCODE_UNAVAILABLE
	case LINEPROTOCOLERRORCODE_INTERNAL_ERROR:
		return ERRORCODE_INTERNAL_ERROR
	case LINEPROTOCOLERRORCODE_INVALID:
		fallthrough
	default:
		return ERRORCODE_INVALID
	}
}

func (o *LineProtocolLengthError) Error() string {
	return o.Message
}

func (o *LineProtocolLengthError) ErrorCode() ErrorCode {
	switch o.Code {
	case LINEPROTOCOLLENGTHERRORCODE_INVALID:
		fallthrough
	default:
		return ERRORCODE_INVALID
	}
}

func (o *TemplateSummaryError) Error() string {
	// If there are no extended errors, simply give the normal error
	if o.Errors == nil || len(*o.Errors) == 0 {
		return o.Message
	}

	// otherwise try to collect the extended error info
	var errMsg []string
	seenErrs := map[string]struct{}{}
	for _, e := range *o.Errors {
		fieldPairs := make([]string, 0, len(e.Fields))
		for i, idx := range e.Indexes {
			field := e.Fields[i]
			if idx == nil || *idx == -1 {
				fieldPairs = append(fieldPairs, field)
				continue
			}
			fieldPairs = append(fieldPairs, fmt.Sprintf("%s[%d]", field, *idx))
		}
		msg := fmt.Sprintf("kind=%s field=%s reason=%q", e.Kind, strings.Join(fieldPairs, "."), e.Reason)
		if _, ok := seenErrs[msg]; ok {
			continue
		}
		seenErrs[msg] = struct{}{}
		errMsg = append(errMsg, msg)
	}

	return strings.Join(errMsg, "\n\t")
}

func (o *TemplateSummaryError) ErrorCode() ErrorCode {
	return ErrorCode(o.Code)
}

func (o *UnauthorizedRequestError) Error() string {
	return o.GetMessage()
}

func (o *UnauthorizedRequestError) ErrorCode() ErrorCode {
	return ErrorCode(o.GetCode())
}

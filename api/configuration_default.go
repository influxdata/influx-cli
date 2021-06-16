package api

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

type ConfigParams struct {
	Host             *url.URL
	UserAgent        string
	Token            *string
	TraceId          *string
	AllowInsecureTLS bool
	Debug            bool
}

// NewAPIConfig builds a configuration tailored to the InfluxDB v2 API.
func NewAPIConfig(params ConfigParams) *Configuration {
	clientTransport := http.DefaultTransport.(*http.Transport).Clone()
	clientTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: params.AllowInsecureTLS}

	apiConfig := NewConfiguration()
	apiConfig.Host = params.Host.Host
	apiConfig.Scheme = params.Host.Scheme
	apiConfig.UserAgent = params.UserAgent
	apiConfig.HTTPClient = &http.Client{Transport: clientTransport}
	if params.Token != nil {
		apiConfig.DefaultHeader["Authorization"] = fmt.Sprintf("Token %s", *params.Token)
	}
	if params.TraceId != nil {
		// NOTE: This is circumventing our codegen. If the header we use for tracing ever changes,
		// we'll need to manually update the string here to match.
		//
		// The alternative is to pass the trace ID to the business logic for every CLI command, and
		// use codegen'd logic to set the header on every HTTP request. Early versions of the CLI
		// used that technique, and we found it to be error-prone and easy to forget during testing.
		apiConfig.DefaultHeader["Zap-Trace-Span"] = *params.TraceId
	}
	apiConfig.Debug = params.Debug

	return apiConfig
}

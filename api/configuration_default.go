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
	Cookie           *string
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

	if params.Host.Path != "" {
		// initialize api server configurations with path of host  url
		apiConfig.Servers = ServerConfigurations{{URL: params.Host.Path}}
		for key := range apiConfig.OperationServers {
			apiConfig.OperationServers[key] = ServerConfigurations{{URL: params.Host.Path}}
		}
	}

	if params.Token != nil && *params.Token != "" {
		apiConfig.DefaultHeader["Authorization"] = fmt.Sprintf("Token %s", *params.Token)
	} else if params.Cookie != nil && *params.Cookie != "" {
		apiConfig.DefaultHeader["Cookie"] = fmt.Sprintf("influxdb-oss-session=%s", *params.Cookie)
	}
	if params.TraceId != nil {
		// NOTE: This is circumventing our codegen. If the header we use for tracing ever changes,
		// we'll need to manually update the string here to match.
		//
		// The alternative is to pass the trace ID to the business logic for every CLI command, and
		// use codegen'd logic to set the header on every HTTP request. Early versions of the CLI
		// used that technique, and we found it to be error-prone and easy to forget during testing.
		apiConfig.DefaultHeader["Zap-Trace-Span"] = *params.TraceId
		apiConfig.DefaultHeader["influx-debug-id"] = *params.TraceId
	}
	apiConfig.Debug = params.Debug

	return apiConfig
}

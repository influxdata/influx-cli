/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	_context "context"
	_fmt "fmt"
	_io "io"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

type ConfigApi interface {

	/*
			 * GetConfig Retrieve runtime configuration
			 * Returns the active runtime configuration of the InfluxDB instance.

		In InfluxDB v2.2+, use this endpoint to view your active runtime configuration,
		including flags and environment variables.

		#### Related guides

		- [View your runtime server configuration]({{% INFLUXDB_DOCS_URL %}}/reference/config-options/#view-your-runtime-server-configuration)

			 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 * @return ApiGetConfigRequest
	*/
	GetConfig(ctx _context.Context) ApiGetConfigRequest

	/*
	 * GetConfigExecute executes the request
	 * @return Config
	 */
	GetConfigExecute(r ApiGetConfigRequest) (Config, error)

	/*
	 * GetConfigExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 * @return Config
	 */
	GetConfigExecuteWithHttpInfo(r ApiGetConfigRequest) (Config, *_nethttp.Response, error)
}

// ConfigApiService ConfigApi service
type ConfigApiService service

type ApiGetConfigRequest struct {
	ctx          _context.Context
	ApiService   ConfigApi
	zapTraceSpan *string
}

func (r ApiGetConfigRequest) ZapTraceSpan(zapTraceSpan string) ApiGetConfigRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetConfigRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetConfigRequest) Execute() (Config, error) {
	return r.ApiService.GetConfigExecute(r)
}

func (r ApiGetConfigRequest) ExecuteWithHttpInfo() (Config, *_nethttp.Response, error) {
	return r.ApiService.GetConfigExecuteWithHttpInfo(r)
}

/*
  - GetConfig Retrieve runtime configuration
  - Returns the active runtime configuration of the InfluxDB instance.

In InfluxDB v2.2+, use this endpoint to view your active runtime configuration,
including flags and environment variables.

#### Related guides

- [View your runtime server configuration]({{% INFLUXDB_DOCS_URL %}}/reference/config-options/#view-your-runtime-server-configuration)

  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @return ApiGetConfigRequest
*/
func (a *ConfigApiService) GetConfig(ctx _context.Context) ApiGetConfigRequest {
	return ApiGetConfigRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Config
 */
func (a *ConfigApiService) GetConfigExecute(r ApiGetConfigRequest) (Config, error) {
	returnVal, _, err := a.GetConfigExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 * @return Config
 */
func (a *ConfigApiService) GetConfigExecuteWithHttpInfo(r ApiGetConfigRequest) (Config, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Config
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigApiService.GetConfig")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.zapTraceSpan != nil {
		localVarHeaderParams["Zap-Trace-Span"] = parameterToString(*r.zapTraceSpan, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	newErr := GenericOpenAPIError{
		buildHeader: localVarHTTPResponse.Header.Get("X-Influxdb-Build"),
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.body = localVarBody
		newErr.error = localVarHTTPResponse.Status
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
			newErr.model = &v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
		newErr.model = &v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	body, err := GunzipIfNeeded(localVarHTTPResponse)
	if err != nil {
		body.Close()
		newErr.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	localVarBody, err := _io.ReadAll(body)
	body.Close()
	if err != nil {
		newErr.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	newErr.body = localVarBody
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

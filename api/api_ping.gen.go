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
	_io "io"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

type PingApi interface {

	/*
			 * GetPing Get the status of the instance
			 * Retrieves the status and InfluxDB version of the instance.

		Use this endpoint to monitor uptime for the InfluxDB instance. The response
		returns a HTTP `204` status code to inform you the instance is available.

		#### InfluxDB Cloud

		- Isn't versioned and doesn't return `X-Influxdb-Version` in the headers.

		#### Related guides

		- [Influx ping]({{% INFLUXDB_DOCS_URL %}}/reference/cli/influx/ping/)

			 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 * @return ApiGetPingRequest
	*/
	GetPing(ctx _context.Context) ApiGetPingRequest

	/*
	 * GetPingExecute executes the request
	 */
	GetPingExecute(r ApiGetPingRequest) error

	/*
	 * GetPingExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 */
	GetPingExecuteWithHttpInfo(r ApiGetPingRequest) (*_nethttp.Response, error)

	/*
			 * HeadPing Get the status of the instance
			 * Returns the status and InfluxDB version of the instance.

		Use this endpoint to monitor uptime for the InfluxDB instance. The response
		returns a HTTP `204` status code to inform you the instance is available.

		#### InfluxDB Cloud

		- Isn't versioned and doesn't return `X-Influxdb-Version` in the headers.

		#### Related guides

		- [Influx ping]({{% INFLUXDB_DOCS_URL %}}/reference/cli/influx/ping/)

			 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 * @return ApiHeadPingRequest
	*/
	HeadPing(ctx _context.Context) ApiHeadPingRequest

	/*
	 * HeadPingExecute executes the request
	 */
	HeadPingExecute(r ApiHeadPingRequest) error

	/*
	 * HeadPingExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 */
	HeadPingExecuteWithHttpInfo(r ApiHeadPingRequest) (*_nethttp.Response, error)
}

// PingApiService PingApi service
type PingApiService service

type ApiGetPingRequest struct {
	ctx        _context.Context
	ApiService PingApi
}

func (r ApiGetPingRequest) Execute() error {
	return r.ApiService.GetPingExecute(r)
}

func (r ApiGetPingRequest) ExecuteWithHttpInfo() (*_nethttp.Response, error) {
	return r.ApiService.GetPingExecuteWithHttpInfo(r)
}

/*
  - GetPing Get the status of the instance
  - Retrieves the status and InfluxDB version of the instance.

Use this endpoint to monitor uptime for the InfluxDB instance. The response
returns a HTTP `204` status code to inform you the instance is available.

#### InfluxDB Cloud

- Isn't versioned and doesn't return `X-Influxdb-Version` in the headers.

#### Related guides

- [Influx ping]({{% INFLUXDB_DOCS_URL %}}/reference/cli/influx/ping/)

  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @return ApiGetPingRequest
*/
func (a *PingApiService) GetPing(ctx _context.Context) ApiGetPingRequest {
	return ApiGetPingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *PingApiService) GetPingExecute(r ApiGetPingRequest) error {
	_, err := a.GetPingExecuteWithHttpInfo(r)
	return err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 */
func (a *PingApiService) GetPingExecuteWithHttpInfo(r ApiGetPingRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PingApiService.GetPing")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ping"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	newErr := GenericOpenAPIError{
		buildHeader: localVarHTTPResponse.Header.Get("X-Influxdb-Build"),
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.body = localVarBody
		newErr.error = localVarHTTPResponse.Status
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiHeadPingRequest struct {
	ctx        _context.Context
	ApiService PingApi
}

func (r ApiHeadPingRequest) Execute() error {
	return r.ApiService.HeadPingExecute(r)
}

func (r ApiHeadPingRequest) ExecuteWithHttpInfo() (*_nethttp.Response, error) {
	return r.ApiService.HeadPingExecuteWithHttpInfo(r)
}

/*
  - HeadPing Get the status of the instance
  - Returns the status and InfluxDB version of the instance.

Use this endpoint to monitor uptime for the InfluxDB instance. The response
returns a HTTP `204` status code to inform you the instance is available.

#### InfluxDB Cloud

- Isn't versioned and doesn't return `X-Influxdb-Version` in the headers.

#### Related guides

- [Influx ping]({{% INFLUXDB_DOCS_URL %}}/reference/cli/influx/ping/)

  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @return ApiHeadPingRequest
*/
func (a *PingApiService) HeadPing(ctx _context.Context) ApiHeadPingRequest {
	return ApiHeadPingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *PingApiService) HeadPingExecute(r ApiHeadPingRequest) error {
	_, err := a.HeadPingExecuteWithHttpInfo(r)
	return err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 */
func (a *PingApiService) HeadPingExecuteWithHttpInfo(r ApiHeadPingRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodHead
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PingApiService.HeadPing")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ping"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	newErr := GenericOpenAPIError{
		buildHeader: localVarHTTPResponse.Header.Get("X-Influxdb-Build"),
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.body = localVarBody
		newErr.error = localVarHTTPResponse.Status
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

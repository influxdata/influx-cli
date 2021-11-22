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
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type TelegrafsApi interface {

	/*
	 * DeleteTelegrafsID Delete a Telegraf configuration
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param telegrafID The Telegraf configuration ID.
	 * @return ApiDeleteTelegrafsIDRequest
	 */
	DeleteTelegrafsID(ctx _context.Context, telegrafID string) ApiDeleteTelegrafsIDRequest

	/*
	 * DeleteTelegrafsIDExecute executes the request
	 */
	DeleteTelegrafsIDExecute(r ApiDeleteTelegrafsIDRequest) error

	/*
	 * DeleteTelegrafsIDExecuteWithHttpInfo executes the request with HTTP response info returned
	 */
	DeleteTelegrafsIDExecuteWithHttpInfo(r ApiDeleteTelegrafsIDRequest) (*_nethttp.Response, error)

	/*
	 * GetTelegrafs List all Telegraf configurations
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiGetTelegrafsRequest
	 */
	GetTelegrafs(ctx _context.Context) ApiGetTelegrafsRequest

	/*
	 * GetTelegrafsExecute executes the request
	 * @return Telegrafs
	 */
	GetTelegrafsExecute(r ApiGetTelegrafsRequest) (Telegrafs, error)

	/*
	   * GetTelegrafsExecuteWithHttpInfo executes the request with HTTP response info returned
	       * @return Telegrafs
	*/
	GetTelegrafsExecuteWithHttpInfo(r ApiGetTelegrafsRequest) (Telegrafs, *_nethttp.Response, error)

	/*
	 * GetTelegrafsID Retrieve a Telegraf configuration
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param telegrafID The Telegraf configuration ID.
	 * @return ApiGetTelegrafsIDRequest
	 */
	GetTelegrafsID(ctx _context.Context, telegrafID string) ApiGetTelegrafsIDRequest

	/*
	 * GetTelegrafsIDExecute executes the request
	 * @return Telegraf
	 */
	GetTelegrafsIDExecute(r ApiGetTelegrafsIDRequest) (Telegraf, error)

	/*
	   * GetTelegrafsIDExecuteWithHttpInfo executes the request with HTTP response info returned
	       * @return Telegraf
	*/
	GetTelegrafsIDExecuteWithHttpInfo(r ApiGetTelegrafsIDRequest) (Telegraf, *_nethttp.Response, error)

	/*
	 * PostTelegrafs Create a Telegraf configuration
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiPostTelegrafsRequest
	 */
	PostTelegrafs(ctx _context.Context) ApiPostTelegrafsRequest

	/*
	 * PostTelegrafsExecute executes the request
	 * @return Telegraf
	 */
	PostTelegrafsExecute(r ApiPostTelegrafsRequest) (Telegraf, error)

	/*
	   * PostTelegrafsExecuteWithHttpInfo executes the request with HTTP response info returned
	       * @return Telegraf
	*/
	PostTelegrafsExecuteWithHttpInfo(r ApiPostTelegrafsRequest) (Telegraf, *_nethttp.Response, error)

	/*
	 * PutTelegrafsID Update a Telegraf configuration
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param telegrafID The Telegraf config ID.
	 * @return ApiPutTelegrafsIDRequest
	 */
	PutTelegrafsID(ctx _context.Context, telegrafID string) ApiPutTelegrafsIDRequest

	/*
	 * PutTelegrafsIDExecute executes the request
	 * @return Telegraf
	 */
	PutTelegrafsIDExecute(r ApiPutTelegrafsIDRequest) (Telegraf, error)

	/*
	   * PutTelegrafsIDExecuteWithHttpInfo executes the request with HTTP response info returned
	       * @return Telegraf
	*/
	PutTelegrafsIDExecuteWithHttpInfo(r ApiPutTelegrafsIDRequest) (Telegraf, *_nethttp.Response, error)

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on OSS
	// servers.
	OnlyOSS() TelegrafsApi

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on cloud
	// servers.
	OnlyCloud() TelegrafsApi
}

// TelegrafsApiService TelegrafsApi service
type TelegrafsApiService service

func (a *TelegrafsApiService) OnlyOSS() TelegrafsApi {
	a.isOnlyOSS = true
	return a
}

func (a *TelegrafsApiService) OnlyCloud() TelegrafsApi {
	a.isOnlyCloud = true
	return a
}

type ApiDeleteTelegrafsIDRequest struct {
	ctx          _context.Context
	ApiService   TelegrafsApi
	telegrafID   string
	zapTraceSpan *string
}

func (r ApiDeleteTelegrafsIDRequest) TelegrafID(telegrafID string) ApiDeleteTelegrafsIDRequest {
	r.telegrafID = telegrafID
	return r
}
func (r ApiDeleteTelegrafsIDRequest) GetTelegrafID() string {
	return r.telegrafID
}

func (r ApiDeleteTelegrafsIDRequest) ZapTraceSpan(zapTraceSpan string) ApiDeleteTelegrafsIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiDeleteTelegrafsIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiDeleteTelegrafsIDRequest) Execute() error {
	return r.ApiService.DeleteTelegrafsIDExecute(r)
}

func (r ApiDeleteTelegrafsIDRequest) ExecuteWithHttpInfo() (*_nethttp.Response, error) {
	return r.ApiService.DeleteTelegrafsIDExecuteWithHttpInfo(r)
}

/*
 * DeleteTelegrafsID Delete a Telegraf configuration
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param telegrafID The Telegraf configuration ID.
 * @return ApiDeleteTelegrafsIDRequest
 */
func (a *TelegrafsApiService) DeleteTelegrafsID(ctx _context.Context, telegrafID string) ApiDeleteTelegrafsIDRequest {
	return ApiDeleteTelegrafsIDRequest{
		ApiService: a,
		ctx:        ctx,
		telegrafID: telegrafID,
	}
}

/*
 * Execute executes the request
 */
func (a *TelegrafsApiService) DeleteTelegrafsIDExecute(r ApiDeleteTelegrafsIDRequest) error {
	_, err := a.DeleteTelegrafsIDExecuteWithHttpInfo(r)
	return err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned
 */
func (a *TelegrafsApiService) DeleteTelegrafsIDExecuteWithHttpInfo(r ApiDeleteTelegrafsIDRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelegrafsApiService.DeleteTelegrafsID")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/telegrafs/{telegrafID}"
	localVarPath = strings.Replace(localVarPath, "{"+"telegrafID"+"}", _neturl.PathEscape(parameterToString(r.telegrafID, "")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	var errorPrefix string
	if a.isOnlyOSS {
		errorPrefix = "InfluxDB OSS-only command failed: "
	} else if a.isOnlyCloud {
		errorPrefix = "InfluxDB Cloud-only command failed: "
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, localVarHTTPResponse.Status),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
			return localVarHTTPResponse, newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
		newErr.model = &v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetTelegrafsRequest struct {
	ctx          _context.Context
	ApiService   TelegrafsApi
	zapTraceSpan *string
	orgID        *string
}

func (r ApiGetTelegrafsRequest) ZapTraceSpan(zapTraceSpan string) ApiGetTelegrafsRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetTelegrafsRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetTelegrafsRequest) OrgID(orgID string) ApiGetTelegrafsRequest {
	r.orgID = &orgID
	return r
}
func (r ApiGetTelegrafsRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiGetTelegrafsRequest) Execute() (Telegrafs, error) {
	return r.ApiService.GetTelegrafsExecute(r)
}

func (r ApiGetTelegrafsRequest) ExecuteWithHttpInfo() (Telegrafs, *_nethttp.Response, error) {
	return r.ApiService.GetTelegrafsExecuteWithHttpInfo(r)
}

/*
 * GetTelegrafs List all Telegraf configurations
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetTelegrafsRequest
 */
func (a *TelegrafsApiService) GetTelegrafs(ctx _context.Context) ApiGetTelegrafsRequest {
	return ApiGetTelegrafsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Telegrafs
 */
func (a *TelegrafsApiService) GetTelegrafsExecute(r ApiGetTelegrafsRequest) (Telegrafs, error) {
	returnVal, _, err := a.GetTelegrafsExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned
 * @return Telegrafs
 */
func (a *TelegrafsApiService) GetTelegrafsExecuteWithHttpInfo(r ApiGetTelegrafsRequest) (Telegrafs, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Telegrafs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelegrafsApiService.GetTelegrafs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/telegrafs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.orgID != nil {
		localVarQueryParams.Add("orgID", parameterToString(*r.orgID, ""))
	}
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

	var errorPrefix string
	if a.isOnlyOSS {
		errorPrefix = "InfluxDB OSS-only command failed: "
	} else if a.isOnlyCloud {
		errorPrefix = "InfluxDB Cloud-only command failed: "
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, localVarHTTPResponse.Status),
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
		return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	localVarBody, err := _io.ReadAll(body)
	body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, err.Error()),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetTelegrafsIDRequest struct {
	ctx          _context.Context
	ApiService   TelegrafsApi
	telegrafID   string
	zapTraceSpan *string
	accept       *string
}

func (r ApiGetTelegrafsIDRequest) TelegrafID(telegrafID string) ApiGetTelegrafsIDRequest {
	r.telegrafID = telegrafID
	return r
}
func (r ApiGetTelegrafsIDRequest) GetTelegrafID() string {
	return r.telegrafID
}

func (r ApiGetTelegrafsIDRequest) ZapTraceSpan(zapTraceSpan string) ApiGetTelegrafsIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetTelegrafsIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetTelegrafsIDRequest) Accept(accept string) ApiGetTelegrafsIDRequest {
	r.accept = &accept
	return r
}
func (r ApiGetTelegrafsIDRequest) GetAccept() *string {
	return r.accept
}

func (r ApiGetTelegrafsIDRequest) Execute() (Telegraf, error) {
	return r.ApiService.GetTelegrafsIDExecute(r)
}

func (r ApiGetTelegrafsIDRequest) ExecuteWithHttpInfo() (Telegraf, *_nethttp.Response, error) {
	return r.ApiService.GetTelegrafsIDExecuteWithHttpInfo(r)
}

/*
 * GetTelegrafsID Retrieve a Telegraf configuration
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param telegrafID The Telegraf configuration ID.
 * @return ApiGetTelegrafsIDRequest
 */
func (a *TelegrafsApiService) GetTelegrafsID(ctx _context.Context, telegrafID string) ApiGetTelegrafsIDRequest {
	return ApiGetTelegrafsIDRequest{
		ApiService: a,
		ctx:        ctx,
		telegrafID: telegrafID,
	}
}

/*
 * Execute executes the request
 * @return Telegraf
 */
func (a *TelegrafsApiService) GetTelegrafsIDExecute(r ApiGetTelegrafsIDRequest) (Telegraf, error) {
	returnVal, _, err := a.GetTelegrafsIDExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned
 * @return Telegraf
 */
func (a *TelegrafsApiService) GetTelegrafsIDExecuteWithHttpInfo(r ApiGetTelegrafsIDRequest) (Telegraf, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Telegraf
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelegrafsApiService.GetTelegrafsID")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/telegrafs/{telegrafID}"
	localVarPath = strings.Replace(localVarPath, "{"+"telegrafID"+"}", _neturl.PathEscape(parameterToString(r.telegrafID, "")), -1)

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
	if r.accept != nil {
		localVarHeaderParams["Accept"] = parameterToString(*r.accept, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var errorPrefix string
	if a.isOnlyOSS {
		errorPrefix = "InfluxDB OSS-only command failed: "
	} else if a.isOnlyCloud {
		errorPrefix = "InfluxDB Cloud-only command failed: "
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, localVarHTTPResponse.Status),
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
		return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	localVarBody, err := _io.ReadAll(body)
	body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, err.Error()),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostTelegrafsRequest struct {
	ctx             _context.Context
	ApiService      TelegrafsApi
	telegrafRequest *TelegrafRequest
	zapTraceSpan    *string
}

func (r ApiPostTelegrafsRequest) TelegrafRequest(telegrafRequest TelegrafRequest) ApiPostTelegrafsRequest {
	r.telegrafRequest = &telegrafRequest
	return r
}
func (r ApiPostTelegrafsRequest) GetTelegrafRequest() *TelegrafRequest {
	return r.telegrafRequest
}

func (r ApiPostTelegrafsRequest) ZapTraceSpan(zapTraceSpan string) ApiPostTelegrafsRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPostTelegrafsRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPostTelegrafsRequest) Execute() (Telegraf, error) {
	return r.ApiService.PostTelegrafsExecute(r)
}

func (r ApiPostTelegrafsRequest) ExecuteWithHttpInfo() (Telegraf, *_nethttp.Response, error) {
	return r.ApiService.PostTelegrafsExecuteWithHttpInfo(r)
}

/*
 * PostTelegrafs Create a Telegraf configuration
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostTelegrafsRequest
 */
func (a *TelegrafsApiService) PostTelegrafs(ctx _context.Context) ApiPostTelegrafsRequest {
	return ApiPostTelegrafsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Telegraf
 */
func (a *TelegrafsApiService) PostTelegrafsExecute(r ApiPostTelegrafsRequest) (Telegraf, error) {
	returnVal, _, err := a.PostTelegrafsExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned
 * @return Telegraf
 */
func (a *TelegrafsApiService) PostTelegrafsExecuteWithHttpInfo(r ApiPostTelegrafsRequest) (Telegraf, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Telegraf
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelegrafsApiService.PostTelegrafs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/telegrafs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.telegrafRequest == nil {
		return localVarReturnValue, nil, reportError("telegrafRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.telegrafRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var errorPrefix string
	if a.isOnlyOSS {
		errorPrefix = "InfluxDB OSS-only command failed: "
	} else if a.isOnlyCloud {
		errorPrefix = "InfluxDB Cloud-only command failed: "
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, localVarHTTPResponse.Status),
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
		return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	localVarBody, err := _io.ReadAll(body)
	body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, err.Error()),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutTelegrafsIDRequest struct {
	ctx             _context.Context
	ApiService      TelegrafsApi
	telegrafID      string
	telegrafRequest *TelegrafRequest
	zapTraceSpan    *string
}

func (r ApiPutTelegrafsIDRequest) TelegrafID(telegrafID string) ApiPutTelegrafsIDRequest {
	r.telegrafID = telegrafID
	return r
}
func (r ApiPutTelegrafsIDRequest) GetTelegrafID() string {
	return r.telegrafID
}

func (r ApiPutTelegrafsIDRequest) TelegrafRequest(telegrafRequest TelegrafRequest) ApiPutTelegrafsIDRequest {
	r.telegrafRequest = &telegrafRequest
	return r
}
func (r ApiPutTelegrafsIDRequest) GetTelegrafRequest() *TelegrafRequest {
	return r.telegrafRequest
}

func (r ApiPutTelegrafsIDRequest) ZapTraceSpan(zapTraceSpan string) ApiPutTelegrafsIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPutTelegrafsIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPutTelegrafsIDRequest) Execute() (Telegraf, error) {
	return r.ApiService.PutTelegrafsIDExecute(r)
}

func (r ApiPutTelegrafsIDRequest) ExecuteWithHttpInfo() (Telegraf, *_nethttp.Response, error) {
	return r.ApiService.PutTelegrafsIDExecuteWithHttpInfo(r)
}

/*
 * PutTelegrafsID Update a Telegraf configuration
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param telegrafID The Telegraf config ID.
 * @return ApiPutTelegrafsIDRequest
 */
func (a *TelegrafsApiService) PutTelegrafsID(ctx _context.Context, telegrafID string) ApiPutTelegrafsIDRequest {
	return ApiPutTelegrafsIDRequest{
		ApiService: a,
		ctx:        ctx,
		telegrafID: telegrafID,
	}
}

/*
 * Execute executes the request
 * @return Telegraf
 */
func (a *TelegrafsApiService) PutTelegrafsIDExecute(r ApiPutTelegrafsIDRequest) (Telegraf, error) {
	returnVal, _, err := a.PutTelegrafsIDExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned
 * @return Telegraf
 */
func (a *TelegrafsApiService) PutTelegrafsIDExecuteWithHttpInfo(r ApiPutTelegrafsIDRequest) (Telegraf, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Telegraf
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelegrafsApiService.PutTelegrafsID")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/telegrafs/{telegrafID}"
	localVarPath = strings.Replace(localVarPath, "{"+"telegrafID"+"}", _neturl.PathEscape(parameterToString(r.telegrafID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.telegrafRequest == nil {
		return localVarReturnValue, nil, reportError("telegrafRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.telegrafRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var errorPrefix string
	if a.isOnlyOSS {
		errorPrefix = "InfluxDB OSS-only command failed: "
	} else if a.isOnlyCloud {
		errorPrefix = "InfluxDB Cloud-only command failed: "
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, localVarHTTPResponse.Status),
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
		return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	localVarBody, err := _io.ReadAll(body)
	body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, err.Error()),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

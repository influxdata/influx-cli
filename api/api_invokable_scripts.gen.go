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

type InvokableScriptsApi interface {

	/*
	 * DeleteScriptsID Delete a script
	 * Deletes a script and all associated records.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param scriptID The ID of the script to delete.
	 * @return ApiDeleteScriptsIDRequest
	 */
	DeleteScriptsID(ctx _context.Context, scriptID string) ApiDeleteScriptsIDRequest

	/*
	 * DeleteScriptsIDExecute executes the request
	 */
	DeleteScriptsIDExecute(r ApiDeleteScriptsIDRequest) error

	/*
	 * DeleteScriptsIDExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 */
	DeleteScriptsIDExecuteWithHttpInfo(r ApiDeleteScriptsIDRequest) (*_nethttp.Response, error)

	/*
	 * GetScripts List scripts
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiGetScriptsRequest
	 */
	GetScripts(ctx _context.Context) ApiGetScriptsRequest

	/*
	 * GetScriptsExecute executes the request
	 * @return Scripts
	 */
	GetScriptsExecute(r ApiGetScriptsRequest) (Scripts, error)

	/*
	 * GetScriptsExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 * @return Scripts
	 */
	GetScriptsExecuteWithHttpInfo(r ApiGetScriptsRequest) (Scripts, *_nethttp.Response, error)

	/*
	 * GetScriptsID Retrieve a script
	 * Uses script ID to retrieve details of an invokable script.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param scriptID The script ID.
	 * @return ApiGetScriptsIDRequest
	 */
	GetScriptsID(ctx _context.Context, scriptID string) ApiGetScriptsIDRequest

	/*
	 * GetScriptsIDExecute executes the request
	 * @return Script
	 */
	GetScriptsIDExecute(r ApiGetScriptsIDRequest) (Script, error)

	/*
	 * GetScriptsIDExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 * @return Script
	 */
	GetScriptsIDExecuteWithHttpInfo(r ApiGetScriptsIDRequest) (Script, *_nethttp.Response, error)

	/*
	 * PatchScriptsID Update a script
	 * Updates properties (`name`, `description`, and `script`) of an invokable script.

	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param scriptID The script ID.
	 * @return ApiPatchScriptsIDRequest
	 */
	PatchScriptsID(ctx _context.Context, scriptID string) ApiPatchScriptsIDRequest

	/*
	 * PatchScriptsIDExecute executes the request
	 * @return Script
	 */
	PatchScriptsIDExecute(r ApiPatchScriptsIDRequest) (Script, error)

	/*
	 * PatchScriptsIDExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 * @return Script
	 */
	PatchScriptsIDExecuteWithHttpInfo(r ApiPatchScriptsIDRequest) (Script, *_nethttp.Response, error)

	/*
	 * PostScripts Create a script
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiPostScriptsRequest
	 */
	PostScripts(ctx _context.Context) ApiPostScriptsRequest

	/*
	 * PostScriptsExecute executes the request
	 * @return Script
	 */
	PostScriptsExecute(r ApiPostScriptsRequest) (Script, error)

	/*
	 * PostScriptsExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 * @return Script
	 */
	PostScriptsExecuteWithHttpInfo(r ApiPostScriptsRequest) (Script, *_nethttp.Response, error)

	/*
	 * PostScriptsIDInvoke Invoke a script
	 * Invokes a script and substitutes `params` keys referenced in the script with `params` key-values sent in the request body.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param scriptID
	 * @return ApiPostScriptsIDInvokeRequest
	 */
	PostScriptsIDInvoke(ctx _context.Context, scriptID string) ApiPostScriptsIDInvokeRequest

	/*
	 * PostScriptsIDInvokeExecute executes the request
	 * @return *os.File
	 */
	PostScriptsIDInvokeExecute(r ApiPostScriptsIDInvokeRequest) (*_nethttp.Response, error)

	/*
	 * PostScriptsIDInvokeExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 * @return *os.File
	 */
	PostScriptsIDInvokeExecuteWithHttpInfo(r ApiPostScriptsIDInvokeRequest) (*_nethttp.Response, *_nethttp.Response, error)
}

// InvokableScriptsApiService InvokableScriptsApi service
type InvokableScriptsApiService service

type ApiDeleteScriptsIDRequest struct {
	ctx        _context.Context
	ApiService InvokableScriptsApi
	scriptID   string
}

func (r ApiDeleteScriptsIDRequest) ScriptID(scriptID string) ApiDeleteScriptsIDRequest {
	r.scriptID = scriptID
	return r
}
func (r ApiDeleteScriptsIDRequest) GetScriptID() string {
	return r.scriptID
}

func (r ApiDeleteScriptsIDRequest) Execute() error {
	return r.ApiService.DeleteScriptsIDExecute(r)
}

func (r ApiDeleteScriptsIDRequest) ExecuteWithHttpInfo() (*_nethttp.Response, error) {
	return r.ApiService.DeleteScriptsIDExecuteWithHttpInfo(r)
}

/*
 * DeleteScriptsID Delete a script
 * Deletes a script and all associated records.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param scriptID The ID of the script to delete.
 * @return ApiDeleteScriptsIDRequest
 */
func (a *InvokableScriptsApiService) DeleteScriptsID(ctx _context.Context, scriptID string) ApiDeleteScriptsIDRequest {
	return ApiDeleteScriptsIDRequest{
		ApiService: a,
		ctx:        ctx,
		scriptID:   scriptID,
	}
}

/*
 * Execute executes the request
 */
func (a *InvokableScriptsApiService) DeleteScriptsIDExecute(r ApiDeleteScriptsIDRequest) error {
	_, err := a.DeleteScriptsIDExecuteWithHttpInfo(r)
	return err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 */
func (a *InvokableScriptsApiService) DeleteScriptsIDExecuteWithHttpInfo(r ApiDeleteScriptsIDRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvokableScriptsApiService.DeleteScriptsID")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scripts/{scriptID}"
	localVarPath = strings.Replace(localVarPath, "{"+"scriptID"+"}", _neturl.PathEscape(parameterToString(r.scriptID, "")), -1)

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

type ApiGetScriptsRequest struct {
	ctx        _context.Context
	ApiService InvokableScriptsApi
	limit      *int32
	offset     *int32
}

func (r ApiGetScriptsRequest) Limit(limit int32) ApiGetScriptsRequest {
	r.limit = &limit
	return r
}
func (r ApiGetScriptsRequest) GetLimit() *int32 {
	return r.limit
}

func (r ApiGetScriptsRequest) Offset(offset int32) ApiGetScriptsRequest {
	r.offset = &offset
	return r
}
func (r ApiGetScriptsRequest) GetOffset() *int32 {
	return r.offset
}

func (r ApiGetScriptsRequest) Execute() (Scripts, error) {
	return r.ApiService.GetScriptsExecute(r)
}

func (r ApiGetScriptsRequest) ExecuteWithHttpInfo() (Scripts, *_nethttp.Response, error) {
	return r.ApiService.GetScriptsExecuteWithHttpInfo(r)
}

/*
 * GetScripts List scripts
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetScriptsRequest
 */
func (a *InvokableScriptsApiService) GetScripts(ctx _context.Context) ApiGetScriptsRequest {
	return ApiGetScriptsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Scripts
 */
func (a *InvokableScriptsApiService) GetScriptsExecute(r ApiGetScriptsRequest) (Scripts, error) {
	returnVal, _, err := a.GetScriptsExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 * @return Scripts
 */
func (a *InvokableScriptsApiService) GetScriptsExecuteWithHttpInfo(r ApiGetScriptsRequest) (Scripts, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Scripts
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvokableScriptsApiService.GetScripts")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scripts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
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

type ApiGetScriptsIDRequest struct {
	ctx        _context.Context
	ApiService InvokableScriptsApi
	scriptID   string
}

func (r ApiGetScriptsIDRequest) ScriptID(scriptID string) ApiGetScriptsIDRequest {
	r.scriptID = scriptID
	return r
}
func (r ApiGetScriptsIDRequest) GetScriptID() string {
	return r.scriptID
}

func (r ApiGetScriptsIDRequest) Execute() (Script, error) {
	return r.ApiService.GetScriptsIDExecute(r)
}

func (r ApiGetScriptsIDRequest) ExecuteWithHttpInfo() (Script, *_nethttp.Response, error) {
	return r.ApiService.GetScriptsIDExecuteWithHttpInfo(r)
}

/*
 * GetScriptsID Retrieve a script
 * Uses script ID to retrieve details of an invokable script.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param scriptID The script ID.
 * @return ApiGetScriptsIDRequest
 */
func (a *InvokableScriptsApiService) GetScriptsID(ctx _context.Context, scriptID string) ApiGetScriptsIDRequest {
	return ApiGetScriptsIDRequest{
		ApiService: a,
		ctx:        ctx,
		scriptID:   scriptID,
	}
}

/*
 * Execute executes the request
 * @return Script
 */
func (a *InvokableScriptsApiService) GetScriptsIDExecute(r ApiGetScriptsIDRequest) (Script, error) {
	returnVal, _, err := a.GetScriptsIDExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 * @return Script
 */
func (a *InvokableScriptsApiService) GetScriptsIDExecuteWithHttpInfo(r ApiGetScriptsIDRequest) (Script, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Script
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvokableScriptsApiService.GetScriptsID")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scripts/{scriptID}"
	localVarPath = strings.Replace(localVarPath, "{"+"scriptID"+"}", _neturl.PathEscape(parameterToString(r.scriptID, "")), -1)

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

type ApiPatchScriptsIDRequest struct {
	ctx                 _context.Context
	ApiService          InvokableScriptsApi
	scriptID            string
	scriptUpdateRequest *ScriptUpdateRequest
}

func (r ApiPatchScriptsIDRequest) ScriptID(scriptID string) ApiPatchScriptsIDRequest {
	r.scriptID = scriptID
	return r
}
func (r ApiPatchScriptsIDRequest) GetScriptID() string {
	return r.scriptID
}

func (r ApiPatchScriptsIDRequest) ScriptUpdateRequest(scriptUpdateRequest ScriptUpdateRequest) ApiPatchScriptsIDRequest {
	r.scriptUpdateRequest = &scriptUpdateRequest
	return r
}
func (r ApiPatchScriptsIDRequest) GetScriptUpdateRequest() *ScriptUpdateRequest {
	return r.scriptUpdateRequest
}

func (r ApiPatchScriptsIDRequest) Execute() (Script, error) {
	return r.ApiService.PatchScriptsIDExecute(r)
}

func (r ApiPatchScriptsIDRequest) ExecuteWithHttpInfo() (Script, *_nethttp.Response, error) {
	return r.ApiService.PatchScriptsIDExecuteWithHttpInfo(r)
}

/*
 * PatchScriptsID Update a script
 * Updates properties (`name`, `description`, and `script`) of an invokable script.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param scriptID The script ID.
 * @return ApiPatchScriptsIDRequest
 */
func (a *InvokableScriptsApiService) PatchScriptsID(ctx _context.Context, scriptID string) ApiPatchScriptsIDRequest {
	return ApiPatchScriptsIDRequest{
		ApiService: a,
		ctx:        ctx,
		scriptID:   scriptID,
	}
}

/*
 * Execute executes the request
 * @return Script
 */
func (a *InvokableScriptsApiService) PatchScriptsIDExecute(r ApiPatchScriptsIDRequest) (Script, error) {
	returnVal, _, err := a.PatchScriptsIDExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 * @return Script
 */
func (a *InvokableScriptsApiService) PatchScriptsIDExecuteWithHttpInfo(r ApiPatchScriptsIDRequest) (Script, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Script
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvokableScriptsApiService.PatchScriptsID")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scripts/{scriptID}"
	localVarPath = strings.Replace(localVarPath, "{"+"scriptID"+"}", _neturl.PathEscape(parameterToString(r.scriptID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.scriptUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("scriptUpdateRequest is required and must be specified")
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
	// body params
	localVarPostBody = r.scriptUpdateRequest
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

type ApiPostScriptsRequest struct {
	ctx                 _context.Context
	ApiService          InvokableScriptsApi
	scriptCreateRequest *ScriptCreateRequest
}

func (r ApiPostScriptsRequest) ScriptCreateRequest(scriptCreateRequest ScriptCreateRequest) ApiPostScriptsRequest {
	r.scriptCreateRequest = &scriptCreateRequest
	return r
}
func (r ApiPostScriptsRequest) GetScriptCreateRequest() *ScriptCreateRequest {
	return r.scriptCreateRequest
}

func (r ApiPostScriptsRequest) Execute() (Script, error) {
	return r.ApiService.PostScriptsExecute(r)
}

func (r ApiPostScriptsRequest) ExecuteWithHttpInfo() (Script, *_nethttp.Response, error) {
	return r.ApiService.PostScriptsExecuteWithHttpInfo(r)
}

/*
 * PostScripts Create a script
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostScriptsRequest
 */
func (a *InvokableScriptsApiService) PostScripts(ctx _context.Context) ApiPostScriptsRequest {
	return ApiPostScriptsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Script
 */
func (a *InvokableScriptsApiService) PostScriptsExecute(r ApiPostScriptsRequest) (Script, error) {
	returnVal, _, err := a.PostScriptsExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 * @return Script
 */
func (a *InvokableScriptsApiService) PostScriptsExecuteWithHttpInfo(r ApiPostScriptsRequest) (Script, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Script
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvokableScriptsApiService.PostScripts")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scripts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.scriptCreateRequest == nil {
		return localVarReturnValue, nil, reportError("scriptCreateRequest is required and must be specified")
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
	// body params
	localVarPostBody = r.scriptCreateRequest
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

type ApiPostScriptsIDInvokeRequest struct {
	ctx                    _context.Context
	ApiService             InvokableScriptsApi
	scriptID               string
	scriptInvocationParams *ScriptInvocationParams
}

func (r ApiPostScriptsIDInvokeRequest) ScriptID(scriptID string) ApiPostScriptsIDInvokeRequest {
	r.scriptID = scriptID
	return r
}
func (r ApiPostScriptsIDInvokeRequest) GetScriptID() string {
	return r.scriptID
}

func (r ApiPostScriptsIDInvokeRequest) ScriptInvocationParams(scriptInvocationParams ScriptInvocationParams) ApiPostScriptsIDInvokeRequest {
	r.scriptInvocationParams = &scriptInvocationParams
	return r
}
func (r ApiPostScriptsIDInvokeRequest) GetScriptInvocationParams() *ScriptInvocationParams {
	return r.scriptInvocationParams
}

func (r ApiPostScriptsIDInvokeRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PostScriptsIDInvokeExecute(r)
}

func (r ApiPostScriptsIDInvokeRequest) ExecuteWithHttpInfo() (*_nethttp.Response, *_nethttp.Response, error) {
	return r.ApiService.PostScriptsIDInvokeExecuteWithHttpInfo(r)
}

/*
 * PostScriptsIDInvoke Invoke a script
 * Invokes a script and substitutes `params` keys referenced in the script with `params` key-values sent in the request body.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param scriptID
 * @return ApiPostScriptsIDInvokeRequest
 */
func (a *InvokableScriptsApiService) PostScriptsIDInvoke(ctx _context.Context, scriptID string) ApiPostScriptsIDInvokeRequest {
	return ApiPostScriptsIDInvokeRequest{
		ApiService: a,
		ctx:        ctx,
		scriptID:   scriptID,
	}
}

/*
 * Execute executes the request
 * @return *os.File
 */
func (a *InvokableScriptsApiService) PostScriptsIDInvokeExecute(r ApiPostScriptsIDInvokeRequest) (*_nethttp.Response, error) {
	returnVal, _, err := a.PostScriptsIDInvokeExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 * @return *os.File
 */
func (a *InvokableScriptsApiService) PostScriptsIDInvokeExecuteWithHttpInfo(r ApiPostScriptsIDInvokeRequest) (*_nethttp.Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *_nethttp.Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvokableScriptsApiService.PostScriptsIDInvoke")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scripts/{scriptID}/invoke"
	localVarPath = strings.Replace(localVarPath, "{"+"scriptID"+"}", _neturl.PathEscape(parameterToString(r.scriptID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	// body params
	localVarPostBody = r.scriptInvocationParams
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

	localVarReturnValue = localVarHTTPResponse

	return localVarReturnValue, localVarHTTPResponse, nil
}

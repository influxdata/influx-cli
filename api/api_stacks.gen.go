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
	"reflect"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type StacksApi interface {

	/*
	 * CreateStack Create a new stack
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiCreateStackRequest
	 */
	CreateStack(ctx _context.Context) ApiCreateStackRequest

	/*
	 * CreateStackExecute executes the request
	 * @return Stack
	 */
	CreateStackExecute(r ApiCreateStackRequest) (Stack, error)

	/*
	 * CreateStackExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 * @return Stack
	 */
	CreateStackExecuteWithHttpInfo(r ApiCreateStackRequest) (Stack, *_nethttp.Response, error)

	/*
	 * DeleteStack Delete a stack and associated resources
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param stackId The identifier of the stack.
	 * @return ApiDeleteStackRequest
	 */
	DeleteStack(ctx _context.Context, stackId string) ApiDeleteStackRequest

	/*
	 * DeleteStackExecute executes the request
	 */
	DeleteStackExecute(r ApiDeleteStackRequest) error

	/*
	 * DeleteStackExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 */
	DeleteStackExecuteWithHttpInfo(r ApiDeleteStackRequest) (*_nethttp.Response, error)

	/*
	 * ListStacks List all installed InfluxDB templates
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiListStacksRequest
	 */
	ListStacks(ctx _context.Context) ApiListStacksRequest

	/*
	 * ListStacksExecute executes the request
	 * @return Stacks
	 */
	ListStacksExecute(r ApiListStacksRequest) (Stacks, error)

	/*
	 * ListStacksExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 * @return Stacks
	 */
	ListStacksExecuteWithHttpInfo(r ApiListStacksRequest) (Stacks, *_nethttp.Response, error)

	/*
	 * ReadStack Retrieve a stack
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param stackId The identifier of the stack.
	 * @return ApiReadStackRequest
	 */
	ReadStack(ctx _context.Context, stackId string) ApiReadStackRequest

	/*
	 * ReadStackExecute executes the request
	 * @return Stack
	 */
	ReadStackExecute(r ApiReadStackRequest) (Stack, error)

	/*
	 * ReadStackExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 * @return Stack
	 */
	ReadStackExecuteWithHttpInfo(r ApiReadStackRequest) (Stack, *_nethttp.Response, error)

	/*
	 * UpdateStack Update an InfluxDB Stack
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param stackId The identifier of the stack.
	 * @return ApiUpdateStackRequest
	 */
	UpdateStack(ctx _context.Context, stackId string) ApiUpdateStackRequest

	/*
	 * UpdateStackExecute executes the request
	 * @return Stack
	 */
	UpdateStackExecute(r ApiUpdateStackRequest) (Stack, error)

	/*
	 * UpdateStackExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 * @return Stack
	 */
	UpdateStackExecuteWithHttpInfo(r ApiUpdateStackRequest) (Stack, *_nethttp.Response, error)
}

// StacksApiService StacksApi service
type StacksApiService service

type ApiCreateStackRequest struct {
	ctx              _context.Context
	ApiService       StacksApi
	stackPostRequest *StackPostRequest
}

func (r ApiCreateStackRequest) StackPostRequest(stackPostRequest StackPostRequest) ApiCreateStackRequest {
	r.stackPostRequest = &stackPostRequest
	return r
}
func (r ApiCreateStackRequest) GetStackPostRequest() *StackPostRequest {
	return r.stackPostRequest
}

func (r ApiCreateStackRequest) Execute() (Stack, error) {
	return r.ApiService.CreateStackExecute(r)
}

func (r ApiCreateStackRequest) ExecuteWithHttpInfo() (Stack, *_nethttp.Response, error) {
	return r.ApiService.CreateStackExecuteWithHttpInfo(r)
}

/*
 * CreateStack Create a new stack
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateStackRequest
 */
func (a *StacksApiService) CreateStack(ctx _context.Context) ApiCreateStackRequest {
	return ApiCreateStackRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Stack
 */
func (a *StacksApiService) CreateStackExecute(r ApiCreateStackRequest) (Stack, error) {
	returnVal, _, err := a.CreateStackExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 * @return Stack
 */
func (a *StacksApiService) CreateStackExecuteWithHttpInfo(r ApiCreateStackRequest) (Stack, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Stack
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StacksApiService.CreateStack")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/stacks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.stackPostRequest == nil {
		return localVarReturnValue, nil, reportError("stackPostRequest is required and must be specified")
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
	localVarPostBody = r.stackPostRequest
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

type ApiDeleteStackRequest struct {
	ctx        _context.Context
	ApiService StacksApi
	stackId    string
	orgID      *string
}

func (r ApiDeleteStackRequest) StackId(stackId string) ApiDeleteStackRequest {
	r.stackId = stackId
	return r
}
func (r ApiDeleteStackRequest) GetStackId() string {
	return r.stackId
}

func (r ApiDeleteStackRequest) OrgID(orgID string) ApiDeleteStackRequest {
	r.orgID = &orgID
	return r
}
func (r ApiDeleteStackRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiDeleteStackRequest) Execute() error {
	return r.ApiService.DeleteStackExecute(r)
}

func (r ApiDeleteStackRequest) ExecuteWithHttpInfo() (*_nethttp.Response, error) {
	return r.ApiService.DeleteStackExecuteWithHttpInfo(r)
}

/*
 * DeleteStack Delete a stack and associated resources
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param stackId The identifier of the stack.
 * @return ApiDeleteStackRequest
 */
func (a *StacksApiService) DeleteStack(ctx _context.Context, stackId string) ApiDeleteStackRequest {
	return ApiDeleteStackRequest{
		ApiService: a,
		ctx:        ctx,
		stackId:    stackId,
	}
}

/*
 * Execute executes the request
 */
func (a *StacksApiService) DeleteStackExecute(r ApiDeleteStackRequest) error {
	_, err := a.DeleteStackExecuteWithHttpInfo(r)
	return err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 */
func (a *StacksApiService) DeleteStackExecuteWithHttpInfo(r ApiDeleteStackRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StacksApiService.DeleteStack")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/stacks/{stack_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"stack_id"+"}", _neturl.PathEscape(parameterToString(r.stackId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.orgID == nil {
		return nil, reportError("orgID is required and must be specified")
	}

	localVarQueryParams.Add("orgID", parameterToString(*r.orgID, ""))
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

type ApiListStacksRequest struct {
	ctx        _context.Context
	ApiService StacksApi
	orgID      *string
	name       *[]string
	stackID    *[]string
}

func (r ApiListStacksRequest) OrgID(orgID string) ApiListStacksRequest {
	r.orgID = &orgID
	return r
}
func (r ApiListStacksRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiListStacksRequest) Name(name []string) ApiListStacksRequest {
	r.name = &name
	return r
}
func (r ApiListStacksRequest) GetName() *[]string {
	return r.name
}

func (r ApiListStacksRequest) StackID(stackID []string) ApiListStacksRequest {
	r.stackID = &stackID
	return r
}
func (r ApiListStacksRequest) GetStackID() *[]string {
	return r.stackID
}

func (r ApiListStacksRequest) Execute() (Stacks, error) {
	return r.ApiService.ListStacksExecute(r)
}

func (r ApiListStacksRequest) ExecuteWithHttpInfo() (Stacks, *_nethttp.Response, error) {
	return r.ApiService.ListStacksExecuteWithHttpInfo(r)
}

/*
 * ListStacks List all installed InfluxDB templates
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListStacksRequest
 */
func (a *StacksApiService) ListStacks(ctx _context.Context) ApiListStacksRequest {
	return ApiListStacksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Stacks
 */
func (a *StacksApiService) ListStacksExecute(r ApiListStacksRequest) (Stacks, error) {
	returnVal, _, err := a.ListStacksExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 * @return Stacks
 */
func (a *StacksApiService) ListStacksExecuteWithHttpInfo(r ApiListStacksRequest) (Stacks, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Stacks
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StacksApiService.ListStacks")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/stacks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.orgID == nil {
		return localVarReturnValue, nil, reportError("orgID is required and must be specified")
	}

	localVarQueryParams.Add("orgID", parameterToString(*r.orgID, ""))
	if r.name != nil {
		t := *r.name
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("name", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("name", parameterToString(t, "multi"))
		}
	}
	if r.stackID != nil {
		t := *r.stackID
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("stackID", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("stackID", parameterToString(t, "multi"))
		}
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

type ApiReadStackRequest struct {
	ctx        _context.Context
	ApiService StacksApi
	stackId    string
}

func (r ApiReadStackRequest) StackId(stackId string) ApiReadStackRequest {
	r.stackId = stackId
	return r
}
func (r ApiReadStackRequest) GetStackId() string {
	return r.stackId
}

func (r ApiReadStackRequest) Execute() (Stack, error) {
	return r.ApiService.ReadStackExecute(r)
}

func (r ApiReadStackRequest) ExecuteWithHttpInfo() (Stack, *_nethttp.Response, error) {
	return r.ApiService.ReadStackExecuteWithHttpInfo(r)
}

/*
 * ReadStack Retrieve a stack
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param stackId The identifier of the stack.
 * @return ApiReadStackRequest
 */
func (a *StacksApiService) ReadStack(ctx _context.Context, stackId string) ApiReadStackRequest {
	return ApiReadStackRequest{
		ApiService: a,
		ctx:        ctx,
		stackId:    stackId,
	}
}

/*
 * Execute executes the request
 * @return Stack
 */
func (a *StacksApiService) ReadStackExecute(r ApiReadStackRequest) (Stack, error) {
	returnVal, _, err := a.ReadStackExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 * @return Stack
 */
func (a *StacksApiService) ReadStackExecuteWithHttpInfo(r ApiReadStackRequest) (Stack, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Stack
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StacksApiService.ReadStack")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/stacks/{stack_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"stack_id"+"}", _neturl.PathEscape(parameterToString(r.stackId, "")), -1)

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

type ApiUpdateStackRequest struct {
	ctx               _context.Context
	ApiService        StacksApi
	stackId           string
	stackPatchRequest *StackPatchRequest
}

func (r ApiUpdateStackRequest) StackId(stackId string) ApiUpdateStackRequest {
	r.stackId = stackId
	return r
}
func (r ApiUpdateStackRequest) GetStackId() string {
	return r.stackId
}

func (r ApiUpdateStackRequest) StackPatchRequest(stackPatchRequest StackPatchRequest) ApiUpdateStackRequest {
	r.stackPatchRequest = &stackPatchRequest
	return r
}
func (r ApiUpdateStackRequest) GetStackPatchRequest() *StackPatchRequest {
	return r.stackPatchRequest
}

func (r ApiUpdateStackRequest) Execute() (Stack, error) {
	return r.ApiService.UpdateStackExecute(r)
}

func (r ApiUpdateStackRequest) ExecuteWithHttpInfo() (Stack, *_nethttp.Response, error) {
	return r.ApiService.UpdateStackExecuteWithHttpInfo(r)
}

/*
 * UpdateStack Update an InfluxDB Stack
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param stackId The identifier of the stack.
 * @return ApiUpdateStackRequest
 */
func (a *StacksApiService) UpdateStack(ctx _context.Context, stackId string) ApiUpdateStackRequest {
	return ApiUpdateStackRequest{
		ApiService: a,
		ctx:        ctx,
		stackId:    stackId,
	}
}

/*
 * Execute executes the request
 * @return Stack
 */
func (a *StacksApiService) UpdateStackExecute(r ApiUpdateStackRequest) (Stack, error) {
	returnVal, _, err := a.UpdateStackExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 * @return Stack
 */
func (a *StacksApiService) UpdateStackExecuteWithHttpInfo(r ApiUpdateStackRequest) (Stack, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Stack
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StacksApiService.UpdateStack")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/stacks/{stack_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"stack_id"+"}", _neturl.PathEscape(parameterToString(r.stackId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.stackPatchRequest == nil {
		return localVarReturnValue, nil, reportError("stackPatchRequest is required and must be specified")
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
	localVarPostBody = r.stackPatchRequest
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

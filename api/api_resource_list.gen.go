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

type ResourceListApi interface {

	/*
	 * GetResources List all known resources
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiGetResourcesRequest
	 */
	GetResources(ctx _context.Context) ApiGetResourcesRequest

	/*
	 * GetResourcesExecute executes the request
	 * @return []string
	 */
	GetResourcesExecute(r ApiGetResourcesRequest) ([]string, error)

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on OSS
	// servers.
	OnlyOSS() ResourceListApi

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on cloud
	// servers.
	OnlyCloud() ResourceListApi
}

// ResourceListApiService ResourceListApi service
type ResourceListApiService service

func (a *ResourceListApiService) OnlyOSS() ResourceListApi {
	a.isOnlyOSS = true
	return a
}

func (a *ResourceListApiService) OnlyCloud() ResourceListApi {
	a.isOnlyCloud = true
	return a
}

type ApiGetResourcesRequest struct {
	ctx          _context.Context
	ApiService   ResourceListApi
	zapTraceSpan *string
}

func (r ApiGetResourcesRequest) ZapTraceSpan(zapTraceSpan string) ApiGetResourcesRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetResourcesRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetResourcesRequest) Execute() ([]string, error) {
	return r.ApiService.GetResourcesExecute(r)
}

/*
 * GetResources List all known resources
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetResourcesRequest
 */
func (a *ResourceListApiService) GetResources(ctx _context.Context) ApiGetResourcesRequest {
	return ApiGetResourcesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return []string
 */
func (a *ResourceListApiService) GetResourcesExecute(r ApiGetResourcesRequest) ([]string, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceListApiService.GetResources")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources"

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
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
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
			return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, localVarHTTPResponse.Status),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
			return localVarReturnValue, newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
		newErr.model = &v
		return localVarReturnValue, newErr
	}

	body, err := GunzipIfNeeded(localVarHTTPResponse)
	if err != nil {
		body.Close()
		return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	localVarBody, err := _io.ReadAll(body)
	body.Close()
	if err != nil {
		return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, err.Error()),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

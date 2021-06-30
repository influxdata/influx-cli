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
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type BucketsApi interface {

	/*
	 * DeleteBucketsID Delete a bucket
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param bucketID The ID of the bucket to delete.
	 * @return ApiDeleteBucketsIDRequest
	 */
	DeleteBucketsID(ctx _context.Context, bucketID string) ApiDeleteBucketsIDRequest

	/*
	 * DeleteBucketsIDExecute executes the request
	 */
	DeleteBucketsIDExecute(r ApiDeleteBucketsIDRequest) error

	/*
	 * GetBuckets List all buckets
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiGetBucketsRequest
	 */
	GetBuckets(ctx _context.Context) ApiGetBucketsRequest

	/*
	 * GetBucketsExecute executes the request
	 * @return Buckets
	 */
	GetBucketsExecute(r ApiGetBucketsRequest) (Buckets, error)

	/*
	 * GetBucketsID Retrieve a bucket
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param bucketID The bucket ID.
	 * @return ApiGetBucketsIDRequest
	 */
	GetBucketsID(ctx _context.Context, bucketID string) ApiGetBucketsIDRequest

	/*
	 * GetBucketsIDExecute executes the request
	 * @return Bucket
	 */
	GetBucketsIDExecute(r ApiGetBucketsIDRequest) (Bucket, error)

	/*
	 * PatchBucketsID Update a bucket
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param bucketID The bucket ID.
	 * @return ApiPatchBucketsIDRequest
	 */
	PatchBucketsID(ctx _context.Context, bucketID string) ApiPatchBucketsIDRequest

	/*
	 * PatchBucketsIDExecute executes the request
	 * @return Bucket
	 */
	PatchBucketsIDExecute(r ApiPatchBucketsIDRequest) (Bucket, error)

	/*
	 * PostBuckets Create a bucket
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiPostBucketsRequest
	 */
	PostBuckets(ctx _context.Context) ApiPostBucketsRequest

	/*
	 * PostBucketsExecute executes the request
	 * @return Bucket
	 */
	PostBucketsExecute(r ApiPostBucketsRequest) (Bucket, error)

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on OSS
	// servers.
	OnlyOSS() BucketsApi

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on cloud
	// servers.
	OnlyCloud() BucketsApi
}

// BucketsApiService BucketsApi service
type BucketsApiService service

func (a *BucketsApiService) OnlyOSS() BucketsApi {
	a.isOnlyOSS = true
	return a
}

func (a *BucketsApiService) OnlyCloud() BucketsApi {
	a.isOnlyCloud = true
	return a
}

type ApiDeleteBucketsIDRequest struct {
	ctx          _context.Context
	ApiService   BucketsApi
	bucketID     string
	zapTraceSpan *string
}

func (r ApiDeleteBucketsIDRequest) BucketID(bucketID string) ApiDeleteBucketsIDRequest {
	r.bucketID = bucketID
	return r
}
func (r ApiDeleteBucketsIDRequest) GetBucketID() string {
	return r.bucketID
}

func (r ApiDeleteBucketsIDRequest) ZapTraceSpan(zapTraceSpan string) ApiDeleteBucketsIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiDeleteBucketsIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiDeleteBucketsIDRequest) Execute() error {
	return r.ApiService.DeleteBucketsIDExecute(r)
}

/*
 * DeleteBucketsID Delete a bucket
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bucketID The ID of the bucket to delete.
 * @return ApiDeleteBucketsIDRequest
 */
func (a *BucketsApiService) DeleteBucketsID(ctx _context.Context, bucketID string) ApiDeleteBucketsIDRequest {
	return ApiDeleteBucketsIDRequest{
		ApiService: a,
		ctx:        ctx,
		bucketID:   bucketID,
	}
}

/*
 * Execute executes the request
 */
func (a *BucketsApiService) DeleteBucketsIDExecute(r ApiDeleteBucketsIDRequest) error {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketsApiService.DeleteBucketsID")
	if err != nil {
		return GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/buckets/{bucketID}"
	localVarPath = strings.Replace(localVarPath, "{"+"bucketID"+"}", _neturl.PathEscape(parameterToString(r.bucketID, "")), -1)

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
		return err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return err
	}

	var errorPrefix string
	if a.isOnlyOSS {
		errorPrefix = "InfluxDB OSS-only command failed"
	} else if a.isOnlyCloud {
		errorPrefix = "InfluxDB Cloud-only command failed"
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return _fmt.Errorf("%s: %w", errorPrefix, err)
		}
		localVarBody, err := _ioutil.ReadAll(body)
		body.Close()
		if err != nil {
			return _fmt.Errorf("%s: %w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s: %s", errorPrefix, localVarHTTPResponse.Status),
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
				return newErr
			}
			v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
			newErr.model = &v
			return newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
			return newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
		newErr.model = &v
		return newErr
	}

	return nil
}

type ApiGetBucketsRequest struct {
	ctx          _context.Context
	ApiService   BucketsApi
	zapTraceSpan *string
	offset       *int32
	limit        *int32
	after        *string
	org          *string
	orgID        *string
	name         *string
	id           *string
}

func (r ApiGetBucketsRequest) ZapTraceSpan(zapTraceSpan string) ApiGetBucketsRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetBucketsRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetBucketsRequest) Offset(offset int32) ApiGetBucketsRequest {
	r.offset = &offset
	return r
}
func (r ApiGetBucketsRequest) GetOffset() *int32 {
	return r.offset
}

func (r ApiGetBucketsRequest) Limit(limit int32) ApiGetBucketsRequest {
	r.limit = &limit
	return r
}
func (r ApiGetBucketsRequest) GetLimit() *int32 {
	return r.limit
}

func (r ApiGetBucketsRequest) After(after string) ApiGetBucketsRequest {
	r.after = &after
	return r
}
func (r ApiGetBucketsRequest) GetAfter() *string {
	return r.after
}

func (r ApiGetBucketsRequest) Org(org string) ApiGetBucketsRequest {
	r.org = &org
	return r
}
func (r ApiGetBucketsRequest) GetOrg() *string {
	return r.org
}

func (r ApiGetBucketsRequest) OrgID(orgID string) ApiGetBucketsRequest {
	r.orgID = &orgID
	return r
}
func (r ApiGetBucketsRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiGetBucketsRequest) Name(name string) ApiGetBucketsRequest {
	r.name = &name
	return r
}
func (r ApiGetBucketsRequest) GetName() *string {
	return r.name
}

func (r ApiGetBucketsRequest) Id(id string) ApiGetBucketsRequest {
	r.id = &id
	return r
}
func (r ApiGetBucketsRequest) GetId() *string {
	return r.id
}

func (r ApiGetBucketsRequest) Execute() (Buckets, error) {
	return r.ApiService.GetBucketsExecute(r)
}

/*
 * GetBuckets List all buckets
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetBucketsRequest
 */
func (a *BucketsApiService) GetBuckets(ctx _context.Context) ApiGetBucketsRequest {
	return ApiGetBucketsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Buckets
 */
func (a *BucketsApiService) GetBucketsExecute(r ApiGetBucketsRequest) (Buckets, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Buckets
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketsApiService.GetBuckets")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/buckets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.org != nil {
		localVarQueryParams.Add("org", parameterToString(*r.org, ""))
	}
	if r.orgID != nil {
		localVarQueryParams.Add("orgID", parameterToString(*r.orgID, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.id != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id, ""))
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
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
	}

	var errorPrefix string
	if a.isOnlyOSS {
		errorPrefix = "InfluxDB OSS-only command failed"
	} else if a.isOnlyCloud {
		errorPrefix = "InfluxDB Cloud-only command failed"
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
		}
		localVarBody, err := _ioutil.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s: %s", errorPrefix, localVarHTTPResponse.Status),
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
		return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
	}
	localVarBody, err := _ioutil.ReadAll(body)
	body.Close()
	if err != nil {
		return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
	}
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s: %s", errorPrefix, err.Error()),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

type ApiGetBucketsIDRequest struct {
	ctx          _context.Context
	ApiService   BucketsApi
	bucketID     string
	zapTraceSpan *string
}

func (r ApiGetBucketsIDRequest) BucketID(bucketID string) ApiGetBucketsIDRequest {
	r.bucketID = bucketID
	return r
}
func (r ApiGetBucketsIDRequest) GetBucketID() string {
	return r.bucketID
}

func (r ApiGetBucketsIDRequest) ZapTraceSpan(zapTraceSpan string) ApiGetBucketsIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetBucketsIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetBucketsIDRequest) Execute() (Bucket, error) {
	return r.ApiService.GetBucketsIDExecute(r)
}

/*
 * GetBucketsID Retrieve a bucket
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bucketID The bucket ID.
 * @return ApiGetBucketsIDRequest
 */
func (a *BucketsApiService) GetBucketsID(ctx _context.Context, bucketID string) ApiGetBucketsIDRequest {
	return ApiGetBucketsIDRequest{
		ApiService: a,
		ctx:        ctx,
		bucketID:   bucketID,
	}
}

/*
 * Execute executes the request
 * @return Bucket
 */
func (a *BucketsApiService) GetBucketsIDExecute(r ApiGetBucketsIDRequest) (Bucket, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Bucket
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketsApiService.GetBucketsID")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/buckets/{bucketID}"
	localVarPath = strings.Replace(localVarPath, "{"+"bucketID"+"}", _neturl.PathEscape(parameterToString(r.bucketID, "")), -1)

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
		errorPrefix = "InfluxDB OSS-only command failed"
	} else if a.isOnlyCloud {
		errorPrefix = "InfluxDB Cloud-only command failed"
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
		}
		localVarBody, err := _ioutil.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s: %s", errorPrefix, localVarHTTPResponse.Status),
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
		return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
	}
	localVarBody, err := _ioutil.ReadAll(body)
	body.Close()
	if err != nil {
		return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
	}
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s: %s", errorPrefix, err.Error()),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

type ApiPatchBucketsIDRequest struct {
	ctx                _context.Context
	ApiService         BucketsApi
	bucketID           string
	patchBucketRequest *PatchBucketRequest
	zapTraceSpan       *string
}

func (r ApiPatchBucketsIDRequest) BucketID(bucketID string) ApiPatchBucketsIDRequest {
	r.bucketID = bucketID
	return r
}
func (r ApiPatchBucketsIDRequest) GetBucketID() string {
	return r.bucketID
}

func (r ApiPatchBucketsIDRequest) PatchBucketRequest(patchBucketRequest PatchBucketRequest) ApiPatchBucketsIDRequest {
	r.patchBucketRequest = &patchBucketRequest
	return r
}
func (r ApiPatchBucketsIDRequest) GetPatchBucketRequest() *PatchBucketRequest {
	return r.patchBucketRequest
}

func (r ApiPatchBucketsIDRequest) ZapTraceSpan(zapTraceSpan string) ApiPatchBucketsIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPatchBucketsIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPatchBucketsIDRequest) Execute() (Bucket, error) {
	return r.ApiService.PatchBucketsIDExecute(r)
}

/*
 * PatchBucketsID Update a bucket
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bucketID The bucket ID.
 * @return ApiPatchBucketsIDRequest
 */
func (a *BucketsApiService) PatchBucketsID(ctx _context.Context, bucketID string) ApiPatchBucketsIDRequest {
	return ApiPatchBucketsIDRequest{
		ApiService: a,
		ctx:        ctx,
		bucketID:   bucketID,
	}
}

/*
 * Execute executes the request
 * @return Bucket
 */
func (a *BucketsApiService) PatchBucketsIDExecute(r ApiPatchBucketsIDRequest) (Bucket, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Bucket
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketsApiService.PatchBucketsID")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/buckets/{bucketID}"
	localVarPath = strings.Replace(localVarPath, "{"+"bucketID"+"}", _neturl.PathEscape(parameterToString(r.bucketID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.patchBucketRequest == nil {
		return localVarReturnValue, reportError("patchBucketRequest is required and must be specified")
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
	localVarPostBody = r.patchBucketRequest
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
		errorPrefix = "InfluxDB OSS-only command failed"
	} else if a.isOnlyCloud {
		errorPrefix = "InfluxDB Cloud-only command failed"
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
		}
		localVarBody, err := _ioutil.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s: %s", errorPrefix, localVarHTTPResponse.Status),
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
		return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
	}
	localVarBody, err := _ioutil.ReadAll(body)
	body.Close()
	if err != nil {
		return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
	}
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s: %s", errorPrefix, err.Error()),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

type ApiPostBucketsRequest struct {
	ctx               _context.Context
	ApiService        BucketsApi
	postBucketRequest *PostBucketRequest
	zapTraceSpan      *string
}

func (r ApiPostBucketsRequest) PostBucketRequest(postBucketRequest PostBucketRequest) ApiPostBucketsRequest {
	r.postBucketRequest = &postBucketRequest
	return r
}
func (r ApiPostBucketsRequest) GetPostBucketRequest() *PostBucketRequest {
	return r.postBucketRequest
}

func (r ApiPostBucketsRequest) ZapTraceSpan(zapTraceSpan string) ApiPostBucketsRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPostBucketsRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPostBucketsRequest) Execute() (Bucket, error) {
	return r.ApiService.PostBucketsExecute(r)
}

/*
 * PostBuckets Create a bucket
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostBucketsRequest
 */
func (a *BucketsApiService) PostBuckets(ctx _context.Context) ApiPostBucketsRequest {
	return ApiPostBucketsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Bucket
 */
func (a *BucketsApiService) PostBucketsExecute(r ApiPostBucketsRequest) (Bucket, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Bucket
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketsApiService.PostBuckets")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/buckets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.postBucketRequest == nil {
		return localVarReturnValue, reportError("postBucketRequest is required and must be specified")
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
	localVarPostBody = r.postBucketRequest
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
		errorPrefix = "InfluxDB OSS-only command failed"
	} else if a.isOnlyCloud {
		errorPrefix = "InfluxDB Cloud-only command failed"
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
		}
		localVarBody, err := _ioutil.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s: %s", errorPrefix, localVarHTTPResponse.Status),
		}
		if localVarHTTPResponse.StatusCode == 422 {
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
		return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
	}
	localVarBody, err := _ioutil.ReadAll(body)
	body.Close()
	if err != nil {
		return localVarReturnValue, _fmt.Errorf("%s: %w", errorPrefix, err)
	}
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s: %s", errorPrefix, err.Error()),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

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

type BucketSchemasApi interface {

	/*
	 * CreateMeasurementSchema Create a measurement schema for a bucket
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param bucketID The identifier of the bucket.
	 * @return ApiCreateMeasurementSchemaRequest
	 */
	CreateMeasurementSchema(ctx _context.Context, bucketID string) ApiCreateMeasurementSchemaRequest

	/*
	 * CreateMeasurementSchemaExecute executes the request
	 * @return MeasurementSchema
	 */
	CreateMeasurementSchemaExecute(r ApiCreateMeasurementSchemaRequest) (MeasurementSchema, error)

	/*
	 * GetMeasurementSchema Retrieve measurement schema information
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param bucketID The identifier of the bucket.
	 * @param measurementID The identifier of the measurement.
	 * @return ApiGetMeasurementSchemaRequest
	 */
	GetMeasurementSchema(ctx _context.Context, bucketID string, measurementID string) ApiGetMeasurementSchemaRequest

	/*
	 * GetMeasurementSchemaExecute executes the request
	 * @return MeasurementSchema
	 */
	GetMeasurementSchemaExecute(r ApiGetMeasurementSchemaRequest) (MeasurementSchema, error)

	/*
	 * GetMeasurementSchemas List all measurement schemas of a bucket
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param bucketID The identifier of the bucket.
	 * @return ApiGetMeasurementSchemasRequest
	 */
	GetMeasurementSchemas(ctx _context.Context, bucketID string) ApiGetMeasurementSchemasRequest

	/*
	 * GetMeasurementSchemasExecute executes the request
	 * @return MeasurementSchemaList
	 */
	GetMeasurementSchemasExecute(r ApiGetMeasurementSchemasRequest) (MeasurementSchemaList, error)

	/*
	 * UpdateMeasurementSchema Update a measurement schema
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param bucketID The identifier of the bucket.
	 * @param measurementID The identifier of the measurement.
	 * @return ApiUpdateMeasurementSchemaRequest
	 */
	UpdateMeasurementSchema(ctx _context.Context, bucketID string, measurementID string) ApiUpdateMeasurementSchemaRequest

	/*
	 * UpdateMeasurementSchemaExecute executes the request
	 * @return MeasurementSchema
	 */
	UpdateMeasurementSchemaExecute(r ApiUpdateMeasurementSchemaRequest) (MeasurementSchema, error)

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on OSS
	// servers.
	OnlyOSS() BucketSchemasApi

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on cloud
	// servers.
	OnlyCloud() BucketSchemasApi
}

// BucketSchemasApiService BucketSchemasApi service
type BucketSchemasApiService service

func (a *BucketSchemasApiService) OnlyOSS() BucketSchemasApi {
	a.isOnlyOSS = true
	return a
}

func (a *BucketSchemasApiService) OnlyCloud() BucketSchemasApi {
	a.isOnlyCloud = true
	return a
}

type ApiCreateMeasurementSchemaRequest struct {
	ctx                            _context.Context
	ApiService                     BucketSchemasApi
	bucketID                       string
	org                            *string
	orgID                          *string
	measurementSchemaCreateRequest *MeasurementSchemaCreateRequest
}

func (r ApiCreateMeasurementSchemaRequest) BucketID(bucketID string) ApiCreateMeasurementSchemaRequest {
	r.bucketID = bucketID
	return r
}
func (r ApiCreateMeasurementSchemaRequest) GetBucketID() string {
	return r.bucketID
}

func (r ApiCreateMeasurementSchemaRequest) Org(org string) ApiCreateMeasurementSchemaRequest {
	r.org = &org
	return r
}
func (r ApiCreateMeasurementSchemaRequest) GetOrg() *string {
	return r.org
}

func (r ApiCreateMeasurementSchemaRequest) OrgID(orgID string) ApiCreateMeasurementSchemaRequest {
	r.orgID = &orgID
	return r
}
func (r ApiCreateMeasurementSchemaRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiCreateMeasurementSchemaRequest) MeasurementSchemaCreateRequest(measurementSchemaCreateRequest MeasurementSchemaCreateRequest) ApiCreateMeasurementSchemaRequest {
	r.measurementSchemaCreateRequest = &measurementSchemaCreateRequest
	return r
}
func (r ApiCreateMeasurementSchemaRequest) GetMeasurementSchemaCreateRequest() *MeasurementSchemaCreateRequest {
	return r.measurementSchemaCreateRequest
}

func (r ApiCreateMeasurementSchemaRequest) Execute() (MeasurementSchema, error) {
	return r.ApiService.CreateMeasurementSchemaExecute(r)
}

/*
 * CreateMeasurementSchema Create a measurement schema for a bucket
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bucketID The identifier of the bucket.
 * @return ApiCreateMeasurementSchemaRequest
 */
func (a *BucketSchemasApiService) CreateMeasurementSchema(ctx _context.Context, bucketID string) ApiCreateMeasurementSchemaRequest {
	return ApiCreateMeasurementSchemaRequest{
		ApiService: a,
		ctx:        ctx,
		bucketID:   bucketID,
	}
}

/*
 * Execute executes the request
 * @return MeasurementSchema
 */
func (a *BucketSchemasApiService) CreateMeasurementSchemaExecute(r ApiCreateMeasurementSchemaRequest) (MeasurementSchema, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MeasurementSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketSchemasApiService.CreateMeasurementSchema")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/buckets/{bucketID}/schema/measurements"
	localVarPath = strings.Replace(localVarPath, "{"+"bucketID"+"}", _neturl.PathEscape(parameterToString(r.bucketID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.org != nil {
		localVarQueryParams.Add("org", parameterToString(*r.org, ""))
	}
	if r.orgID != nil {
		localVarQueryParams.Add("orgID", parameterToString(*r.orgID, ""))
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
	localVarPostBody = r.measurementSchemaCreateRequest
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
			error: _fmt.Sprintf("%s: code %s", errorPrefix, localVarHTTPResponse.Status),
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = _fmt.Sprintf("%s: %s", errorPrefix, err.Error())
				return localVarReturnValue, newErr
			}
			v.SetMessage(_fmt.Sprintf("%s: %s", errorPrefix, v.GetMessage()))
			newErr.model = &v
		}
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

type ApiGetMeasurementSchemaRequest struct {
	ctx           _context.Context
	ApiService    BucketSchemasApi
	bucketID      string
	measurementID string
	org           *string
	orgID         *string
}

func (r ApiGetMeasurementSchemaRequest) BucketID(bucketID string) ApiGetMeasurementSchemaRequest {
	r.bucketID = bucketID
	return r
}
func (r ApiGetMeasurementSchemaRequest) GetBucketID() string {
	return r.bucketID
}

func (r ApiGetMeasurementSchemaRequest) MeasurementID(measurementID string) ApiGetMeasurementSchemaRequest {
	r.measurementID = measurementID
	return r
}
func (r ApiGetMeasurementSchemaRequest) GetMeasurementID() string {
	return r.measurementID
}

func (r ApiGetMeasurementSchemaRequest) Org(org string) ApiGetMeasurementSchemaRequest {
	r.org = &org
	return r
}
func (r ApiGetMeasurementSchemaRequest) GetOrg() *string {
	return r.org
}

func (r ApiGetMeasurementSchemaRequest) OrgID(orgID string) ApiGetMeasurementSchemaRequest {
	r.orgID = &orgID
	return r
}
func (r ApiGetMeasurementSchemaRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiGetMeasurementSchemaRequest) Execute() (MeasurementSchema, error) {
	return r.ApiService.GetMeasurementSchemaExecute(r)
}

/*
 * GetMeasurementSchema Retrieve measurement schema information
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bucketID The identifier of the bucket.
 * @param measurementID The identifier of the measurement.
 * @return ApiGetMeasurementSchemaRequest
 */
func (a *BucketSchemasApiService) GetMeasurementSchema(ctx _context.Context, bucketID string, measurementID string) ApiGetMeasurementSchemaRequest {
	return ApiGetMeasurementSchemaRequest{
		ApiService:    a,
		ctx:           ctx,
		bucketID:      bucketID,
		measurementID: measurementID,
	}
}

/*
 * Execute executes the request
 * @return MeasurementSchema
 */
func (a *BucketSchemasApiService) GetMeasurementSchemaExecute(r ApiGetMeasurementSchemaRequest) (MeasurementSchema, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MeasurementSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketSchemasApiService.GetMeasurementSchema")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/buckets/{bucketID}/schema/measurements/{measurementID}"
	localVarPath = strings.Replace(localVarPath, "{"+"bucketID"+"}", _neturl.PathEscape(parameterToString(r.bucketID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"measurementID"+"}", _neturl.PathEscape(parameterToString(r.measurementID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.org != nil {
		localVarQueryParams.Add("org", parameterToString(*r.org, ""))
	}
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
			error: _fmt.Sprintf("%s: code %s", errorPrefix, localVarHTTPResponse.Status),
		}
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

type ApiGetMeasurementSchemasRequest struct {
	ctx        _context.Context
	ApiService BucketSchemasApi
	bucketID   string
	org        *string
	orgID      *string
	name       *string
}

func (r ApiGetMeasurementSchemasRequest) BucketID(bucketID string) ApiGetMeasurementSchemasRequest {
	r.bucketID = bucketID
	return r
}
func (r ApiGetMeasurementSchemasRequest) GetBucketID() string {
	return r.bucketID
}

func (r ApiGetMeasurementSchemasRequest) Org(org string) ApiGetMeasurementSchemasRequest {
	r.org = &org
	return r
}
func (r ApiGetMeasurementSchemasRequest) GetOrg() *string {
	return r.org
}

func (r ApiGetMeasurementSchemasRequest) OrgID(orgID string) ApiGetMeasurementSchemasRequest {
	r.orgID = &orgID
	return r
}
func (r ApiGetMeasurementSchemasRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiGetMeasurementSchemasRequest) Name(name string) ApiGetMeasurementSchemasRequest {
	r.name = &name
	return r
}
func (r ApiGetMeasurementSchemasRequest) GetName() *string {
	return r.name
}

func (r ApiGetMeasurementSchemasRequest) Execute() (MeasurementSchemaList, error) {
	return r.ApiService.GetMeasurementSchemasExecute(r)
}

/*
 * GetMeasurementSchemas List all measurement schemas of a bucket
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bucketID The identifier of the bucket.
 * @return ApiGetMeasurementSchemasRequest
 */
func (a *BucketSchemasApiService) GetMeasurementSchemas(ctx _context.Context, bucketID string) ApiGetMeasurementSchemasRequest {
	return ApiGetMeasurementSchemasRequest{
		ApiService: a,
		ctx:        ctx,
		bucketID:   bucketID,
	}
}

/*
 * Execute executes the request
 * @return MeasurementSchemaList
 */
func (a *BucketSchemasApiService) GetMeasurementSchemasExecute(r ApiGetMeasurementSchemasRequest) (MeasurementSchemaList, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MeasurementSchemaList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketSchemasApiService.GetMeasurementSchemas")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/buckets/{bucketID}/schema/measurements"
	localVarPath = strings.Replace(localVarPath, "{"+"bucketID"+"}", _neturl.PathEscape(parameterToString(r.bucketID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.org != nil {
		localVarQueryParams.Add("org", parameterToString(*r.org, ""))
	}
	if r.orgID != nil {
		localVarQueryParams.Add("orgID", parameterToString(*r.orgID, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
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
			error: _fmt.Sprintf("%s: code %s", errorPrefix, localVarHTTPResponse.Status),
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = _fmt.Sprintf("%s: %s", errorPrefix, err.Error())
				return localVarReturnValue, newErr
			}
			v.SetMessage(_fmt.Sprintf("%s: %s", errorPrefix, v.GetMessage()))
			newErr.model = &v
		}
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

type ApiUpdateMeasurementSchemaRequest struct {
	ctx                            _context.Context
	ApiService                     BucketSchemasApi
	bucketID                       string
	measurementID                  string
	org                            *string
	orgID                          *string
	measurementSchemaUpdateRequest *MeasurementSchemaUpdateRequest
}

func (r ApiUpdateMeasurementSchemaRequest) BucketID(bucketID string) ApiUpdateMeasurementSchemaRequest {
	r.bucketID = bucketID
	return r
}
func (r ApiUpdateMeasurementSchemaRequest) GetBucketID() string {
	return r.bucketID
}

func (r ApiUpdateMeasurementSchemaRequest) MeasurementID(measurementID string) ApiUpdateMeasurementSchemaRequest {
	r.measurementID = measurementID
	return r
}
func (r ApiUpdateMeasurementSchemaRequest) GetMeasurementID() string {
	return r.measurementID
}

func (r ApiUpdateMeasurementSchemaRequest) Org(org string) ApiUpdateMeasurementSchemaRequest {
	r.org = &org
	return r
}
func (r ApiUpdateMeasurementSchemaRequest) GetOrg() *string {
	return r.org
}

func (r ApiUpdateMeasurementSchemaRequest) OrgID(orgID string) ApiUpdateMeasurementSchemaRequest {
	r.orgID = &orgID
	return r
}
func (r ApiUpdateMeasurementSchemaRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiUpdateMeasurementSchemaRequest) MeasurementSchemaUpdateRequest(measurementSchemaUpdateRequest MeasurementSchemaUpdateRequest) ApiUpdateMeasurementSchemaRequest {
	r.measurementSchemaUpdateRequest = &measurementSchemaUpdateRequest
	return r
}
func (r ApiUpdateMeasurementSchemaRequest) GetMeasurementSchemaUpdateRequest() *MeasurementSchemaUpdateRequest {
	return r.measurementSchemaUpdateRequest
}

func (r ApiUpdateMeasurementSchemaRequest) Execute() (MeasurementSchema, error) {
	return r.ApiService.UpdateMeasurementSchemaExecute(r)
}

/*
 * UpdateMeasurementSchema Update a measurement schema
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bucketID The identifier of the bucket.
 * @param measurementID The identifier of the measurement.
 * @return ApiUpdateMeasurementSchemaRequest
 */
func (a *BucketSchemasApiService) UpdateMeasurementSchema(ctx _context.Context, bucketID string, measurementID string) ApiUpdateMeasurementSchemaRequest {
	return ApiUpdateMeasurementSchemaRequest{
		ApiService:    a,
		ctx:           ctx,
		bucketID:      bucketID,
		measurementID: measurementID,
	}
}

/*
 * Execute executes the request
 * @return MeasurementSchema
 */
func (a *BucketSchemasApiService) UpdateMeasurementSchemaExecute(r ApiUpdateMeasurementSchemaRequest) (MeasurementSchema, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MeasurementSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketSchemasApiService.UpdateMeasurementSchema")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/buckets/{bucketID}/schema/measurements/{measurementID}"
	localVarPath = strings.Replace(localVarPath, "{"+"bucketID"+"}", _neturl.PathEscape(parameterToString(r.bucketID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"measurementID"+"}", _neturl.PathEscape(parameterToString(r.measurementID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.org != nil {
		localVarQueryParams.Add("org", parameterToString(*r.org, ""))
	}
	if r.orgID != nil {
		localVarQueryParams.Add("orgID", parameterToString(*r.orgID, ""))
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
	localVarPostBody = r.measurementSchemaUpdateRequest
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
			error: _fmt.Sprintf("%s: code %s", errorPrefix, localVarHTTPResponse.Status),
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = _fmt.Sprintf("%s: %s", errorPrefix, err.Error())
				return localVarReturnValue, newErr
			}
			v.SetMessage(_fmt.Sprintf("%s: %s", errorPrefix, v.GetMessage()))
			newErr.model = &v
		}
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

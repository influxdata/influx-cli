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
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

type DeleteApi interface {

	/*
	 * PostDelete Delete time series data from InfluxDB
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiPostDeleteRequest
	 */
	PostDelete(ctx _context.Context) ApiPostDeleteRequest

	/*
	 * PostDeleteExecute executes the request
	 */
	PostDeleteExecute(r ApiPostDeleteRequest) error
}

// DeleteApiService DeleteApi service
type DeleteApiService service

type ApiPostDeleteRequest struct {
	ctx                    _context.Context
	ApiService             DeleteApi
	deletePredicateRequest *DeletePredicateRequest
	zapTraceSpan           *string
	org                    *string
	bucket                 *string
	orgID                  *string
	bucketID               *string
}

func (r ApiPostDeleteRequest) DeletePredicateRequest(deletePredicateRequest DeletePredicateRequest) ApiPostDeleteRequest {
	r.deletePredicateRequest = &deletePredicateRequest
	return r
}
func (r ApiPostDeleteRequest) GetDeletePredicateRequest() *DeletePredicateRequest {
	return r.deletePredicateRequest
}

func (r ApiPostDeleteRequest) ZapTraceSpan(zapTraceSpan string) ApiPostDeleteRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPostDeleteRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPostDeleteRequest) Org(org string) ApiPostDeleteRequest {
	r.org = &org
	return r
}
func (r ApiPostDeleteRequest) GetOrg() *string {
	return r.org
}

func (r ApiPostDeleteRequest) Bucket(bucket string) ApiPostDeleteRequest {
	r.bucket = &bucket
	return r
}
func (r ApiPostDeleteRequest) GetBucket() *string {
	return r.bucket
}

func (r ApiPostDeleteRequest) OrgID(orgID string) ApiPostDeleteRequest {
	r.orgID = &orgID
	return r
}
func (r ApiPostDeleteRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiPostDeleteRequest) BucketID(bucketID string) ApiPostDeleteRequest {
	r.bucketID = &bucketID
	return r
}
func (r ApiPostDeleteRequest) GetBucketID() *string {
	return r.bucketID
}

func (r ApiPostDeleteRequest) Execute() error {
	return r.ApiService.PostDeleteExecute(r)
}

/*
 * PostDelete Delete time series data from InfluxDB
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostDeleteRequest
 */
func (a *DeleteApiService) PostDelete(ctx _context.Context) ApiPostDeleteRequest {
	return ApiPostDeleteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *DeleteApiService) PostDeleteExecute(r ApiPostDeleteRequest) error {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeleteApiService.PostDelete")
	if err != nil {
		return GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.deletePredicateRequest == nil {
		return reportError("deletePredicateRequest is required and must be specified")
	}

	if r.org != nil {
		localVarQueryParams.Add("org", parameterToString(*r.org, ""))
	}
	if r.bucket != nil {
		localVarQueryParams.Add("bucket", parameterToString(*r.bucket, ""))
	}
	if r.orgID != nil {
		localVarQueryParams.Add("orgID", parameterToString(*r.orgID, ""))
	}
	if r.bucketID != nil {
		localVarQueryParams.Add("bucketID", parameterToString(*r.bucketID, ""))
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
	localVarPostBody = r.deletePredicateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return err
		}
		localVarBody, err := _ioutil.ReadAll(body)
		body.Close()
		if err != nil {
			return err
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return newErr
			}
			newErr.model = &v
			return newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return newErr
			}
			newErr.model = &v
			return newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return newErr
			}
			newErr.model = &v
			return newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return newErr
		}
		newErr.model = &v
		return newErr
	}

	return nil
}

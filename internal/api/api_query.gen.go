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
	_gzip "compress/gzip"
	_context "context"
	_io "io"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

type QueryApi interface {

	/*
	 * PostQuery Query InfluxDB
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiPostQueryRequest
	 */
	PostQuery(ctx _context.Context) ApiPostQueryRequest

	/*
	 * PostQueryExecute executes the request
	 * @return *os.File
	 */
	PostQueryExecute(r ApiPostQueryRequest) (_io.ReadCloser, error)
}

// queryApiGzipReadCloser supports streaming gzip response-bodies directly from the server.
type queryApiGzipReadCloser struct {
	underlying _io.ReadCloser
	gzip       _io.ReadCloser
}

func (gzrc *queryApiGzipReadCloser) Read(p []byte) (int, error) {
	return gzrc.gzip.Read(p)
}
func (gzrc *queryApiGzipReadCloser) Close() error {
	if err := gzrc.gzip.Close(); err != nil {
		return err
	}
	return gzrc.underlying.Close()
}

// QueryApiService QueryApi service
type QueryApiService service

type ApiPostQueryRequest struct {
	ctx            _context.Context
	ApiService     QueryApi
	zapTraceSpan   *string
	acceptEncoding *string
	contentType    *string
	org            *string
	orgID          *string
	query          *Query
}

func (r ApiPostQueryRequest) ZapTraceSpan(zapTraceSpan string) ApiPostQueryRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPostQueryRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPostQueryRequest) AcceptEncoding(acceptEncoding string) ApiPostQueryRequest {
	r.acceptEncoding = &acceptEncoding
	return r
}
func (r ApiPostQueryRequest) GetAcceptEncoding() *string {
	return r.acceptEncoding
}

func (r ApiPostQueryRequest) ContentType(contentType string) ApiPostQueryRequest {
	r.contentType = &contentType
	return r
}
func (r ApiPostQueryRequest) GetContentType() *string {
	return r.contentType
}

func (r ApiPostQueryRequest) Org(org string) ApiPostQueryRequest {
	r.org = &org
	return r
}
func (r ApiPostQueryRequest) GetOrg() *string {
	return r.org
}

func (r ApiPostQueryRequest) OrgID(orgID string) ApiPostQueryRequest {
	r.orgID = &orgID
	return r
}
func (r ApiPostQueryRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiPostQueryRequest) Query(query Query) ApiPostQueryRequest {
	r.query = &query
	return r
}
func (r ApiPostQueryRequest) GetQuery() *Query {
	return r.query
}

func (r ApiPostQueryRequest) Execute() (_io.ReadCloser, error) {
	return r.ApiService.PostQueryExecute(r)
}

/*
 * PostQuery Query InfluxDB
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostQueryRequest
 */
func (a *QueryApiService) PostQuery(ctx _context.Context) ApiPostQueryRequest {
	return ApiPostQueryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return *os.File
 */
func (a *QueryApiService) PostQueryExecute(r ApiPostQueryRequest) (_io.ReadCloser, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  _io.ReadCloser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryApiService.PostQuery")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/query"

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
	localVarHTTPHeaderAccepts := []string{"text/csv", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.zapTraceSpan != nil {
		localVarHeaderParams["Zap-Trace-Span"] = parameterToString(*r.zapTraceSpan, "")
	}
	if r.acceptEncoding != nil {
		localVarHeaderParams["Accept-Encoding"] = parameterToString(*r.acceptEncoding, "")
	}
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
	}
	// body params
	localVarPostBody = r.query
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
	}

	var body _io.ReadCloser = localVarHTTPResponse.Body
	if localVarHTTPResponse.Header.Get("Content-Encoding") == "gzip" {
		gzr, err := _gzip.NewReader(body)
		if err != nil {
			body.Close()
			return localVarReturnValue, err
		}
		body = &queryApiGzipReadCloser{underlying: body, gzip: gzr}
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, err := _ioutil.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarReturnValue, err
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, newErr
		}
		newErr.model = &v
		return localVarReturnValue, newErr
	}

	localVarReturnValue = body

	return localVarReturnValue, nil
}

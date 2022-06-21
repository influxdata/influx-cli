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

type LegacyQueryApi interface {

	/*
	 * GetLegacyQuery Query with the 1.x compatibility API
	 * Queries InfluxDB using InfluxQL.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiGetLegacyQueryRequest
	 */
	GetLegacyQuery(ctx _context.Context) ApiGetLegacyQueryRequest

	/*
	 * GetLegacyQueryExecute executes the request
	 * @return string
	 */
	GetLegacyQueryExecute(r ApiGetLegacyQueryRequest) (string, error)

	/*
	 * GetLegacyQueryExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 * @return string
	 */
	GetLegacyQueryExecuteWithHttpInfo(r ApiGetLegacyQueryRequest) (string, *_nethttp.Response, error)
}

// LegacyQueryApiService LegacyQueryApi service
type LegacyQueryApiService service

type ApiGetLegacyQueryRequest struct {
	ctx            _context.Context
	ApiService     LegacyQueryApi
	db             *string
	q              *string
	zapTraceSpan   *string
	accept         *string
	acceptEncoding *string
	contentType    *string
	u              *string
	p              *string
	rp             *string
	epoch          *string
}

func (r ApiGetLegacyQueryRequest) Db(db string) ApiGetLegacyQueryRequest {
	r.db = &db
	return r
}
func (r ApiGetLegacyQueryRequest) GetDb() *string {
	return r.db
}

func (r ApiGetLegacyQueryRequest) Q(q string) ApiGetLegacyQueryRequest {
	r.q = &q
	return r
}
func (r ApiGetLegacyQueryRequest) GetQ() *string {
	return r.q
}

func (r ApiGetLegacyQueryRequest) ZapTraceSpan(zapTraceSpan string) ApiGetLegacyQueryRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetLegacyQueryRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetLegacyQueryRequest) Accept(accept string) ApiGetLegacyQueryRequest {
	r.accept = &accept
	return r
}
func (r ApiGetLegacyQueryRequest) GetAccept() *string {
	return r.accept
}

func (r ApiGetLegacyQueryRequest) AcceptEncoding(acceptEncoding string) ApiGetLegacyQueryRequest {
	r.acceptEncoding = &acceptEncoding
	return r
}
func (r ApiGetLegacyQueryRequest) GetAcceptEncoding() *string {
	return r.acceptEncoding
}

func (r ApiGetLegacyQueryRequest) ContentType(contentType string) ApiGetLegacyQueryRequest {
	r.contentType = &contentType
	return r
}
func (r ApiGetLegacyQueryRequest) GetContentType() *string {
	return r.contentType
}

func (r ApiGetLegacyQueryRequest) U(u string) ApiGetLegacyQueryRequest {
	r.u = &u
	return r
}
func (r ApiGetLegacyQueryRequest) GetU() *string {
	return r.u
}

func (r ApiGetLegacyQueryRequest) P(p string) ApiGetLegacyQueryRequest {
	r.p = &p
	return r
}
func (r ApiGetLegacyQueryRequest) GetP() *string {
	return r.p
}

func (r ApiGetLegacyQueryRequest) Rp(rp string) ApiGetLegacyQueryRequest {
	r.rp = &rp
	return r
}
func (r ApiGetLegacyQueryRequest) GetRp() *string {
	return r.rp
}

func (r ApiGetLegacyQueryRequest) Epoch(epoch string) ApiGetLegacyQueryRequest {
	r.epoch = &epoch
	return r
}
func (r ApiGetLegacyQueryRequest) GetEpoch() *string {
	return r.epoch
}

func (r ApiGetLegacyQueryRequest) Execute() (string, error) {
	return r.ApiService.GetLegacyQueryExecute(r)
}

func (r ApiGetLegacyQueryRequest) ExecuteWithHttpInfo() (string, *_nethttp.Response, error) {
	return r.ApiService.GetLegacyQueryExecuteWithHttpInfo(r)
}

/*
 * GetLegacyQuery Query with the 1.x compatibility API
 * Queries InfluxDB using InfluxQL.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetLegacyQueryRequest
 */
func (a *LegacyQueryApiService) GetLegacyQuery(ctx _context.Context) ApiGetLegacyQueryRequest {
	return ApiGetLegacyQueryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return string
 */
func (a *LegacyQueryApiService) GetLegacyQueryExecute(r ApiGetLegacyQueryRequest) (string, error) {
	returnVal, _, err := a.GetLegacyQueryExecuteWithHttpInfo(r)
	return returnVal, err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 * @return string
 */
func (a *LegacyQueryApiService) GetLegacyQueryExecuteWithHttpInfo(r ApiGetLegacyQueryRequest) (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyQueryApiService.GetLegacyQuery")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.db == nil {
		return localVarReturnValue, nil, reportError("db is required and must be specified")
	}
	if r.q == nil {
		return localVarReturnValue, nil, reportError("q is required and must be specified")
	}

	if r.u != nil {
		localVarQueryParams.Add("u", parameterToString(*r.u, ""))
	}
	if r.p != nil {
		localVarQueryParams.Add("p", parameterToString(*r.p, ""))
	}
	localVarQueryParams.Add("db", parameterToString(*r.db, ""))
	if r.rp != nil {
		localVarQueryParams.Add("rp", parameterToString(*r.rp, ""))
	}
	localVarQueryParams.Add("q", parameterToString(*r.q, ""))
	if r.epoch != nil {
		localVarQueryParams.Add("epoch", parameterToString(*r.epoch, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/csv", "application/csv", "application/json", "application/x-msgpack"}

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
	if r.acceptEncoding != nil {
		localVarHeaderParams["Accept-Encoding"] = parameterToString(*r.acceptEncoding, "")
	}
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
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

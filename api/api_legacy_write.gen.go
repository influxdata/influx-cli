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

type LegacyWriteApi interface {

	/*
	 * PostLegacyWrite Write time series data into InfluxDB in a V1-compatible format
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiPostLegacyWriteRequest
	 */
	PostLegacyWrite(ctx _context.Context) ApiPostLegacyWriteRequest

	/*
	 * PostLegacyWriteExecute executes the request
	 */
	PostLegacyWriteExecute(r ApiPostLegacyWriteRequest) error

	/*
	 * PostLegacyWriteExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not
	 * available on the returned HTTP response as it will have already been read and closed; access to the response body
	 * content should be achieved through the returned response model if applicable.
	 */
	PostLegacyWriteExecuteWithHttpInfo(r ApiPostLegacyWriteRequest) (*_nethttp.Response, error)
}

// LegacyWriteApiService LegacyWriteApi service
type LegacyWriteApiService service

type ApiPostLegacyWriteRequest struct {
	ctx             _context.Context
	ApiService      LegacyWriteApi
	db              *string
	body            *string
	zapTraceSpan    *string
	u               *string
	p               *string
	rp              *string
	precision       *string
	contentEncoding *string
}

func (r ApiPostLegacyWriteRequest) Db(db string) ApiPostLegacyWriteRequest {
	r.db = &db
	return r
}
func (r ApiPostLegacyWriteRequest) GetDb() *string {
	return r.db
}

func (r ApiPostLegacyWriteRequest) Body(body string) ApiPostLegacyWriteRequest {
	r.body = &body
	return r
}
func (r ApiPostLegacyWriteRequest) GetBody() *string {
	return r.body
}

func (r ApiPostLegacyWriteRequest) ZapTraceSpan(zapTraceSpan string) ApiPostLegacyWriteRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPostLegacyWriteRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPostLegacyWriteRequest) U(u string) ApiPostLegacyWriteRequest {
	r.u = &u
	return r
}
func (r ApiPostLegacyWriteRequest) GetU() *string {
	return r.u
}

func (r ApiPostLegacyWriteRequest) P(p string) ApiPostLegacyWriteRequest {
	r.p = &p
	return r
}
func (r ApiPostLegacyWriteRequest) GetP() *string {
	return r.p
}

func (r ApiPostLegacyWriteRequest) Rp(rp string) ApiPostLegacyWriteRequest {
	r.rp = &rp
	return r
}
func (r ApiPostLegacyWriteRequest) GetRp() *string {
	return r.rp
}

func (r ApiPostLegacyWriteRequest) Precision(precision string) ApiPostLegacyWriteRequest {
	r.precision = &precision
	return r
}
func (r ApiPostLegacyWriteRequest) GetPrecision() *string {
	return r.precision
}

func (r ApiPostLegacyWriteRequest) ContentEncoding(contentEncoding string) ApiPostLegacyWriteRequest {
	r.contentEncoding = &contentEncoding
	return r
}
func (r ApiPostLegacyWriteRequest) GetContentEncoding() *string {
	return r.contentEncoding
}

func (r ApiPostLegacyWriteRequest) Execute() error {
	return r.ApiService.PostLegacyWriteExecute(r)
}

func (r ApiPostLegacyWriteRequest) ExecuteWithHttpInfo() (*_nethttp.Response, error) {
	return r.ApiService.PostLegacyWriteExecuteWithHttpInfo(r)
}

/*
 * PostLegacyWrite Write time series data into InfluxDB in a V1-compatible format
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostLegacyWriteRequest
 */
func (a *LegacyWriteApiService) PostLegacyWrite(ctx _context.Context) ApiPostLegacyWriteRequest {
	return ApiPostLegacyWriteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *LegacyWriteApiService) PostLegacyWriteExecute(r ApiPostLegacyWriteRequest) error {
	_, err := a.PostLegacyWriteExecuteWithHttpInfo(r)
	return err
}

/*
 * ExecuteWithHttpInfo executes the request with HTTP response info returned. The response body is not available on the
 * returned HTTP response as it will have already been read and closed; access to the response body content should be
 * achieved through the returned response model if applicable.
 */
func (a *LegacyWriteApiService) PostLegacyWriteExecuteWithHttpInfo(r ApiPostLegacyWriteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyWriteApiService.PostLegacyWrite")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/write"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.db == nil {
		return nil, reportError("db is required and must be specified")
	}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
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
	if r.precision != nil {
		localVarQueryParams.Add("precision", parameterToString(*r.precision, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

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
	if r.contentEncoding != nil {
		localVarHeaderParams["Content-Encoding"] = parameterToString(*r.contentEncoding, "")
	}
	// body params
	localVarPostBody = r.body
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v LineProtocolError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
				return localVarHTTPResponse, newErr
			}
			v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
			newErr.model = &v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
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
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 413 {
			var v LineProtocolLengthError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
				return localVarHTTPResponse, newErr
			}
			v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
			newErr.model = &v
			return localVarHTTPResponse, newErr
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

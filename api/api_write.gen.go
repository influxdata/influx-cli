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

type WriteApi interface {

	/*
			 * PostWrite Write data
			 * Writes data to a bucket.

		To write data into InfluxDB, you need the following:
		- **organization** – _See [View organizations]({{% INFLUXDB_DOCS_URL %}}/organizations/view-orgs/#view-your-organization-id) for instructions on viewing your organization ID._
		- **bucket** – _See [View buckets]({{% INFLUXDB_DOCS_URL %}}/organizations/buckets/view-buckets/) for
		 instructions on viewing your bucket ID._
		- **API token** – _See [View tokens]({{% INFLUXDB_DOCS_URL %}}/security/tokens/view-tokens/)
		 for instructions on viewing your API token._
		- **InfluxDB URL** – _See [InfluxDB URLs]({{% INFLUXDB_DOCS_URL %}}/reference/urls/)_.
		- data in [line protocol]({{% INFLUXDB_DOCS_URL %}}/reference/syntax/line-protocol) format.

		For more information and examples, see [Write data with the InfluxDB API]({{% INFLUXDB_DOCS_URL %}}/write-data/developer-tools/api).

			 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 * @return ApiPostWriteRequest
	*/
	PostWrite(ctx _context.Context) ApiPostWriteRequest

	/*
	 * PostWriteExecute executes the request
	 */
	PostWriteExecute(r ApiPostWriteRequest) error

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on OSS
	// servers.
	OnlyOSS() WriteApi

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on cloud
	// servers.
	OnlyCloud() WriteApi
}

// WriteApiService WriteApi service
type WriteApiService service

func (a *WriteApiService) OnlyOSS() WriteApi {
	a.isOnlyOSS = true
	return a
}

func (a *WriteApiService) OnlyCloud() WriteApi {
	a.isOnlyCloud = true
	return a
}

type ApiPostWriteRequest struct {
	ctx             _context.Context
	ApiService      WriteApi
	org             *string
	bucket          *string
	body            []byte
	zapTraceSpan    *string
	contentEncoding *string
	contentType     *string
	contentLength   *int32
	accept          *string
	orgID           *string
	precision       *WritePrecision
}

func (r ApiPostWriteRequest) Org(org string) ApiPostWriteRequest {
	r.org = &org
	return r
}
func (r ApiPostWriteRequest) GetOrg() *string {
	return r.org
}

func (r ApiPostWriteRequest) Bucket(bucket string) ApiPostWriteRequest {
	r.bucket = &bucket
	return r
}
func (r ApiPostWriteRequest) GetBucket() *string {
	return r.bucket
}

func (r ApiPostWriteRequest) Body(body []byte) ApiPostWriteRequest {
	r.body = body
	return r
}
func (r ApiPostWriteRequest) GetBody() []byte {
	return r.body
}

func (r ApiPostWriteRequest) ZapTraceSpan(zapTraceSpan string) ApiPostWriteRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPostWriteRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPostWriteRequest) ContentEncoding(contentEncoding string) ApiPostWriteRequest {
	r.contentEncoding = &contentEncoding
	return r
}
func (r ApiPostWriteRequest) GetContentEncoding() *string {
	return r.contentEncoding
}

func (r ApiPostWriteRequest) ContentType(contentType string) ApiPostWriteRequest {
	r.contentType = &contentType
	return r
}
func (r ApiPostWriteRequest) GetContentType() *string {
	return r.contentType
}

func (r ApiPostWriteRequest) ContentLength(contentLength int32) ApiPostWriteRequest {
	r.contentLength = &contentLength
	return r
}
func (r ApiPostWriteRequest) GetContentLength() *int32 {
	return r.contentLength
}

func (r ApiPostWriteRequest) Accept(accept string) ApiPostWriteRequest {
	r.accept = &accept
	return r
}
func (r ApiPostWriteRequest) GetAccept() *string {
	return r.accept
}

func (r ApiPostWriteRequest) OrgID(orgID string) ApiPostWriteRequest {
	r.orgID = &orgID
	return r
}
func (r ApiPostWriteRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiPostWriteRequest) Precision(precision WritePrecision) ApiPostWriteRequest {
	r.precision = &precision
	return r
}
func (r ApiPostWriteRequest) GetPrecision() *WritePrecision {
	return r.precision
}

func (r ApiPostWriteRequest) Execute() error {
	return r.ApiService.PostWriteExecute(r)
}

/*
 * PostWrite Write data
 * Writes data to a bucket.

To write data into InfluxDB, you need the following:
- **organization** – _See [View organizations]({{% INFLUXDB_DOCS_URL %}}/organizations/view-orgs/#view-your-organization-id) for instructions on viewing your organization ID._
- **bucket** – _See [View buckets]({{% INFLUXDB_DOCS_URL %}}/organizations/buckets/view-buckets/) for
 instructions on viewing your bucket ID._
- **API token** – _See [View tokens]({{% INFLUXDB_DOCS_URL %}}/security/tokens/view-tokens/)
 for instructions on viewing your API token._
- **InfluxDB URL** – _See [InfluxDB URLs]({{% INFLUXDB_DOCS_URL %}}/reference/urls/)_.
- data in [line protocol]({{% INFLUXDB_DOCS_URL %}}/reference/syntax/line-protocol) format.

For more information and examples, see [Write data with the InfluxDB API]({{% INFLUXDB_DOCS_URL %}}/write-data/developer-tools/api).

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostWriteRequest
*/
func (a *WriteApiService) PostWrite(ctx _context.Context) ApiPostWriteRequest {
	return ApiPostWriteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *WriteApiService) PostWriteExecute(r ApiPostWriteRequest) error {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WriteApiService.PostWrite")
	if err != nil {
		return GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/write"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.org == nil {
		return reportError("org is required and must be specified")
	}
	if r.bucket == nil {
		return reportError("bucket is required and must be specified")
	}
	if r.body == nil {
		return reportError("body is required and must be specified")
	}

	localVarQueryParams.Add("org", parameterToString(*r.org, ""))
	if r.orgID != nil {
		localVarQueryParams.Add("orgID", parameterToString(*r.orgID, ""))
	}
	localVarQueryParams.Add("bucket", parameterToString(*r.bucket, ""))
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
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
	}
	if r.contentLength != nil {
		localVarHeaderParams["Content-Length"] = parameterToString(*r.contentLength, "")
	}
	if r.accept != nil {
		localVarHeaderParams["Accept"] = parameterToString(*r.accept, "")
	}
	// body params
	localVarPostBody = r.body
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
		errorPrefix = "InfluxDB OSS-only command failed: "
	} else if a.isOnlyCloud {
		errorPrefix = "InfluxDB Cloud-only command failed: "
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return _fmt.Errorf("%s%w", errorPrefix, err)
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			return _fmt.Errorf("%s%w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, localVarHTTPResponse.Status),
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v LineProtocolError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
				return newErr
			}
			v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
			newErr.model = &v
			return newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
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
		if localVarHTTPResponse.StatusCode == 413 {
			var v LineProtocolLengthError
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

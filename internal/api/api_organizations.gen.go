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
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

type OrganizationsApi interface {

	/*
	 * GetOrgs List all organizations
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiGetOrgsRequest
	 */
	GetOrgs(ctx _context.Context) ApiGetOrgsRequest

	/*
	 * GetOrgsExecute executes the request
	 * @return Organizations
	 */
	GetOrgsExecute(r ApiGetOrgsRequest) (Organizations, *_nethttp.Response, error)

	/*
	 * PostOrgs Create an organization
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiPostOrgsRequest
	 */
	PostOrgs(ctx _context.Context) ApiPostOrgsRequest

	/*
	 * PostOrgsExecute executes the request
	 * @return Organization
	 */
	PostOrgsExecute(r ApiPostOrgsRequest) (Organization, *_nethttp.Response, error)
}

// OrganizationsApiService OrganizationsApi service
type OrganizationsApiService service

type ApiGetOrgsRequest struct {
	ctx          _context.Context
	ApiService   OrganizationsApi
	zapTraceSpan *string
	offset       *int32
	limit        *int32
	descending   *bool
	org          *string
	orgID        *string
	userID       *string
}

func (r ApiGetOrgsRequest) ZapTraceSpan(zapTraceSpan string) ApiGetOrgsRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetOrgsRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetOrgsRequest) Offset(offset int32) ApiGetOrgsRequest {
	r.offset = &offset
	return r
}
func (r ApiGetOrgsRequest) GetOffset() *int32 {
	return r.offset
}

func (r ApiGetOrgsRequest) Limit(limit int32) ApiGetOrgsRequest {
	r.limit = &limit
	return r
}
func (r ApiGetOrgsRequest) GetLimit() *int32 {
	return r.limit
}

func (r ApiGetOrgsRequest) Descending(descending bool) ApiGetOrgsRequest {
	r.descending = &descending
	return r
}
func (r ApiGetOrgsRequest) GetDescending() *bool {
	return r.descending
}

func (r ApiGetOrgsRequest) Org(org string) ApiGetOrgsRequest {
	r.org = &org
	return r
}
func (r ApiGetOrgsRequest) GetOrg() *string {
	return r.org
}

func (r ApiGetOrgsRequest) OrgID(orgID string) ApiGetOrgsRequest {
	r.orgID = &orgID
	return r
}
func (r ApiGetOrgsRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiGetOrgsRequest) UserID(userID string) ApiGetOrgsRequest {
	r.userID = &userID
	return r
}
func (r ApiGetOrgsRequest) GetUserID() *string {
	return r.userID
}

func (r ApiGetOrgsRequest) Execute() (Organizations, *_nethttp.Response, error) {
	return r.ApiService.GetOrgsExecute(r)
}

/*
 * GetOrgs List all organizations
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetOrgsRequest
 */
func (a *OrganizationsApiService) GetOrgs(ctx _context.Context) ApiGetOrgsRequest {
	return ApiGetOrgsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Organizations
 */
func (a *OrganizationsApiService) GetOrgsExecute(r ApiGetOrgsRequest) (Organizations, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Organizations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.GetOrgs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.descending != nil {
		localVarQueryParams.Add("descending", parameterToString(*r.descending, ""))
	}
	if r.org != nil {
		localVarQueryParams.Add("org", parameterToString(*r.org, ""))
	}
	if r.orgID != nil {
		localVarQueryParams.Add("orgID", parameterToString(*r.orgID, ""))
	}
	if r.userID != nil {
		localVarQueryParams.Add("userID", parameterToString(*r.userID, ""))
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
		localVarHTTPResponse.Body.Close()
		localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
		if err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = &v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostOrgsRequest struct {
	ctx          _context.Context
	ApiService   OrganizationsApi
	organization *Organization
	zapTraceSpan *string
}

func (r ApiPostOrgsRequest) Organization(organization Organization) ApiPostOrgsRequest {
	r.organization = &organization
	return r
}
func (r ApiPostOrgsRequest) GetOrganization() *Organization {
	return r.organization
}

func (r ApiPostOrgsRequest) ZapTraceSpan(zapTraceSpan string) ApiPostOrgsRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPostOrgsRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPostOrgsRequest) Execute() (Organization, *_nethttp.Response, error) {
	return r.ApiService.PostOrgsExecute(r)
}

/*
 * PostOrgs Create an organization
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostOrgsRequest
 */
func (a *OrganizationsApiService) PostOrgs(ctx _context.Context) ApiPostOrgsRequest {
	return ApiPostOrgsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Organization
 */
func (a *OrganizationsApiService) PostOrgsExecute(r ApiPostOrgsRequest) (Organization, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Organization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationsApiService.PostOrgs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.organization == nil {
		return localVarReturnValue, nil, reportError("organization is required and must be specified")
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
	localVarPostBody = r.organization
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
		localVarHTTPResponse.Body.Close()
		localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
		if err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = &v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

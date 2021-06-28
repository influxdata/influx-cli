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

type LegacyAuthorizationsApi interface {

	/*
	 * DeleteLegacyAuthorizationsID Delete a legacy authorization
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param authID The ID of the legacy authorization to delete.
	 * @return ApiDeleteLegacyAuthorizationsIDRequest
	 */
	DeleteLegacyAuthorizationsID(ctx _context.Context, authID string) ApiDeleteLegacyAuthorizationsIDRequest

	/*
	 * DeleteLegacyAuthorizationsIDExecute executes the request
	 */
	DeleteLegacyAuthorizationsIDExecute(r ApiDeleteLegacyAuthorizationsIDRequest) error

	/*
	 * GetLegacyAuthorizations List all legacy authorizations
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiGetLegacyAuthorizationsRequest
	 */
	GetLegacyAuthorizations(ctx _context.Context) ApiGetLegacyAuthorizationsRequest

	/*
	 * GetLegacyAuthorizationsExecute executes the request
	 * @return Authorizations
	 */
	GetLegacyAuthorizationsExecute(r ApiGetLegacyAuthorizationsRequest) (Authorizations, error)

	/*
	 * GetLegacyAuthorizationsID Retrieve a legacy authorization
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param authID The ID of the legacy authorization to get.
	 * @return ApiGetLegacyAuthorizationsIDRequest
	 */
	GetLegacyAuthorizationsID(ctx _context.Context, authID string) ApiGetLegacyAuthorizationsIDRequest

	/*
	 * GetLegacyAuthorizationsIDExecute executes the request
	 * @return Authorization
	 */
	GetLegacyAuthorizationsIDExecute(r ApiGetLegacyAuthorizationsIDRequest) (Authorization, error)

	/*
	 * PatchLegacyAuthorizationsID Update a legacy authorization to be active or inactive
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param authID The ID of the legacy authorization to update.
	 * @return ApiPatchLegacyAuthorizationsIDRequest
	 */
	PatchLegacyAuthorizationsID(ctx _context.Context, authID string) ApiPatchLegacyAuthorizationsIDRequest

	/*
	 * PatchLegacyAuthorizationsIDExecute executes the request
	 * @return Authorization
	 */
	PatchLegacyAuthorizationsIDExecute(r ApiPatchLegacyAuthorizationsIDRequest) (Authorization, error)

	/*
	 * PostLegacyAuthorizations Create a legacy authorization
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiPostLegacyAuthorizationsRequest
	 */
	PostLegacyAuthorizations(ctx _context.Context) ApiPostLegacyAuthorizationsRequest

	/*
	 * PostLegacyAuthorizationsExecute executes the request
	 * @return Authorization
	 */
	PostLegacyAuthorizationsExecute(r ApiPostLegacyAuthorizationsRequest) (Authorization, error)

	/*
	 * PostLegacyAuthorizationsIDPassword Set a legacy authorization password
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param authID The ID of the legacy authorization to update.
	 * @return ApiPostLegacyAuthorizationsIDPasswordRequest
	 */
	PostLegacyAuthorizationsIDPassword(ctx _context.Context, authID string) ApiPostLegacyAuthorizationsIDPasswordRequest

	/*
	 * PostLegacyAuthorizationsIDPasswordExecute executes the request
	 */
	PostLegacyAuthorizationsIDPasswordExecute(r ApiPostLegacyAuthorizationsIDPasswordRequest) error

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on OSS
	// servers.
	OnlyOSS() LegacyAuthorizationsApi

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on cloud
	// servers.
	OnlyCloud() LegacyAuthorizationsApi
}

// LegacyAuthorizationsApiService LegacyAuthorizationsApi service
type LegacyAuthorizationsApiService service

func (a *LegacyAuthorizationsApiService) OnlyOSS() LegacyAuthorizationsApi {
	a.isOnlyOSS = true
	return a
}

func (a *LegacyAuthorizationsApiService) OnlyCloud() LegacyAuthorizationsApi {
	a.isOnlyCloud = true
	return a
}

type ApiDeleteLegacyAuthorizationsIDRequest struct {
	ctx          _context.Context
	ApiService   LegacyAuthorizationsApi
	authID       string
	zapTraceSpan *string
}

func (r ApiDeleteLegacyAuthorizationsIDRequest) AuthID(authID string) ApiDeleteLegacyAuthorizationsIDRequest {
	r.authID = authID
	return r
}
func (r ApiDeleteLegacyAuthorizationsIDRequest) GetAuthID() string {
	return r.authID
}

func (r ApiDeleteLegacyAuthorizationsIDRequest) ZapTraceSpan(zapTraceSpan string) ApiDeleteLegacyAuthorizationsIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiDeleteLegacyAuthorizationsIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiDeleteLegacyAuthorizationsIDRequest) Execute() error {
	return r.ApiService.DeleteLegacyAuthorizationsIDExecute(r)
}

/*
 * DeleteLegacyAuthorizationsID Delete a legacy authorization
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authID The ID of the legacy authorization to delete.
 * @return ApiDeleteLegacyAuthorizationsIDRequest
 */
func (a *LegacyAuthorizationsApiService) DeleteLegacyAuthorizationsID(ctx _context.Context, authID string) ApiDeleteLegacyAuthorizationsIDRequest {
	return ApiDeleteLegacyAuthorizationsIDRequest{
		ApiService: a,
		ctx:        ctx,
		authID:     authID,
	}
}

/*
 * Execute executes the request
 */
func (a *LegacyAuthorizationsApiService) DeleteLegacyAuthorizationsIDExecute(r ApiDeleteLegacyAuthorizationsIDRequest) error {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyAuthorizationsApiService.DeleteLegacyAuthorizationsID")
	if err != nil {
		return GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/legacy/authorizations/{authID}"
	localVarPath = strings.Replace(localVarPath, "{"+"authID"+"}", _neturl.PathEscape(parameterToString(r.authID, "")), -1)

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
		localVarBody, err := _ioutil.ReadAll(body)
		body.Close()
		if err != nil {
			return _fmt.Errorf("%s%w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, localVarHTTPResponse.Status),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s%v", errorPrefix, err.Error())
			return newErr
		}
		newErr.model = &v
		newErr.error = _fmt.Sprintf("%s%v", errorPrefix, v.Error())
		return newErr
	}

	return nil
}

type ApiGetLegacyAuthorizationsRequest struct {
	ctx          _context.Context
	ApiService   LegacyAuthorizationsApi
	zapTraceSpan *string
	userID       *string
	user         *string
	orgID        *string
	org          *string
	token        *string
	authID       *string
}

func (r ApiGetLegacyAuthorizationsRequest) ZapTraceSpan(zapTraceSpan string) ApiGetLegacyAuthorizationsRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetLegacyAuthorizationsRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetLegacyAuthorizationsRequest) UserID(userID string) ApiGetLegacyAuthorizationsRequest {
	r.userID = &userID
	return r
}
func (r ApiGetLegacyAuthorizationsRequest) GetUserID() *string {
	return r.userID
}

func (r ApiGetLegacyAuthorizationsRequest) User(user string) ApiGetLegacyAuthorizationsRequest {
	r.user = &user
	return r
}
func (r ApiGetLegacyAuthorizationsRequest) GetUser() *string {
	return r.user
}

func (r ApiGetLegacyAuthorizationsRequest) OrgID(orgID string) ApiGetLegacyAuthorizationsRequest {
	r.orgID = &orgID
	return r
}
func (r ApiGetLegacyAuthorizationsRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiGetLegacyAuthorizationsRequest) Org(org string) ApiGetLegacyAuthorizationsRequest {
	r.org = &org
	return r
}
func (r ApiGetLegacyAuthorizationsRequest) GetOrg() *string {
	return r.org
}

func (r ApiGetLegacyAuthorizationsRequest) Token(token string) ApiGetLegacyAuthorizationsRequest {
	r.token = &token
	return r
}
func (r ApiGetLegacyAuthorizationsRequest) GetToken() *string {
	return r.token
}

func (r ApiGetLegacyAuthorizationsRequest) AuthID(authID string) ApiGetLegacyAuthorizationsRequest {
	r.authID = &authID
	return r
}
func (r ApiGetLegacyAuthorizationsRequest) GetAuthID() *string {
	return r.authID
}

func (r ApiGetLegacyAuthorizationsRequest) Execute() (Authorizations, error) {
	return r.ApiService.GetLegacyAuthorizationsExecute(r)
}

/*
 * GetLegacyAuthorizations List all legacy authorizations
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetLegacyAuthorizationsRequest
 */
func (a *LegacyAuthorizationsApiService) GetLegacyAuthorizations(ctx _context.Context) ApiGetLegacyAuthorizationsRequest {
	return ApiGetLegacyAuthorizationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Authorizations
 */
func (a *LegacyAuthorizationsApiService) GetLegacyAuthorizationsExecute(r ApiGetLegacyAuthorizationsRequest) (Authorizations, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Authorizations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyAuthorizationsApiService.GetLegacyAuthorizations")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/legacy/authorizations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.userID != nil {
		localVarQueryParams.Add("userID", parameterToString(*r.userID, ""))
	}
	if r.user != nil {
		localVarQueryParams.Add("user", parameterToString(*r.user, ""))
	}
	if r.orgID != nil {
		localVarQueryParams.Add("orgID", parameterToString(*r.orgID, ""))
	}
	if r.org != nil {
		localVarQueryParams.Add("org", parameterToString(*r.org, ""))
	}
	if r.token != nil {
		localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.authID != nil {
		localVarQueryParams.Add("authID", parameterToString(*r.authID, ""))
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
		localVarBody, err := _ioutil.ReadAll(body)
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
			newErr.error = _fmt.Sprintf("%s%v", errorPrefix, err.Error())
			return localVarReturnValue, newErr
		}
		newErr.model = &v
		newErr.error = _fmt.Sprintf("%s%v", errorPrefix, v.Error())
		return localVarReturnValue, newErr
	}

	body, err := GunzipIfNeeded(localVarHTTPResponse)
	if err != nil {
		body.Close()
		return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	localVarBody, err := _ioutil.ReadAll(body)
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

type ApiGetLegacyAuthorizationsIDRequest struct {
	ctx          _context.Context
	ApiService   LegacyAuthorizationsApi
	authID       string
	zapTraceSpan *string
}

func (r ApiGetLegacyAuthorizationsIDRequest) AuthID(authID string) ApiGetLegacyAuthorizationsIDRequest {
	r.authID = authID
	return r
}
func (r ApiGetLegacyAuthorizationsIDRequest) GetAuthID() string {
	return r.authID
}

func (r ApiGetLegacyAuthorizationsIDRequest) ZapTraceSpan(zapTraceSpan string) ApiGetLegacyAuthorizationsIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetLegacyAuthorizationsIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetLegacyAuthorizationsIDRequest) Execute() (Authorization, error) {
	return r.ApiService.GetLegacyAuthorizationsIDExecute(r)
}

/*
 * GetLegacyAuthorizationsID Retrieve a legacy authorization
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authID The ID of the legacy authorization to get.
 * @return ApiGetLegacyAuthorizationsIDRequest
 */
func (a *LegacyAuthorizationsApiService) GetLegacyAuthorizationsID(ctx _context.Context, authID string) ApiGetLegacyAuthorizationsIDRequest {
	return ApiGetLegacyAuthorizationsIDRequest{
		ApiService: a,
		ctx:        ctx,
		authID:     authID,
	}
}

/*
 * Execute executes the request
 * @return Authorization
 */
func (a *LegacyAuthorizationsApiService) GetLegacyAuthorizationsIDExecute(r ApiGetLegacyAuthorizationsIDRequest) (Authorization, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Authorization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyAuthorizationsApiService.GetLegacyAuthorizationsID")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/legacy/authorizations/{authID}"
	localVarPath = strings.Replace(localVarPath, "{"+"authID"+"}", _neturl.PathEscape(parameterToString(r.authID, "")), -1)

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
		localVarBody, err := _ioutil.ReadAll(body)
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
			newErr.error = _fmt.Sprintf("%s%v", errorPrefix, err.Error())
			return localVarReturnValue, newErr
		}
		newErr.model = &v
		newErr.error = _fmt.Sprintf("%s%v", errorPrefix, v.Error())
		return localVarReturnValue, newErr
	}

	body, err := GunzipIfNeeded(localVarHTTPResponse)
	if err != nil {
		body.Close()
		return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	localVarBody, err := _ioutil.ReadAll(body)
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

type ApiPatchLegacyAuthorizationsIDRequest struct {
	ctx                        _context.Context
	ApiService                 LegacyAuthorizationsApi
	authID                     string
	authorizationUpdateRequest *AuthorizationUpdateRequest
	zapTraceSpan               *string
}

func (r ApiPatchLegacyAuthorizationsIDRequest) AuthID(authID string) ApiPatchLegacyAuthorizationsIDRequest {
	r.authID = authID
	return r
}
func (r ApiPatchLegacyAuthorizationsIDRequest) GetAuthID() string {
	return r.authID
}

func (r ApiPatchLegacyAuthorizationsIDRequest) AuthorizationUpdateRequest(authorizationUpdateRequest AuthorizationUpdateRequest) ApiPatchLegacyAuthorizationsIDRequest {
	r.authorizationUpdateRequest = &authorizationUpdateRequest
	return r
}
func (r ApiPatchLegacyAuthorizationsIDRequest) GetAuthorizationUpdateRequest() *AuthorizationUpdateRequest {
	return r.authorizationUpdateRequest
}

func (r ApiPatchLegacyAuthorizationsIDRequest) ZapTraceSpan(zapTraceSpan string) ApiPatchLegacyAuthorizationsIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPatchLegacyAuthorizationsIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPatchLegacyAuthorizationsIDRequest) Execute() (Authorization, error) {
	return r.ApiService.PatchLegacyAuthorizationsIDExecute(r)
}

/*
 * PatchLegacyAuthorizationsID Update a legacy authorization to be active or inactive
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authID The ID of the legacy authorization to update.
 * @return ApiPatchLegacyAuthorizationsIDRequest
 */
func (a *LegacyAuthorizationsApiService) PatchLegacyAuthorizationsID(ctx _context.Context, authID string) ApiPatchLegacyAuthorizationsIDRequest {
	return ApiPatchLegacyAuthorizationsIDRequest{
		ApiService: a,
		ctx:        ctx,
		authID:     authID,
	}
}

/*
 * Execute executes the request
 * @return Authorization
 */
func (a *LegacyAuthorizationsApiService) PatchLegacyAuthorizationsIDExecute(r ApiPatchLegacyAuthorizationsIDRequest) (Authorization, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Authorization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyAuthorizationsApiService.PatchLegacyAuthorizationsID")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/legacy/authorizations/{authID}"
	localVarPath = strings.Replace(localVarPath, "{"+"authID"+"}", _neturl.PathEscape(parameterToString(r.authID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.authorizationUpdateRequest == nil {
		return localVarReturnValue, reportError("authorizationUpdateRequest is required and must be specified")
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
	localVarPostBody = r.authorizationUpdateRequest
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
		localVarBody, err := _ioutil.ReadAll(body)
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
			newErr.error = _fmt.Sprintf("%s%v", errorPrefix, err.Error())
			return localVarReturnValue, newErr
		}
		newErr.model = &v
		newErr.error = _fmt.Sprintf("%s%v", errorPrefix, v.Error())
		return localVarReturnValue, newErr
	}

	body, err := GunzipIfNeeded(localVarHTTPResponse)
	if err != nil {
		body.Close()
		return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	localVarBody, err := _ioutil.ReadAll(body)
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

type ApiPostLegacyAuthorizationsRequest struct {
	ctx                            _context.Context
	ApiService                     LegacyAuthorizationsApi
	legacyAuthorizationPostRequest *LegacyAuthorizationPostRequest
	zapTraceSpan                   *string
}

func (r ApiPostLegacyAuthorizationsRequest) LegacyAuthorizationPostRequest(legacyAuthorizationPostRequest LegacyAuthorizationPostRequest) ApiPostLegacyAuthorizationsRequest {
	r.legacyAuthorizationPostRequest = &legacyAuthorizationPostRequest
	return r
}
func (r ApiPostLegacyAuthorizationsRequest) GetLegacyAuthorizationPostRequest() *LegacyAuthorizationPostRequest {
	return r.legacyAuthorizationPostRequest
}

func (r ApiPostLegacyAuthorizationsRequest) ZapTraceSpan(zapTraceSpan string) ApiPostLegacyAuthorizationsRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPostLegacyAuthorizationsRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPostLegacyAuthorizationsRequest) Execute() (Authorization, error) {
	return r.ApiService.PostLegacyAuthorizationsExecute(r)
}

/*
 * PostLegacyAuthorizations Create a legacy authorization
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostLegacyAuthorizationsRequest
 */
func (a *LegacyAuthorizationsApiService) PostLegacyAuthorizations(ctx _context.Context) ApiPostLegacyAuthorizationsRequest {
	return ApiPostLegacyAuthorizationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Authorization
 */
func (a *LegacyAuthorizationsApiService) PostLegacyAuthorizationsExecute(r ApiPostLegacyAuthorizationsRequest) (Authorization, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Authorization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyAuthorizationsApiService.PostLegacyAuthorizations")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/legacy/authorizations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.legacyAuthorizationPostRequest == nil {
		return localVarReturnValue, reportError("legacyAuthorizationPostRequest is required and must be specified")
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
	localVarPostBody = r.legacyAuthorizationPostRequest
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
		localVarBody, err := _ioutil.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, localVarHTTPResponse.Status),
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = _fmt.Sprintf("%s%v", errorPrefix, err.Error())
				return localVarReturnValue, newErr
			}
			newErr.model = &v
			newErr.error = _fmt.Sprintf("%s%v", errorPrefix, v.Error())
			return localVarReturnValue, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s%v", errorPrefix, err.Error())
			return localVarReturnValue, newErr
		}
		newErr.model = &v
		newErr.error = _fmt.Sprintf("%s%v", errorPrefix, v.Error())
		return localVarReturnValue, newErr
	}

	body, err := GunzipIfNeeded(localVarHTTPResponse)
	if err != nil {
		body.Close()
		return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	localVarBody, err := _ioutil.ReadAll(body)
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

type ApiPostLegacyAuthorizationsIDPasswordRequest struct {
	ctx               _context.Context
	ApiService        LegacyAuthorizationsApi
	authID            string
	passwordResetBody *PasswordResetBody
	zapTraceSpan      *string
}

func (r ApiPostLegacyAuthorizationsIDPasswordRequest) AuthID(authID string) ApiPostLegacyAuthorizationsIDPasswordRequest {
	r.authID = authID
	return r
}
func (r ApiPostLegacyAuthorizationsIDPasswordRequest) GetAuthID() string {
	return r.authID
}

func (r ApiPostLegacyAuthorizationsIDPasswordRequest) PasswordResetBody(passwordResetBody PasswordResetBody) ApiPostLegacyAuthorizationsIDPasswordRequest {
	r.passwordResetBody = &passwordResetBody
	return r
}
func (r ApiPostLegacyAuthorizationsIDPasswordRequest) GetPasswordResetBody() *PasswordResetBody {
	return r.passwordResetBody
}

func (r ApiPostLegacyAuthorizationsIDPasswordRequest) ZapTraceSpan(zapTraceSpan string) ApiPostLegacyAuthorizationsIDPasswordRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPostLegacyAuthorizationsIDPasswordRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPostLegacyAuthorizationsIDPasswordRequest) Execute() error {
	return r.ApiService.PostLegacyAuthorizationsIDPasswordExecute(r)
}

/*
 * PostLegacyAuthorizationsIDPassword Set a legacy authorization password
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authID The ID of the legacy authorization to update.
 * @return ApiPostLegacyAuthorizationsIDPasswordRequest
 */
func (a *LegacyAuthorizationsApiService) PostLegacyAuthorizationsIDPassword(ctx _context.Context, authID string) ApiPostLegacyAuthorizationsIDPasswordRequest {
	return ApiPostLegacyAuthorizationsIDPasswordRequest{
		ApiService: a,
		ctx:        ctx,
		authID:     authID,
	}
}

/*
 * Execute executes the request
 */
func (a *LegacyAuthorizationsApiService) PostLegacyAuthorizationsIDPasswordExecute(r ApiPostLegacyAuthorizationsIDPasswordRequest) error {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyAuthorizationsApiService.PostLegacyAuthorizationsIDPassword")
	if err != nil {
		return GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/legacy/authorizations/{authID}/password"
	localVarPath = strings.Replace(localVarPath, "{"+"authID"+"}", _neturl.PathEscape(parameterToString(r.authID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.passwordResetBody == nil {
		return reportError("passwordResetBody is required and must be specified")
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
	localVarPostBody = r.passwordResetBody
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
		localVarBody, err := _ioutil.ReadAll(body)
		body.Close()
		if err != nil {
			return _fmt.Errorf("%s%w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, localVarHTTPResponse.Status),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s%v", errorPrefix, err.Error())
			return newErr
		}
		newErr.model = &v
		newErr.error = _fmt.Sprintf("%s%v", errorPrefix, v.Error())
		return newErr
	}

	return nil
}

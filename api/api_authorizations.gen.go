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

type AuthorizationsApi interface {

	/*
	 * DeleteAuthorizationsID Delete an authorization
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param authID The ID of the authorization to delete.
	 * @return ApiDeleteAuthorizationsIDRequest
	 */
	DeleteAuthorizationsID(ctx _context.Context, authID string) ApiDeleteAuthorizationsIDRequest

	/*
	 * DeleteAuthorizationsIDExecute executes the request
	 */
	DeleteAuthorizationsIDExecute(r ApiDeleteAuthorizationsIDRequest) error

	/*
	 * GetAuthorizations List all authorizations
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiGetAuthorizationsRequest
	 */
	GetAuthorizations(ctx _context.Context) ApiGetAuthorizationsRequest

	/*
	 * GetAuthorizationsExecute executes the request
	 * @return Authorizations
	 */
	GetAuthorizationsExecute(r ApiGetAuthorizationsRequest) (Authorizations, error)

	/*
	 * GetAuthorizationsID Retrieve an authorization
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param authID The ID of the authorization to get.
	 * @return ApiGetAuthorizationsIDRequest
	 */
	GetAuthorizationsID(ctx _context.Context, authID string) ApiGetAuthorizationsIDRequest

	/*
	 * GetAuthorizationsIDExecute executes the request
	 * @return Authorization
	 */
	GetAuthorizationsIDExecute(r ApiGetAuthorizationsIDRequest) (Authorization, error)

	/*
	 * PatchAuthorizationsID Update an authorization to be active or inactive
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param authID The ID of the authorization to update.
	 * @return ApiPatchAuthorizationsIDRequest
	 */
	PatchAuthorizationsID(ctx _context.Context, authID string) ApiPatchAuthorizationsIDRequest

	/*
	 * PatchAuthorizationsIDExecute executes the request
	 * @return Authorization
	 */
	PatchAuthorizationsIDExecute(r ApiPatchAuthorizationsIDRequest) (Authorization, error)

	/*
	 * PostAuthorizations Create an authorization
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiPostAuthorizationsRequest
	 */
	PostAuthorizations(ctx _context.Context) ApiPostAuthorizationsRequest

	/*
	 * PostAuthorizationsExecute executes the request
	 * @return Authorization
	 */
	PostAuthorizationsExecute(r ApiPostAuthorizationsRequest) (Authorization, error)

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on OSS
	// servers.
	OnlyOSS() AuthorizationsApi

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on cloud
	// servers.
	OnlyCloud() AuthorizationsApi
}

// AuthorizationsApiService AuthorizationsApi service
type AuthorizationsApiService service

func (a *AuthorizationsApiService) OnlyOSS() AuthorizationsApi {
	a.isOnlyOSS = true
	return a
}

func (a *AuthorizationsApiService) OnlyCloud() AuthorizationsApi {
	a.isOnlyCloud = true
	return a
}

type ApiDeleteAuthorizationsIDRequest struct {
	ctx          _context.Context
	ApiService   AuthorizationsApi
	authID       string
	zapTraceSpan *string
}

func (r ApiDeleteAuthorizationsIDRequest) AuthID(authID string) ApiDeleteAuthorizationsIDRequest {
	r.authID = authID
	return r
}
func (r ApiDeleteAuthorizationsIDRequest) GetAuthID() string {
	return r.authID
}

func (r ApiDeleteAuthorizationsIDRequest) ZapTraceSpan(zapTraceSpan string) ApiDeleteAuthorizationsIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiDeleteAuthorizationsIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiDeleteAuthorizationsIDRequest) Execute() error {
	return r.ApiService.DeleteAuthorizationsIDExecute(r)
}

/*
 * DeleteAuthorizationsID Delete an authorization
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authID The ID of the authorization to delete.
 * @return ApiDeleteAuthorizationsIDRequest
 */
func (a *AuthorizationsApiService) DeleteAuthorizationsID(ctx _context.Context, authID string) ApiDeleteAuthorizationsIDRequest {
	return ApiDeleteAuthorizationsIDRequest{
		ApiService: a,
		ctx:        ctx,
		authID:     authID,
	}
}

/*
 * Execute executes the request
 */
func (a *AuthorizationsApiService) DeleteAuthorizationsIDExecute(r ApiDeleteAuthorizationsIDRequest) error {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationsApiService.DeleteAuthorizationsID")
	if err != nil {
		return GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/authorizations/{authID}"
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

type ApiGetAuthorizationsRequest struct {
	ctx          _context.Context
	ApiService   AuthorizationsApi
	zapTraceSpan *string
	userID       *string
	user         *string
	orgID        *string
	org          *string
}

func (r ApiGetAuthorizationsRequest) ZapTraceSpan(zapTraceSpan string) ApiGetAuthorizationsRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetAuthorizationsRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetAuthorizationsRequest) UserID(userID string) ApiGetAuthorizationsRequest {
	r.userID = &userID
	return r
}
func (r ApiGetAuthorizationsRequest) GetUserID() *string {
	return r.userID
}

func (r ApiGetAuthorizationsRequest) User(user string) ApiGetAuthorizationsRequest {
	r.user = &user
	return r
}
func (r ApiGetAuthorizationsRequest) GetUser() *string {
	return r.user
}

func (r ApiGetAuthorizationsRequest) OrgID(orgID string) ApiGetAuthorizationsRequest {
	r.orgID = &orgID
	return r
}
func (r ApiGetAuthorizationsRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiGetAuthorizationsRequest) Org(org string) ApiGetAuthorizationsRequest {
	r.org = &org
	return r
}
func (r ApiGetAuthorizationsRequest) GetOrg() *string {
	return r.org
}

func (r ApiGetAuthorizationsRequest) Execute() (Authorizations, error) {
	return r.ApiService.GetAuthorizationsExecute(r)
}

/*
 * GetAuthorizations List all authorizations
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetAuthorizationsRequest
 */
func (a *AuthorizationsApiService) GetAuthorizations(ctx _context.Context) ApiGetAuthorizationsRequest {
	return ApiGetAuthorizationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Authorizations
 */
func (a *AuthorizationsApiService) GetAuthorizationsExecute(r ApiGetAuthorizationsRequest) (Authorizations, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Authorizations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationsApiService.GetAuthorizations")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/authorizations"

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

type ApiGetAuthorizationsIDRequest struct {
	ctx          _context.Context
	ApiService   AuthorizationsApi
	authID       string
	zapTraceSpan *string
}

func (r ApiGetAuthorizationsIDRequest) AuthID(authID string) ApiGetAuthorizationsIDRequest {
	r.authID = authID
	return r
}
func (r ApiGetAuthorizationsIDRequest) GetAuthID() string {
	return r.authID
}

func (r ApiGetAuthorizationsIDRequest) ZapTraceSpan(zapTraceSpan string) ApiGetAuthorizationsIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetAuthorizationsIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetAuthorizationsIDRequest) Execute() (Authorization, error) {
	return r.ApiService.GetAuthorizationsIDExecute(r)
}

/*
 * GetAuthorizationsID Retrieve an authorization
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authID The ID of the authorization to get.
 * @return ApiGetAuthorizationsIDRequest
 */
func (a *AuthorizationsApiService) GetAuthorizationsID(ctx _context.Context, authID string) ApiGetAuthorizationsIDRequest {
	return ApiGetAuthorizationsIDRequest{
		ApiService: a,
		ctx:        ctx,
		authID:     authID,
	}
}

/*
 * Execute executes the request
 * @return Authorization
 */
func (a *AuthorizationsApiService) GetAuthorizationsIDExecute(r ApiGetAuthorizationsIDRequest) (Authorization, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Authorization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationsApiService.GetAuthorizationsID")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/authorizations/{authID}"
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

type ApiPatchAuthorizationsIDRequest struct {
	ctx                        _context.Context
	ApiService                 AuthorizationsApi
	authID                     string
	authorizationUpdateRequest *AuthorizationUpdateRequest
	zapTraceSpan               *string
}

func (r ApiPatchAuthorizationsIDRequest) AuthID(authID string) ApiPatchAuthorizationsIDRequest {
	r.authID = authID
	return r
}
func (r ApiPatchAuthorizationsIDRequest) GetAuthID() string {
	return r.authID
}

func (r ApiPatchAuthorizationsIDRequest) AuthorizationUpdateRequest(authorizationUpdateRequest AuthorizationUpdateRequest) ApiPatchAuthorizationsIDRequest {
	r.authorizationUpdateRequest = &authorizationUpdateRequest
	return r
}
func (r ApiPatchAuthorizationsIDRequest) GetAuthorizationUpdateRequest() *AuthorizationUpdateRequest {
	return r.authorizationUpdateRequest
}

func (r ApiPatchAuthorizationsIDRequest) ZapTraceSpan(zapTraceSpan string) ApiPatchAuthorizationsIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPatchAuthorizationsIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPatchAuthorizationsIDRequest) Execute() (Authorization, error) {
	return r.ApiService.PatchAuthorizationsIDExecute(r)
}

/*
 * PatchAuthorizationsID Update an authorization to be active or inactive
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authID The ID of the authorization to update.
 * @return ApiPatchAuthorizationsIDRequest
 */
func (a *AuthorizationsApiService) PatchAuthorizationsID(ctx _context.Context, authID string) ApiPatchAuthorizationsIDRequest {
	return ApiPatchAuthorizationsIDRequest{
		ApiService: a,
		ctx:        ctx,
		authID:     authID,
	}
}

/*
 * Execute executes the request
 * @return Authorization
 */
func (a *AuthorizationsApiService) PatchAuthorizationsIDExecute(r ApiPatchAuthorizationsIDRequest) (Authorization, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Authorization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationsApiService.PatchAuthorizationsID")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/authorizations/{authID}"
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

type ApiPostAuthorizationsRequest struct {
	ctx                      _context.Context
	ApiService               AuthorizationsApi
	authorizationPostRequest *AuthorizationPostRequest
	zapTraceSpan             *string
}

func (r ApiPostAuthorizationsRequest) AuthorizationPostRequest(authorizationPostRequest AuthorizationPostRequest) ApiPostAuthorizationsRequest {
	r.authorizationPostRequest = &authorizationPostRequest
	return r
}
func (r ApiPostAuthorizationsRequest) GetAuthorizationPostRequest() *AuthorizationPostRequest {
	return r.authorizationPostRequest
}

func (r ApiPostAuthorizationsRequest) ZapTraceSpan(zapTraceSpan string) ApiPostAuthorizationsRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPostAuthorizationsRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPostAuthorizationsRequest) Execute() (Authorization, error) {
	return r.ApiService.PostAuthorizationsExecute(r)
}

/*
 * PostAuthorizations Create an authorization
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostAuthorizationsRequest
 */
func (a *AuthorizationsApiService) PostAuthorizations(ctx _context.Context) ApiPostAuthorizationsRequest {
	return ApiPostAuthorizationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Authorization
 */
func (a *AuthorizationsApiService) PostAuthorizationsExecute(r ApiPostAuthorizationsRequest) (Authorization, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Authorization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationsApiService.PostAuthorizations")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/authorizations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.authorizationPostRequest == nil {
		return localVarReturnValue, reportError("authorizationPostRequest is required and must be specified")
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
	localVarPostBody = r.authorizationPostRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
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

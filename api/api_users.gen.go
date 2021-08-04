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
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type UsersApi interface {

	/*
	 * DeleteUsersID Delete a user
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param userID The ID of the user to delete.
	 * @return ApiDeleteUsersIDRequest
	 */
	DeleteUsersID(ctx _context.Context, userID string) ApiDeleteUsersIDRequest

	/*
	 * DeleteUsersIDExecute executes the request
	 */
	DeleteUsersIDExecute(r ApiDeleteUsersIDRequest) error

	/*
	 * GetUsers List all users
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiGetUsersRequest
	 */
	GetUsers(ctx _context.Context) ApiGetUsersRequest

	/*
	 * GetUsersExecute executes the request
	 * @return Users
	 */
	GetUsersExecute(r ApiGetUsersRequest) (Users, error)

	/*
	 * GetUsersID Retrieve a user
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param userID The user ID.
	 * @return ApiGetUsersIDRequest
	 */
	GetUsersID(ctx _context.Context, userID string) ApiGetUsersIDRequest

	/*
	 * GetUsersIDExecute executes the request
	 * @return UserResponse
	 */
	GetUsersIDExecute(r ApiGetUsersIDRequest) (UserResponse, error)

	/*
	 * PatchUsersID Update a user
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param userID The ID of the user to update.
	 * @return ApiPatchUsersIDRequest
	 */
	PatchUsersID(ctx _context.Context, userID string) ApiPatchUsersIDRequest

	/*
	 * PatchUsersIDExecute executes the request
	 * @return UserResponse
	 */
	PatchUsersIDExecute(r ApiPatchUsersIDRequest) (UserResponse, error)

	/*
	 * PostUsers Create a user
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiPostUsersRequest
	 */
	PostUsers(ctx _context.Context) ApiPostUsersRequest

	/*
	 * PostUsersExecute executes the request
	 * @return UserResponse
	 */
	PostUsersExecute(r ApiPostUsersRequest) (UserResponse, error)

	/*
	 * PostUsersIDPassword Update a password
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param userID The user ID.
	 * @return ApiPostUsersIDPasswordRequest
	 */
	PostUsersIDPassword(ctx _context.Context, userID string) ApiPostUsersIDPasswordRequest

	/*
	 * PostUsersIDPasswordExecute executes the request
	 */
	PostUsersIDPasswordExecute(r ApiPostUsersIDPasswordRequest) error

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on OSS
	// servers.
	OnlyOSS() UsersApi

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on cloud
	// servers.
	OnlyCloud() UsersApi
}

// UsersApiService UsersApi service
type UsersApiService service

func (a *UsersApiService) OnlyOSS() UsersApi {
	a.isOnlyOSS = true
	return a
}

func (a *UsersApiService) OnlyCloud() UsersApi {
	a.isOnlyCloud = true
	return a
}

type ApiDeleteUsersIDRequest struct {
	ctx          _context.Context
	ApiService   UsersApi
	userID       string
	zapTraceSpan *string
}

func (r ApiDeleteUsersIDRequest) UserID(userID string) ApiDeleteUsersIDRequest {
	r.userID = userID
	return r
}
func (r ApiDeleteUsersIDRequest) GetUserID() string {
	return r.userID
}

func (r ApiDeleteUsersIDRequest) ZapTraceSpan(zapTraceSpan string) ApiDeleteUsersIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiDeleteUsersIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiDeleteUsersIDRequest) Execute() error {
	return r.ApiService.DeleteUsersIDExecute(r)
}

/*
 * DeleteUsersID Delete a user
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userID The ID of the user to delete.
 * @return ApiDeleteUsersIDRequest
 */
func (a *UsersApiService) DeleteUsersID(ctx _context.Context, userID string) ApiDeleteUsersIDRequest {
	return ApiDeleteUsersIDRequest{
		ApiService: a,
		ctx:        ctx,
		userID:     userID,
	}
}

/*
 * Execute executes the request
 */
func (a *UsersApiService) DeleteUsersIDExecute(r ApiDeleteUsersIDRequest) error {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.DeleteUsersID")
	if err != nil {
		return GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{userID}"
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)

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
		localVarBody, err := _io.ReadAll(body)
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
			newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
			return newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
		newErr.model = &v
		return newErr
	}

	return nil
}

type ApiGetUsersRequest struct {
	ctx          _context.Context
	ApiService   UsersApi
	zapTraceSpan *string
	offset       *int32
	limit        *int32
	after        *string
	name         *string
	id           *string
}

func (r ApiGetUsersRequest) ZapTraceSpan(zapTraceSpan string) ApiGetUsersRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetUsersRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetUsersRequest) Offset(offset int32) ApiGetUsersRequest {
	r.offset = &offset
	return r
}
func (r ApiGetUsersRequest) GetOffset() *int32 {
	return r.offset
}

func (r ApiGetUsersRequest) Limit(limit int32) ApiGetUsersRequest {
	r.limit = &limit
	return r
}
func (r ApiGetUsersRequest) GetLimit() *int32 {
	return r.limit
}

func (r ApiGetUsersRequest) After(after string) ApiGetUsersRequest {
	r.after = &after
	return r
}
func (r ApiGetUsersRequest) GetAfter() *string {
	return r.after
}

func (r ApiGetUsersRequest) Name(name string) ApiGetUsersRequest {
	r.name = &name
	return r
}
func (r ApiGetUsersRequest) GetName() *string {
	return r.name
}

func (r ApiGetUsersRequest) Id(id string) ApiGetUsersRequest {
	r.id = &id
	return r
}
func (r ApiGetUsersRequest) GetId() *string {
	return r.id
}

func (r ApiGetUsersRequest) Execute() (Users, error) {
	return r.ApiService.GetUsersExecute(r)
}

/*
 * GetUsers List all users
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetUsersRequest
 */
func (a *UsersApiService) GetUsers(ctx _context.Context) ApiGetUsersRequest {
	return ApiGetUsersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Users
 */
func (a *UsersApiService) GetUsersExecute(r ApiGetUsersRequest) (Users, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Users
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.GetUsers")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users"

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
		localVarBody, err := _io.ReadAll(body)
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
		return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	localVarBody, err := _io.ReadAll(body)
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

type ApiGetUsersIDRequest struct {
	ctx          _context.Context
	ApiService   UsersApi
	userID       string
	zapTraceSpan *string
}

func (r ApiGetUsersIDRequest) UserID(userID string) ApiGetUsersIDRequest {
	r.userID = userID
	return r
}
func (r ApiGetUsersIDRequest) GetUserID() string {
	return r.userID
}

func (r ApiGetUsersIDRequest) ZapTraceSpan(zapTraceSpan string) ApiGetUsersIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetUsersIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetUsersIDRequest) Execute() (UserResponse, error) {
	return r.ApiService.GetUsersIDExecute(r)
}

/*
 * GetUsersID Retrieve a user
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userID The user ID.
 * @return ApiGetUsersIDRequest
 */
func (a *UsersApiService) GetUsersID(ctx _context.Context, userID string) ApiGetUsersIDRequest {
	return ApiGetUsersIDRequest{
		ApiService: a,
		ctx:        ctx,
		userID:     userID,
	}
}

/*
 * Execute executes the request
 * @return UserResponse
 */
func (a *UsersApiService) GetUsersIDExecute(r ApiGetUsersIDRequest) (UserResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.GetUsersID")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{userID}"
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)

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
		localVarBody, err := _io.ReadAll(body)
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
		return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	localVarBody, err := _io.ReadAll(body)
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

type ApiPatchUsersIDRequest struct {
	ctx          _context.Context
	ApiService   UsersApi
	userID       string
	user         *User
	zapTraceSpan *string
}

func (r ApiPatchUsersIDRequest) UserID(userID string) ApiPatchUsersIDRequest {
	r.userID = userID
	return r
}
func (r ApiPatchUsersIDRequest) GetUserID() string {
	return r.userID
}

func (r ApiPatchUsersIDRequest) User(user User) ApiPatchUsersIDRequest {
	r.user = &user
	return r
}
func (r ApiPatchUsersIDRequest) GetUser() *User {
	return r.user
}

func (r ApiPatchUsersIDRequest) ZapTraceSpan(zapTraceSpan string) ApiPatchUsersIDRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPatchUsersIDRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPatchUsersIDRequest) Execute() (UserResponse, error) {
	return r.ApiService.PatchUsersIDExecute(r)
}

/*
 * PatchUsersID Update a user
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userID The ID of the user to update.
 * @return ApiPatchUsersIDRequest
 */
func (a *UsersApiService) PatchUsersID(ctx _context.Context, userID string) ApiPatchUsersIDRequest {
	return ApiPatchUsersIDRequest{
		ApiService: a,
		ctx:        ctx,
		userID:     userID,
	}
}

/*
 * Execute executes the request
 * @return UserResponse
 */
func (a *UsersApiService) PatchUsersIDExecute(r ApiPatchUsersIDRequest) (UserResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.PatchUsersID")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{userID}"
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.user == nil {
		return localVarReturnValue, reportError("user is required and must be specified")
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
	localVarPostBody = r.user
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
		localVarBody, err := _io.ReadAll(body)
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
		return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	localVarBody, err := _io.ReadAll(body)
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

type ApiPostUsersRequest struct {
	ctx          _context.Context
	ApiService   UsersApi
	user         *User
	zapTraceSpan *string
}

func (r ApiPostUsersRequest) User(user User) ApiPostUsersRequest {
	r.user = &user
	return r
}
func (r ApiPostUsersRequest) GetUser() *User {
	return r.user
}

func (r ApiPostUsersRequest) ZapTraceSpan(zapTraceSpan string) ApiPostUsersRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPostUsersRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPostUsersRequest) Execute() (UserResponse, error) {
	return r.ApiService.PostUsersExecute(r)
}

/*
 * PostUsers Create a user
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostUsersRequest
 */
func (a *UsersApiService) PostUsers(ctx _context.Context) ApiPostUsersRequest {
	return ApiPostUsersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return UserResponse
 */
func (a *UsersApiService) PostUsersExecute(r ApiPostUsersRequest) (UserResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.PostUsers")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.user == nil {
		return localVarReturnValue, reportError("user is required and must be specified")
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
	localVarPostBody = r.user
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
		localVarBody, err := _io.ReadAll(body)
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
		return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	localVarBody, err := _io.ReadAll(body)
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

type ApiPostUsersIDPasswordRequest struct {
	ctx               _context.Context
	ApiService        UsersApi
	userID            string
	passwordResetBody *PasswordResetBody
	zapTraceSpan      *string
}

func (r ApiPostUsersIDPasswordRequest) UserID(userID string) ApiPostUsersIDPasswordRequest {
	r.userID = userID
	return r
}
func (r ApiPostUsersIDPasswordRequest) GetUserID() string {
	return r.userID
}

func (r ApiPostUsersIDPasswordRequest) PasswordResetBody(passwordResetBody PasswordResetBody) ApiPostUsersIDPasswordRequest {
	r.passwordResetBody = &passwordResetBody
	return r
}
func (r ApiPostUsersIDPasswordRequest) GetPasswordResetBody() *PasswordResetBody {
	return r.passwordResetBody
}

func (r ApiPostUsersIDPasswordRequest) ZapTraceSpan(zapTraceSpan string) ApiPostUsersIDPasswordRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPostUsersIDPasswordRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPostUsersIDPasswordRequest) Execute() error {
	return r.ApiService.PostUsersIDPasswordExecute(r)
}

/*
 * PostUsersIDPassword Update a password
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userID The user ID.
 * @return ApiPostUsersIDPasswordRequest
 */
func (a *UsersApiService) PostUsersIDPassword(ctx _context.Context, userID string) ApiPostUsersIDPasswordRequest {
	return ApiPostUsersIDPasswordRequest{
		ApiService: a,
		ctx:        ctx,
		userID:     userID,
	}
}

/*
 * Execute executes the request
 */
func (a *UsersApiService) PostUsersIDPasswordExecute(r ApiPostUsersIDPasswordRequest) error {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.PostUsersIDPassword")
	if err != nil {
		return GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{userID}/password"
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", _neturl.PathEscape(parameterToString(r.userID, "")), -1)

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
		localVarBody, err := _io.ReadAll(body)
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
			newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
			return newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
		newErr.model = &v
		return newErr
	}

	return nil
}

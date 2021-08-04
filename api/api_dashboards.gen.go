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
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

type DashboardsApi interface {

	/*
	 * GetDashboards List all dashboards
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiGetDashboardsRequest
	 */
	GetDashboards(ctx _context.Context) ApiGetDashboardsRequest

	/*
	 * GetDashboardsExecute executes the request
	 * @return Dashboards
	 */
	GetDashboardsExecute(r ApiGetDashboardsRequest) (Dashboards, error)

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on OSS
	// servers.
	OnlyOSS() DashboardsApi

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on cloud
	// servers.
	OnlyCloud() DashboardsApi
}

// DashboardsApiService DashboardsApi service
type DashboardsApiService service

func (a *DashboardsApiService) OnlyOSS() DashboardsApi {
	a.isOnlyOSS = true
	return a
}

func (a *DashboardsApiService) OnlyCloud() DashboardsApi {
	a.isOnlyCloud = true
	return a
}

type ApiGetDashboardsRequest struct {
	ctx          _context.Context
	ApiService   DashboardsApi
	zapTraceSpan *string
	offset       *int32
	limit        *int32
	descending   *bool
	owner        *string
	sortBy       *string
	id           *[]string
	orgID        *string
	org          *string
}

func (r ApiGetDashboardsRequest) ZapTraceSpan(zapTraceSpan string) ApiGetDashboardsRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetDashboardsRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetDashboardsRequest) Offset(offset int32) ApiGetDashboardsRequest {
	r.offset = &offset
	return r
}
func (r ApiGetDashboardsRequest) GetOffset() *int32 {
	return r.offset
}

func (r ApiGetDashboardsRequest) Limit(limit int32) ApiGetDashboardsRequest {
	r.limit = &limit
	return r
}
func (r ApiGetDashboardsRequest) GetLimit() *int32 {
	return r.limit
}

func (r ApiGetDashboardsRequest) Descending(descending bool) ApiGetDashboardsRequest {
	r.descending = &descending
	return r
}
func (r ApiGetDashboardsRequest) GetDescending() *bool {
	return r.descending
}

func (r ApiGetDashboardsRequest) Owner(owner string) ApiGetDashboardsRequest {
	r.owner = &owner
	return r
}
func (r ApiGetDashboardsRequest) GetOwner() *string {
	return r.owner
}

func (r ApiGetDashboardsRequest) SortBy(sortBy string) ApiGetDashboardsRequest {
	r.sortBy = &sortBy
	return r
}
func (r ApiGetDashboardsRequest) GetSortBy() *string {
	return r.sortBy
}

func (r ApiGetDashboardsRequest) Id(id []string) ApiGetDashboardsRequest {
	r.id = &id
	return r
}
func (r ApiGetDashboardsRequest) GetId() *[]string {
	return r.id
}

func (r ApiGetDashboardsRequest) OrgID(orgID string) ApiGetDashboardsRequest {
	r.orgID = &orgID
	return r
}
func (r ApiGetDashboardsRequest) GetOrgID() *string {
	return r.orgID
}

func (r ApiGetDashboardsRequest) Org(org string) ApiGetDashboardsRequest {
	r.org = &org
	return r
}
func (r ApiGetDashboardsRequest) GetOrg() *string {
	return r.org
}

func (r ApiGetDashboardsRequest) Execute() (Dashboards, error) {
	return r.ApiService.GetDashboardsExecute(r)
}

/*
 * GetDashboards List all dashboards
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetDashboardsRequest
 */
func (a *DashboardsApiService) GetDashboards(ctx _context.Context) ApiGetDashboardsRequest {
	return ApiGetDashboardsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Dashboards
 */
func (a *DashboardsApiService) GetDashboardsExecute(r ApiGetDashboardsRequest) (Dashboards, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Dashboards
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardsApiService.GetDashboards")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dashboards"

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
	if r.owner != nil {
		localVarQueryParams.Add("owner", parameterToString(*r.owner, ""))
	}
	if r.sortBy != nil {
		localVarQueryParams.Add("sortBy", parameterToString(*r.sortBy, ""))
	}
	if r.id != nil {
		t := *r.id
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("id", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("id", parameterToString(t, "multi"))
		}
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

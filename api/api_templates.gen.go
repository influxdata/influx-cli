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
)

// Linger please
var (
	_ _context.Context
)

type TemplatesApi interface {

	/*
	 * ApplyTemplate Apply or dry-run an InfluxDB Template
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiApplyTemplateRequest
	 */
	ApplyTemplate(ctx _context.Context) ApiApplyTemplateRequest

	/*
	 * ApplyTemplateExecute executes the request
	 * @return TemplateSummary
	 */
	ApplyTemplateExecute(r ApiApplyTemplateRequest) (TemplateSummary, error)

	/*
	 * ExportTemplate Export a new Influx Template
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiExportTemplateRequest
	 */
	ExportTemplate(ctx _context.Context) ApiExportTemplateRequest

	/*
	 * ExportTemplateExecute executes the request
	 * @return []TemplateEntry
	 */
	ExportTemplateExecute(r ApiExportTemplateRequest) ([]TemplateEntry, error)

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on OSS
	// servers.
	OnlyOSS() TemplatesApi

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on cloud
	// servers.
	OnlyCloud() TemplatesApi
}

// TemplatesApiService TemplatesApi service
type TemplatesApiService service

func (a *TemplatesApiService) OnlyOSS() TemplatesApi {
	a.isOnlyOSS = true
	return a
}

func (a *TemplatesApiService) OnlyCloud() TemplatesApi {
	a.isOnlyCloud = true
	return a
}

type ApiApplyTemplateRequest struct {
	ctx           _context.Context
	ApiService    TemplatesApi
	templateApply *TemplateApply
}

func (r ApiApplyTemplateRequest) TemplateApply(templateApply TemplateApply) ApiApplyTemplateRequest {
	r.templateApply = &templateApply
	return r
}
func (r ApiApplyTemplateRequest) GetTemplateApply() *TemplateApply {
	return r.templateApply
}

func (r ApiApplyTemplateRequest) Execute() (TemplateSummary, error) {
	return r.ApiService.ApplyTemplateExecute(r)
}

/*
 * ApplyTemplate Apply or dry-run an InfluxDB Template
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiApplyTemplateRequest
 */
func (a *TemplatesApiService) ApplyTemplate(ctx _context.Context) ApiApplyTemplateRequest {
	return ApiApplyTemplateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return TemplateSummary
 */
func (a *TemplatesApiService) ApplyTemplateExecute(r ApiApplyTemplateRequest) (TemplateSummary, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TemplateSummary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplatesApiService.ApplyTemplate")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/templates/apply"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.templateApply == nil {
		return localVarReturnValue, reportError("templateApply is required and must be specified")
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
	// body params
	localVarPostBody = r.templateApply
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
			error: _fmt.Sprintf("%s: code %s", errorPrefix, localVarHTTPResponse.Status),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s: %s", errorPrefix, err.Error())
			return localVarReturnValue, newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", errorPrefix, v.GetMessage()))
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

type ApiExportTemplateRequest struct {
	ctx            _context.Context
	ApiService     TemplatesApi
	templateExport *TemplateExport
}

func (r ApiExportTemplateRequest) TemplateExport(templateExport TemplateExport) ApiExportTemplateRequest {
	r.templateExport = &templateExport
	return r
}
func (r ApiExportTemplateRequest) GetTemplateExport() *TemplateExport {
	return r.templateExport
}

func (r ApiExportTemplateRequest) Execute() ([]TemplateEntry, error) {
	return r.ApiService.ExportTemplateExecute(r)
}

/*
 * ExportTemplate Export a new Influx Template
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiExportTemplateRequest
 */
func (a *TemplatesApiService) ExportTemplate(ctx _context.Context) ApiExportTemplateRequest {
	return ApiExportTemplateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return []TemplateEntry
 */
func (a *TemplatesApiService) ExportTemplateExecute(r ApiExportTemplateRequest) ([]TemplateEntry, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TemplateEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplatesApiService.ExportTemplate")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/templates/export"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	// body params
	localVarPostBody = r.templateExport
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
			error: _fmt.Sprintf("%s: code %s", errorPrefix, localVarHTTPResponse.Status),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s: %s", errorPrefix, err.Error())
			return localVarReturnValue, newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", errorPrefix, v.GetMessage()))
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

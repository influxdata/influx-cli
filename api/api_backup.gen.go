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
	"time"
)

// Linger please
var (
	_ _context.Context
)

type BackupApi interface {

	/*
	 * GetBackupMetadata Download snapshot of all metadata in the server
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiGetBackupMetadataRequest
	 */
	GetBackupMetadata(ctx _context.Context) ApiGetBackupMetadataRequest

	/*
	 * GetBackupMetadataExecute executes the request
	 * @return *os.File
	 */
	GetBackupMetadataExecute(r ApiGetBackupMetadataRequest) (*_nethttp.Response, error)

	/*
	 * GetBackupShardId Download snapshot of all TSM data in a shard
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param shardID The shard ID.
	 * @return ApiGetBackupShardIdRequest
	 */
	GetBackupShardId(ctx _context.Context, shardID int64) ApiGetBackupShardIdRequest

	/*
	 * GetBackupShardIdExecute executes the request
	 * @return *os.File
	 */
	GetBackupShardIdExecute(r ApiGetBackupShardIdRequest) (*_nethttp.Response, error)

	// Sets the intention of the API to only work for InfluxDB OSS servers - for logging error messages
	OnlyOSS() BackupApi

	// Sets the intention of the API to only work for InfluxDB Cloud servers - for logging error messages
	OnlyCloud() BackupApi
}

// BackupApiService BackupApi service
type BackupApiService service

func (a *BackupApiService) OnlyOSS() BackupApi {
	a.isOnlyOSS = true
	return a
}

func (a *BackupApiService) OnlyCloud() BackupApi {
	a.isOnlyCloud = true
	return a
}

type ApiGetBackupMetadataRequest struct {
	ctx            _context.Context
	ApiService     BackupApi
	zapTraceSpan   *string
	acceptEncoding *string
	isOnlyOSS      bool
	isOnlyCloud    bool
}

func (r ApiGetBackupMetadataRequest) ZapTraceSpan(zapTraceSpan string) ApiGetBackupMetadataRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetBackupMetadataRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetBackupMetadataRequest) AcceptEncoding(acceptEncoding string) ApiGetBackupMetadataRequest {
	r.acceptEncoding = &acceptEncoding
	return r
}
func (r ApiGetBackupMetadataRequest) GetAcceptEncoding() *string {
	return r.acceptEncoding
}

func (r ApiGetBackupMetadataRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetBackupMetadataExecute(r)
}

func (r ApiGetBackupMetadataRequest) OnlyOSS() ApiGetBackupMetadataRequest {
	r.isOnlyOSS = true
	return r
}

func (r ApiGetBackupMetadataRequest) OnlyCloud() ApiGetBackupMetadataRequest {
	r.isOnlyCloud = true
	return r
}

/*
 * GetBackupMetadata Download snapshot of all metadata in the server
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetBackupMetadataRequest
 */
func (a *BackupApiService) GetBackupMetadata(ctx _context.Context) ApiGetBackupMetadataRequest {
	return ApiGetBackupMetadataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return *os.File
 */
func (a *BackupApiService) GetBackupMetadataExecute(r ApiGetBackupMetadataRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *_nethttp.Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupApiService.GetBackupMetadata")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/backup/metadata"

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
	localVarHTTPHeaderAccepts := []string{"multipart/mixed", "application/json"}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
	}

	var errorPrefix string
	if r.isOnlyOSS || a.isOnlyOSS {
		errorPrefix = "InfluxDB OSS-only command failed: "
	} else if r.isOnlyCloud || a.isOnlyCloud {
		errorPrefix = "InfluxDB Cloud-only command failed: "
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return localVarReturnValue, _fmt.Errorf("%s%v", errorPrefix, err)
		}
		localVarBody, err := _ioutil.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarReturnValue, _fmt.Errorf("%s%v", errorPrefix, err)
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
		newErr.error = _fmt.Sprintf("%s%v", errorPrefix, v.Error())
		return localVarReturnValue, newErr
	}

	localVarReturnValue = localVarHTTPResponse

	return localVarReturnValue, nil
}

type ApiGetBackupShardIdRequest struct {
	ctx            _context.Context
	ApiService     BackupApi
	shardID        int64
	zapTraceSpan   *string
	acceptEncoding *string
	since          *time.Time
	isOnlyOSS      bool
	isOnlyCloud    bool
}

func (r ApiGetBackupShardIdRequest) ShardID(shardID int64) ApiGetBackupShardIdRequest {
	r.shardID = shardID
	return r
}
func (r ApiGetBackupShardIdRequest) GetShardID() int64 {
	return r.shardID
}

func (r ApiGetBackupShardIdRequest) ZapTraceSpan(zapTraceSpan string) ApiGetBackupShardIdRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetBackupShardIdRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetBackupShardIdRequest) AcceptEncoding(acceptEncoding string) ApiGetBackupShardIdRequest {
	r.acceptEncoding = &acceptEncoding
	return r
}
func (r ApiGetBackupShardIdRequest) GetAcceptEncoding() *string {
	return r.acceptEncoding
}

func (r ApiGetBackupShardIdRequest) Since(since time.Time) ApiGetBackupShardIdRequest {
	r.since = &since
	return r
}
func (r ApiGetBackupShardIdRequest) GetSince() *time.Time {
	return r.since
}

func (r ApiGetBackupShardIdRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetBackupShardIdExecute(r)
}

func (r ApiGetBackupShardIdRequest) OnlyOSS() ApiGetBackupShardIdRequest {
	r.isOnlyOSS = true
	return r
}

func (r ApiGetBackupShardIdRequest) OnlyCloud() ApiGetBackupShardIdRequest {
	r.isOnlyCloud = true
	return r
}

/*
 * GetBackupShardId Download snapshot of all TSM data in a shard
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param shardID The shard ID.
 * @return ApiGetBackupShardIdRequest
 */
func (a *BackupApiService) GetBackupShardId(ctx _context.Context, shardID int64) ApiGetBackupShardIdRequest {
	return ApiGetBackupShardIdRequest{
		ApiService: a,
		ctx:        ctx,
		shardID:    shardID,
	}
}

/*
 * Execute executes the request
 * @return *os.File
 */
func (a *BackupApiService) GetBackupShardIdExecute(r ApiGetBackupShardIdRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *_nethttp.Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupApiService.GetBackupShardId")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/backup/shards/{shardID}"
	localVarPath = strings.Replace(localVarPath, "{"+"shardID"+"}", _neturl.PathEscape(parameterToString(r.shardID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.since != nil {
		localVarQueryParams.Add("since", parameterToString(*r.since, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/octet-stream", "application/json"}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
	}

	var errorPrefix string
	if r.isOnlyOSS || a.isOnlyOSS {
		errorPrefix = "InfluxDB OSS-only command failed: "
	} else if r.isOnlyCloud || a.isOnlyCloud {
		errorPrefix = "InfluxDB Cloud-only command failed: "
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return localVarReturnValue, _fmt.Errorf("%s%v", errorPrefix, err)
		}
		localVarBody, err := _ioutil.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarReturnValue, _fmt.Errorf("%s%v", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, localVarHTTPResponse.Status),
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = _fmt.Sprintf("%s%v", errorPrefix, err.Error())
				return localVarReturnValue, newErr
			}
			newErr.error = _fmt.Sprintf("%s%v", errorPrefix, v.Error())
			return localVarReturnValue, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s%v", errorPrefix, err.Error())
			return localVarReturnValue, newErr
		}
		newErr.error = _fmt.Sprintf("%s%v", errorPrefix, v.Error())
		return localVarReturnValue, newErr
	}

	localVarReturnValue = localVarHTTPResponse

	return localVarReturnValue, nil
}

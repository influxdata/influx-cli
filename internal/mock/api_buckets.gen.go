// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influx-cli/v2/api (interfaces: BucketsApi)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/influxdata/influx-cli/v2/api"
)

// MockBucketsApi is a mock of BucketsApi interface.
type MockBucketsApi struct {
	ctrl     *gomock.Controller
	recorder *MockBucketsApiMockRecorder
}

// MockBucketsApiMockRecorder is the mock recorder for MockBucketsApi.
type MockBucketsApiMockRecorder struct {
	mock *MockBucketsApi
}

// NewMockBucketsApi creates a new mock instance.
func NewMockBucketsApi(ctrl *gomock.Controller) *MockBucketsApi {
	mock := &MockBucketsApi{ctrl: ctrl}
	mock.recorder = &MockBucketsApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBucketsApi) EXPECT() *MockBucketsApiMockRecorder {
	return m.recorder
}

// DeleteBucketsID mocks base method.
func (m *MockBucketsApi) DeleteBucketsID(arg0 context.Context, arg1 string) api.ApiDeleteBucketsIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucketsID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiDeleteBucketsIDRequest)
	return ret0
}

// DeleteBucketsID indicates an expected call of DeleteBucketsID.
func (mr *MockBucketsApiMockRecorder) DeleteBucketsID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucketsID", reflect.TypeOf((*MockBucketsApi)(nil).DeleteBucketsID), arg0, arg1)
}

// DeleteBucketsIDExecute mocks base method.
func (m *MockBucketsApi) DeleteBucketsIDExecute(arg0 api.ApiDeleteBucketsIDRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucketsIDExecute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBucketsIDExecute indicates an expected call of DeleteBucketsIDExecute.
func (mr *MockBucketsApiMockRecorder) DeleteBucketsIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucketsIDExecute", reflect.TypeOf((*MockBucketsApi)(nil).DeleteBucketsIDExecute), arg0)
}

// DeleteBucketsIDExecuteWithHttpInfo mocks base method.
func (m *MockBucketsApi) DeleteBucketsIDExecuteWithHttpInfo(arg0 api.ApiDeleteBucketsIDRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucketsIDExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBucketsIDExecuteWithHttpInfo indicates an expected call of DeleteBucketsIDExecuteWithHttpInfo.
func (mr *MockBucketsApiMockRecorder) DeleteBucketsIDExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucketsIDExecuteWithHttpInfo", reflect.TypeOf((*MockBucketsApi)(nil).DeleteBucketsIDExecuteWithHttpInfo), arg0)
}

// GetBuckets mocks base method.
func (m *MockBucketsApi) GetBuckets(arg0 context.Context) api.ApiGetBucketsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBuckets", arg0)
	ret0, _ := ret[0].(api.ApiGetBucketsRequest)
	return ret0
}

// GetBuckets indicates an expected call of GetBuckets.
func (mr *MockBucketsApiMockRecorder) GetBuckets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuckets", reflect.TypeOf((*MockBucketsApi)(nil).GetBuckets), arg0)
}

// GetBucketsExecute mocks base method.
func (m *MockBucketsApi) GetBucketsExecute(arg0 api.ApiGetBucketsRequest) (api.Buckets, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBucketsExecute", arg0)
	ret0, _ := ret[0].(api.Buckets)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketsExecute indicates an expected call of GetBucketsExecute.
func (mr *MockBucketsApiMockRecorder) GetBucketsExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketsExecute", reflect.TypeOf((*MockBucketsApi)(nil).GetBucketsExecute), arg0)
}

// GetBucketsExecuteWithHttpInfo mocks base method.
func (m *MockBucketsApi) GetBucketsExecuteWithHttpInfo(arg0 api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBucketsExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.Buckets)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBucketsExecuteWithHttpInfo indicates an expected call of GetBucketsExecuteWithHttpInfo.
func (mr *MockBucketsApiMockRecorder) GetBucketsExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketsExecuteWithHttpInfo", reflect.TypeOf((*MockBucketsApi)(nil).GetBucketsExecuteWithHttpInfo), arg0)
}

// GetBucketsID mocks base method.
func (m *MockBucketsApi) GetBucketsID(arg0 context.Context, arg1 string) api.ApiGetBucketsIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBucketsID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiGetBucketsIDRequest)
	return ret0
}

// GetBucketsID indicates an expected call of GetBucketsID.
func (mr *MockBucketsApiMockRecorder) GetBucketsID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketsID", reflect.TypeOf((*MockBucketsApi)(nil).GetBucketsID), arg0, arg1)
}

// GetBucketsIDExecute mocks base method.
func (m *MockBucketsApi) GetBucketsIDExecute(arg0 api.ApiGetBucketsIDRequest) (api.Bucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBucketsIDExecute", arg0)
	ret0, _ := ret[0].(api.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketsIDExecute indicates an expected call of GetBucketsIDExecute.
func (mr *MockBucketsApiMockRecorder) GetBucketsIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketsIDExecute", reflect.TypeOf((*MockBucketsApi)(nil).GetBucketsIDExecute), arg0)
}

// GetBucketsIDExecuteWithHttpInfo mocks base method.
func (m *MockBucketsApi) GetBucketsIDExecuteWithHttpInfo(arg0 api.ApiGetBucketsIDRequest) (api.Bucket, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBucketsIDExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.Bucket)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBucketsIDExecuteWithHttpInfo indicates an expected call of GetBucketsIDExecuteWithHttpInfo.
func (mr *MockBucketsApiMockRecorder) GetBucketsIDExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketsIDExecuteWithHttpInfo", reflect.TypeOf((*MockBucketsApi)(nil).GetBucketsIDExecuteWithHttpInfo), arg0)
}

// OnlyCloud mocks base method.
func (m *MockBucketsApi) OnlyCloud() api.BucketsApi {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnlyCloud")
	ret0, _ := ret[0].(api.BucketsApi)
	return ret0
}

// OnlyCloud indicates an expected call of OnlyCloud.
func (mr *MockBucketsApiMockRecorder) OnlyCloud() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnlyCloud", reflect.TypeOf((*MockBucketsApi)(nil).OnlyCloud))
}

// OnlyOSS mocks base method.
func (m *MockBucketsApi) OnlyOSS() api.BucketsApi {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnlyOSS")
	ret0, _ := ret[0].(api.BucketsApi)
	return ret0
}

// OnlyOSS indicates an expected call of OnlyOSS.
func (mr *MockBucketsApiMockRecorder) OnlyOSS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnlyOSS", reflect.TypeOf((*MockBucketsApi)(nil).OnlyOSS))
}

// PatchBucketsID mocks base method.
func (m *MockBucketsApi) PatchBucketsID(arg0 context.Context, arg1 string) api.ApiPatchBucketsIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchBucketsID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiPatchBucketsIDRequest)
	return ret0
}

// PatchBucketsID indicates an expected call of PatchBucketsID.
func (mr *MockBucketsApiMockRecorder) PatchBucketsID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchBucketsID", reflect.TypeOf((*MockBucketsApi)(nil).PatchBucketsID), arg0, arg1)
}

// PatchBucketsIDExecute mocks base method.
func (m *MockBucketsApi) PatchBucketsIDExecute(arg0 api.ApiPatchBucketsIDRequest) (api.Bucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchBucketsIDExecute", arg0)
	ret0, _ := ret[0].(api.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchBucketsIDExecute indicates an expected call of PatchBucketsIDExecute.
func (mr *MockBucketsApiMockRecorder) PatchBucketsIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchBucketsIDExecute", reflect.TypeOf((*MockBucketsApi)(nil).PatchBucketsIDExecute), arg0)
}

// PatchBucketsIDExecuteWithHttpInfo mocks base method.
func (m *MockBucketsApi) PatchBucketsIDExecuteWithHttpInfo(arg0 api.ApiPatchBucketsIDRequest) (api.Bucket, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchBucketsIDExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.Bucket)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PatchBucketsIDExecuteWithHttpInfo indicates an expected call of PatchBucketsIDExecuteWithHttpInfo.
func (mr *MockBucketsApiMockRecorder) PatchBucketsIDExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchBucketsIDExecuteWithHttpInfo", reflect.TypeOf((*MockBucketsApi)(nil).PatchBucketsIDExecuteWithHttpInfo), arg0)
}

// PostBuckets mocks base method.
func (m *MockBucketsApi) PostBuckets(arg0 context.Context) api.ApiPostBucketsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostBuckets", arg0)
	ret0, _ := ret[0].(api.ApiPostBucketsRequest)
	return ret0
}

// PostBuckets indicates an expected call of PostBuckets.
func (mr *MockBucketsApiMockRecorder) PostBuckets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBuckets", reflect.TypeOf((*MockBucketsApi)(nil).PostBuckets), arg0)
}

// PostBucketsExecute mocks base method.
func (m *MockBucketsApi) PostBucketsExecute(arg0 api.ApiPostBucketsRequest) (api.Bucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostBucketsExecute", arg0)
	ret0, _ := ret[0].(api.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostBucketsExecute indicates an expected call of PostBucketsExecute.
func (mr *MockBucketsApiMockRecorder) PostBucketsExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBucketsExecute", reflect.TypeOf((*MockBucketsApi)(nil).PostBucketsExecute), arg0)
}

// PostBucketsExecuteWithHttpInfo mocks base method.
func (m *MockBucketsApi) PostBucketsExecuteWithHttpInfo(arg0 api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostBucketsExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.Bucket)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PostBucketsExecuteWithHttpInfo indicates an expected call of PostBucketsExecuteWithHttpInfo.
func (mr *MockBucketsApiMockRecorder) PostBucketsExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBucketsExecuteWithHttpInfo", reflect.TypeOf((*MockBucketsApi)(nil).PostBucketsExecuteWithHttpInfo), arg0)
}

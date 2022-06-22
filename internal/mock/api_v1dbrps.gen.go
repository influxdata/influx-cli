// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influx-cli/v2/api (interfaces: DBRPsApi)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/influxdata/influx-cli/v2/api"
)

// MockDBRPsApi is a mock of DBRPsApi interface.
type MockDBRPsApi struct {
	ctrl     *gomock.Controller
	recorder *MockDBRPsApiMockRecorder
}

// MockDBRPsApiMockRecorder is the mock recorder for MockDBRPsApi.
type MockDBRPsApiMockRecorder struct {
	mock *MockDBRPsApi
}

// NewMockDBRPsApi creates a new mock instance.
func NewMockDBRPsApi(ctrl *gomock.Controller) *MockDBRPsApi {
	mock := &MockDBRPsApi{ctrl: ctrl}
	mock.recorder = &MockDBRPsApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBRPsApi) EXPECT() *MockDBRPsApiMockRecorder {
	return m.recorder
}

// DeleteDBRPID mocks base method.
func (m *MockDBRPsApi) DeleteDBRPID(arg0 context.Context, arg1 string) api.ApiDeleteDBRPIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDBRPID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiDeleteDBRPIDRequest)
	return ret0
}

// DeleteDBRPID indicates an expected call of DeleteDBRPID.
func (mr *MockDBRPsApiMockRecorder) DeleteDBRPID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDBRPID", reflect.TypeOf((*MockDBRPsApi)(nil).DeleteDBRPID), arg0, arg1)
}

// DeleteDBRPIDExecute mocks base method.
func (m *MockDBRPsApi) DeleteDBRPIDExecute(arg0 api.ApiDeleteDBRPIDRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDBRPIDExecute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDBRPIDExecute indicates an expected call of DeleteDBRPIDExecute.
func (mr *MockDBRPsApiMockRecorder) DeleteDBRPIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDBRPIDExecute", reflect.TypeOf((*MockDBRPsApi)(nil).DeleteDBRPIDExecute), arg0)
}

// DeleteDBRPIDExecuteWithHttpInfo mocks base method.
func (m *MockDBRPsApi) DeleteDBRPIDExecuteWithHttpInfo(arg0 api.ApiDeleteDBRPIDRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDBRPIDExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDBRPIDExecuteWithHttpInfo indicates an expected call of DeleteDBRPIDExecuteWithHttpInfo.
func (mr *MockDBRPsApiMockRecorder) DeleteDBRPIDExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDBRPIDExecuteWithHttpInfo", reflect.TypeOf((*MockDBRPsApi)(nil).DeleteDBRPIDExecuteWithHttpInfo), arg0)
}

// GetDBRPs mocks base method.
func (m *MockDBRPsApi) GetDBRPs(arg0 context.Context) api.ApiGetDBRPsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDBRPs", arg0)
	ret0, _ := ret[0].(api.ApiGetDBRPsRequest)
	return ret0
}

// GetDBRPs indicates an expected call of GetDBRPs.
func (mr *MockDBRPsApiMockRecorder) GetDBRPs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDBRPs", reflect.TypeOf((*MockDBRPsApi)(nil).GetDBRPs), arg0)
}

// GetDBRPsExecute mocks base method.
func (m *MockDBRPsApi) GetDBRPsExecute(arg0 api.ApiGetDBRPsRequest) (api.DBRPs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDBRPsExecute", arg0)
	ret0, _ := ret[0].(api.DBRPs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDBRPsExecute indicates an expected call of GetDBRPsExecute.
func (mr *MockDBRPsApiMockRecorder) GetDBRPsExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDBRPsExecute", reflect.TypeOf((*MockDBRPsApi)(nil).GetDBRPsExecute), arg0)
}

// GetDBRPsExecuteWithHttpInfo mocks base method.
func (m *MockDBRPsApi) GetDBRPsExecuteWithHttpInfo(arg0 api.ApiGetDBRPsRequest) (api.DBRPs, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDBRPsExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.DBRPs)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDBRPsExecuteWithHttpInfo indicates an expected call of GetDBRPsExecuteWithHttpInfo.
func (mr *MockDBRPsApiMockRecorder) GetDBRPsExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDBRPsExecuteWithHttpInfo", reflect.TypeOf((*MockDBRPsApi)(nil).GetDBRPsExecuteWithHttpInfo), arg0)
}

// GetDBRPsID mocks base method.
func (m *MockDBRPsApi) GetDBRPsID(arg0 context.Context, arg1 string) api.ApiGetDBRPsIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDBRPsID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiGetDBRPsIDRequest)
	return ret0
}

// GetDBRPsID indicates an expected call of GetDBRPsID.
func (mr *MockDBRPsApiMockRecorder) GetDBRPsID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDBRPsID", reflect.TypeOf((*MockDBRPsApi)(nil).GetDBRPsID), arg0, arg1)
}

// GetDBRPsIDExecute mocks base method.
func (m *MockDBRPsApi) GetDBRPsIDExecute(arg0 api.ApiGetDBRPsIDRequest) (api.DBRPGet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDBRPsIDExecute", arg0)
	ret0, _ := ret[0].(api.DBRPGet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDBRPsIDExecute indicates an expected call of GetDBRPsIDExecute.
func (mr *MockDBRPsApiMockRecorder) GetDBRPsIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDBRPsIDExecute", reflect.TypeOf((*MockDBRPsApi)(nil).GetDBRPsIDExecute), arg0)
}

// GetDBRPsIDExecuteWithHttpInfo mocks base method.
func (m *MockDBRPsApi) GetDBRPsIDExecuteWithHttpInfo(arg0 api.ApiGetDBRPsIDRequest) (api.DBRPGet, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDBRPsIDExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.DBRPGet)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDBRPsIDExecuteWithHttpInfo indicates an expected call of GetDBRPsIDExecuteWithHttpInfo.
func (mr *MockDBRPsApiMockRecorder) GetDBRPsIDExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDBRPsIDExecuteWithHttpInfo", reflect.TypeOf((*MockDBRPsApi)(nil).GetDBRPsIDExecuteWithHttpInfo), arg0)
}

// PatchDBRPID mocks base method.
func (m *MockDBRPsApi) PatchDBRPID(arg0 context.Context, arg1 string) api.ApiPatchDBRPIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchDBRPID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiPatchDBRPIDRequest)
	return ret0
}

// PatchDBRPID indicates an expected call of PatchDBRPID.
func (mr *MockDBRPsApiMockRecorder) PatchDBRPID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchDBRPID", reflect.TypeOf((*MockDBRPsApi)(nil).PatchDBRPID), arg0, arg1)
}

// PatchDBRPIDExecute mocks base method.
func (m *MockDBRPsApi) PatchDBRPIDExecute(arg0 api.ApiPatchDBRPIDRequest) (api.DBRPGet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchDBRPIDExecute", arg0)
	ret0, _ := ret[0].(api.DBRPGet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchDBRPIDExecute indicates an expected call of PatchDBRPIDExecute.
func (mr *MockDBRPsApiMockRecorder) PatchDBRPIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchDBRPIDExecute", reflect.TypeOf((*MockDBRPsApi)(nil).PatchDBRPIDExecute), arg0)
}

// PatchDBRPIDExecuteWithHttpInfo mocks base method.
func (m *MockDBRPsApi) PatchDBRPIDExecuteWithHttpInfo(arg0 api.ApiPatchDBRPIDRequest) (api.DBRPGet, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchDBRPIDExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.DBRPGet)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PatchDBRPIDExecuteWithHttpInfo indicates an expected call of PatchDBRPIDExecuteWithHttpInfo.
func (mr *MockDBRPsApiMockRecorder) PatchDBRPIDExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchDBRPIDExecuteWithHttpInfo", reflect.TypeOf((*MockDBRPsApi)(nil).PatchDBRPIDExecuteWithHttpInfo), arg0)
}

// PostDBRP mocks base method.
func (m *MockDBRPsApi) PostDBRP(arg0 context.Context) api.ApiPostDBRPRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostDBRP", arg0)
	ret0, _ := ret[0].(api.ApiPostDBRPRequest)
	return ret0
}

// PostDBRP indicates an expected call of PostDBRP.
func (mr *MockDBRPsApiMockRecorder) PostDBRP(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostDBRP", reflect.TypeOf((*MockDBRPsApi)(nil).PostDBRP), arg0)
}

// PostDBRPExecute mocks base method.
func (m *MockDBRPsApi) PostDBRPExecute(arg0 api.ApiPostDBRPRequest) (api.DBRP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostDBRPExecute", arg0)
	ret0, _ := ret[0].(api.DBRP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostDBRPExecute indicates an expected call of PostDBRPExecute.
func (mr *MockDBRPsApiMockRecorder) PostDBRPExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostDBRPExecute", reflect.TypeOf((*MockDBRPsApi)(nil).PostDBRPExecute), arg0)
}

// PostDBRPExecuteWithHttpInfo mocks base method.
func (m *MockDBRPsApi) PostDBRPExecuteWithHttpInfo(arg0 api.ApiPostDBRPRequest) (api.DBRP, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostDBRPExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.DBRP)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PostDBRPExecuteWithHttpInfo indicates an expected call of PostDBRPExecuteWithHttpInfo.
func (mr *MockDBRPsApiMockRecorder) PostDBRPExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostDBRPExecuteWithHttpInfo", reflect.TypeOf((*MockDBRPsApi)(nil).PostDBRPExecuteWithHttpInfo), arg0)
}

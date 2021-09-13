// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influx-cli/v2/api (interfaces: RemoteConnectionsApi)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/influxdata/influx-cli/v2/api"
)

// MockRemoteConnectionsApi is a mock of RemoteConnectionsApi interface.
type MockRemoteConnectionsApi struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteConnectionsApiMockRecorder
}

// MockRemoteConnectionsApiMockRecorder is the mock recorder for MockRemoteConnectionsApi.
type MockRemoteConnectionsApiMockRecorder struct {
	mock *MockRemoteConnectionsApi
}

// NewMockRemoteConnectionsApi creates a new mock instance.
func NewMockRemoteConnectionsApi(ctrl *gomock.Controller) *MockRemoteConnectionsApi {
	mock := &MockRemoteConnectionsApi{ctrl: ctrl}
	mock.recorder = &MockRemoteConnectionsApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteConnectionsApi) EXPECT() *MockRemoteConnectionsApiMockRecorder {
	return m.recorder
}

// DeleteRemoteConnectionByID mocks base method.
func (m *MockRemoteConnectionsApi) DeleteRemoteConnectionByID(arg0 context.Context, arg1 string) api.ApiDeleteRemoteConnectionByIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRemoteConnectionByID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiDeleteRemoteConnectionByIDRequest)
	return ret0
}

// DeleteRemoteConnectionByID indicates an expected call of DeleteRemoteConnectionByID.
func (mr *MockRemoteConnectionsApiMockRecorder) DeleteRemoteConnectionByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRemoteConnectionByID", reflect.TypeOf((*MockRemoteConnectionsApi)(nil).DeleteRemoteConnectionByID), arg0, arg1)
}

// DeleteRemoteConnectionByIDExecute mocks base method.
func (m *MockRemoteConnectionsApi) DeleteRemoteConnectionByIDExecute(arg0 api.ApiDeleteRemoteConnectionByIDRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRemoteConnectionByIDExecute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRemoteConnectionByIDExecute indicates an expected call of DeleteRemoteConnectionByIDExecute.
func (mr *MockRemoteConnectionsApiMockRecorder) DeleteRemoteConnectionByIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRemoteConnectionByIDExecute", reflect.TypeOf((*MockRemoteConnectionsApi)(nil).DeleteRemoteConnectionByIDExecute), arg0)
}

// GetRemoteConnectionByID mocks base method.
func (m *MockRemoteConnectionsApi) GetRemoteConnectionByID(arg0 context.Context, arg1 string) api.ApiGetRemoteConnectionByIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteConnectionByID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiGetRemoteConnectionByIDRequest)
	return ret0
}

// GetRemoteConnectionByID indicates an expected call of GetRemoteConnectionByID.
func (mr *MockRemoteConnectionsApiMockRecorder) GetRemoteConnectionByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteConnectionByID", reflect.TypeOf((*MockRemoteConnectionsApi)(nil).GetRemoteConnectionByID), arg0, arg1)
}

// GetRemoteConnectionByIDExecute mocks base method.
func (m *MockRemoteConnectionsApi) GetRemoteConnectionByIDExecute(arg0 api.ApiGetRemoteConnectionByIDRequest) (api.RemoteConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteConnectionByIDExecute", arg0)
	ret0, _ := ret[0].(api.RemoteConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRemoteConnectionByIDExecute indicates an expected call of GetRemoteConnectionByIDExecute.
func (mr *MockRemoteConnectionsApiMockRecorder) GetRemoteConnectionByIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteConnectionByIDExecute", reflect.TypeOf((*MockRemoteConnectionsApi)(nil).GetRemoteConnectionByIDExecute), arg0)
}

// GetRemoteConnections mocks base method.
func (m *MockRemoteConnectionsApi) GetRemoteConnections(arg0 context.Context) api.ApiGetRemoteConnectionsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteConnections", arg0)
	ret0, _ := ret[0].(api.ApiGetRemoteConnectionsRequest)
	return ret0
}

// GetRemoteConnections indicates an expected call of GetRemoteConnections.
func (mr *MockRemoteConnectionsApiMockRecorder) GetRemoteConnections(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteConnections", reflect.TypeOf((*MockRemoteConnectionsApi)(nil).GetRemoteConnections), arg0)
}

// GetRemoteConnectionsExecute mocks base method.
func (m *MockRemoteConnectionsApi) GetRemoteConnectionsExecute(arg0 api.ApiGetRemoteConnectionsRequest) (api.RemoteConnections, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteConnectionsExecute", arg0)
	ret0, _ := ret[0].(api.RemoteConnections)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRemoteConnectionsExecute indicates an expected call of GetRemoteConnectionsExecute.
func (mr *MockRemoteConnectionsApiMockRecorder) GetRemoteConnectionsExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteConnectionsExecute", reflect.TypeOf((*MockRemoteConnectionsApi)(nil).GetRemoteConnectionsExecute), arg0)
}

// OnlyCloud mocks base method.
func (m *MockRemoteConnectionsApi) OnlyCloud() api.RemoteConnectionsApi {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnlyCloud")
	ret0, _ := ret[0].(api.RemoteConnectionsApi)
	return ret0
}

// OnlyCloud indicates an expected call of OnlyCloud.
func (mr *MockRemoteConnectionsApiMockRecorder) OnlyCloud() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnlyCloud", reflect.TypeOf((*MockRemoteConnectionsApi)(nil).OnlyCloud))
}

// OnlyOSS mocks base method.
func (m *MockRemoteConnectionsApi) OnlyOSS() api.RemoteConnectionsApi {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnlyOSS")
	ret0, _ := ret[0].(api.RemoteConnectionsApi)
	return ret0
}

// OnlyOSS indicates an expected call of OnlyOSS.
func (mr *MockRemoteConnectionsApiMockRecorder) OnlyOSS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnlyOSS", reflect.TypeOf((*MockRemoteConnectionsApi)(nil).OnlyOSS))
}

// PatchRemoteConnectionByID mocks base method.
func (m *MockRemoteConnectionsApi) PatchRemoteConnectionByID(arg0 context.Context, arg1 string) api.ApiPatchRemoteConnectionByIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchRemoteConnectionByID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiPatchRemoteConnectionByIDRequest)
	return ret0
}

// PatchRemoteConnectionByID indicates an expected call of PatchRemoteConnectionByID.
func (mr *MockRemoteConnectionsApiMockRecorder) PatchRemoteConnectionByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchRemoteConnectionByID", reflect.TypeOf((*MockRemoteConnectionsApi)(nil).PatchRemoteConnectionByID), arg0, arg1)
}

// PatchRemoteConnectionByIDExecute mocks base method.
func (m *MockRemoteConnectionsApi) PatchRemoteConnectionByIDExecute(arg0 api.ApiPatchRemoteConnectionByIDRequest) (api.RemoteConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchRemoteConnectionByIDExecute", arg0)
	ret0, _ := ret[0].(api.RemoteConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchRemoteConnectionByIDExecute indicates an expected call of PatchRemoteConnectionByIDExecute.
func (mr *MockRemoteConnectionsApiMockRecorder) PatchRemoteConnectionByIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchRemoteConnectionByIDExecute", reflect.TypeOf((*MockRemoteConnectionsApi)(nil).PatchRemoteConnectionByIDExecute), arg0)
}

// PostRemoteConnection mocks base method.
func (m *MockRemoteConnectionsApi) PostRemoteConnection(arg0 context.Context) api.ApiPostRemoteConnectionRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostRemoteConnection", arg0)
	ret0, _ := ret[0].(api.ApiPostRemoteConnectionRequest)
	return ret0
}

// PostRemoteConnection indicates an expected call of PostRemoteConnection.
func (mr *MockRemoteConnectionsApiMockRecorder) PostRemoteConnection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostRemoteConnection", reflect.TypeOf((*MockRemoteConnectionsApi)(nil).PostRemoteConnection), arg0)
}

// PostRemoteConnectionExecute mocks base method.
func (m *MockRemoteConnectionsApi) PostRemoteConnectionExecute(arg0 api.ApiPostRemoteConnectionRequest) (api.RemoteConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostRemoteConnectionExecute", arg0)
	ret0, _ := ret[0].(api.RemoteConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostRemoteConnectionExecute indicates an expected call of PostRemoteConnectionExecute.
func (mr *MockRemoteConnectionsApiMockRecorder) PostRemoteConnectionExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostRemoteConnectionExecute", reflect.TypeOf((*MockRemoteConnectionsApi)(nil).PostRemoteConnectionExecute), arg0)
}

// PostValidateRemoteConnectionByID mocks base method.
func (m *MockRemoteConnectionsApi) PostValidateRemoteConnectionByID(arg0 context.Context, arg1 string) api.ApiPostValidateRemoteConnectionByIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostValidateRemoteConnectionByID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiPostValidateRemoteConnectionByIDRequest)
	return ret0
}

// PostValidateRemoteConnectionByID indicates an expected call of PostValidateRemoteConnectionByID.
func (mr *MockRemoteConnectionsApiMockRecorder) PostValidateRemoteConnectionByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostValidateRemoteConnectionByID", reflect.TypeOf((*MockRemoteConnectionsApi)(nil).PostValidateRemoteConnectionByID), arg0, arg1)
}

// PostValidateRemoteConnectionByIDExecute mocks base method.
func (m *MockRemoteConnectionsApi) PostValidateRemoteConnectionByIDExecute(arg0 api.ApiPostValidateRemoteConnectionByIDRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostValidateRemoteConnectionByIDExecute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostValidateRemoteConnectionByIDExecute indicates an expected call of PostValidateRemoteConnectionByIDExecute.
func (mr *MockRemoteConnectionsApiMockRecorder) PostValidateRemoteConnectionByIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostValidateRemoteConnectionByIDExecute", reflect.TypeOf((*MockRemoteConnectionsApi)(nil).PostValidateRemoteConnectionByIDExecute), arg0)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influx-cli/v2/internal/api (interfaces: SetupApi)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/influxdata/influx-cli/v2/internal/api"
)

// MockSetupApi is a mock of SetupApi interface.
type MockSetupApi struct {
	ctrl     *gomock.Controller
	recorder *MockSetupApiMockRecorder
}

// MockSetupApiMockRecorder is the mock recorder for MockSetupApi.
type MockSetupApiMockRecorder struct {
	mock *MockSetupApi
}

// NewMockSetupApi creates a new mock instance.
func NewMockSetupApi(ctrl *gomock.Controller) *MockSetupApi {
	mock := &MockSetupApi{ctrl: ctrl}
	mock.recorder = &MockSetupApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSetupApi) EXPECT() *MockSetupApiMockRecorder {
	return m.recorder
}

// GetSetup mocks base method.
func (m *MockSetupApi) GetSetup(arg0 context.Context) api.ApiGetSetupRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSetup", arg0)
	ret0, _ := ret[0].(api.ApiGetSetupRequest)
	return ret0
}

// GetSetup indicates an expected call of GetSetup.
func (mr *MockSetupApiMockRecorder) GetSetup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSetup", reflect.TypeOf((*MockSetupApi)(nil).GetSetup), arg0)
}

// GetSetupExecute mocks base method.
func (m *MockSetupApi) GetSetupExecute(arg0 api.ApiGetSetupRequest) (api.InlineResponse200, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSetupExecute", arg0)
	ret0, _ := ret[0].(api.InlineResponse200)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSetupExecute indicates an expected call of GetSetupExecute.
func (mr *MockSetupApiMockRecorder) GetSetupExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSetupExecute", reflect.TypeOf((*MockSetupApi)(nil).GetSetupExecute), arg0)
}

// PostSetup mocks base method.
func (m *MockSetupApi) PostSetup(arg0 context.Context) api.ApiPostSetupRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostSetup", arg0)
	ret0, _ := ret[0].(api.ApiPostSetupRequest)
	return ret0
}

// PostSetup indicates an expected call of PostSetup.
func (mr *MockSetupApiMockRecorder) PostSetup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostSetup", reflect.TypeOf((*MockSetupApi)(nil).PostSetup), arg0)
}

// PostSetupExecute mocks base method.
func (m *MockSetupApi) PostSetupExecute(arg0 api.ApiPostSetupRequest) (api.OnboardingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostSetupExecute", arg0)
	ret0, _ := ret[0].(api.OnboardingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostSetupExecute indicates an expected call of PostSetupExecute.
func (mr *MockSetupApiMockRecorder) PostSetupExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostSetupExecute", reflect.TypeOf((*MockSetupApi)(nil).PostSetupExecute), arg0)
}

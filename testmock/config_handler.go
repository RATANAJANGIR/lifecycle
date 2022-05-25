// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/buildpacks/lifecycle (interfaces: ConfigHandler)

// Package testmock is a generated GoMock package.
package testmock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	buildpack "github.com/buildpacks/lifecycle/buildpack"
)

// MockConfigHandler is a mock of ConfigHandler interface.
type MockConfigHandler struct {
	ctrl     *gomock.Controller
	recorder *MockConfigHandlerMockRecorder
}

// MockConfigHandlerMockRecorder is the mock recorder for MockConfigHandler.
type MockConfigHandlerMockRecorder struct {
	mock *MockConfigHandler
}

// NewMockConfigHandler creates a new mock instance.
func NewMockConfigHandler(ctrl *gomock.Controller) *MockConfigHandler {
	mock := &MockConfigHandler{ctrl: ctrl}
	mock.recorder = &MockConfigHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigHandler) EXPECT() *MockConfigHandlerMockRecorder {
	return m.recorder
}

// ReadGroup mocks base method.
func (m *MockConfigHandler) ReadGroup(arg0 string) ([]buildpack.GroupBuildpack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadGroup", arg0)
	ret0, _ := ret[0].([]buildpack.GroupBuildpack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadGroup indicates an expected call of ReadGroup.
func (mr *MockConfigHandlerMockRecorder) ReadGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadGroup", reflect.TypeOf((*MockConfigHandler)(nil).ReadGroup), arg0)
}
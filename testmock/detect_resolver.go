// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/buildpacks/lifecycle (interfaces: DetectResolver)

// Package testmock is a generated GoMock package.
package testmock

import (
	reflect "reflect"
	sync "sync"

	gomock "github.com/golang/mock/gomock"

	buildpack "github.com/buildpacks/lifecycle/buildpack"
	"github.com/buildpacks/lifecycle/platform/files"
)

// MockDetectResolver is a mock of DetectResolver interface.
type MockDetectResolver struct {
	ctrl     *gomock.Controller
	recorder *MockDetectResolverMockRecorder
}

// MockDetectResolverMockRecorder is the mock recorder for MockDetectResolver.
type MockDetectResolverMockRecorder struct {
	mock *MockDetectResolver
}

// NewMockDetectResolver creates a new mock instance.
func NewMockDetectResolver(ctrl *gomock.Controller) *MockDetectResolver {
	mock := &MockDetectResolver{ctrl: ctrl}
	mock.recorder = &MockDetectResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDetectResolver) EXPECT() *MockDetectResolverMockRecorder {
	return m.recorder
}

// Resolve mocks base method.
func (m *MockDetectResolver) Resolve(arg0 []buildpack.GroupElement, arg1 *sync.Map) ([]buildpack.GroupElement, []files.BuildPlanEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolve", arg0, arg1)
	ret0, _ := ret[0].([]buildpack.GroupElement)
	ret1, _ := ret[1].([]files.BuildPlanEntry)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Resolve indicates an expected call of Resolve.
func (mr *MockDetectResolverMockRecorder) Resolve(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockDetectResolver)(nil).Resolve), arg0, arg1)
}

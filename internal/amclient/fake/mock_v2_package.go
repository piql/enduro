// Code generated by MockGen. DO NOT EDIT.
// Source:github.com/penwern/enduro/internal/amclient (interfaces: PackageService)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	amclient "github.com/penwern/enduro/internal/amclient"
	gomock "github.com/golang/mock/gomock"
)

// MockPackageService is a mock of PackageService interface.
type MockPackageService struct {
	ctrl     *gomock.Controller
	recorder *MockPackageServiceMockRecorder
}

// MockPackageServiceMockRecorder is the mock recorder for MockPackageService.
type MockPackageServiceMockRecorder struct {
	mock *MockPackageService
}

// NewMockPackageService creates a new mock instance.
func NewMockPackageService(ctrl *gomock.Controller) *MockPackageService {
	mock := &MockPackageService{ctrl: ctrl}
	mock.recorder = &MockPackageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPackageService) EXPECT() *MockPackageServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPackageService) Create(arg0 context.Context, arg1 *amclient.PackageCreateRequest) (*amclient.PackageCreateResponse, *amclient.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*amclient.PackageCreateResponse)
	ret1, _ := ret[1].(*amclient.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockPackageServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPackageService)(nil).Create), arg0, arg1)
}

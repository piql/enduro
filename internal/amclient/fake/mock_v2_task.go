// Code generated by MockGen. DO NOT EDIT.
// Source:github.com/penwern/enduro/internal/amclient (interfaces: TaskService)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	amclient "github.com/penwern/enduro/internal/amclient"
	gomock "github.com/golang/mock/gomock"
)

// MockTaskService is a mock of TaskService interface.
type MockTaskService struct {
	ctrl     *gomock.Controller
	recorder *MockTaskServiceMockRecorder
}

// MockTaskServiceMockRecorder is the mock recorder for MockTaskService.
type MockTaskServiceMockRecorder struct {
	mock *MockTaskService
}

// NewMockTaskService creates a new mock instance.
func NewMockTaskService(ctrl *gomock.Controller) *MockTaskService {
	mock := &MockTaskService{ctrl: ctrl}
	mock.recorder = &MockTaskServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskService) EXPECT() *MockTaskServiceMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockTaskService) Read(arg0 context.Context, arg1 string) (*amclient.TaskDetailed, *amclient.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1)
	ret0, _ := ret[0].(*amclient.TaskDetailed)
	ret1, _ := ret[1].(*amclient.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Read indicates an expected call of Read.
func (mr *MockTaskServiceMockRecorder) Read(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockTaskService)(nil).Read), arg0, arg1)
}

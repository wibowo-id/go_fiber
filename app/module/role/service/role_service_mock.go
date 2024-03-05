// Code generated by MockGen. DO NOT EDIT.
// Source: dlh_oss_be/app/module/role/service (interfaces: RoleService)

// Package service is a generated GoMock package.
package service

import (
	schema "dlh_oss_be/app/database/schema"
	request "dlh_oss_be/app/module/role/request"
	request0 "dlh_oss_be/app/request"
	paginator "dlh_oss_be/utils/paginator"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRoleService is a mock of RoleService interface.
type MockRoleService struct {
	ctrl     *gomock.Controller
	recorder *MockRoleServiceMockRecorder
}

// MockRoleServiceMockRecorder is the mock recorder for MockRoleService.
type MockRoleServiceMockRecorder struct {
	mock *MockRoleService
}

// NewMockRoleService creates a new mock instance.
func NewMockRoleService(ctrl *gomock.Controller) *MockRoleService {
	mock := &MockRoleService{ctrl: ctrl}
	mock.recorder = &MockRoleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleService) EXPECT() *MockRoleServiceMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockRoleService) All(arg0 request0.Pagination) ([]*schema.Role, paginator.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", arg0)
	ret0, _ := ret[0].([]*schema.Role)
	ret1, _ := ret[1].(paginator.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// All indicates an expected call of All.
func (mr *MockRoleServiceMockRecorder) All(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockRoleService)(nil).All), arg0)
}

// Delete mocks base method.
func (m *MockRoleService) Delete(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRoleServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleService)(nil).Delete), arg0, arg1)
}

// Show mocks base method.
func (m *MockRoleService) Show(arg0 string) (*schema.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", arg0)
	ret0, _ := ret[0].(*schema.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show.
func (mr *MockRoleServiceMockRecorder) Show(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockRoleService)(nil).Show), arg0)
}

// Store mocks base method.
func (m *MockRoleService) Store(arg0 request.RoleRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockRoleServiceMockRecorder) Store(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockRoleService)(nil).Store), arg0)
}

// Update mocks base method.
func (m *MockRoleService) Update(arg0 string, arg1 request.RoleRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRoleServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRoleService)(nil).Update), arg0, arg1)
}
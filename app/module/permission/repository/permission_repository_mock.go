// Code generated by MockGen. DO NOT EDIT.
// Source: dlh_oss_be/app/module/permission/repository (interfaces: PermissionRepository)

// Package repository is a generated GoMock package.
package repository

import (
	schema "dlh_oss_be/app/database/schema"
	request "dlh_oss_be/app/module/permission/request"
	request0 "dlh_oss_be/app/request"
	paginator "dlh_oss_be/utils/paginator"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPermissionRepository is a mock of PermissionRepository interface.
type MockPermissionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPermissionRepositoryMockRecorder
}

// MockPermissionRepositoryMockRecorder is the mock recorder for MockPermissionRepository.
type MockPermissionRepositoryMockRecorder struct {
	mock *MockPermissionRepository
}

// NewMockPermissionRepository creates a new mock instance.
func NewMockPermissionRepository(ctrl *gomock.Controller) *MockPermissionRepository {
	mock := &MockPermissionRepository{ctrl: ctrl}
	mock.recorder = &MockPermissionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPermissionRepository) EXPECT() *MockPermissionRepositoryMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockPermissionRepository) All(arg0 request0.Pagination) ([]*schema.Permission, paginator.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", arg0)
	ret0, _ := ret[0].([]*schema.Permission)
	ret1, _ := ret[1].(paginator.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// All indicates an expected call of All.
func (mr *MockPermissionRepositoryMockRecorder) All(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockPermissionRepository)(nil).All), arg0)
}

// CheckExist mocks base method.
func (m *MockPermissionRepository) CheckExist(arg0 string) (*schema.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExist", arg0)
	ret0, _ := ret[0].(*schema.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExist indicates an expected call of CheckExist.
func (mr *MockPermissionRepositoryMockRecorder) CheckExist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExist", reflect.TypeOf((*MockPermissionRepository)(nil).CheckExist), arg0)
}

// Create mocks base method.
func (m *MockPermissionRepository) Create(arg0 *schema.Permission) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPermissionRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPermissionRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockPermissionRepository) Delete(arg0 request.PermissionDeleteRequest, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPermissionRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPermissionRepository)(nil).Delete), arg0, arg1)
}

// FindByUserId mocks base method.
func (m *MockPermissionRepository) FindByUserId(arg0 string) ([]*schema.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserId", arg0)
	ret0, _ := ret[0].([]*schema.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserId indicates an expected call of FindByUserId.
func (mr *MockPermissionRepositoryMockRecorder) FindByUserId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserId", reflect.TypeOf((*MockPermissionRepository)(nil).FindByUserId), arg0)
}

// FindOne mocks base method.
func (m *MockPermissionRepository) FindOne(arg0 string) (*schema.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0)
	ret0, _ := ret[0].(*schema.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockPermissionRepositoryMockRecorder) FindOne(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockPermissionRepository)(nil).FindOne), arg0)
}

// Update mocks base method.
func (m *MockPermissionRepository) Update(arg0 string, arg1 *schema.Permission) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPermissionRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPermissionRepository)(nil).Update), arg0, arg1)
}
// Code generated by MockGen. DO NOT EDIT.
// Source: go_fiber_wibowo/app/module/version/repository (interfaces: VersionRepository)

// Package repository is a generated GoMock package.
package repository

import (
	schema "go_fiber_wibowo/app/database/schema"
	request "go_fiber_wibowo/app/module/version/request"
	paginator "go_fiber_wibowo/utils/paginator"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockVersionRepository is a mock of VersionRepository interface.
type MockVersionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockVersionRepositoryMockRecorder
}

// MockVersionRepositoryMockRecorder is the mock recorder for MockVersionRepository.
type MockVersionRepositoryMockRecorder struct {
	mock *MockVersionRepository
}

// NewMockVersionRepository creates a new mock instance.
func NewMockVersionRepository(ctrl *gomock.Controller) *MockVersionRepository {
	mock := &MockVersionRepository{ctrl: ctrl}
	mock.recorder = &MockVersionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVersionRepository) EXPECT() *MockVersionRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockVersionRepository) Create(arg0 *schema.Version) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockVersionRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVersionRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockVersionRepository) Delete(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockVersionRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVersionRepository)(nil).Delete), arg0)
}

// FindOne mocks base method.
func (m *MockVersionRepository) FindOne(arg0 uuid.UUID) (*schema.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0)
	ret0, _ := ret[0].(*schema.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockVersionRepositoryMockRecorder) FindOne(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockVersionRepository)(nil).FindOne), arg0)
}

// GetOneVersion mocks base method.
func (m *MockVersionRepository) GetOneVersion() (*schema.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneVersion")
	ret0, _ := ret[0].(*schema.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneVersion indicates an expected call of GetOneVersion.
func (mr *MockVersionRepositoryMockRecorder) GetOneVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneVersion", reflect.TypeOf((*MockVersionRepository)(nil).GetOneVersion))
}

// GetVersions mocks base method.
func (m *MockVersionRepository) GetVersions(arg0 request.VersionsRequest) ([]*schema.Version, paginator.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersions", arg0)
	ret0, _ := ret[0].([]*schema.Version)
	ret1, _ := ret[1].(paginator.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetVersions indicates an expected call of GetVersions.
func (mr *MockVersionRepositoryMockRecorder) GetVersions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersions", reflect.TypeOf((*MockVersionRepository)(nil).GetVersions), arg0)
}

// Update mocks base method.
func (m *MockVersionRepository) Update(arg0 uuid.UUID, arg1 *schema.Version) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockVersionRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVersionRepository)(nil).Update), arg0, arg1)
}

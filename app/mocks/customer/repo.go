// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/layers/repositories/customer/init.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/gobbeston/trainning-golang-microservice/app/entities"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepo is a mock of Repo interface
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockRepo) CreateUser(input *entities.Users) (*entities.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", input)
	ret0, _ := ret[0].(*entities.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockRepoMockRecorder) CreateUser(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepo)(nil).CreateUser), input)
}

// CreateRoles mocks base method
func (m *MockRepo) CreateRoles(input *entities.Roles) (*entities.Roles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoles", input)
	ret0, _ := ret[0].(*entities.Roles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRoles indicates an expected call of CreateRoles
func (mr *MockRepoMockRecorder) CreateRoles(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoles", reflect.TypeOf((*MockRepo)(nil).CreateRoles), input)
}

// FindOneUser mocks base method
func (m *MockRepo) FindOneUser(filter *entities.UsersFilter) (*entities.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneUser", filter)
	ret0, _ := ret[0].(*entities.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneUser indicates an expected call of FindOneUser
func (mr *MockRepoMockRecorder) FindOneUser(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneUser", reflect.TypeOf((*MockRepo)(nil).FindOneUser), filter)
}

// FindOneRole mocks base method
func (m *MockRepo) FindOneRole(filter *entities.RolesFilter) (*entities.Roles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneRole", filter)
	ret0, _ := ret[0].(*entities.Roles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneRole indicates an expected call of FindOneRole
func (mr *MockRepoMockRecorder) FindOneRole(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneRole", reflect.TypeOf((*MockRepo)(nil).FindOneRole), filter)
}

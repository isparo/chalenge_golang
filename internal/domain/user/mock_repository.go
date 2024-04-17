// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go

// Package user is a generated GoMock package.
package user

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// GetUserInfoByLogin mocks base method.
func (m *MockUserRepository) GetUserInfoByLogin(user, password string) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfoByLogin", user, password)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfoByLogin indicates an expected call of GetUserInfoByLogin.
func (mr *MockUserRepositoryMockRecorder) GetUserInfoByLogin(user, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfoByLogin", reflect.TypeOf((*MockUserRepository)(nil).GetUserInfoByLogin), user, password)
}

// SaveUser mocks base method.
func (m *MockUserRepository) SaveUser(user User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUser indicates an expected call of SaveUser.
func (mr *MockUserRepositoryMockRecorder) SaveUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockUserRepository)(nil).SaveUser), user)
}

// ValidateUserUniqueInfo mocks base method.
func (m *MockUserRepository) ValidateUserUniqueInfo(phone, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateUserUniqueInfo", phone, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateUserUniqueInfo indicates an expected call of ValidateUserUniqueInfo.
func (mr *MockUserRepositoryMockRecorder) ValidateUserUniqueInfo(phone, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateUserUniqueInfo", reflect.TypeOf((*MockUserRepository)(nil).ValidateUserUniqueInfo), phone, email)
}

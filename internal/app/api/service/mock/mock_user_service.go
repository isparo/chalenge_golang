// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_service.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	user_entity "github.com/josue/chalenge_golang/internal/domain/user"
)

// MockuserDomainService is a mock of userDomainService interface.
type MockuserDomainService struct {
	ctrl     *gomock.Controller
	recorder *MockuserDomainServiceMockRecorder
}

// MockuserDomainServiceMockRecorder is the mock recorder for MockuserDomainService.
type MockuserDomainServiceMockRecorder struct {
	mock *MockuserDomainService
}

// NewMockuserDomainService creates a new mock instance.
func NewMockuserDomainService(ctrl *gomock.Controller) *MockuserDomainService {
	mock := &MockuserDomainService{ctrl: ctrl}
	mock.recorder = &MockuserDomainServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockuserDomainService) EXPECT() *MockuserDomainServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockuserDomainService) CreateUser(userData user_entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", userData)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockuserDomainServiceMockRecorder) CreateUser(userData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockuserDomainService)(nil).CreateUser), userData)
}

// GetUserInfoByLogin mocks base method.
func (m *MockuserDomainService) GetUserInfoByLogin(user, password string) (*user_entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfoByLogin", user, password)
	ret0, _ := ret[0].(*user_entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfoByLogin indicates an expected call of GetUserInfoByLogin.
func (mr *MockuserDomainServiceMockRecorder) GetUserInfoByLogin(user, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfoByLogin", reflect.TypeOf((*MockuserDomainService)(nil).GetUserInfoByLogin), user, password)
}

// ValidateUserByPhoneAndEmail mocks base method.
func (m *MockuserDomainService) ValidateUserByPhoneAndEmail(phone, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateUserByPhoneAndEmail", phone, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateUserByPhoneAndEmail indicates an expected call of ValidateUserByPhoneAndEmail.
func (mr *MockuserDomainServiceMockRecorder) ValidateUserByPhoneAndEmail(phone, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateUserByPhoneAndEmail", reflect.TypeOf((*MockuserDomainService)(nil).ValidateUserByPhoneAndEmail), phone, email)
}
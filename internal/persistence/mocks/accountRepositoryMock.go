// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yaroslavnayug/go-payment-system/internal/domain/model (interfaces: AccountRepositoryInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/yaroslavnayug/go-payment-system/internal/domain/model"
	reflect "reflect"
)

// MockAccountRepositoryInterface is a mock of AccountRepositoryInterface interface
type MockAccountRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryInterfaceMockRecorder
}

// MockAccountRepositoryInterfaceMockRecorder is the mock recorder for MockAccountRepositoryInterface
type MockAccountRepositoryInterfaceMockRecorder struct {
	mock *MockAccountRepositoryInterface
}

// NewMockAccountRepositoryInterface creates a new mock instance
func NewMockAccountRepositoryInterface(ctrl *gomock.Controller) *MockAccountRepositoryInterface {
	mock := &MockAccountRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountRepositoryInterface) EXPECT() *MockAccountRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method
func (m *MockAccountRepositoryInterface) CreateAccount(arg0 model.Account) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockAccountRepositoryInterfaceMockRecorder) CreateAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountRepositoryInterface)(nil).CreateAccount), arg0)
}

// GetAccountByPassportData mocks base method
func (m *MockAccountRepositoryInterface) GetAccountByPassportData(arg0 string) (*model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByPassportData", arg0)
	ret0, _ := ret[0].(*model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByPassportData indicates an expected call of GetAccountByPassportData
func (mr *MockAccountRepositoryInterfaceMockRecorder) GetAccountByPassportData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByPassportData", reflect.TypeOf((*MockAccountRepositoryInterface)(nil).GetAccountByPassportData), arg0)
}
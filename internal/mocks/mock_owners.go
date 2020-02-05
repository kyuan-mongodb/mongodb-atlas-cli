// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/owners.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	cloudmanager "github.com/mongodb-labs/pcgc/cloudmanager"
	reflect "reflect"
)

// MockOwnerCreator is a mock of OwnerCreator interface
type MockOwnerCreator struct {
	ctrl     *gomock.Controller
	recorder *MockOwnerCreatorMockRecorder
}

// MockOwnerCreatorMockRecorder is the mock recorder for MockOwnerCreator
type MockOwnerCreatorMockRecorder struct {
	mock *MockOwnerCreator
}

// NewMockOwnerCreator creates a new mock instance
func NewMockOwnerCreator(ctrl *gomock.Controller) *MockOwnerCreator {
	mock := &MockOwnerCreator{ctrl: ctrl}
	mock.recorder = &MockOwnerCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOwnerCreator) EXPECT() *MockOwnerCreatorMockRecorder {
	return m.recorder
}

// CreateOwner mocks base method
func (m *MockOwnerCreator) CreateOwner(arg0 *cloudmanager.User, arg1 string) (*cloudmanager.CreateUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOwner", arg0, arg1)
	ret0, _ := ret[0].(*cloudmanager.CreateUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOwner indicates an expected call of CreateOwner
func (mr *MockOwnerCreatorMockRecorder) CreateOwner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOwner", reflect.TypeOf((*MockOwnerCreator)(nil).CreateOwner), arg0, arg1)
}
// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/core/app/signal/service (interfaces: Manager)

// Package mockservice is a generated GoMock package.
package mockservice

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// StopAll mocks base method
func (m *MockManager) StopAll() {
	m.ctrl.Call(m, "StopAll")
}

// StopAll indicates an expected call of StopAll
func (mr *MockManagerMockRecorder) StopAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopAll", reflect.TypeOf((*MockManager)(nil).StopAll))
}

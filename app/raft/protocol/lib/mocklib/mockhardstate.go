// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/go-node/app/raft/protocol/lib (interfaces: HardState)

// Package mocklib is a generated GoMock package.
package mocklib

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHardState is a mock of HardState interface
type MockHardState struct {
	ctrl     *gomock.Controller
	recorder *MockHardStateMockRecorder
}

// MockHardStateMockRecorder is the mock recorder for MockHardState
type MockHardStateMockRecorder struct {
	mock *MockHardState
}

// NewMockHardState creates a new mock instance
func NewMockHardState(ctrl *gomock.Controller) *MockHardState {
	mock := &MockHardState{ctrl: ctrl}
	mock.recorder = &MockHardStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHardState) EXPECT() *MockHardStateMockRecorder {
	return m.recorder
}

// Commit mocks base method
func (m *MockHardState) Commit() uint64 {
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockHardStateMockRecorder) Commit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockHardState)(nil).Commit))
}

// Term mocks base method
func (m *MockHardState) Term() uint64 {
	ret := m.ctrl.Call(m, "Term")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Term indicates an expected call of Term
func (mr *MockHardStateMockRecorder) Term() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Term", reflect.TypeOf((*MockHardState)(nil).Term))
}

// Vote mocks base method
func (m *MockHardState) Vote() uint64 {
	ret := m.ctrl.Call(m, "Vote")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Vote indicates an expected call of Vote
func (mr *MockHardStateMockRecorder) Vote() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vote", reflect.TypeOf((*MockHardState)(nil).Vote))
}

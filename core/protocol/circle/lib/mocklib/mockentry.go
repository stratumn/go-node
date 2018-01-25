// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/core/protocol/circle/lib (interfaces: Entry)

// Package mocklib is a generated GoMock package.
package mocklib

import (
	gomock "github.com/golang/mock/gomock"
	lib "github.com/stratumn/alice/core/protocol/circle/lib"
	reflect "reflect"
)

// MockEntry is a mock of Entry interface
type MockEntry struct {
	ctrl     *gomock.Controller
	recorder *MockEntryMockRecorder
}

// MockEntryMockRecorder is the mock recorder for MockEntry
type MockEntryMockRecorder struct {
	mock *MockEntry
}

// NewMockEntry creates a new mock instance
func NewMockEntry(ctrl *gomock.Controller) *MockEntry {
	mock := &MockEntry{ctrl: ctrl}
	mock.recorder = &MockEntryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEntry) EXPECT() *MockEntryMockRecorder {
	return m.recorder
}

// Data mocks base method
func (m *MockEntry) Data() []byte {
	ret := m.ctrl.Call(m, "Data")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Data indicates an expected call of Data
func (mr *MockEntryMockRecorder) Data() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*MockEntry)(nil).Data))
}

// Index mocks base method
func (m *MockEntry) Index() uint64 {
	ret := m.ctrl.Call(m, "Index")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Index indicates an expected call of Index
func (mr *MockEntryMockRecorder) Index() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockEntry)(nil).Index))
}

// Term mocks base method
func (m *MockEntry) Term() uint64 {
	ret := m.ctrl.Call(m, "Term")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Term indicates an expected call of Term
func (mr *MockEntryMockRecorder) Term() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Term", reflect.TypeOf((*MockEntry)(nil).Term))
}

// Type mocks base method
func (m *MockEntry) Type() lib.EntryType {
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(lib.EntryType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockEntryMockRecorder) Type() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockEntry)(nil).Type))
}

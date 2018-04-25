// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/go-indigocore/fossilizer (interfaces: Adapter)

// Package mockfossilizer is a generated GoMock package.
package mockfossilizer

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	fossilizer "github.com/stratumn/alice/vendor/github.com/stratumn/go-indigocore/fossilizer"
	reflect "reflect"
)

// MockAdapter is a mock of Adapter interface
type MockAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockAdapterMockRecorder
}

// MockAdapterMockRecorder is the mock recorder for MockAdapter
type MockAdapterMockRecorder struct {
	mock *MockAdapter
}

// NewMockAdapter creates a new mock instance
func NewMockAdapter(ctrl *gomock.Controller) *MockAdapter {
	mock := &MockAdapter{ctrl: ctrl}
	mock.recorder = &MockAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdapter) EXPECT() *MockAdapterMockRecorder {
	return m.recorder
}

// AddFossilizerEventChan mocks base method
func (m *MockAdapter) AddFossilizerEventChan(arg0 chan *fossilizer.Event) {
	m.ctrl.Call(m, "AddFossilizerEventChan", arg0)
}

// AddFossilizerEventChan indicates an expected call of AddFossilizerEventChan
func (mr *MockAdapterMockRecorder) AddFossilizerEventChan(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFossilizerEventChan", reflect.TypeOf((*MockAdapter)(nil).AddFossilizerEventChan), arg0)
}

// Fossilize mocks base method
func (m *MockAdapter) Fossilize(arg0 context.Context, arg1, arg2 []byte) error {
	ret := m.ctrl.Call(m, "Fossilize", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Fossilize indicates an expected call of Fossilize
func (mr *MockAdapterMockRecorder) Fossilize(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fossilize", reflect.TypeOf((*MockAdapter)(nil).Fossilize), arg0, arg1, arg2)
}

// GetInfo mocks base method
func (m *MockAdapter) GetInfo(arg0 context.Context) (interface{}, error) {
	ret := m.ctrl.Call(m, "GetInfo", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo
func (mr *MockAdapterMockRecorder) GetInfo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockAdapter)(nil).GetInfo), arg0)
}

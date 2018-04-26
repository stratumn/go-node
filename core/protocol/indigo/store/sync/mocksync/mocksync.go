// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/core/protocol/indigo/store/sync (interfaces: Engine)

package mocksync

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	cs "github.com/stratumn/go-indigocore/cs"
	store "github.com/stratumn/go-indigocore/store"
)

// MockEngine is a mock of Engine interface
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return _m.recorder
}

// Close mocks base method
func (_m *MockEngine) Close(_param0 context.Context) {
	_m.ctrl.Call(_m, "Close", _param0)
}

// Close indicates an expected call of Close
func (_mr *MockEngineMockRecorder) Close(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockEngine)(nil).Close), arg0)
}

// GetMissingLinks mocks base method
func (_m *MockEngine) GetMissingLinks(_param0 context.Context, _param1 *cs.Link, _param2 store.SegmentReader) ([]*cs.Link, error) {
	ret := _m.ctrl.Call(_m, "GetMissingLinks", _param0, _param1, _param2)
	ret0, _ := ret[0].([]*cs.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMissingLinks indicates an expected call of GetMissingLinks
func (_mr *MockEngineMockRecorder) GetMissingLinks(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetMissingLinks", reflect.TypeOf((*MockEngine)(nil).GetMissingLinks), arg0, arg1, arg2)
}

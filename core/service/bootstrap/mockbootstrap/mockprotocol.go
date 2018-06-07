// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/core/protocol/bootstrap (interfaces: Handler)

package mockbootstrap

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	go_multiaddr "gx/ipfs/QmWWQ2Txc2c6tqjsBpzg5Ar652cHPGNsQQp2SejkNmkUMb/go-multiaddr"
	go_libp2p_peer "gx/ipfs/QmcJukH2sAFjY3HdBKq35WDzWoL3UUu2gt9wdfqZTUyM74/go-libp2p-peer"
	reflect "reflect"
)

// MockHandler is a mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return _m.recorder
}

// Accept mocks base method
func (_m *MockHandler) Accept(_param0 context.Context, _param1 go_libp2p_peer.ID) error {
	ret := _m.ctrl.Call(_m, "Accept", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Accept indicates an expected call of Accept
func (_mr *MockHandlerMockRecorder) Accept(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Accept", reflect.TypeOf((*MockHandler)(nil).Accept), arg0, arg1)
}

// AddNode mocks base method
func (_m *MockHandler) AddNode(_param0 context.Context, _param1 go_libp2p_peer.ID, _param2 go_multiaddr.Multiaddr, _param3 []byte) error {
	ret := _m.ctrl.Call(_m, "AddNode", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNode indicates an expected call of AddNode
func (_mr *MockHandlerMockRecorder) AddNode(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "AddNode", reflect.TypeOf((*MockHandler)(nil).AddNode), arg0, arg1, arg2, arg3)
}

// Close mocks base method
func (_m *MockHandler) Close(_param0 context.Context) {
	_m.ctrl.Call(_m, "Close", _param0)
}

// Close indicates an expected call of Close
func (_mr *MockHandlerMockRecorder) Close(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockHandler)(nil).Close), arg0)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/go-node/core/app/bootstrap/protocol (interfaces: Handler)

// Package mockprotocol is a generated GoMock package.
package mockprotocol

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	go_libp2p_peer "github.com/libp2p/go-libp2p-peer"
	go_multiaddr "github.com/multiformats/go-multiaddr"
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
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockHandler) Accept(arg0 context.Context, arg1 go_libp2p_peer.ID) error {
	ret := m.ctrl.Call(m, "Accept", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockHandlerMockRecorder) Accept(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockHandler)(nil).Accept), arg0, arg1)
}

// AddNode mocks base method
func (m *MockHandler) AddNode(arg0 context.Context, arg1 go_libp2p_peer.ID, arg2 go_multiaddr.Multiaddr, arg3 []byte) error {
	ret := m.ctrl.Call(m, "AddNode", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNode indicates an expected call of AddNode
func (mr *MockHandlerMockRecorder) AddNode(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNode", reflect.TypeOf((*MockHandler)(nil).AddNode), arg0, arg1, arg2, arg3)
}

// Close mocks base method
func (m *MockHandler) Close(arg0 context.Context) {
	m.ctrl.Call(m, "Close", arg0)
}

// Close indicates an expected call of Close
func (mr *MockHandlerMockRecorder) Close(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockHandler)(nil).Close), arg0)
}

// CompleteBootstrap mocks base method
func (m *MockHandler) CompleteBootstrap(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "CompleteBootstrap", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteBootstrap indicates an expected call of CompleteBootstrap
func (mr *MockHandlerMockRecorder) CompleteBootstrap(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteBootstrap", reflect.TypeOf((*MockHandler)(nil).CompleteBootstrap), arg0)
}

// Handshake mocks base method
func (m *MockHandler) Handshake(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Handshake", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handshake indicates an expected call of Handshake
func (mr *MockHandlerMockRecorder) Handshake(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handshake", reflect.TypeOf((*MockHandler)(nil).Handshake), arg0)
}

// Reject mocks base method
func (m *MockHandler) Reject(arg0 context.Context, arg1 go_libp2p_peer.ID) error {
	ret := m.ctrl.Call(m, "Reject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reject indicates an expected call of Reject
func (mr *MockHandlerMockRecorder) Reject(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reject", reflect.TypeOf((*MockHandler)(nil).Reject), arg0, arg1)
}

// RemoveNode mocks base method
func (m *MockHandler) RemoveNode(arg0 context.Context, arg1 go_libp2p_peer.ID) error {
	ret := m.ctrl.Call(m, "RemoveNode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNode indicates an expected call of RemoveNode
func (mr *MockHandlerMockRecorder) RemoveNode(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNode", reflect.TypeOf((*MockHandler)(nil).RemoveNode), arg0, arg1)
}

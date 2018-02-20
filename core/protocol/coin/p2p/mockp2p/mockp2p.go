// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/core/protocol/coin/p2p (interfaces: P2P)

// Package mockchain is a generated GoMock package.
package mockchain

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	coin "github.com/stratumn/alice/pb/coin"
	go_libp2p_peer "gx/ipfs/Qma7H6RW8wRrfZpNSXwxYGcd1E149s42FpWNpDNieSVrnU/go-libp2p-peer"
	reflect "reflect"
)

// MockP2P is a mock of P2P interface
type MockP2P struct {
	ctrl     *gomock.Controller
	recorder *MockP2PMockRecorder
}

// MockP2PMockRecorder is the mock recorder for MockP2P
type MockP2PMockRecorder struct {
	mock *MockP2P
}

// NewMockP2P creates a new mock instance
func NewMockP2P(ctrl *gomock.Controller) *MockP2P {
	mock := &MockP2P{ctrl: ctrl}
	mock.recorder = &MockP2PMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockP2P) EXPECT() *MockP2PMockRecorder {
	return m.recorder
}

// RequestBlockByHash mocks base method
func (m *MockP2P) RequestBlockByHash(arg0 context.Context, arg1 go_libp2p_peer.ID, arg2 []byte) (*coin.Block, error) {
	ret := m.ctrl.Call(m, "RequestBlockByHash", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coin.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestBlockByHash indicates an expected call of RequestBlockByHash
func (mr *MockP2PMockRecorder) RequestBlockByHash(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestBlockByHash", reflect.TypeOf((*MockP2P)(nil).RequestBlockByHash), arg0, arg1, arg2)
}

// RequestBlocksByNumber mocks base method
func (m *MockP2P) RequestBlocksByNumber(arg0 context.Context, arg1 go_libp2p_peer.ID, arg2, arg3 uint64) ([]*coin.Block, error) {
	ret := m.ctrl.Call(m, "RequestBlocksByNumber", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*coin.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestBlocksByNumber indicates an expected call of RequestBlocksByNumber
func (mr *MockP2PMockRecorder) RequestBlocksByNumber(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestBlocksByNumber", reflect.TypeOf((*MockP2P)(nil).RequestBlocksByNumber), arg0, arg1, arg2, arg3)
}

// RequestHeaderByHash mocks base method
func (m *MockP2P) RequestHeaderByHash(arg0 context.Context, arg1 go_libp2p_peer.ID, arg2 []byte) (*coin.Header, error) {
	ret := m.ctrl.Call(m, "RequestHeaderByHash", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coin.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestHeaderByHash indicates an expected call of RequestHeaderByHash
func (mr *MockP2PMockRecorder) RequestHeaderByHash(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestHeaderByHash", reflect.TypeOf((*MockP2P)(nil).RequestHeaderByHash), arg0, arg1, arg2)
}

// RequestHeadersByNumber mocks base method
func (m *MockP2P) RequestHeadersByNumber(arg0 context.Context, arg1 go_libp2p_peer.ID, arg2, arg3 uint64) ([]*coin.Header, error) {
	ret := m.ctrl.Call(m, "RequestHeadersByNumber", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*coin.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestHeadersByNumber indicates an expected call of RequestHeadersByNumber
func (mr *MockP2PMockRecorder) RequestHeadersByNumber(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestHeadersByNumber", reflect.TypeOf((*MockP2P)(nil).RequestHeadersByNumber), arg0, arg1, arg2, arg3)
}

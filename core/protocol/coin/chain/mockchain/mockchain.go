// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/core/protocol/coin/chain (interfaces: Chain)

// Package mockchain is a generated GoMock package.
package mockchain

import (
	gomock "github.com/golang/mock/gomock"
	chain "github.com/stratumn/alice/core/protocol/coin/chain"
	coin "github.com/stratumn/alice/pb/coin"
	reflect "reflect"
)

// MockChain is a mock of Chain interface
type MockChain struct {
	ctrl     *gomock.Controller
	recorder *MockChainMockRecorder
}

// MockChainMockRecorder is the mock recorder for MockChain
type MockChainMockRecorder struct {
	mock *MockChain
}

// NewMockChain creates a new mock instance
func NewMockChain(ctrl *gomock.Controller) *MockChain {
	mock := &MockChain{ctrl: ctrl}
	mock.recorder = &MockChainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChain) EXPECT() *MockChainMockRecorder {
	return m.recorder
}

// AddBlock mocks base method
func (m *MockChain) AddBlock(arg0 *coin.Block) error {
	ret := m.ctrl.Call(m, "AddBlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBlock indicates an expected call of AddBlock
func (mr *MockChainMockRecorder) AddBlock(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBlock", reflect.TypeOf((*MockChain)(nil).AddBlock), arg0)
}

// Config mocks base method
func (m *MockChain) Config() *chain.Config {
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*chain.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockChainMockRecorder) Config() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockChain)(nil).Config))
}

// CurrentBlock mocks base method
func (m *MockChain) CurrentBlock() (*coin.Block, error) {
	ret := m.ctrl.Call(m, "CurrentBlock")
	ret0, _ := ret[0].(*coin.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentBlock indicates an expected call of CurrentBlock
func (mr *MockChainMockRecorder) CurrentBlock() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentBlock", reflect.TypeOf((*MockChain)(nil).CurrentBlock))
}

// CurrentHeader mocks base method
func (m *MockChain) CurrentHeader() (*coin.Header, error) {
	ret := m.ctrl.Call(m, "CurrentHeader")
	ret0, _ := ret[0].(*coin.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentHeader indicates an expected call of CurrentHeader
func (mr *MockChainMockRecorder) CurrentHeader() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentHeader", reflect.TypeOf((*MockChain)(nil).CurrentHeader))
}

// GetBlock mocks base method
func (m *MockChain) GetBlock(arg0 []byte, arg1 uint64) (*coin.Block, error) {
	ret := m.ctrl.Call(m, "GetBlock", arg0, arg1)
	ret0, _ := ret[0].(*coin.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock
func (mr *MockChainMockRecorder) GetBlock(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockChain)(nil).GetBlock), arg0, arg1)
}

// GetBlockByHash mocks base method
func (m *MockChain) GetBlockByHash(arg0 []byte) (*coin.Block, error) {
	ret := m.ctrl.Call(m, "GetBlockByHash", arg0)
	ret0, _ := ret[0].(*coin.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByHash indicates an expected call of GetBlockByHash
func (mr *MockChainMockRecorder) GetBlockByHash(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHash", reflect.TypeOf((*MockChain)(nil).GetBlockByHash), arg0)
}

// GetBlockByNumber mocks base method
func (m *MockChain) GetBlockByNumber(arg0 uint64) (*coin.Block, error) {
	ret := m.ctrl.Call(m, "GetBlockByNumber", arg0)
	ret0, _ := ret[0].(*coin.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByNumber indicates an expected call of GetBlockByNumber
func (mr *MockChainMockRecorder) GetBlockByNumber(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByNumber", reflect.TypeOf((*MockChain)(nil).GetBlockByNumber), arg0)
}

// GetHeaderByHash mocks base method
func (m *MockChain) GetHeaderByHash(arg0 []byte) (*coin.Header, error) {
	ret := m.ctrl.Call(m, "GetHeaderByHash", arg0)
	ret0, _ := ret[0].(*coin.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeaderByHash indicates an expected call of GetHeaderByHash
func (mr *MockChainMockRecorder) GetHeaderByHash(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderByHash", reflect.TypeOf((*MockChain)(nil).GetHeaderByHash), arg0)
}

// GetHeaderByNumber mocks base method
func (m *MockChain) GetHeaderByNumber(arg0 uint64) (*coin.Header, error) {
	ret := m.ctrl.Call(m, "GetHeaderByNumber", arg0)
	ret0, _ := ret[0].(*coin.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeaderByNumber indicates an expected call of GetHeaderByNumber
func (mr *MockChainMockRecorder) GetHeaderByNumber(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderByNumber", reflect.TypeOf((*MockChain)(nil).GetHeaderByNumber), arg0)
}

// GetHeadersByNumber mocks base method
func (m *MockChain) GetHeadersByNumber(arg0 uint64) ([]*coin.Header, error) {
	ret := m.ctrl.Call(m, "GetHeadersByNumber", arg0)
	ret0, _ := ret[0].([]*coin.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeadersByNumber indicates an expected call of GetHeadersByNumber
func (mr *MockChainMockRecorder) GetHeadersByNumber(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeadersByNumber", reflect.TypeOf((*MockChain)(nil).GetHeadersByNumber), arg0)
}

// GetParentBlock mocks base method
func (m *MockChain) GetParentBlock(arg0 *coin.Header) (*coin.Block, error) {
	ret := m.ctrl.Call(m, "GetParentBlock", arg0)
	ret0, _ := ret[0].(*coin.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParentBlock indicates an expected call of GetParentBlock
func (mr *MockChainMockRecorder) GetParentBlock(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParentBlock", reflect.TypeOf((*MockChain)(nil).GetParentBlock), arg0)
}

// SetHead mocks base method
func (m *MockChain) SetHead(arg0 *coin.Block) error {
	ret := m.ctrl.Call(m, "SetHead", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHead indicates an expected call of SetHead
func (mr *MockChainMockRecorder) SetHead(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHead", reflect.TypeOf((*MockChain)(nil).SetHead), arg0)
}

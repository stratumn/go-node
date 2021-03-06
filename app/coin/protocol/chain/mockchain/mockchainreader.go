// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/go-node/app/coin/protocol/chain (interfaces: Reader)

// Package mockchain is a generated GoMock package.
package mockchain

import (
	gomock "github.com/golang/mock/gomock"
	pb "github.com/stratumn/go-node/app/coin/pb"
	chain "github.com/stratumn/go-node/app/coin/protocol/chain"
	reflect "reflect"
)

// MockReader is a mock of Reader interface
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// Config mocks base method
func (m *MockReader) Config() *chain.Config {
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*chain.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockReaderMockRecorder) Config() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockReader)(nil).Config))
}

// CurrentBlock mocks base method
func (m *MockReader) CurrentBlock() (*pb.Block, error) {
	ret := m.ctrl.Call(m, "CurrentBlock")
	ret0, _ := ret[0].(*pb.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentBlock indicates an expected call of CurrentBlock
func (mr *MockReaderMockRecorder) CurrentBlock() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentBlock", reflect.TypeOf((*MockReader)(nil).CurrentBlock))
}

// CurrentHeader mocks base method
func (m *MockReader) CurrentHeader() (*pb.Header, error) {
	ret := m.ctrl.Call(m, "CurrentHeader")
	ret0, _ := ret[0].(*pb.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentHeader indicates an expected call of CurrentHeader
func (mr *MockReaderMockRecorder) CurrentHeader() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentHeader", reflect.TypeOf((*MockReader)(nil).CurrentHeader))
}

// GetBlock mocks base method
func (m *MockReader) GetBlock(arg0 []byte, arg1 uint64) (*pb.Block, error) {
	ret := m.ctrl.Call(m, "GetBlock", arg0, arg1)
	ret0, _ := ret[0].(*pb.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock
func (mr *MockReaderMockRecorder) GetBlock(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockReader)(nil).GetBlock), arg0, arg1)
}

// GetBlockByHash mocks base method
func (m *MockReader) GetBlockByHash(arg0 []byte) (*pb.Block, error) {
	ret := m.ctrl.Call(m, "GetBlockByHash", arg0)
	ret0, _ := ret[0].(*pb.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByHash indicates an expected call of GetBlockByHash
func (mr *MockReaderMockRecorder) GetBlockByHash(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHash", reflect.TypeOf((*MockReader)(nil).GetBlockByHash), arg0)
}

// GetBlockByNumber mocks base method
func (m *MockReader) GetBlockByNumber(arg0 uint64) (*pb.Block, error) {
	ret := m.ctrl.Call(m, "GetBlockByNumber", arg0)
	ret0, _ := ret[0].(*pb.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByNumber indicates an expected call of GetBlockByNumber
func (mr *MockReaderMockRecorder) GetBlockByNumber(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByNumber", reflect.TypeOf((*MockReader)(nil).GetBlockByNumber), arg0)
}

// GetHeaderByHash mocks base method
func (m *MockReader) GetHeaderByHash(arg0 []byte) (*pb.Header, error) {
	ret := m.ctrl.Call(m, "GetHeaderByHash", arg0)
	ret0, _ := ret[0].(*pb.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeaderByHash indicates an expected call of GetHeaderByHash
func (mr *MockReaderMockRecorder) GetHeaderByHash(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderByHash", reflect.TypeOf((*MockReader)(nil).GetHeaderByHash), arg0)
}

// GetHeaderByNumber mocks base method
func (m *MockReader) GetHeaderByNumber(arg0 uint64) (*pb.Header, error) {
	ret := m.ctrl.Call(m, "GetHeaderByNumber", arg0)
	ret0, _ := ret[0].(*pb.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeaderByNumber indicates an expected call of GetHeaderByNumber
func (mr *MockReaderMockRecorder) GetHeaderByNumber(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderByNumber", reflect.TypeOf((*MockReader)(nil).GetHeaderByNumber), arg0)
}

// GetHeadersByNumber mocks base method
func (m *MockReader) GetHeadersByNumber(arg0 uint64) ([]*pb.Header, error) {
	ret := m.ctrl.Call(m, "GetHeadersByNumber", arg0)
	ret0, _ := ret[0].([]*pb.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeadersByNumber indicates an expected call of GetHeadersByNumber
func (mr *MockReaderMockRecorder) GetHeadersByNumber(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeadersByNumber", reflect.TypeOf((*MockReader)(nil).GetHeadersByNumber), arg0)
}

// GetParentBlock mocks base method
func (m *MockReader) GetParentBlock(arg0 *pb.Header) (*pb.Block, error) {
	ret := m.ctrl.Call(m, "GetParentBlock", arg0)
	ret0, _ := ret[0].(*pb.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParentBlock indicates an expected call of GetParentBlock
func (mr *MockReaderMockRecorder) GetParentBlock(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParentBlock", reflect.TypeOf((*MockReader)(nil).GetParentBlock), arg0)
}

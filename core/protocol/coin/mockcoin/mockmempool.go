// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/core/protocol/coin (interfaces: Mempool)

package mockcoin

import (
	gomock "github.com/golang/mock/gomock"
	coin "github.com/stratumn/alice/pb/coin"
	reflect "reflect"
)

// MockMempool is a mock of Mempool interface
type MockMempool struct {
	ctrl     *gomock.Controller
	recorder *MockMempoolMockRecorder
}

// MockMempoolMockRecorder is the mock recorder for MockMempool
type MockMempoolMockRecorder struct {
	mock *MockMempool
}

// NewMockMempool creates a new mock instance
func NewMockMempool(ctrl *gomock.Controller) *MockMempool {
	mock := &MockMempool{ctrl: ctrl}
	mock.recorder = &MockMempoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockMempool) EXPECT() *MockMempoolMockRecorder {
	return _m.recorder
}

// AddTransaction mocks base method
func (_m *MockMempool) AddTransaction(_param0 *coin.Transaction) error {
	ret := _m.ctrl.Call(_m, "AddTransaction", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTransaction indicates an expected call of AddTransaction
func (_mr *MockMempoolMockRecorder) AddTransaction(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "AddTransaction", reflect.TypeOf((*MockMempool)(nil).AddTransaction), arg0)
}

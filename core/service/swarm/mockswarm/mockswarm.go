// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/core/service/swarm (interfaces: Transport)

package mockswarm

import (
	gomock "github.com/golang/mock/gomock"
	go_stream_muxer "gx/ipfs/QmY9JXR3FupnYAYJWK9aMr9bCpqWKcToQ1tz8DVGTrHpHw/go-stream-muxer"
	net "net"
	reflect "reflect"
)

// MockTransport is a mock of Transport interface
type MockTransport struct {
	ctrl     *gomock.Controller
	recorder *MockTransportMockRecorder
}

// MockTransportMockRecorder is the mock recorder for MockTransport
type MockTransportMockRecorder struct {
	mock *MockTransport
}

// NewMockTransport creates a new mock instance
func NewMockTransport(ctrl *gomock.Controller) *MockTransport {
	mock := &MockTransport{ctrl: ctrl}
	mock.recorder = &MockTransportMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockTransport) EXPECT() *MockTransportMockRecorder {
	return _m.recorder
}

// NewConn mocks base method
func (_m *MockTransport) NewConn(_param0 net.Conn, _param1 bool) (go_stream_muxer.Conn, error) {
	ret := _m.ctrl.Call(_m, "NewConn", _param0, _param1)
	ret0, _ := ret[0].(go_stream_muxer.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewConn indicates an expected call of NewConn
func (_mr *MockTransportMockRecorder) NewConn(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "NewConn", reflect.TypeOf((*MockTransport)(nil).NewConn), arg0, arg1)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: gx/ipfs/QmfZTdmunzKzAGJrSvXXQbQ5kLLUiEMX5vdwux7iXkdk7D/go-libp2p-host (interfaces: Host)

package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	go_multistream "gx/ipfs/QmTnsezaB1wWNRHeHnYrm8K4d5i9wtyj3GsqjC3Rt5b5v5/go-multistream"
	go_multiaddr "gx/ipfs/QmWWQ2Txc2c6tqjsBpzg5Ar652cHPGNsQQp2SejkNmkUMb/go-multiaddr"
	go_libp2p_net "gx/ipfs/QmXoz9o2PT3tEzf7hicegwex5UgVP54n3k82K7jrWFyN86/go-libp2p-net"
	go_libp2p_protocol "gx/ipfs/QmZNkThpqfVXs9GNbexPrfBbXSLNYeKrE7jwFM2oqHbyqN/go-libp2p-protocol"
	go_libp2p_peer "gx/ipfs/QmcJukH2sAFjY3HdBKq35WDzWoL3UUu2gt9wdfqZTUyM74/go-libp2p-peer"
	go_libp2p_peerstore "gx/ipfs/QmdeiKhUy1TVGBaKxt7y1QmBDLBdisSrLJ1x58Eoj4PXUh/go-libp2p-peerstore"
	go_libp2p_interface_connmgr "gx/ipfs/QmfQNieWBPwmnUjXWPZbjJPzhNwFFabTb5RQ79dyVWGujQ/go-libp2p-interface-connmgr"
	reflect "reflect"
)

// MockHost is a mock of Host interface
type MockHost struct {
	ctrl     *gomock.Controller
	recorder *MockHostMockRecorder
}

// MockHostMockRecorder is the mock recorder for MockHost
type MockHostMockRecorder struct {
	mock *MockHost
}

// NewMockHost creates a new mock instance
func NewMockHost(ctrl *gomock.Controller) *MockHost {
	mock := &MockHost{ctrl: ctrl}
	mock.recorder = &MockHostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockHost) EXPECT() *MockHostMockRecorder {
	return _m.recorder
}

// Addrs mocks base method
func (_m *MockHost) Addrs() []go_multiaddr.Multiaddr {
	ret := _m.ctrl.Call(_m, "Addrs")
	ret0, _ := ret[0].([]go_multiaddr.Multiaddr)
	return ret0
}

// Addrs indicates an expected call of Addrs
func (_mr *MockHostMockRecorder) Addrs() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Addrs", reflect.TypeOf((*MockHost)(nil).Addrs))
}

// Close mocks base method
func (_m *MockHost) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (_mr *MockHostMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockHost)(nil).Close))
}

// ConnManager mocks base method
func (_m *MockHost) ConnManager() go_libp2p_interface_connmgr.ConnManager {
	ret := _m.ctrl.Call(_m, "ConnManager")
	ret0, _ := ret[0].(go_libp2p_interface_connmgr.ConnManager)
	return ret0
}

// ConnManager indicates an expected call of ConnManager
func (_mr *MockHostMockRecorder) ConnManager() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ConnManager", reflect.TypeOf((*MockHost)(nil).ConnManager))
}

// Connect mocks base method
func (_m *MockHost) Connect(_param0 context.Context, _param1 go_libp2p_peerstore.PeerInfo) error {
	ret := _m.ctrl.Call(_m, "Connect", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (_mr *MockHostMockRecorder) Connect(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Connect", reflect.TypeOf((*MockHost)(nil).Connect), arg0, arg1)
}

// ID mocks base method
func (_m *MockHost) ID() go_libp2p_peer.ID {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(go_libp2p_peer.ID)
	return ret0
}

// ID indicates an expected call of ID
func (_mr *MockHostMockRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ID", reflect.TypeOf((*MockHost)(nil).ID))
}

// Mux mocks base method
func (_m *MockHost) Mux() *go_multistream.MultistreamMuxer {
	ret := _m.ctrl.Call(_m, "Mux")
	ret0, _ := ret[0].(*go_multistream.MultistreamMuxer)
	return ret0
}

// Mux indicates an expected call of Mux
func (_mr *MockHostMockRecorder) Mux() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Mux", reflect.TypeOf((*MockHost)(nil).Mux))
}

// Network mocks base method
func (_m *MockHost) Network() go_libp2p_net.Network {
	ret := _m.ctrl.Call(_m, "Network")
	ret0, _ := ret[0].(go_libp2p_net.Network)
	return ret0
}

// Network indicates an expected call of Network
func (_mr *MockHostMockRecorder) Network() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Network", reflect.TypeOf((*MockHost)(nil).Network))
}

// NewStream mocks base method
func (_m *MockHost) NewStream(_param0 context.Context, _param1 go_libp2p_peer.ID, _param2 ...go_libp2p_protocol.ID) (go_libp2p_net.Stream, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "NewStream", _s...)
	ret0, _ := ret[0].(go_libp2p_net.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewStream indicates an expected call of NewStream
func (_mr *MockHostMockRecorder) NewStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "NewStream", reflect.TypeOf((*MockHost)(nil).NewStream), _s...)
}

// Peerstore mocks base method
func (_m *MockHost) Peerstore() go_libp2p_peerstore.Peerstore {
	ret := _m.ctrl.Call(_m, "Peerstore")
	ret0, _ := ret[0].(go_libp2p_peerstore.Peerstore)
	return ret0
}

// Peerstore indicates an expected call of Peerstore
func (_mr *MockHostMockRecorder) Peerstore() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Peerstore", reflect.TypeOf((*MockHost)(nil).Peerstore))
}

// RemoveStreamHandler mocks base method
func (_m *MockHost) RemoveStreamHandler(_param0 go_libp2p_protocol.ID) {
	_m.ctrl.Call(_m, "RemoveStreamHandler", _param0)
}

// RemoveStreamHandler indicates an expected call of RemoveStreamHandler
func (_mr *MockHostMockRecorder) RemoveStreamHandler(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RemoveStreamHandler", reflect.TypeOf((*MockHost)(nil).RemoveStreamHandler), arg0)
}

// SetStreamHandler mocks base method
func (_m *MockHost) SetStreamHandler(_param0 go_libp2p_protocol.ID, _param1 go_libp2p_net.StreamHandler) {
	_m.ctrl.Call(_m, "SetStreamHandler", _param0, _param1)
}

// SetStreamHandler indicates an expected call of SetStreamHandler
func (_mr *MockHostMockRecorder) SetStreamHandler(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetStreamHandler", reflect.TypeOf((*MockHost)(nil).SetStreamHandler), arg0, arg1)
}

// SetStreamHandlerMatch mocks base method
func (_m *MockHost) SetStreamHandlerMatch(_param0 go_libp2p_protocol.ID, _param1 func(string) bool, _param2 go_libp2p_net.StreamHandler) {
	_m.ctrl.Call(_m, "SetStreamHandlerMatch", _param0, _param1, _param2)
}

// SetStreamHandlerMatch indicates an expected call of SetStreamHandlerMatch
func (_mr *MockHostMockRecorder) SetStreamHandlerMatch(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetStreamHandlerMatch", reflect.TypeOf((*MockHost)(nil).SetStreamHandlerMatch), arg0, arg1, arg2)
}
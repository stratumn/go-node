// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/cli (interfaces: CLI)

// Package mockcli is a generated GoMock package.
package mockcli

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	cli "github.com/stratumn/alice/cli"
	reflect "reflect"
)

// MockCLI is a mock of CLI interface
type MockCLI struct {
	ctrl     *gomock.Controller
	recorder *MockCLIMockRecorder
}

// MockCLIMockRecorder is the mock recorder for MockCLI
type MockCLIMockRecorder struct {
	mock *MockCLI
}

// NewMockCLI creates a new mock instance
func NewMockCLI(ctrl *gomock.Controller) *MockCLI {
	mock := &MockCLI{ctrl: ctrl}
	mock.recorder = &MockCLIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCLI) EXPECT() *MockCLIMockRecorder {
	return m.recorder
}

// Address mocks base method
func (m *MockCLI) Address() string {
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(string)
	return ret0
}

// Address indicates an expected call of Address
func (mr *MockCLIMockRecorder) Address() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockCLI)(nil).Address))
}

// Commands mocks base method
func (m *MockCLI) Commands() []cli.Cmd {
	ret := m.ctrl.Call(m, "Commands")
	ret0, _ := ret[0].([]cli.Cmd)
	return ret0
}

// Commands indicates an expected call of Commands
func (mr *MockCLIMockRecorder) Commands() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commands", reflect.TypeOf((*MockCLI)(nil).Commands))
}

// Config mocks base method
func (m *MockCLI) Config() cli.Config {
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(cli.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockCLIMockRecorder) Config() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockCLI)(nil).Config))
}

// Connect mocks base method
func (m *MockCLI) Connect(arg0 context.Context, arg1 string) error {
	ret := m.ctrl.Call(m, "Connect", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockCLIMockRecorder) Connect(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockCLI)(nil).Connect), arg0, arg1)
}

// Console mocks base method
func (m *MockCLI) Console() *cli.Console {
	ret := m.ctrl.Call(m, "Console")
	ret0, _ := ret[0].(*cli.Console)
	return ret0
}

// Console indicates an expected call of Console
func (mr *MockCLIMockRecorder) Console() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Console", reflect.TypeOf((*MockCLI)(nil).Console))
}

// DidJustExecute mocks base method
func (m *MockCLI) DidJustExecute() bool {
	ret := m.ctrl.Call(m, "DidJustExecute")
	ret0, _ := ret[0].(bool)
	return ret0
}

// DidJustExecute indicates an expected call of DidJustExecute
func (mr *MockCLIMockRecorder) DidJustExecute() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DidJustExecute", reflect.TypeOf((*MockCLI)(nil).DidJustExecute))
}

// Disconnect mocks base method
func (m *MockCLI) Disconnect() error {
	ret := m.ctrl.Call(m, "Disconnect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnect indicates an expected call of Disconnect
func (mr *MockCLIMockRecorder) Disconnect() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockCLI)(nil).Disconnect))
}

// Eval mocks base method
func (m *MockCLI) Eval(arg0 context.Context, arg1 string) error {
	ret := m.ctrl.Call(m, "Eval", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Eval indicates an expected call of Eval
func (mr *MockCLIMockRecorder) Eval(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eval", reflect.TypeOf((*MockCLI)(nil).Eval), arg0, arg1)
}

// Exec mocks base method
func (m *MockCLI) Exec(arg0 context.Context, arg1 string) {
	m.ctrl.Call(m, "Exec", arg0, arg1)
}

// Exec indicates an expected call of Exec
func (mr *MockCLIMockRecorder) Exec(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockCLI)(nil).Exec), arg0, arg1)
}

// Run mocks base method
func (m *MockCLI) Run(arg0 context.Context) {
	m.ctrl.Call(m, "Run", arg0)
}

// Run indicates an expected call of Run
func (mr *MockCLIMockRecorder) Run(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockCLI)(nil).Run), arg0)
}

// Suggest mocks base method
func (m *MockCLI) Suggest(arg0 cli.Content) []cli.Suggest {
	ret := m.ctrl.Call(m, "Suggest", arg0)
	ret0, _ := ret[0].([]cli.Suggest)
	return ret0
}

// Suggest indicates an expected call of Suggest
func (mr *MockCLIMockRecorder) Suggest(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suggest", reflect.TypeOf((*MockCLI)(nil).Suggest), arg0)
}

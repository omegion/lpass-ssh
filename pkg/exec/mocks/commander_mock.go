// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/omegion/bw-ssh/pkg/exec (interfaces: CommanderInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCommanderInterface is a mock of CommanderInterface interface
type MockCommanderInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCommanderInterfaceMockRecorder
}

// MockCommanderInterfaceMockRecorder is the mock recorder for MockCommanderInterface
type MockCommanderInterfaceMockRecorder struct {
	mock *MockCommanderInterface
}

// NewMockCommanderInterface creates a new mock instance
func NewMockCommanderInterface(ctrl *gomock.Controller) *MockCommanderInterface {
	mock := &MockCommanderInterface{ctrl: ctrl}
	mock.recorder = &MockCommanderInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommanderInterface) EXPECT() *MockCommanderInterfaceMockRecorder {
	return m.recorder
}

// Output mocks base method
func (m *MockCommanderInterface) Output(arg0 string, arg1 ...string) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Output", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Output indicates an expected call of Output
func (mr *MockCommanderInterfaceMockRecorder) Output(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockCommanderInterface)(nil).Output), varargs...)
}
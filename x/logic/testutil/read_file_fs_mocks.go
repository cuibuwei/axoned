// Code generated by MockGen. DO NOT EDIT.
// Source: io/fs (interfaces: ReadFileFS)

// Package testutil is a generated GoMock package.
package testutil

import (
	fs "io/fs"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockReadFileFS is a mock of ReadFileFS interface.
type MockReadFileFS struct {
	ctrl     *gomock.Controller
	recorder *MockReadFileFSMockRecorder
}

// MockReadFileFSMockRecorder is the mock recorder for MockReadFileFS.
type MockReadFileFSMockRecorder struct {
	mock *MockReadFileFS
}

// NewMockReadFileFS creates a new mock instance.
func NewMockReadFileFS(ctrl *gomock.Controller) *MockReadFileFS {
	mock := &MockReadFileFS{ctrl: ctrl}
	mock.recorder = &MockReadFileFSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadFileFS) EXPECT() *MockReadFileFSMockRecorder {
	return m.recorder
}

// Open mocks base method.
func (m *MockReadFileFS) Open(arg0 string) (fs.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0)
	ret0, _ := ret[0].(fs.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockReadFileFSMockRecorder) Open(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockReadFileFS)(nil).Open), arg0)
}

// ReadFile mocks base method.
func (m *MockReadFileFS) ReadFile(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MockReadFileFSMockRecorder) ReadFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockReadFileFS)(nil).ReadFile), arg0)
}

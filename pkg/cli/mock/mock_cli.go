// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cli/cli.go

// Package mock_cli is a generated GoMock package.
package mock_cli

import (
	reflect "reflect"

	readline "github.com/chzyer/readline"
	gomock "github.com/golang/mock/gomock"
	preflight "github.com/replicatedhq/kurl/pkg/preflight"
	afero "github.com/spf13/afero"
	viper "github.com/spf13/viper"
)

// MockCLI is a mock of CLI interface.
type MockCLI struct {
	ctrl     *gomock.Controller
	recorder *MockCLIMockRecorder
}

// MockCLIMockRecorder is the mock recorder for MockCLI.
type MockCLIMockRecorder struct {
	mock *MockCLI
}

// NewMockCLI creates a new mock instance.
func NewMockCLI(ctrl *gomock.Controller) *MockCLI {
	mock := &MockCLI{ctrl: ctrl}
	mock.recorder = &MockCLIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCLI) EXPECT() *MockCLIMockRecorder {
	return m.recorder
}

// GetFS mocks base method.
func (m *MockCLI) GetFS() afero.Fs {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFS")
	ret0, _ := ret[0].(afero.Fs)
	return ret0
}

// GetFS indicates an expected call of GetFS.
func (mr *MockCLIMockRecorder) GetFS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFS", reflect.TypeOf((*MockCLI)(nil).GetFS))
}

// GetPreflightRunner mocks base method.
func (m *MockCLI) GetPreflightRunner() preflight.Runner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPreflightRunner")
	ret0, _ := ret[0].(preflight.Runner)
	return ret0
}

// GetPreflightRunner indicates an expected call of GetPreflightRunner.
func (mr *MockCLIMockRecorder) GetPreflightRunner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPreflightRunner", reflect.TypeOf((*MockCLI)(nil).GetPreflightRunner))
}

// GetReadline mocks base method.
func (m *MockCLI) GetReadline() *readline.Instance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReadline")
	ret0, _ := ret[0].(*readline.Instance)
	return ret0
}

// GetReadline indicates an expected call of GetReadline.
func (mr *MockCLIMockRecorder) GetReadline() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReadline", reflect.TypeOf((*MockCLI)(nil).GetReadline))
}

// GetViper mocks base method.
func (m *MockCLI) GetViper() *viper.Viper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetViper")
	ret0, _ := ret[0].(*viper.Viper)
	return ret0
}

// GetViper indicates an expected call of GetViper.
func (mr *MockCLIMockRecorder) GetViper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetViper", reflect.TypeOf((*MockCLI)(nil).GetViper))
}
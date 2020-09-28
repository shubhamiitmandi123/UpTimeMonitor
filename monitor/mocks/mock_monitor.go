// Code generated by MockGen. DO NOT EDIT.
// Source: up_time_monitor/moniter (interfaces: Monitor)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	structures "up_time_monitor/structures"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

// MockMonitor is a mock of Monitor interface
type MockMonitor struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorMockRecorder
}

// MockMonitorMockRecorder is the mock recorder for MockMonitor
type MockMonitorMockRecorder struct {
	mock *MockMonitor
}

// NewMockMonitor creates a new mock instance
func NewMockMonitor(ctrl *gomock.Controller) *MockMonitor {
	mock := &MockMonitor{ctrl: ctrl}
	mock.recorder = &MockMonitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMonitor) EXPECT() *MockMonitorMockRecorder {
	return m.recorder
}

// StartMonitoring mocks base method
func (m *MockMonitor) StartMonitoring(arg0 structures.URLInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartMonitoring", arg0)
}

// StartMonitoring indicates an expected call of StartMonitoring
func (mr *MockMonitorMockRecorder) StartMonitoring(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartMonitoring", reflect.TypeOf((*MockMonitor)(nil).StartMonitoring), arg0)
}

// StartMonitoringFromDatabase mocks base method
func (m *MockMonitor) StartMonitoringFromDatabase() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartMonitoringFromDatabase")
}

// StartMonitoringFromDatabase indicates an expected call of StartMonitoringFromDatabase
func (mr *MockMonitorMockRecorder) StartMonitoringFromDatabase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartMonitoringFromDatabase", reflect.TypeOf((*MockMonitor)(nil).StartMonitoringFromDatabase))
}

// StopMonitoring mocks base method
func (m *MockMonitor) StopMonitoring(arg0 uuid.UUID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopMonitoring", arg0)
}

// StopMonitoring indicates an expected call of StopMonitoring
func (mr *MockMonitorMockRecorder) StopMonitoring(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopMonitoring", reflect.TypeOf((*MockMonitor)(nil).StopMonitoring), arg0)
}
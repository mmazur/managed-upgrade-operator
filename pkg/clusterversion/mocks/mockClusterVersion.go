// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/clusterversion (interfaces: ClusterVersion)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/openshift/api/config/v1"
)

// MockClusterVersion is a mock of ClusterVersion interface
type MockClusterVersion struct {
	ctrl     *gomock.Controller
	recorder *MockClusterVersionMockRecorder
}

// MockClusterVersionMockRecorder is the mock recorder for MockClusterVersion
type MockClusterVersionMockRecorder struct {
	mock *MockClusterVersion
}

// NewMockClusterVersion creates a new mock instance
func NewMockClusterVersion(ctrl *gomock.Controller) *MockClusterVersion {
	mock := &MockClusterVersion{ctrl: ctrl}
	mock.recorder = &MockClusterVersionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterVersion) EXPECT() *MockClusterVersionMockRecorder {
	return m.recorder
}

// GetClusterVersion mocks base method
func (m *MockClusterVersion) GetClusterVersion() (*v1.ClusterVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterVersion")
	ret0, _ := ret[0].(*v1.ClusterVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterVersion indicates an expected call of GetClusterVersion
func (mr *MockClusterVersionMockRecorder) GetClusterVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterVersion", reflect.TypeOf((*MockClusterVersion)(nil).GetClusterVersion))
}

// HasUpgradeCommenced mocks base method
func (m *MockClusterVersion) HasUpgradeCommenced(arg0 *v1alpha1.UpgradeConfig) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUpgradeCommenced", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1

// HasUpgradeCommenced indicates an expected call of HasUpgradeCommenced
func (mr *MockClusterVersionMockRecorder) HasUpgradeCommenced(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUpgradeCommenced", reflect.TypeOf((*MockClusterVersion)(nil).HasUpgradeCommenced), arg0)
}
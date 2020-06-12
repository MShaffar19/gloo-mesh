// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/controller (interfaces: VirtualMeshEventWatcher,TrafficPolicyEventWatcher,AccessControlPolicyEventWatcher)

// Package mock_smh_networking is a generated GoMock package.
package mock_smh_networking

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/controller"
	reflect "reflect"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockVirtualMeshEventWatcher is a mock of VirtualMeshEventWatcher interface.
type MockVirtualMeshEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshEventWatcherMockRecorder
}

// MockVirtualMeshEventWatcherMockRecorder is the mock recorder for MockVirtualMeshEventWatcher.
type MockVirtualMeshEventWatcherMockRecorder struct {
	mock *MockVirtualMeshEventWatcher
}

// NewMockVirtualMeshEventWatcher creates a new mock instance.
func NewMockVirtualMeshEventWatcher(ctrl *gomock.Controller) *MockVirtualMeshEventWatcher {
	mock := &MockVirtualMeshEventWatcher{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshEventWatcher) EXPECT() *MockVirtualMeshEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockVirtualMeshEventWatcher) AddEventHandler(arg0 context.Context, arg1 controller.VirtualMeshEventHandler, arg2 ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockVirtualMeshEventWatcherMockRecorder) AddEventHandler(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockVirtualMeshEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockTrafficPolicyEventWatcher is a mock of TrafficPolicyEventWatcher interface.
type MockTrafficPolicyEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicyEventWatcherMockRecorder
}

// MockTrafficPolicyEventWatcherMockRecorder is the mock recorder for MockTrafficPolicyEventWatcher.
type MockTrafficPolicyEventWatcherMockRecorder struct {
	mock *MockTrafficPolicyEventWatcher
}

// NewMockTrafficPolicyEventWatcher creates a new mock instance.
func NewMockTrafficPolicyEventWatcher(ctrl *gomock.Controller) *MockTrafficPolicyEventWatcher {
	mock := &MockTrafficPolicyEventWatcher{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicyEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficPolicyEventWatcher) EXPECT() *MockTrafficPolicyEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockTrafficPolicyEventWatcher) AddEventHandler(arg0 context.Context, arg1 controller.TrafficPolicyEventHandler, arg2 ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockTrafficPolicyEventWatcherMockRecorder) AddEventHandler(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockTrafficPolicyEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockAccessControlPolicyEventWatcher is a mock of AccessControlPolicyEventWatcher interface.
type MockAccessControlPolicyEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockAccessControlPolicyEventWatcherMockRecorder
}

// MockAccessControlPolicyEventWatcherMockRecorder is the mock recorder for MockAccessControlPolicyEventWatcher.
type MockAccessControlPolicyEventWatcherMockRecorder struct {
	mock *MockAccessControlPolicyEventWatcher
}

// NewMockAccessControlPolicyEventWatcher creates a new mock instance.
func NewMockAccessControlPolicyEventWatcher(ctrl *gomock.Controller) *MockAccessControlPolicyEventWatcher {
	mock := &MockAccessControlPolicyEventWatcher{ctrl: ctrl}
	mock.recorder = &MockAccessControlPolicyEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessControlPolicyEventWatcher) EXPECT() *MockAccessControlPolicyEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockAccessControlPolicyEventWatcher) AddEventHandler(arg0 context.Context, arg1 controller.AccessControlPolicyEventHandler, arg2 ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockAccessControlPolicyEventWatcherMockRecorder) AddEventHandler(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockAccessControlPolicyEventWatcher)(nil).AddEventHandler), varargs...)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/external-apis/pkg/api/k8s/core/v1/controller (interfaces: ServiceEventWatcher)

// Package mock_corev1 is a generated GoMock package.
package mock_corev1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockServiceEventWatcher is a mock of ServiceEventWatcher interface.
type MockServiceEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceEventWatcherMockRecorder
}

// MockServiceEventWatcherMockRecorder is the mock recorder for MockServiceEventWatcher.
type MockServiceEventWatcherMockRecorder struct {
	mock *MockServiceEventWatcher
}

// NewMockServiceEventWatcher creates a new mock instance.
func NewMockServiceEventWatcher(ctrl *gomock.Controller) *MockServiceEventWatcher {
	mock := &MockServiceEventWatcher{ctrl: ctrl}
	mock.recorder = &MockServiceEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceEventWatcher) EXPECT() *MockServiceEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockServiceEventWatcher) AddEventHandler(arg0 context.Context, arg1 controller.ServiceEventHandler, arg2 ...predicate.Predicate) error {
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
func (mr *MockServiceEventWatcherMockRecorder) AddEventHandler(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockServiceEventWatcher)(nil).AddEventHandler), varargs...)
}

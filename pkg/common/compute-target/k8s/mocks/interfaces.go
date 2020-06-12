// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_mc_manager is a generated GoMock package.
package mock_mc_manager

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	mc_manager "github.com/solo-io/service-mesh-hub/pkg/common/compute-target/k8s"
	rest "k8s.io/client-go/rest"
	reflect "reflect"
	manager "sigs.k8s.io/controller-runtime/pkg/manager"
)

// MockAsyncManagerHandler is a mock of AsyncManagerHandler interface.
type MockAsyncManagerHandler struct {
	ctrl     *gomock.Controller
	recorder *MockAsyncManagerHandlerMockRecorder
}

// MockAsyncManagerHandlerMockRecorder is the mock recorder for MockAsyncManagerHandler.
type MockAsyncManagerHandlerMockRecorder struct {
	mock *MockAsyncManagerHandler
}

// NewMockAsyncManagerHandler creates a new mock instance.
func NewMockAsyncManagerHandler(ctrl *gomock.Controller) *MockAsyncManagerHandler {
	mock := &MockAsyncManagerHandler{ctrl: ctrl}
	mock.recorder = &MockAsyncManagerHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAsyncManagerHandler) EXPECT() *MockAsyncManagerHandlerMockRecorder {
	return m.recorder
}

// ClusterAdded mocks base method.
func (m *MockAsyncManagerHandler) ClusterAdded(ctx context.Context, mgr mc_manager.AsyncManager, clusterName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterAdded", ctx, mgr, clusterName)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClusterAdded indicates an expected call of ClusterAdded.
func (mr *MockAsyncManagerHandlerMockRecorder) ClusterAdded(ctx, mgr, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterAdded", reflect.TypeOf((*MockAsyncManagerHandler)(nil).ClusterAdded), ctx, mgr, clusterName)
}

// ClusterRemoved mocks base method.
func (m *MockAsyncManagerHandler) ClusterRemoved(cluster string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterRemoved", cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClusterRemoved indicates an expected call of ClusterRemoved.
func (mr *MockAsyncManagerHandlerMockRecorder) ClusterRemoved(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterRemoved", reflect.TypeOf((*MockAsyncManagerHandler)(nil).ClusterRemoved), cluster)
}

// MockAsyncManagerInformer is a mock of AsyncManagerInformer interface.
type MockAsyncManagerInformer struct {
	ctrl     *gomock.Controller
	recorder *MockAsyncManagerInformerMockRecorder
}

// MockAsyncManagerInformerMockRecorder is the mock recorder for MockAsyncManagerInformer.
type MockAsyncManagerInformerMockRecorder struct {
	mock *MockAsyncManagerInformer
}

// NewMockAsyncManagerInformer creates a new mock instance.
func NewMockAsyncManagerInformer(ctrl *gomock.Controller) *MockAsyncManagerInformer {
	mock := &MockAsyncManagerInformer{ctrl: ctrl}
	mock.recorder = &MockAsyncManagerInformerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAsyncManagerInformer) EXPECT() *MockAsyncManagerInformerMockRecorder {
	return m.recorder
}

// AddHandler mocks base method.
func (m *MockAsyncManagerInformer) AddHandler(informer mc_manager.AsyncManagerHandler, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddHandler", informer, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddHandler indicates an expected call of AddHandler.
func (mr *MockAsyncManagerInformerMockRecorder) AddHandler(informer, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHandler", reflect.TypeOf((*MockAsyncManagerInformer)(nil).AddHandler), informer, name)
}

// RemoveHandler mocks base method.
func (m *MockAsyncManagerInformer) RemoveHandler(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveHandler", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveHandler indicates an expected call of RemoveHandler.
func (mr *MockAsyncManagerInformerMockRecorder) RemoveHandler(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveHandler", reflect.TypeOf((*MockAsyncManagerInformer)(nil).RemoveHandler), name)
}

// MockAsyncManager is a mock of AsyncManager interface.
type MockAsyncManager struct {
	ctrl     *gomock.Controller
	recorder *MockAsyncManagerMockRecorder
}

// MockAsyncManagerMockRecorder is the mock recorder for MockAsyncManager.
type MockAsyncManagerMockRecorder struct {
	mock *MockAsyncManager
}

// NewMockAsyncManager creates a new mock instance.
func NewMockAsyncManager(ctrl *gomock.Controller) *MockAsyncManager {
	mock := &MockAsyncManager{ctrl: ctrl}
	mock.recorder = &MockAsyncManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAsyncManager) EXPECT() *MockAsyncManagerMockRecorder {
	return m.recorder
}

// Manager mocks base method.
func (m *MockAsyncManager) Manager() manager.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Manager")
	ret0, _ := ret[0].(manager.Manager)
	return ret0
}

// Manager indicates an expected call of Manager.
func (mr *MockAsyncManagerMockRecorder) Manager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Manager", reflect.TypeOf((*MockAsyncManager)(nil).Manager))
}

// Context mocks base method.
func (m *MockAsyncManager) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAsyncManagerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAsyncManager)(nil).Context))
}

// Error mocks base method.
func (m *MockAsyncManager) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockAsyncManagerMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockAsyncManager)(nil).Error))
}

// GotError mocks base method.
func (m *MockAsyncManager) GotError() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GotError")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// GotError indicates an expected call of GotError.
func (mr *MockAsyncManagerMockRecorder) GotError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GotError", reflect.TypeOf((*MockAsyncManager)(nil).GotError))
}

// Start mocks base method.
func (m *MockAsyncManager) Start(opts ...mc_manager.AsyncManagerStartOptionsFunc) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Start", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockAsyncManagerMockRecorder) Start(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockAsyncManager)(nil).Start), opts...)
}

// Stop mocks base method.
func (m *MockAsyncManager) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockAsyncManagerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockAsyncManager)(nil).Stop))
}

// MockAsyncManagerFactory is a mock of AsyncManagerFactory interface.
type MockAsyncManagerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockAsyncManagerFactoryMockRecorder
}

// MockAsyncManagerFactoryMockRecorder is the mock recorder for MockAsyncManagerFactory.
type MockAsyncManagerFactoryMockRecorder struct {
	mock *MockAsyncManagerFactory
}

// NewMockAsyncManagerFactory creates a new mock instance.
func NewMockAsyncManagerFactory(ctrl *gomock.Controller) *MockAsyncManagerFactory {
	mock := &MockAsyncManagerFactory{ctrl: ctrl}
	mock.recorder = &MockAsyncManagerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAsyncManagerFactory) EXPECT() *MockAsyncManagerFactoryMockRecorder {
	return m.recorder
}

// New mocks base method.
func (m *MockAsyncManagerFactory) New(parentCtx context.Context, cfg *rest.Config, opts mc_manager.AsyncManagerOptions) (mc_manager.AsyncManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", parentCtx, cfg, opts)
	ret0, _ := ret[0].(mc_manager.AsyncManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New.
func (mr *MockAsyncManagerFactoryMockRecorder) New(parentCtx, cfg, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockAsyncManagerFactory)(nil).New), parentCtx, cfg, opts)
}

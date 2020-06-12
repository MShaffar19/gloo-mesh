// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_meshes is a generated GoMock package.
package mock_meshes

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	v1alpha10 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
	dns "github.com/solo-io/service-mesh-hub/pkg/mesh-networking/federation/dns"
	reflect "reflect"
)

// MockMeshFederationClient is a mock of MeshFederationClient interface.
type MockMeshFederationClient struct {
	ctrl     *gomock.Controller
	recorder *MockMeshFederationClientMockRecorder
}

// MockMeshFederationClientMockRecorder is the mock recorder for MockMeshFederationClient.
type MockMeshFederationClientMockRecorder struct {
	mock *MockMeshFederationClient
}

// NewMockMeshFederationClient creates a new mock instance.
func NewMockMeshFederationClient(ctrl *gomock.Controller) *MockMeshFederationClient {
	mock := &MockMeshFederationClient{ctrl: ctrl}
	mock.recorder = &MockMeshFederationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshFederationClient) EXPECT() *MockMeshFederationClientMockRecorder {
	return m.recorder
}

// FederateServiceSide mocks base method.
func (m *MockMeshFederationClient) FederateServiceSide(ctx context.Context, installationNamespace string, virtualMesh *v1alpha10.VirtualMesh, meshService *v1alpha1.MeshService) (dns.ExternalAccessPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FederateServiceSide", ctx, installationNamespace, virtualMesh, meshService)
	ret0, _ := ret[0].(dns.ExternalAccessPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FederateServiceSide indicates an expected call of FederateServiceSide.
func (mr *MockMeshFederationClientMockRecorder) FederateServiceSide(ctx, installationNamespace, virtualMesh, meshService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FederateServiceSide", reflect.TypeOf((*MockMeshFederationClient)(nil).FederateServiceSide), ctx, installationNamespace, virtualMesh, meshService)
}

// FederateClientSide mocks base method.
func (m *MockMeshFederationClient) FederateClientSide(ctx context.Context, installationNamespace string, eap dns.ExternalAccessPoint, meshService *v1alpha1.MeshService, meshWorkload *v1alpha1.MeshWorkload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FederateClientSide", ctx, installationNamespace, eap, meshService, meshWorkload)
	ret0, _ := ret[0].(error)
	return ret0
}

// FederateClientSide indicates an expected call of FederateClientSide.
func (mr *MockMeshFederationClientMockRecorder) FederateClientSide(ctx, installationNamespace, eap, meshService, meshWorkload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FederateClientSide", reflect.TypeOf((*MockMeshFederationClient)(nil).FederateClientSide), ctx, installationNamespace, eap, meshService, meshWorkload)
}

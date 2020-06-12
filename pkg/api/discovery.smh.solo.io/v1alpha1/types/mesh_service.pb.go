// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/mesh_service.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
	types1 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/types"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//*
//The MeshService is an abstraction for a service which we have discovered to be part of a
//given mesh. The Mesh object has references to the MeshServices which belong to it.
type MeshServiceSpec struct {
	// Metadata about the kube-native service backing this MeshService.
	KubeService *MeshServiceSpec_KubeService `protobuf:"bytes,1,opt,name=kube_service,json=kubeService,proto3" json:"kube_service,omitempty"`
	// The mesh with which this service is associated.
	Mesh *types.ResourceRef `protobuf:"bytes,2,opt,name=mesh,proto3" json:"mesh,omitempty"`
	// Subsets for routing, based on labels.
	Subsets map[string]*MeshServiceSpec_Subset `protobuf:"bytes,3,rep,name=subsets,proto3" json:"subsets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Metadata about the decisions that Service Mesh Hub has made about what workloads this service is federated to.
	Federation           *MeshServiceSpec_Federation `protobuf:"bytes,4,opt,name=federation,proto3" json:"federation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MeshServiceSpec) Reset()         { *m = MeshServiceSpec{} }
func (m *MeshServiceSpec) String() string { return proto.CompactTextString(m) }
func (*MeshServiceSpec) ProtoMessage()    {}
func (*MeshServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d29223946a40ddd2, []int{0}
}
func (m *MeshServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceSpec.Unmarshal(m, b)
}
func (m *MeshServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceSpec.Marshal(b, m, deterministic)
}
func (m *MeshServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceSpec.Merge(m, src)
}
func (m *MeshServiceSpec) XXX_Size() int {
	return xxx_messageInfo_MeshServiceSpec.Size(m)
}
func (m *MeshServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceSpec proto.InternalMessageInfo

func (m *MeshServiceSpec) GetKubeService() *MeshServiceSpec_KubeService {
	if m != nil {
		return m.KubeService
	}
	return nil
}

func (m *MeshServiceSpec) GetMesh() *types.ResourceRef {
	if m != nil {
		return m.Mesh
	}
	return nil
}

func (m *MeshServiceSpec) GetSubsets() map[string]*MeshServiceSpec_Subset {
	if m != nil {
		return m.Subsets
	}
	return nil
}

func (m *MeshServiceSpec) GetFederation() *MeshServiceSpec_Federation {
	if m != nil {
		return m.Federation
	}
	return nil
}

type MeshServiceSpec_Federation struct {
	//*
	//For any workload that this service has federated to (i.e., any MeshWorkload whose ref appears in `federated_to_workloads`),
	//a client in that workload will be able to reach this service at this DNS name. This includes workloads on clusters other than
	//the one hosting this service.
	MulticlusterDnsName string `protobuf:"bytes,1,opt,name=multicluster_dns_name,json=multiclusterDnsName,proto3" json:"multicluster_dns_name,omitempty"`
	// The list of MeshWorkloads which are able to resolve this service's `multicluster_dns_name`.
	FederatedToWorkloads []*types.ResourceRef `protobuf:"bytes,2,rep,name=federated_to_workloads,json=federatedToWorkloads,proto3" json:"federated_to_workloads,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MeshServiceSpec_Federation) Reset()         { *m = MeshServiceSpec_Federation{} }
func (m *MeshServiceSpec_Federation) String() string { return proto.CompactTextString(m) }
func (*MeshServiceSpec_Federation) ProtoMessage()    {}
func (*MeshServiceSpec_Federation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d29223946a40ddd2, []int{0, 0}
}
func (m *MeshServiceSpec_Federation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceSpec_Federation.Unmarshal(m, b)
}
func (m *MeshServiceSpec_Federation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceSpec_Federation.Marshal(b, m, deterministic)
}
func (m *MeshServiceSpec_Federation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceSpec_Federation.Merge(m, src)
}
func (m *MeshServiceSpec_Federation) XXX_Size() int {
	return xxx_messageInfo_MeshServiceSpec_Federation.Size(m)
}
func (m *MeshServiceSpec_Federation) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceSpec_Federation.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceSpec_Federation proto.InternalMessageInfo

func (m *MeshServiceSpec_Federation) GetMulticlusterDnsName() string {
	if m != nil {
		return m.MulticlusterDnsName
	}
	return ""
}

func (m *MeshServiceSpec_Federation) GetFederatedToWorkloads() []*types.ResourceRef {
	if m != nil {
		return m.FederatedToWorkloads
	}
	return nil
}

type MeshServiceSpec_KubeService struct {
	// A reference to the kube-native service that this MeshService represents.
	Ref *types.ResourceRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// Selectors for the set of pods targeted by the k8s Service.
	WorkloadSelectorLabels map[string]string `protobuf:"bytes,2,rep,name=workload_selector_labels,json=workloadSelectorLabels,proto3" json:"workload_selector_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Labels on the underlying k8s Service itself.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The ports exposed by the underlying service.
	Ports                []*MeshServiceSpec_KubeService_KubeServicePort `protobuf:"bytes,4,rep,name=ports,proto3" json:"ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *MeshServiceSpec_KubeService) Reset()         { *m = MeshServiceSpec_KubeService{} }
func (m *MeshServiceSpec_KubeService) String() string { return proto.CompactTextString(m) }
func (*MeshServiceSpec_KubeService) ProtoMessage()    {}
func (*MeshServiceSpec_KubeService) Descriptor() ([]byte, []int) {
	return fileDescriptor_d29223946a40ddd2, []int{0, 1}
}
func (m *MeshServiceSpec_KubeService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceSpec_KubeService.Unmarshal(m, b)
}
func (m *MeshServiceSpec_KubeService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceSpec_KubeService.Marshal(b, m, deterministic)
}
func (m *MeshServiceSpec_KubeService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceSpec_KubeService.Merge(m, src)
}
func (m *MeshServiceSpec_KubeService) XXX_Size() int {
	return xxx_messageInfo_MeshServiceSpec_KubeService.Size(m)
}
func (m *MeshServiceSpec_KubeService) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceSpec_KubeService.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceSpec_KubeService proto.InternalMessageInfo

func (m *MeshServiceSpec_KubeService) GetRef() *types.ResourceRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *MeshServiceSpec_KubeService) GetWorkloadSelectorLabels() map[string]string {
	if m != nil {
		return m.WorkloadSelectorLabels
	}
	return nil
}

func (m *MeshServiceSpec_KubeService) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MeshServiceSpec_KubeService) GetPorts() []*MeshServiceSpec_KubeService_KubeServicePort {
	if m != nil {
		return m.Ports
	}
	return nil
}

type MeshServiceSpec_KubeService_KubeServicePort struct {
	// external-facing port for this service (i.e., NOT the service's target port on the backing pods)
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Protocol             string   `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshServiceSpec_KubeService_KubeServicePort) Reset() {
	*m = MeshServiceSpec_KubeService_KubeServicePort{}
}
func (m *MeshServiceSpec_KubeService_KubeServicePort) String() string {
	return proto.CompactTextString(m)
}
func (*MeshServiceSpec_KubeService_KubeServicePort) ProtoMessage() {}
func (*MeshServiceSpec_KubeService_KubeServicePort) Descriptor() ([]byte, []int) {
	return fileDescriptor_d29223946a40ddd2, []int{0, 1, 0}
}
func (m *MeshServiceSpec_KubeService_KubeServicePort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceSpec_KubeService_KubeServicePort.Unmarshal(m, b)
}
func (m *MeshServiceSpec_KubeService_KubeServicePort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceSpec_KubeService_KubeServicePort.Marshal(b, m, deterministic)
}
func (m *MeshServiceSpec_KubeService_KubeServicePort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceSpec_KubeService_KubeServicePort.Merge(m, src)
}
func (m *MeshServiceSpec_KubeService_KubeServicePort) XXX_Size() int {
	return xxx_messageInfo_MeshServiceSpec_KubeService_KubeServicePort.Size(m)
}
func (m *MeshServiceSpec_KubeService_KubeServicePort) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceSpec_KubeService_KubeServicePort.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceSpec_KubeService_KubeServicePort proto.InternalMessageInfo

func (m *MeshServiceSpec_KubeService_KubeServicePort) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *MeshServiceSpec_KubeService_KubeServicePort) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MeshServiceSpec_KubeService_KubeServicePort) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type MeshServiceSpec_Subset struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshServiceSpec_Subset) Reset()         { *m = MeshServiceSpec_Subset{} }
func (m *MeshServiceSpec_Subset) String() string { return proto.CompactTextString(m) }
func (*MeshServiceSpec_Subset) ProtoMessage()    {}
func (*MeshServiceSpec_Subset) Descriptor() ([]byte, []int) {
	return fileDescriptor_d29223946a40ddd2, []int{0, 2}
}
func (m *MeshServiceSpec_Subset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceSpec_Subset.Unmarshal(m, b)
}
func (m *MeshServiceSpec_Subset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceSpec_Subset.Marshal(b, m, deterministic)
}
func (m *MeshServiceSpec_Subset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceSpec_Subset.Merge(m, src)
}
func (m *MeshServiceSpec_Subset) XXX_Size() int {
	return xxx_messageInfo_MeshServiceSpec_Subset.Size(m)
}
func (m *MeshServiceSpec_Subset) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceSpec_Subset.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceSpec_Subset proto.InternalMessageInfo

func (m *MeshServiceSpec_Subset) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type MeshServiceStatus struct {
	// The status of federation artifacts being written to remote clusters as a result of the federation metadata on this object's Spec.
	FederationStatus         *types.Status                               `protobuf:"bytes,1,opt,name=federation_status,json=federationStatus,proto3" json:"federation_status,omitempty"`
	ValidatedTrafficPolicies []*MeshServiceStatus_ValidatedTrafficPolicy `protobuf:"bytes,2,rep,name=validated_traffic_policies,json=validatedTrafficPolicies,proto3" json:"validated_traffic_policies,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                                    `json:"-"`
	XXX_unrecognized         []byte                                      `json:"-"`
	XXX_sizecache            int32                                       `json:"-"`
}

func (m *MeshServiceStatus) Reset()         { *m = MeshServiceStatus{} }
func (m *MeshServiceStatus) String() string { return proto.CompactTextString(m) }
func (*MeshServiceStatus) ProtoMessage()    {}
func (*MeshServiceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_d29223946a40ddd2, []int{1}
}
func (m *MeshServiceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceStatus.Unmarshal(m, b)
}
func (m *MeshServiceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceStatus.Marshal(b, m, deterministic)
}
func (m *MeshServiceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceStatus.Merge(m, src)
}
func (m *MeshServiceStatus) XXX_Size() int {
	return xxx_messageInfo_MeshServiceStatus.Size(m)
}
func (m *MeshServiceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceStatus proto.InternalMessageInfo

func (m *MeshServiceStatus) GetFederationStatus() *types.Status {
	if m != nil {
		return m.FederationStatus
	}
	return nil
}

func (m *MeshServiceStatus) GetValidatedTrafficPolicies() []*MeshServiceStatus_ValidatedTrafficPolicy {
	if m != nil {
		return m.ValidatedTrafficPolicies
	}
	return nil
}

type MeshServiceStatus_ValidatedTrafficPolicy struct {
	Ref                  *types.ResourceRef        `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	TrafficPolicySpec    *types1.TrafficPolicySpec `protobuf:"bytes,2,opt,name=traffic_policy_spec,json=trafficPolicySpec,proto3" json:"traffic_policy_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MeshServiceStatus_ValidatedTrafficPolicy) Reset() {
	*m = MeshServiceStatus_ValidatedTrafficPolicy{}
}
func (m *MeshServiceStatus_ValidatedTrafficPolicy) String() string { return proto.CompactTextString(m) }
func (*MeshServiceStatus_ValidatedTrafficPolicy) ProtoMessage()    {}
func (*MeshServiceStatus_ValidatedTrafficPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_d29223946a40ddd2, []int{1, 0}
}
func (m *MeshServiceStatus_ValidatedTrafficPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshServiceStatus_ValidatedTrafficPolicy.Unmarshal(m, b)
}
func (m *MeshServiceStatus_ValidatedTrafficPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshServiceStatus_ValidatedTrafficPolicy.Marshal(b, m, deterministic)
}
func (m *MeshServiceStatus_ValidatedTrafficPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshServiceStatus_ValidatedTrafficPolicy.Merge(m, src)
}
func (m *MeshServiceStatus_ValidatedTrafficPolicy) XXX_Size() int {
	return xxx_messageInfo_MeshServiceStatus_ValidatedTrafficPolicy.Size(m)
}
func (m *MeshServiceStatus_ValidatedTrafficPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshServiceStatus_ValidatedTrafficPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_MeshServiceStatus_ValidatedTrafficPolicy proto.InternalMessageInfo

func (m *MeshServiceStatus_ValidatedTrafficPolicy) GetRef() *types.ResourceRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *MeshServiceStatus_ValidatedTrafficPolicy) GetTrafficPolicySpec() *types1.TrafficPolicySpec {
	if m != nil {
		return m.TrafficPolicySpec
	}
	return nil
}

func init() {
	proto.RegisterType((*MeshServiceSpec)(nil), "discovery.smh.solo.io.MeshServiceSpec")
	proto.RegisterMapType((map[string]*MeshServiceSpec_Subset)(nil), "discovery.smh.solo.io.MeshServiceSpec.SubsetsEntry")
	proto.RegisterType((*MeshServiceSpec_Federation)(nil), "discovery.smh.solo.io.MeshServiceSpec.Federation")
	proto.RegisterType((*MeshServiceSpec_KubeService)(nil), "discovery.smh.solo.io.MeshServiceSpec.KubeService")
	proto.RegisterMapType((map[string]string)(nil), "discovery.smh.solo.io.MeshServiceSpec.KubeService.LabelsEntry")
	proto.RegisterMapType((map[string]string)(nil), "discovery.smh.solo.io.MeshServiceSpec.KubeService.WorkloadSelectorLabelsEntry")
	proto.RegisterType((*MeshServiceSpec_KubeService_KubeServicePort)(nil), "discovery.smh.solo.io.MeshServiceSpec.KubeService.KubeServicePort")
	proto.RegisterType((*MeshServiceSpec_Subset)(nil), "discovery.smh.solo.io.MeshServiceSpec.Subset")
	proto.RegisterType((*MeshServiceStatus)(nil), "discovery.smh.solo.io.MeshServiceStatus")
	proto.RegisterType((*MeshServiceStatus_ValidatedTrafficPolicy)(nil), "discovery.smh.solo.io.MeshServiceStatus.ValidatedTrafficPolicy")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/mesh_service.proto", fileDescriptor_d29223946a40ddd2)
}

var fileDescriptor_d29223946a40ddd2 = []byte{
	// 725 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xc7, 0xe5, 0x24, 0xcd, 0xbd, 0x3d, 0xe9, 0x55, 0xdb, 0x69, 0x1b, 0x59, 0xbe, 0xba, 0x57,
	0x51, 0x57, 0x45, 0x22, 0xb6, 0x92, 0x6e, 0x00, 0x09, 0x90, 0x0a, 0x45, 0x42, 0xd0, 0xaa, 0x38,
	0x6d, 0xf9, 0xd8, 0x58, 0x8e, 0x73, 0x92, 0x8c, 0xe2, 0x64, 0xac, 0x99, 0x71, 0xaa, 0x2c, 0x90,
	0x58, 0xf2, 0x02, 0xbc, 0x01, 0x0b, 0x5e, 0x82, 0xf7, 0x60, 0xcd, 0x93, 0x20, 0xcf, 0xd8, 0x89,
	0x53, 0x2c, 0x9a, 0x76, 0x37, 0x1f, 0xe7, 0xff, 0x3b, 0x67, 0xe6, 0x9c, 0x33, 0x03, 0x27, 0x03,
	0x2a, 0x87, 0x71, 0xd7, 0x0e, 0xd8, 0xd8, 0x11, 0x2c, 0x64, 0x4d, 0xca, 0x1c, 0x81, 0x7c, 0x4a,
	0x03, 0x6c, 0x8e, 0x51, 0x0c, 0x9b, 0xc3, 0xb8, 0xeb, 0xf8, 0x11, 0x75, 0x7a, 0x54, 0x04, 0x6c,
	0x8a, 0x7c, 0xe6, 0x4c, 0x5b, 0x7e, 0x18, 0x0d, 0xfd, 0x96, 0x93, 0xec, 0x7b, 0xa9, 0xb1, 0x1d,
	0x71, 0x26, 0x19, 0xd9, 0x9b, 0x9b, 0xd9, 0x62, 0x3c, 0xb4, 0x13, 0xa2, 0x4d, 0x99, 0x75, 0xbf,
	0x10, 0x19, 0x30, 0x8e, 0x0b, 0x1a, 0xc7, 0xbe, 0x86, 0x58, 0xce, 0x0a, 0xd6, 0x42, 0xfa, 0x32,
	0x16, 0xa9, 0xe0, 0x71, 0xa1, 0x60, 0x82, 0xf2, 0x8a, 0xf1, 0x11, 0x9d, 0x0c, 0x16, 0x32, 0xc9,
	0xfd, 0x7e, 0x9f, 0x06, 0x5e, 0xc4, 0x42, 0x1a, 0xcc, 0x52, 0xf9, 0xee, 0x80, 0x0d, 0x98, 0x1a,
	0x3a, 0xc9, 0x48, 0xaf, 0xee, 0x7f, 0x5f, 0x87, 0xcd, 0x13, 0x14, 0xc3, 0x8e, 0x66, 0x77, 0x22,
	0x0c, 0xc8, 0x05, 0x6c, 0x8c, 0xe2, 0x2e, 0x66, 0x87, 0x36, 0x8d, 0x86, 0x71, 0x50, 0x6b, 0xb7,
	0xed, 0xc2, 0x53, 0xdb, 0xd7, 0xd4, 0xf6, 0xab, 0xb8, 0x8b, 0xe9, 0xdc, 0xad, 0x8d, 0x16, 0x13,
	0xd2, 0x82, 0x4a, 0x12, 0xb9, 0x59, 0x52, 0xb8, 0xff, 0xec, 0xe4, 0xa8, 0x4b, 0x24, 0x17, 0x05,
	0x8b, 0x79, 0x80, 0x2e, 0xf6, 0x5d, 0x65, 0x4a, 0x4e, 0xe0, 0x2f, 0x11, 0x77, 0x05, 0x4a, 0x61,
	0x96, 0x1b, 0xe5, 0x83, 0x5a, 0xfb, 0x70, 0xc5, 0x20, 0x3a, 0x5a, 0x75, 0x3c, 0x91, 0x7c, 0xe6,
	0x66, 0x0c, 0xf2, 0x06, 0xa0, 0x8f, 0x3d, 0xe4, 0xbe, 0xa4, 0x6c, 0x62, 0x56, 0x54, 0x1c, 0xad,
	0x15, 0x89, 0x2f, 0xe6, 0x42, 0x37, 0x07, 0xb1, 0xbe, 0x18, 0x00, 0x8b, 0x2d, 0xd2, 0x86, 0xbd,
	0x71, 0x1c, 0x4a, 0x1a, 0x84, 0xb1, 0x90, 0xc8, 0xbd, 0xde, 0x44, 0x78, 0x13, 0x7f, 0xac, 0xef,
	0x70, 0xdd, 0xdd, 0xc9, 0x6f, 0x3e, 0x9f, 0x88, 0x53, 0x7f, 0x8c, 0xa4, 0x03, 0xf5, 0x14, 0x88,
	0x3d, 0x4f, 0x32, 0x2f, 0x49, 0x67, 0xc8, 0xfc, 0x9e, 0x30, 0x4b, 0xea, 0xcc, 0x37, 0xdc, 0xd4,
	0xee, 0x5c, 0x7c, 0xce, 0xde, 0x66, 0x52, 0xeb, 0x47, 0x05, 0x6a, 0xb9, 0x4c, 0x10, 0x07, 0xca,
	0x1c, 0xfb, 0x69, 0x2a, 0x6f, 0x20, 0x26, 0x96, 0xe4, 0xb3, 0x01, 0x66, 0x16, 0x89, 0x27, 0x30,
	0xc4, 0x40, 0x32, 0xee, 0x85, 0x7e, 0x17, 0xc3, 0x2c, 0xb0, 0xd3, 0xdb, 0x57, 0x84, 0x9d, 0x45,
	0xd8, 0x49, 0x89, 0xaf, 0x15, 0x50, 0xe7, 0xa9, 0x7e, 0x55, 0xb8, 0x49, 0x2e, 0xa1, 0x9a, 0xfa,
	0xd5, 0x45, 0xf0, 0xe4, 0x0e, 0x7e, 0xf3, 0x7e, 0x52, 0x1a, 0x79, 0x07, 0x6b, 0x11, 0xe3, 0x52,
	0x98, 0x15, 0x85, 0x3d, 0xba, 0x03, 0x36, 0x37, 0x3e, 0x63, 0x5c, 0xba, 0x1a, 0x68, 0x5d, 0xc0,
	0xe6, 0xb5, 0x1d, 0x42, 0xa0, 0x92, 0xec, 0xa9, 0x0c, 0xfc, 0xe3, 0xaa, 0x71, 0xb2, 0xa6, 0x8a,
	0xa3, 0xa4, 0x8a, 0x43, 0x8d, 0x89, 0x05, 0x7f, 0xab, 0xce, 0x0c, 0x58, 0x68, 0x96, 0xd5, 0xfa,
	0x7c, 0x6e, 0xbd, 0x84, 0x7f, 0xff, 0x70, 0x7f, 0x64, 0x0b, 0xca, 0x23, 0x9c, 0xa5, 0xa5, 0x96,
	0x0c, 0xc9, 0x2e, 0xac, 0x4d, 0xfd, 0x30, 0xce, 0x3c, 0xe8, 0xc9, 0xa3, 0xd2, 0x03, 0xc3, 0x7a,
	0x08, 0xb5, 0xbb, 0x4a, 0x1b, 0x50, 0xd5, 0xed, 0x45, 0xea, 0x50, 0x55, 0xcb, 0xc2, 0x34, 0x1a,
	0xe5, 0x83, 0x75, 0x37, 0x9d, 0x59, 0x14, 0x36, 0xf2, 0x0d, 0x58, 0x40, 0x7f, 0x96, 0xa7, 0xd7,
	0xda, 0xcd, 0x5b, 0xb5, 0x75, 0x2e, 0x98, 0xfd, 0x4f, 0x65, 0xd8, 0xce, 0x5b, 0xa9, 0x07, 0x93,
	0x1c, 0xc3, 0xf6, 0xa2, 0x47, 0x3d, 0xfd, 0x8a, 0xa6, 0xb5, 0x6f, 0xfe, 0x5e, 0xfb, 0x5a, 0xe4,
	0x6e, 0x2d, 0x24, 0x29, 0xe6, 0x23, 0x58, 0x53, 0x3f, 0xa4, 0x3d, 0xdd, 0x99, 0xf9, 0x47, 0x95,
	0x62, 0xd6, 0x04, 0x4f, 0x57, 0x08, 0x5d, 0xbf, 0xe2, 0x97, 0x19, 0xea, 0x5c, 0x93, 0xce, 0xd4,
	0xeb, 0xec, 0x9a, 0xd3, 0xa2, 0x75, 0x8a, 0xc2, 0xfa, 0x6a, 0x40, 0xbd, 0x58, 0x74, 0xfb, 0x76,
	0x7e, 0x0f, 0x3b, 0xcb, 0xbf, 0x82, 0x27, 0x22, 0x0c, 0xd2, 0xeb, 0xbf, 0x67, 0x2f, 0x7e, 0x91,
	0x25, 0xcc, 0x92, 0xd3, 0x24, 0x03, 0xee, 0xb6, 0xbc, 0xbe, 0x74, 0x74, 0xf9, 0xed, 0xe7, 0xff,
	0xc6, 0x87, 0xb3, 0x55, 0xbe, 0xd8, 0x68, 0x34, 0x58, 0xfe, 0x66, 0xf3, 0xde, 0x72, 0xff, 0xd7,
	0x2c, 0x42, 0xd1, 0xad, 0xaa, 0xba, 0x3f, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x44, 0x5f, 0x97,
	0x8b, 0xbd, 0x07, 0x00, 0x00,
}

func (this *MeshServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceSpec)
	if !ok {
		that2, ok := that.(MeshServiceSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.KubeService.Equal(that1.KubeService) {
		return false
	}
	if !this.Mesh.Equal(that1.Mesh) {
		return false
	}
	if len(this.Subsets) != len(that1.Subsets) {
		return false
	}
	for i := range this.Subsets {
		if !this.Subsets[i].Equal(that1.Subsets[i]) {
			return false
		}
	}
	if !this.Federation.Equal(that1.Federation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshServiceSpec_Federation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceSpec_Federation)
	if !ok {
		that2, ok := that.(MeshServiceSpec_Federation)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MulticlusterDnsName != that1.MulticlusterDnsName {
		return false
	}
	if len(this.FederatedToWorkloads) != len(that1.FederatedToWorkloads) {
		return false
	}
	for i := range this.FederatedToWorkloads {
		if !this.FederatedToWorkloads[i].Equal(that1.FederatedToWorkloads[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshServiceSpec_KubeService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceSpec_KubeService)
	if !ok {
		that2, ok := that.(MeshServiceSpec_KubeService)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Ref.Equal(that1.Ref) {
		return false
	}
	if len(this.WorkloadSelectorLabels) != len(that1.WorkloadSelectorLabels) {
		return false
	}
	for i := range this.WorkloadSelectorLabels {
		if this.WorkloadSelectorLabels[i] != that1.WorkloadSelectorLabels[i] {
			return false
		}
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if len(this.Ports) != len(that1.Ports) {
		return false
	}
	for i := range this.Ports {
		if !this.Ports[i].Equal(that1.Ports[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshServiceSpec_KubeService_KubeServicePort) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceSpec_KubeService_KubeServicePort)
	if !ok {
		that2, ok := that.(MeshServiceSpec_KubeService_KubeServicePort)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Protocol != that1.Protocol {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshServiceSpec_Subset) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceSpec_Subset)
	if !ok {
		that2, ok := that.(MeshServiceSpec_Subset)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Values) != len(that1.Values) {
		return false
	}
	for i := range this.Values {
		if this.Values[i] != that1.Values[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshServiceStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceStatus)
	if !ok {
		that2, ok := that.(MeshServiceStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.FederationStatus.Equal(that1.FederationStatus) {
		return false
	}
	if len(this.ValidatedTrafficPolicies) != len(that1.ValidatedTrafficPolicies) {
		return false
	}
	for i := range this.ValidatedTrafficPolicies {
		if !this.ValidatedTrafficPolicies[i].Equal(that1.ValidatedTrafficPolicies[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshServiceStatus_ValidatedTrafficPolicy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshServiceStatus_ValidatedTrafficPolicy)
	if !ok {
		that2, ok := that.(MeshServiceStatus_ValidatedTrafficPolicy)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Ref.Equal(that1.Ref) {
		return false
	}
	if !this.TrafficPolicySpec.Equal(that1.TrafficPolicySpec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

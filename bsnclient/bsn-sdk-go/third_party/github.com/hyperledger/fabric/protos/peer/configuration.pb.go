/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/configuration.proto

package peer // import "bsn-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AnchorPeers simply represents list of anchor peers which is used in ConfigurationItem
type AnchorPeers struct {
	AnchorPeers          []*AnchorPeer `protobuf:"bytes,1,rep,name=anchor_peers,json=anchorPeers,proto3" json:"anchor_peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AnchorPeers) Reset()         { *m = AnchorPeers{} }
func (m *AnchorPeers) String() string { return proto.CompactTextString(m) }
func (*AnchorPeers) ProtoMessage()    {}
func (*AnchorPeers) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_d9ec63ae33c182ef, []int{0}
}
func (m *AnchorPeers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnchorPeers.Unmarshal(m, b)
}
func (m *AnchorPeers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnchorPeers.Marshal(b, m, deterministic)
}
func (dst *AnchorPeers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnchorPeers.Merge(dst, src)
}
func (m *AnchorPeers) XXX_Size() int {
	return xxx_messageInfo_AnchorPeers.Size(m)
}
func (m *AnchorPeers) XXX_DiscardUnknown() {
	xxx_messageInfo_AnchorPeers.DiscardUnknown(m)
}

var xxx_messageInfo_AnchorPeers proto.InternalMessageInfo

func (m *AnchorPeers) GetAnchorPeers() []*AnchorPeer {
	if m != nil {
		return m.AnchorPeers
	}
	return nil
}

// AnchorPeer message structure which provides information about anchor peer, it includes host name,
// port number and peer certificate.
type AnchorPeer struct {
	// DNS host name of the anchor peer
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// The port number
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnchorPeer) Reset()         { *m = AnchorPeer{} }
func (m *AnchorPeer) String() string { return proto.CompactTextString(m) }
func (*AnchorPeer) ProtoMessage()    {}
func (*AnchorPeer) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_d9ec63ae33c182ef, []int{1}
}
func (m *AnchorPeer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnchorPeer.Unmarshal(m, b)
}
func (m *AnchorPeer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnchorPeer.Marshal(b, m, deterministic)
}
func (dst *AnchorPeer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnchorPeer.Merge(dst, src)
}
func (m *AnchorPeer) XXX_Size() int {
	return xxx_messageInfo_AnchorPeer.Size(m)
}
func (m *AnchorPeer) XXX_DiscardUnknown() {
	xxx_messageInfo_AnchorPeer.DiscardUnknown(m)
}

var xxx_messageInfo_AnchorPeer proto.InternalMessageInfo

func (m *AnchorPeer) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *AnchorPeer) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// APIResource represents an API resource in the peer whose ACL
// is determined by the policy_ref field
type APIResource struct {
	PolicyRef            string   `protobuf:"bytes,1,opt,name=policy_ref,json=policyRef,proto3" json:"policy_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIResource) Reset()         { *m = APIResource{} }
func (m *APIResource) String() string { return proto.CompactTextString(m) }
func (*APIResource) ProtoMessage()    {}
func (*APIResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_d9ec63ae33c182ef, []int{2}
}
func (m *APIResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIResource.Unmarshal(m, b)
}
func (m *APIResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIResource.Marshal(b, m, deterministic)
}
func (dst *APIResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIResource.Merge(dst, src)
}
func (m *APIResource) XXX_Size() int {
	return xxx_messageInfo_APIResource.Size(m)
}
func (m *APIResource) XXX_DiscardUnknown() {
	xxx_messageInfo_APIResource.DiscardUnknown(m)
}

var xxx_messageInfo_APIResource proto.InternalMessageInfo

func (m *APIResource) GetPolicyRef() string {
	if m != nil {
		return m.PolicyRef
	}
	return ""
}

// ACLs provides mappings for resources in a channel. APIResource encapsulates
// reference to a policy used to determine ACL for the resource
type ACLs struct {
	Acls                 map[string]*APIResource `protobuf:"bytes,1,rep,name=acls,proto3" json:"acls,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ACLs) Reset()         { *m = ACLs{} }
func (m *ACLs) String() string { return proto.CompactTextString(m) }
func (*ACLs) ProtoMessage()    {}
func (*ACLs) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_d9ec63ae33c182ef, []int{3}
}
func (m *ACLs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACLs.Unmarshal(m, b)
}
func (m *ACLs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACLs.Marshal(b, m, deterministic)
}
func (dst *ACLs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACLs.Merge(dst, src)
}
func (m *ACLs) XXX_Size() int {
	return xxx_messageInfo_ACLs.Size(m)
}
func (m *ACLs) XXX_DiscardUnknown() {
	xxx_messageInfo_ACLs.DiscardUnknown(m)
}

var xxx_messageInfo_ACLs proto.InternalMessageInfo

func (m *ACLs) GetAcls() map[string]*APIResource {
	if m != nil {
		return m.Acls
	}
	return nil
}

func init() {
	proto.RegisterType((*AnchorPeers)(nil), "sdk.protos.AnchorPeers")
	proto.RegisterType((*AnchorPeer)(nil), "sdk.protos.AnchorPeer")
	proto.RegisterType((*APIResource)(nil), "sdk.protos.APIResource")
	proto.RegisterType((*ACLs)(nil), "sdk.protos.ACLs")
	proto.RegisterMapType((map[string]*APIResource)(nil), "sdk.protos.ACLs.AclsEntry")
}

func init() {
	proto.RegisterFile("peer/configuration.proto", fileDescriptor_configuration_d9ec63ae33c182ef)
}

var fileDescriptor_configuration_d9ec63ae33c182ef = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xdf, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0xe9, 0x7e, 0x7c, 0x61, 0xb7, 0xdf, 0x07, 0x89, 0x20, 0x45, 0x10, 0x46, 0x9f, 0x36,
	0x91, 0x14, 0xa6, 0x82, 0xf8, 0x56, 0xa7, 0x0f, 0xc2, 0xc0, 0x91, 0x47, 0x5f, 0x46, 0x16, 0x6f,
	0x7f, 0x60, 0x6d, 0xca, 0x4d, 0x2a, 0xf4, 0xcd, 0x3f, 0x5d, 0x9a, 0x6c, 0xab, 0x4f, 0x39, 0x39,
	0xf9, 0x9c, 0xcb, 0x21, 0x17, 0xa2, 0x06, 0x91, 0x12, 0xa5, 0xeb, 0xac, 0xcc, 0x5b, 0x92, 0xb6,
	0xd4, 0x35, 0x6f, 0x48, 0x5b, 0xcd, 0xfe, 0xb9, 0xc3, 0xc4, 0xcf, 0x10, 0xa6, 0xb5, 0x2a, 0x34,
	0x6d, 0x11, 0xc9, 0xb0, 0x7b, 0xf8, 0x2f, 0xdd, 0x75, 0xd7, 0x27, 0x4d, 0x14, 0xcc, 0xc7, 0x8b,
	0x70, 0xc5, 0x7c, 0xc8, 0xf0, 0x01, 0x15, 0xa1, 0x1c, 0x62, 0xf1, 0x1d, 0xc0, 0xf0, 0xc4, 0x18,
	0x4c, 0x0a, 0x6d, 0x6c, 0x14, 0xcc, 0x83, 0xc5, 0x4c, 0x38, 0xdd, 0x7b, 0x8d, 0x26, 0x1b, 0x8d,
	0xe6, 0xc1, 0x62, 0x2a, 0x9c, 0x8e, 0x6f, 0x20, 0x4c, 0xb7, 0xaf, 0x02, 0x8d, 0x6e, 0x49, 0x21,
	0xbb, 0x02, 0x68, 0x74, 0x55, 0xaa, 0x6e, 0x47, 0x98, 0x1d, 0xc2, 0x33, 0xef, 0x08, 0xcc, 0xe2,
	0x9f, 0x00, 0x26, 0xe9, 0x7a, 0x63, 0xd8, 0x35, 0x4c, 0xa4, 0xaa, 0x8e, 0xdd, 0x2e, 0x4e, 0xdd,
	0xd6, 0x1b, 0xc3, 0x53, 0x55, 0x99, 0x97, 0xda, 0x52, 0x27, 0x1c, 0x73, 0xb9, 0x81, 0xd9, 0xc9,
	0x62, 0x67, 0x30, 0xfe, 0xc4, 0xee, 0x30, 0xb9, 0x97, 0x6c, 0x09, 0xd3, 0x6f, 0x59, 0xb5, 0xe8,
	0x6a, 0x85, 0xab, 0xf3, 0xd3, 0xac, 0xa1, 0x96, 0xf0, 0xc4, 0xe3, 0xe8, 0x21, 0x78, 0x7a, 0x83,
	0x58, 0x53, 0xce, 0x8b, 0xae, 0x41, 0xaa, 0xf0, 0x23, 0x47, 0xe2, 0x99, 0xdc, 0x53, 0xa9, 0x8e,
	0xb9, 0xfe, 0xd3, 0xde, 0x97, 0x79, 0x69, 0x8b, 0x76, 0xcf, 0x95, 0xfe, 0x4a, 0xfe, 0xa0, 0x89,
	0x47, 0x13, 0x8f, 0x26, 0x3d, 0xba, 0xf7, 0x5b, 0xb8, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xae,
	0x0a, 0x1d, 0x41, 0xa8, 0x01, 0x00, 0x00,
}

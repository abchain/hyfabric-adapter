/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/query.proto

package peer // import "hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"

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

// ChaincodeQueryResponse returns information about each chaincode that pertains
// to a query in lscc.go, such as GetChaincodes (returns all chaincodes
// instantiated on a channel), and GetInstalledChaincodes (returns all chaincodes
// installed on a peer)
type ChaincodeQueryResponse struct {
	Chaincodes           []*ChaincodeInfo `protobuf:"bytes,1,rep,name=chaincodes,proto3" json:"chaincodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ChaincodeQueryResponse) Reset()         { *m = ChaincodeQueryResponse{} }
func (m *ChaincodeQueryResponse) String() string { return proto.CompactTextString(m) }
func (*ChaincodeQueryResponse) ProtoMessage()    {}
func (*ChaincodeQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_37e5d407118e73f5, []int{0}
}
func (m *ChaincodeQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeQueryResponse.Unmarshal(m, b)
}
func (m *ChaincodeQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeQueryResponse.Marshal(b, m, deterministic)
}
func (dst *ChaincodeQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeQueryResponse.Merge(dst, src)
}
func (m *ChaincodeQueryResponse) XXX_Size() int {
	return xxx_messageInfo_ChaincodeQueryResponse.Size(m)
}
func (m *ChaincodeQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeQueryResponse proto.InternalMessageInfo

func (m *ChaincodeQueryResponse) GetChaincodes() []*ChaincodeInfo {
	if m != nil {
		return m.Chaincodes
	}
	return nil
}

// ChaincodeInfo contains general information about an installed/instantiated
// chaincode
type ChaincodeInfo struct {
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// the path as specified by the install/instantiate transaction
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// the chaincode function upon instantiation and its arguments. This will be
	// blank if the query is returning information about installed chaincodes.
	Input string `protobuf:"bytes,4,opt,name=input,proto3" json:"input,omitempty"`
	// the name of the ESCC for this chaincode. This will be
	// blank if the query is returning information about installed chaincodes.
	Escc string `protobuf:"bytes,5,opt,name=escc,proto3" json:"escc,omitempty"`
	// the name of the VSCC for this chaincode. This will be
	// blank if the query is returning information about installed chaincodes.
	Vscc string `protobuf:"bytes,6,opt,name=vscc,proto3" json:"vscc,omitempty"`
	// the chaincode unique id.
	// computed as: H(
	//                H(name || version) ||
	//                H(CodePackage)
	//              )
	Id                   []byte   `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeInfo) Reset()         { *m = ChaincodeInfo{} }
func (m *ChaincodeInfo) String() string { return proto.CompactTextString(m) }
func (*ChaincodeInfo) ProtoMessage()    {}
func (*ChaincodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_37e5d407118e73f5, []int{1}
}
func (m *ChaincodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeInfo.Unmarshal(m, b)
}
func (m *ChaincodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeInfo.Marshal(b, m, deterministic)
}
func (dst *ChaincodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeInfo.Merge(dst, src)
}
func (m *ChaincodeInfo) XXX_Size() int {
	return xxx_messageInfo_ChaincodeInfo.Size(m)
}
func (m *ChaincodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeInfo proto.InternalMessageInfo

func (m *ChaincodeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ChaincodeInfo) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ChaincodeInfo) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *ChaincodeInfo) GetEscc() string {
	if m != nil {
		return m.Escc
	}
	return ""
}

func (m *ChaincodeInfo) GetVscc() string {
	if m != nil {
		return m.Vscc
	}
	return ""
}

func (m *ChaincodeInfo) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

// ChannelQueryResponse returns information about each channel that pertains
// to a query in lscc.go, such as GetChannels (returns all channels for a
// given peer)
type ChannelQueryResponse struct {
	Channels             []*ChannelInfo `protobuf:"bytes,1,rep,name=channels,proto3" json:"channels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ChannelQueryResponse) Reset()         { *m = ChannelQueryResponse{} }
func (m *ChannelQueryResponse) String() string { return proto.CompactTextString(m) }
func (*ChannelQueryResponse) ProtoMessage()    {}
func (*ChannelQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_37e5d407118e73f5, []int{2}
}
func (m *ChannelQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelQueryResponse.Unmarshal(m, b)
}
func (m *ChannelQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelQueryResponse.Marshal(b, m, deterministic)
}
func (dst *ChannelQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelQueryResponse.Merge(dst, src)
}
func (m *ChannelQueryResponse) XXX_Size() int {
	return xxx_messageInfo_ChannelQueryResponse.Size(m)
}
func (m *ChannelQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelQueryResponse proto.InternalMessageInfo

func (m *ChannelQueryResponse) GetChannels() []*ChannelInfo {
	if m != nil {
		return m.Channels
	}
	return nil
}

// ChannelInfo contains general information about channels
type ChannelInfo struct {
	ChannelId            string   `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelInfo) Reset()         { *m = ChannelInfo{} }
func (m *ChannelInfo) String() string { return proto.CompactTextString(m) }
func (*ChannelInfo) ProtoMessage()    {}
func (*ChannelInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_37e5d407118e73f5, []int{3}
}
func (m *ChannelInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelInfo.Unmarshal(m, b)
}
func (m *ChannelInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelInfo.Marshal(b, m, deterministic)
}
func (dst *ChannelInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelInfo.Merge(dst, src)
}
func (m *ChannelInfo) XXX_Size() int {
	return xxx_messageInfo_ChannelInfo.Size(m)
}
func (m *ChannelInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelInfo proto.InternalMessageInfo

func (m *ChannelInfo) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func init() {
	proto.RegisterType((*ChaincodeQueryResponse)(nil), "sdk.protos.ChaincodeQueryResponse")
	proto.RegisterType((*ChaincodeInfo)(nil), "sdk.protos.ChaincodeInfo")
	proto.RegisterType((*ChannelQueryResponse)(nil), "sdk.protos.ChannelQueryResponse")
	proto.RegisterType((*ChannelInfo)(nil), "sdk.protos.ChannelInfo")
}

func init() { proto.RegisterFile("peer/query.proto", fileDescriptor_query_37e5d407118e73f5) }

var fileDescriptor_query_37e5d407118e73f5 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0xe9, 0xfe, 0xba, 0x33, 0x15, 0x89, 0x53, 0x72, 0x23, 0x8c, 0x5e, 0x4d, 0x90, 0x16,
	0x14, 0x5f, 0xc0, 0x5d, 0xc8, 0xae, 0x86, 0xbd, 0xf4, 0x46, 0xba, 0xe4, 0x6c, 0x0d, 0x6c, 0x49,
	0x4c, 0xba, 0xc1, 0x9e, 0xc6, 0x57, 0x95, 0x93, 0xac, 0xa3, 0xbb, 0xea, 0x39, 0xbf, 0xef, 0x17,
	0xca, 0x97, 0xc0, 0x9d, 0x45, 0x74, 0xf9, 0xef, 0x1e, 0xdd, 0x31, 0xb3, 0xce, 0xd4, 0x86, 0x0d,
	0xc2, 0xc7, 0xa7, 0x4b, 0x78, 0x9c, 0x57, 0xa5, 0xd2, 0xc2, 0x48, 0xfc, 0xa2, 0xbc, 0x40, 0x6f,
	0x8d, 0xf6, 0xc8, 0xde, 0x01, 0x44, 0x93, 0x78, 0x9e, 0x4c, 0xbb, 0xb3, 0xf1, 0xeb, 0x43, 0x3c,
	0xed, 0xb3, 0xf3, 0x99, 0x85, 0x5e, 0x9b, 0xa2, 0x25, 0xa6, 0x7f, 0x09, 0xdc, 0x5c, 0xa4, 0x8c,
	0x41, 0x4f, 0x97, 0x3b, 0xe4, 0xc9, 0x34, 0x99, 0x8d, 0x8a, 0x30, 0x33, 0x0e, 0xc3, 0x03, 0x3a,
	0xaf, 0x8c, 0xe6, 0x9d, 0x80, 0x9b, 0x95, 0x6c, 0x5b, 0xd6, 0x15, 0xef, 0x46, 0x9b, 0x66, 0x36,
	0x81, 0xbe, 0xd2, 0x76, 0x5f, 0xf3, 0x5e, 0x80, 0x71, 0x21, 0x13, 0xbd, 0x10, 0xbc, 0x1f, 0x4d,
	0x9a, 0x89, 0x1d, 0x88, 0x0d, 0x22, 0xa3, 0x99, 0xdd, 0x42, 0x47, 0x49, 0x3e, 0x9c, 0x26, 0xb3,
	0xeb, 0xa2, 0xa3, 0x64, 0xfa, 0x09, 0x93, 0x79, 0x55, 0x6a, 0x8d, 0xdb, 0xcb, 0xc2, 0x39, 0x5c,
	0x89, 0xc8, 0x9b, 0xba, 0xf7, 0xad, 0xba, 0xc4, 0x43, 0xd9, 0xb3, 0x94, 0xbe, 0xc0, 0xb8, 0x15,
	0xb0, 0xa7, 0x70, 0x61, 0xb4, 0xfe, 0x28, 0x79, 0x6a, 0x3b, 0x3a, 0x91, 0x85, 0xfc, 0x58, 0x42,
	0x6a, 0xdc, 0x26, 0xab, 0x8e, 0x16, 0xdd, 0x16, 0xe5, 0x06, 0x5d, 0xb6, 0x2e, 0x57, 0x4e, 0x89,
	0xe6, 0x27, 0xf4, 0x46, 0xdf, 0xcf, 0x1b, 0x55, 0x57, 0xfb, 0x55, 0x26, 0xcc, 0x2e, 0x6f, 0xa9,
	0x79, 0x54, 0xf3, 0xa8, 0xe6, 0xa4, 0xae, 0xe2, 0x13, 0xbe, 0xfd, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x79, 0x94, 0xb3, 0xbd, 0xdd, 0x01, 0x00, 0x00,
}

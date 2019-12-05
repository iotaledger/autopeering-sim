// Code generated by protoc-gen-go. DO NOT EDIT.
// source: discover/proto/message.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto3 "github.com/iotaledger/autopeering-sim/peer/proto"
	proto1 "github.com/iotaledger/autopeering-sim/peer/service/proto"
	proto2 "github.com/iotaledger/autopeering-sim/salt/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Ping struct {
	// protocol version number
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// string form of the return address (e.g. "192.0.2.1:25", "[2001:db8::1]:80")
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	// string form of the recipient address
	To string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	// unix time
	Timestamp            int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_43f14146485f66eb, []int{0}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ping) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Ping) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Ping) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Pong struct {
	// hash of the ping packet
	PingHash []byte `protobuf:"bytes,1,opt,name=ping_hash,json=pingHash,proto3" json:"ping_hash,omitempty"`
	// string form of the recipient address
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// services supported by the sender
	Services             *proto1.ServiceMap `protobuf:"bytes,3,opt,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_43f14146485f66eb, []int{1}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetPingHash() []byte {
	if m != nil {
		return m.PingHash
	}
	return nil
}

func (m *Pong) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Pong) GetServices() *proto1.ServiceMap {
	if m != nil {
		return m.Services
	}
	return nil
}

type DiscoveryRequest struct {
	// string form of the recipient address
	To string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	// unix time
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// salt of the requester
	Salt                 *proto2.Salt `protobuf:"bytes,3,opt,name=salt,proto3" json:"salt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DiscoveryRequest) Reset()         { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()    {}
func (*DiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_43f14146485f66eb, []int{2}
}

func (m *DiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryRequest.Unmarshal(m, b)
}
func (m *DiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryRequest.Marshal(b, m, deterministic)
}
func (m *DiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryRequest.Merge(m, src)
}
func (m *DiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoveryRequest.Size(m)
}
func (m *DiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryRequest proto.InternalMessageInfo

func (m *DiscoveryRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *DiscoveryRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *DiscoveryRequest) GetSalt() *proto2.Salt {
	if m != nil {
		return m.Salt
	}
	return nil
}

type DiscoveryResponse struct {
	// hash of the corresponding request
	ReqHash []byte `protobuf:"bytes,1,opt,name=req_hash,json=reqHash,proto3" json:"req_hash,omitempty"`
	// list of peers
	Peers                []*proto3.Peer `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DiscoveryResponse) Reset()         { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()    {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_43f14146485f66eb, []int{3}
}

func (m *DiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryResponse.Unmarshal(m, b)
}
func (m *DiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryResponse.Marshal(b, m, deterministic)
}
func (m *DiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryResponse.Merge(m, src)
}
func (m *DiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_DiscoveryResponse.Size(m)
}
func (m *DiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryResponse proto.InternalMessageInfo

func (m *DiscoveryResponse) GetReqHash() []byte {
	if m != nil {
		return m.ReqHash
	}
	return nil
}

func (m *DiscoveryResponse) GetPeers() []*proto3.Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func init() {
	proto.RegisterType((*Ping)(nil), "proto.Ping")
	proto.RegisterType((*Pong)(nil), "proto.Pong")
	proto.RegisterType((*DiscoveryRequest)(nil), "proto.DiscoveryRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "proto.DiscoveryResponse")
}

func init() { proto.RegisterFile("discover/proto/message.proto", fileDescriptor_43f14146485f66eb) }

var fileDescriptor_43f14146485f66eb = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x4d, 0x4b, 0x11, 0x18, 0xd4, 0xc8, 0x26, 0x24, 0x15, 0x49, 0xac, 0x9c, 0xb8, 0x40, 0x13,
	0x35, 0x7e, 0x80, 0xf1, 0xe0, 0xc5, 0x04, 0xd7, 0x9b, 0x17, 0xb3, 0xc0, 0xb8, 0xdd, 0xa4, 0xed,
	0x96, 0x9d, 0x85, 0xc4, 0xbf, 0x37, 0xdd, 0x2e, 0x88, 0xc6, 0xd3, 0xce, 0xbc, 0xf7, 0xf2, 0xde,
	0xcc, 0xb4, 0x30, 0x5e, 0x2b, 0x5a, 0xe9, 0x1d, 0x9a, 0xb4, 0x32, 0xda, 0xea, 0xb4, 0x40, 0x22,
	0x21, 0x71, 0xee, 0x3a, 0xd6, 0x76, 0xcf, 0x68, 0x58, 0xe1, 0x41, 0x50, 0x97, 0x0d, 0x3b, 0x4a,
	0x1c, 0x4c, 0x68, 0x76, 0x6a, 0x85, 0x9e, 0xf6, 0x9d, 0x57, 0x0c, 0x49, 0xe4, 0x76, 0xcf, 0x88,
	0xdc, 0x36, 0xf0, 0x64, 0x09, 0xd1, 0x42, 0x95, 0x92, 0xc5, 0xd0, 0xd9, 0xa1, 0x21, 0xa5, 0xcb,
	0x38, 0x48, 0x82, 0xe9, 0x19, 0xdf, 0xb7, 0x8c, 0x41, 0xf4, 0x69, 0x74, 0x11, 0x87, 0x49, 0x30,
	0xed, 0x71, 0x57, 0xb3, 0x73, 0x08, 0xad, 0x8e, 0x5b, 0x0e, 0x09, 0xad, 0x66, 0x63, 0xe8, 0x59,
	0x55, 0x20, 0x59, 0x51, 0x54, 0x71, 0x94, 0x04, 0xd3, 0x16, 0xff, 0x01, 0x5c, 0x86, 0x2e, 0x25,
	0xbb, 0x82, 0x5e, 0xa5, 0x4a, 0xf9, 0x91, 0x09, 0xca, 0x5c, 0xca, 0x29, 0xef, 0xd6, 0xc0, 0xb3,
	0xa0, 0xcc, 0x5b, 0x86, 0x07, 0xcb, 0x19, 0x74, 0xfd, 0x02, 0xe4, 0x82, 0xfa, 0xb7, 0x83, 0x66,
	0xe4, 0xf9, 0x5b, 0x03, 0xbf, 0x88, 0x8a, 0x1f, 0x24, 0x13, 0x01, 0x17, 0x4f, 0xfe, 0x7c, 0x5f,
	0x1c, 0x37, 0x5b, 0x24, 0xeb, 0x2d, 0x83, 0xff, 0xa7, 0x0c, 0xff, 0x4c, 0xc9, 0xae, 0x21, 0xaa,
	0xef, 0xe2, 0xc3, 0xfa, 0xfb, 0x30, 0x91, 0x5b, 0xee, 0x88, 0xc9, 0x2b, 0x0c, 0x8e, 0x22, 0xa8,
	0xd2, 0x25, 0x21, 0xbb, 0x84, 0xae, 0xc1, 0xcd, 0xf1, 0x4a, 0x1d, 0x83, 0x1b, 0xb7, 0xd1, 0x0d,
	0xb4, 0xeb, 0xaf, 0x42, 0x71, 0x98, 0xb4, 0x8e, 0x1c, 0x17, 0x88, 0x86, 0x37, 0xcc, 0xe3, 0xc3,
	0xfb, 0xbd, 0x54, 0x36, 0xdb, 0x2e, 0xe7, 0x2b, 0x5d, 0xa4, 0x4a, 0x5b, 0x91, 0xe3, 0x5a, 0xa2,
	0x49, 0xc5, 0xd6, 0xea, 0x5a, 0xa2, 0x4a, 0x39, 0x23, 0x55, 0xa4, 0xbf, 0x7f, 0x8d, 0xe5, 0x89,
	0x7b, 0xee, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x34, 0xc4, 0x71, 0x81, 0x33, 0x02, 0x00, 0x00,
}

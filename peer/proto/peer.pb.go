// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/proto/peer.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Minimal encoding of a peer
type Peer struct {
	// public key used for signing
	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// address of autopeering protocol (e.g. "192.0.2.1:25", "[2001:db8::1]:80")
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_155860cd2f47eba7, []int{0}
}

func (m *Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peer.Unmarshal(m, b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return xxx_messageInfo_Peer.Size(m)
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Peer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Peer)(nil), "proto.peer")
}

func init() { proto.RegisterFile("peer/proto/peer.proto", fileDescriptor_155860cd2f47eba7) }

var fileDescriptor_155860cd2f47eba7 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x48, 0x4d, 0x2d,
	0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x07, 0x31, 0xf5, 0xc0, 0x4c, 0x21, 0x56, 0x30, 0xa5,
	0x64, 0xcf, 0xc5, 0x02, 0x12, 0x14, 0x92, 0xe5, 0xe2, 0x2a, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0x8e,
	0xcf, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0xe2, 0x84, 0x88, 0x78, 0xa7, 0x56,
	0x0a, 0x49, 0x70, 0xb1, 0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x29, 0x30, 0x6a,
	0x70, 0x06, 0xc1, 0xb8, 0x4e, 0x5a, 0x51, 0x1a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0xe5, 0xf9, 0x39, 0x39, 0x89, 0xc9, 0xfa, 0x89, 0xa5, 0x25, 0xf9, 0x20, 0x63,
	0x33, 0xf3, 0xd2, 0xf5, 0x11, 0xd6, 0x27, 0xb1, 0x81, 0x29, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x8b, 0x77, 0x50, 0x63, 0x93, 0x00, 0x00, 0x00,
}

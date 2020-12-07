// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test

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

type PairUint32 struct {
	First                *uint32  `protobuf:"varint,1,opt,name=first" json:"first,omitempty"`
	Second               *uint32  `protobuf:"varint,2,opt,name=second" json:"second,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PairUint32) Reset()         { *m = PairUint32{} }
func (m *PairUint32) String() string { return proto.CompactTextString(m) }
func (*PairUint32) ProtoMessage()    {}
func (*PairUint32) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *PairUint32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairUint32.Unmarshal(m, b)
}
func (m *PairUint32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairUint32.Marshal(b, m, deterministic)
}
func (m *PairUint32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairUint32.Merge(m, src)
}
func (m *PairUint32) XXX_Size() int {
	return xxx_messageInfo_PairUint32.Size(m)
}
func (m *PairUint32) XXX_DiscardUnknown() {
	xxx_messageInfo_PairUint32.DiscardUnknown(m)
}

var xxx_messageInfo_PairUint32 proto.InternalMessageInfo

func (m *PairUint32) GetFirst() uint32 {
	if m != nil && m.First != nil {
		return *m.First
	}
	return 0
}

func (m *PairUint32) GetSecond() uint32 {
	if m != nil && m.Second != nil {
		return *m.Second
	}
	return 0
}

type PairInt32 struct {
	First                *int32   `protobuf:"varint,1,opt,name=first" json:"first,omitempty"`
	Second               *int32   `protobuf:"varint,2,opt,name=second" json:"second,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PairInt32) Reset()         { *m = PairInt32{} }
func (m *PairInt32) String() string { return proto.CompactTextString(m) }
func (*PairInt32) ProtoMessage()    {}
func (*PairInt32) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *PairInt32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairInt32.Unmarshal(m, b)
}
func (m *PairInt32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairInt32.Marshal(b, m, deterministic)
}
func (m *PairInt32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairInt32.Merge(m, src)
}
func (m *PairInt32) XXX_Size() int {
	return xxx_messageInfo_PairInt32.Size(m)
}
func (m *PairInt32) XXX_DiscardUnknown() {
	xxx_messageInfo_PairInt32.DiscardUnknown(m)
}

var xxx_messageInfo_PairInt32 proto.InternalMessageInfo

func (m *PairInt32) GetFirst() int32 {
	if m != nil && m.First != nil {
		return *m.First
	}
	return 0
}

func (m *PairInt32) GetSecond() int32 {
	if m != nil && m.Second != nil {
		return *m.Second
	}
	return 0
}

func init() {
	proto.RegisterType((*PairUint32)(nil), "test.PairUint32")
	proto.RegisterType((*PairInt32)(nil), "test.PairInt32")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xac, 0xb8, 0xb8, 0x02, 0x12,
	0x33, 0x8b, 0x42, 0x33, 0xf3, 0x4a, 0x8c, 0x8d, 0x84, 0x44, 0xb8, 0x58, 0xd3, 0x32, 0x8b, 0x8a,
	0x4b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x83, 0x20, 0x1c, 0x21, 0x31, 0x2e, 0xb6, 0xe2, 0xd4,
	0xe4, 0xfc, 0xbc, 0x14, 0x09, 0x26, 0xb0, 0x30, 0x94, 0xa7, 0x64, 0xc9, 0xc5, 0x09, 0xd2, 0xeb,
	0x89, 0xa9, 0x95, 0x15, 0xbb, 0x56, 0x56, 0x98, 0x56, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42,
	0x7e, 0x11, 0x9f, 0x89, 0x00, 0x00, 0x00,
}

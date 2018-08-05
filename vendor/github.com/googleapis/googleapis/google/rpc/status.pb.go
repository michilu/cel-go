// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/rpc/status.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Status struct {
	Code                 int32      `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Details              []*any.Any `protobuf:"bytes,3,rep,name=details" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_a9e33ab1692c5f4c, []int{0}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Status) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "google.rpc.Status")
}

func init() { proto.RegisterFile("google/rpc/status.proto", fileDescriptor_status_a9e33ab1692c5f4c) }

var fileDescriptor_status_a9e33ab1692c5f4c = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0x48, 0xe8, 0x15, 0x15, 0x24, 0x4b, 0x49, 0x42, 0x15, 0x81,
	0x65, 0x92, 0x4a, 0xd3, 0xf4, 0x13, 0xf3, 0x2a, 0x21, 0xca, 0x94, 0xd2, 0xb8, 0xd8, 0x82, 0xc1,
	0xda, 0x84, 0x84, 0xb8, 0x58, 0x92, 0xf3, 0x53, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83,
	0xc0, 0x6c, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x26, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x48, 0x8f, 0x8b, 0x3d, 0x25, 0xb5, 0x24, 0x31, 0x33, 0xa7,
	0x58, 0x82, 0x59, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x44, 0x0f, 0x6a, 0x21, 0xcc, 0x12, 0x3d, 0xc7,
	0xbc, 0xca, 0x20, 0x98, 0x22, 0x27, 0x79, 0x2e, 0xbe, 0xe4, 0xfc, 0x5c, 0x3d, 0x84, 0xa3, 0x9c,
	0xb8, 0x21, 0xf6, 0x06, 0x80, 0x94, 0x07, 0x30, 0x26, 0xb1, 0x81, 0xf5, 0x19, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x7f, 0x92, 0x58, 0x7f, 0xd1, 0x00, 0x00, 0x00,
}

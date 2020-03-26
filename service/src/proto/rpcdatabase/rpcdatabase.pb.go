// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpcdatabase.proto

package rpcdatabase

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

////////////////////////////////////////////
// mysql
////////////////////////////////////////////
type Prop struct {
	Propid               int32    `protobuf:"varint,1,opt,name=propid,proto3" json:"propid,omitempty"`
	Proptype             int32    `protobuf:"varint,2,opt,name=proptype,proto3" json:"proptype,omitempty"`
	Expire               int32    `protobuf:"varint,3,opt,name=expire,proto3" json:"expire,omitempty"`
	Count                int64    `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Prop) Reset()         { *m = Prop{} }
func (m *Prop) String() string { return proto.CompactTextString(m) }
func (*Prop) ProtoMessage()    {}
func (*Prop) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe2a72ba7550c07e, []int{0}
}

func (m *Prop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prop.Unmarshal(m, b)
}
func (m *Prop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prop.Marshal(b, m, deterministic)
}
func (m *Prop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prop.Merge(m, src)
}
func (m *Prop) XXX_Size() int {
	return xxx_messageInfo_Prop.Size(m)
}
func (m *Prop) XXX_DiscardUnknown() {
	xxx_messageInfo_Prop.DiscardUnknown(m)
}

var xxx_messageInfo_Prop proto.InternalMessageInfo

func (m *Prop) GetPropid() int32 {
	if m != nil {
		return m.Propid
	}
	return 0
}

func (m *Prop) GetProptype() int32 {
	if m != nil {
		return m.Proptype
	}
	return 0
}

func (m *Prop) GetExpire() int32 {
	if m != nil {
		return m.Expire
	}
	return 0
}

func (m *Prop) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ReqPlayerInfo struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqPlayerInfo) Reset()         { *m = ReqPlayerInfo{} }
func (m *ReqPlayerInfo) String() string { return proto.CompactTextString(m) }
func (*ReqPlayerInfo) ProtoMessage()    {}
func (*ReqPlayerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe2a72ba7550c07e, []int{1}
}

func (m *ReqPlayerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqPlayerInfo.Unmarshal(m, b)
}
func (m *ReqPlayerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqPlayerInfo.Marshal(b, m, deterministic)
}
func (m *ReqPlayerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqPlayerInfo.Merge(m, src)
}
func (m *ReqPlayerInfo) XXX_Size() int {
	return xxx_messageInfo_ReqPlayerInfo.Size(m)
}
func (m *ReqPlayerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqPlayerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReqPlayerInfo proto.InternalMessageInfo

func (m *ReqPlayerInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RespPlayerInfo struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Rich                 int64    `protobuf:"varint,3,opt,name=rich,proto3" json:"rich,omitempty"`
	Sex                  int32    `protobuf:"varint,4,opt,name=sex,proto3" json:"sex,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Bag                  []*Prop  `protobuf:"bytes,6,rep,name=bag,proto3" json:"bag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespPlayerInfo) Reset()         { *m = RespPlayerInfo{} }
func (m *RespPlayerInfo) String() string { return proto.CompactTextString(m) }
func (*RespPlayerInfo) ProtoMessage()    {}
func (*RespPlayerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe2a72ba7550c07e, []int{2}
}

func (m *RespPlayerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespPlayerInfo.Unmarshal(m, b)
}
func (m *RespPlayerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespPlayerInfo.Marshal(b, m, deterministic)
}
func (m *RespPlayerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespPlayerInfo.Merge(m, src)
}
func (m *RespPlayerInfo) XXX_Size() int {
	return xxx_messageInfo_RespPlayerInfo.Size(m)
}
func (m *RespPlayerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RespPlayerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RespPlayerInfo proto.InternalMessageInfo

func (m *RespPlayerInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RespPlayerInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RespPlayerInfo) GetRich() int64 {
	if m != nil {
		return m.Rich
	}
	return 0
}

func (m *RespPlayerInfo) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *RespPlayerInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RespPlayerInfo) GetBag() []*Prop {
	if m != nil {
		return m.Bag
	}
	return nil
}

func init() {
	proto.RegisterType((*Prop)(nil), "rpcdatabase.Prop")
	proto.RegisterType((*ReqPlayerInfo)(nil), "rpcdatabase.ReqPlayerInfo")
	proto.RegisterType((*RespPlayerInfo)(nil), "rpcdatabase.RespPlayerInfo")
}

func init() {
	proto.RegisterFile("rpcdatabase.proto", fileDescriptor_fe2a72ba7550c07e)
}

var fileDescriptor_fe2a72ba7550c07e = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x69, 0xd3, 0x16, 0x9d, 0xc5, 0xc5, 0x1d, 0x44, 0x82, 0x17, 0x4b, 0xbd, 0xf4, 0xb4,
	0x07, 0x7d, 0x0a, 0x6f, 0xcb, 0xbc, 0x41, 0xda, 0x8e, 0xb6, 0xa0, 0xc9, 0x98, 0x46, 0xd8, 0x7d,
	0x10, 0xdf, 0x57, 0x92, 0xc8, 0xb2, 0x17, 0x6f, 0xff, 0x37, 0x33, 0x3f, 0xf3, 0xf3, 0xc3, 0xce,
	0xcb, 0x38, 0x99, 0x60, 0x06, 0xb3, 0xf2, 0x5e, 0xbc, 0x0b, 0x0e, 0x37, 0x17, 0xa3, 0x6e, 0x86,
	0xea, 0xe0, 0x9d, 0xe0, 0x3d, 0x34, 0xe2, 0x9d, 0x2c, 0x93, 0x2e, 0xda, 0xa2, 0xaf, 0xe9, 0x8f,
	0xf0, 0x01, 0xae, 0xa2, 0x0a, 0x27, 0x61, 0x5d, 0xa6, 0xcd, 0x99, 0xa3, 0x87, 0x8f, 0xb2, 0x78,
	0xd6, 0x2a, 0x7b, 0x32, 0xe1, 0x1d, 0xd4, 0xa3, 0xfb, 0xb6, 0x41, 0x57, 0x6d, 0xd1, 0x2b, 0xca,
	0xd0, 0x3d, 0xc2, 0x0d, 0xf1, 0xd7, 0xe1, 0xc3, 0x9c, 0xd8, 0xbf, 0xda, 0x37, 0x87, 0x5b, 0x28,
	0xcf, 0xef, 0xca, 0x65, 0xea, 0x7e, 0x0a, 0xd8, 0x12, 0xaf, 0xf2, 0xff, 0x09, 0x22, 0x54, 0xd6,
	0x7c, 0xe6, 0x24, 0xd7, 0x94, 0x74, 0x9c, 0xf9, 0x65, 0x9c, 0x53, 0x06, 0x45, 0x49, 0xe3, 0x2d,
	0xa8, 0x95, 0x8f, 0xe9, 0x7f, 0x4d, 0x51, 0xc6, 0x4c, 0x32, 0x3b, 0xcb, 0xba, 0x4e, 0xd6, 0x0c,
	0xf8, 0x04, 0x6a, 0x30, 0xef, 0xba, 0x69, 0x55, 0xbf, 0x79, 0xde, 0xed, 0x2f, 0xbb, 0x8a, 0xad,
	0x50, 0xdc, 0x0e, 0x4d, 0xaa, 0xed, 0xe5, 0x37, 0x00, 0x00, 0xff, 0xff, 0xed, 0x0c, 0x62, 0x2c,
	0x4b, 0x01, 0x00, 0x00,
}

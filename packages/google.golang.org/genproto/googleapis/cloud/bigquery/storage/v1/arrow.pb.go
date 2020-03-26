// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/bigquery/storage/v1/arrow.proto

package storage

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Arrow schema as specified in
// https://arrow.apache.org/docs/python/api/datatypes.html
// and serialized to bytes using IPC:
// https://arrow.apache.org/docs/ipc.html.
//
// See code samples on how this message can be deserialized.
type ArrowSchema struct {
	// IPC serialized Arrow schema.
	SerializedSchema     []byte   `protobuf:"bytes,1,opt,name=serialized_schema,json=serializedSchema,proto3" json:"serialized_schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrowSchema) Reset()         { *m = ArrowSchema{} }
func (m *ArrowSchema) String() string { return proto.CompactTextString(m) }
func (*ArrowSchema) ProtoMessage()    {}
func (*ArrowSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d5179f91c6841b3, []int{0}
}

func (m *ArrowSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrowSchema.Unmarshal(m, b)
}
func (m *ArrowSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrowSchema.Marshal(b, m, deterministic)
}
func (m *ArrowSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrowSchema.Merge(m, src)
}
func (m *ArrowSchema) XXX_Size() int {
	return xxx_messageInfo_ArrowSchema.Size(m)
}
func (m *ArrowSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrowSchema.DiscardUnknown(m)
}

var xxx_messageInfo_ArrowSchema proto.InternalMessageInfo

func (m *ArrowSchema) GetSerializedSchema() []byte {
	if m != nil {
		return m.SerializedSchema
	}
	return nil
}

// Arrow RecordBatch.
type ArrowRecordBatch struct {
	// IPC-serialized Arrow RecordBatch.
	SerializedRecordBatch []byte `protobuf:"bytes,1,opt,name=serialized_record_batch,json=serializedRecordBatch,proto3" json:"serialized_record_batch,omitempty"`
	// The count of rows in `serialized_record_batch`.
	RowCount             int64    `protobuf:"varint,2,opt,name=row_count,json=rowCount,proto3" json:"row_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrowRecordBatch) Reset()         { *m = ArrowRecordBatch{} }
func (m *ArrowRecordBatch) String() string { return proto.CompactTextString(m) }
func (*ArrowRecordBatch) ProtoMessage()    {}
func (*ArrowRecordBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d5179f91c6841b3, []int{1}
}

func (m *ArrowRecordBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrowRecordBatch.Unmarshal(m, b)
}
func (m *ArrowRecordBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrowRecordBatch.Marshal(b, m, deterministic)
}
func (m *ArrowRecordBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrowRecordBatch.Merge(m, src)
}
func (m *ArrowRecordBatch) XXX_Size() int {
	return xxx_messageInfo_ArrowRecordBatch.Size(m)
}
func (m *ArrowRecordBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrowRecordBatch.DiscardUnknown(m)
}

var xxx_messageInfo_ArrowRecordBatch proto.InternalMessageInfo

func (m *ArrowRecordBatch) GetSerializedRecordBatch() []byte {
	if m != nil {
		return m.SerializedRecordBatch
	}
	return nil
}

func (m *ArrowRecordBatch) GetRowCount() int64 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ArrowSchema)(nil), "google.cloud.bigquery.storage.v1.ArrowSchema")
	proto.RegisterType((*ArrowRecordBatch)(nil), "google.cloud.bigquery.storage.v1.ArrowRecordBatch")
}

func init() {
	proto.RegisterFile("google/cloud/bigquery/storage/v1/arrow.proto", fileDescriptor_0d5179f91c6841b3)
}

var fileDescriptor_0d5179f91c6841b3 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x49, 0x7f, 0xf8, 0xd1, 0xd1, 0x45, 0x0d, 0x88, 0x05, 0x37, 0xa1, 0xb8, 0x28, 0x28,
	0x33, 0x04, 0xc1, 0x85, 0xae, 0x4c, 0x17, 0xdd, 0xd6, 0x14, 0xba, 0x90, 0x40, 0x98, 0x4c, 0x86,
	0xe9, 0x40, 0x9a, 0x5b, 0x6f, 0x92, 0x06, 0x7d, 0x24, 0x5f, 0xc3, 0x9d, 0x4f, 0x25, 0xb9, 0x99,
	0xd0, 0x6e, 0xc4, 0xdd, 0xcc, 0x9c, 0xef, 0x9c, 0x33, 0xf7, 0xb2, 0x3b, 0x03, 0x60, 0x0a, 0x2d,
	0x54, 0x01, 0x4d, 0x2e, 0x32, 0x6b, 0xde, 0x1a, 0x8d, 0xef, 0xa2, 0xaa, 0x01, 0xa5, 0xd1, 0x62,
	0x1f, 0x0a, 0x89, 0x08, 0x2d, 0xdf, 0x21, 0xd4, 0xe0, 0x07, 0x3d, 0xcd, 0x89, 0xe6, 0x03, 0xcd,
	0x1d, 0xcd, 0xf7, 0xe1, 0xf4, 0x91, 0x9d, 0x3d, 0x77, 0x86, 0x95, 0xda, 0xe8, 0xad, 0xf4, 0x6f,
	0xd9, 0x45, 0xa5, 0xd1, 0xca, 0xc2, 0x7e, 0xe8, 0x3c, 0xad, 0xe8, 0x71, 0xe2, 0x05, 0xde, 0xec,
	0x3c, 0x1e, 0x1f, 0x84, 0x1e, 0x9e, 0x1a, 0x36, 0x26, 0x6f, 0xac, 0x15, 0x60, 0x1e, 0xc9, 0x5a,
	0x6d, 0xfc, 0x07, 0x76, 0x75, 0x14, 0x80, 0xa4, 0xa4, 0x59, 0x27, 0xb9, 0x98, 0xcb, 0x83, 0x7c,
	0xec, 0xbb, 0x66, 0xa7, 0x08, 0x6d, 0xaa, 0xa0, 0x29, 0xeb, 0xc9, 0x28, 0xf0, 0x66, 0xff, 0xe2,
	0x13, 0x84, 0x76, 0xde, 0xdd, 0xa3, 0x2f, 0x8f, 0xdd, 0x28, 0xd8, 0xf2, 0xbf, 0xa6, 0x89, 0x18,
	0xfd, 0x67, 0xd9, 0xcd, 0xbe, 0xf4, 0x5e, 0x17, 0x8e, 0x37, 0x50, 0xc8, 0xd2, 0x70, 0x40, 0x23,
	0x8c, 0x2e, 0x69, 0x33, 0xa2, 0x97, 0xe4, 0xce, 0x56, 0xbf, 0xaf, 0xf2, 0xc9, 0x1d, 0x3f, 0x47,
	0xc1, 0xa2, 0x4f, 0x9a, 0x53, 0x73, 0x64, 0xcd, 0x0b, 0x35, 0xaf, 0x5c, 0xf3, 0x3a, 0xfc, 0x1e,
	0x90, 0x84, 0x90, 0x64, 0x40, 0x12, 0x87, 0x24, 0xeb, 0x30, 0xfb, 0x4f, 0xcd, 0xf7, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x11, 0xc3, 0xfb, 0xfa, 0xc3, 0x01, 0x00, 0x00,
}

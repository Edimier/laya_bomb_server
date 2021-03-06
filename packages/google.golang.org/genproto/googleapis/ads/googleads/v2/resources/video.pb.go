// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/video.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A video.
type Video struct {
	// Output only. The resource name of the video.
	// Video resource names have the form:
	//
	// `customers/{customer_id}/videos/{video_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the video.
	Id *wrappers.StringValue `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The owner channel id of the video.
	ChannelId *wrappers.StringValue `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// Output only. The duration of the video in milliseconds.
	DurationMillis *wrappers.Int64Value `protobuf:"bytes,4,opt,name=duration_millis,json=durationMillis,proto3" json:"duration_millis,omitempty"`
	// Output only. The title of the video.
	Title                *wrappers.StringValue `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Video) Reset()         { *m = Video{} }
func (m *Video) String() string { return proto.CompactTextString(m) }
func (*Video) ProtoMessage()    {}
func (*Video) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ce7c184e83c587c, []int{0}
}

func (m *Video) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Video.Unmarshal(m, b)
}
func (m *Video) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Video.Marshal(b, m, deterministic)
}
func (m *Video) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Video.Merge(m, src)
}
func (m *Video) XXX_Size() int {
	return xxx_messageInfo_Video.Size(m)
}
func (m *Video) XXX_DiscardUnknown() {
	xxx_messageInfo_Video.DiscardUnknown(m)
}

var xxx_messageInfo_Video proto.InternalMessageInfo

func (m *Video) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Video) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Video) GetChannelId() *wrappers.StringValue {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *Video) GetDurationMillis() *wrappers.Int64Value {
	if m != nil {
		return m.DurationMillis
	}
	return nil
}

func (m *Video) GetTitle() *wrappers.StringValue {
	if m != nil {
		return m.Title
	}
	return nil
}

func init() {
	proto.RegisterType((*Video)(nil), "google.ads.googleads.v2.resources.Video")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/video.proto", fileDescriptor_6ce7c184e83c587c)
}

var fileDescriptor_6ce7c184e83c587c = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x6b, 0xd4, 0x40,
	0x1c, 0xc6, 0x49, 0xe2, 0x0a, 0x1d, 0xdf, 0x20, 0xa7, 0x58, 0x4b, 0xdd, 0x2a, 0x95, 0x5e, 0x9c,
	0xc1, 0x28, 0x0a, 0xe3, 0xc5, 0x59, 0x90, 0xba, 0x8a, 0x52, 0x56, 0xc8, 0x41, 0x16, 0x96, 0xd9,
	0xcc, 0x34, 0x1d, 0x48, 0x66, 0xc2, 0x4c, 0xb2, 0x1e, 0x4a, 0x3f, 0x87, 0x77, 0x8f, 0x7e, 0x14,
	0x3f, 0x45, 0xcf, 0xfd, 0x02, 0x82, 0x27, 0xd9, 0x79, 0x49, 0x17, 0x84, 0xb6, 0xa7, 0x3c, 0xe1,
	0xf9, 0x3d, 0x0f, 0x4f, 0xc2, 0x1f, 0x3c, 0xaf, 0x94, 0xaa, 0x6a, 0x8e, 0x28, 0x33, 0xc8, 0xc9,
	0xb5, 0x5a, 0xe5, 0x48, 0x73, 0xa3, 0x7a, 0x5d, 0x72, 0x83, 0x56, 0x82, 0x71, 0x05, 0x5b, 0xad,
	0x3a, 0x95, 0xee, 0x39, 0x06, 0x52, 0x66, 0xe0, 0x80, 0xc3, 0x55, 0x0e, 0x07, 0x7c, 0xfb, 0x71,
	0x68, 0x6c, 0x05, 0x3a, 0x16, 0xbc, 0x66, 0x8b, 0x25, 0x3f, 0xa1, 0x2b, 0xa1, 0xb4, 0xeb, 0xd8,
	0x7e, 0xb8, 0x01, 0x84, 0x98, 0xb7, 0x76, 0xbd, 0x65, 0xdf, 0x96, 0xfd, 0x31, 0xfa, 0xae, 0x69,
	0xdb, 0x72, 0x6d, 0xbc, 0xbf, 0xb3, 0x11, 0xa5, 0x52, 0xaa, 0x8e, 0x76, 0x42, 0x49, 0xef, 0x3e,
	0xf9, 0x91, 0x80, 0x51, 0xb1, 0x1e, 0x9b, 0x7e, 0x02, 0xf7, 0x42, 0xf3, 0x42, 0xd2, 0x86, 0x67,
	0xd1, 0x38, 0x3a, 0xd8, 0x9a, 0x3c, 0x3b, 0x27, 0xc9, 0x5f, 0x32, 0x06, 0xbb, 0x97, 0xd3, 0xbd,
	0x6a, 0x85, 0x81, 0xa5, 0x6a, 0x90, 0x8d, 0xcf, 0xee, 0x86, 0xf0, 0x17, 0xda, 0xf0, 0xf4, 0x05,
	0x88, 0x05, 0xcb, 0xe2, 0x71, 0x74, 0x70, 0x27, 0xdf, 0xf1, 0x01, 0x18, 0x16, 0xc2, 0xaf, 0x9d,
	0x16, 0xb2, 0x2a, 0x68, 0xdd, 0xf3, 0x49, 0x72, 0x4e, 0x92, 0x59, 0x2c, 0x58, 0xfa, 0x0e, 0x80,
	0xf2, 0x84, 0x4a, 0xc9, 0xeb, 0x85, 0x60, 0x59, 0x72, 0xd3, 0xe8, 0x96, 0x0f, 0x4d, 0x59, 0x3a,
	0x05, 0x0f, 0x58, 0xaf, 0xed, 0xe7, 0x2d, 0x1a, 0x51, 0xd7, 0xc2, 0x64, 0xb7, 0x6c, 0xcd, 0xa3,
	0xff, 0x6a, 0xa6, 0xb2, 0x7b, 0xfd, 0x6a, 0xa3, 0xe5, 0x7e, 0x08, 0x7e, 0xb6, 0xb9, 0xf4, 0x0d,
	0x18, 0x75, 0xa2, 0xab, 0x79, 0x36, 0xba, 0xe9, 0x0e, 0xc7, 0xe3, 0x0f, 0x17, 0xe4, 0xfd, 0x75,
	0xff, 0x2a, 0x7d, 0x5a, 0xf6, 0xa6, 0x53, 0x0d, 0xd7, 0x06, 0x9d, 0x06, 0x79, 0xe6, 0x6e, 0xc6,
	0xa0, 0x53, 0xfb, 0x3c, 0x9b, 0xfc, 0x89, 0xc0, 0x7e, 0xa9, 0x1a, 0x78, 0xed, 0xf5, 0x4c, 0x80,
	0x6d, 0x3d, 0x5a, 0x4f, 0x3b, 0x8a, 0xbe, 0x7d, 0xf4, 0x81, 0x4a, 0xd5, 0x54, 0x56, 0x50, 0xe9,
	0x0a, 0x55, 0x5c, 0xda, 0xe1, 0xe8, 0x72, 0xcd, 0x15, 0xc7, 0xfb, 0x76, 0x50, 0x3f, 0xe3, 0xe4,
	0x90, 0x90, 0x5f, 0xf1, 0xde, 0xa1, 0xab, 0x24, 0xcc, 0x40, 0x27, 0xd7, 0xaa, 0xc8, 0xe1, 0x2c,
	0x90, 0xbf, 0x03, 0x33, 0x27, 0xcc, 0xcc, 0x07, 0x66, 0x5e, 0xe4, 0xf3, 0x81, 0xb9, 0x88, 0xf7,
	0x9d, 0x81, 0x31, 0x61, 0x06, 0xe3, 0x81, 0xc2, 0xb8, 0xc8, 0x31, 0x1e, 0xb8, 0xe5, 0x6d, 0x3b,
	0xf6, 0xe5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x2c, 0x8e, 0x0d, 0x68, 0x03, 0x00, 0x00,
}

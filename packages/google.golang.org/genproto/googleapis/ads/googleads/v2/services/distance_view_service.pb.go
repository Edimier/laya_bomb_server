// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/distance_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [DistanceViewService.GetDistanceView][google.ads.googleads.v2.services.DistanceViewService.GetDistanceView].
type GetDistanceViewRequest struct {
	// Required. The resource name of the distance view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDistanceViewRequest) Reset()         { *m = GetDistanceViewRequest{} }
func (m *GetDistanceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetDistanceViewRequest) ProtoMessage()    {}
func (*GetDistanceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4de60c8631d2d9d8, []int{0}
}

func (m *GetDistanceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDistanceViewRequest.Unmarshal(m, b)
}
func (m *GetDistanceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDistanceViewRequest.Marshal(b, m, deterministic)
}
func (m *GetDistanceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDistanceViewRequest.Merge(m, src)
}
func (m *GetDistanceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetDistanceViewRequest.Size(m)
}
func (m *GetDistanceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDistanceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDistanceViewRequest proto.InternalMessageInfo

func (m *GetDistanceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDistanceViewRequest)(nil), "google.ads.googleads.v2.services.GetDistanceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/distance_view_service.proto", fileDescriptor_4de60c8631d2d9d8)
}

var fileDescriptor_4de60c8631d2d9d8 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xbd, 0x8a, 0xd5, 0x40,
	0x14, 0x26, 0x59, 0x10, 0x0c, 0x8a, 0x10, 0x41, 0x2f, 0x51, 0xf0, 0xb2, 0x6c, 0xb1, 0x6c, 0x31,
	0x83, 0x11, 0x41, 0x46, 0x2d, 0x26, 0x08, 0xb1, 0x92, 0x65, 0x85, 0x14, 0x12, 0x08, 0xb3, 0xc9,
	0x31, 0x0e, 0x24, 0x33, 0xd7, 0x9c, 0xb9, 0xd9, 0x42, 0x6c, 0x7c, 0x05, 0xdf, 0xc0, 0xd2, 0xf7,
	0xb0, 0xd9, 0xd6, 0xce, 0xca, 0xc2, 0xca, 0x47, 0xb0, 0x51, 0xb2, 0x93, 0xc9, 0x66, 0xdd, 0xbd,
	0xdc, 0xee, 0x23, 0xdf, 0xcf, 0x39, 0xe7, 0x9b, 0x04, 0xcf, 0x6a, 0xad, 0xeb, 0x06, 0xa8, 0xa8,
	0x90, 0x5a, 0x38, 0xa0, 0x3e, 0xa6, 0x08, 0x5d, 0x2f, 0x4b, 0x40, 0x5a, 0x49, 0x34, 0x42, 0x95,
	0x50, 0xf4, 0x12, 0x4e, 0x8a, 0xf1, 0x33, 0x59, 0x75, 0xda, 0xe8, 0x70, 0x69, 0x2d, 0x44, 0x54,
	0x48, 0x26, 0x37, 0xe9, 0x63, 0xe2, 0xdc, 0xd1, 0xe3, 0x4d, 0xf9, 0x1d, 0xa0, 0x5e, 0x77, 0x97,
	0x06, 0xd8, 0xe0, 0xe8, 0xbe, 0xb3, 0xad, 0x24, 0x15, 0x4a, 0x69, 0x23, 0x8c, 0xd4, 0x0a, 0x47,
	0xf6, 0xee, 0x8c, 0x2d, 0x1b, 0x09, 0xca, 0x8c, 0xc4, 0x83, 0x19, 0xf1, 0x56, 0x42, 0x53, 0x15,
	0xc7, 0xf0, 0x4e, 0xf4, 0x52, 0x77, 0x56, 0xb0, 0x9b, 0x04, 0x77, 0x52, 0x30, 0x2f, 0xc6, 0x89,
	0x99, 0x84, 0x93, 0x23, 0x78, 0xbf, 0x06, 0x34, 0xe1, 0x7e, 0x70, 0xd3, 0xad, 0x54, 0x28, 0xd1,
	0xc2, 0xc2, 0x5b, 0x7a, 0xfb, 0xd7, 0x93, 0x9d, 0x9f, 0xdc, 0x3f, 0xba, 0xe1, 0x98, 0x57, 0xa2,
	0x85, 0xf8, 0x8f, 0x17, 0xdc, 0x9e, 0x27, 0xbc, 0xb6, 0xb7, 0x86, 0xdf, 0xbc, 0xe0, 0xd6, 0x7f,
	0xe1, 0xe1, 0x13, 0xb2, 0xad, 0x21, 0x72, 0xf5, 0x3e, 0x11, 0xdd, 0xe8, 0x9c, 0x9a, 0x23, 0x73,
	0xdf, 0x6e, 0xfa, 0x83, 0x5f, 0xbc, 0xe0, 0xd3, 0xf7, 0x5f, 0x9f, 0xfd, 0x87, 0x21, 0x1d, 0xda,
	0xfe, 0x70, 0x81, 0x79, 0x5e, 0xae, 0xd1, 0xe8, 0x16, 0x3a, 0xa4, 0x07, 0x53, 0xfd, 0x43, 0x08,
	0xd2, 0x83, 0x8f, 0xd1, 0xbd, 0x53, 0xbe, 0x38, 0x1f, 0x38, 0xa2, 0x95, 0x44, 0x52, 0xea, 0x36,
	0xf9, 0xeb, 0x05, 0x7b, 0xa5, 0x6e, 0xb7, 0x9e, 0x95, 0x2c, 0xae, 0xa8, 0xe8, 0x70, 0x78, 0x83,
	0x43, 0xef, 0xcd, 0xcb, 0xd1, 0x5d, 0xeb, 0x46, 0xa8, 0x9a, 0xe8, 0xae, 0xa6, 0x35, 0xa8, 0xb3,
	0x17, 0xa2, 0xe7, 0xf3, 0x36, 0xff, 0x93, 0x4f, 0x1d, 0xf8, 0xe2, 0xef, 0xa4, 0x9c, 0x7f, 0xf5,
	0x97, 0xa9, 0x0d, 0xe4, 0x15, 0x12, 0x0b, 0x07, 0x94, 0xc5, 0x64, 0x1c, 0x8c, 0xa7, 0x4e, 0x92,
	0xf3, 0x0a, 0xf3, 0x49, 0x92, 0x67, 0x71, 0xee, 0x24, 0xbf, 0xfd, 0x3d, 0xfb, 0x9d, 0x31, 0x5e,
	0x21, 0x63, 0x93, 0x88, 0xb1, 0x2c, 0x66, 0xcc, 0xc9, 0x8e, 0xaf, 0x9d, 0xed, 0xf9, 0xe8, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x32, 0x90, 0xf2, 0x3a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DistanceViewServiceClient is the client API for DistanceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DistanceViewServiceClient interface {
	// Returns the attributes of the requested distance view.
	GetDistanceView(ctx context.Context, in *GetDistanceViewRequest, opts ...grpc.CallOption) (*resources.DistanceView, error)
}

type distanceViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDistanceViewServiceClient(cc grpc.ClientConnInterface) DistanceViewServiceClient {
	return &distanceViewServiceClient{cc}
}

func (c *distanceViewServiceClient) GetDistanceView(ctx context.Context, in *GetDistanceViewRequest, opts ...grpc.CallOption) (*resources.DistanceView, error) {
	out := new(resources.DistanceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.DistanceViewService/GetDistanceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DistanceViewServiceServer is the server API for DistanceViewService service.
type DistanceViewServiceServer interface {
	// Returns the attributes of the requested distance view.
	GetDistanceView(context.Context, *GetDistanceViewRequest) (*resources.DistanceView, error)
}

// UnimplementedDistanceViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDistanceViewServiceServer struct {
}

func (*UnimplementedDistanceViewServiceServer) GetDistanceView(ctx context.Context, req *GetDistanceViewRequest) (*resources.DistanceView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDistanceView not implemented")
}

func RegisterDistanceViewServiceServer(s *grpc.Server, srv DistanceViewServiceServer) {
	s.RegisterService(&_DistanceViewService_serviceDesc, srv)
}

func _DistanceViewService_GetDistanceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDistanceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistanceViewServiceServer).GetDistanceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.DistanceViewService/GetDistanceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistanceViewServiceServer).GetDistanceView(ctx, req.(*GetDistanceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DistanceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.DistanceViewService",
	HandlerType: (*DistanceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDistanceView",
			Handler:    _DistanceViewService_GetDistanceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/distance_view_service.proto",
}

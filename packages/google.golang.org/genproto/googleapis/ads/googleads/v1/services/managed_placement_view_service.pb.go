// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/managed_placement_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
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

// Request message for [ManagedPlacementViewService.GetManagedPlacementView][google.ads.googleads.v1.services.ManagedPlacementViewService.GetManagedPlacementView].
type GetManagedPlacementViewRequest struct {
	// Required. The resource name of the Managed Placement View to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetManagedPlacementViewRequest) Reset()         { *m = GetManagedPlacementViewRequest{} }
func (m *GetManagedPlacementViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetManagedPlacementViewRequest) ProtoMessage()    {}
func (*GetManagedPlacementViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c301455ca7db209f, []int{0}
}

func (m *GetManagedPlacementViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetManagedPlacementViewRequest.Unmarshal(m, b)
}
func (m *GetManagedPlacementViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetManagedPlacementViewRequest.Marshal(b, m, deterministic)
}
func (m *GetManagedPlacementViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetManagedPlacementViewRequest.Merge(m, src)
}
func (m *GetManagedPlacementViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetManagedPlacementViewRequest.Size(m)
}
func (m *GetManagedPlacementViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetManagedPlacementViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetManagedPlacementViewRequest proto.InternalMessageInfo

func (m *GetManagedPlacementViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetManagedPlacementViewRequest)(nil), "google.ads.googleads.v1.services.GetManagedPlacementViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/managed_placement_view_service.proto", fileDescriptor_c301455ca7db209f)
}

var fileDescriptor_c301455ca7db209f = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x8b, 0xd4, 0x40,
	0x18, 0x25, 0x39, 0x10, 0x0c, 0xda, 0xa4, 0xb9, 0x23, 0x27, 0x1a, 0x8e, 0x2b, 0x8e, 0x2b, 0x66,
	0x88, 0x16, 0x87, 0x23, 0x8a, 0xb3, 0x20, 0x2b, 0x82, 0xba, 0x9c, 0x90, 0x42, 0x02, 0x61, 0x2e,
	0xf9, 0x8c, 0x03, 0xc9, 0x4c, 0xcc, 0xcc, 0xe6, 0x0a, 0xb1, 0xb1, 0xf1, 0x07, 0xf8, 0x0f, 0xae,
	0xf4, 0xa7, 0x6c, 0x6b, 0x67, 0x65, 0xa1, 0x8d, 0xbf, 0x42, 0xb2, 0x93, 0xc9, 0xee, 0xc2, 0x66,
	0xb7, 0x7b, 0xe4, 0xbd, 0xef, 0xbd, 0x6f, 0xbe, 0x17, 0xef, 0x45, 0x21, 0x65, 0x51, 0x02, 0x66,
	0xb9, 0xc2, 0x06, 0x76, 0xa8, 0x8d, 0xb0, 0x82, 0xa6, 0xe5, 0x19, 0x28, 0x5c, 0x31, 0xc1, 0x0a,
	0xc8, 0xd3, 0xba, 0x64, 0x19, 0x54, 0x20, 0x74, 0xda, 0x72, 0xb8, 0x4e, 0x7b, 0x1e, 0xd5, 0x8d,
	0xd4, 0xd2, 0x0f, 0xcd, 0x2c, 0x62, 0xb9, 0x42, 0x83, 0x0d, 0x6a, 0x23, 0x64, 0x6d, 0x82, 0x67,
	0x63, 0x41, 0x0d, 0x28, 0x39, 0x6f, 0xc6, 0x93, 0x4c, 0x42, 0x70, 0xcf, 0xce, 0xd7, 0x1c, 0x33,
	0x21, 0xa4, 0x66, 0x9a, 0x4b, 0xa1, 0x7a, 0xf6, 0x70, 0x8d, 0xcd, 0x4a, 0x0e, 0x42, 0xf7, 0xc4,
	0x83, 0x35, 0xe2, 0x03, 0x87, 0x32, 0x4f, 0xaf, 0xe0, 0x23, 0x6b, 0xb9, 0x6c, 0x8c, 0xe0, 0xe4,
	0x95, 0x77, 0x7f, 0x0a, 0xfa, 0xb5, 0x89, 0x9e, 0xd9, 0xe4, 0x98, 0xc3, 0xf5, 0x25, 0x7c, 0x9a,
	0x83, 0xd2, 0xfe, 0x99, 0x77, 0xd7, 0xee, 0x98, 0x0a, 0x56, 0xc1, 0x91, 0x13, 0x3a, 0x67, 0xb7,
	0x27, 0x07, 0xbf, 0xa9, 0x7b, 0x79, 0xc7, 0x32, 0x6f, 0x58, 0x05, 0x0f, 0x6f, 0x5c, 0xef, 0x78,
	0x9b, 0xd3, 0x3b, 0x73, 0x04, 0xff, 0xaf, 0xe3, 0x1d, 0x8e, 0x84, 0xf9, 0xcf, 0xd1, 0xbe, 0x13,
	0xa2, 0xdd, 0x7b, 0x06, 0x17, 0xa3, 0x0e, 0xc3, 0x89, 0xd1, 0xb6, 0xf9, 0x93, 0xb7, 0xbf, 0xe8,
	0xe6, 0x0b, 0xbf, 0xfe, 0xfc, 0xf3, 0xdd, 0x7d, 0xec, 0x5f, 0x74, 0xf5, 0x7c, 0xde, 0x60, 0x9e,
	0x66, 0x73, 0xa5, 0x65, 0x05, 0x8d, 0xc2, 0xe7, 0xb6, 0xaf, 0x0d, 0x33, 0x85, 0xcf, 0xbf, 0x04,
	0xc7, 0x0b, 0x7a, 0xb4, 0x5a, 0xa0, 0x47, 0x35, 0x57, 0x28, 0x93, 0xd5, 0xe4, 0x9b, 0xeb, 0x9d,
	0x66, 0xb2, 0xda, 0xfb, 0xdc, 0x49, 0xb8, 0xe3, 0x94, 0xb3, 0xae, 0xbb, 0x99, 0xf3, 0xfe, 0x65,
	0xef, 0x52, 0xc8, 0x92, 0x89, 0x02, 0xc9, 0xa6, 0xc0, 0x05, 0x88, 0x65, 0xb3, 0x78, 0x95, 0x3b,
	0xfe, 0x77, 0x3f, 0xb1, 0xe0, 0xc6, 0x3d, 0x98, 0x52, 0xfa, 0xc3, 0x0d, 0xa7, 0xc6, 0x90, 0xe6,
	0x0a, 0x19, 0xd8, 0xa1, 0x38, 0x42, 0x7d, 0xb0, 0x5a, 0x58, 0x49, 0x42, 0x73, 0x95, 0x0c, 0x92,
	0x24, 0x8e, 0x12, 0x2b, 0xf9, 0xe7, 0x9e, 0x9a, 0xef, 0x84, 0xd0, 0x5c, 0x11, 0x32, 0x88, 0x08,
	0x89, 0x23, 0x42, 0xac, 0xec, 0xea, 0xd6, 0x72, 0xcf, 0x47, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x05, 0xd6, 0xe4, 0x76, 0x84, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ManagedPlacementViewServiceClient is the client API for ManagedPlacementViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ManagedPlacementViewServiceClient interface {
	// Returns the requested Managed Placement view in full detail.
	GetManagedPlacementView(ctx context.Context, in *GetManagedPlacementViewRequest, opts ...grpc.CallOption) (*resources.ManagedPlacementView, error)
}

type managedPlacementViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagedPlacementViewServiceClient(cc grpc.ClientConnInterface) ManagedPlacementViewServiceClient {
	return &managedPlacementViewServiceClient{cc}
}

func (c *managedPlacementViewServiceClient) GetManagedPlacementView(ctx context.Context, in *GetManagedPlacementViewRequest, opts ...grpc.CallOption) (*resources.ManagedPlacementView, error) {
	out := new(resources.ManagedPlacementView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ManagedPlacementViewService/GetManagedPlacementView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagedPlacementViewServiceServer is the server API for ManagedPlacementViewService service.
type ManagedPlacementViewServiceServer interface {
	// Returns the requested Managed Placement view in full detail.
	GetManagedPlacementView(context.Context, *GetManagedPlacementViewRequest) (*resources.ManagedPlacementView, error)
}

// UnimplementedManagedPlacementViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedManagedPlacementViewServiceServer struct {
}

func (*UnimplementedManagedPlacementViewServiceServer) GetManagedPlacementView(ctx context.Context, req *GetManagedPlacementViewRequest) (*resources.ManagedPlacementView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManagedPlacementView not implemented")
}

func RegisterManagedPlacementViewServiceServer(s *grpc.Server, srv ManagedPlacementViewServiceServer) {
	s.RegisterService(&_ManagedPlacementViewService_serviceDesc, srv)
}

func _ManagedPlacementViewService_GetManagedPlacementView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManagedPlacementViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagedPlacementViewServiceServer).GetManagedPlacementView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ManagedPlacementViewService/GetManagedPlacementView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagedPlacementViewServiceServer).GetManagedPlacementView(ctx, req.(*GetManagedPlacementViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ManagedPlacementViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ManagedPlacementViewService",
	HandlerType: (*ManagedPlacementViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetManagedPlacementView",
			Handler:    _ManagedPlacementViewService_GetManagedPlacementView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/managed_placement_view_service.proto",
}

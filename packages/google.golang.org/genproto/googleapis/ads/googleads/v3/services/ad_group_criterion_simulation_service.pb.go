// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/ad_group_criterion_simulation_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
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

// Request message for
// [AdGroupCriterionSimulationService.GetAdGroupCriterionSimulation][google.ads.googleads.v3.services.AdGroupCriterionSimulationService.GetAdGroupCriterionSimulation].
type GetAdGroupCriterionSimulationRequest struct {
	// Required. The resource name of the ad group criterion simulation to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupCriterionSimulationRequest) Reset()         { *m = GetAdGroupCriterionSimulationRequest{} }
func (m *GetAdGroupCriterionSimulationRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupCriterionSimulationRequest) ProtoMessage()    {}
func (*GetAdGroupCriterionSimulationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da600f48a422603f, []int{0}
}

func (m *GetAdGroupCriterionSimulationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupCriterionSimulationRequest.Unmarshal(m, b)
}
func (m *GetAdGroupCriterionSimulationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupCriterionSimulationRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupCriterionSimulationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupCriterionSimulationRequest.Merge(m, src)
}
func (m *GetAdGroupCriterionSimulationRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupCriterionSimulationRequest.Size(m)
}
func (m *GetAdGroupCriterionSimulationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupCriterionSimulationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupCriterionSimulationRequest proto.InternalMessageInfo

func (m *GetAdGroupCriterionSimulationRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupCriterionSimulationRequest)(nil), "google.ads.googleads.v3.services.GetAdGroupCriterionSimulationRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/ad_group_criterion_simulation_service.proto", fileDescriptor_da600f48a422603f)
}

var fileDescriptor_da600f48a422603f = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xb1, 0x8b, 0xd4, 0x40,
	0x14, 0xc6, 0xd9, 0x1c, 0x08, 0x06, 0x6d, 0xd2, 0x78, 0x44, 0xc5, 0xf5, 0x58, 0xe4, 0xb8, 0x62,
	0x06, 0x4c, 0x37, 0xb2, 0xc8, 0xac, 0x68, 0x2c, 0x44, 0x96, 0x3b, 0xdc, 0x42, 0x02, 0x61, 0x2e,
	0x33, 0xc6, 0x81, 0x24, 0x13, 0xe7, 0x4d, 0xd2, 0x88, 0x8d, 0xbd, 0x95, 0xff, 0x81, 0xa5, 0x7f,
	0x85, 0xf5, 0xb5, 0x76, 0x56, 0x16, 0x56, 0xfe, 0x0f, 0x82, 0x24, 0x93, 0xc9, 0xdd, 0x16, 0x31,
	0x76, 0x1f, 0xf9, 0x3e, 0x7e, 0xdf, 0xcb, 0x7b, 0x89, 0xff, 0x22, 0x57, 0x2a, 0x2f, 0x04, 0x66,
	0x1c, 0xb0, 0x95, 0x9d, 0x6a, 0x23, 0x0c, 0x42, 0xb7, 0x32, 0x13, 0x80, 0x19, 0x4f, 0x73, 0xad,
	0x9a, 0x3a, 0xcd, 0xb4, 0x34, 0x42, 0x4b, 0x55, 0xa5, 0x20, 0xcb, 0xa6, 0x60, 0xa6, 0x97, 0x36,
	0x86, 0x6a, 0xad, 0x8c, 0x0a, 0x96, 0x16, 0x81, 0x18, 0x07, 0x34, 0xd2, 0x50, 0x1b, 0x21, 0x47,
	0x0b, 0x9f, 0x4e, 0xf5, 0x69, 0x01, 0xaa, 0xd1, 0xb3, 0x85, 0xb6, 0x28, 0xbc, 0xe3, 0x30, 0xb5,
	0xc4, 0xac, 0xaa, 0x94, 0xe9, 0x4d, 0x18, 0xdc, 0x5b, 0x57, 0xdc, 0xac, 0x90, 0xa2, 0x32, 0x83,
	0x71, 0xef, 0x8a, 0xf1, 0x46, 0x8a, 0x82, 0xa7, 0xe7, 0xe2, 0x2d, 0x6b, 0xa5, 0xd2, 0x36, 0x70,
	0xb4, 0xf5, 0x57, 0xb1, 0x30, 0x94, 0xc7, 0xdd, 0x00, 0x4f, 0x5c, 0xff, 0xd9, 0x58, 0x7f, 0x2a,
	0xde, 0x35, 0x02, 0x4c, 0x70, 0xec, 0xdf, 0x74, 0x03, 0xa7, 0x15, 0x2b, 0xc5, 0xe1, 0x62, 0xb9,
	0x38, 0xbe, 0xbe, 0x39, 0xf8, 0x49, 0xbd, 0xd3, 0x1b, 0xce, 0x79, 0xc9, 0x4a, 0xf1, 0xf0, 0x9b,
	0xe7, 0xdf, 0x9f, 0xe6, 0x9d, 0xd9, 0xbd, 0x04, 0x7f, 0x16, 0xfe, 0xdd, 0x7f, 0x16, 0x07, 0xcf,
	0xd0, 0xdc, 0x6e, 0xd1, 0xff, 0x4c, 0x1e, 0xae, 0x27, 0x39, 0xe3, 0x05, 0xd0, 0x34, 0xe5, 0xe8,
	0xd5, 0x0f, 0xba, 0xff, 0xe6, 0x1f, 0xbf, 0xff, 0xfa, 0xec, 0x3d, 0x0e, 0xd6, 0xdd, 0x0d, 0xdf,
	0xef, 0x39, 0xeb, 0xac, 0x01, 0xa3, 0x4a, 0xa1, 0x01, 0x9f, 0x60, 0x36, 0x89, 0x04, 0x7c, 0xf2,
	0x21, 0xbc, 0x7d, 0x41, 0x0f, 0x2f, 0x87, 0x19, 0x54, 0x2d, 0x01, 0x65, 0xaa, 0xdc, 0x7c, 0xf2,
	0xfc, 0x55, 0xa6, 0xca, 0xd9, 0x05, 0x6c, 0x1e, 0xcc, 0x2e, 0x7a, 0xdb, 0x5d, 0x79, 0xbb, 0x78,
	0xfd, 0x7c, 0x60, 0xe5, 0xaa, 0x60, 0x55, 0x8e, 0x94, 0xce, 0x71, 0x2e, 0xaa, 0xfe, 0x1b, 0xc0,
	0x97, 0xed, 0xd3, 0x7f, 0xc5, 0x23, 0x27, 0xbe, 0x78, 0x07, 0x31, 0xa5, 0x5f, 0xbd, 0x65, 0x6c,
	0x81, 0x94, 0x03, 0xb2, 0xb2, 0x53, 0xbb, 0x08, 0x0d, 0xc5, 0x70, 0xe1, 0x22, 0x09, 0xe5, 0x90,
	0x8c, 0x91, 0x64, 0x17, 0x25, 0x2e, 0xf2, 0xdb, 0x5b, 0xd9, 0xe7, 0x84, 0x50, 0x0e, 0x84, 0x8c,
	0x21, 0x42, 0x76, 0x11, 0x21, 0x2e, 0x76, 0x7e, 0xad, 0x9f, 0x33, 0xfa, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0x6b, 0x09, 0x38, 0x11, 0xbc, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupCriterionSimulationServiceClient is the client API for AdGroupCriterionSimulationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupCriterionSimulationServiceClient interface {
	// Returns the requested ad group criterion simulation in full detail.
	GetAdGroupCriterionSimulation(ctx context.Context, in *GetAdGroupCriterionSimulationRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterionSimulation, error)
}

type adGroupCriterionSimulationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupCriterionSimulationServiceClient(cc grpc.ClientConnInterface) AdGroupCriterionSimulationServiceClient {
	return &adGroupCriterionSimulationServiceClient{cc}
}

func (c *adGroupCriterionSimulationServiceClient) GetAdGroupCriterionSimulation(ctx context.Context, in *GetAdGroupCriterionSimulationRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterionSimulation, error) {
	out := new(resources.AdGroupCriterionSimulation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.AdGroupCriterionSimulationService/GetAdGroupCriterionSimulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupCriterionSimulationServiceServer is the server API for AdGroupCriterionSimulationService service.
type AdGroupCriterionSimulationServiceServer interface {
	// Returns the requested ad group criterion simulation in full detail.
	GetAdGroupCriterionSimulation(context.Context, *GetAdGroupCriterionSimulationRequest) (*resources.AdGroupCriterionSimulation, error)
}

// UnimplementedAdGroupCriterionSimulationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupCriterionSimulationServiceServer struct {
}

func (*UnimplementedAdGroupCriterionSimulationServiceServer) GetAdGroupCriterionSimulation(ctx context.Context, req *GetAdGroupCriterionSimulationRequest) (*resources.AdGroupCriterionSimulation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdGroupCriterionSimulation not implemented")
}

func RegisterAdGroupCriterionSimulationServiceServer(s *grpc.Server, srv AdGroupCriterionSimulationServiceServer) {
	s.RegisterService(&_AdGroupCriterionSimulationService_serviceDesc, srv)
}

func _AdGroupCriterionSimulationService_GetAdGroupCriterionSimulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupCriterionSimulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionSimulationServiceServer).GetAdGroupCriterionSimulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.AdGroupCriterionSimulationService/GetAdGroupCriterionSimulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionSimulationServiceServer).GetAdGroupCriterionSimulation(ctx, req.(*GetAdGroupCriterionSimulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupCriterionSimulationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.AdGroupCriterionSimulationService",
	HandlerType: (*AdGroupCriterionSimulationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupCriterionSimulation",
			Handler:    _AdGroupCriterionSimulationService_GetAdGroupCriterionSimulation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/ad_group_criterion_simulation_service.proto",
}

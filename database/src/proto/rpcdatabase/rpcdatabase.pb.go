// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpcdatabase.proto

package rpcdatabase

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ClientRequest struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientRequest) Reset()         { *m = ClientRequest{} }
func (m *ClientRequest) String() string { return proto.CompactTextString(m) }
func (*ClientRequest) ProtoMessage()    {}
func (*ClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe2a72ba7550c07e, []int{0}
}

func (m *ClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientRequest.Unmarshal(m, b)
}
func (m *ClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientRequest.Marshal(b, m, deterministic)
}
func (m *ClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientRequest.Merge(m, src)
}
func (m *ClientRequest) XXX_Size() int {
	return xxx_messageInfo_ClientRequest.Size(m)
}
func (m *ClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientRequest proto.InternalMessageInfo

func (m *ClientRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ServerResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerResponse) Reset()         { *m = ServerResponse{} }
func (m *ServerResponse) String() string { return proto.CompactTextString(m) }
func (*ServerResponse) ProtoMessage()    {}
func (*ServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe2a72ba7550c07e, []int{1}
}

func (m *ServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerResponse.Unmarshal(m, b)
}
func (m *ServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerResponse.Marshal(b, m, deterministic)
}
func (m *ServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerResponse.Merge(m, src)
}
func (m *ServerResponse) XXX_Size() int {
	return xxx_messageInfo_ServerResponse.Size(m)
}
func (m *ServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerResponse proto.InternalMessageInfo

func (m *ServerResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ClientStream struct {
	AgentId              int32    `protobuf:"varint,1,opt,name=agentId,proto3" json:"agentId,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStream) Reset()         { *m = ClientStream{} }
func (m *ClientStream) String() string { return proto.CompactTextString(m) }
func (*ClientStream) ProtoMessage()    {}
func (*ClientStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe2a72ba7550c07e, []int{2}
}

func (m *ClientStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStream.Unmarshal(m, b)
}
func (m *ClientStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStream.Marshal(b, m, deterministic)
}
func (m *ClientStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStream.Merge(m, src)
}
func (m *ClientStream) XXX_Size() int {
	return xxx_messageInfo_ClientStream.Size(m)
}
func (m *ClientStream) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStream.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStream proto.InternalMessageInfo

func (m *ClientStream) GetAgentId() int32 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *ClientStream) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ServerStream struct {
	AgentId              int32    `protobuf:"varint,1,opt,name=agentId,proto3" json:"agentId,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerStream) Reset()         { *m = ServerStream{} }
func (m *ServerStream) String() string { return proto.CompactTextString(m) }
func (*ServerStream) ProtoMessage()    {}
func (*ServerStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe2a72ba7550c07e, []int{3}
}

func (m *ServerStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStream.Unmarshal(m, b)
}
func (m *ServerStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStream.Marshal(b, m, deterministic)
}
func (m *ServerStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStream.Merge(m, src)
}
func (m *ServerStream) XXX_Size() int {
	return xxx_messageInfo_ServerStream.Size(m)
}
func (m *ServerStream) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStream.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStream proto.InternalMessageInfo

func (m *ServerStream) GetAgentId() int32 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *ServerStream) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientRequest)(nil), "rpcdatabase.ClientRequest")
	proto.RegisterType((*ServerResponse)(nil), "rpcdatabase.ServerResponse")
	proto.RegisterType((*ClientStream)(nil), "rpcdatabase.ClientStream")
	proto.RegisterType((*ServerStream)(nil), "rpcdatabase.ServerStream")
}

func init() {
	proto.RegisterFile("rpcdatabase.proto", fileDescriptor_fe2a72ba7550c07e)
}

var fileDescriptor_fe2a72ba7550c07e = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2a, 0x48, 0x4e,
	0x49, 0x2c, 0x49, 0x4c, 0x4a, 0x2c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46,
	0x12, 0x52, 0x52, 0xe6, 0xe2, 0x75, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x09, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0x01, 0x49, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x81,
	0xd9, 0x4a, 0x2a, 0x5c, 0x7c, 0xc1, 0xa9, 0x45, 0x65, 0xa9, 0x45, 0x41, 0xa9, 0xc5, 0x05, 0xf9,
	0x79, 0xc5, 0xa9, 0x58, 0x55, 0xd9, 0x70, 0xf1, 0x40, 0x8c, 0x0a, 0x2e, 0x29, 0x4a, 0x4d, 0xcc,
	0x15, 0x92, 0xe0, 0x62, 0x4f, 0x4c, 0x4f, 0xcd, 0x2b, 0xf1, 0x4c, 0x01, 0x2b, 0x63, 0x0d, 0x82,
	0x71, 0xe1, 0xba, 0x99, 0x50, 0x75, 0x43, 0xec, 0x20, 0x47, 0xb7, 0xd1, 0x0c, 0x46, 0x2e, 0x7e,
	0x97, 0xc4, 0x92, 0x44, 0xa7, 0xc4, 0xe2, 0x54, 0x90, 0x31, 0x99, 0xc9, 0xa9, 0x42, 0x8e, 0x5c,
	0x2c, 0xce, 0x89, 0x39, 0x39, 0x42, 0x52, 0x7a, 0xc8, 0x61, 0x80, 0xe2, 0x5b, 0x29, 0x69, 0x14,
	0x39, 0x54, 0x4f, 0x2a, 0x31, 0x08, 0x39, 0x71, 0xb1, 0x04, 0xa7, 0xe6, 0xa5, 0x08, 0x49, 0x62,
	0x31, 0x02, 0xe2, 0x4e, 0x29, 0x49, 0x2c, 0x26, 0x40, 0xa4, 0x94, 0x18, 0x34, 0x18, 0x0d, 0x18,
	0x93, 0xd8, 0xc0, 0xa1, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x31, 0x63, 0x80, 0xbc, 0x8a,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DataBaseServiceClient is the client API for DataBaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataBaseServiceClient interface {
	Call(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ServerResponse, error)
	Send(ctx context.Context, opts ...grpc.CallOption) (DataBaseService_SendClient, error)
}

type dataBaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataBaseServiceClient(cc grpc.ClientConnInterface) DataBaseServiceClient {
	return &dataBaseServiceClient{cc}
}

func (c *dataBaseServiceClient) Call(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/rpcdatabase.DataBaseService/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataBaseServiceClient) Send(ctx context.Context, opts ...grpc.CallOption) (DataBaseService_SendClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataBaseService_serviceDesc.Streams[0], "/rpcdatabase.DataBaseService/Send", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataBaseServiceSendClient{stream}
	return x, nil
}

type DataBaseService_SendClient interface {
	Send(*ClientStream) error
	Recv() (*ServerStream, error)
	grpc.ClientStream
}

type dataBaseServiceSendClient struct {
	grpc.ClientStream
}

func (x *dataBaseServiceSendClient) Send(m *ClientStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataBaseServiceSendClient) Recv() (*ServerStream, error) {
	m := new(ServerStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataBaseServiceServer is the server API for DataBaseService service.
type DataBaseServiceServer interface {
	Call(context.Context, *ClientRequest) (*ServerResponse, error)
	Send(DataBaseService_SendServer) error
}

// UnimplementedDataBaseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDataBaseServiceServer struct {
}

func (*UnimplementedDataBaseServiceServer) Call(ctx context.Context, req *ClientRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (*UnimplementedDataBaseServiceServer) Send(srv DataBaseService_SendServer) error {
	return status.Errorf(codes.Unimplemented, "method Send not implemented")
}

func RegisterDataBaseServiceServer(s *grpc.Server, srv DataBaseServiceServer) {
	s.RegisterService(&_DataBaseService_serviceDesc, srv)
}

func _DataBaseService_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBaseServiceServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcdatabase.DataBaseService/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBaseServiceServer).Call(ctx, req.(*ClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataBaseService_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataBaseServiceServer).Send(&dataBaseServiceSendServer{stream})
}

type DataBaseService_SendServer interface {
	Send(*ServerStream) error
	Recv() (*ClientStream, error)
	grpc.ServerStream
}

type dataBaseServiceSendServer struct {
	grpc.ServerStream
}

func (x *dataBaseServiceSendServer) Send(m *ServerStream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataBaseServiceSendServer) Recv() (*ClientStream, error) {
	m := new(ClientStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DataBaseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcdatabase.DataBaseService",
	HandlerType: (*DataBaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _DataBaseService_Call_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Send",
			Handler:       _DataBaseService_Send_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rpcdatabase.proto",
}

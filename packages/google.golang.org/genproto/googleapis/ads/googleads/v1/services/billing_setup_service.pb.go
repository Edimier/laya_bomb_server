// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/billing_setup_service.proto

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

// Request message for
// [BillingSetupService.GetBillingSetup][google.ads.googleads.v1.services.BillingSetupService.GetBillingSetup].
type GetBillingSetupRequest struct {
	// Required. The resource name of the billing setup to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBillingSetupRequest) Reset()         { *m = GetBillingSetupRequest{} }
func (m *GetBillingSetupRequest) String() string { return proto.CompactTextString(m) }
func (*GetBillingSetupRequest) ProtoMessage()    {}
func (*GetBillingSetupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a012b9d34a4d410, []int{0}
}

func (m *GetBillingSetupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBillingSetupRequest.Unmarshal(m, b)
}
func (m *GetBillingSetupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBillingSetupRequest.Marshal(b, m, deterministic)
}
func (m *GetBillingSetupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBillingSetupRequest.Merge(m, src)
}
func (m *GetBillingSetupRequest) XXX_Size() int {
	return xxx_messageInfo_GetBillingSetupRequest.Size(m)
}
func (m *GetBillingSetupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBillingSetupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBillingSetupRequest proto.InternalMessageInfo

func (m *GetBillingSetupRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for billing setup mutate operations.
type MutateBillingSetupRequest struct {
	// Required. Id of the customer to apply the billing setup mutate operation to.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The operation to perform.
	Operation            *BillingSetupOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MutateBillingSetupRequest) Reset()         { *m = MutateBillingSetupRequest{} }
func (m *MutateBillingSetupRequest) String() string { return proto.CompactTextString(m) }
func (*MutateBillingSetupRequest) ProtoMessage()    {}
func (*MutateBillingSetupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a012b9d34a4d410, []int{1}
}

func (m *MutateBillingSetupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateBillingSetupRequest.Unmarshal(m, b)
}
func (m *MutateBillingSetupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateBillingSetupRequest.Marshal(b, m, deterministic)
}
func (m *MutateBillingSetupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateBillingSetupRequest.Merge(m, src)
}
func (m *MutateBillingSetupRequest) XXX_Size() int {
	return xxx_messageInfo_MutateBillingSetupRequest.Size(m)
}
func (m *MutateBillingSetupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateBillingSetupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateBillingSetupRequest proto.InternalMessageInfo

func (m *MutateBillingSetupRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateBillingSetupRequest) GetOperation() *BillingSetupOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

// A single operation on a billing setup, which describes the cancellation of an
// existing billing setup.
type BillingSetupOperation struct {
	// Only one of these operations can be set. "Update" operations are not
	// supported.
	//
	// Types that are valid to be assigned to Operation:
	//	*BillingSetupOperation_Create
	//	*BillingSetupOperation_Remove
	Operation            isBillingSetupOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *BillingSetupOperation) Reset()         { *m = BillingSetupOperation{} }
func (m *BillingSetupOperation) String() string { return proto.CompactTextString(m) }
func (*BillingSetupOperation) ProtoMessage()    {}
func (*BillingSetupOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a012b9d34a4d410, []int{2}
}

func (m *BillingSetupOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillingSetupOperation.Unmarshal(m, b)
}
func (m *BillingSetupOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillingSetupOperation.Marshal(b, m, deterministic)
}
func (m *BillingSetupOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillingSetupOperation.Merge(m, src)
}
func (m *BillingSetupOperation) XXX_Size() int {
	return xxx_messageInfo_BillingSetupOperation.Size(m)
}
func (m *BillingSetupOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_BillingSetupOperation.DiscardUnknown(m)
}

var xxx_messageInfo_BillingSetupOperation proto.InternalMessageInfo

type isBillingSetupOperation_Operation interface {
	isBillingSetupOperation_Operation()
}

type BillingSetupOperation_Create struct {
	Create *resources.BillingSetup `protobuf:"bytes,2,opt,name=create,proto3,oneof"`
}

type BillingSetupOperation_Remove struct {
	Remove string `protobuf:"bytes,1,opt,name=remove,proto3,oneof"`
}

func (*BillingSetupOperation_Create) isBillingSetupOperation_Operation() {}

func (*BillingSetupOperation_Remove) isBillingSetupOperation_Operation() {}

func (m *BillingSetupOperation) GetOperation() isBillingSetupOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *BillingSetupOperation) GetCreate() *resources.BillingSetup {
	if x, ok := m.GetOperation().(*BillingSetupOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *BillingSetupOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*BillingSetupOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BillingSetupOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BillingSetupOperation_Create)(nil),
		(*BillingSetupOperation_Remove)(nil),
	}
}

// Response message for a billing setup operation.
type MutateBillingSetupResponse struct {
	// A result that identifies the resource affected by the mutate request.
	Result               *MutateBillingSetupResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MutateBillingSetupResponse) Reset()         { *m = MutateBillingSetupResponse{} }
func (m *MutateBillingSetupResponse) String() string { return proto.CompactTextString(m) }
func (*MutateBillingSetupResponse) ProtoMessage()    {}
func (*MutateBillingSetupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a012b9d34a4d410, []int{3}
}

func (m *MutateBillingSetupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateBillingSetupResponse.Unmarshal(m, b)
}
func (m *MutateBillingSetupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateBillingSetupResponse.Marshal(b, m, deterministic)
}
func (m *MutateBillingSetupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateBillingSetupResponse.Merge(m, src)
}
func (m *MutateBillingSetupResponse) XXX_Size() int {
	return xxx_messageInfo_MutateBillingSetupResponse.Size(m)
}
func (m *MutateBillingSetupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateBillingSetupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateBillingSetupResponse proto.InternalMessageInfo

func (m *MutateBillingSetupResponse) GetResult() *MutateBillingSetupResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// Result for a single billing setup mutate.
type MutateBillingSetupResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateBillingSetupResult) Reset()         { *m = MutateBillingSetupResult{} }
func (m *MutateBillingSetupResult) String() string { return proto.CompactTextString(m) }
func (*MutateBillingSetupResult) ProtoMessage()    {}
func (*MutateBillingSetupResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a012b9d34a4d410, []int{4}
}

func (m *MutateBillingSetupResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateBillingSetupResult.Unmarshal(m, b)
}
func (m *MutateBillingSetupResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateBillingSetupResult.Marshal(b, m, deterministic)
}
func (m *MutateBillingSetupResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateBillingSetupResult.Merge(m, src)
}
func (m *MutateBillingSetupResult) XXX_Size() int {
	return xxx_messageInfo_MutateBillingSetupResult.Size(m)
}
func (m *MutateBillingSetupResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateBillingSetupResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateBillingSetupResult proto.InternalMessageInfo

func (m *MutateBillingSetupResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBillingSetupRequest)(nil), "google.ads.googleads.v1.services.GetBillingSetupRequest")
	proto.RegisterType((*MutateBillingSetupRequest)(nil), "google.ads.googleads.v1.services.MutateBillingSetupRequest")
	proto.RegisterType((*BillingSetupOperation)(nil), "google.ads.googleads.v1.services.BillingSetupOperation")
	proto.RegisterType((*MutateBillingSetupResponse)(nil), "google.ads.googleads.v1.services.MutateBillingSetupResponse")
	proto.RegisterType((*MutateBillingSetupResult)(nil), "google.ads.googleads.v1.services.MutateBillingSetupResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/billing_setup_service.proto", fileDescriptor_0a012b9d34a4d410)
}

var fileDescriptor_0a012b9d34a4d410 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0x36, 0x29, 0x14, 0x76, 0xaa, 0x08, 0x23, 0xab, 0xb1, 0x0a, 0x96, 0xd8, 0x43, 0x29, 0x92,
	0x21, 0x15, 0x59, 0x99, 0xdd, 0x45, 0x92, 0x4b, 0x77, 0x0f, 0xea, 0xd2, 0xc5, 0x1e, 0x96, 0x42,
	0x49, 0x93, 0x31, 0x06, 0x92, 0x4c, 0xcc, 0x4c, 0x7a, 0x59, 0xf6, 0xe2, 0xcd, 0xb3, 0x67, 0x2f,
	0x1e, 0xfd, 0x3f, 0xbc, 0xec, 0xd5, 0xdb, 0x9e, 0x04, 0xc5, 0x83, 0xff, 0x84, 0x92, 0x4c, 0xa6,
	0x4d, 0xdd, 0x94, 0xb2, 0x7b, 0x7b, 0x33, 0xef, 0x7b, 0xdf, 0xf7, 0x7e, 0xcd, 0x80, 0x3d, 0x9f,
	0x52, 0x3f, 0x24, 0xc8, 0xf1, 0x18, 0x12, 0x66, 0x6e, 0xcd, 0x4d, 0xc4, 0x48, 0x3a, 0x0f, 0x5c,
	0xc2, 0xd0, 0x2c, 0x08, 0xc3, 0x20, 0xf6, 0xa7, 0x8c, 0xf0, 0x2c, 0x99, 0x96, 0xd7, 0x46, 0x92,
	0x52, 0x4e, 0x61, 0x47, 0x84, 0x18, 0x8e, 0xc7, 0x8c, 0x45, 0xb4, 0x31, 0x37, 0x0d, 0x19, 0xdd,
	0x7e, 0xb6, 0x8e, 0x3f, 0x25, 0x8c, 0x66, 0xe9, 0x25, 0x01, 0x41, 0xdc, 0x7e, 0x28, 0xc3, 0x92,
	0x00, 0x39, 0x71, 0x4c, 0xb9, 0xc3, 0x03, 0x1a, 0xb3, 0xd2, 0x7b, 0xaf, 0xe2, 0x75, 0xc3, 0x80,
	0xc4, 0xbc, 0x74, 0x3c, 0xaa, 0x38, 0xde, 0x06, 0x24, 0xf4, 0xa6, 0x33, 0xf2, 0xce, 0x99, 0x07,
	0x34, 0x15, 0x00, 0xdd, 0x06, 0x77, 0x87, 0x84, 0xdb, 0x42, 0xf1, 0x38, 0x17, 0x1c, 0x91, 0xf7,
	0x19, 0x61, 0x1c, 0xf6, 0xc0, 0x2d, 0x99, 0xd2, 0x34, 0x76, 0x22, 0xa2, 0x29, 0x1d, 0xa5, 0xb7,
	0x65, 0x37, 0x7e, 0x58, 0xea, 0xe8, 0xa6, 0xf4, 0xbc, 0x72, 0x22, 0xa2, 0x7f, 0x56, 0xc0, 0xfd,
	0x97, 0x19, 0x77, 0x38, 0xa9, 0xe3, 0xe9, 0x82, 0x96, 0x9b, 0x31, 0x4e, 0x23, 0x92, 0x4e, 0x03,
	0xaf, 0xca, 0x02, 0xe4, 0xfd, 0xa1, 0x07, 0x4f, 0xc0, 0x16, 0x4d, 0x48, 0x5a, 0x54, 0xa5, 0xa9,
	0x1d, 0xa5, 0xd7, 0x1a, 0xec, 0x18, 0x9b, 0x9a, 0x69, 0x54, 0xf5, 0x5e, 0xcb, 0x70, 0x41, 0xbe,
	0xa4, 0xd3, 0x3f, 0x2a, 0x60, 0xbb, 0x16, 0x09, 0x0f, 0x41, 0xd3, 0x4d, 0x89, 0xc3, 0x49, 0x29,
	0x89, 0xd6, 0x4a, 0x2e, 0xa6, 0xb3, 0xa2, 0x79, 0x70, 0x63, 0x54, 0x12, 0x40, 0x0d, 0x34, 0x53,
	0x12, 0xd1, 0x79, 0xd9, 0xa7, 0xdc, 0x23, 0xce, 0x76, 0xab, 0x52, 0x9a, 0x9e, 0x80, 0x76, 0x5d,
	0xab, 0x58, 0x42, 0x63, 0x46, 0xe0, 0x28, 0x27, 0x61, 0x59, 0xc8, 0x0b, 0x92, 0xd6, 0x00, 0x6f,
	0x6e, 0x41, 0x2d, 0x5b, 0x16, 0xf2, 0x51, 0xc9, 0xa4, 0xbf, 0x00, 0xda, 0x3a, 0x0c, 0x7c, 0x5c,
	0x3b, 0xe3, 0xd5, 0xf1, 0x0e, 0x7e, 0x36, 0xc0, 0x9d, 0x6a, 0xec, 0xb1, 0x90, 0x86, 0xdf, 0x14,
	0x70, 0xfb, 0xbf, 0xdd, 0x81, 0xcf, 0x37, 0x27, 0x5c, 0xbf, 0x6e, 0xed, 0xab, 0xb6, 0x5e, 0x1f,
	0x5e, 0x58, 0xab, 0xc9, 0x7f, 0xf8, 0xfe, 0xeb, 0x93, 0x6a, 0x42, 0x94, 0x3f, 0xa6, 0xd3, 0x15,
	0xcf, 0xbe, 0x5c, 0x34, 0x86, 0xfa, 0xf2, 0x75, 0x15, 0x24, 0x0c, 0xf5, 0xcf, 0xe0, 0x6f, 0x05,
	0xc0, 0xcb, 0x1d, 0x82, 0xbb, 0xd7, 0xeb, 0xbd, 0xa8, 0x66, 0xef, 0x9a, 0x83, 0x2b, 0xd6, 0x40,
	0x7f, 0x73, 0x61, 0x6d, 0x57, 0xde, 0xcc, 0x93, 0xc5, 0xfa, 0x14, 0x25, 0xee, 0xe8, 0x83, 0xbc,
	0xc4, 0x65, 0x4d, 0xa7, 0x15, 0xec, 0x7e, 0xff, 0x6c, 0xb5, 0x42, 0x1c, 0x15, 0x3a, 0x58, 0xe9,
	0xb7, 0x1f, 0x9c, 0x5b, 0xda, 0x32, 0x97, 0xd2, 0x4a, 0x02, 0x66, 0xb8, 0x34, 0xb2, 0xff, 0x2a,
	0xa0, 0xeb, 0xd2, 0x68, 0x63, 0xde, 0xb6, 0x56, 0xb3, 0x0b, 0x47, 0xf9, 0x5f, 0x72, 0xa4, 0x9c,
	0x1c, 0x94, 0xd1, 0x3e, 0x0d, 0x9d, 0xd8, 0x37, 0x68, 0xea, 0x23, 0x9f, 0xc4, 0xc5, 0x4f, 0x83,
	0x96, 0x7a, 0xeb, 0xff, 0xd6, 0x5d, 0x69, 0x7c, 0x51, 0x1b, 0x43, 0xcb, 0xfa, 0xaa, 0x76, 0x86,
	0x82, 0xd0, 0xf2, 0x98, 0x21, 0xcc, 0xdc, 0x1a, 0x9b, 0x46, 0x29, 0xcc, 0xce, 0x25, 0x64, 0x62,
	0x79, 0x6c, 0xb2, 0x80, 0x4c, 0xc6, 0xe6, 0x44, 0x42, 0xfe, 0xa8, 0x5d, 0x71, 0x8f, 0xb1, 0xe5,
	0x31, 0x8c, 0x17, 0x20, 0x8c, 0xc7, 0x26, 0xc6, 0x12, 0x36, 0x6b, 0x16, 0x79, 0x3e, 0xfd, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x37, 0x84, 0x5c, 0x47, 0x02, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BillingSetupServiceClient is the client API for BillingSetupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BillingSetupServiceClient interface {
	// Returns a billing setup.
	GetBillingSetup(ctx context.Context, in *GetBillingSetupRequest, opts ...grpc.CallOption) (*resources.BillingSetup, error)
	// Creates a billing setup, or cancels an existing billing setup.
	MutateBillingSetup(ctx context.Context, in *MutateBillingSetupRequest, opts ...grpc.CallOption) (*MutateBillingSetupResponse, error)
}

type billingSetupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBillingSetupServiceClient(cc grpc.ClientConnInterface) BillingSetupServiceClient {
	return &billingSetupServiceClient{cc}
}

func (c *billingSetupServiceClient) GetBillingSetup(ctx context.Context, in *GetBillingSetupRequest, opts ...grpc.CallOption) (*resources.BillingSetup, error) {
	out := new(resources.BillingSetup)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.BillingSetupService/GetBillingSetup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingSetupServiceClient) MutateBillingSetup(ctx context.Context, in *MutateBillingSetupRequest, opts ...grpc.CallOption) (*MutateBillingSetupResponse, error) {
	out := new(MutateBillingSetupResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.BillingSetupService/MutateBillingSetup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingSetupServiceServer is the server API for BillingSetupService service.
type BillingSetupServiceServer interface {
	// Returns a billing setup.
	GetBillingSetup(context.Context, *GetBillingSetupRequest) (*resources.BillingSetup, error)
	// Creates a billing setup, or cancels an existing billing setup.
	MutateBillingSetup(context.Context, *MutateBillingSetupRequest) (*MutateBillingSetupResponse, error)
}

// UnimplementedBillingSetupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBillingSetupServiceServer struct {
}

func (*UnimplementedBillingSetupServiceServer) GetBillingSetup(ctx context.Context, req *GetBillingSetupRequest) (*resources.BillingSetup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBillingSetup not implemented")
}
func (*UnimplementedBillingSetupServiceServer) MutateBillingSetup(ctx context.Context, req *MutateBillingSetupRequest) (*MutateBillingSetupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateBillingSetup not implemented")
}

func RegisterBillingSetupServiceServer(s *grpc.Server, srv BillingSetupServiceServer) {
	s.RegisterService(&_BillingSetupService_serviceDesc, srv)
}

func _BillingSetupService_GetBillingSetup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBillingSetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingSetupServiceServer).GetBillingSetup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.BillingSetupService/GetBillingSetup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingSetupServiceServer).GetBillingSetup(ctx, req.(*GetBillingSetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingSetupService_MutateBillingSetup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateBillingSetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingSetupServiceServer).MutateBillingSetup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.BillingSetupService/MutateBillingSetup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingSetupServiceServer).MutateBillingSetup(ctx, req.(*MutateBillingSetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BillingSetupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.BillingSetupService",
	HandlerType: (*BillingSetupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBillingSetup",
			Handler:    _BillingSetupService_GetBillingSetup_Handler,
		},
		{
			MethodName: "MutateBillingSetup",
			Handler:    _BillingSetupService_MutateBillingSetup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/billing_setup_service.proto",
}

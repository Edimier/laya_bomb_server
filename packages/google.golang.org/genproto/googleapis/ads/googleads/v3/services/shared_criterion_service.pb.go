// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/shared_criterion_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [SharedCriterionService.GetSharedCriterion][google.ads.googleads.v3.services.SharedCriterionService.GetSharedCriterion].
type GetSharedCriterionRequest struct {
	// Required. The resource name of the shared criterion to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSharedCriterionRequest) Reset()         { *m = GetSharedCriterionRequest{} }
func (m *GetSharedCriterionRequest) String() string { return proto.CompactTextString(m) }
func (*GetSharedCriterionRequest) ProtoMessage()    {}
func (*GetSharedCriterionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_759802b7b62c0df4, []int{0}
}

func (m *GetSharedCriterionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSharedCriterionRequest.Unmarshal(m, b)
}
func (m *GetSharedCriterionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSharedCriterionRequest.Marshal(b, m, deterministic)
}
func (m *GetSharedCriterionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSharedCriterionRequest.Merge(m, src)
}
func (m *GetSharedCriterionRequest) XXX_Size() int {
	return xxx_messageInfo_GetSharedCriterionRequest.Size(m)
}
func (m *GetSharedCriterionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSharedCriterionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSharedCriterionRequest proto.InternalMessageInfo

func (m *GetSharedCriterionRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [SharedCriterionService.MutateSharedCriteria][google.ads.googleads.v3.services.SharedCriterionService.MutateSharedCriteria].
type MutateSharedCriteriaRequest struct {
	// Required. The ID of the customer whose shared criteria are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual shared criteria.
	Operations []*SharedCriterionOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateSharedCriteriaRequest) Reset()         { *m = MutateSharedCriteriaRequest{} }
func (m *MutateSharedCriteriaRequest) String() string { return proto.CompactTextString(m) }
func (*MutateSharedCriteriaRequest) ProtoMessage()    {}
func (*MutateSharedCriteriaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_759802b7b62c0df4, []int{1}
}

func (m *MutateSharedCriteriaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateSharedCriteriaRequest.Unmarshal(m, b)
}
func (m *MutateSharedCriteriaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateSharedCriteriaRequest.Marshal(b, m, deterministic)
}
func (m *MutateSharedCriteriaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateSharedCriteriaRequest.Merge(m, src)
}
func (m *MutateSharedCriteriaRequest) XXX_Size() int {
	return xxx_messageInfo_MutateSharedCriteriaRequest.Size(m)
}
func (m *MutateSharedCriteriaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateSharedCriteriaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateSharedCriteriaRequest proto.InternalMessageInfo

func (m *MutateSharedCriteriaRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateSharedCriteriaRequest) GetOperations() []*SharedCriterionOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateSharedCriteriaRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateSharedCriteriaRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on an shared criterion.
type SharedCriterionOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*SharedCriterionOperation_Create
	//	*SharedCriterionOperation_Remove
	Operation            isSharedCriterionOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *SharedCriterionOperation) Reset()         { *m = SharedCriterionOperation{} }
func (m *SharedCriterionOperation) String() string { return proto.CompactTextString(m) }
func (*SharedCriterionOperation) ProtoMessage()    {}
func (*SharedCriterionOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_759802b7b62c0df4, []int{2}
}

func (m *SharedCriterionOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedCriterionOperation.Unmarshal(m, b)
}
func (m *SharedCriterionOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedCriterionOperation.Marshal(b, m, deterministic)
}
func (m *SharedCriterionOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedCriterionOperation.Merge(m, src)
}
func (m *SharedCriterionOperation) XXX_Size() int {
	return xxx_messageInfo_SharedCriterionOperation.Size(m)
}
func (m *SharedCriterionOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedCriterionOperation.DiscardUnknown(m)
}

var xxx_messageInfo_SharedCriterionOperation proto.InternalMessageInfo

type isSharedCriterionOperation_Operation interface {
	isSharedCriterionOperation_Operation()
}

type SharedCriterionOperation_Create struct {
	Create *resources.SharedCriterion `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type SharedCriterionOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*SharedCriterionOperation_Create) isSharedCriterionOperation_Operation() {}

func (*SharedCriterionOperation_Remove) isSharedCriterionOperation_Operation() {}

func (m *SharedCriterionOperation) GetOperation() isSharedCriterionOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *SharedCriterionOperation) GetCreate() *resources.SharedCriterion {
	if x, ok := m.GetOperation().(*SharedCriterionOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *SharedCriterionOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*SharedCriterionOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SharedCriterionOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SharedCriterionOperation_Create)(nil),
		(*SharedCriterionOperation_Remove)(nil),
	}
}

// Response message for a shared criterion mutate.
type MutateSharedCriteriaResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateSharedCriterionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MutateSharedCriteriaResponse) Reset()         { *m = MutateSharedCriteriaResponse{} }
func (m *MutateSharedCriteriaResponse) String() string { return proto.CompactTextString(m) }
func (*MutateSharedCriteriaResponse) ProtoMessage()    {}
func (*MutateSharedCriteriaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_759802b7b62c0df4, []int{3}
}

func (m *MutateSharedCriteriaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateSharedCriteriaResponse.Unmarshal(m, b)
}
func (m *MutateSharedCriteriaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateSharedCriteriaResponse.Marshal(b, m, deterministic)
}
func (m *MutateSharedCriteriaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateSharedCriteriaResponse.Merge(m, src)
}
func (m *MutateSharedCriteriaResponse) XXX_Size() int {
	return xxx_messageInfo_MutateSharedCriteriaResponse.Size(m)
}
func (m *MutateSharedCriteriaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateSharedCriteriaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateSharedCriteriaResponse proto.InternalMessageInfo

func (m *MutateSharedCriteriaResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateSharedCriteriaResponse) GetResults() []*MutateSharedCriterionResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the shared criterion mutate.
type MutateSharedCriterionResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateSharedCriterionResult) Reset()         { *m = MutateSharedCriterionResult{} }
func (m *MutateSharedCriterionResult) String() string { return proto.CompactTextString(m) }
func (*MutateSharedCriterionResult) ProtoMessage()    {}
func (*MutateSharedCriterionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_759802b7b62c0df4, []int{4}
}

func (m *MutateSharedCriterionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateSharedCriterionResult.Unmarshal(m, b)
}
func (m *MutateSharedCriterionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateSharedCriterionResult.Marshal(b, m, deterministic)
}
func (m *MutateSharedCriterionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateSharedCriterionResult.Merge(m, src)
}
func (m *MutateSharedCriterionResult) XXX_Size() int {
	return xxx_messageInfo_MutateSharedCriterionResult.Size(m)
}
func (m *MutateSharedCriterionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateSharedCriterionResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateSharedCriterionResult proto.InternalMessageInfo

func (m *MutateSharedCriterionResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSharedCriterionRequest)(nil), "google.ads.googleads.v3.services.GetSharedCriterionRequest")
	proto.RegisterType((*MutateSharedCriteriaRequest)(nil), "google.ads.googleads.v3.services.MutateSharedCriteriaRequest")
	proto.RegisterType((*SharedCriterionOperation)(nil), "google.ads.googleads.v3.services.SharedCriterionOperation")
	proto.RegisterType((*MutateSharedCriteriaResponse)(nil), "google.ads.googleads.v3.services.MutateSharedCriteriaResponse")
	proto.RegisterType((*MutateSharedCriterionResult)(nil), "google.ads.googleads.v3.services.MutateSharedCriterionResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/shared_criterion_service.proto", fileDescriptor_759802b7b62c0df4)
}

var fileDescriptor_759802b7b62c0df4 = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0x37, 0x59, 0xa9, 0x76, 0xb6, 0x55, 0x18, 0xb5, 0xc6, 0x6d, 0xc1, 0x25, 0x16, 0x5c, 0x16,
	0x49, 0x64, 0x73, 0x29, 0x29, 0x55, 0xb2, 0xd2, 0x3f, 0x82, 0xda, 0xb2, 0x85, 0x16, 0x64, 0x21,
	0x4c, 0x93, 0xe9, 0x36, 0x90, 0x64, 0xe2, 0xcc, 0x64, 0xa1, 0x94, 0x5e, 0xfa, 0x05, 0x3c, 0xf8,
	0x0d, 0x3c, 0xfa, 0x25, 0xbc, 0xf7, 0xe2, 0xc1, 0x5b, 0x4f, 0x1e, 0x3c, 0x79, 0xf0, 0xe0, 0xc9,
	0xa3, 0x24, 0x93, 0xd9, 0x7f, 0x66, 0x59, 0xe8, 0xed, 0x65, 0xde, 0x7b, 0xbf, 0xf7, 0x7b, 0xf3,
	0x7e, 0x6f, 0x02, 0x5e, 0xf6, 0x08, 0xe9, 0x85, 0xd8, 0x44, 0x3e, 0x33, 0x85, 0x99, 0x59, 0x7d,
	0xcb, 0x64, 0x98, 0xf6, 0x03, 0x0f, 0x33, 0x93, 0x9d, 0x20, 0x8a, 0x7d, 0xd7, 0xa3, 0x01, 0xc7,
	0x34, 0x20, 0xb1, 0x5b, 0x78, 0x8c, 0x84, 0x12, 0x4e, 0x60, 0x5d, 0x64, 0x19, 0xc8, 0x67, 0xc6,
	0x00, 0xc0, 0xe8, 0x5b, 0x86, 0x04, 0xa8, 0xad, 0x4d, 0x2b, 0x41, 0x31, 0x23, 0x29, 0x2d, 0xab,
	0x21, 0xb0, 0x6b, 0x2b, 0x32, 0x33, 0x09, 0x4c, 0x14, 0xc7, 0x84, 0x23, 0x1e, 0x90, 0x98, 0x15,
	0xde, 0x87, 0x23, 0x5e, 0x2f, 0x0c, 0x70, 0xcc, 0x0b, 0xc7, 0xe3, 0x11, 0xc7, 0x71, 0x80, 0x43,
	0xdf, 0x3d, 0xc2, 0x27, 0xa8, 0x1f, 0x10, 0x3a, 0x91, 0x49, 0x13, 0xcf, 0x64, 0x1c, 0xf1, 0xb4,
	0x80, 0xd4, 0x37, 0xc1, 0xa3, 0x6d, 0xcc, 0xf7, 0x73, 0x36, 0xaf, 0x24, 0x99, 0x0e, 0xfe, 0x90,
	0x62, 0xc6, 0x61, 0x03, 0x2c, 0x4a, 0xc6, 0x6e, 0x8c, 0x22, 0xac, 0x29, 0x75, 0xa5, 0x31, 0xdf,
	0xae, 0xfc, 0x70, 0xd4, 0xce, 0x82, 0xf4, 0xbc, 0x43, 0x11, 0xd6, 0xff, 0x28, 0x60, 0xf9, 0x6d,
	0xca, 0x11, 0xc7, 0x63, 0x50, 0x48, 0x22, 0xad, 0x82, 0xaa, 0x97, 0x32, 0x4e, 0x22, 0x4c, 0xdd,
	0xc0, 0x1f, 0xc5, 0x01, 0xf2, 0xfc, 0xb5, 0x0f, 0x5d, 0x00, 0x48, 0x82, 0xa9, 0xe8, 0x59, 0x53,
	0xeb, 0x95, 0x46, 0xb5, 0x65, 0x1b, 0xb3, 0xae, 0xdb, 0x98, 0x60, 0xbf, 0x2b, 0x21, 0x8a, 0x02,
	0x43, 0x48, 0xf8, 0x14, 0xdc, 0x4d, 0x10, 0xe5, 0x01, 0x0a, 0xdd, 0x63, 0x14, 0x84, 0x29, 0xc5,
	0x5a, 0xa5, 0xae, 0x34, 0x6e, 0x77, 0xee, 0x14, 0xc7, 0x5b, 0xe2, 0x14, 0x3e, 0x01, 0x8b, 0x7d,
	0x14, 0x06, 0x3e, 0xe2, 0xd8, 0x25, 0x71, 0x78, 0xaa, 0xdd, 0xcc, 0xc3, 0x16, 0xe4, 0xe1, 0x6e,
	0x1c, 0x9e, 0xea, 0x1f, 0x15, 0xa0, 0x4d, 0xab, 0x0d, 0xdf, 0x80, 0x39, 0x8f, 0x62, 0xc4, 0xc5,
	0xa5, 0x55, 0x5b, 0xad, 0xa9, 0x7d, 0x0c, 0x44, 0x31, 0xd9, 0xc8, 0xce, 0x8d, 0x4e, 0x81, 0x01,
	0x35, 0x30, 0x47, 0x71, 0x44, 0xfa, 0x82, 0xef, 0x7c, 0xe6, 0x11, 0xdf, 0xed, 0x2a, 0x98, 0x1f,
	0x34, 0xa8, 0x7f, 0x55, 0xc0, 0x4a, 0xf9, 0x18, 0x58, 0x42, 0x62, 0x86, 0xe1, 0x16, 0x78, 0x30,
	0x71, 0x01, 0x2e, 0xa6, 0x94, 0xd0, 0x1c, 0xb6, 0xda, 0x82, 0x92, 0x24, 0x4d, 0x3c, 0x63, 0x3f,
	0xd7, 0x49, 0xe7, 0xde, 0xf8, 0xd5, 0x6c, 0x66, 0xe1, 0xf0, 0x10, 0xdc, 0xa2, 0x98, 0xa5, 0x21,
	0x97, 0x63, 0xda, 0x98, 0x3d, 0xa6, 0x12, 0x62, 0x99, 0xd4, 0x32, 0x94, 0x8e, 0x44, 0xd3, 0xdb,
	0xa5, 0x3a, 0x92, 0x71, 0xd9, 0x5c, 0x4a, 0x14, 0x39, 0x2e, 0xc6, 0xd6, 0xdf, 0x0a, 0x58, 0x9a,
	0x48, 0xdf, 0x17, 0x24, 0xe0, 0x37, 0x05, 0xc0, 0xff, 0xf5, 0x0e, 0xd7, 0x67, 0xb3, 0x9f, 0xba,
	0x25, 0xb5, 0x6b, 0x4c, 0x56, 0xdf, 0xb9, 0x72, 0xc6, 0x1b, 0xb9, 0xf8, 0xfe, 0xf3, 0x93, 0xda,
	0x82, 0xcf, 0xb3, 0x57, 0xe2, 0x6c, 0xcc, 0xb3, 0x21, 0x17, 0x84, 0x99, 0xcd, 0xe2, 0xd9, 0x90,
	0x63, 0x35, 0x9b, 0xe7, 0xf0, 0xb7, 0x02, 0xee, 0x97, 0x8d, 0x1c, 0x5e, 0x6f, 0x22, 0x72, 0x63,
	0x6b, 0x2f, 0xae, 0x9b, 0x2e, 0x94, 0xa6, 0x1f, 0x5e, 0x39, 0x4b, 0x23, 0x2b, 0xff, 0x6c, 0xb8,
	0x85, 0x79, 0xab, 0x6b, 0xba, 0x95, 0xb5, 0x3a, 0xec, 0xed, 0x6c, 0x24, 0x78, 0xa3, 0x79, 0x3e,
	0xd1, 0xa9, 0x1d, 0xe5, 0xb5, 0x6c, 0xa5, 0x59, 0x5b, 0xbe, 0x74, 0xb4, 0x21, 0x9f, 0xc2, 0x4a,
	0x02, 0x66, 0x78, 0x24, 0x6a, 0x5f, 0xa8, 0x60, 0xd5, 0x23, 0xd1, 0x4c, 0xee, 0xed, 0xe5, 0x72,
	0x81, 0xec, 0x65, 0x8f, 0xe2, 0x9e, 0xf2, 0x7e, 0xa7, 0x00, 0xe8, 0x91, 0x10, 0xc5, 0x3d, 0x83,
	0xd0, 0x9e, 0xd9, 0xc3, 0x71, 0xfe, 0x64, 0x9a, 0xc3, 0x92, 0xd3, 0xff, 0x21, 0xeb, 0xd2, 0xf8,
	0xac, 0x56, 0xb6, 0x1d, 0xe7, 0x8b, 0x5a, 0xdf, 0x16, 0x80, 0x8e, 0xcf, 0x0c, 0x61, 0x66, 0xd6,
	0x81, 0x65, 0x14, 0x85, 0xd9, 0xa5, 0x0c, 0xe9, 0x3a, 0x3e, 0xeb, 0x0e, 0x42, 0xba, 0x07, 0x56,
	0x57, 0x86, 0xfc, 0x52, 0x57, 0xc5, 0xb9, 0x6d, 0x3b, 0x3e, 0xb3, 0xed, 0x41, 0x90, 0x6d, 0x1f,
	0x58, 0xb6, 0x2d, 0xc3, 0x8e, 0xe6, 0x72, 0x9e, 0xd6, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x96,
	0xed, 0xaf, 0x7a, 0xea, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SharedCriterionServiceClient is the client API for SharedCriterionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SharedCriterionServiceClient interface {
	// Returns the requested shared criterion in full detail.
	GetSharedCriterion(ctx context.Context, in *GetSharedCriterionRequest, opts ...grpc.CallOption) (*resources.SharedCriterion, error)
	// Creates or removes shared criteria. Operation statuses are returned.
	MutateSharedCriteria(ctx context.Context, in *MutateSharedCriteriaRequest, opts ...grpc.CallOption) (*MutateSharedCriteriaResponse, error)
}

type sharedCriterionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSharedCriterionServiceClient(cc grpc.ClientConnInterface) SharedCriterionServiceClient {
	return &sharedCriterionServiceClient{cc}
}

func (c *sharedCriterionServiceClient) GetSharedCriterion(ctx context.Context, in *GetSharedCriterionRequest, opts ...grpc.CallOption) (*resources.SharedCriterion, error) {
	out := new(resources.SharedCriterion)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.SharedCriterionService/GetSharedCriterion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedCriterionServiceClient) MutateSharedCriteria(ctx context.Context, in *MutateSharedCriteriaRequest, opts ...grpc.CallOption) (*MutateSharedCriteriaResponse, error) {
	out := new(MutateSharedCriteriaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.SharedCriterionService/MutateSharedCriteria", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharedCriterionServiceServer is the server API for SharedCriterionService service.
type SharedCriterionServiceServer interface {
	// Returns the requested shared criterion in full detail.
	GetSharedCriterion(context.Context, *GetSharedCriterionRequest) (*resources.SharedCriterion, error)
	// Creates or removes shared criteria. Operation statuses are returned.
	MutateSharedCriteria(context.Context, *MutateSharedCriteriaRequest) (*MutateSharedCriteriaResponse, error)
}

// UnimplementedSharedCriterionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSharedCriterionServiceServer struct {
}

func (*UnimplementedSharedCriterionServiceServer) GetSharedCriterion(ctx context.Context, req *GetSharedCriterionRequest) (*resources.SharedCriterion, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetSharedCriterion not implemented")
}
func (*UnimplementedSharedCriterionServiceServer) MutateSharedCriteria(ctx context.Context, req *MutateSharedCriteriaRequest) (*MutateSharedCriteriaResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateSharedCriteria not implemented")
}

func RegisterSharedCriterionServiceServer(s *grpc.Server, srv SharedCriterionServiceServer) {
	s.RegisterService(&_SharedCriterionService_serviceDesc, srv)
}

func _SharedCriterionService_GetSharedCriterion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSharedCriterionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedCriterionServiceServer).GetSharedCriterion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.SharedCriterionService/GetSharedCriterion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedCriterionServiceServer).GetSharedCriterion(ctx, req.(*GetSharedCriterionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedCriterionService_MutateSharedCriteria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateSharedCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedCriterionServiceServer).MutateSharedCriteria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.SharedCriterionService/MutateSharedCriteria",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedCriterionServiceServer).MutateSharedCriteria(ctx, req.(*MutateSharedCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SharedCriterionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.SharedCriterionService",
	HandlerType: (*SharedCriterionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSharedCriterion",
			Handler:    _SharedCriterionService_GetSharedCriterion_Handler,
		},
		{
			MethodName: "MutateSharedCriteria",
			Handler:    _SharedCriterionService_MutateSharedCriteria_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/shared_criterion_service.proto",
}

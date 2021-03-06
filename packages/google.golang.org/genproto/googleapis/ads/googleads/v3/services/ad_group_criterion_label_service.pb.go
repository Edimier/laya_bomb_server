// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/ad_group_criterion_label_service.proto

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

// Request message for
// [AdGroupCriterionLabelService.GetAdGroupCriterionLabel][google.ads.googleads.v3.services.AdGroupCriterionLabelService.GetAdGroupCriterionLabel].
type GetAdGroupCriterionLabelRequest struct {
	// Required. The resource name of the ad group criterion label to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupCriterionLabelRequest) Reset()         { *m = GetAdGroupCriterionLabelRequest{} }
func (m *GetAdGroupCriterionLabelRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupCriterionLabelRequest) ProtoMessage()    {}
func (*GetAdGroupCriterionLabelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9d794cf6cbc5ec4, []int{0}
}

func (m *GetAdGroupCriterionLabelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupCriterionLabelRequest.Unmarshal(m, b)
}
func (m *GetAdGroupCriterionLabelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupCriterionLabelRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupCriterionLabelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupCriterionLabelRequest.Merge(m, src)
}
func (m *GetAdGroupCriterionLabelRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupCriterionLabelRequest.Size(m)
}
func (m *GetAdGroupCriterionLabelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupCriterionLabelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupCriterionLabelRequest proto.InternalMessageInfo

func (m *GetAdGroupCriterionLabelRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [AdGroupCriterionLabelService.MutateAdGroupCriterionLabels][google.ads.googleads.v3.services.AdGroupCriterionLabelService.MutateAdGroupCriterionLabels].
type MutateAdGroupCriterionLabelsRequest struct {
	// Required. ID of the customer whose ad group criterion labels are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on ad group criterion labels.
	Operations []*AdGroupCriterionLabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateAdGroupCriterionLabelsRequest) Reset()         { *m = MutateAdGroupCriterionLabelsRequest{} }
func (m *MutateAdGroupCriterionLabelsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupCriterionLabelsRequest) ProtoMessage()    {}
func (*MutateAdGroupCriterionLabelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9d794cf6cbc5ec4, []int{1}
}

func (m *MutateAdGroupCriterionLabelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupCriterionLabelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupCriterionLabelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupCriterionLabelsRequest.Merge(m, src)
}
func (m *MutateAdGroupCriterionLabelsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsRequest.Size(m)
}
func (m *MutateAdGroupCriterionLabelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupCriterionLabelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupCriterionLabelsRequest proto.InternalMessageInfo

func (m *MutateAdGroupCriterionLabelsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupCriterionLabelsRequest) GetOperations() []*AdGroupCriterionLabelOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupCriterionLabelsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupCriterionLabelsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on an ad group criterion label.
type AdGroupCriterionLabelOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupCriterionLabelOperation_Create
	//	*AdGroupCriterionLabelOperation_Remove
	Operation            isAdGroupCriterionLabelOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *AdGroupCriterionLabelOperation) Reset()         { *m = AdGroupCriterionLabelOperation{} }
func (m *AdGroupCriterionLabelOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterionLabelOperation) ProtoMessage()    {}
func (*AdGroupCriterionLabelOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9d794cf6cbc5ec4, []int{2}
}

func (m *AdGroupCriterionLabelOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterionLabelOperation.Unmarshal(m, b)
}
func (m *AdGroupCriterionLabelOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterionLabelOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupCriterionLabelOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterionLabelOperation.Merge(m, src)
}
func (m *AdGroupCriterionLabelOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterionLabelOperation.Size(m)
}
func (m *AdGroupCriterionLabelOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterionLabelOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterionLabelOperation proto.InternalMessageInfo

type isAdGroupCriterionLabelOperation_Operation interface {
	isAdGroupCriterionLabelOperation_Operation()
}

type AdGroupCriterionLabelOperation_Create struct {
	Create *resources.AdGroupCriterionLabel `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupCriterionLabelOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*AdGroupCriterionLabelOperation_Create) isAdGroupCriterionLabelOperation_Operation() {}

func (*AdGroupCriterionLabelOperation_Remove) isAdGroupCriterionLabelOperation_Operation() {}

func (m *AdGroupCriterionLabelOperation) GetOperation() isAdGroupCriterionLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupCriterionLabelOperation) GetCreate() *resources.AdGroupCriterionLabel {
	if x, ok := m.GetOperation().(*AdGroupCriterionLabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupCriterionLabelOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupCriterionLabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupCriterionLabelOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupCriterionLabelOperation_Create)(nil),
		(*AdGroupCriterionLabelOperation_Remove)(nil),
	}
}

// Response message for an ad group criterion labels mutate.
type MutateAdGroupCriterionLabelsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupCriterionLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *MutateAdGroupCriterionLabelsResponse) Reset()         { *m = MutateAdGroupCriterionLabelsResponse{} }
func (m *MutateAdGroupCriterionLabelsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupCriterionLabelsResponse) ProtoMessage()    {}
func (*MutateAdGroupCriterionLabelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9d794cf6cbc5ec4, []int{3}
}

func (m *MutateAdGroupCriterionLabelsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupCriterionLabelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupCriterionLabelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupCriterionLabelsResponse.Merge(m, src)
}
func (m *MutateAdGroupCriterionLabelsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupCriterionLabelsResponse.Size(m)
}
func (m *MutateAdGroupCriterionLabelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupCriterionLabelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupCriterionLabelsResponse proto.InternalMessageInfo

func (m *MutateAdGroupCriterionLabelsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupCriterionLabelsResponse) GetResults() []*MutateAdGroupCriterionLabelResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for an ad group criterion label mutate.
type MutateAdGroupCriterionLabelResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupCriterionLabelResult) Reset()         { *m = MutateAdGroupCriterionLabelResult{} }
func (m *MutateAdGroupCriterionLabelResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupCriterionLabelResult) ProtoMessage()    {}
func (*MutateAdGroupCriterionLabelResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9d794cf6cbc5ec4, []int{4}
}

func (m *MutateAdGroupCriterionLabelResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupCriterionLabelResult.Unmarshal(m, b)
}
func (m *MutateAdGroupCriterionLabelResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupCriterionLabelResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupCriterionLabelResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupCriterionLabelResult.Merge(m, src)
}
func (m *MutateAdGroupCriterionLabelResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupCriterionLabelResult.Size(m)
}
func (m *MutateAdGroupCriterionLabelResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupCriterionLabelResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupCriterionLabelResult proto.InternalMessageInfo

func (m *MutateAdGroupCriterionLabelResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupCriterionLabelRequest)(nil), "google.ads.googleads.v3.services.GetAdGroupCriterionLabelRequest")
	proto.RegisterType((*MutateAdGroupCriterionLabelsRequest)(nil), "google.ads.googleads.v3.services.MutateAdGroupCriterionLabelsRequest")
	proto.RegisterType((*AdGroupCriterionLabelOperation)(nil), "google.ads.googleads.v3.services.AdGroupCriterionLabelOperation")
	proto.RegisterType((*MutateAdGroupCriterionLabelsResponse)(nil), "google.ads.googleads.v3.services.MutateAdGroupCriterionLabelsResponse")
	proto.RegisterType((*MutateAdGroupCriterionLabelResult)(nil), "google.ads.googleads.v3.services.MutateAdGroupCriterionLabelResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/ad_group_criterion_label_service.proto", fileDescriptor_f9d794cf6cbc5ec4)
}

var fileDescriptor_f9d794cf6cbc5ec4 = []byte{
	// 725 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcb, 0x6a, 0xdb, 0x4c,
	0x14, 0xfe, 0x25, 0x87, 0xfc, 0xcd, 0x38, 0x69, 0x61, 0x4a, 0x5b, 0xe1, 0x86, 0xc6, 0x51, 0x0c,
	0x35, 0xa6, 0x48, 0x60, 0x6f, 0x82, 0x42, 0x20, 0x72, 0x48, 0x9c, 0xd2, 0x4b, 0x82, 0x02, 0x59,
	0x14, 0x17, 0x31, 0x96, 0x26, 0xaa, 0x40, 0xd2, 0xa8, 0x33, 0x23, 0x43, 0x08, 0x81, 0xd2, 0x55,
	0x57, 0xdd, 0x74, 0xd9, 0x5d, 0x97, 0x7d, 0x82, 0x3e, 0x43, 0xa0, 0xab, 0xee, 0xb2, 0xea, 0xa2,
	0x9b, 0xf6, 0x29, 0x8a, 0x2e, 0xe3, 0x4b, 0xf0, 0x05, 0x9a, 0xdd, 0xf1, 0x39, 0x9f, 0xbe, 0xf3,
	0x9d, 0xcb, 0x1c, 0x83, 0x8e, 0x47, 0x88, 0x17, 0x60, 0x1d, 0xb9, 0x4c, 0xcf, 0xcd, 0xd4, 0xea,
	0xb7, 0x74, 0x86, 0x69, 0xdf, 0x77, 0x30, 0xd3, 0x91, 0x6b, 0x7b, 0x94, 0x24, 0xb1, 0xed, 0x50,
	0x9f, 0x63, 0xea, 0x93, 0xc8, 0x0e, 0x50, 0x0f, 0x07, 0x76, 0x81, 0xd0, 0x62, 0x4a, 0x38, 0x81,
	0xd5, 0xfc, 0x6b, 0x0d, 0xb9, 0x4c, 0x1b, 0x10, 0x69, 0xfd, 0x96, 0x26, 0x88, 0x2a, 0x3b, 0xd3,
	0x52, 0x51, 0xcc, 0x48, 0x42, 0x67, 0xe5, 0xca, 0x73, 0x54, 0x56, 0x05, 0x43, 0xec, 0xeb, 0x28,
	0x8a, 0x08, 0x47, 0xdc, 0x27, 0x11, 0x2b, 0xa2, 0x0f, 0x46, 0xa2, 0x4e, 0xe0, 0xe3, 0x88, 0x17,
	0x81, 0xb5, 0x91, 0xc0, 0xa9, 0x8f, 0x03, 0xd7, 0xee, 0xe1, 0x37, 0xa8, 0xef, 0x13, 0x7a, 0xed,
	0x4b, 0x1a, 0x3b, 0x3a, 0xe3, 0x88, 0x27, 0x05, 0xa5, 0xfa, 0x0c, 0xac, 0x75, 0x30, 0x37, 0xdd,
	0x4e, 0x2a, 0x6a, 0x57, 0x68, 0x7a, 0x9e, 0x4a, 0xb2, 0xf0, 0xdb, 0x04, 0x33, 0x0e, 0xeb, 0x60,
	0x45, 0xe8, 0xb7, 0x23, 0x14, 0x62, 0x45, 0xaa, 0x4a, 0xf5, 0xa5, 0x76, 0xe9, 0xa7, 0x29, 0x5b,
	0xcb, 0x22, 0xf2, 0x12, 0x85, 0x58, 0x7d, 0x27, 0x83, 0x8d, 0x17, 0x09, 0x47, 0x1c, 0x4f, 0x24,
	0x64, 0x82, 0xb1, 0x06, 0xca, 0x4e, 0xc2, 0x38, 0x09, 0x31, 0xb5, 0x7d, 0x77, 0x94, 0x0f, 0x08,
	0xff, 0x53, 0x17, 0x62, 0x00, 0x48, 0x8c, 0x69, 0xde, 0x01, 0x45, 0xae, 0x96, 0xea, 0xe5, 0xe6,
	0x8e, 0x36, 0x6f, 0x08, 0xda, 0xc4, 0xd4, 0x87, 0x82, 0xa8, 0x48, 0x33, 0x24, 0x86, 0x8f, 0xc1,
	0x9d, 0x18, 0x51, 0xee, 0xa3, 0xc0, 0x3e, 0x45, 0x7e, 0x90, 0x50, 0xac, 0x94, 0xaa, 0x52, 0xfd,
	0x96, 0x75, 0xbb, 0x70, 0xef, 0xe7, 0x5e, 0xb8, 0x01, 0x56, 0xfa, 0x28, 0xf0, 0x5d, 0xc4, 0xb1,
	0x4d, 0xa2, 0xe0, 0x4c, 0x59, 0xc8, 0x60, 0xcb, 0xc2, 0x79, 0x18, 0x05, 0x67, 0xea, 0x67, 0x09,
	0x3c, 0x9a, 0xad, 0x00, 0x5a, 0x60, 0xd1, 0xa1, 0x18, 0xf1, 0xbc, 0x91, 0xe5, 0xe6, 0xe6, 0xd4,
	0x9a, 0x06, 0x6b, 0x33, 0xb9, 0xa8, 0x83, 0xff, 0xac, 0x82, 0x09, 0x2a, 0x60, 0x91, 0xe2, 0x90,
	0xf4, 0xb1, 0x22, 0xa7, 0xcd, 0x4c, 0x23, 0xf9, 0xef, 0x76, 0x19, 0x2c, 0x0d, 0x8a, 0x55, 0xbf,
	0x4b, 0xa0, 0x36, 0x7b, 0x40, 0x2c, 0x26, 0x11, 0xc3, 0x70, 0x1f, 0xdc, 0xbb, 0xd6, 0x14, 0x1b,
	0x53, 0x4a, 0x68, 0xd6, 0x9a, 0x72, 0x13, 0x0a, 0xc9, 0x34, 0x76, 0xb4, 0xe3, 0x6c, 0x9f, 0xac,
	0xbb, 0xe3, 0xed, 0xda, 0x4b, 0xe1, 0xf0, 0x35, 0xf8, 0x9f, 0x62, 0x96, 0x04, 0x5c, 0x0c, 0x70,
	0x77, 0xfe, 0x00, 0x67, 0x08, 0xb4, 0x32, 0x2e, 0x4b, 0x70, 0xaa, 0x07, 0x60, 0x7d, 0x2e, 0x3a,
	0x9d, 0xdb, 0x84, 0xfd, 0x1d, 0x5f, 0xdd, 0xe6, 0xb7, 0x05, 0xb0, 0x3a, 0x91, 0xe4, 0x38, 0x97,
	0x05, 0x7f, 0x4b, 0x40, 0x99, 0xf6, 0x52, 0xa0, 0x39, 0xbf, 0xaa, 0x39, 0xaf, 0xac, 0xf2, 0xcf,
	0x5b, 0xa0, 0x1e, 0x5d, 0x99, 0xe3, 0x05, 0xbe, 0xff, 0xf1, 0xeb, 0x93, 0x6c, 0xc0, 0xcd, 0xf4,
	0xf2, 0x9c, 0x8f, 0x45, 0xb6, 0xc5, 0xf3, 0x62, 0x7a, 0x43, 0x47, 0x13, 0x57, 0x40, 0x6f, 0x5c,
	0xc0, 0x8f, 0x32, 0x58, 0x9d, 0xb5, 0x26, 0x70, 0xef, 0x46, 0x53, 0x14, 0x77, 0xa0, 0xb2, 0x7f,
	0x53, 0x9a, 0x7c, 0x5b, 0x55, 0xfb, 0xca, 0xbc, 0x3f, 0x72, 0x50, 0x9e, 0x0c, 0x5f, 0x77, 0xd6,
	0x8a, 0x1d, 0x75, 0x2b, 0x6d, 0xc5, 0xb0, 0xf6, 0xf3, 0x11, 0xf0, 0x76, 0xe3, 0x62, 0x4a, 0x27,
	0x8c, 0x30, 0xcb, 0x6d, 0x48, 0x8d, 0xca, 0xc3, 0x4b, 0x53, 0x19, 0xea, 0x2b, 0xac, 0xd8, 0x67,
	0x9a, 0x43, 0xc2, 0xf6, 0x07, 0x19, 0xd4, 0x1c, 0x12, 0xce, 0xad, 0xa5, 0xbd, 0x3e, 0x6b, 0xc1,
	0x8e, 0xd2, 0x73, 0x7c, 0x24, 0xbd, 0x3a, 0x28, 0x68, 0x3c, 0x12, 0xa0, 0xc8, 0xd3, 0x08, 0xf5,
	0x74, 0x0f, 0x47, 0xd9, 0xb1, 0xd6, 0x87, 0x89, 0xa7, 0xff, 0x9b, 0x6d, 0x09, 0xe3, 0x8b, 0x5c,
	0xea, 0x98, 0xe6, 0x57, 0xb9, 0xda, 0xc9, 0x09, 0x4d, 0x97, 0x69, 0xb9, 0x99, 0x5a, 0x27, 0x2d,
	0xad, 0x48, 0xcc, 0x2e, 0x05, 0xa4, 0x6b, 0xba, 0xac, 0x3b, 0x80, 0x74, 0x4f, 0x5a, 0x5d, 0x01,
	0xf9, 0x23, 0xd7, 0x72, 0xbf, 0x61, 0x98, 0x2e, 0x33, 0x8c, 0x01, 0xc8, 0x30, 0x4e, 0x5a, 0x86,
	0x21, 0x60, 0xbd, 0xc5, 0x4c, 0x67, 0xeb, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x77, 0xe8,
	0x94, 0x74, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupCriterionLabelServiceClient is the client API for AdGroupCriterionLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupCriterionLabelServiceClient interface {
	// Returns the requested ad group criterion label in full detail.
	GetAdGroupCriterionLabel(ctx context.Context, in *GetAdGroupCriterionLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterionLabel, error)
	// Creates and removes ad group criterion labels.
	// Operation statuses are returned.
	MutateAdGroupCriterionLabels(ctx context.Context, in *MutateAdGroupCriterionLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupCriterionLabelsResponse, error)
}

type adGroupCriterionLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupCriterionLabelServiceClient(cc grpc.ClientConnInterface) AdGroupCriterionLabelServiceClient {
	return &adGroupCriterionLabelServiceClient{cc}
}

func (c *adGroupCriterionLabelServiceClient) GetAdGroupCriterionLabel(ctx context.Context, in *GetAdGroupCriterionLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterionLabel, error) {
	out := new(resources.AdGroupCriterionLabel)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.AdGroupCriterionLabelService/GetAdGroupCriterionLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupCriterionLabelServiceClient) MutateAdGroupCriterionLabels(ctx context.Context, in *MutateAdGroupCriterionLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupCriterionLabelsResponse, error) {
	out := new(MutateAdGroupCriterionLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.AdGroupCriterionLabelService/MutateAdGroupCriterionLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupCriterionLabelServiceServer is the server API for AdGroupCriterionLabelService service.
type AdGroupCriterionLabelServiceServer interface {
	// Returns the requested ad group criterion label in full detail.
	GetAdGroupCriterionLabel(context.Context, *GetAdGroupCriterionLabelRequest) (*resources.AdGroupCriterionLabel, error)
	// Creates and removes ad group criterion labels.
	// Operation statuses are returned.
	MutateAdGroupCriterionLabels(context.Context, *MutateAdGroupCriterionLabelsRequest) (*MutateAdGroupCriterionLabelsResponse, error)
}

// UnimplementedAdGroupCriterionLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupCriterionLabelServiceServer struct {
}

func (*UnimplementedAdGroupCriterionLabelServiceServer) GetAdGroupCriterionLabel(ctx context.Context, req *GetAdGroupCriterionLabelRequest) (*resources.AdGroupCriterionLabel, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAdGroupCriterionLabel not implemented")
}
func (*UnimplementedAdGroupCriterionLabelServiceServer) MutateAdGroupCriterionLabels(ctx context.Context, req *MutateAdGroupCriterionLabelsRequest) (*MutateAdGroupCriterionLabelsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupCriterionLabels not implemented")
}

func RegisterAdGroupCriterionLabelServiceServer(s *grpc.Server, srv AdGroupCriterionLabelServiceServer) {
	s.RegisterService(&_AdGroupCriterionLabelService_serviceDesc, srv)
}

func _AdGroupCriterionLabelService_GetAdGroupCriterionLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupCriterionLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionLabelServiceServer).GetAdGroupCriterionLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.AdGroupCriterionLabelService/GetAdGroupCriterionLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionLabelServiceServer).GetAdGroupCriterionLabel(ctx, req.(*GetAdGroupCriterionLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupCriterionLabelService_MutateAdGroupCriterionLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupCriterionLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionLabelServiceServer).MutateAdGroupCriterionLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.AdGroupCriterionLabelService/MutateAdGroupCriterionLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionLabelServiceServer).MutateAdGroupCriterionLabels(ctx, req.(*MutateAdGroupCriterionLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupCriterionLabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.AdGroupCriterionLabelService",
	HandlerType: (*AdGroupCriterionLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupCriterionLabel",
			Handler:    _AdGroupCriterionLabelService_GetAdGroupCriterionLabel_Handler,
		},
		{
			MethodName: "MutateAdGroupCriterionLabels",
			Handler:    _AdGroupCriterionLabelService_MutateAdGroupCriterionLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/ad_group_criterion_label_service.proto",
}

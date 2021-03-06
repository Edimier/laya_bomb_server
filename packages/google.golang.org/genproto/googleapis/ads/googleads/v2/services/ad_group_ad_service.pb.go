// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/ad_group_ad_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// Request message for [AdGroupAdService.GetAdGroupAd][google.ads.googleads.v2.services.AdGroupAdService.GetAdGroupAd].
type GetAdGroupAdRequest struct {
	// Required. The resource name of the ad to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupAdRequest) Reset()         { *m = GetAdGroupAdRequest{} }
func (m *GetAdGroupAdRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupAdRequest) ProtoMessage()    {}
func (*GetAdGroupAdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_756c28aeb001c0ab, []int{0}
}

func (m *GetAdGroupAdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupAdRequest.Unmarshal(m, b)
}
func (m *GetAdGroupAdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupAdRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupAdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupAdRequest.Merge(m, src)
}
func (m *GetAdGroupAdRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupAdRequest.Size(m)
}
func (m *GetAdGroupAdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupAdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupAdRequest proto.InternalMessageInfo

func (m *GetAdGroupAdRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdGroupAdService.MutateAdGroupAds][google.ads.googleads.v2.services.AdGroupAdService.MutateAdGroupAds].
type MutateAdGroupAdsRequest struct {
	// Required. The ID of the customer whose ads are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual ads.
	Operations []*AdGroupAdOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateAdGroupAdsRequest) Reset()         { *m = MutateAdGroupAdsRequest{} }
func (m *MutateAdGroupAdsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupAdsRequest) ProtoMessage()    {}
func (*MutateAdGroupAdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_756c28aeb001c0ab, []int{1}
}

func (m *MutateAdGroupAdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupAdsRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupAdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupAdsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupAdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupAdsRequest.Merge(m, src)
}
func (m *MutateAdGroupAdsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupAdsRequest.Size(m)
}
func (m *MutateAdGroupAdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupAdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupAdsRequest proto.InternalMessageInfo

func (m *MutateAdGroupAdsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupAdsRequest) GetOperations() []*AdGroupAdOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupAdsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupAdsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an ad group ad.
type AdGroupAdOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Configuration for how policies are validated.
	PolicyValidationParameter *common.PolicyValidationParameter `protobuf:"bytes,5,opt,name=policy_validation_parameter,json=policyValidationParameter,proto3" json:"policy_validation_parameter,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupAdOperation_Create
	//	*AdGroupAdOperation_Update
	//	*AdGroupAdOperation_Remove
	Operation            isAdGroupAdOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *AdGroupAdOperation) Reset()         { *m = AdGroupAdOperation{} }
func (m *AdGroupAdOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdOperation) ProtoMessage()    {}
func (*AdGroupAdOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_756c28aeb001c0ab, []int{2}
}

func (m *AdGroupAdOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdOperation.Unmarshal(m, b)
}
func (m *AdGroupAdOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupAdOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdOperation.Merge(m, src)
}
func (m *AdGroupAdOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdOperation.Size(m)
}
func (m *AdGroupAdOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdOperation proto.InternalMessageInfo

func (m *AdGroupAdOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *AdGroupAdOperation) GetPolicyValidationParameter() *common.PolicyValidationParameter {
	if m != nil {
		return m.PolicyValidationParameter
	}
	return nil
}

type isAdGroupAdOperation_Operation interface {
	isAdGroupAdOperation_Operation()
}

type AdGroupAdOperation_Create struct {
	Create *resources.AdGroupAd `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupAdOperation_Update struct {
	Update *resources.AdGroupAd `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AdGroupAdOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AdGroupAdOperation_Create) isAdGroupAdOperation_Operation() {}

func (*AdGroupAdOperation_Update) isAdGroupAdOperation_Operation() {}

func (*AdGroupAdOperation_Remove) isAdGroupAdOperation_Operation() {}

func (m *AdGroupAdOperation) GetOperation() isAdGroupAdOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupAdOperation) GetCreate() *resources.AdGroupAd {
	if x, ok := m.GetOperation().(*AdGroupAdOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupAdOperation) GetUpdate() *resources.AdGroupAd {
	if x, ok := m.GetOperation().(*AdGroupAdOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *AdGroupAdOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupAdOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupAdOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupAdOperation_Create)(nil),
		(*AdGroupAdOperation_Update)(nil),
		(*AdGroupAdOperation_Remove)(nil),
	}
}

// Response message for an ad group ad mutate.
type MutateAdGroupAdsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupAdResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MutateAdGroupAdsResponse) Reset()         { *m = MutateAdGroupAdsResponse{} }
func (m *MutateAdGroupAdsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupAdsResponse) ProtoMessage()    {}
func (*MutateAdGroupAdsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_756c28aeb001c0ab, []int{3}
}

func (m *MutateAdGroupAdsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupAdsResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupAdsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupAdsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupAdsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupAdsResponse.Merge(m, src)
}
func (m *MutateAdGroupAdsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupAdsResponse.Size(m)
}
func (m *MutateAdGroupAdsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupAdsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupAdsResponse proto.InternalMessageInfo

func (m *MutateAdGroupAdsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupAdsResponse) GetResults() []*MutateAdGroupAdResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the ad mutate.
type MutateAdGroupAdResult struct {
	// The resource name returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupAdResult) Reset()         { *m = MutateAdGroupAdResult{} }
func (m *MutateAdGroupAdResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupAdResult) ProtoMessage()    {}
func (*MutateAdGroupAdResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_756c28aeb001c0ab, []int{4}
}

func (m *MutateAdGroupAdResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupAdResult.Unmarshal(m, b)
}
func (m *MutateAdGroupAdResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupAdResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupAdResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupAdResult.Merge(m, src)
}
func (m *MutateAdGroupAdResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupAdResult.Size(m)
}
func (m *MutateAdGroupAdResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupAdResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupAdResult proto.InternalMessageInfo

func (m *MutateAdGroupAdResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupAdRequest)(nil), "google.ads.googleads.v2.services.GetAdGroupAdRequest")
	proto.RegisterType((*MutateAdGroupAdsRequest)(nil), "google.ads.googleads.v2.services.MutateAdGroupAdsRequest")
	proto.RegisterType((*AdGroupAdOperation)(nil), "google.ads.googleads.v2.services.AdGroupAdOperation")
	proto.RegisterType((*MutateAdGroupAdsResponse)(nil), "google.ads.googleads.v2.services.MutateAdGroupAdsResponse")
	proto.RegisterType((*MutateAdGroupAdResult)(nil), "google.ads.googleads.v2.services.MutateAdGroupAdResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/ad_group_ad_service.proto", fileDescriptor_756c28aeb001c0ab)
}

var fileDescriptor_756c28aeb001c0ab = []byte{
	// 816 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6b, 0xe3, 0x46,
	0x14, 0xaf, 0xec, 0x36, 0x6d, 0xc6, 0x49, 0x1b, 0x26, 0xa4, 0x51, 0x9d, 0x42, 0x8d, 0x1a, 0xa8,
	0x71, 0x83, 0x04, 0x4a, 0x4a, 0x89, 0xd2, 0x50, 0x64, 0xda, 0x38, 0x3d, 0xa4, 0x71, 0x15, 0x30,
	0xb4, 0x18, 0xc4, 0x44, 0x9a, 0xb8, 0x22, 0x92, 0x46, 0x9d, 0x19, 0x19, 0x4c, 0xc8, 0xa5, 0xfd,
	0x08, 0xfd, 0x06, 0x7b, 0xdc, 0xfb, 0x1e, 0xf6, 0x2b, 0x84, 0xbd, 0xed, 0x2d, 0x87, 0x65, 0x0f,
	0x7b, 0x58, 0xf6, 0x33, 0xec, 0x61, 0x91, 0x46, 0x23, 0xff, 0x49, 0x8c, 0x49, 0x6e, 0x4f, 0xf3,
	0x7e, 0xbf, 0xdf, 0x7b, 0x6f, 0xde, 0x7b, 0x23, 0x60, 0x0d, 0x08, 0x19, 0x84, 0xd8, 0x40, 0x3e,
	0x33, 0x84, 0x99, 0x59, 0x43, 0xd3, 0x60, 0x98, 0x0e, 0x03, 0x0f, 0x33, 0x03, 0xf9, 0xee, 0x80,
	0x92, 0x34, 0x71, 0x91, 0xef, 0x16, 0x87, 0x7a, 0x42, 0x09, 0x27, 0xb0, 0x21, 0x08, 0x3a, 0xf2,
	0x99, 0x5e, 0x72, 0xf5, 0xa1, 0xa9, 0x4b, 0x6e, 0xfd, 0xfb, 0x79, 0xea, 0x1e, 0x89, 0x22, 0x12,
	0x1b, 0x09, 0x09, 0x03, 0x6f, 0x24, 0xe4, 0xea, 0xbb, 0xf3, 0xc0, 0x14, 0x33, 0x92, 0xd2, 0x99,
	0x5c, 0x0a, 0xd2, 0xd7, 0x92, 0x94, 0x04, 0x06, 0x8a, 0x63, 0xc2, 0x11, 0x0f, 0x48, 0xcc, 0x0a,
	0xef, 0xe6, 0x84, 0xd7, 0x0b, 0x03, 0x1c, 0xf3, 0xc2, 0xf1, 0xcd, 0x84, 0xe3, 0x22, 0xc0, 0xa1,
	0xef, 0x9e, 0xe3, 0xbf, 0xd1, 0x30, 0x20, 0xb4, 0x00, 0x14, 0xb5, 0x19, 0xf9, 0xd7, 0x79, 0x7a,
	0x51, 0xa0, 0x22, 0xc4, 0x2e, 0x67, 0xb4, 0x69, 0xe2, 0x19, 0x8c, 0x23, 0x9e, 0x16, 0x41, 0xb5,
	0x9f, 0xc1, 0x7a, 0x07, 0x73, 0xdb, 0xef, 0x64, 0x99, 0xda, 0xbe, 0x83, 0xff, 0x49, 0x31, 0xe3,
	0xb0, 0x09, 0x56, 0x65, 0x21, 0x6e, 0x8c, 0x22, 0xac, 0x2a, 0x0d, 0xa5, 0xb9, 0xdc, 0xae, 0xbe,
	0xb6, 0x2b, 0xce, 0x8a, 0xf4, 0xfc, 0x8e, 0x22, 0xac, 0xbd, 0x55, 0xc0, 0xe6, 0x49, 0xca, 0x11,
	0xc7, 0xa5, 0x08, 0x93, 0x2a, 0xdb, 0xa0, 0xe6, 0xa5, 0x8c, 0x93, 0x08, 0x53, 0x37, 0xf0, 0x27,
	0x35, 0x80, 0x3c, 0xff, 0xcd, 0x87, 0x7f, 0x02, 0x40, 0x12, 0x4c, 0xc5, 0x5d, 0xa8, 0x95, 0x46,
	0xb5, 0x59, 0x33, 0xf7, 0xf4, 0x45, 0xed, 0xd2, 0xcb, 0x70, 0xa7, 0x92, 0x5c, 0x48, 0x8f, 0xc5,
	0xe0, 0x77, 0xe0, 0x8b, 0x04, 0x51, 0x1e, 0xa0, 0xd0, 0xbd, 0x40, 0x41, 0x98, 0x52, 0xac, 0x56,
	0x1b, 0x4a, 0xf3, 0x33, 0xe7, 0xf3, 0xe2, 0xf8, 0x48, 0x9c, 0xc2, 0x6f, 0xc1, 0xea, 0x10, 0x85,
	0x81, 0x8f, 0x38, 0x76, 0x49, 0x1c, 0x8e, 0xd4, 0x8f, 0x73, 0xd8, 0x8a, 0x3c, 0x3c, 0x8d, 0xc3,
	0x91, 0xf6, 0x5f, 0x15, 0xc0, 0xbb, 0x51, 0xe1, 0x01, 0xa8, 0xa5, 0x49, 0xce, 0xcc, 0x2e, 0x3c,
	0x67, 0xd6, 0xcc, 0xba, 0x2c, 0x40, 0xf6, 0x44, 0x3f, 0xca, 0x7a, 0x72, 0x82, 0xd8, 0xa5, 0x03,
	0x04, 0x3c, 0xb3, 0xe1, 0x08, 0x6c, 0x89, 0xb9, 0x72, 0x8b, 0x50, 0x01, 0x89, 0xdd, 0x04, 0x51,
	0x14, 0x61, 0x8e, 0xa9, 0xfa, 0x49, 0x2e, 0xb6, 0x3f, 0xf7, 0x36, 0xc4, 0x68, 0xea, 0xdd, 0x5c,
	0xa2, 0x57, 0x2a, 0x74, 0xa5, 0x80, 0xf3, 0x55, 0x32, 0xcf, 0x05, 0x8f, 0xc0, 0x92, 0x47, 0x31,
	0xe2, 0xa2, 0xb9, 0x35, 0x73, 0x67, 0x6e, 0x94, 0x72, 0xa6, 0xc7, 0x97, 0x7e, 0xfc, 0x91, 0x53,
	0xb0, 0x33, 0x1d, 0x51, 0x90, 0x5a, 0x79, 0x9c, 0x8e, 0x60, 0x43, 0x15, 0x2c, 0x51, 0x1c, 0x91,
	0xa1, 0xe8, 0xd1, 0x72, 0xe6, 0x11, 0xdf, 0xed, 0x1a, 0x58, 0x2e, 0x9b, 0xaa, 0x3d, 0x53, 0x80,
	0x7a, 0x77, 0xe0, 0x58, 0x42, 0x62, 0x96, 0xe5, 0xb2, 0x31, 0xd3, 0x70, 0x17, 0x53, 0x4a, 0x68,
	0x2e, 0x59, 0x33, 0xa1, 0x4c, 0x8d, 0x26, 0x9e, 0x7e, 0x96, 0xef, 0x81, 0xb3, 0x3e, 0x3d, 0x0a,
	0xbf, 0x66, 0x70, 0xf8, 0x07, 0xf8, 0x94, 0x62, 0x96, 0x86, 0x5c, 0x0e, 0xe4, 0x8f, 0x8b, 0x07,
	0x72, 0x26, 0x29, 0x27, 0xe7, 0x3b, 0x52, 0x47, 0xfb, 0x09, 0x6c, 0xdc, 0x8b, 0xc8, 0x66, 0xef,
	0x9e, 0x5d, 0x9b, 0x5e, 0x33, 0xf3, 0x45, 0x15, 0xac, 0x95, 0xc4, 0x33, 0x11, 0x12, 0x3e, 0x57,
	0xc0, 0xca, 0xe4, 0xf6, 0xc2, 0x1f, 0x16, 0x67, 0x79, 0xcf, 0xb6, 0xd7, 0x1f, 0xd4, 0x31, 0xed,
	0x97, 0x5b, 0x7b, 0x3a, 0xe1, 0x7f, 0x5f, 0xbe, 0xf9, 0xbf, 0xa2, 0xc3, 0x9d, 0xec, 0xf9, 0xbb,
	0x9a, 0xf2, 0x1c, 0xca, 0x35, 0x67, 0x46, 0xcb, 0x40, 0x65, 0xbb, 0x8c, 0xd6, 0x35, 0x7c, 0xa5,
	0x80, 0xb5, 0xd9, 0x36, 0xc2, 0xfd, 0x07, 0xdf, 0xb2, 0x7c, 0x6b, 0xea, 0xd6, 0x63, 0xa8, 0x62,
	0x6a, 0xb4, 0xb3, 0x5b, 0xfb, 0xcb, 0x89, 0x87, 0x6a, 0x67, 0xfc, 0x82, 0xe4, 0xa5, 0xed, 0x69,
	0x46, 0xfe, 0x1b, 0x28, 0x6b, 0xb9, 0x9a, 0x00, 0x1f, 0xb6, 0xae, 0x27, 0x2a, 0xb3, 0xa2, 0x3c,
	0x86, 0xa5, 0xb4, 0xea, 0x5b, 0x37, 0xb6, 0x3a, 0xce, 0xa3, 0xb0, 0x92, 0x80, 0x65, 0x3b, 0xdb,
	0x7e, 0xaf, 0x80, 0x6d, 0x8f, 0x44, 0x0b, 0x73, 0x6e, 0x6f, 0xcc, 0x36, 0xbd, 0x9b, 0xbd, 0x27,
	0x5d, 0xe5, 0xaf, 0xe3, 0x82, 0x3a, 0x20, 0x21, 0x8a, 0x07, 0x3a, 0xa1, 0x03, 0x63, 0x80, 0xe3,
	0xfc, 0xb5, 0x31, 0xc6, 0xc1, 0xe6, 0xff, 0x2a, 0x0f, 0xa4, 0xf1, 0xa4, 0x52, 0xed, 0xd8, 0xf6,
	0xd3, 0x4a, 0xa3, 0x23, 0x04, 0x6d, 0x9f, 0xe9, 0xc2, 0xcc, 0xac, 0x9e, 0xa9, 0x17, 0x81, 0xd9,
	0x8d, 0x84, 0xf4, 0x6d, 0x9f, 0xf5, 0x4b, 0x48, 0xbf, 0x67, 0xf6, 0x25, 0xe4, 0x5d, 0x65, 0x5b,
	0x9c, 0x5b, 0x56, 0x76, 0x19, 0x56, 0x09, 0xb2, 0xac, 0x9e, 0x69, 0x59, 0x12, 0x76, 0xbe, 0x94,
	0xe7, 0xb9, 0xfb, 0x21, 0x00, 0x00, 0xff, 0xff, 0x95, 0x9e, 0xb0, 0x46, 0xd1, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupAdServiceClient is the client API for AdGroupAdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupAdServiceClient interface {
	// Returns the requested ad in full detail.
	GetAdGroupAd(ctx context.Context, in *GetAdGroupAdRequest, opts ...grpc.CallOption) (*resources.AdGroupAd, error)
	// Creates, updates, or removes ads. Operation statuses are returned.
	MutateAdGroupAds(ctx context.Context, in *MutateAdGroupAdsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdsResponse, error)
}

type adGroupAdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupAdServiceClient(cc grpc.ClientConnInterface) AdGroupAdServiceClient {
	return &adGroupAdServiceClient{cc}
}

func (c *adGroupAdServiceClient) GetAdGroupAd(ctx context.Context, in *GetAdGroupAdRequest, opts ...grpc.CallOption) (*resources.AdGroupAd, error) {
	out := new(resources.AdGroupAd)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdGroupAdService/GetAdGroupAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupAdServiceClient) MutateAdGroupAds(ctx context.Context, in *MutateAdGroupAdsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdsResponse, error) {
	out := new(MutateAdGroupAdsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdGroupAdService/MutateAdGroupAds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupAdServiceServer is the server API for AdGroupAdService service.
type AdGroupAdServiceServer interface {
	// Returns the requested ad in full detail.
	GetAdGroupAd(context.Context, *GetAdGroupAdRequest) (*resources.AdGroupAd, error)
	// Creates, updates, or removes ads. Operation statuses are returned.
	MutateAdGroupAds(context.Context, *MutateAdGroupAdsRequest) (*MutateAdGroupAdsResponse, error)
}

// UnimplementedAdGroupAdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupAdServiceServer struct {
}

func (*UnimplementedAdGroupAdServiceServer) GetAdGroupAd(ctx context.Context, req *GetAdGroupAdRequest) (*resources.AdGroupAd, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAdGroupAd not implemented")
}
func (*UnimplementedAdGroupAdServiceServer) MutateAdGroupAds(ctx context.Context, req *MutateAdGroupAdsRequest) (*MutateAdGroupAdsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupAds not implemented")
}

func RegisterAdGroupAdServiceServer(s *grpc.Server, srv AdGroupAdServiceServer) {
	s.RegisterService(&_AdGroupAdService_serviceDesc, srv)
}

func _AdGroupAdService_GetAdGroupAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAdServiceServer).GetAdGroupAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdGroupAdService/GetAdGroupAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAdServiceServer).GetAdGroupAd(ctx, req.(*GetAdGroupAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupAdService_MutateAdGroupAds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupAdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAdServiceServer).MutateAdGroupAds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdGroupAdService/MutateAdGroupAds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAdServiceServer).MutateAdGroupAds(ctx, req.(*MutateAdGroupAdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupAdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.AdGroupAdService",
	HandlerType: (*AdGroupAdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupAd",
			Handler:    _AdGroupAdService_GetAdGroupAd_Handler,
		},
		{
			MethodName: "MutateAdGroupAds",
			Handler:    _AdGroupAdService_MutateAdGroupAds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/ad_group_ad_service.proto",
}

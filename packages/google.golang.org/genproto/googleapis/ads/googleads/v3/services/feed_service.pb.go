// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/feed_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
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

// Request message for [FeedService.GetFeed][google.ads.googleads.v3.services.FeedService.GetFeed].
type GetFeedRequest struct {
	// Required. The resource name of the feed to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedRequest) Reset()         { *m = GetFeedRequest{} }
func (m *GetFeedRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedRequest) ProtoMessage()    {}
func (*GetFeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b4d1072354ada7f, []int{0}
}

func (m *GetFeedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedRequest.Unmarshal(m, b)
}
func (m *GetFeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedRequest.Merge(m, src)
}
func (m *GetFeedRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedRequest.Size(m)
}
func (m *GetFeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedRequest proto.InternalMessageInfo

func (m *GetFeedRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [FeedService.MutateFeeds][google.ads.googleads.v3.services.FeedService.MutateFeeds].
type MutateFeedsRequest struct {
	// Required. The ID of the customer whose feeds are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual feeds.
	Operations []*FeedOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateFeedsRequest) Reset()         { *m = MutateFeedsRequest{} }
func (m *MutateFeedsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateFeedsRequest) ProtoMessage()    {}
func (*MutateFeedsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b4d1072354ada7f, []int{1}
}

func (m *MutateFeedsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedsRequest.Unmarshal(m, b)
}
func (m *MutateFeedsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedsRequest.Marshal(b, m, deterministic)
}
func (m *MutateFeedsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedsRequest.Merge(m, src)
}
func (m *MutateFeedsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateFeedsRequest.Size(m)
}
func (m *MutateFeedsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedsRequest proto.InternalMessageInfo

func (m *MutateFeedsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateFeedsRequest) GetOperations() []*FeedOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateFeedsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateFeedsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an feed.
type FeedOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*FeedOperation_Create
	//	*FeedOperation_Update
	//	*FeedOperation_Remove
	Operation            isFeedOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *FeedOperation) Reset()         { *m = FeedOperation{} }
func (m *FeedOperation) String() string { return proto.CompactTextString(m) }
func (*FeedOperation) ProtoMessage()    {}
func (*FeedOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b4d1072354ada7f, []int{2}
}

func (m *FeedOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedOperation.Unmarshal(m, b)
}
func (m *FeedOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedOperation.Marshal(b, m, deterministic)
}
func (m *FeedOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedOperation.Merge(m, src)
}
func (m *FeedOperation) XXX_Size() int {
	return xxx_messageInfo_FeedOperation.Size(m)
}
func (m *FeedOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedOperation.DiscardUnknown(m)
}

var xxx_messageInfo_FeedOperation proto.InternalMessageInfo

func (m *FeedOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isFeedOperation_Operation interface {
	isFeedOperation_Operation()
}

type FeedOperation_Create struct {
	Create *resources.Feed `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type FeedOperation_Update struct {
	Update *resources.Feed `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type FeedOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*FeedOperation_Create) isFeedOperation_Operation() {}

func (*FeedOperation_Update) isFeedOperation_Operation() {}

func (*FeedOperation_Remove) isFeedOperation_Operation() {}

func (m *FeedOperation) GetOperation() isFeedOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *FeedOperation) GetCreate() *resources.Feed {
	if x, ok := m.GetOperation().(*FeedOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *FeedOperation) GetUpdate() *resources.Feed {
	if x, ok := m.GetOperation().(*FeedOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *FeedOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*FeedOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FeedOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FeedOperation_Create)(nil),
		(*FeedOperation_Update)(nil),
		(*FeedOperation_Remove)(nil),
	}
}

// Response message for an feed mutate.
type MutateFeedsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateFeedResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MutateFeedsResponse) Reset()         { *m = MutateFeedsResponse{} }
func (m *MutateFeedsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateFeedsResponse) ProtoMessage()    {}
func (*MutateFeedsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b4d1072354ada7f, []int{3}
}

func (m *MutateFeedsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedsResponse.Unmarshal(m, b)
}
func (m *MutateFeedsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedsResponse.Marshal(b, m, deterministic)
}
func (m *MutateFeedsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedsResponse.Merge(m, src)
}
func (m *MutateFeedsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateFeedsResponse.Size(m)
}
func (m *MutateFeedsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedsResponse proto.InternalMessageInfo

func (m *MutateFeedsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateFeedsResponse) GetResults() []*MutateFeedResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the feed mutate.
type MutateFeedResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedResult) Reset()         { *m = MutateFeedResult{} }
func (m *MutateFeedResult) String() string { return proto.CompactTextString(m) }
func (*MutateFeedResult) ProtoMessage()    {}
func (*MutateFeedResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b4d1072354ada7f, []int{4}
}

func (m *MutateFeedResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedResult.Unmarshal(m, b)
}
func (m *MutateFeedResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedResult.Marshal(b, m, deterministic)
}
func (m *MutateFeedResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedResult.Merge(m, src)
}
func (m *MutateFeedResult) XXX_Size() int {
	return xxx_messageInfo_MutateFeedResult.Size(m)
}
func (m *MutateFeedResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedResult proto.InternalMessageInfo

func (m *MutateFeedResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedRequest)(nil), "google.ads.googleads.v3.services.GetFeedRequest")
	proto.RegisterType((*MutateFeedsRequest)(nil), "google.ads.googleads.v3.services.MutateFeedsRequest")
	proto.RegisterType((*FeedOperation)(nil), "google.ads.googleads.v3.services.FeedOperation")
	proto.RegisterType((*MutateFeedsResponse)(nil), "google.ads.googleads.v3.services.MutateFeedsResponse")
	proto.RegisterType((*MutateFeedResult)(nil), "google.ads.googleads.v3.services.MutateFeedResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/feed_service.proto", fileDescriptor_4b4d1072354ada7f)
}

var fileDescriptor_4b4d1072354ada7f = []byte{
	// 745 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x6a, 0xdb, 0x4a,
	0x14, 0xbe, 0xb2, 0x2f, 0xc9, 0xcd, 0x28, 0xc9, 0x0d, 0x13, 0xee, 0xad, 0x71, 0x0b, 0x35, 0x6a,
	0x20, 0xae, 0x09, 0x52, 0xb1, 0x5b, 0x0a, 0x0a, 0xa1, 0xc8, 0x50, 0x27, 0x85, 0xe6, 0x07, 0xa5,
	0x64, 0x51, 0x0c, 0x66, 0x22, 0x8d, 0x5d, 0x11, 0x49, 0xa3, 0xce, 0x8c, 0x0c, 0x21, 0x64, 0xd3,
	0x57, 0xe8, 0x03, 0x14, 0x0a, 0xdd, 0xb4, 0x6f, 0x92, 0x6d, 0x77, 0x59, 0x94, 0x2e, 0xba, 0xea,
	0x0b, 0x94, 0xee, 0xca, 0x68, 0x66, 0xfc, 0x13, 0x1a, 0xdc, 0xec, 0x8e, 0xe6, 0x7c, 0xdf, 0x77,
	0xce, 0x7c, 0x67, 0x8e, 0x0d, 0x5a, 0x03, 0x42, 0x06, 0x31, 0x76, 0x50, 0xc8, 0x1c, 0x19, 0x8a,
	0x68, 0xd8, 0x72, 0x18, 0xa6, 0xc3, 0x28, 0xc0, 0xcc, 0xe9, 0x63, 0x1c, 0xf6, 0xd4, 0x97, 0x9d,
	0x51, 0xc2, 0x09, 0xac, 0x49, 0xa4, 0x8d, 0x42, 0x66, 0x8f, 0x48, 0xf6, 0xb0, 0x65, 0x6b, 0x52,
	0x75, 0xe3, 0x3a, 0x59, 0x8a, 0x19, 0xc9, 0xa9, 0xd6, 0x95, 0x7a, 0xd5, 0x3b, 0x1a, 0x9d, 0x45,
	0x0e, 0x4a, 0x53, 0xc2, 0x11, 0x8f, 0x48, 0xca, 0x54, 0xf6, 0xd6, 0x44, 0x36, 0x88, 0x23, 0x9c,
	0x72, 0x95, 0xb8, 0x3b, 0x91, 0xe8, 0x47, 0x38, 0x0e, 0x7b, 0xc7, 0xf8, 0x15, 0x1a, 0x46, 0x84,
	0x2a, 0x80, 0xea, 0xd3, 0x29, 0xbe, 0x8e, 0xf3, 0xbe, 0x42, 0x25, 0x88, 0x9d, 0x5c, 0xd1, 0xa6,
	0x59, 0xe0, 0x30, 0x8e, 0x78, 0xae, 0x8a, 0x5a, 0x2e, 0x58, 0xde, 0xc6, 0xbc, 0x83, 0x71, 0xe8,
	0xe3, 0xd7, 0x39, 0x66, 0x1c, 0xd6, 0xc1, 0x92, 0x6e, 0xbe, 0x97, 0xa2, 0x04, 0x57, 0x8c, 0x9a,
	0x51, 0x5f, 0x68, 0x97, 0xbf, 0x7a, 0x25, 0x7f, 0x51, 0x67, 0xf6, 0x50, 0x82, 0xad, 0x2f, 0x06,
	0x80, 0xbb, 0x39, 0x47, 0x1c, 0x0b, 0x3e, 0xd3, 0x02, 0x6b, 0xc0, 0x0c, 0x72, 0xc6, 0x49, 0x82,
	0x69, 0x2f, 0x0a, 0x27, 0xe9, 0x40, 0x9f, 0x3f, 0x0b, 0xe1, 0x0b, 0x00, 0x48, 0x86, 0xa9, 0x74,
	0xa0, 0x52, 0xaa, 0x95, 0xeb, 0x66, 0xd3, 0xb1, 0x67, 0x19, 0x6e, 0x8b, 0x4a, 0xfb, 0x9a, 0xa7,
	0x54, 0xc7, 0x3a, 0x70, 0x1d, 0xfc, 0x9b, 0x21, 0xca, 0x23, 0x14, 0xf7, 0xfa, 0x28, 0x8a, 0x73,
	0x8a, 0x2b, 0xe5, 0x9a, 0x51, 0xff, 0xc7, 0x5f, 0x56, 0xc7, 0x1d, 0x79, 0x0a, 0xef, 0x81, 0xa5,
	0x21, 0x8a, 0xa3, 0x10, 0x71, 0xdc, 0x23, 0x69, 0x7c, 0x5a, 0xf9, 0xbb, 0x80, 0x2d, 0xea, 0xc3,
	0xfd, 0x34, 0x3e, 0xb5, 0x7e, 0x1a, 0x60, 0x69, 0xaa, 0x20, 0xdc, 0x04, 0x66, 0x9e, 0x15, 0x24,
	0x61, 0x6e, 0x41, 0x32, 0x9b, 0x55, 0xdd, 0xb6, 0xf6, 0xdf, 0xee, 0x08, 0xff, 0x77, 0x11, 0x3b,
	0xf1, 0x81, 0x84, 0x8b, 0x18, 0x7a, 0x60, 0x2e, 0xa0, 0x18, 0x71, 0x69, 0xa9, 0xd9, 0x5c, 0xbf,
	0xf6, 0xba, 0xa3, 0xd7, 0x53, 0xdc, 0x77, 0xe7, 0x2f, 0x5f, 0x11, 0x85, 0x84, 0x14, 0xac, 0x94,
	0x6e, 0x2c, 0x21, 0x89, 0xb0, 0x02, 0xe6, 0x28, 0x4e, 0xc8, 0x50, 0x3a, 0xb3, 0x20, 0x32, 0xf2,
	0xbb, 0x6d, 0x82, 0x85, 0x91, 0x95, 0xd6, 0x27, 0x03, 0xac, 0x4e, 0x0d, 0x97, 0x65, 0x24, 0x65,
	0x18, 0x76, 0xc0, 0x7f, 0x57, 0x1c, 0xee, 0x61, 0x4a, 0x09, 0x2d, 0xd4, 0xcc, 0x26, 0xd4, 0x0d,
	0xd1, 0x2c, 0xb0, 0x0f, 0x8b, 0x97, 0xe6, 0xaf, 0x4e, 0x7b, 0xff, 0x54, 0xc0, 0xe1, 0x73, 0x30,
	0x4f, 0x31, 0xcb, 0x63, 0xae, 0x87, 0xdf, 0x9c, 0x3d, 0xfc, 0x71, 0x3f, 0x7e, 0x41, 0xf5, 0xb5,
	0x84, 0xf5, 0x18, 0xac, 0x5c, 0x4d, 0x8a, 0x11, 0xff, 0xe6, 0x21, 0x4f, 0xbf, 0xe1, 0xe6, 0xbb,
	0x32, 0x30, 0x05, 0xe7, 0x50, 0xd6, 0x80, 0x1f, 0x0c, 0x30, 0xaf, 0x16, 0x02, 0x3e, 0x98, 0xdd,
	0xd1, 0xf4, 0xee, 0x54, 0xff, 0x74, 0x1c, 0xd6, 0x93, 0x4b, 0x6f, 0xba, 0xb9, 0x37, 0x9f, 0xbf,
	0xbd, 0x2d, 0xdd, 0x87, 0xeb, 0xe2, 0xb7, 0xe3, 0x6c, 0x2a, 0xb3, 0xa5, 0x97, 0x86, 0x39, 0x8d,
	0xe2, 0xc7, 0x84, 0x39, 0x8d, 0x73, 0x78, 0x61, 0x00, 0x73, 0x62, 0x3c, 0xf0, 0xe1, 0x4d, 0xdc,
	0xd3, 0xab, 0x5a, 0x7d, 0x74, 0x43, 0x96, 0x7c, 0x03, 0xd6, 0xde, 0xa5, 0xf7, 0xff, 0xc4, 0x8a,
	0x6f, 0x8c, 0x17, 0xb0, 0xb8, 0x86, 0x63, 0x35, 0xc4, 0x35, 0xc6, 0x7d, 0x9f, 0x4d, 0x80, 0xb7,
	0x1a, 0xe7, 0xf2, 0x16, 0x6e, 0x52, 0x28, 0xbb, 0x46, 0xa3, 0x7a, 0xfb, 0xc2, 0xab, 0x8c, 0xab,
	0xab, 0x28, 0x8b, 0x98, 0x1d, 0x90, 0xa4, 0xfd, 0xc3, 0x00, 0x6b, 0x01, 0x49, 0x66, 0x76, 0xda,
	0x5e, 0x99, 0x98, 0xe3, 0x81, 0xd8, 0xc4, 0x03, 0xe3, 0xe5, 0x8e, 0x62, 0x0d, 0x48, 0x8c, 0xd2,
	0x81, 0x4d, 0xe8, 0xc0, 0x19, 0xe0, 0xb4, 0xd8, 0x53, 0x67, 0x5c, 0xe7, 0xfa, 0x7f, 0x85, 0x4d,
	0x1d, 0xbc, 0x2f, 0x95, 0xb7, 0x3d, 0xef, 0x63, 0xa9, 0xb6, 0x2d, 0x05, 0xbd, 0x90, 0xd9, 0x32,
	0x14, 0xd1, 0x51, 0xcb, 0x56, 0x85, 0xd9, 0x85, 0x86, 0x74, 0xbd, 0x90, 0x75, 0x47, 0x90, 0xee,
	0x51, 0xab, 0xab, 0x21, 0xdf, 0x4b, 0x6b, 0xf2, 0xdc, 0x75, 0xbd, 0x90, 0xb9, 0xee, 0x08, 0xe4,
	0xba, 0x47, 0x2d, 0xd7, 0xd5, 0xb0, 0xe3, 0xb9, 0xa2, 0xcf, 0xd6, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x43, 0xb9, 0x76, 0xde, 0xbc, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FeedServiceClient is the client API for FeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedServiceClient interface {
	// Returns the requested feed in full detail.
	GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*resources.Feed, error)
	// Creates, updates, or removes feeds. Operation statuses are
	// returned.
	MutateFeeds(ctx context.Context, in *MutateFeedsRequest, opts ...grpc.CallOption) (*MutateFeedsResponse, error)
}

type feedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedServiceClient(cc grpc.ClientConnInterface) FeedServiceClient {
	return &feedServiceClient{cc}
}

func (c *feedServiceClient) GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*resources.Feed, error) {
	out := new(resources.Feed)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.FeedService/GetFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) MutateFeeds(ctx context.Context, in *MutateFeedsRequest, opts ...grpc.CallOption) (*MutateFeedsResponse, error) {
	out := new(MutateFeedsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.FeedService/MutateFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServiceServer is the server API for FeedService service.
type FeedServiceServer interface {
	// Returns the requested feed in full detail.
	GetFeed(context.Context, *GetFeedRequest) (*resources.Feed, error)
	// Creates, updates, or removes feeds. Operation statuses are
	// returned.
	MutateFeeds(context.Context, *MutateFeedsRequest) (*MutateFeedsResponse, error)
}

// UnimplementedFeedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFeedServiceServer struct {
}

func (*UnimplementedFeedServiceServer) GetFeed(ctx context.Context, req *GetFeedRequest) (*resources.Feed, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetFeed not implemented")
}
func (*UnimplementedFeedServiceServer) MutateFeeds(ctx context.Context, req *MutateFeedsRequest) (*MutateFeedsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateFeeds not implemented")
}

func RegisterFeedServiceServer(s *grpc.Server, srv FeedServiceServer) {
	s.RegisterService(&_FeedService_serviceDesc, srv)
}

func _FeedService_GetFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).GetFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.FeedService/GetFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).GetFeed(ctx, req.(*GetFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_MutateFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).MutateFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.FeedService/MutateFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).MutateFeeds(ctx, req.(*MutateFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.FeedService",
	HandlerType: (*FeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeed",
			Handler:    _FeedService_GetFeed_Handler,
		},
		{
			MethodName: "MutateFeeds",
			Handler:    _FeedService_MutateFeeds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/feed_service.proto",
}

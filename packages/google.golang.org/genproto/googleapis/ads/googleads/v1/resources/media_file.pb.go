// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/media_file.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// A media file.
type MediaFile struct {
	// Immutable. The resource name of the media file.
	// Media file resource names have the form:
	//
	// `customers/{customer_id}/mediaFiles/{media_file_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the media file.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Immutable. Type of the media file.
	Type enums.MediaTypeEnum_MediaType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.MediaTypeEnum_MediaType" json:"type,omitempty"`
	// Output only. The mime type of the media file.
	MimeType enums.MimeTypeEnum_MimeType `protobuf:"varint,6,opt,name=mime_type,json=mimeType,proto3,enum=google.ads.googleads.v1.enums.MimeTypeEnum_MimeType" json:"mime_type,omitempty"`
	// Immutable. The URL of where the original media file was downloaded from (or a file
	// name). Only used for media of type AUDIO and IMAGE.
	SourceUrl *wrappers.StringValue `protobuf:"bytes,7,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	// Immutable. The name of the media file. The name can be used by clients to help
	// identify previously uploaded media.
	Name *wrappers.StringValue `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The size of the media file in bytes.
	FileSize *wrappers.Int64Value `protobuf:"bytes,9,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	// The specific type of the media file.
	//
	// Types that are valid to be assigned to Mediatype:
	//	*MediaFile_Image
	//	*MediaFile_MediaBundle
	//	*MediaFile_Audio
	//	*MediaFile_Video
	Mediatype            isMediaFile_Mediatype `protobuf_oneof:"mediatype"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MediaFile) Reset()         { *m = MediaFile{} }
func (m *MediaFile) String() string { return proto.CompactTextString(m) }
func (*MediaFile) ProtoMessage()    {}
func (*MediaFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ffee613242502eb, []int{0}
}

func (m *MediaFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaFile.Unmarshal(m, b)
}
func (m *MediaFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaFile.Marshal(b, m, deterministic)
}
func (m *MediaFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaFile.Merge(m, src)
}
func (m *MediaFile) XXX_Size() int {
	return xxx_messageInfo_MediaFile.Size(m)
}
func (m *MediaFile) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaFile.DiscardUnknown(m)
}

var xxx_messageInfo_MediaFile proto.InternalMessageInfo

func (m *MediaFile) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MediaFile) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MediaFile) GetType() enums.MediaTypeEnum_MediaType {
	if m != nil {
		return m.Type
	}
	return enums.MediaTypeEnum_UNSPECIFIED
}

func (m *MediaFile) GetMimeType() enums.MimeTypeEnum_MimeType {
	if m != nil {
		return m.MimeType
	}
	return enums.MimeTypeEnum_UNSPECIFIED
}

func (m *MediaFile) GetSourceUrl() *wrappers.StringValue {
	if m != nil {
		return m.SourceUrl
	}
	return nil
}

func (m *MediaFile) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MediaFile) GetFileSize() *wrappers.Int64Value {
	if m != nil {
		return m.FileSize
	}
	return nil
}

type isMediaFile_Mediatype interface {
	isMediaFile_Mediatype()
}

type MediaFile_Image struct {
	Image *MediaImage `protobuf:"bytes,3,opt,name=image,proto3,oneof"`
}

type MediaFile_MediaBundle struct {
	MediaBundle *MediaBundle `protobuf:"bytes,4,opt,name=media_bundle,json=mediaBundle,proto3,oneof"`
}

type MediaFile_Audio struct {
	Audio *MediaAudio `protobuf:"bytes,10,opt,name=audio,proto3,oneof"`
}

type MediaFile_Video struct {
	Video *MediaVideo `protobuf:"bytes,11,opt,name=video,proto3,oneof"`
}

func (*MediaFile_Image) isMediaFile_Mediatype() {}

func (*MediaFile_MediaBundle) isMediaFile_Mediatype() {}

func (*MediaFile_Audio) isMediaFile_Mediatype() {}

func (*MediaFile_Video) isMediaFile_Mediatype() {}

func (m *MediaFile) GetMediatype() isMediaFile_Mediatype {
	if m != nil {
		return m.Mediatype
	}
	return nil
}

func (m *MediaFile) GetImage() *MediaImage {
	if x, ok := m.GetMediatype().(*MediaFile_Image); ok {
		return x.Image
	}
	return nil
}

func (m *MediaFile) GetMediaBundle() *MediaBundle {
	if x, ok := m.GetMediatype().(*MediaFile_MediaBundle); ok {
		return x.MediaBundle
	}
	return nil
}

func (m *MediaFile) GetAudio() *MediaAudio {
	if x, ok := m.GetMediatype().(*MediaFile_Audio); ok {
		return x.Audio
	}
	return nil
}

func (m *MediaFile) GetVideo() *MediaVideo {
	if x, ok := m.GetMediatype().(*MediaFile_Video); ok {
		return x.Video
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MediaFile) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MediaFile_Image)(nil),
		(*MediaFile_MediaBundle)(nil),
		(*MediaFile_Audio)(nil),
		(*MediaFile_Video)(nil),
	}
}

// Encapsulates an Image.
type MediaImage struct {
	// Immutable. Raw image data.
	Data                 *wrappers.BytesValue `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MediaImage) Reset()         { *m = MediaImage{} }
func (m *MediaImage) String() string { return proto.CompactTextString(m) }
func (*MediaImage) ProtoMessage()    {}
func (*MediaImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ffee613242502eb, []int{1}
}

func (m *MediaImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaImage.Unmarshal(m, b)
}
func (m *MediaImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaImage.Marshal(b, m, deterministic)
}
func (m *MediaImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaImage.Merge(m, src)
}
func (m *MediaImage) XXX_Size() int {
	return xxx_messageInfo_MediaImage.Size(m)
}
func (m *MediaImage) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaImage.DiscardUnknown(m)
}

var xxx_messageInfo_MediaImage proto.InternalMessageInfo

func (m *MediaImage) GetData() *wrappers.BytesValue {
	if m != nil {
		return m.Data
	}
	return nil
}

// Represents a ZIP archive media the content of which contains HTML5 assets.
type MediaBundle struct {
	// Immutable. Raw zipped data.
	Data                 *wrappers.BytesValue `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MediaBundle) Reset()         { *m = MediaBundle{} }
func (m *MediaBundle) String() string { return proto.CompactTextString(m) }
func (*MediaBundle) ProtoMessage()    {}
func (*MediaBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ffee613242502eb, []int{2}
}

func (m *MediaBundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaBundle.Unmarshal(m, b)
}
func (m *MediaBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaBundle.Marshal(b, m, deterministic)
}
func (m *MediaBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaBundle.Merge(m, src)
}
func (m *MediaBundle) XXX_Size() int {
	return xxx_messageInfo_MediaBundle.Size(m)
}
func (m *MediaBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaBundle.DiscardUnknown(m)
}

var xxx_messageInfo_MediaBundle proto.InternalMessageInfo

func (m *MediaBundle) GetData() *wrappers.BytesValue {
	if m != nil {
		return m.Data
	}
	return nil
}

// Encapsulates an Audio.
type MediaAudio struct {
	// Output only. The duration of the Audio in milliseconds.
	AdDurationMillis     *wrappers.Int64Value `protobuf:"bytes,1,opt,name=ad_duration_millis,json=adDurationMillis,proto3" json:"ad_duration_millis,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MediaAudio) Reset()         { *m = MediaAudio{} }
func (m *MediaAudio) String() string { return proto.CompactTextString(m) }
func (*MediaAudio) ProtoMessage()    {}
func (*MediaAudio) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ffee613242502eb, []int{3}
}

func (m *MediaAudio) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaAudio.Unmarshal(m, b)
}
func (m *MediaAudio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaAudio.Marshal(b, m, deterministic)
}
func (m *MediaAudio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaAudio.Merge(m, src)
}
func (m *MediaAudio) XXX_Size() int {
	return xxx_messageInfo_MediaAudio.Size(m)
}
func (m *MediaAudio) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaAudio.DiscardUnknown(m)
}

var xxx_messageInfo_MediaAudio proto.InternalMessageInfo

func (m *MediaAudio) GetAdDurationMillis() *wrappers.Int64Value {
	if m != nil {
		return m.AdDurationMillis
	}
	return nil
}

// Encapsulates a Video.
type MediaVideo struct {
	// Output only. The duration of the Video in milliseconds.
	AdDurationMillis *wrappers.Int64Value `protobuf:"bytes,1,opt,name=ad_duration_millis,json=adDurationMillis,proto3" json:"ad_duration_millis,omitempty"`
	// Immutable. The YouTube video ID (as seen in YouTube URLs).
	YoutubeVideoId *wrappers.StringValue `protobuf:"bytes,2,opt,name=youtube_video_id,json=youtubeVideoId,proto3" json:"youtube_video_id,omitempty"`
	// Output only. The Advertising Digital Identification code for this video, as defined by
	// the American Association of Advertising Agencies, used mainly for
	// television commercials.
	AdvertisingIdCode *wrappers.StringValue `protobuf:"bytes,3,opt,name=advertising_id_code,json=advertisingIdCode,proto3" json:"advertising_id_code,omitempty"`
	// Output only. The Industry Standard Commercial Identifier code for this video, used
	// mainly for television commercials.
	IsciCode             *wrappers.StringValue `protobuf:"bytes,4,opt,name=isci_code,json=isciCode,proto3" json:"isci_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MediaVideo) Reset()         { *m = MediaVideo{} }
func (m *MediaVideo) String() string { return proto.CompactTextString(m) }
func (*MediaVideo) ProtoMessage()    {}
func (*MediaVideo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ffee613242502eb, []int{4}
}

func (m *MediaVideo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaVideo.Unmarshal(m, b)
}
func (m *MediaVideo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaVideo.Marshal(b, m, deterministic)
}
func (m *MediaVideo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaVideo.Merge(m, src)
}
func (m *MediaVideo) XXX_Size() int {
	return xxx_messageInfo_MediaVideo.Size(m)
}
func (m *MediaVideo) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaVideo.DiscardUnknown(m)
}

var xxx_messageInfo_MediaVideo proto.InternalMessageInfo

func (m *MediaVideo) GetAdDurationMillis() *wrappers.Int64Value {
	if m != nil {
		return m.AdDurationMillis
	}
	return nil
}

func (m *MediaVideo) GetYoutubeVideoId() *wrappers.StringValue {
	if m != nil {
		return m.YoutubeVideoId
	}
	return nil
}

func (m *MediaVideo) GetAdvertisingIdCode() *wrappers.StringValue {
	if m != nil {
		return m.AdvertisingIdCode
	}
	return nil
}

func (m *MediaVideo) GetIsciCode() *wrappers.StringValue {
	if m != nil {
		return m.IsciCode
	}
	return nil
}

func init() {
	proto.RegisterType((*MediaFile)(nil), "google.ads.googleads.v1.resources.MediaFile")
	proto.RegisterType((*MediaImage)(nil), "google.ads.googleads.v1.resources.MediaImage")
	proto.RegisterType((*MediaBundle)(nil), "google.ads.googleads.v1.resources.MediaBundle")
	proto.RegisterType((*MediaAudio)(nil), "google.ads.googleads.v1.resources.MediaAudio")
	proto.RegisterType((*MediaVideo)(nil), "google.ads.googleads.v1.resources.MediaVideo")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/media_file.proto", fileDescriptor_7ffee613242502eb)
}

var fileDescriptor_7ffee613242502eb = []byte{
	// 779 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc7, 0xc9, 0x57, 0x1b, 0x4f, 0x4a, 0x55, 0x86, 0x1b, 0x53, 0x2a, 0x68, 0x23, 0x2a, 0x55,
	0x88, 0x8e, 0x49, 0x29, 0xbd, 0x30, 0x02, 0x61, 0x17, 0x28, 0x29, 0xea, 0x07, 0x2e, 0x8d, 0x10,
	0x8a, 0x64, 0x4d, 0x32, 0x53, 0x33, 0x92, 0xed, 0x89, 0x3c, 0x76, 0x50, 0x5a, 0xf5, 0x21, 0x78,
	0x85, 0xbd, 0xdc, 0x67, 0xd8, 0x27, 0xd8, 0xa7, 0xe8, 0x75, 0x1f, 0x61, 0xf7, 0x66, 0x35, 0xe3,
	0xb1, 0x13, 0xed, 0x6e, 0x37, 0xe9, 0x6a, 0xef, 0xce, 0xe4, 0xfc, 0xff, 0x3f, 0x9f, 0x39, 0x73,
	0x66, 0x02, 0xf6, 0x02, 0xce, 0x83, 0x90, 0x5a, 0x98, 0x08, 0x2b, 0x0f, 0x65, 0x34, 0xee, 0x58,
	0x09, 0x15, 0x3c, 0x4b, 0x86, 0x54, 0x58, 0x11, 0x25, 0x0c, 0xfb, 0x57, 0x2c, 0xa4, 0x68, 0x94,
	0xf0, 0x94, 0xc3, 0xad, 0x5c, 0x88, 0x30, 0x11, 0xa8, 0xf4, 0xa0, 0x71, 0x07, 0x95, 0x9e, 0x75,
	0xf4, 0x10, 0x96, 0xc6, 0x59, 0x54, 0x20, 0xd3, 0xc9, 0x48, 0x23, 0xd7, 0x77, 0xe7, 0xe8, 0x59,
	0x44, 0x67, 0xe5, 0x5f, 0x16, 0xf2, 0x11, 0xb3, 0xae, 0x18, 0x0d, 0x89, 0x3f, 0xa0, 0xff, 0xe2,
	0x31, 0xe3, 0x89, 0x16, 0x7c, 0x36, 0x23, 0x28, 0xaa, 0xd2, 0xa9, 0x2f, 0x74, 0x4a, 0xad, 0x06,
	0xd9, 0x95, 0xf5, 0x5f, 0x82, 0x47, 0x23, 0x9a, 0x08, 0x9d, 0xdf, 0x98, 0xb1, 0xe2, 0x38, 0xe6,
	0x29, 0x4e, 0x19, 0x8f, 0x75, 0xb6, 0xfd, 0xff, 0x32, 0x30, 0x4e, 0x64, 0xf5, 0xbf, 0xb1, 0x90,
	0xc2, 0x33, 0xf0, 0x71, 0x41, 0xf7, 0x63, 0x1c, 0x51, 0xb3, 0xb2, 0x59, 0xd9, 0x31, 0xdc, 0xaf,
	0xef, 0x9c, 0xc6, 0x0b, 0xe7, 0x2b, 0xd0, 0x9e, 0x76, 0x47, 0x47, 0x23, 0x26, 0xd0, 0x90, 0x47,
	0x56, 0x89, 0xf0, 0x56, 0x0a, 0xc0, 0x29, 0x8e, 0x28, 0xfc, 0x16, 0x54, 0x19, 0x31, 0xab, 0x9b,
	0x95, 0x9d, 0xd6, 0xde, 0xe7, 0xda, 0x84, 0x8a, 0x4a, 0x51, 0x37, 0x4e, 0x0f, 0xf6, 0x7b, 0x38,
	0xcc, 0xa8, 0x5b, 0xbb, 0x73, 0x6a, 0x5e, 0x95, 0x11, 0x78, 0x06, 0xea, 0xb2, 0x31, 0x66, 0x63,
	0xb3, 0xb2, 0xb3, 0xba, 0x77, 0x80, 0x1e, 0x3a, 0x1b, 0xd5, 0x48, 0xa4, 0xbe, 0xfb, 0xd7, 0x64,
	0x44, 0x7f, 0x8d, 0xb3, 0x68, 0xba, 0x92, 0xb8, 0x86, 0xa7, 0x40, 0xb0, 0x07, 0x8c, 0xb2, 0xdd,
	0xe6, 0x92, 0xa2, 0xee, 0xcf, 0xa3, 0xb2, 0x88, 0x4e, 0xa1, 0x7a, 0x91, 0x97, 0xd8, 0x8c, 0xf4,
	0x12, 0xfe, 0x0c, 0x80, 0xee, 0x54, 0x96, 0x84, 0xe6, 0xb2, 0xda, 0xe2, 0xc6, 0x1b, 0x5b, 0xbc,
	0x48, 0x13, 0x16, 0x07, 0xe5, 0x1e, 0x1b, 0x9e, 0x91, 0x9b, 0x2e, 0x93, 0x10, 0x7e, 0x0f, 0xea,
	0xaa, 0xc9, 0xcd, 0x45, 0xbd, 0x4a, 0x0e, 0x7f, 0x04, 0x86, 0x1c, 0x5e, 0x5f, 0xb0, 0x6b, 0x6a,
	0x1a, 0x0b, 0xb6, 0xb6, 0x29, 0x2d, 0x17, 0xec, 0x9a, 0xc2, 0x63, 0xd0, 0x60, 0x11, 0x0e, 0xa8,
	0x59, 0x53, 0xd6, 0x5d, 0x34, 0x77, 0xfa, 0xf3, 0xbe, 0x76, 0xa5, 0x49, 0xd5, 0xf1, 0xfb, 0x47,
	0x5e, 0x8e, 0x80, 0x7f, 0x83, 0x95, 0x7c, 0xf4, 0x07, 0x59, 0x4c, 0x42, 0x6a, 0xd6, 0x15, 0x12,
	0x2d, 0x8a, 0x74, 0x95, 0xab, 0x60, 0xb6, 0xa2, 0xe9, 0x6f, 0xb2, 0x4a, 0x9c, 0x11, 0xc6, 0x4d,
	0xf0, 0xb8, 0x2a, 0x1d, 0x69, 0x52, 0x5b, 0x96, 0x55, 0x2a, 0x84, 0x64, 0x8d, 0x19, 0xa1, 0xdc,
	0x6c, 0x3d, 0x8e, 0xd5, 0x93, 0xa6, 0x72, 0xc7, 0x0a, 0x61, 0x5f, 0xde, 0x3b, 0xde, 0x22, 0xf7,
	0x00, 0x7e, 0x33, 0xcc, 0x44, 0xca, 0x23, 0x9a, 0x08, 0xeb, 0xa6, 0x08, 0x6f, 0xf3, 0x87, 0x42,
	0xe6, 0x85, 0x75, 0x33, 0x7d, 0x87, 0x6e, 0xdd, 0x16, 0x30, 0xd4, 0x52, 0x0e, 0x69, 0xdb, 0x05,
	0x60, 0xda, 0x71, 0xb8, 0x0f, 0xea, 0x04, 0xa7, 0x58, 0x5d, 0xc5, 0xb7, 0x9d, 0xb4, 0x3b, 0x49,
	0xa9, 0x98, 0x1d, 0x12, 0xa9, 0x6e, 0x1f, 0x82, 0xd6, 0x4c, 0x8b, 0xdf, 0x13, 0xd2, 0xd7, 0x85,
	0xa8, 0xa6, 0xc2, 0x53, 0x00, 0x31, 0xf1, 0x49, 0x96, 0xa8, 0x07, 0xc4, 0x8f, 0x58, 0x18, 0x32,
	0xf1, 0x20, 0xf1, 0xf5, 0x01, 0x5c, 0xc3, 0xe4, 0x17, 0x6d, 0x3d, 0x51, 0xce, 0xf6, 0xb3, 0xaa,
	0xc6, 0xab, 0x3e, 0x7f, 0x68, 0x3c, 0xfc, 0x03, 0xac, 0x4d, 0x78, 0x96, 0x66, 0x03, 0xea, 0xab,
	0xa3, 0xf3, 0xcb, 0x87, 0x68, 0x81, 0x9b, 0xb6, 0xaa, 0xad, 0xaa, 0xb4, 0x2e, 0x81, 0x7f, 0x82,
	0x4f, 0x31, 0x19, 0xd3, 0x24, 0x65, 0x82, 0xc5, 0x81, 0xcf, 0x88, 0x3f, 0xe4, 0xa4, 0xb8, 0x42,
	0x73, 0x79, 0x35, 0xef, 0x93, 0x19, 0x77, 0x97, 0x1c, 0x72, 0x42, 0xe1, 0x4f, 0xc0, 0x60, 0x62,
	0xc8, 0x72, 0x50, 0x7d, 0x51, 0x50, 0x53, 0x7a, 0xa4, 0xdf, 0x7d, 0x59, 0x01, 0xdb, 0x43, 0x1e,
	0xcd, 0x1f, 0x66, 0x77, 0xb5, 0x9c, 0xca, 0x73, 0xc9, 0x3d, 0xaf, 0xfc, 0x73, 0xac, 0x4d, 0x01,
	0x0f, 0x71, 0x1c, 0x20, 0x9e, 0x04, 0x56, 0x40, 0x63, 0xf5, 0x55, 0x6b, 0x3a, 0xd1, 0xef, 0xf8,
	0x13, 0xfd, 0xa1, 0x8c, 0x9e, 0x54, 0x6b, 0x47, 0x8e, 0xf3, 0xb4, 0xba, 0x75, 0x94, 0x23, 0x1d,
	0x22, 0x50, 0x1e, 0xca, 0xa8, 0xd7, 0x41, 0x5e, 0xa1, 0x7c, 0x5e, 0x68, 0xfa, 0x0e, 0x11, 0xfd,
	0x52, 0xd3, 0xef, 0x75, 0xfa, 0xa5, 0xe6, 0xbe, 0xba, 0x9d, 0x27, 0x6c, 0xdb, 0x21, 0xc2, 0xb6,
	0x4b, 0x95, 0x6d, 0xf7, 0x3a, 0xb6, 0x5d, 0xea, 0x06, 0x4b, 0xaa, 0xd8, 0xef, 0x5e, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xdb, 0xb2, 0x17, 0xf0, 0xf0, 0x07, 0x00, 0x00,
}
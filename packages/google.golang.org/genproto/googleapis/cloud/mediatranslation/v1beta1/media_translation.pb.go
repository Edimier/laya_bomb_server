// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/mediatranslation/v1beta1/media_translation.proto

package mediatranslation

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Indicates the type of speech event.
type StreamingTranslateSpeechResponse_SpeechEventType int32

const (
	// No speech event specified.
	StreamingTranslateSpeechResponse_SPEECH_EVENT_TYPE_UNSPECIFIED StreamingTranslateSpeechResponse_SpeechEventType = 0
	// This event indicates that the server has detected the end of the user's
	// speech utterance and expects no additional speech. Therefore, the server
	// will not process additional audio (although it may subsequently return
	// additional results). When the client receives 'END_OF_SINGLE_UTTERANCE'
	// event, the client should stop sending the requests. However, clients
	// should keep receiving remaining responses until the stream is terminated.
	// To construct the complete sentence in a streaming way, one should
	// override (if 'is_final' of previous response is false), or append (if
	// 'is_final' of previous response is true). This event is only sent if
	// `single_utterance` was set to `true`, and is not used otherwise.
	StreamingTranslateSpeechResponse_END_OF_SINGLE_UTTERANCE StreamingTranslateSpeechResponse_SpeechEventType = 1
)

var StreamingTranslateSpeechResponse_SpeechEventType_name = map[int32]string{
	0: "SPEECH_EVENT_TYPE_UNSPECIFIED",
	1: "END_OF_SINGLE_UTTERANCE",
}

var StreamingTranslateSpeechResponse_SpeechEventType_value = map[string]int32{
	"SPEECH_EVENT_TYPE_UNSPECIFIED": 0,
	"END_OF_SINGLE_UTTERANCE":       1,
}

func (x StreamingTranslateSpeechResponse_SpeechEventType) String() string {
	return proto.EnumName(StreamingTranslateSpeechResponse_SpeechEventType_name, int32(x))
}

func (StreamingTranslateSpeechResponse_SpeechEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b2f67d966c0ba6c, []int{4, 0}
}

// Provides information to the speech translation that specifies how to process
// the request.
type TranslateSpeechConfig struct {
	// Required. Encoding of audio data.
	// Supported formats:
	//
	// - `linear16`
	//
	//   Uncompressed 16-bit signed little-endian samples (Linear PCM).
	//
	//
	AudioEncoding string `protobuf:"bytes,1,opt,name=audio_encoding,json=audioEncoding,proto3" json:"audio_encoding,omitempty"`
	// Required. Source language code (BCP-47) of the input audio.
	SourceLanguageCode string `protobuf:"bytes,2,opt,name=source_language_code,json=sourceLanguageCode,proto3" json:"source_language_code,omitempty"`
	// Optional. A list of up to 3 additional language codes (BCP-47), listing possible
	// alternative languages of the supplied audio. If alternative source
	// languages are listed, speech translation result will translate in the most
	// likely language detected including the main source_language_code. The
	// translated result will include the language code of the language detected
	// in the audio.
	AlternativeSourceLanguageCodes []string `protobuf:"bytes,6,rep,name=alternative_source_language_codes,json=alternativeSourceLanguageCodes,proto3" json:"alternative_source_language_codes,omitempty"`
	// Required. Target language code (BCP-47) of the output.
	TargetLanguageCode string `protobuf:"bytes,3,opt,name=target_language_code,json=targetLanguageCode,proto3" json:"target_language_code,omitempty"`
	// Optional. Sample rate in Hertz of the audio data. Valid values are:
	// 8000-48000. 16000 is optimal. For best results, set the sampling rate of
	// the audio source to 16000 Hz. If that's not possible, use the native sample
	// rate of the audio source (instead of re-sampling). This field can only be
	// omitted for `FLAC` and `WAV` audio files.
	SampleRateHertz int32 `protobuf:"varint,4,opt,name=sample_rate_hertz,json=sampleRateHertz,proto3" json:"sample_rate_hertz,omitempty"`
	// Optional.
	Model                string   `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranslateSpeechConfig) Reset()         { *m = TranslateSpeechConfig{} }
func (m *TranslateSpeechConfig) String() string { return proto.CompactTextString(m) }
func (*TranslateSpeechConfig) ProtoMessage()    {}
func (*TranslateSpeechConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2f67d966c0ba6c, []int{0}
}

func (m *TranslateSpeechConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslateSpeechConfig.Unmarshal(m, b)
}
func (m *TranslateSpeechConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslateSpeechConfig.Marshal(b, m, deterministic)
}
func (m *TranslateSpeechConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslateSpeechConfig.Merge(m, src)
}
func (m *TranslateSpeechConfig) XXX_Size() int {
	return xxx_messageInfo_TranslateSpeechConfig.Size(m)
}
func (m *TranslateSpeechConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslateSpeechConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TranslateSpeechConfig proto.InternalMessageInfo

func (m *TranslateSpeechConfig) GetAudioEncoding() string {
	if m != nil {
		return m.AudioEncoding
	}
	return ""
}

func (m *TranslateSpeechConfig) GetSourceLanguageCode() string {
	if m != nil {
		return m.SourceLanguageCode
	}
	return ""
}

func (m *TranslateSpeechConfig) GetAlternativeSourceLanguageCodes() []string {
	if m != nil {
		return m.AlternativeSourceLanguageCodes
	}
	return nil
}

func (m *TranslateSpeechConfig) GetTargetLanguageCode() string {
	if m != nil {
		return m.TargetLanguageCode
	}
	return ""
}

func (m *TranslateSpeechConfig) GetSampleRateHertz() int32 {
	if m != nil {
		return m.SampleRateHertz
	}
	return 0
}

func (m *TranslateSpeechConfig) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

// Config used for streaming translation.
type StreamingTranslateSpeechConfig struct {
	// Required. The common config for all the following audio contents.
	AudioConfig *TranslateSpeechConfig `protobuf:"bytes,1,opt,name=audio_config,json=audioConfig,proto3" json:"audio_config,omitempty"`
	// Optional. If `false` or omitted, the system performs
	// continuous translation (continuing to wait for and process audio even if
	// the user pauses speaking) until the client closes the input stream (gRPC
	// API) or until the maximum time limit has been reached. May return multiple
	// `StreamingTranslateSpeechResult`s with the `is_final` flag set to `true`.
	//
	// If `true`, the speech translator will detect a single spoken utterance.
	// When it detects that the user has paused or stopped speaking, it will
	// return an `END_OF_SINGLE_UTTERANCE` event and cease translation.
	// When the client receives 'END_OF_SINGLE_UTTERANCE' event, the client should
	// stop sending the requests. However, clients should keep receiving remaining
	// responses until the stream is terminated. To construct the complete
	// sentence in a streaming way, one should override (if 'is_final' of previous
	// response is false), or append (if 'is_final' of previous response is true).
	SingleUtterance      bool     `protobuf:"varint,2,opt,name=single_utterance,json=singleUtterance,proto3" json:"single_utterance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingTranslateSpeechConfig) Reset()         { *m = StreamingTranslateSpeechConfig{} }
func (m *StreamingTranslateSpeechConfig) String() string { return proto.CompactTextString(m) }
func (*StreamingTranslateSpeechConfig) ProtoMessage()    {}
func (*StreamingTranslateSpeechConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2f67d966c0ba6c, []int{1}
}

func (m *StreamingTranslateSpeechConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingTranslateSpeechConfig.Unmarshal(m, b)
}
func (m *StreamingTranslateSpeechConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingTranslateSpeechConfig.Marshal(b, m, deterministic)
}
func (m *StreamingTranslateSpeechConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingTranslateSpeechConfig.Merge(m, src)
}
func (m *StreamingTranslateSpeechConfig) XXX_Size() int {
	return xxx_messageInfo_StreamingTranslateSpeechConfig.Size(m)
}
func (m *StreamingTranslateSpeechConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingTranslateSpeechConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingTranslateSpeechConfig proto.InternalMessageInfo

func (m *StreamingTranslateSpeechConfig) GetAudioConfig() *TranslateSpeechConfig {
	if m != nil {
		return m.AudioConfig
	}
	return nil
}

func (m *StreamingTranslateSpeechConfig) GetSingleUtterance() bool {
	if m != nil {
		return m.SingleUtterance
	}
	return false
}

// The top-level message sent by the client for the `StreamingTranslateSpeech`
// method. Multiple `StreamingTranslateSpeechRequest` messages are sent. The
// first message must contain a `streaming_config` message and must not contain
// `audio_content` data. All subsequent messages must contain `audio_content`
// data and must not contain a `streaming_config` message.
type StreamingTranslateSpeechRequest struct {
	// The streaming request, which is either a streaming config or content.
	//
	// Types that are valid to be assigned to StreamingRequest:
	//	*StreamingTranslateSpeechRequest_StreamingConfig
	//	*StreamingTranslateSpeechRequest_AudioContent
	StreamingRequest     isStreamingTranslateSpeechRequest_StreamingRequest `protobuf_oneof:"streaming_request"`
	XXX_NoUnkeyedLiteral struct{}                                           `json:"-"`
	XXX_unrecognized     []byte                                             `json:"-"`
	XXX_sizecache        int32                                              `json:"-"`
}

func (m *StreamingTranslateSpeechRequest) Reset()         { *m = StreamingTranslateSpeechRequest{} }
func (m *StreamingTranslateSpeechRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingTranslateSpeechRequest) ProtoMessage()    {}
func (*StreamingTranslateSpeechRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2f67d966c0ba6c, []int{2}
}

func (m *StreamingTranslateSpeechRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingTranslateSpeechRequest.Unmarshal(m, b)
}
func (m *StreamingTranslateSpeechRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingTranslateSpeechRequest.Marshal(b, m, deterministic)
}
func (m *StreamingTranslateSpeechRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingTranslateSpeechRequest.Merge(m, src)
}
func (m *StreamingTranslateSpeechRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingTranslateSpeechRequest.Size(m)
}
func (m *StreamingTranslateSpeechRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingTranslateSpeechRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingTranslateSpeechRequest proto.InternalMessageInfo

type isStreamingTranslateSpeechRequest_StreamingRequest interface {
	isStreamingTranslateSpeechRequest_StreamingRequest()
}

type StreamingTranslateSpeechRequest_StreamingConfig struct {
	StreamingConfig *StreamingTranslateSpeechConfig `protobuf:"bytes,1,opt,name=streaming_config,json=streamingConfig,proto3,oneof"`
}

type StreamingTranslateSpeechRequest_AudioContent struct {
	AudioContent []byte `protobuf:"bytes,2,opt,name=audio_content,json=audioContent,proto3,oneof"`
}

func (*StreamingTranslateSpeechRequest_StreamingConfig) isStreamingTranslateSpeechRequest_StreamingRequest() {
}

func (*StreamingTranslateSpeechRequest_AudioContent) isStreamingTranslateSpeechRequest_StreamingRequest() {
}

func (m *StreamingTranslateSpeechRequest) GetStreamingRequest() isStreamingTranslateSpeechRequest_StreamingRequest {
	if m != nil {
		return m.StreamingRequest
	}
	return nil
}

func (m *StreamingTranslateSpeechRequest) GetStreamingConfig() *StreamingTranslateSpeechConfig {
	if x, ok := m.GetStreamingRequest().(*StreamingTranslateSpeechRequest_StreamingConfig); ok {
		return x.StreamingConfig
	}
	return nil
}

func (m *StreamingTranslateSpeechRequest) GetAudioContent() []byte {
	if x, ok := m.GetStreamingRequest().(*StreamingTranslateSpeechRequest_AudioContent); ok {
		return x.AudioContent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StreamingTranslateSpeechRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StreamingTranslateSpeechRequest_StreamingConfig)(nil),
		(*StreamingTranslateSpeechRequest_AudioContent)(nil),
	}
}

// A streaming speech translation result corresponding to a portion of the audio
// that is currently being processed.
type StreamingTranslateSpeechResult struct {
	// Translation result.
	//
	// Use oneof field to reserve for future tts result.
	//
	// Types that are valid to be assigned to Result:
	//	*StreamingTranslateSpeechResult_TextTranslationResult_
	Result isStreamingTranslateSpeechResult_Result `protobuf_oneof:"result"`
	// Output only. The debug only recognition result in original language. This field is debug
	// only and will be set to empty string if not available.
	// This is implementation detail and will not be backward compatible.
	//
	// Still need to decide whether to expose this field by default.
	RecognitionResult    string   `protobuf:"bytes,3,opt,name=recognition_result,json=recognitionResult,proto3" json:"recognition_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingTranslateSpeechResult) Reset()         { *m = StreamingTranslateSpeechResult{} }
func (m *StreamingTranslateSpeechResult) String() string { return proto.CompactTextString(m) }
func (*StreamingTranslateSpeechResult) ProtoMessage()    {}
func (*StreamingTranslateSpeechResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2f67d966c0ba6c, []int{3}
}

func (m *StreamingTranslateSpeechResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingTranslateSpeechResult.Unmarshal(m, b)
}
func (m *StreamingTranslateSpeechResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingTranslateSpeechResult.Marshal(b, m, deterministic)
}
func (m *StreamingTranslateSpeechResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingTranslateSpeechResult.Merge(m, src)
}
func (m *StreamingTranslateSpeechResult) XXX_Size() int {
	return xxx_messageInfo_StreamingTranslateSpeechResult.Size(m)
}
func (m *StreamingTranslateSpeechResult) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingTranslateSpeechResult.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingTranslateSpeechResult proto.InternalMessageInfo

type isStreamingTranslateSpeechResult_Result interface {
	isStreamingTranslateSpeechResult_Result()
}

type StreamingTranslateSpeechResult_TextTranslationResult_ struct {
	TextTranslationResult *StreamingTranslateSpeechResult_TextTranslationResult `protobuf:"bytes,1,opt,name=text_translation_result,json=textTranslationResult,proto3,oneof"`
}

func (*StreamingTranslateSpeechResult_TextTranslationResult_) isStreamingTranslateSpeechResult_Result() {
}

func (m *StreamingTranslateSpeechResult) GetResult() isStreamingTranslateSpeechResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *StreamingTranslateSpeechResult) GetTextTranslationResult() *StreamingTranslateSpeechResult_TextTranslationResult {
	if x, ok := m.GetResult().(*StreamingTranslateSpeechResult_TextTranslationResult_); ok {
		return x.TextTranslationResult
	}
	return nil
}

func (m *StreamingTranslateSpeechResult) GetRecognitionResult() string {
	if m != nil {
		return m.RecognitionResult
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StreamingTranslateSpeechResult) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StreamingTranslateSpeechResult_TextTranslationResult_)(nil),
	}
}

// Text translation result.
type StreamingTranslateSpeechResult_TextTranslationResult struct {
	// Output only. The translated sentence.
	Translation string `protobuf:"bytes,1,opt,name=translation,proto3" json:"translation,omitempty"`
	// Output only. If `false`, this `StreamingTranslateSpeechResult` represents
	// an interim result that may change. If `true`, this is the final time the
	// translation service will return this particular
	// `StreamingTranslateSpeechResult`, the streaming translator will not
	// return any further hypotheses for this portion of the transcript and
	// corresponding audio.
	IsFinal bool `protobuf:"varint,2,opt,name=is_final,json=isFinal,proto3" json:"is_final,omitempty"`
	// Output only. The source language code (BCP-47) detected in the audio. Speech
	// translation result will translate in the most likely language detected
	// including the alternative source languages and main source_language_code.
	DetectedSourceLanguageCode string   `protobuf:"bytes,3,opt,name=detected_source_language_code,json=detectedSourceLanguageCode,proto3" json:"detected_source_language_code,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *StreamingTranslateSpeechResult_TextTranslationResult) Reset() {
	*m = StreamingTranslateSpeechResult_TextTranslationResult{}
}
func (m *StreamingTranslateSpeechResult_TextTranslationResult) String() string {
	return proto.CompactTextString(m)
}
func (*StreamingTranslateSpeechResult_TextTranslationResult) ProtoMessage() {}
func (*StreamingTranslateSpeechResult_TextTranslationResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2f67d966c0ba6c, []int{3, 0}
}

func (m *StreamingTranslateSpeechResult_TextTranslationResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingTranslateSpeechResult_TextTranslationResult.Unmarshal(m, b)
}
func (m *StreamingTranslateSpeechResult_TextTranslationResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingTranslateSpeechResult_TextTranslationResult.Marshal(b, m, deterministic)
}
func (m *StreamingTranslateSpeechResult_TextTranslationResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingTranslateSpeechResult_TextTranslationResult.Merge(m, src)
}
func (m *StreamingTranslateSpeechResult_TextTranslationResult) XXX_Size() int {
	return xxx_messageInfo_StreamingTranslateSpeechResult_TextTranslationResult.Size(m)
}
func (m *StreamingTranslateSpeechResult_TextTranslationResult) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingTranslateSpeechResult_TextTranslationResult.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingTranslateSpeechResult_TextTranslationResult proto.InternalMessageInfo

func (m *StreamingTranslateSpeechResult_TextTranslationResult) GetTranslation() string {
	if m != nil {
		return m.Translation
	}
	return ""
}

func (m *StreamingTranslateSpeechResult_TextTranslationResult) GetIsFinal() bool {
	if m != nil {
		return m.IsFinal
	}
	return false
}

func (m *StreamingTranslateSpeechResult_TextTranslationResult) GetDetectedSourceLanguageCode() string {
	if m != nil {
		return m.DetectedSourceLanguageCode
	}
	return ""
}

// A streaming speech translation response corresponding to a portion of
// the audio currently processed.
type StreamingTranslateSpeechResponse struct {
	// Output only. If set, returns a [google.rpc.Status][google.rpc.Status] message that
	// specifies the error for the operation.
	Error *status.Status `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// Output only. The translation result that is currently being processed (is_final could be
	// true or false).
	Result *StreamingTranslateSpeechResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	// Output only. Indicates the type of speech event.
	SpeechEventType      StreamingTranslateSpeechResponse_SpeechEventType `protobuf:"varint,3,opt,name=speech_event_type,json=speechEventType,proto3,enum=google.cloud.mediatranslation.v1beta1.StreamingTranslateSpeechResponse_SpeechEventType" json:"speech_event_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                         `json:"-"`
	XXX_unrecognized     []byte                                           `json:"-"`
	XXX_sizecache        int32                                            `json:"-"`
}

func (m *StreamingTranslateSpeechResponse) Reset()         { *m = StreamingTranslateSpeechResponse{} }
func (m *StreamingTranslateSpeechResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingTranslateSpeechResponse) ProtoMessage()    {}
func (*StreamingTranslateSpeechResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2f67d966c0ba6c, []int{4}
}

func (m *StreamingTranslateSpeechResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingTranslateSpeechResponse.Unmarshal(m, b)
}
func (m *StreamingTranslateSpeechResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingTranslateSpeechResponse.Marshal(b, m, deterministic)
}
func (m *StreamingTranslateSpeechResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingTranslateSpeechResponse.Merge(m, src)
}
func (m *StreamingTranslateSpeechResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingTranslateSpeechResponse.Size(m)
}
func (m *StreamingTranslateSpeechResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingTranslateSpeechResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingTranslateSpeechResponse proto.InternalMessageInfo

func (m *StreamingTranslateSpeechResponse) GetError() *status.Status {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *StreamingTranslateSpeechResponse) GetResult() *StreamingTranslateSpeechResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *StreamingTranslateSpeechResponse) GetSpeechEventType() StreamingTranslateSpeechResponse_SpeechEventType {
	if m != nil {
		return m.SpeechEventType
	}
	return StreamingTranslateSpeechResponse_SPEECH_EVENT_TYPE_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("google.cloud.mediatranslation.v1beta1.StreamingTranslateSpeechResponse_SpeechEventType", StreamingTranslateSpeechResponse_SpeechEventType_name, StreamingTranslateSpeechResponse_SpeechEventType_value)
	proto.RegisterType((*TranslateSpeechConfig)(nil), "google.cloud.mediatranslation.v1beta1.TranslateSpeechConfig")
	proto.RegisterType((*StreamingTranslateSpeechConfig)(nil), "google.cloud.mediatranslation.v1beta1.StreamingTranslateSpeechConfig")
	proto.RegisterType((*StreamingTranslateSpeechRequest)(nil), "google.cloud.mediatranslation.v1beta1.StreamingTranslateSpeechRequest")
	proto.RegisterType((*StreamingTranslateSpeechResult)(nil), "google.cloud.mediatranslation.v1beta1.StreamingTranslateSpeechResult")
	proto.RegisterType((*StreamingTranslateSpeechResult_TextTranslationResult)(nil), "google.cloud.mediatranslation.v1beta1.StreamingTranslateSpeechResult.TextTranslationResult")
	proto.RegisterType((*StreamingTranslateSpeechResponse)(nil), "google.cloud.mediatranslation.v1beta1.StreamingTranslateSpeechResponse")
}

func init() {
	proto.RegisterFile("google/cloud/mediatranslation/v1beta1/media_translation.proto", fileDescriptor_7b2f67d966c0ba6c)
}

var fileDescriptor_7b2f67d966c0ba6c = []byte{
	// 917 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x41, 0x6f, 0xe3, 0x44,
	0x1b, 0x8e, 0x93, 0xaf, 0xfd, 0xca, 0xb4, 0xbb, 0x6d, 0x07, 0xaa, 0x86, 0xa0, 0x6d, 0xbb, 0x91,
	0x2a, 0x15, 0xa4, 0xb5, 0x69, 0x11, 0x97, 0x00, 0x87, 0x26, 0xeb, 0x34, 0x95, 0x96, 0x10, 0x9c,
	0xb4, 0x08, 0xa8, 0x34, 0x9a, 0xda, 0x6f, 0x9d, 0x91, 0x1c, 0x8f, 0x99, 0x99, 0xa4, 0xbb, 0xfc,
	0x02, 0xee, 0x88, 0x03, 0x27, 0x0e, 0x48, 0x1c, 0xb8, 0x71, 0xe1, 0x47, 0xac, 0xc4, 0x01, 0x7e,
	0x01, 0x67, 0x7e, 0x01, 0x47, 0xe4, 0x19, 0xa7, 0x71, 0xdd, 0x74, 0xa9, 0x54, 0x8e, 0x9e, 0xe7,
	0x79, 0x9f, 0xf7, 0x79, 0xde, 0x99, 0xc9, 0x04, 0x7d, 0x14, 0x72, 0x1e, 0x46, 0xe0, 0xf8, 0x11,
	0x1f, 0x07, 0xce, 0x08, 0x02, 0x46, 0x95, 0xa0, 0xb1, 0x8c, 0xa8, 0x62, 0x3c, 0x76, 0x26, 0xfb,
	0xe7, 0xa0, 0xe8, 0xbe, 0x01, 0x48, 0x0e, 0xb1, 0x13, 0xc1, 0x15, 0xc7, 0xbb, 0xa6, 0xdc, 0xd6,
	0xe5, 0x76, 0xb1, 0xdc, 0xce, 0xca, 0x6b, 0xdb, 0x59, 0x17, 0x9a, 0x30, 0xe7, 0x82, 0x41, 0x14,
	0x90, 0x73, 0x18, 0xd2, 0x09, 0xe3, 0xc2, 0xe8, 0xd4, 0x36, 0x33, 0x82, 0x48, 0x7c, 0x47, 0x2a,
	0xaa, 0xc6, 0xb2, 0x00, 0xa4, 0x95, 0x7e, 0xc4, 0x20, 0x56, 0x06, 0xa8, 0xff, 0x56, 0x46, 0x1b,
	0x83, 0xac, 0x15, 0xf4, 0x13, 0x00, 0x7f, 0xd8, 0xe2, 0xf1, 0x05, 0x0b, 0xf1, 0x3b, 0xe8, 0x21,
	0x1d, 0x07, 0x8c, 0x13, 0x88, 0x7d, 0x1e, 0xb0, 0x38, 0xac, 0x5a, 0x3b, 0xd6, 0xde, 0x6b, 0xcd,
	0xca, 0x9f, 0x87, 0x65, 0xef, 0x81, 0x86, 0xdc, 0x0c, 0xc1, 0xef, 0xa3, 0x37, 0x24, 0x1f, 0x0b,
	0x1f, 0x48, 0x44, 0xe3, 0x70, 0x4c, 0x43, 0x20, 0x3e, 0x0f, 0xa0, 0x5a, 0x9e, 0x55, 0x60, 0x43,
	0x78, 0x96, 0xe1, 0x2d, 0x1e, 0x00, 0xee, 0xa2, 0xc7, 0x34, 0x52, 0x20, 0x62, 0xaa, 0xd8, 0x04,
	0xc8, 0x3c, 0x09, 0x59, 0x5d, 0xdc, 0xa9, 0x18, 0x0d, 0xcb, 0xdb, 0xca, 0xb1, 0xfb, 0x37, 0xe4,
	0x64, 0x6a, 0x43, 0x51, 0x11, 0x82, 0x2a, 0xd8, 0xa8, 0xe4, 0x6c, 0x18, 0xc2, 0x35, 0x1b, 0x0e,
	0x5a, 0x97, 0x74, 0x94, 0x44, 0x40, 0x04, 0x55, 0x40, 0x86, 0x20, 0xd4, 0xd7, 0xd5, 0xff, 0xed,
	0x58, 0x7b, 0x0b, 0xa6, 0xed, 0xaa, 0x41, 0x3d, 0xaa, 0xa0, 0x93, 0x62, 0xf8, 0x4d, 0xb4, 0x30,
	0xe2, 0x01, 0x44, 0xd5, 0x85, 0xa9, 0xb0, 0xe5, 0x99, 0x95, 0xfa, 0xaf, 0x16, 0xda, 0xea, 0x2b,
	0x01, 0x74, 0xc4, 0xe2, 0x70, 0xfe, 0x60, 0x03, 0xb4, 0x62, 0x06, 0xeb, 0xeb, 0x6f, 0x3d, 0xd6,
	0xe5, 0x83, 0x0f, 0xed, 0x3b, 0x9d, 0x01, 0x7b, 0xae, 0xa6, 0xc9, 0xb6, 0xac, 0x65, 0xb3, 0x2e,
	0x36, 0x5a, 0x93, 0x2c, 0x0e, 0x23, 0x20, 0x63, 0xa5, 0x40, 0xd0, 0xd8, 0x37, 0xdb, 0xb1, 0x34,
	0xcd, 0xa4, 0xc1, 0x93, 0x29, 0x56, 0xff, 0xdd, 0x42, 0xdb, 0xb7, 0x19, 0xf7, 0xe0, 0xab, 0x31,
	0x48, 0x85, 0x05, 0x5a, 0x93, 0x53, 0xca, 0x75, 0xf7, 0xee, 0x1d, 0xdd, 0xbf, 0x7a, 0x34, 0x9d,
	0x92, 0xb7, 0x7a, 0xd5, 0x20, 0xcb, 0xb1, 0x8b, 0x1e, 0x5c, 0x4d, 0x4b, 0x41, 0xac, 0x74, 0x88,
	0x95, 0x4e, 0xc9, 0x5b, 0x99, 0xa6, 0x4d, 0x57, 0x9b, 0xaf, 0xa3, 0xf5, 0x99, 0x35, 0x61, 0xfc,
	0xd6, 0x7f, 0xa8, 0xdc, 0xbe, 0x19, 0x1e, 0xc8, 0x71, 0xa4, 0xf0, 0x77, 0x16, 0xda, 0x54, 0xf0,
	0x5c, 0xe5, 0x2f, 0x25, 0x11, 0x1a, 0xcb, 0xa2, 0x7d, 0x79, 0xcf, 0x68, 0xa6, 0x91, 0x3d, 0x80,
	0xe7, 0x6a, 0x30, 0x2b, 0x33, 0xab, 0x9d, 0x92, 0xb7, 0xa1, 0xe6, 0x01, 0xf8, 0x00, 0x61, 0x01,
	0x3e, 0x0f, 0x63, 0x96, 0x77, 0x74, 0x75, 0x90, 0x2b, 0xde, 0x7a, 0x0e, 0x36, 0x35, 0xb5, 0x9f,
	0x2c, 0xb4, 0x31, 0xb7, 0x0d, 0xde, 0x45, 0xcb, 0x39, 0xcb, 0xb3, 0x8b, 0x5c, 0xf1, 0xf2, 0xeb,
	0x78, 0x0b, 0x2d, 0x31, 0x49, 0x2e, 0x58, 0x4c, 0xa3, 0xd9, 0x59, 0xa9, 0x78, 0xff, 0x67, 0xb2,
	0x9d, 0xae, 0xe1, 0x36, 0x7a, 0x14, 0x80, 0x02, 0x5f, 0x41, 0x30, 0xf7, 0xb2, 0xe6, 0xfd, 0xd5,
	0xa6, 0xcc, 0x9b, 0x17, 0xb5, 0xb9, 0x84, 0x16, 0x4d, 0xa0, 0xfa, 0xb7, 0x15, 0xb4, 0xf3, 0x8a,
	0xc1, 0x25, 0x3c, 0x96, 0x80, 0x9f, 0xa0, 0x05, 0x10, 0x82, 0x8b, 0x6c, 0x43, 0xf0, 0x74, 0x43,
	0x44, 0xe2, 0xdb, 0x7d, 0xfd, 0x2b, 0x67, 0x5a, 0x1a, 0x16, 0x3e, 0x9f, 0xaa, 0xeb, 0x0c, 0xf7,
	0x3f, 0x9b, 0x66, 0x86, 0xa6, 0x45, 0xa6, 0x8c, 0xbf, 0xb1, 0xd0, 0xba, 0xd4, 0x28, 0x81, 0x09,
	0xc4, 0x8a, 0xa8, 0x17, 0x89, 0x89, 0xff, 0xf0, 0xe0, 0xb3, 0xfb, 0xf7, 0xd3, 0xb9, 0x6d, 0xf3,
	0xe9, 0xa6, 0xfa, 0x83, 0x17, 0x09, 0x18, 0x07, 0xab, 0xf2, 0xfa, 0x6a, 0xfd, 0x53, 0xb4, 0x5a,
	0x20, 0xe2, 0xc7, 0xe8, 0x51, 0xbf, 0xe7, 0xba, 0xad, 0x0e, 0x71, 0x4f, 0xdd, 0xee, 0x80, 0x0c,
	0x3e, 0xef, 0xb9, 0xe4, 0xa4, 0xdb, 0xef, 0xb9, 0xad, 0xe3, 0xf6, 0xb1, 0xfb, 0x74, 0xad, 0x84,
	0xdf, 0x42, 0x9b, 0x6e, 0xf7, 0x29, 0xf9, 0xa4, 0x4d, 0xfa, 0xc7, 0xdd, 0xa3, 0x67, 0x2e, 0x39,
	0x19, 0x0c, 0x5c, 0xef, 0xb0, 0xdb, 0x72, 0xd7, 0xac, 0x83, 0x1f, 0xcb, 0xa8, 0x6a, 0x34, 0x73,
	0x47, 0xa9, 0x0f, 0x62, 0xc2, 0x7c, 0xc0, 0xbf, 0x58, 0xa8, 0x7a, 0x9b, 0x75, 0xdc, 0xbe, 0x77,
	0x76, 0x7d, 0x73, 0x6b, 0x47, 0xff, 0xd1, 0x0c, 0xeb, 0xa5, 0x3d, 0xeb, 0x5d, 0xab, 0xd6, 0x7f,
	0x79, 0xb8, 0x7d, 0x43, 0xc3, 0x74, 0xa0, 0x09, 0x93, 0xb6, 0xcf, 0x47, 0x7f, 0x1c, 0xda, 0x43,
	0xa5, 0x12, 0xd9, 0x70, 0x9c, 0xcb, 0xcb, 0xcb, 0x02, 0xe8, 0xd0, 0xb1, 0x1a, 0x9a, 0xc7, 0xfd,
	0x49, 0x12, 0x51, 0x75, 0xc1, 0xc5, 0xa8, 0xf9, 0x7d, 0x19, 0xbd, 0xed, 0xf3, 0xd1, 0xdd, 0x9c,
	0x36, 0x37, 0x3e, 0x4e, 0x91, 0xdc, 0x3c, 0x7b, 0xe9, 0xfb, 0xdb, 0xb3, 0xbe, 0x38, 0xc9, 0xea,
	0x43, 0x9e, 0xde, 0x25, 0x9b, 0x8b, 0xd0, 0x09, 0x21, 0xd6, 0xaf, 0xb3, 0x33, 0x73, 0xf1, 0x2f,
	0xff, 0x2c, 0x3e, 0x28, 0x02, 0x7f, 0x5b, 0xd6, 0xcf, 0xe5, 0xdd, 0x23, 0x23, 0xdd, 0xd2, 0xd6,
	0x8a, 0x06, 0xec, 0xd3, 0xfd, 0x66, 0x5a, 0xfb, 0x72, 0xca, 0x3b, 0xd3, 0xbc, 0xb3, 0x22, 0xef,
	0xec, 0xd4, 0xf4, 0xf8, 0xab, 0xbc, 0x67, 0x78, 0x8d, 0x86, 0x26, 0x36, 0x1a, 0x45, 0x66, 0xa3,
	0x91, 0x51, 0xcf, 0x17, 0xb5, 0xff, 0xf7, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xa7, 0xc3,
	0x21, 0x18, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SpeechTranslationServiceClient is the client API for SpeechTranslationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpeechTranslationServiceClient interface {
	// Performs bidirectional streaming speech translation: receive results while
	// sending audio. This method is only available via the gRPC API (not REST).
	StreamingTranslateSpeech(ctx context.Context, opts ...grpc.CallOption) (SpeechTranslationService_StreamingTranslateSpeechClient, error)
}

type speechTranslationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpeechTranslationServiceClient(cc grpc.ClientConnInterface) SpeechTranslationServiceClient {
	return &speechTranslationServiceClient{cc}
}

func (c *speechTranslationServiceClient) StreamingTranslateSpeech(ctx context.Context, opts ...grpc.CallOption) (SpeechTranslationService_StreamingTranslateSpeechClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SpeechTranslationService_serviceDesc.Streams[0], "/google.cloud.mediatranslation.v1beta1.SpeechTranslationService/StreamingTranslateSpeech", opts...)
	if err != nil {
		return nil, err
	}
	x := &speechTranslationServiceStreamingTranslateSpeechClient{stream}
	return x, nil
}

type SpeechTranslationService_StreamingTranslateSpeechClient interface {
	Send(*StreamingTranslateSpeechRequest) error
	Recv() (*StreamingTranslateSpeechResponse, error)
	grpc.ClientStream
}

type speechTranslationServiceStreamingTranslateSpeechClient struct {
	grpc.ClientStream
}

func (x *speechTranslationServiceStreamingTranslateSpeechClient) Send(m *StreamingTranslateSpeechRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *speechTranslationServiceStreamingTranslateSpeechClient) Recv() (*StreamingTranslateSpeechResponse, error) {
	m := new(StreamingTranslateSpeechResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpeechTranslationServiceServer is the server API for SpeechTranslationService service.
type SpeechTranslationServiceServer interface {
	// Performs bidirectional streaming speech translation: receive results while
	// sending audio. This method is only available via the gRPC API (not REST).
	StreamingTranslateSpeech(SpeechTranslationService_StreamingTranslateSpeechServer) error
}

// UnimplementedSpeechTranslationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSpeechTranslationServiceServer struct {
}

func (*UnimplementedSpeechTranslationServiceServer) StreamingTranslateSpeech(srv SpeechTranslationService_StreamingTranslateSpeechServer) error {
	return status1.Errorf(codes.Unimplemented, "method StreamingTranslateSpeech not implemented")
}

func RegisterSpeechTranslationServiceServer(s *grpc.Server, srv SpeechTranslationServiceServer) {
	s.RegisterService(&_SpeechTranslationService_serviceDesc, srv)
}

func _SpeechTranslationService_StreamingTranslateSpeech_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SpeechTranslationServiceServer).StreamingTranslateSpeech(&speechTranslationServiceStreamingTranslateSpeechServer{stream})
}

type SpeechTranslationService_StreamingTranslateSpeechServer interface {
	Send(*StreamingTranslateSpeechResponse) error
	Recv() (*StreamingTranslateSpeechRequest, error)
	grpc.ServerStream
}

type speechTranslationServiceStreamingTranslateSpeechServer struct {
	grpc.ServerStream
}

func (x *speechTranslationServiceStreamingTranslateSpeechServer) Send(m *StreamingTranslateSpeechResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *speechTranslationServiceStreamingTranslateSpeechServer) Recv() (*StreamingTranslateSpeechRequest, error) {
	m := new(StreamingTranslateSpeechRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SpeechTranslationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.mediatranslation.v1beta1.SpeechTranslationService",
	HandlerType: (*SpeechTranslationServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingTranslateSpeech",
			Handler:       _SpeechTranslationService_StreamingTranslateSpeech_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "google/cloud/mediatranslation/v1beta1/media_translation.proto",
}

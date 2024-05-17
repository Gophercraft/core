// DO NOT EDIT: this file was auto-generated by Gophercraft/protoss

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: bgs/low/pb/client/voice_types.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VoiceJoinType int32

const (
	VoiceJoinType_VOICE_JOIN_NORMAL VoiceJoinType = 0
	VoiceJoinType_VOICE_JOIN_MUTED  VoiceJoinType = 1
)

// Enum value maps for VoiceJoinType.
var (
	VoiceJoinType_name = map[int32]string{
		0: "VOICE_JOIN_NORMAL",
		1: "VOICE_JOIN_MUTED",
	}
	VoiceJoinType_value = map[string]int32{
		"VOICE_JOIN_NORMAL": 0,
		"VOICE_JOIN_MUTED":  1,
	}
)

func (x VoiceJoinType) Enum() *VoiceJoinType {
	p := new(VoiceJoinType)
	*p = x
	return p
}

func (x VoiceJoinType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VoiceJoinType) Descriptor() protoreflect.EnumDescriptor {
	return file_bgs_low_pb_client_voice_types_proto_enumTypes[0].Descriptor()
}

func (VoiceJoinType) Type() protoreflect.EnumType {
	return &file_bgs_low_pb_client_voice_types_proto_enumTypes[0]
}

func (x VoiceJoinType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *VoiceJoinType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = VoiceJoinType(num)
	return nil
}

// Deprecated: Use VoiceJoinType.Descriptor instead.
func (VoiceJoinType) EnumDescriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_voice_types_proto_rawDescGZIP(), []int{0}
}

type VoiceMuteReason int32

const (
	VoiceMuteReason_VOICE_MUTE_REASON_NONE                         VoiceMuteReason = 0
	VoiceMuteReason_VOICE_MUTE_REASON_PARENTAL_CONTROL_LISTEN_ONLY VoiceMuteReason = 1
	VoiceMuteReason_VOICE_MUTE_REASON_REQUESTED                    VoiceMuteReason = 2
	VoiceMuteReason_VOICE_MUTE_REASON_SQUELCHED                    VoiceMuteReason = 3
)

// Enum value maps for VoiceMuteReason.
var (
	VoiceMuteReason_name = map[int32]string{
		0: "VOICE_MUTE_REASON_NONE",
		1: "VOICE_MUTE_REASON_PARENTAL_CONTROL_LISTEN_ONLY",
		2: "VOICE_MUTE_REASON_REQUESTED",
		3: "VOICE_MUTE_REASON_SQUELCHED",
	}
	VoiceMuteReason_value = map[string]int32{
		"VOICE_MUTE_REASON_NONE":                         0,
		"VOICE_MUTE_REASON_PARENTAL_CONTROL_LISTEN_ONLY": 1,
		"VOICE_MUTE_REASON_REQUESTED":                    2,
		"VOICE_MUTE_REASON_SQUELCHED":                    3,
	}
)

func (x VoiceMuteReason) Enum() *VoiceMuteReason {
	p := new(VoiceMuteReason)
	*p = x
	return p
}

func (x VoiceMuteReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VoiceMuteReason) Descriptor() protoreflect.EnumDescriptor {
	return file_bgs_low_pb_client_voice_types_proto_enumTypes[1].Descriptor()
}

func (VoiceMuteReason) Type() protoreflect.EnumType {
	return &file_bgs_low_pb_client_voice_types_proto_enumTypes[1]
}

func (x VoiceMuteReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *VoiceMuteReason) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = VoiceMuteReason(num)
	return nil
}

// Deprecated: Use VoiceMuteReason.Descriptor instead.
func (VoiceMuteReason) EnumDescriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_voice_types_proto_rawDescGZIP(), []int{1}
}

type VoiceProviderVersion int32

const (
	VoiceProviderVersion_VOICE_PROVIDER_V4 VoiceProviderVersion = 0
	VoiceProviderVersion_VOICE_PROVIDER_V5 VoiceProviderVersion = 1
)

// Enum value maps for VoiceProviderVersion.
var (
	VoiceProviderVersion_name = map[int32]string{
		0: "VOICE_PROVIDER_V4",
		1: "VOICE_PROVIDER_V5",
	}
	VoiceProviderVersion_value = map[string]int32{
		"VOICE_PROVIDER_V4": 0,
		"VOICE_PROVIDER_V5": 1,
	}
)

func (x VoiceProviderVersion) Enum() *VoiceProviderVersion {
	p := new(VoiceProviderVersion)
	*p = x
	return p
}

func (x VoiceProviderVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VoiceProviderVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_bgs_low_pb_client_voice_types_proto_enumTypes[2].Descriptor()
}

func (VoiceProviderVersion) Type() protoreflect.EnumType {
	return &file_bgs_low_pb_client_voice_types_proto_enumTypes[2]
}

func (x VoiceProviderVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *VoiceProviderVersion) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = VoiceProviderVersion(num)
	return nil
}

// Deprecated: Use VoiceProviderVersion.Descriptor instead.
func (VoiceProviderVersion) EnumDescriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_voice_types_proto_rawDescGZIP(), []int{2}
}

type VoiceCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoiceId    *string          `protobuf:"bytes,1,opt,name=voice_id,json=voiceId" json:"voice_id,omitempty"`
	Token      *string          `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Url        *string          `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	JoinType   *VoiceJoinType   `protobuf:"varint,4,opt,name=join_type,json=joinType,enum=bgs.protocol.VoiceJoinType" json:"join_type,omitempty"`
	MuteReason *VoiceMuteReason `protobuf:"varint,5,opt,name=mute_reason,json=muteReason,enum=bgs.protocol.VoiceMuteReason" json:"mute_reason,omitempty"`
}

func (x *VoiceCredentials) Reset() {
	*x = VoiceCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_voice_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoiceCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceCredentials) ProtoMessage() {}

func (x *VoiceCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_voice_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoiceCredentials.ProtoReflect.Descriptor instead.
func (*VoiceCredentials) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_voice_types_proto_rawDescGZIP(), []int{0}
}

func (x *VoiceCredentials) GetVoiceId() string {
	if x != nil && x.VoiceId != nil {
		return *x.VoiceId
	}
	return ""
}

func (x *VoiceCredentials) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *VoiceCredentials) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *VoiceCredentials) GetJoinType() VoiceJoinType {
	if x != nil && x.JoinType != nil {
		return *x.JoinType
	}
	return VoiceJoinType_VOICE_JOIN_NORMAL
}

func (x *VoiceCredentials) GetMuteReason() VoiceMuteReason {
	if x != nil && x.MuteReason != nil {
		return *x.MuteReason
	}
	return VoiceMuteReason_VOICE_MUTE_REASON_NONE
}

var File_bgs_low_pb_client_voice_types_proto protoreflect.FileDescriptor

var file_bgs_low_pb_client_voice_types_proto_rawDesc = []byte{
	0x0a, 0x23, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x22, 0xcf, 0x01, 0x0a, 0x10, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x38, 0x0a, 0x09, 0x6a,
	0x6f, 0x69, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6a, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x6d, 0x75, 0x74, 0x65, 0x5f, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x62, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x4d,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x0a, 0x6d, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2a, 0x3c, 0x0a, 0x0d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x4a, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f,
	0x4a, 0x4f, 0x49, 0x4e, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x5f, 0x4d, 0x55, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x2a, 0xa3, 0x01, 0x0a, 0x0f, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x4d, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x56, 0x4f, 0x49, 0x43, 0x45,
	0x5f, 0x4d, 0x55, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x32, 0x0a, 0x2e, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x4d, 0x55, 0x54,
	0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x41,
	0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x4e,
	0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x56, 0x4f, 0x49, 0x43, 0x45,
	0x5f, 0x4d, 0x55, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x56, 0x4f, 0x49, 0x43,
	0x45, 0x5f, 0x4d, 0x55, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x53, 0x51,
	0x55, 0x45, 0x4c, 0x43, 0x48, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x44, 0x0a, 0x14, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49,
	0x44, 0x45, 0x52, 0x5f, 0x56, 0x34, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x4f, 0x49, 0x43,
	0x45, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x56, 0x35, 0x10, 0x01, 0x42,
	0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f,
	0x70, 0x68, 0x65, 0x72, 0x63, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62,
	0x6e, 0x65, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c,
}

var (
	file_bgs_low_pb_client_voice_types_proto_rawDescOnce sync.Once
	file_bgs_low_pb_client_voice_types_proto_rawDescData = file_bgs_low_pb_client_voice_types_proto_rawDesc
)

func file_bgs_low_pb_client_voice_types_proto_rawDescGZIP() []byte {
	file_bgs_low_pb_client_voice_types_proto_rawDescOnce.Do(func() {
		file_bgs_low_pb_client_voice_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgs_low_pb_client_voice_types_proto_rawDescData)
	})
	return file_bgs_low_pb_client_voice_types_proto_rawDescData
}

var file_bgs_low_pb_client_voice_types_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_bgs_low_pb_client_voice_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bgs_low_pb_client_voice_types_proto_goTypes = []interface{}{
	(VoiceJoinType)(0),        // 0: bgs.protocol.VoiceJoinType
	(VoiceMuteReason)(0),      // 1: bgs.protocol.VoiceMuteReason
	(VoiceProviderVersion)(0), // 2: bgs.protocol.VoiceProviderVersion
	(*VoiceCredentials)(nil),  // 3: bgs.protocol.VoiceCredentials
}
var file_bgs_low_pb_client_voice_types_proto_depIdxs = []int32{
	0, // 0: bgs.protocol.VoiceCredentials.join_type:type_name -> bgs.protocol.VoiceJoinType
	1, // 1: bgs.protocol.VoiceCredentials.mute_reason:type_name -> bgs.protocol.VoiceMuteReason
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bgs_low_pb_client_voice_types_proto_init() }
func file_bgs_low_pb_client_voice_types_proto_init() {
	if File_bgs_low_pb_client_voice_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bgs_low_pb_client_voice_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoiceCredentials); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bgs_low_pb_client_voice_types_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bgs_low_pb_client_voice_types_proto_goTypes,
		DependencyIndexes: file_bgs_low_pb_client_voice_types_proto_depIdxs,
		EnumInfos:         file_bgs_low_pb_client_voice_types_proto_enumTypes,
		MessageInfos:      file_bgs_low_pb_client_voice_types_proto_msgTypes,
	}.Build()
	File_bgs_low_pb_client_voice_types_proto = out.File
	file_bgs_low_pb_client_voice_types_proto_rawDesc = nil
	file_bgs_low_pb_client_voice_types_proto_goTypes = nil
	file_bgs_low_pb_client_voice_types_proto_depIdxs = nil
}
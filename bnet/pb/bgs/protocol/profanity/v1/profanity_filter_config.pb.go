// DO NOT EDIT: this file was auto-generated by Gophercraft/protoss

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: bgs/low/pb/client/profanity_filter_config.proto

package v1

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

type WordFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  *string `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	Regex *string `protobuf:"bytes,2,req,name=regex" json:"regex,omitempty"`
}

func (x *WordFilter) Reset() {
	*x = WordFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_profanity_filter_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WordFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WordFilter) ProtoMessage() {}

func (x *WordFilter) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_profanity_filter_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WordFilter.ProtoReflect.Descriptor instead.
func (*WordFilter) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_profanity_filter_config_proto_rawDescGZIP(), []int{0}
}

func (x *WordFilter) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *WordFilter) GetRegex() string {
	if x != nil && x.Regex != nil {
		return *x.Regex
	}
	return ""
}

type WordFilters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters []*WordFilter `protobuf:"bytes,1,rep,name=filters" json:"filters,omitempty"`
}

func (x *WordFilters) Reset() {
	*x = WordFilters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgs_low_pb_client_profanity_filter_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WordFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WordFilters) ProtoMessage() {}

func (x *WordFilters) ProtoReflect() protoreflect.Message {
	mi := &file_bgs_low_pb_client_profanity_filter_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WordFilters.ProtoReflect.Descriptor instead.
func (*WordFilters) Descriptor() ([]byte, []int) {
	return file_bgs_low_pb_client_profanity_filter_config_proto_rawDescGZIP(), []int{1}
}

func (x *WordFilters) GetFilters() []*WordFilter {
	if x != nil {
		return x.Filters
	}
	return nil
}

var File_bgs_low_pb_client_profanity_filter_config_proto protoreflect.FileDescriptor

var file_bgs_low_pb_client_profanity_filter_config_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x62, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x36, 0x0a, 0x0a,
	0x57, 0x6f, 0x72, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x72,
	0x65, 0x67, 0x65, 0x78, 0x22, 0x4e, 0x0a, 0x0b, 0x57, 0x6f, 0x72, 0x64, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x3f, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x6f, 0x72, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x63, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x62, 0x6e, 0x65, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x67, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x61, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x76, 0x31,
}

var (
	file_bgs_low_pb_client_profanity_filter_config_proto_rawDescOnce sync.Once
	file_bgs_low_pb_client_profanity_filter_config_proto_rawDescData = file_bgs_low_pb_client_profanity_filter_config_proto_rawDesc
)

func file_bgs_low_pb_client_profanity_filter_config_proto_rawDescGZIP() []byte {
	file_bgs_low_pb_client_profanity_filter_config_proto_rawDescOnce.Do(func() {
		file_bgs_low_pb_client_profanity_filter_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgs_low_pb_client_profanity_filter_config_proto_rawDescData)
	})
	return file_bgs_low_pb_client_profanity_filter_config_proto_rawDescData
}

var file_bgs_low_pb_client_profanity_filter_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bgs_low_pb_client_profanity_filter_config_proto_goTypes = []interface{}{
	(*WordFilter)(nil),  // 0: bgs.protocol.profanity.v1.WordFilter
	(*WordFilters)(nil), // 1: bgs.protocol.profanity.v1.WordFilters
}
var file_bgs_low_pb_client_profanity_filter_config_proto_depIdxs = []int32{
	0, // 0: bgs.protocol.profanity.v1.WordFilters.filters:type_name -> bgs.protocol.profanity.v1.WordFilter
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bgs_low_pb_client_profanity_filter_config_proto_init() }
func file_bgs_low_pb_client_profanity_filter_config_proto_init() {
	if File_bgs_low_pb_client_profanity_filter_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bgs_low_pb_client_profanity_filter_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WordFilter); i {
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
		file_bgs_low_pb_client_profanity_filter_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WordFilters); i {
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
			RawDescriptor: file_bgs_low_pb_client_profanity_filter_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bgs_low_pb_client_profanity_filter_config_proto_goTypes,
		DependencyIndexes: file_bgs_low_pb_client_profanity_filter_config_proto_depIdxs,
		MessageInfos:      file_bgs_low_pb_client_profanity_filter_config_proto_msgTypes,
	}.Build()
	File_bgs_low_pb_client_profanity_filter_config_proto = out.File
	file_bgs_low_pb_client_profanity_filter_config_proto_rawDesc = nil
	file_bgs_low_pb_client_profanity_filter_config_proto_goTypes = nil
	file_bgs_low_pb_client_profanity_filter_config_proto_depIdxs = nil
}
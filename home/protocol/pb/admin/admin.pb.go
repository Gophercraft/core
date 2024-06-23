// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: admin.proto

package admin

import (
	_ "github.com/Gophercraft/core/home/protocol/pb/auth"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type AccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Credential string `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	Id         uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AccountRequest) Reset() {
	*x = AccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountRequest) ProtoMessage() {}

func (x *AccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountRequest.ProtoReflect.Descriptor instead.
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{0}
}

func (x *AccountRequest) GetCredential() string {
	if x != nil {
		return x.Credential
	}
	return ""
}

func (x *AccountRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SuspendAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Credential         string `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	Id                 uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	SuspensionDuration int64  `protobuf:"varint,3,opt,name=suspension_duration,json=suspensionDuration,proto3" json:"suspension_duration,omitempty"`
}

func (x *SuspendAccountRequest) Reset() {
	*x = SuspendAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuspendAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuspendAccountRequest) ProtoMessage() {}

func (x *SuspendAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuspendAccountRequest.ProtoReflect.Descriptor instead.
func (*SuspendAccountRequest) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{1}
}

func (x *SuspendAccountRequest) GetCredential() string {
	if x != nil {
		return x.Credential
	}
	return ""
}

func (x *SuspendAccountRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SuspendAccountRequest) GetSuspensionDuration() int64 {
	if x != nil {
		return x.SuspensionDuration
	}
	return 0
}

type BanStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Banned        bool                 `protobuf:"varint,2,opt,name=banned,proto3" json:"banned,omitempty"`
	Suspended     bool                 `protobuf:"varint,3,opt,name=suspended,proto3" json:"suspended,omitempty"`
	UnsuspendTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=unsuspend_time,json=unsuspendTime,proto3" json:"unsuspend_time,omitempty"`
}

func (x *BanStatus) Reset() {
	*x = BanStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BanStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BanStatus) ProtoMessage() {}

func (x *BanStatus) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BanStatus.ProtoReflect.Descriptor instead.
func (*BanStatus) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{2}
}

func (x *BanStatus) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BanStatus) GetBanned() bool {
	if x != nil {
		return x.Banned
	}
	return false
}

func (x *BanStatus) GetSuspended() bool {
	if x != nil {
		return x.Suspended
	}
	return false
}

func (x *BanStatus) GetUnsuspendTime() *timestamp.Timestamp {
	if x != nil {
		return x.UnsuspendTime
	}
	return nil
}

type LockStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locked bool `protobuf:"varint,1,opt,name=locked,proto3" json:"locked,omitempty"`
}

func (x *LockStatus) Reset() {
	*x = LockStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockStatus) ProtoMessage() {}

func (x *LockStatus) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockStatus.ProtoReflect.Descriptor instead.
func (*LockStatus) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{3}
}

func (x *LockStatus) GetLocked() bool {
	if x != nil {
		return x.Locked
	}
	return false
}

type FileHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name like "backup-____.zip"
	// cannot contain any slashes
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FileHeader) Reset() {
	*x = FileHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileHeader) ProtoMessage() {}

func (x *FileHeader) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileHeader.ProtoReflect.Descriptor instead.
func (*FileHeader) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{4}
}

func (x *FileHeader) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileHeader) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FileChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{5}
}

func (x *FileChunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type FileDownload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*FileDownload_Header
	//	*FileDownload_Chunk
	Data isFileDownload_Data `protobuf_oneof:"data"`
}

func (x *FileDownload) Reset() {
	*x = FileDownload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDownload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDownload) ProtoMessage() {}

func (x *FileDownload) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDownload.ProtoReflect.Descriptor instead.
func (*FileDownload) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{6}
}

func (m *FileDownload) GetData() isFileDownload_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *FileDownload) GetHeader() *FileHeader {
	if x, ok := x.GetData().(*FileDownload_Header); ok {
		return x.Header
	}
	return nil
}

func (x *FileDownload) GetChunk() *FileChunk {
	if x, ok := x.GetData().(*FileDownload_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isFileDownload_Data interface {
	isFileDownload_Data()
}

type FileDownload_Header struct {
	Header *FileHeader `protobuf:"bytes,1,opt,name=header,proto3,oneof"`
}

type FileDownload_Chunk struct {
	Chunk *FileChunk `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*FileDownload_Header) isFileDownload_Data() {}

func (*FileDownload_Chunk) isFileDownload_Data() {}

type TakeBackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TakeBackupRequest) Reset() {
	*x = TakeBackupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeBackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeBackupRequest) ProtoMessage() {}

func (x *TakeBackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeBackupRequest.ProtoReflect.Descriptor instead.
func (*TakeBackupRequest) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{7}
}

var File_admin_proto protoreflect.FileDescriptor

var file_admin_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x40, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x78, 0x0a, 0x15, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x13,
	0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x73, 0x75, 0x73, 0x70, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x94, 0x01,
	0x0a, 0x09, 0x42, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x61, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x62, 0x61, 0x6e,
	0x6e, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x12, 0x41, 0x0a, 0x0e, 0x75, 0x6e, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x75, 0x6e, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x0a, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x22, 0x34, 0x0a, 0x0a, 0x46, 0x69,
	0x6c, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x22, 0x1f, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x6d, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48,
	0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x13, 0x0a, 0x11, 0x54, 0x61, 0x6b, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xc5, 0x05, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x54, 0x61, 0x6b, 0x65, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x12, 0x18, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x6b,
	0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x00, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x0a, 0x42, 0x61, 0x6e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x42, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x0c, 0x55, 0x6e, 0x62, 0x61, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x42, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0e, 0x53,
	0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x42, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x10, 0x55, 0x6e, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x42, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x0b, 0x4c, 0x6f, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x63,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0d, 0x55, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x47, 0x61, 0x6d,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x42, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x10, 0x55, 0x6e, 0x62, 0x61, 0x6e, 0x47, 0x61, 0x6d, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x42, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x12, 0x46, 0x0a, 0x12, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x47, 0x61, 0x6d,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x42,
	0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x14, 0x55, 0x6e,
	0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x42, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x70, 0x68,
	0x65, 0x72, 0x63, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x6f, 0x6d,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_proto_rawDescOnce sync.Once
	file_admin_proto_rawDescData = file_admin_proto_rawDesc
)

func file_admin_proto_rawDescGZIP() []byte {
	file_admin_proto_rawDescOnce.Do(func() {
		file_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_proto_rawDescData)
	})
	return file_admin_proto_rawDescData
}

var file_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_admin_proto_goTypes = []interface{}{
	(*AccountRequest)(nil),        // 0: admin.AccountRequest
	(*SuspendAccountRequest)(nil), // 1: admin.SuspendAccountRequest
	(*BanStatus)(nil),             // 2: admin.BanStatus
	(*LockStatus)(nil),            // 3: admin.LockStatus
	(*FileHeader)(nil),            // 4: admin.FileHeader
	(*FileChunk)(nil),             // 5: admin.FileChunk
	(*FileDownload)(nil),          // 6: admin.FileDownload
	(*TakeBackupRequest)(nil),     // 7: admin.TakeBackupRequest
	(*timestamp.Timestamp)(nil),   // 8: google.protobuf.Timestamp
}
var file_admin_proto_depIdxs = []int32{
	8,  // 0: admin.BanStatus.unsuspend_time:type_name -> google.protobuf.Timestamp
	4,  // 1: admin.FileDownload.header:type_name -> admin.FileHeader
	5,  // 2: admin.FileDownload.chunk:type_name -> admin.FileChunk
	7,  // 3: admin.AdminService.TakeBackup:input_type -> admin.TakeBackupRequest
	0,  // 4: admin.AdminService.BanAccount:input_type -> admin.AccountRequest
	0,  // 5: admin.AdminService.UnbanAccount:input_type -> admin.AccountRequest
	1,  // 6: admin.AdminService.SuspendAccount:input_type -> admin.SuspendAccountRequest
	0,  // 7: admin.AdminService.UnsuspendAccount:input_type -> admin.AccountRequest
	0,  // 8: admin.AdminService.LockAccount:input_type -> admin.AccountRequest
	0,  // 9: admin.AdminService.UnlockAccount:input_type -> admin.AccountRequest
	0,  // 10: admin.AdminService.BanGameAccount:input_type -> admin.AccountRequest
	0,  // 11: admin.AdminService.UnbanGameAccount:input_type -> admin.AccountRequest
	1,  // 12: admin.AdminService.SuspendGameAccount:input_type -> admin.SuspendAccountRequest
	0,  // 13: admin.AdminService.UnsuspendGameAccount:input_type -> admin.AccountRequest
	6,  // 14: admin.AdminService.TakeBackup:output_type -> admin.FileDownload
	2,  // 15: admin.AdminService.BanAccount:output_type -> admin.BanStatus
	2,  // 16: admin.AdminService.UnbanAccount:output_type -> admin.BanStatus
	2,  // 17: admin.AdminService.SuspendAccount:output_type -> admin.BanStatus
	2,  // 18: admin.AdminService.UnsuspendAccount:output_type -> admin.BanStatus
	3,  // 19: admin.AdminService.LockAccount:output_type -> admin.LockStatus
	3,  // 20: admin.AdminService.UnlockAccount:output_type -> admin.LockStatus
	2,  // 21: admin.AdminService.BanGameAccount:output_type -> admin.BanStatus
	2,  // 22: admin.AdminService.UnbanGameAccount:output_type -> admin.BanStatus
	2,  // 23: admin.AdminService.SuspendGameAccount:output_type -> admin.BanStatus
	2,  // 24: admin.AdminService.UnsuspendGameAccount:output_type -> admin.BanStatus
	14, // [14:25] is the sub-list for method output_type
	3,  // [3:14] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_admin_proto_init() }
func file_admin_proto_init() {
	if File_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountRequest); i {
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
		file_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuspendAccountRequest); i {
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
		file_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BanStatus); i {
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
		file_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockStatus); i {
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
		file_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileHeader); i {
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
		file_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileChunk); i {
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
		file_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDownload); i {
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
		file_admin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeBackupRequest); i {
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
	file_admin_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*FileDownload_Header)(nil),
		(*FileDownload_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_proto_goTypes,
		DependencyIndexes: file_admin_proto_depIdxs,
		MessageInfos:      file_admin_proto_msgTypes,
	}.Build()
	File_admin_proto = out.File
	file_admin_proto_rawDesc = nil
	file_admin_proto_goTypes = nil
	file_admin_proto_depIdxs = nil
}

// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: github.com/Gophercraft/core/bnet/public_protos/RealmList.proto

package realmlist

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type RealmListTicketIdentity struct {
	GameAccountID        *uint32  `protobuf:"fixed32,1,req,name=gameAccountID" json:"gameAccountID,omitempty"`
	GameAccountRegion    *uint32  `protobuf:"varint,2,req,name=gameAccountRegion" json:"gameAccountRegion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RealmListTicketIdentity) Reset()         { *m = RealmListTicketIdentity{} }
func (m *RealmListTicketIdentity) String() string { return proto.CompactTextString(m) }
func (*RealmListTicketIdentity) ProtoMessage()    {}
func (*RealmListTicketIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_366349c60ad91e0b, []int{0}
}

func (m *RealmListTicketIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealmListTicketIdentity.Unmarshal(m, b)
}
func (m *RealmListTicketIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealmListTicketIdentity.Marshal(b, m, deterministic)
}
func (m *RealmListTicketIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealmListTicketIdentity.Merge(m, src)
}
func (m *RealmListTicketIdentity) XXX_Size() int {
	return xxx_messageInfo_RealmListTicketIdentity.Size(m)
}
func (m *RealmListTicketIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_RealmListTicketIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_RealmListTicketIdentity proto.InternalMessageInfo

func (m *RealmListTicketIdentity) GetGameAccountID() uint32 {
	if m != nil && m.GameAccountID != nil {
		return *m.GameAccountID
	}
	return 0
}

func (m *RealmListTicketIdentity) GetGameAccountRegion() uint32 {
	if m != nil && m.GameAccountRegion != nil {
		return *m.GameAccountRegion
	}
	return 0
}

type ClientVersion struct {
	VersionMajor         *uint32  `protobuf:"varint,1,req,name=versionMajor" json:"versionMajor,omitempty"`
	VersionMinor         *uint32  `protobuf:"varint,2,req,name=versionMinor" json:"versionMinor,omitempty"`
	VersionRevision      *uint32  `protobuf:"varint,3,req,name=versionRevision" json:"versionRevision,omitempty"`
	VersionBuild         *uint32  `protobuf:"fixed32,4,req,name=versionBuild" json:"versionBuild,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientVersion) Reset()         { *m = ClientVersion{} }
func (m *ClientVersion) String() string { return proto.CompactTextString(m) }
func (*ClientVersion) ProtoMessage()    {}
func (*ClientVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_366349c60ad91e0b, []int{1}
}

func (m *ClientVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientVersion.Unmarshal(m, b)
}
func (m *ClientVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientVersion.Marshal(b, m, deterministic)
}
func (m *ClientVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientVersion.Merge(m, src)
}
func (m *ClientVersion) XXX_Size() int {
	return xxx_messageInfo_ClientVersion.Size(m)
}
func (m *ClientVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientVersion.DiscardUnknown(m)
}

var xxx_messageInfo_ClientVersion proto.InternalMessageInfo

func (m *ClientVersion) GetVersionMajor() uint32 {
	if m != nil && m.VersionMajor != nil {
		return *m.VersionMajor
	}
	return 0
}

func (m *ClientVersion) GetVersionMinor() uint32 {
	if m != nil && m.VersionMinor != nil {
		return *m.VersionMinor
	}
	return 0
}

func (m *ClientVersion) GetVersionRevision() uint32 {
	if m != nil && m.VersionRevision != nil {
		return *m.VersionRevision
	}
	return 0
}

func (m *ClientVersion) GetVersionBuild() uint32 {
	if m != nil && m.VersionBuild != nil {
		return *m.VersionBuild
	}
	return 0
}

type ClientInformation struct {
	Platform             *uint32        `protobuf:"fixed32,1,req,name=platform" json:"platform,omitempty"`
	BuildVariant         *string        `protobuf:"bytes,2,req,name=buildVariant" json:"buildVariant,omitempty"`
	Type                 *uint32        `protobuf:"fixed32,3,req,name=type" json:"type,omitempty"`
	TimeZone             *string        `protobuf:"bytes,4,req,name=timeZone" json:"timeZone,omitempty"`
	CurrentTime          *uint32        `protobuf:"varint,5,req,name=currentTime" json:"currentTime,omitempty"`
	TextLocale           *uint32        `protobuf:"fixed32,6,req,name=textLocale" json:"textLocale,omitempty"`
	AudioLocale          *uint32        `protobuf:"fixed32,7,req,name=audioLocale" json:"audioLocale,omitempty"`
	VersionDataBuild     *uint32        `protobuf:"fixed32,8,req,name=versionDataBuild" json:"versionDataBuild,omitempty"`
	Version              *ClientVersion `protobuf:"bytes,9,req,name=version" json:"version,omitempty"`
	Secret               []byte         `protobuf:"bytes,10,req,name=secret" json:"secret,omitempty"`
	ClientArch           *uint32        `protobuf:"fixed32,11,req,name=clientArch" json:"clientArch,omitempty"`
	SystemVersion        *string        `protobuf:"bytes,12,req,name=systemVersion" json:"systemVersion,omitempty"`
	PlatformType         *uint32        `protobuf:"fixed32,13,req,name=platformType" json:"platformType,omitempty"`
	SystemArch           *uint32        `protobuf:"fixed32,14,req,name=systemArch" json:"systemArch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ClientInformation) Reset()         { *m = ClientInformation{} }
func (m *ClientInformation) String() string { return proto.CompactTextString(m) }
func (*ClientInformation) ProtoMessage()    {}
func (*ClientInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_366349c60ad91e0b, []int{2}
}

func (m *ClientInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInformation.Unmarshal(m, b)
}
func (m *ClientInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInformation.Marshal(b, m, deterministic)
}
func (m *ClientInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInformation.Merge(m, src)
}
func (m *ClientInformation) XXX_Size() int {
	return xxx_messageInfo_ClientInformation.Size(m)
}
func (m *ClientInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInformation.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInformation proto.InternalMessageInfo

func (m *ClientInformation) GetPlatform() uint32 {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return 0
}

func (m *ClientInformation) GetBuildVariant() string {
	if m != nil && m.BuildVariant != nil {
		return *m.BuildVariant
	}
	return ""
}

func (m *ClientInformation) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *ClientInformation) GetTimeZone() string {
	if m != nil && m.TimeZone != nil {
		return *m.TimeZone
	}
	return ""
}

func (m *ClientInformation) GetCurrentTime() uint32 {
	if m != nil && m.CurrentTime != nil {
		return *m.CurrentTime
	}
	return 0
}

func (m *ClientInformation) GetTextLocale() uint32 {
	if m != nil && m.TextLocale != nil {
		return *m.TextLocale
	}
	return 0
}

func (m *ClientInformation) GetAudioLocale() uint32 {
	if m != nil && m.AudioLocale != nil {
		return *m.AudioLocale
	}
	return 0
}

func (m *ClientInformation) GetVersionDataBuild() uint32 {
	if m != nil && m.VersionDataBuild != nil {
		return *m.VersionDataBuild
	}
	return 0
}

func (m *ClientInformation) GetVersion() *ClientVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ClientInformation) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *ClientInformation) GetClientArch() uint32 {
	if m != nil && m.ClientArch != nil {
		return *m.ClientArch
	}
	return 0
}

func (m *ClientInformation) GetSystemVersion() string {
	if m != nil && m.SystemVersion != nil {
		return *m.SystemVersion
	}
	return ""
}

func (m *ClientInformation) GetPlatformType() uint32 {
	if m != nil && m.PlatformType != nil {
		return *m.PlatformType
	}
	return 0
}

func (m *ClientInformation) GetSystemArch() uint32 {
	if m != nil && m.SystemArch != nil {
		return *m.SystemArch
	}
	return 0
}

type RealmListTicketClientInformation struct {
	Info                 *ClientInformation `protobuf:"bytes,1,req,name=info" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RealmListTicketClientInformation) Reset()         { *m = RealmListTicketClientInformation{} }
func (m *RealmListTicketClientInformation) String() string { return proto.CompactTextString(m) }
func (*RealmListTicketClientInformation) ProtoMessage()    {}
func (*RealmListTicketClientInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_366349c60ad91e0b, []int{3}
}

func (m *RealmListTicketClientInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealmListTicketClientInformation.Unmarshal(m, b)
}
func (m *RealmListTicketClientInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealmListTicketClientInformation.Marshal(b, m, deterministic)
}
func (m *RealmListTicketClientInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealmListTicketClientInformation.Merge(m, src)
}
func (m *RealmListTicketClientInformation) XXX_Size() int {
	return xxx_messageInfo_RealmListTicketClientInformation.Size(m)
}
func (m *RealmListTicketClientInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_RealmListTicketClientInformation.DiscardUnknown(m)
}

var xxx_messageInfo_RealmListTicketClientInformation proto.InternalMessageInfo

func (m *RealmListTicketClientInformation) GetInfo() *ClientInformation {
	if m != nil {
		return m.Info
	}
	return nil
}

type RealmCharacterCountEntry struct {
	WowRealmAddress      *uint32  `protobuf:"fixed32,1,req,name=wowRealmAddress" json:"wowRealmAddress,omitempty"`
	Count                *uint32  `protobuf:"varint,2,req,name=count" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RealmCharacterCountEntry) Reset()         { *m = RealmCharacterCountEntry{} }
func (m *RealmCharacterCountEntry) String() string { return proto.CompactTextString(m) }
func (*RealmCharacterCountEntry) ProtoMessage()    {}
func (*RealmCharacterCountEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_366349c60ad91e0b, []int{4}
}

func (m *RealmCharacterCountEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealmCharacterCountEntry.Unmarshal(m, b)
}
func (m *RealmCharacterCountEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealmCharacterCountEntry.Marshal(b, m, deterministic)
}
func (m *RealmCharacterCountEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealmCharacterCountEntry.Merge(m, src)
}
func (m *RealmCharacterCountEntry) XXX_Size() int {
	return xxx_messageInfo_RealmCharacterCountEntry.Size(m)
}
func (m *RealmCharacterCountEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_RealmCharacterCountEntry.DiscardUnknown(m)
}

var xxx_messageInfo_RealmCharacterCountEntry proto.InternalMessageInfo

func (m *RealmCharacterCountEntry) GetWowRealmAddress() uint32 {
	if m != nil && m.WowRealmAddress != nil {
		return *m.WowRealmAddress
	}
	return 0
}

func (m *RealmCharacterCountEntry) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

type RealmCharacterCountList struct {
	Counts               []*RealmCharacterCountEntry `protobuf:"bytes,1,rep,name=counts" json:"counts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *RealmCharacterCountList) Reset()         { *m = RealmCharacterCountList{} }
func (m *RealmCharacterCountList) String() string { return proto.CompactTextString(m) }
func (*RealmCharacterCountList) ProtoMessage()    {}
func (*RealmCharacterCountList) Descriptor() ([]byte, []int) {
	return fileDescriptor_366349c60ad91e0b, []int{5}
}

func (m *RealmCharacterCountList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealmCharacterCountList.Unmarshal(m, b)
}
func (m *RealmCharacterCountList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealmCharacterCountList.Marshal(b, m, deterministic)
}
func (m *RealmCharacterCountList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealmCharacterCountList.Merge(m, src)
}
func (m *RealmCharacterCountList) XXX_Size() int {
	return xxx_messageInfo_RealmCharacterCountList.Size(m)
}
func (m *RealmCharacterCountList) XXX_DiscardUnknown() {
	xxx_messageInfo_RealmCharacterCountList.DiscardUnknown(m)
}

var xxx_messageInfo_RealmCharacterCountList proto.InternalMessageInfo

func (m *RealmCharacterCountList) GetCounts() []*RealmCharacterCountEntry {
	if m != nil {
		return m.Counts
	}
	return nil
}

type RealmEntry struct {
	WowRealmAddress      *uint32        `protobuf:"fixed32,1,req,name=wowRealmAddress" json:"wowRealmAddress,omitempty"`
	CfgTimezonesID       *uint32        `protobuf:"varint,2,req,name=cfgTimezonesID" json:"cfgTimezonesID,omitempty"`
	PopulationState      *uint32        `protobuf:"varint,3,req,name=populationState" json:"populationState,omitempty"`
	CfgCategoriesID      *uint32        `protobuf:"varint,4,req,name=cfgCategoriesID" json:"cfgCategoriesID,omitempty"`
	Version              *ClientVersion `protobuf:"bytes,5,req,name=version" json:"version,omitempty"`
	CfgRealmsID          *uint32        `protobuf:"varint,6,req,name=cfgRealmsID" json:"cfgRealmsID,omitempty"`
	Flags                *uint32        `protobuf:"varint,7,req,name=flags" json:"flags,omitempty"`
	Name                 *string        `protobuf:"bytes,8,req,name=name" json:"name,omitempty"`
	CfgConfigsID         *uint32        `protobuf:"varint,9,req,name=cfgConfigsID" json:"cfgConfigsID,omitempty"`
	CfgLanguagesID       *uint32        `protobuf:"varint,10,req,name=cfgLanguagesID" json:"cfgLanguagesID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RealmEntry) Reset()         { *m = RealmEntry{} }
func (m *RealmEntry) String() string { return proto.CompactTextString(m) }
func (*RealmEntry) ProtoMessage()    {}
func (*RealmEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_366349c60ad91e0b, []int{6}
}

func (m *RealmEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealmEntry.Unmarshal(m, b)
}
func (m *RealmEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealmEntry.Marshal(b, m, deterministic)
}
func (m *RealmEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealmEntry.Merge(m, src)
}
func (m *RealmEntry) XXX_Size() int {
	return xxx_messageInfo_RealmEntry.Size(m)
}
func (m *RealmEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_RealmEntry.DiscardUnknown(m)
}

var xxx_messageInfo_RealmEntry proto.InternalMessageInfo

func (m *RealmEntry) GetWowRealmAddress() uint32 {
	if m != nil && m.WowRealmAddress != nil {
		return *m.WowRealmAddress
	}
	return 0
}

func (m *RealmEntry) GetCfgTimezonesID() uint32 {
	if m != nil && m.CfgTimezonesID != nil {
		return *m.CfgTimezonesID
	}
	return 0
}

func (m *RealmEntry) GetPopulationState() uint32 {
	if m != nil && m.PopulationState != nil {
		return *m.PopulationState
	}
	return 0
}

func (m *RealmEntry) GetCfgCategoriesID() uint32 {
	if m != nil && m.CfgCategoriesID != nil {
		return *m.CfgCategoriesID
	}
	return 0
}

func (m *RealmEntry) GetVersion() *ClientVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *RealmEntry) GetCfgRealmsID() uint32 {
	if m != nil && m.CfgRealmsID != nil {
		return *m.CfgRealmsID
	}
	return 0
}

func (m *RealmEntry) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *RealmEntry) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *RealmEntry) GetCfgConfigsID() uint32 {
	if m != nil && m.CfgConfigsID != nil {
		return *m.CfgConfigsID
	}
	return 0
}

func (m *RealmEntry) GetCfgLanguagesID() uint32 {
	if m != nil && m.CfgLanguagesID != nil {
		return *m.CfgLanguagesID
	}
	return 0
}

type RealmState struct {
	Update               *RealmEntry `protobuf:"bytes,1,opt,name=update" json:"update,omitempty"`
	Deleting             *bool       `protobuf:"varint,2,req,name=deleting" json:"deleting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RealmState) Reset()         { *m = RealmState{} }
func (m *RealmState) String() string { return proto.CompactTextString(m) }
func (*RealmState) ProtoMessage()    {}
func (*RealmState) Descriptor() ([]byte, []int) {
	return fileDescriptor_366349c60ad91e0b, []int{7}
}

func (m *RealmState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealmState.Unmarshal(m, b)
}
func (m *RealmState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealmState.Marshal(b, m, deterministic)
}
func (m *RealmState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealmState.Merge(m, src)
}
func (m *RealmState) XXX_Size() int {
	return xxx_messageInfo_RealmState.Size(m)
}
func (m *RealmState) XXX_DiscardUnknown() {
	xxx_messageInfo_RealmState.DiscardUnknown(m)
}

var xxx_messageInfo_RealmState proto.InternalMessageInfo

func (m *RealmState) GetUpdate() *RealmEntry {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *RealmState) GetDeleting() bool {
	if m != nil && m.Deleting != nil {
		return *m.Deleting
	}
	return false
}

type RealmListUpdates struct {
	Updates              []*RealmState `protobuf:"bytes,1,rep,name=updates" json:"updates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RealmListUpdates) Reset()         { *m = RealmListUpdates{} }
func (m *RealmListUpdates) String() string { return proto.CompactTextString(m) }
func (*RealmListUpdates) ProtoMessage()    {}
func (*RealmListUpdates) Descriptor() ([]byte, []int) {
	return fileDescriptor_366349c60ad91e0b, []int{8}
}

func (m *RealmListUpdates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealmListUpdates.Unmarshal(m, b)
}
func (m *RealmListUpdates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealmListUpdates.Marshal(b, m, deterministic)
}
func (m *RealmListUpdates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealmListUpdates.Merge(m, src)
}
func (m *RealmListUpdates) XXX_Size() int {
	return xxx_messageInfo_RealmListUpdates.Size(m)
}
func (m *RealmListUpdates) XXX_DiscardUnknown() {
	xxx_messageInfo_RealmListUpdates.DiscardUnknown(m)
}

var xxx_messageInfo_RealmListUpdates proto.InternalMessageInfo

func (m *RealmListUpdates) GetUpdates() []*RealmState {
	if m != nil {
		return m.Updates
	}
	return nil
}

type IPAddress struct {
	Ip                   *string  `protobuf:"bytes,1,req,name=ip" json:"ip,omitempty"`
	Port                 *uint32  `protobuf:"varint,2,req,name=port" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPAddress) Reset()         { *m = IPAddress{} }
func (m *IPAddress) String() string { return proto.CompactTextString(m) }
func (*IPAddress) ProtoMessage()    {}
func (*IPAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_366349c60ad91e0b, []int{9}
}

func (m *IPAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPAddress.Unmarshal(m, b)
}
func (m *IPAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPAddress.Marshal(b, m, deterministic)
}
func (m *IPAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPAddress.Merge(m, src)
}
func (m *IPAddress) XXX_Size() int {
	return xxx_messageInfo_IPAddress.Size(m)
}
func (m *IPAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_IPAddress.DiscardUnknown(m)
}

var xxx_messageInfo_IPAddress proto.InternalMessageInfo

func (m *IPAddress) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *IPAddress) GetPort() uint32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

type RealmIPAddressFamily struct {
	Family               *uint32      `protobuf:"varint,1,req,name=family" json:"family,omitempty"`
	Addresses            []*IPAddress `protobuf:"bytes,2,rep,name=addresses" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RealmIPAddressFamily) Reset()         { *m = RealmIPAddressFamily{} }
func (m *RealmIPAddressFamily) String() string { return proto.CompactTextString(m) }
func (*RealmIPAddressFamily) ProtoMessage()    {}
func (*RealmIPAddressFamily) Descriptor() ([]byte, []int) {
	return fileDescriptor_366349c60ad91e0b, []int{10}
}

func (m *RealmIPAddressFamily) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealmIPAddressFamily.Unmarshal(m, b)
}
func (m *RealmIPAddressFamily) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealmIPAddressFamily.Marshal(b, m, deterministic)
}
func (m *RealmIPAddressFamily) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealmIPAddressFamily.Merge(m, src)
}
func (m *RealmIPAddressFamily) XXX_Size() int {
	return xxx_messageInfo_RealmIPAddressFamily.Size(m)
}
func (m *RealmIPAddressFamily) XXX_DiscardUnknown() {
	xxx_messageInfo_RealmIPAddressFamily.DiscardUnknown(m)
}

var xxx_messageInfo_RealmIPAddressFamily proto.InternalMessageInfo

func (m *RealmIPAddressFamily) GetFamily() uint32 {
	if m != nil && m.Family != nil {
		return *m.Family
	}
	return 0
}

func (m *RealmIPAddressFamily) GetAddresses() []*IPAddress {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type RealmListServerIPAddresses struct {
	Families             []*RealmIPAddressFamily `protobuf:"bytes,1,rep,name=families" json:"families,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RealmListServerIPAddresses) Reset()         { *m = RealmListServerIPAddresses{} }
func (m *RealmListServerIPAddresses) String() string { return proto.CompactTextString(m) }
func (*RealmListServerIPAddresses) ProtoMessage()    {}
func (*RealmListServerIPAddresses) Descriptor() ([]byte, []int) {
	return fileDescriptor_366349c60ad91e0b, []int{11}
}

func (m *RealmListServerIPAddresses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealmListServerIPAddresses.Unmarshal(m, b)
}
func (m *RealmListServerIPAddresses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealmListServerIPAddresses.Marshal(b, m, deterministic)
}
func (m *RealmListServerIPAddresses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealmListServerIPAddresses.Merge(m, src)
}
func (m *RealmListServerIPAddresses) XXX_Size() int {
	return xxx_messageInfo_RealmListServerIPAddresses.Size(m)
}
func (m *RealmListServerIPAddresses) XXX_DiscardUnknown() {
	xxx_messageInfo_RealmListServerIPAddresses.DiscardUnknown(m)
}

var xxx_messageInfo_RealmListServerIPAddresses proto.InternalMessageInfo

func (m *RealmListServerIPAddresses) GetFamilies() []*RealmIPAddressFamily {
	if m != nil {
		return m.Families
	}
	return nil
}

func init() {
	proto.RegisterType((*RealmListTicketIdentity)(nil), "JSON.RealmList.RealmListTicketIdentity")
	proto.RegisterType((*ClientVersion)(nil), "JSON.RealmList.ClientVersion")
	proto.RegisterType((*ClientInformation)(nil), "JSON.RealmList.ClientInformation")
	proto.RegisterType((*RealmListTicketClientInformation)(nil), "JSON.RealmList.RealmListTicketClientInformation")
	proto.RegisterType((*RealmCharacterCountEntry)(nil), "JSON.RealmList.RealmCharacterCountEntry")
	proto.RegisterType((*RealmCharacterCountList)(nil), "JSON.RealmList.RealmCharacterCountList")
	proto.RegisterType((*RealmEntry)(nil), "JSON.RealmList.RealmEntry")
	proto.RegisterType((*RealmState)(nil), "JSON.RealmList.RealmState")
	proto.RegisterType((*RealmListUpdates)(nil), "JSON.RealmList.RealmListUpdates")
	proto.RegisterType((*IPAddress)(nil), "JSON.RealmList.IPAddress")
	proto.RegisterType((*RealmIPAddressFamily)(nil), "JSON.RealmList.RealmIPAddressFamily")
	proto.RegisterType((*RealmListServerIPAddresses)(nil), "JSON.RealmList.RealmListServerIPAddresses")
}

func init() {
	proto.RegisterFile("github.com/Gophercraft/core/bnet/public_protos/RealmList.proto", fileDescriptor_366349c60ad91e0b)
}

var fileDescriptor_366349c60ad91e0b = []byte{
	// 843 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xed, 0x6e, 0xe3, 0x44,
	0x14, 0x55, 0xdd, 0x8f, 0x34, 0x37, 0x4d, 0xd9, 0x1d, 0xad, 0xc0, 0x54, 0x02, 0x05, 0x6b, 0x85,
	0x22, 0x84, 0x92, 0x55, 0x00, 0xed, 0xdf, 0xed, 0xa6, 0xa0, 0x0d, 0x2a, 0x1f, 0x9a, 0x96, 0x95,
	0x28, 0x08, 0x34, 0x75, 0xae, 0xdd, 0x01, 0x7b, 0xc6, 0x1a, 0x8f, 0xbb, 0x84, 0xbf, 0x3c, 0x0b,
	0xaf, 0xc7, 0x33, 0xa0, 0xb9, 0x9e, 0xb8, 0xb6, 0x5b, 0x24, 0xf6, 0x57, 0xe7, 0x9e, 0x9e, 0x7b,
	0xe6, 0xcc, 0xfd, 0x70, 0x60, 0x99, 0x4a, 0x7b, 0x53, 0x5d, 0xcf, 0x62, 0x9d, 0xcf, 0xcb, 0xaa,
	0x40, 0x53, 0x3c, 0x7b, 0x66, 0xe7, 0xa9, 0x2e, 0x6e, 0xd0, 0xc4, 0x46, 0x24, 0x76, 0x7e, 0xad,
	0xd0, 0xce, 0x8b, 0xea, 0x3a, 0x93, 0xf1, 0xaf, 0x85, 0xd1, 0x56, 0x97, 0x73, 0x8e, 0x22, 0xcb,
	0xcf, 0x65, 0x69, 0x67, 0x04, 0xb0, 0xe3, 0xaf, 0x2f, 0xbe, 0xfb, 0x76, 0xd6, 0xa0, 0x51, 0x0e,
	0xef, 0x35, 0xc1, 0xa5, 0x8c, 0x7f, 0x47, 0xbb, 0x5a, 0xa3, 0xb2, 0xd2, 0x6e, 0xd8, 0x53, 0x18,
	0xa7, 0x22, 0xc7, 0xd3, 0x38, 0xd6, 0x95, 0xb2, 0xab, 0xb3, 0x70, 0x67, 0x12, 0x4c, 0x07, 0xbc,
	0x0b, 0xb2, 0x4f, 0xe1, 0x71, 0x0b, 0xe0, 0x98, 0x4a, 0xad, 0xc2, 0x60, 0x12, 0x4c, 0xc7, 0xfc,
	0xfe, 0x3f, 0xa2, 0xbf, 0x77, 0x60, 0xbc, 0xcc, 0x24, 0x2a, 0xfb, 0x1a, 0x4d, 0x29, 0xb5, 0x62,
	0x11, 0x1c, 0xdd, 0xd6, 0xc7, 0x6f, 0xc4, 0x6f, 0xda, 0xd0, 0x25, 0x63, 0xde, 0xc1, 0xda, 0x1c,
	0xa9, 0xb4, 0xf1, 0xf2, 0x1d, 0x8c, 0x4d, 0xe1, 0x1d, 0x1f, 0x73, 0xbc, 0x95, 0xee, 0x6f, 0xb8,
	0x4b, 0xb4, 0x3e, 0xdc, 0x52, 0x7b, 0x59, 0xc9, 0x6c, 0x1d, 0xee, 0xd1, 0xb3, 0x3a, 0x58, 0xf4,
	0xcf, 0x2e, 0x3c, 0xae, 0x7d, 0xae, 0x54, 0xa2, 0x4d, 0x2e, 0xac, 0xcb, 0x3c, 0x81, 0xc3, 0x22,
	0x13, 0xd6, 0x01, 0xbe, 0x18, 0x4d, 0xec, 0x54, 0xaf, 0x5d, 0xea, 0x6b, 0x61, 0xa4, 0x50, 0x96,
	0x3c, 0x0e, 0x79, 0x07, 0x63, 0x0c, 0xf6, 0xec, 0xa6, 0x40, 0x32, 0x36, 0xe0, 0x74, 0x76, 0x9a,
	0x56, 0xe6, 0x78, 0xa5, 0x15, 0x92, 0x93, 0x21, 0x6f, 0x62, 0x36, 0x81, 0x51, 0x5c, 0x19, 0x83,
	0xca, 0x5e, 0xca, 0x1c, 0xc3, 0x7d, 0x7a, 0x4f, 0x1b, 0x62, 0x1f, 0x02, 0x58, 0xfc, 0xc3, 0x9e,
	0xeb, 0x58, 0x64, 0x18, 0x1e, 0x90, 0x6e, 0x0b, 0x71, 0x0a, 0xa2, 0x5a, 0x4b, 0xed, 0x09, 0x03,
	0x22, 0xb4, 0x21, 0xf6, 0x09, 0x3c, 0xf2, 0x2f, 0x3f, 0x13, 0x56, 0xd4, 0x15, 0x39, 0x24, 0xda,
	0x3d, 0x9c, 0x3d, 0x87, 0x81, 0xc7, 0xc2, 0xe1, 0x24, 0x98, 0x8e, 0x16, 0x1f, 0xcc, 0xba, 0xe3,
	0x34, 0xeb, 0xf4, 0x96, 0x6f, 0xd9, 0xec, 0x5d, 0x38, 0x28, 0x31, 0x36, 0x68, 0x43, 0x98, 0x04,
	0xd3, 0x23, 0xee, 0x23, 0x67, 0x3f, 0xa6, 0x8c, 0x53, 0x13, 0xdf, 0x84, 0xa3, 0xda, 0xfe, 0x1d,
	0xe2, 0x46, 0xb0, 0xdc, 0x94, 0x16, 0x73, 0xaf, 0x18, 0x1e, 0x51, 0x85, 0xba, 0xa0, 0x2b, 0xfd,
	0xb6, 0x0d, 0x97, 0xae, 0xbc, 0xe3, 0xba, 0xa1, 0x6d, 0xcc, 0xdd, 0x54, 0x27, 0xd1, 0x4d, 0xc7,
	0xf5, 0x4d, 0x77, 0x48, 0xf4, 0x23, 0x4c, 0x7a, 0x7b, 0x70, 0xbf, 0xfd, 0x5f, 0xc0, 0x9e, 0x54,
	0x89, 0xa6, 0xd6, 0x8f, 0x16, 0x1f, 0x3d, 0xfc, 0xf6, 0x56, 0x02, 0x27, 0x7a, 0x74, 0x05, 0x21,
	0x91, 0x96, 0x37, 0xc2, 0x88, 0xd8, 0xa2, 0x59, 0xba, 0x85, 0xf8, 0x52, 0x59, 0xb3, 0x71, 0x53,
	0xfb, 0x46, 0xbf, 0xa1, 0x7f, 0x9f, 0xae, 0xd7, 0x06, 0xcb, 0xd2, 0x0f, 0x56, 0x1f, 0x66, 0x4f,
	0x60, 0x9f, 0x16, 0xc9, 0x0f, 0x7f, 0x1d, 0x44, 0x3f, 0xf9, 0xf5, 0xed, 0x6a, 0x3b, 0x3b, 0xec,
	0x05, 0x1c, 0x10, 0xc7, 0x29, 0xee, 0x4e, 0x47, 0x8b, 0x69, 0xdf, 0xef, 0x7f, 0x99, 0xe2, 0x3e,
	0x2f, 0xfa, 0x6b, 0x17, 0x80, 0x48, 0x6f, 0xeb, 0xf5, 0x63, 0x38, 0x8e, 0x93, 0xd4, 0x0d, 0xe8,
	0x9f, 0x5a, 0x61, 0xb9, 0x3a, 0xf3, 0xa6, 0x7b, 0xa8, 0x53, 0x2c, 0x74, 0x51, 0x65, 0x54, 0xad,
	0x0b, 0x2b, 0x2c, 0x6e, 0x77, 0xb6, 0x07, 0x3b, 0x66, 0x9c, 0xa4, 0x4b, 0x61, 0x31, 0xd5, 0x46,
	0x92, 0xe4, 0x5e, 0xcd, 0xec, 0xc1, 0xed, 0x19, 0xdd, 0x7f, 0xab, 0x19, 0x75, 0xcb, 0x96, 0xa4,
	0x44, 0x73, 0xf2, 0x07, 0x7e, 0xd9, 0xee, 0x20, 0xd7, 0x82, 0x24, 0x13, 0x69, 0x49, 0x6b, 0x34,
	0xe6, 0x75, 0xe0, 0x96, 0x5a, 0x89, 0x1c, 0x69, 0x69, 0x86, 0x9c, 0xce, 0x6e, 0x22, 0x9d, 0x2f,
	0xad, 0x12, 0x99, 0x3a, 0xb1, 0x61, 0xfd, 0xc1, 0x6a, 0x63, 0xbe, 0x48, 0xe7, 0x42, 0xa5, 0x95,
	0x48, 0xe9, 0x45, 0xd0, 0x14, 0xa9, 0x85, 0x46, 0x3f, 0xfb, 0x26, 0xd4, 0x85, 0x58, 0xc0, 0x41,
	0x55, 0xac, 0x5d, 0xa5, 0x76, 0x26, 0x3b, 0xd3, 0xd1, 0xe2, 0xe4, 0xc1, 0xae, 0xfa, 0x3e, 0xd6,
	0x4c, 0xf7, 0x89, 0x59, 0x63, 0x86, 0x56, 0xaa, 0x94, 0x1a, 0x71, 0xc8, 0x9b, 0x38, 0x7a, 0x05,
	0x8f, 0x9a, 0xdc, 0x1f, 0x88, 0x5e, 0xb2, 0xcf, 0x61, 0x50, 0x67, 0x6e, 0x47, 0xe7, 0xe1, 0x4b,
	0xc8, 0x10, 0xdf, 0x52, 0xa3, 0x39, 0x0c, 0x57, 0xdf, 0x6f, 0x27, 0xe0, 0x18, 0x02, 0x59, 0xd0,
	0x78, 0x0c, 0x79, 0x20, 0x0b, 0x57, 0xa4, 0x42, 0x9b, 0xed, 0xf0, 0xd2, 0x39, 0x4a, 0xe1, 0x09,
	0xe9, 0x34, 0x59, 0x5f, 0x89, 0x5c, 0x66, 0x1b, 0xf7, 0xb1, 0x48, 0xe8, 0xe4, 0x7f, 0x0b, 0x7c,
	0xc4, 0x9e, 0xc3, 0x50, 0xd4, 0x44, 0x2c, 0xc3, 0x80, 0x8c, 0xbd, 0xdf, 0x37, 0xd6, 0x68, 0xf1,
	0x3b, 0x6e, 0xf4, 0x0b, 0x9c, 0x34, 0x8c, 0x0b, 0x34, 0xb7, 0x68, 0x1a, 0x1a, 0x96, 0xec, 0x05,
	0x1c, 0xd2, 0x05, 0xb2, 0x79, 0xee, 0xd3, 0x07, 0x9f, 0xdb, 0xb3, 0xc9, 0x9b, 0xac, 0x97, 0x9f,
	0xbd, 0x0a, 0xae, 0xe6, 0xff, 0xef, 0xc7, 0xd9, 0x38, 0xa5, 0x4c, 0x96, 0xf6, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xaa, 0x0e, 0xf5, 0x57, 0xce, 0x07, 0x00, 0x00,
}

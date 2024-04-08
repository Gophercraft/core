// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: club_membership_listener.proto

package membership

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "github.com/Gophercraft/core/bnet/bgs/protocol"
	v11 "github.com/Gophercraft/core/bnet/bgs/protocol/account/v1"
	v1 "github.com/Gophercraft/core/bnet/bgs/protocol/club/v1"
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

type ClubAddedNotification struct {
	AgentId              *v1.MemberId                  `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Membership           *v1.ClubMembershipDescription `protobuf:"bytes,3,opt,name=membership" json:"membership,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ClubAddedNotification) Reset()         { *m = ClubAddedNotification{} }
func (m *ClubAddedNotification) String() string { return proto.CompactTextString(m) }
func (*ClubAddedNotification) ProtoMessage()    {}
func (*ClubAddedNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_17767d2abc085b26, []int{0}
}

func (m *ClubAddedNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClubAddedNotification.Unmarshal(m, b)
}
func (m *ClubAddedNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClubAddedNotification.Marshal(b, m, deterministic)
}
func (m *ClubAddedNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClubAddedNotification.Merge(m, src)
}
func (m *ClubAddedNotification) XXX_Size() int {
	return xxx_messageInfo_ClubAddedNotification.Size(m)
}
func (m *ClubAddedNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_ClubAddedNotification.DiscardUnknown(m)
}

var xxx_messageInfo_ClubAddedNotification proto.InternalMessageInfo

func (m *ClubAddedNotification) GetAgentId() *v1.MemberId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *ClubAddedNotification) GetMembership() *v1.ClubMembershipDescription {
	if m != nil {
		return m.Membership
	}
	return nil
}

type ClubRemovedNotification struct {
	AgentId              *v1.MemberId          `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	MemberId             *v1.MemberId          `protobuf:"bytes,3,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
	ClubId               *uint64               `protobuf:"varint,4,opt,name=club_id,json=clubId" json:"club_id,omitempty"`
	Reason               *v1.ClubRemovedReason `protobuf:"varint,5,opt,name=reason,enum=bgs.protocol.club.v1.ClubRemovedReason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClubRemovedNotification) Reset()         { *m = ClubRemovedNotification{} }
func (m *ClubRemovedNotification) String() string { return proto.CompactTextString(m) }
func (*ClubRemovedNotification) ProtoMessage()    {}
func (*ClubRemovedNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_17767d2abc085b26, []int{1}
}

func (m *ClubRemovedNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClubRemovedNotification.Unmarshal(m, b)
}
func (m *ClubRemovedNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClubRemovedNotification.Marshal(b, m, deterministic)
}
func (m *ClubRemovedNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClubRemovedNotification.Merge(m, src)
}
func (m *ClubRemovedNotification) XXX_Size() int {
	return xxx_messageInfo_ClubRemovedNotification.Size(m)
}
func (m *ClubRemovedNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_ClubRemovedNotification.DiscardUnknown(m)
}

var xxx_messageInfo_ClubRemovedNotification proto.InternalMessageInfo

func (m *ClubRemovedNotification) GetAgentId() *v1.MemberId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *ClubRemovedNotification) GetMemberId() *v1.MemberId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

func (m *ClubRemovedNotification) GetClubId() uint64 {
	if m != nil && m.ClubId != nil {
		return *m.ClubId
	}
	return 0
}

func (m *ClubRemovedNotification) GetReason() v1.ClubRemovedReason {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return v1.ClubRemovedReason_CLUB_REMOVED_REASON_NONE
}

type ReceivedInvitationAddedNotification struct {
	AgentId              *v1.MemberId       `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Invitation           *v1.ClubInvitation `protobuf:"bytes,3,opt,name=invitation" json:"invitation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ReceivedInvitationAddedNotification) Reset()         { *m = ReceivedInvitationAddedNotification{} }
func (m *ReceivedInvitationAddedNotification) String() string { return proto.CompactTextString(m) }
func (*ReceivedInvitationAddedNotification) ProtoMessage()    {}
func (*ReceivedInvitationAddedNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_17767d2abc085b26, []int{2}
}

func (m *ReceivedInvitationAddedNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceivedInvitationAddedNotification.Unmarshal(m, b)
}
func (m *ReceivedInvitationAddedNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceivedInvitationAddedNotification.Marshal(b, m, deterministic)
}
func (m *ReceivedInvitationAddedNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceivedInvitationAddedNotification.Merge(m, src)
}
func (m *ReceivedInvitationAddedNotification) XXX_Size() int {
	return xxx_messageInfo_ReceivedInvitationAddedNotification.Size(m)
}
func (m *ReceivedInvitationAddedNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceivedInvitationAddedNotification.DiscardUnknown(m)
}

var xxx_messageInfo_ReceivedInvitationAddedNotification proto.InternalMessageInfo

func (m *ReceivedInvitationAddedNotification) GetAgentId() *v1.MemberId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *ReceivedInvitationAddedNotification) GetInvitation() *v1.ClubInvitation {
	if m != nil {
		return m.Invitation
	}
	return nil
}

type ReceivedInvitationRemovedNotification struct {
	AgentId              *v1.MemberId                      `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	InvitationId         *uint64                           `protobuf:"fixed64,3,opt,name=invitation_id,json=invitationId" json:"invitation_id,omitempty"`
	Reason               *protocol.InvitationRemovedReason `protobuf:"varint,4,opt,name=reason,enum=bgs.protocol.InvitationRemovedReason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ReceivedInvitationRemovedNotification) Reset()         { *m = ReceivedInvitationRemovedNotification{} }
func (m *ReceivedInvitationRemovedNotification) String() string { return proto.CompactTextString(m) }
func (*ReceivedInvitationRemovedNotification) ProtoMessage()    {}
func (*ReceivedInvitationRemovedNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_17767d2abc085b26, []int{3}
}

func (m *ReceivedInvitationRemovedNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceivedInvitationRemovedNotification.Unmarshal(m, b)
}
func (m *ReceivedInvitationRemovedNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceivedInvitationRemovedNotification.Marshal(b, m, deterministic)
}
func (m *ReceivedInvitationRemovedNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceivedInvitationRemovedNotification.Merge(m, src)
}
func (m *ReceivedInvitationRemovedNotification) XXX_Size() int {
	return xxx_messageInfo_ReceivedInvitationRemovedNotification.Size(m)
}
func (m *ReceivedInvitationRemovedNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceivedInvitationRemovedNotification.DiscardUnknown(m)
}

var xxx_messageInfo_ReceivedInvitationRemovedNotification proto.InternalMessageInfo

func (m *ReceivedInvitationRemovedNotification) GetAgentId() *v1.MemberId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *ReceivedInvitationRemovedNotification) GetInvitationId() uint64 {
	if m != nil && m.InvitationId != nil {
		return *m.InvitationId
	}
	return 0
}

func (m *ReceivedInvitationRemovedNotification) GetReason() protocol.InvitationRemovedReason {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return protocol.InvitationRemovedReason_INVITATION_REMOVED_REASON_ACCEPTED
}

type SharedSettingsChangedNotification struct {
	AgentId              *v11.AccountId                   `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Assignment           *v1.ClubSharedSettingsAssignment `protobuf:"bytes,4,opt,name=assignment" json:"assignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *SharedSettingsChangedNotification) Reset()         { *m = SharedSettingsChangedNotification{} }
func (m *SharedSettingsChangedNotification) String() string { return proto.CompactTextString(m) }
func (*SharedSettingsChangedNotification) ProtoMessage()    {}
func (*SharedSettingsChangedNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_17767d2abc085b26, []int{4}
}

func (m *SharedSettingsChangedNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedSettingsChangedNotification.Unmarshal(m, b)
}
func (m *SharedSettingsChangedNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedSettingsChangedNotification.Marshal(b, m, deterministic)
}
func (m *SharedSettingsChangedNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedSettingsChangedNotification.Merge(m, src)
}
func (m *SharedSettingsChangedNotification) XXX_Size() int {
	return xxx_messageInfo_SharedSettingsChangedNotification.Size(m)
}
func (m *SharedSettingsChangedNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedSettingsChangedNotification.DiscardUnknown(m)
}

var xxx_messageInfo_SharedSettingsChangedNotification proto.InternalMessageInfo

func (m *SharedSettingsChangedNotification) GetAgentId() *v11.AccountId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SharedSettingsChangedNotification) GetAssignment() *v1.ClubSharedSettingsAssignment {
	if m != nil {
		return m.Assignment
	}
	return nil
}

type StreamMentionAddedNotification struct {
	AgentId              *v1.MemberId      `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Mention              *v1.StreamMention `protobuf:"bytes,3,opt,name=mention" json:"mention,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StreamMentionAddedNotification) Reset()         { *m = StreamMentionAddedNotification{} }
func (m *StreamMentionAddedNotification) String() string { return proto.CompactTextString(m) }
func (*StreamMentionAddedNotification) ProtoMessage()    {}
func (*StreamMentionAddedNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_17767d2abc085b26, []int{5}
}

func (m *StreamMentionAddedNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMentionAddedNotification.Unmarshal(m, b)
}
func (m *StreamMentionAddedNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMentionAddedNotification.Marshal(b, m, deterministic)
}
func (m *StreamMentionAddedNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMentionAddedNotification.Merge(m, src)
}
func (m *StreamMentionAddedNotification) XXX_Size() int {
	return xxx_messageInfo_StreamMentionAddedNotification.Size(m)
}
func (m *StreamMentionAddedNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMentionAddedNotification.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMentionAddedNotification proto.InternalMessageInfo

func (m *StreamMentionAddedNotification) GetAgentId() *v1.MemberId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *StreamMentionAddedNotification) GetMention() *v1.StreamMention {
	if m != nil {
		return m.Mention
	}
	return nil
}

type StreamMentionRemovedNotification struct {
	AgentId              *v11.AccountId         `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	MentionId            *protocol.TimeSeriesId `protobuf:"bytes,3,opt,name=mention_id,json=mentionId" json:"mention_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StreamMentionRemovedNotification) Reset()         { *m = StreamMentionRemovedNotification{} }
func (m *StreamMentionRemovedNotification) String() string { return proto.CompactTextString(m) }
func (*StreamMentionRemovedNotification) ProtoMessage()    {}
func (*StreamMentionRemovedNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_17767d2abc085b26, []int{6}
}

func (m *StreamMentionRemovedNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMentionRemovedNotification.Unmarshal(m, b)
}
func (m *StreamMentionRemovedNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMentionRemovedNotification.Marshal(b, m, deterministic)
}
func (m *StreamMentionRemovedNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMentionRemovedNotification.Merge(m, src)
}
func (m *StreamMentionRemovedNotification) XXX_Size() int {
	return xxx_messageInfo_StreamMentionRemovedNotification.Size(m)
}
func (m *StreamMentionRemovedNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMentionRemovedNotification.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMentionRemovedNotification proto.InternalMessageInfo

func (m *StreamMentionRemovedNotification) GetAgentId() *v11.AccountId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *StreamMentionRemovedNotification) GetMentionId() *protocol.TimeSeriesId {
	if m != nil {
		return m.MentionId
	}
	return nil
}

type StreamMentionAdvanceViewTimeNotification struct {
	AgentId              *v11.AccountId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	ViewTime             *uint64        `protobuf:"varint,3,opt,name=view_time,json=viewTime" json:"view_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StreamMentionAdvanceViewTimeNotification) Reset() {
	*m = StreamMentionAdvanceViewTimeNotification{}
}
func (m *StreamMentionAdvanceViewTimeNotification) String() string { return proto.CompactTextString(m) }
func (*StreamMentionAdvanceViewTimeNotification) ProtoMessage()    {}
func (*StreamMentionAdvanceViewTimeNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_17767d2abc085b26, []int{7}
}

func (m *StreamMentionAdvanceViewTimeNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMentionAdvanceViewTimeNotification.Unmarshal(m, b)
}
func (m *StreamMentionAdvanceViewTimeNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMentionAdvanceViewTimeNotification.Marshal(b, m, deterministic)
}
func (m *StreamMentionAdvanceViewTimeNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMentionAdvanceViewTimeNotification.Merge(m, src)
}
func (m *StreamMentionAdvanceViewTimeNotification) XXX_Size() int {
	return xxx_messageInfo_StreamMentionAdvanceViewTimeNotification.Size(m)
}
func (m *StreamMentionAdvanceViewTimeNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMentionAdvanceViewTimeNotification.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMentionAdvanceViewTimeNotification proto.InternalMessageInfo

func (m *StreamMentionAdvanceViewTimeNotification) GetAgentId() *v11.AccountId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *StreamMentionAdvanceViewTimeNotification) GetViewTime() uint64 {
	if m != nil && m.ViewTime != nil {
		return *m.ViewTime
	}
	return 0
}

func init() {
	proto.RegisterType((*ClubAddedNotification)(nil), "bgs.protocol.club.v1.membership.ClubAddedNotification")
	proto.RegisterType((*ClubRemovedNotification)(nil), "bgs.protocol.club.v1.membership.ClubRemovedNotification")
	proto.RegisterType((*ReceivedInvitationAddedNotification)(nil), "bgs.protocol.club.v1.membership.ReceivedInvitationAddedNotification")
	proto.RegisterType((*ReceivedInvitationRemovedNotification)(nil), "bgs.protocol.club.v1.membership.ReceivedInvitationRemovedNotification")
	proto.RegisterType((*SharedSettingsChangedNotification)(nil), "bgs.protocol.club.v1.membership.SharedSettingsChangedNotification")
	proto.RegisterType((*StreamMentionAddedNotification)(nil), "bgs.protocol.club.v1.membership.StreamMentionAddedNotification")
	proto.RegisterType((*StreamMentionRemovedNotification)(nil), "bgs.protocol.club.v1.membership.StreamMentionRemovedNotification")
	proto.RegisterType((*StreamMentionAdvanceViewTimeNotification)(nil), "bgs.protocol.club.v1.membership.StreamMentionAdvanceViewTimeNotification")
}

func init() { proto.RegisterFile("club_membership_listener.proto", fileDescriptor_17767d2abc085b26) }

var fileDescriptor_17767d2abc085b26 = []byte{
	// 757 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcb, 0x6e, 0xd3, 0x4a,
	0x18, 0x3e, 0xd3, 0xa6, 0x49, 0xfa, 0xe7, 0xf4, 0xe8, 0x68, 0x74, 0x4e, 0x9b, 0xa6, 0x52, 0x08,
	0x2e, 0x15, 0x91, 0x8a, 0xec, 0x36, 0x0b, 0x44, 0x41, 0x55, 0x95, 0x5e, 0x90, 0x8c, 0x68, 0x83,
	0x26, 0x88, 0x05, 0x9b, 0xc8, 0xb1, 0xa7, 0xce, 0x48, 0xf1, 0x38, 0xb2, 0x27, 0xae, 0x60, 0x81,
	0xaa, 0x0a, 0x21, 0xc4, 0x82, 0x07, 0x40, 0xec, 0x90, 0x78, 0x04, 0x1e, 0x81, 0x25, 0x2f, 0xc2,
	0x13, 0x74, 0x87, 0xec, 0x38, 0x17, 0x27, 0x0e, 0x6e, 0x4b, 0x76, 0x93, 0xe4, 0xff, 0x2e, 0xff,
	0x37, 0xa3, 0x4f, 0x81, 0xa2, 0xde, 0xee, 0x36, 0x1b, 0x16, 0xb5, 0x9a, 0xd4, 0x71, 0x5b, 0xac,
	0xd3, 0x68, 0x33, 0x57, 0x50, 0x4e, 0x1d, 0xb9, 0xe3, 0xd8, 0xc2, 0xc6, 0xb7, 0x9a, 0xa6, 0xdb,
	0x3b, 0xea, 0x76, 0x5b, 0xf6, 0x87, 0x65, 0x6f, 0x5b, 0x1e, 0xce, 0x17, 0xfe, 0x0d, 0x08, 0xc4,
	0xab, 0x0e, 0x0d, 0xe7, 0xa4, 0x2f, 0x08, 0xfe, 0x3f, 0x68, 0x77, 0x9b, 0x55, 0xc3, 0xa0, 0xc6,
	0x89, 0x2d, 0xd8, 0x29, 0xd3, 0x35, 0xc1, 0x6c, 0x8e, 0x77, 0x20, 0xab, 0x99, 0x94, 0x8b, 0x06,
	0x33, 0xf2, 0xa8, 0x84, 0xca, 0xb9, 0x4a, 0x51, 0x8e, 0xe5, 0x3f, 0x0e, 0xf8, 0x55, 0x83, 0x64,
	0x82, 0x79, 0xd5, 0xc0, 0x35, 0x80, 0xa1, 0x68, 0x7e, 0x3e, 0x00, 0x2b, 0xf1, 0x60, 0x5f, 0xfb,
	0x78, 0x30, 0x7b, 0x48, 0x5d, 0xdd, 0x61, 0x1d, 0x5f, 0x9f, 0x8c, 0x50, 0x48, 0x3f, 0x11, 0xac,
	0xf8, 0x93, 0x84, 0x5a, 0xb6, 0x37, 0x3b, 0x9f, 0x8f, 0x60, 0xb1, 0x27, 0xe2, 0x63, 0xe7, 0xaf,
	0x84, 0xcd, 0x5a, 0xe1, 0x09, 0xaf, 0x40, 0x26, 0x48, 0x93, 0x19, 0xf9, 0x54, 0x09, 0x95, 0x53,
	0x24, 0xed, 0x7f, 0x54, 0x0d, 0xbc, 0x07, 0x69, 0x87, 0x6a, 0xae, 0xcd, 0xf3, 0x0b, 0x25, 0x54,
	0xfe, 0xa7, 0x72, 0x77, 0xfa, 0xe6, 0xe1, 0x3e, 0x24, 0x18, 0x27, 0x21, 0x4c, 0xfa, 0x8a, 0x60,
	0x9d, 0x50, 0x9d, 0x32, 0x8f, 0x1a, 0x2a, 0xf7, 0x98, 0x08, 0x16, 0x9d, 0xe9, 0x0d, 0x1d, 0x02,
	0xb0, 0x01, 0x73, 0xb8, 0xfa, 0x9d, 0xe9, 0x3e, 0x87, 0x2e, 0xc8, 0x08, 0x4e, 0xfa, 0x8e, 0x60,
	0x63, 0xd2, 0xe8, 0x8c, 0x2f, 0x69, 0x1d, 0x96, 0x86, 0x92, 0xfd, 0x8b, 0x4a, 0x93, 0xbf, 0x87,
	0x5f, 0xaa, 0x06, 0xde, 0x1d, 0x64, 0x9e, 0x0a, 0x32, 0xdf, 0x88, 0xb2, 0x4f, 0x98, 0x1b, 0x4b,
	0xfc, 0x1b, 0x82, 0xdb, 0xf5, 0x96, 0xe6, 0x50, 0xa3, 0x4e, 0x85, 0x60, 0xdc, 0x74, 0x0f, 0x5a,
	0x1a, 0x37, 0xc7, 0x96, 0xd8, 0x9d, 0x58, 0x42, 0x8a, 0xca, 0x68, 0xba, 0x6e, 0x77, 0xb9, 0xf0,
	0xf7, 0xa8, 0xf6, 0x8e, 0xa3, 0x8b, 0x10, 0x00, 0xcd, 0x75, 0x99, 0xc9, 0x2d, 0xca, 0x45, 0xe0,
	0x33, 0x57, 0xa9, 0x4c, 0xcf, 0x3c, 0xea, 0xa7, 0x3a, 0x40, 0x92, 0x11, 0x16, 0xe9, 0x13, 0x82,
	0x62, 0x5d, 0x38, 0x54, 0xb3, 0x8e, 0x29, 0x9f, 0xf9, 0x2b, 0xd9, 0x85, 0x8c, 0xd5, 0xa3, 0x0d,
	0x9f, 0xc8, 0x7a, 0x3c, 0x32, 0xe2, 0x80, 0xf4, 0x31, 0xd2, 0x67, 0x04, 0xa5, 0xe8, 0x4f, 0x31,
	0x2f, 0xe3, 0x0f, 0x43, 0xdd, 0xf1, 0xab, 0x86, 0x8f, 0x3e, 0x8d, 0x5c, 0xa5, 0x10, 0x25, 0x78,
	0xce, 0x2c, 0x5a, 0xa7, 0x0e, 0xa3, 0xae, 0x6a, 0x90, 0xc5, 0x70, 0x5a, 0x35, 0xa4, 0x77, 0x08,
	0xca, 0x63, 0xd9, 0x79, 0x1a, 0xd7, 0xe9, 0x0b, 0x46, 0xcf, 0x7c, 0xcc, 0x2c, 0x6d, 0xae, 0xc1,
	0xa2, 0xc7, 0xe8, 0x59, 0x43, 0x30, 0x8b, 0x06, 0x2e, 0x53, 0x24, 0xeb, 0x85, 0x3a, 0x95, 0x1f,
	0x59, 0x58, 0x8e, 0xf6, 0xe0, 0xd3, 0xb0, 0xd7, 0x71, 0x0b, 0x72, 0x35, 0x3e, 0xe8, 0x67, 0x7c,
	0x5f, 0x4e, 0x68, 0x78, 0x39, 0xb6, 0xcb, 0x0b, 0xab, 0x51, 0xdc, 0x49, 0xad, 0x41, 0x8e, 0xea,
	0xcf, 0x6a, 0x27, 0xf5, 0x23, 0x29, 0x7d, 0x71, 0xb9, 0x39, 0x97, 0x45, 0xb8, 0x0d, 0x4b, 0x3d,
	0xa5, 0xf0, 0x92, 0xf0, 0x83, 0x2b, 0x69, 0xc5, 0x5c, 0x69, 0xb2, 0xda, 0x1c, 0x7e, 0x8b, 0x60,
	0xb5, 0xc6, 0xa7, 0x94, 0x1c, 0x3e, 0x4c, 0x94, 0xbe, 0x42, 0x3d, 0x26, 0xdb, 0x98, 0xc7, 0xef,
	0x11, 0xac, 0xc5, 0xd9, 0xe8, 0x67, 0xf0, 0xf8, 0x06, 0x46, 0x6e, 0x94, 0x48, 0x0a, 0x9f, 0x23,
	0x58, 0xa9, 0xf1, 0xd8, 0x12, 0xc2, 0xfb, 0x89, 0x36, 0x12, 0xcb, 0x2b, 0xd9, 0xc2, 0x02, 0x7e,
	0x0d, 0xff, 0xd5, 0xf8, 0x64, 0x9b, 0xe0, 0xbd, 0x64, 0xf9, 0xdf, 0x56, 0x50, 0xb2, 0x76, 0x1a,
	0xbf, 0x81, 0xe5, 0x31, 0xed, 0xfe, 0x1d, 0x54, 0xaf, 0xa7, 0x7e, 0xa3, 0xf8, 0x33, 0xf8, 0x23,
	0x82, 0xe2, 0xc4, 0xf2, 0x91, 0x3a, 0xc0, 0xea, 0x75, 0x63, 0x98, 0xda, 0x26, 0xc9, 0x86, 0xb2,
	0x85, 0x87, 0x17, 0x97, 0x9b, 0x32, 0xdc, 0x6b, 0x72, 0x2a, 0x92, 0xfe, 0x35, 0xf5, 0xdb, 0xe2,
	0xc3, 0xe5, 0x66, 0x2a, 0x8b, 0xf2, 0x68, 0xff, 0xc9, 0xcb, 0x23, 0x93, 0x89, 0x56, 0xb7, 0x29,
	0xeb, 0xb6, 0xa5, 0xb8, 0xdd, 0x0e, 0x75, 0x3a, 0x5b, 0x5b, 0x42, 0x31, 0xed, 0x4e, 0x8b, 0x3a,
	0xba, 0xa3, 0x9d, 0x0a, 0xc5, 0x27, 0x56, 0x9a, 0xa6, 0xab, 0xf4, 0xc9, 0x15, 0x9f, 0x5c, 0xf1,
	0xb6, 0x95, 0xe1, 0x42, 0xe7, 0xe8, 0xaf, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xe0, 0x31,
	0xfd, 0x74, 0x0a, 0x00, 0x00,
}
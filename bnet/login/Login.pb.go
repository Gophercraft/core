// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: github.com/Gophercraft/core/bnet/public_protos/Login.proto

package login

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

type FormType int32

const (
	FormType_LOGIN_FORM FormType = 1
)

var FormType_name = map[int32]string{
	1: "LOGIN_FORM",
}

var FormType_value = map[string]int32{
	"LOGIN_FORM": 1,
}

func (x FormType) Enum() *FormType {
	p := new(FormType)
	*p = x
	return p
}

func (x FormType) String() string {
	return proto.EnumName(FormType_name, int32(x))
}

func (x *FormType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FormType_value, data, "FormType")
	if err != nil {
		return err
	}
	*x = FormType(value)
	return nil
}

func (FormType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d28fea78584271b, []int{0}
}

type AuthenticationState int32

const (
	AuthenticationState_LOGIN         AuthenticationState = 1
	AuthenticationState_LEGAL         AuthenticationState = 2
	AuthenticationState_AUTHENTICATOR AuthenticationState = 3
	AuthenticationState_DONE          AuthenticationState = 4
)

var AuthenticationState_name = map[int32]string{
	1: "LOGIN",
	2: "LEGAL",
	3: "AUTHENTICATOR",
	4: "DONE",
}

var AuthenticationState_value = map[string]int32{
	"LOGIN":         1,
	"LEGAL":         2,
	"AUTHENTICATOR": 3,
	"DONE":          4,
}

func (x AuthenticationState) Enum() *AuthenticationState {
	p := new(AuthenticationState)
	*p = x
	return p
}

func (x AuthenticationState) String() string {
	return proto.EnumName(AuthenticationState_name, int32(x))
}

func (x *AuthenticationState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AuthenticationState_value, data, "AuthenticationState")
	if err != nil {
		return err
	}
	*x = AuthenticationState(value)
	return nil
}

func (AuthenticationState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d28fea78584271b, []int{1}
}

type ErrorResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d28fea78584271b, []int{0}
}

func (m *ErrorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorResponse.Unmarshal(m, b)
}
func (m *ErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorResponse.Marshal(b, m, deterministic)
}
func (m *ErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorResponse.Merge(m, src)
}
func (m *ErrorResponse) XXX_Size() int {
	return xxx_messageInfo_ErrorResponse.Size(m)
}
func (m *ErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorResponse proto.InternalMessageInfo

type FormInput struct {
	InputId              *string  `protobuf:"bytes,1,req,name=input_id,json=inputId" json:"input_id,omitempty"`
	Type                 *string  `protobuf:"bytes,2,req,name=type" json:"type,omitempty"`
	Label                *string  `protobuf:"bytes,3,req,name=label" json:"label,omitempty"`
	MaxLength            *uint32  `protobuf:"varint,4,opt,name=max_length,json=maxLength" json:"max_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FormInput) Reset()         { *m = FormInput{} }
func (m *FormInput) String() string { return proto.CompactTextString(m) }
func (*FormInput) ProtoMessage()    {}
func (*FormInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d28fea78584271b, []int{1}
}

func (m *FormInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FormInput.Unmarshal(m, b)
}
func (m *FormInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FormInput.Marshal(b, m, deterministic)
}
func (m *FormInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FormInput.Merge(m, src)
}
func (m *FormInput) XXX_Size() int {
	return xxx_messageInfo_FormInput.Size(m)
}
func (m *FormInput) XXX_DiscardUnknown() {
	xxx_messageInfo_FormInput.DiscardUnknown(m)
}

var xxx_messageInfo_FormInput proto.InternalMessageInfo

func (m *FormInput) GetInputId() string {
	if m != nil && m.InputId != nil {
		return *m.InputId
	}
	return ""
}

func (m *FormInput) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *FormInput) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *FormInput) GetMaxLength() uint32 {
	if m != nil && m.MaxLength != nil {
		return *m.MaxLength
	}
	return 0
}

type FormInputs struct {
	Type                 *FormType    `protobuf:"varint,1,req,name=type,enum=Battlenet.JSON.Login.FormType" json:"type,omitempty"`
	Inputs               []*FormInput `protobuf:"bytes,2,rep,name=inputs" json:"inputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FormInputs) Reset()         { *m = FormInputs{} }
func (m *FormInputs) String() string { return proto.CompactTextString(m) }
func (*FormInputs) ProtoMessage()    {}
func (*FormInputs) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d28fea78584271b, []int{2}
}

func (m *FormInputs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FormInputs.Unmarshal(m, b)
}
func (m *FormInputs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FormInputs.Marshal(b, m, deterministic)
}
func (m *FormInputs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FormInputs.Merge(m, src)
}
func (m *FormInputs) XXX_Size() int {
	return xxx_messageInfo_FormInputs.Size(m)
}
func (m *FormInputs) XXX_DiscardUnknown() {
	xxx_messageInfo_FormInputs.DiscardUnknown(m)
}

var xxx_messageInfo_FormInputs proto.InternalMessageInfo

func (m *FormInputs) GetType() FormType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return FormType_LOGIN_FORM
}

func (m *FormInputs) GetInputs() []*FormInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

type FormInputValue struct {
	InputId              *string  `protobuf:"bytes,1,req,name=input_id,json=inputId" json:"input_id,omitempty"`
	Value                *string  `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FormInputValue) Reset()         { *m = FormInputValue{} }
func (m *FormInputValue) String() string { return proto.CompactTextString(m) }
func (*FormInputValue) ProtoMessage()    {}
func (*FormInputValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d28fea78584271b, []int{3}
}

func (m *FormInputValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FormInputValue.Unmarshal(m, b)
}
func (m *FormInputValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FormInputValue.Marshal(b, m, deterministic)
}
func (m *FormInputValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FormInputValue.Merge(m, src)
}
func (m *FormInputValue) XXX_Size() int {
	return xxx_messageInfo_FormInputValue.Size(m)
}
func (m *FormInputValue) XXX_DiscardUnknown() {
	xxx_messageInfo_FormInputValue.DiscardUnknown(m)
}

var xxx_messageInfo_FormInputValue proto.InternalMessageInfo

func (m *FormInputValue) GetInputId() string {
	if m != nil && m.InputId != nil {
		return *m.InputId
	}
	return ""
}

func (m *FormInputValue) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type LoginForm struct {
	PlatformId           *string           `protobuf:"bytes,1,req,name=platform_id,json=platformId" json:"platform_id,omitempty"`
	ProgramId            *string           `protobuf:"bytes,2,req,name=program_id,json=programId" json:"program_id,omitempty"`
	Version              *string           `protobuf:"bytes,3,req,name=version" json:"version,omitempty"`
	Inputs               []*FormInputValue `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LoginForm) Reset()         { *m = LoginForm{} }
func (m *LoginForm) String() string { return proto.CompactTextString(m) }
func (*LoginForm) ProtoMessage()    {}
func (*LoginForm) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d28fea78584271b, []int{4}
}

func (m *LoginForm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginForm.Unmarshal(m, b)
}
func (m *LoginForm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginForm.Marshal(b, m, deterministic)
}
func (m *LoginForm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginForm.Merge(m, src)
}
func (m *LoginForm) XXX_Size() int {
	return xxx_messageInfo_LoginForm.Size(m)
}
func (m *LoginForm) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginForm.DiscardUnknown(m)
}

var xxx_messageInfo_LoginForm proto.InternalMessageInfo

func (m *LoginForm) GetPlatformId() string {
	if m != nil && m.PlatformId != nil {
		return *m.PlatformId
	}
	return ""
}

func (m *LoginForm) GetProgramId() string {
	if m != nil && m.ProgramId != nil {
		return *m.ProgramId
	}
	return ""
}

func (m *LoginForm) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *LoginForm) GetInputs() []*FormInputValue {
	if m != nil {
		return m.Inputs
	}
	return nil
}

type LoginResult struct {
	AuthenticationState  *AuthenticationState `protobuf:"varint,1,req,name=authentication_state,json=authenticationState,enum=Battlenet.JSON.Login.AuthenticationState" json:"authentication_state,omitempty"`
	ErrorCode            *string              `protobuf:"bytes,2,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	ErrorMessage         *string              `protobuf:"bytes,3,opt,name=error_message,json=errorMessage" json:"error_message,omitempty"`
	Url                  *string              `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
	LoginTicket          *string              `protobuf:"bytes,5,opt,name=login_ticket,json=loginTicket" json:"login_ticket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LoginResult) Reset()         { *m = LoginResult{} }
func (m *LoginResult) String() string { return proto.CompactTextString(m) }
func (*LoginResult) ProtoMessage()    {}
func (*LoginResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d28fea78584271b, []int{5}
}

func (m *LoginResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResult.Unmarshal(m, b)
}
func (m *LoginResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResult.Marshal(b, m, deterministic)
}
func (m *LoginResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResult.Merge(m, src)
}
func (m *LoginResult) XXX_Size() int {
	return xxx_messageInfo_LoginResult.Size(m)
}
func (m *LoginResult) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResult.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResult proto.InternalMessageInfo

func (m *LoginResult) GetAuthenticationState() AuthenticationState {
	if m != nil && m.AuthenticationState != nil {
		return *m.AuthenticationState
	}
	return AuthenticationState_LOGIN
}

func (m *LoginResult) GetErrorCode() string {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return ""
}

func (m *LoginResult) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

func (m *LoginResult) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *LoginResult) GetLoginTicket() string {
	if m != nil && m.LoginTicket != nil {
		return *m.LoginTicket
	}
	return ""
}

type LoginRefreshResult struct {
	LoginTicketExpiry    *uint64  `protobuf:"varint,1,req,name=login_ticket_expiry,json=loginTicketExpiry" json:"login_ticket_expiry,omitempty"`
	IsExpired            *bool    `protobuf:"varint,2,opt,name=is_expired,json=isExpired" json:"is_expired,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRefreshResult) Reset()         { *m = LoginRefreshResult{} }
func (m *LoginRefreshResult) String() string { return proto.CompactTextString(m) }
func (*LoginRefreshResult) ProtoMessage()    {}
func (*LoginRefreshResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d28fea78584271b, []int{6}
}

func (m *LoginRefreshResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRefreshResult.Unmarshal(m, b)
}
func (m *LoginRefreshResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRefreshResult.Marshal(b, m, deterministic)
}
func (m *LoginRefreshResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRefreshResult.Merge(m, src)
}
func (m *LoginRefreshResult) XXX_Size() int {
	return xxx_messageInfo_LoginRefreshResult.Size(m)
}
func (m *LoginRefreshResult) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRefreshResult.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRefreshResult proto.InternalMessageInfo

func (m *LoginRefreshResult) GetLoginTicketExpiry() uint64 {
	if m != nil && m.LoginTicketExpiry != nil {
		return *m.LoginTicketExpiry
	}
	return 0
}

func (m *LoginRefreshResult) GetIsExpired() bool {
	if m != nil && m.IsExpired != nil {
		return *m.IsExpired
	}
	return false
}

type GameAccountInfo struct {
	DisplayName          *string  `protobuf:"bytes,1,req,name=display_name,json=displayName" json:"display_name,omitempty"`
	Expansion            *uint32  `protobuf:"varint,2,req,name=expansion" json:"expansion,omitempty"`
	IsSuspended          *bool    `protobuf:"varint,3,opt,name=is_suspended,json=isSuspended" json:"is_suspended,omitempty"`
	IsBanned             *bool    `protobuf:"varint,4,opt,name=is_banned,json=isBanned" json:"is_banned,omitempty"`
	SuspensionExpires    *uint64  `protobuf:"varint,5,opt,name=suspension_expires,json=suspensionExpires" json:"suspension_expires,omitempty"`
	SuspensionReason     *string  `protobuf:"bytes,6,opt,name=suspension_reason,json=suspensionReason" json:"suspension_reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameAccountInfo) Reset()         { *m = GameAccountInfo{} }
func (m *GameAccountInfo) String() string { return proto.CompactTextString(m) }
func (*GameAccountInfo) ProtoMessage()    {}
func (*GameAccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d28fea78584271b, []int{7}
}

func (m *GameAccountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameAccountInfo.Unmarshal(m, b)
}
func (m *GameAccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameAccountInfo.Marshal(b, m, deterministic)
}
func (m *GameAccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameAccountInfo.Merge(m, src)
}
func (m *GameAccountInfo) XXX_Size() int {
	return xxx_messageInfo_GameAccountInfo.Size(m)
}
func (m *GameAccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GameAccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GameAccountInfo proto.InternalMessageInfo

func (m *GameAccountInfo) GetDisplayName() string {
	if m != nil && m.DisplayName != nil {
		return *m.DisplayName
	}
	return ""
}

func (m *GameAccountInfo) GetExpansion() uint32 {
	if m != nil && m.Expansion != nil {
		return *m.Expansion
	}
	return 0
}

func (m *GameAccountInfo) GetIsSuspended() bool {
	if m != nil && m.IsSuspended != nil {
		return *m.IsSuspended
	}
	return false
}

func (m *GameAccountInfo) GetIsBanned() bool {
	if m != nil && m.IsBanned != nil {
		return *m.IsBanned
	}
	return false
}

func (m *GameAccountInfo) GetSuspensionExpires() uint64 {
	if m != nil && m.SuspensionExpires != nil {
		return *m.SuspensionExpires
	}
	return 0
}

func (m *GameAccountInfo) GetSuspensionReason() string {
	if m != nil && m.SuspensionReason != nil {
		return *m.SuspensionReason
	}
	return ""
}

type GameAccountList struct {
	GameAccounts         []*GameAccountInfo `protobuf:"bytes,1,rep,name=game_accounts,json=gameAccounts" json:"game_accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GameAccountList) Reset()         { *m = GameAccountList{} }
func (m *GameAccountList) String() string { return proto.CompactTextString(m) }
func (*GameAccountList) ProtoMessage()    {}
func (*GameAccountList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d28fea78584271b, []int{8}
}

func (m *GameAccountList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameAccountList.Unmarshal(m, b)
}
func (m *GameAccountList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameAccountList.Marshal(b, m, deterministic)
}
func (m *GameAccountList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameAccountList.Merge(m, src)
}
func (m *GameAccountList) XXX_Size() int {
	return xxx_messageInfo_GameAccountList.Size(m)
}
func (m *GameAccountList) XXX_DiscardUnknown() {
	xxx_messageInfo_GameAccountList.DiscardUnknown(m)
}

var xxx_messageInfo_GameAccountList proto.InternalMessageInfo

func (m *GameAccountList) GetGameAccounts() []*GameAccountInfo {
	if m != nil {
		return m.GameAccounts
	}
	return nil
}

func init() {
	proto.RegisterEnum("Battlenet.JSON.Login.FormType", FormType_name, FormType_value)
	proto.RegisterEnum("Battlenet.JSON.Login.AuthenticationState", AuthenticationState_name, AuthenticationState_value)
	proto.RegisterType((*ErrorResponse)(nil), "Battlenet.JSON.Login.ErrorResponse")
	proto.RegisterType((*FormInput)(nil), "Battlenet.JSON.Login.FormInput")
	proto.RegisterType((*FormInputs)(nil), "Battlenet.JSON.Login.FormInputs")
	proto.RegisterType((*FormInputValue)(nil), "Battlenet.JSON.Login.FormInputValue")
	proto.RegisterType((*LoginForm)(nil), "Battlenet.JSON.Login.LoginForm")
	proto.RegisterType((*LoginResult)(nil), "Battlenet.JSON.Login.LoginResult")
	proto.RegisterType((*LoginRefreshResult)(nil), "Battlenet.JSON.Login.LoginRefreshResult")
	proto.RegisterType((*GameAccountInfo)(nil), "Battlenet.JSON.Login.GameAccountInfo")
	proto.RegisterType((*GameAccountList)(nil), "Battlenet.JSON.Login.GameAccountList")
}

func init() {
	proto.RegisterFile("github.com/Gophercraft/core/bnet/public_protos/Login.proto", fileDescriptor_8d28fea78584271b)
}

var fileDescriptor_8d28fea78584271b = []byte{
	// 729 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcd, 0x6e, 0xe3, 0x36,
	0x10, 0xc7, 0x21, 0xdb, 0xd9, 0xb5, 0xc6, 0x71, 0xd6, 0x61, 0x72, 0x50, 0x3f, 0x37, 0x75, 0x5b,
	0x20, 0xdd, 0x45, 0xed, 0x45, 0x2e, 0xbd, 0x14, 0x28, 0x9c, 0xad, 0x37, 0xf1, 0xc2, 0x6b, 0x03,
	0x8c, 0xdb, 0x43, 0xd1, 0x42, 0xa0, 0xa5, 0x89, 0x4d, 0x54, 0x22, 0x05, 0x92, 0x5a, 0xc4, 0xcf,
	0xd3, 0x57, 0xeb, 0xb9, 0xcf, 0x50, 0x70, 0x24, 0xc7, 0x49, 0xeb, 0x7e, 0xdc, 0x38, 0xbf, 0xf9,
	0xe0, 0xfc, 0x87, 0x23, 0xc1, 0x77, 0x2b, 0xe9, 0xd6, 0xe5, 0x72, 0x90, 0xe8, 0x7c, 0x68, 0xcb,
	0x02, 0x4d, 0xf1, 0xea, 0x95, 0x1b, 0xae, 0x74, 0xb1, 0x46, 0x93, 0x18, 0x71, 0xeb, 0x86, 0x4b,
	0x85, 0x6e, 0x58, 0x94, 0xcb, 0x4c, 0x26, 0x71, 0x61, 0xb4, 0xd3, 0x76, 0x38, 0xd5, 0x2b, 0xa9,
	0x06, 0x64, 0xb0, 0xd3, 0x4b, 0xe1, 0x5c, 0x86, 0x0a, 0xdd, 0xe0, 0xed, 0xcd, 0x7c, 0x36, 0x20,
	0x5f, 0xff, 0x19, 0x74, 0xc7, 0xc6, 0x68, 0xc3, 0xd1, 0x16, 0x5a, 0x59, 0xec, 0x6b, 0x08, 0xdf,
	0x68, 0x93, 0x4f, 0x54, 0x51, 0x3a, 0xf6, 0x01, 0xb4, 0xa5, 0x3f, 0xc4, 0x32, 0x8d, 0x82, 0xb3,
	0xc6, 0x79, 0xc8, 0x9f, 0x92, 0x3d, 0x49, 0x19, 0x83, 0x96, 0xdb, 0x14, 0x18, 0x35, 0x08, 0xd3,
	0x99, 0x9d, 0xc2, 0x41, 0x26, 0x96, 0x98, 0x45, 0x4d, 0x82, 0x95, 0xc1, 0x3e, 0x01, 0xc8, 0xc5,
	0x5d, 0x9c, 0xa1, 0x5a, 0xb9, 0x75, 0xd4, 0x3a, 0x0b, 0xce, 0xbb, 0x3c, 0xcc, 0xc5, 0xdd, 0x94,
	0x40, 0x7f, 0x03, 0x70, 0x7f, 0xa1, 0x65, 0x17, 0x75, 0x59, 0x7f, 0xdb, 0xd1, 0xc5, 0xa7, 0x83,
	0x7d, 0x4d, 0x0f, 0x7c, 0xfc, 0x62, 0x53, 0x60, 0x7d, 0xed, 0x37, 0xf0, 0x84, 0xba, 0xb2, 0x51,
	0xe3, 0xac, 0x79, 0xde, 0xb9, 0x78, 0xfe, 0xcf, 0x59, 0x74, 0x0b, 0xaf, 0xc3, 0xfb, 0x23, 0x38,
	0xba, 0x87, 0x3f, 0x8a, 0xac, 0xc4, 0x7f, 0x13, 0x7c, 0x0a, 0x07, 0xef, 0x7d, 0x4c, 0xad, 0xb8,
	0x32, 0xfa, 0xbf, 0x05, 0x10, 0x52, 0x79, 0x5f, 0x88, 0x3d, 0x87, 0x4e, 0x91, 0x09, 0x77, 0xab,
	0x4d, 0xbe, 0xab, 0x00, 0x5b, 0x34, 0x49, 0xfd, 0x2c, 0x0a, 0xa3, 0x57, 0x46, 0x90, 0xbf, 0xaa,
	0x14, 0xd6, 0x64, 0x92, 0xb2, 0x08, 0x9e, 0xbe, 0x47, 0x63, 0xa5, 0x56, 0xf5, 0x08, 0xb7, 0x26,
	0xfb, 0xf6, 0x5e, 0x63, 0x8b, 0x34, 0x7e, 0xf1, 0x1f, 0x1a, 0x49, 0xce, 0xbd, 0xd0, 0xdf, 0x03,
	0xe8, 0x50, 0x00, 0x47, 0x5b, 0x66, 0x8e, 0xfd, 0x0c, 0xa7, 0xa2, 0x74, 0x6b, 0x54, 0x4e, 0x26,
	0xc2, 0x49, 0xad, 0x62, 0xeb, 0x84, 0xdb, 0x4e, 0xfd, 0xab, 0xfd, 0xb5, 0x47, 0x8f, 0x32, 0x6e,
	0x7c, 0x02, 0x3f, 0x11, 0x7f, 0x87, 0x5e, 0x24, 0xfa, 0x9d, 0x8a, 0x13, 0x9d, 0xfa, 0x71, 0x05,
	0x5e, 0x24, 0x91, 0xd7, 0x3a, 0x45, 0xf6, 0x39, 0x74, 0x2b, 0x77, 0x8e, 0xd6, 0x8a, 0x15, 0x46,
	0x4d, 0x8a, 0x38, 0x24, 0xf8, 0xae, 0x62, 0xac, 0x07, 0xcd, 0xd2, 0x64, 0xb4, 0x2d, 0x21, 0xf7,
	0x47, 0xf6, 0x19, 0x1c, 0x66, 0xbe, 0x8f, 0xd8, 0xc9, 0xe4, 0x57, 0x74, 0xd1, 0x01, 0xb9, 0x3a,
	0xc4, 0x16, 0x84, 0xfa, 0x09, 0xb0, 0x5a, 0xe5, 0xad, 0x41, 0xbb, 0xae, 0xc5, 0x0e, 0xe0, 0xe4,
	0x61, 0x62, 0x8c, 0x77, 0x85, 0x34, 0x1b, 0xd2, 0xda, 0xe2, 0xc7, 0x0f, 0xf2, 0xc7, 0xe4, 0xf0,
	0xed, 0x4b, 0x5b, 0x45, 0x61, 0x4a, 0xed, 0xb7, 0x79, 0x28, 0xed, 0xb8, 0x02, 0xfd, 0x3f, 0x02,
	0x78, 0x76, 0x25, 0x72, 0x1c, 0x25, 0x89, 0x2e, 0x95, 0x9b, 0xa8, 0x5b, 0xed, 0x7b, 0x4b, 0xa5,
	0x2d, 0x32, 0xb1, 0x89, 0x95, 0xc8, 0xb1, 0x7e, 0xf8, 0x4e, 0xcd, 0x66, 0x22, 0x47, 0xf6, 0x31,
	0x84, 0x78, 0x57, 0x08, 0x45, 0x8f, 0xeb, 0x1f, 0xbe, 0xcb, 0x77, 0xc0, 0x17, 0x90, 0x36, 0xb6,
	0xa5, 0x2d, 0x50, 0xa5, 0x98, 0xd2, 0x48, 0xda, 0xbc, 0x23, 0xed, 0xcd, 0x16, 0xb1, 0x8f, 0x20,
	0x94, 0x36, 0x5e, 0x0a, 0xa5, 0x30, 0xa5, 0xb9, 0xb4, 0x79, 0x5b, 0xda, 0x4b, 0xb2, 0xd9, 0xd7,
	0xc0, 0xaa, 0x64, 0x5f, 0xad, 0xee, 0xdd, 0xd2, 0x88, 0x5a, 0xfc, 0x78, 0xe7, 0xa9, 0x34, 0x58,
	0xf6, 0x12, 0x1e, 0xc0, 0xd8, 0xa0, 0xb0, 0x5a, 0x45, 0x4f, 0x68, 0xa0, 0xbd, 0x9d, 0x83, 0x13,
	0xef, 0xff, 0xf2, 0x48, 0xef, 0x54, 0x5a, 0xc7, 0xde, 0x42, 0x77, 0x25, 0x72, 0x8c, 0x45, 0xc5,
	0x6c, 0x14, 0xd0, 0x52, 0x7e, 0xb9, 0x7f, 0x71, 0xfe, 0x32, 0x2d, 0x7e, 0xb8, 0xda, 0x01, 0xfb,
	0xe2, 0x43, 0x68, 0x6f, 0xbf, 0x67, 0x76, 0x04, 0x30, 0x9d, 0x5f, 0x4d, 0x66, 0xf1, 0x9b, 0x39,
	0x7f, 0xd7, 0x0b, 0x5e, 0x5c, 0xc3, 0xc9, 0x9e, 0xad, 0x63, 0x21, 0x1c, 0x50, 0x58, 0x2f, 0xa0,
	0xe3, 0xf8, 0x6a, 0x34, 0xed, 0x35, 0xd8, 0x31, 0x74, 0x47, 0x3f, 0x2c, 0xae, 0xc7, 0xb3, 0xc5,
	0xe4, 0xf5, 0x68, 0x31, 0xe7, 0xbd, 0x26, 0x6b, 0x43, 0xeb, 0xfb, 0xf9, 0x6c, 0xdc, 0x6b, 0x5d,
	0x0e, 0xaf, 0x1b, 0x3f, 0xbd, 0xfc, 0x7f, 0xbf, 0x50, 0xda, 0x88, 0x3f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xa1, 0xdc, 0x80, 0x76, 0x70, 0x05, 0x00, 0x00,
}

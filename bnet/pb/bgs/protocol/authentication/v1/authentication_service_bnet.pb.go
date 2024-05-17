// DO NOT EDIT: this file was auto-generated by Gophercraft/protoss

// Code generated by protoc-gen-go-bnet. DO NOT EDIT.
// versions:
// - protoc-gen-go-bnet v1.3.0
// - protoc             v4.25.2
// source: bgs/low/pb/client/authentication_service.proto

package v1

import (
	context "context"
	protocol "github.com/Gophercraft/core/bnet/pb/bgs/protocol"
	rpc "github.com/Gophercraft/core/bnet/rpc"
	codes "github.com/Gophercraft/core/bnet/rpc/codes"
	status "github.com/Gophercraft/core/bnet/rpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the bnet package it is being compiled against.
const _ = rpc.SupportPackageIsVersion1

const (
	AuthenticationListener_OnServerStateChange_FullMethodName  = "/bgs.protocol.authentication.v1.AuthenticationListener/OnServerStateChange"
	AuthenticationListener_OnLogonComplete_FullMethodName      = "/bgs.protocol.authentication.v1.AuthenticationListener/OnLogonComplete"
	AuthenticationListener_OnLogonUpdate_FullMethodName        = "/bgs.protocol.authentication.v1.AuthenticationListener/OnLogonUpdate"
	AuthenticationListener_OnVersionInfoUpdated_FullMethodName = "/bgs.protocol.authentication.v1.AuthenticationListener/OnVersionInfoUpdated"
	AuthenticationListener_OnLogonQueueUpdate_FullMethodName   = "/bgs.protocol.authentication.v1.AuthenticationListener/OnLogonQueueUpdate"
	AuthenticationListener_OnLogonQueueEnd_FullMethodName      = "/bgs.protocol.authentication.v1.AuthenticationListener/OnLogonQueueEnd"
)

// AuthenticationListenerClient is the client API for AuthenticationListener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationListenerClient interface {
	OnServerStateChange(ctx context.Context, in *ServerStateChangeRequest, opts ...rpc.CallOption) (*protocol.NO_RESPONSE, error)
	OnLogonComplete(ctx context.Context, in *LogonResult, opts ...rpc.CallOption) (*protocol.NO_RESPONSE, error)
	OnLogonUpdate(ctx context.Context, in *LogonUpdateRequest, opts ...rpc.CallOption) (*protocol.NO_RESPONSE, error)
	OnVersionInfoUpdated(ctx context.Context, in *VersionInfoNotification, opts ...rpc.CallOption) (*protocol.NO_RESPONSE, error)
	OnLogonQueueUpdate(ctx context.Context, in *LogonQueueUpdateRequest, opts ...rpc.CallOption) (*protocol.NO_RESPONSE, error)
	OnLogonQueueEnd(ctx context.Context, in *protocol.NoData, opts ...rpc.CallOption) (*protocol.NO_RESPONSE, error)
}

type authenticationListenerClient struct {
	cc rpc.ClientConnectionInterface
}

func NewAuthenticationListenerClient(cc rpc.ClientConnectionInterface) AuthenticationListenerClient {
	return &authenticationListenerClient{cc}
}

func (c *authenticationListenerClient) OnServerStateChange(ctx context.Context, in *ServerStateChangeRequest, opts ...rpc.CallOption) (*protocol.NO_RESPONSE, error) {
	out := new(protocol.NO_RESPONSE)
	err := c.cc.Invoke(ctx, AuthenticationListener_OnServerStateChange_FullMethodName, 0x71240E35, 4, false, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationListenerClient) OnLogonComplete(ctx context.Context, in *LogonResult, opts ...rpc.CallOption) (*protocol.NO_RESPONSE, error) {
	out := new(protocol.NO_RESPONSE)
	err := c.cc.Invoke(ctx, AuthenticationListener_OnLogonComplete_FullMethodName, 0x71240E35, 5, false, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationListenerClient) OnLogonUpdate(ctx context.Context, in *LogonUpdateRequest, opts ...rpc.CallOption) (*protocol.NO_RESPONSE, error) {
	out := new(protocol.NO_RESPONSE)
	err := c.cc.Invoke(ctx, AuthenticationListener_OnLogonUpdate_FullMethodName, 0x71240E35, 10, false, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationListenerClient) OnVersionInfoUpdated(ctx context.Context, in *VersionInfoNotification, opts ...rpc.CallOption) (*protocol.NO_RESPONSE, error) {
	out := new(protocol.NO_RESPONSE)
	err := c.cc.Invoke(ctx, AuthenticationListener_OnVersionInfoUpdated_FullMethodName, 0x71240E35, 11, false, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationListenerClient) OnLogonQueueUpdate(ctx context.Context, in *LogonQueueUpdateRequest, opts ...rpc.CallOption) (*protocol.NO_RESPONSE, error) {
	out := new(protocol.NO_RESPONSE)
	err := c.cc.Invoke(ctx, AuthenticationListener_OnLogonQueueUpdate_FullMethodName, 0x71240E35, 12, false, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationListenerClient) OnLogonQueueEnd(ctx context.Context, in *protocol.NoData, opts ...rpc.CallOption) (*protocol.NO_RESPONSE, error) {
	out := new(protocol.NO_RESPONSE)
	err := c.cc.Invoke(ctx, AuthenticationListener_OnLogonQueueEnd_FullMethodName, 0x71240E35, 13, false, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationListenerServer is the server API for AuthenticationListener service.
// All implementations must embed UnimplementedAuthenticationListenerServer
// for forward compatibility
type AuthenticationListenerServer interface {
	OnServerStateChange(context.Context, *ServerStateChangeRequest) (*protocol.NO_RESPONSE, error)
	OnLogonComplete(context.Context, *LogonResult) (*protocol.NO_RESPONSE, error)
	OnLogonUpdate(context.Context, *LogonUpdateRequest) (*protocol.NO_RESPONSE, error)
	OnVersionInfoUpdated(context.Context, *VersionInfoNotification) (*protocol.NO_RESPONSE, error)
	OnLogonQueueUpdate(context.Context, *LogonQueueUpdateRequest) (*protocol.NO_RESPONSE, error)
	OnLogonQueueEnd(context.Context, *protocol.NoData) (*protocol.NO_RESPONSE, error)
	mustEmbedUnimplementedAuthenticationListenerServer()
}

// UnimplementedAuthenticationListenerServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationListenerServer struct {
}

func (UnimplementedAuthenticationListenerServer) OnServerStateChange(context.Context, *ServerStateChangeRequest) (*protocol.NO_RESPONSE, error) {
	return nil, status.Errorf(codes.ERROR_RPC_NOT_IMPLEMENTED, "method OnServerStateChange not implemented")
}
func (UnimplementedAuthenticationListenerServer) OnLogonComplete(context.Context, *LogonResult) (*protocol.NO_RESPONSE, error) {
	return nil, status.Errorf(codes.ERROR_RPC_NOT_IMPLEMENTED, "method OnLogonComplete not implemented")
}
func (UnimplementedAuthenticationListenerServer) OnLogonUpdate(context.Context, *LogonUpdateRequest) (*protocol.NO_RESPONSE, error) {
	return nil, status.Errorf(codes.ERROR_RPC_NOT_IMPLEMENTED, "method OnLogonUpdate not implemented")
}
func (UnimplementedAuthenticationListenerServer) OnVersionInfoUpdated(context.Context, *VersionInfoNotification) (*protocol.NO_RESPONSE, error) {
	return nil, status.Errorf(codes.ERROR_RPC_NOT_IMPLEMENTED, "method OnVersionInfoUpdated not implemented")
}
func (UnimplementedAuthenticationListenerServer) OnLogonQueueUpdate(context.Context, *LogonQueueUpdateRequest) (*protocol.NO_RESPONSE, error) {
	return nil, status.Errorf(codes.ERROR_RPC_NOT_IMPLEMENTED, "method OnLogonQueueUpdate not implemented")
}
func (UnimplementedAuthenticationListenerServer) OnLogonQueueEnd(context.Context, *protocol.NoData) (*protocol.NO_RESPONSE, error) {
	return nil, status.Errorf(codes.ERROR_RPC_NOT_IMPLEMENTED, "method OnLogonQueueEnd not implemented")
}
func (UnimplementedAuthenticationListenerServer) mustEmbedUnimplementedAuthenticationListenerServer() {
}

// UnsafeAuthenticationListenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationListenerServer will
// result in compilation errors.
type UnsafeAuthenticationListenerServer interface {
	mustEmbedUnimplementedAuthenticationListenerServer()
}

func RegisterAuthenticationListenerServer(s rpc.ServiceRegistrar, srv AuthenticationListenerServer) {
	s.RegisterService(&AuthenticationListener_ServiceDesc, srv)
}

func _AuthenticationListener_OnServerStateChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor rpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerStateChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationListenerServer).OnServerStateChange(ctx, in)
	}
	info := &rpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationListener_OnServerStateChange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationListenerServer).OnServerStateChange(ctx, req.(*ServerStateChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationListener_OnLogonComplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor rpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogonResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationListenerServer).OnLogonComplete(ctx, in)
	}
	info := &rpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationListener_OnLogonComplete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationListenerServer).OnLogonComplete(ctx, req.(*LogonResult))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationListener_OnLogonUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor rpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogonUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationListenerServer).OnLogonUpdate(ctx, in)
	}
	info := &rpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationListener_OnLogonUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationListenerServer).OnLogonUpdate(ctx, req.(*LogonUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationListener_OnVersionInfoUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor rpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionInfoNotification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationListenerServer).OnVersionInfoUpdated(ctx, in)
	}
	info := &rpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationListener_OnVersionInfoUpdated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationListenerServer).OnVersionInfoUpdated(ctx, req.(*VersionInfoNotification))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationListener_OnLogonQueueUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor rpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogonQueueUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationListenerServer).OnLogonQueueUpdate(ctx, in)
	}
	info := &rpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationListener_OnLogonQueueUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationListenerServer).OnLogonQueueUpdate(ctx, req.(*LogonQueueUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationListener_OnLogonQueueEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor rpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol.NoData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationListenerServer).OnLogonQueueEnd(ctx, in)
	}
	info := &rpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationListener_OnLogonQueueEnd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationListenerServer).OnLogonQueueEnd(ctx, req.(*protocol.NoData))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationListener_ServiceDesc is the rpc.ServiceDesc for AuthenticationListener service.
// It's only intended for direct use with rpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationListener_ServiceDesc = rpc.ServiceDesc{
	ServiceHash: 0x71240E35,
	ServiceName: "bgs.protocol.authentication.v1.AuthenticationListener",
	HandlerType: (*AuthenticationListenerServer)(nil),
	Methods: []rpc.MethodDesc{
		{
			MethodName: "OnServerStateChange",
			MethodId:   4,
			NoResponse: true,
			Handler:    _AuthenticationListener_OnServerStateChange_Handler,
		},
		{
			MethodName: "OnLogonComplete",
			MethodId:   5,
			NoResponse: true,
			Handler:    _AuthenticationListener_OnLogonComplete_Handler,
		},
		{
			MethodName: "OnLogonUpdate",
			MethodId:   10,
			NoResponse: true,
			Handler:    _AuthenticationListener_OnLogonUpdate_Handler,
		},
		{
			MethodName: "OnVersionInfoUpdated",
			MethodId:   11,
			NoResponse: true,
			Handler:    _AuthenticationListener_OnVersionInfoUpdated_Handler,
		},
		{
			MethodName: "OnLogonQueueUpdate",
			MethodId:   12,
			NoResponse: true,
			Handler:    _AuthenticationListener_OnLogonQueueUpdate_Handler,
		},
		{
			MethodName: "OnLogonQueueEnd",
			MethodId:   13,
			NoResponse: true,
			Handler:    _AuthenticationListener_OnLogonQueueEnd_Handler,
		},
	},
	Metadata: "bgs/low/pb/client/authentication_service.proto",
}

const (
	AuthenticationService_Logon_FullMethodName                  = "/bgs.protocol.authentication.v1.AuthenticationService/Logon"
	AuthenticationService_VerifyWebCredentials_FullMethodName   = "/bgs.protocol.authentication.v1.AuthenticationService/VerifyWebCredentials"
	AuthenticationService_GenerateWebCredentials_FullMethodName = "/bgs.protocol.authentication.v1.AuthenticationService/GenerateWebCredentials"
)

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	Logon(ctx context.Context, in *LogonRequest, opts ...rpc.CallOption) (*protocol.NoData, error)
	VerifyWebCredentials(ctx context.Context, in *VerifyWebCredentialsRequest, opts ...rpc.CallOption) (*protocol.NoData, error)
	GenerateWebCredentials(ctx context.Context, in *GenerateWebCredentialsRequest, opts ...rpc.CallOption) (*GenerateWebCredentialsResponse, error)
}

type authenticationServiceClient struct {
	cc rpc.ClientConnectionInterface
}

func NewAuthenticationServiceClient(cc rpc.ClientConnectionInterface) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) Logon(ctx context.Context, in *LogonRequest, opts ...rpc.CallOption) (*protocol.NoData, error) {
	out := new(protocol.NoData)
	err := c.cc.Invoke(ctx, AuthenticationService_Logon_FullMethodName, 0x0DECFC01, 1, true, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) VerifyWebCredentials(ctx context.Context, in *VerifyWebCredentialsRequest, opts ...rpc.CallOption) (*protocol.NoData, error) {
	out := new(protocol.NoData)
	err := c.cc.Invoke(ctx, AuthenticationService_VerifyWebCredentials_FullMethodName, 0x0DECFC01, 7, true, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) GenerateWebCredentials(ctx context.Context, in *GenerateWebCredentialsRequest, opts ...rpc.CallOption) (*GenerateWebCredentialsResponse, error) {
	out := new(GenerateWebCredentialsResponse)
	err := c.cc.Invoke(ctx, AuthenticationService_GenerateWebCredentials_FullMethodName, 0x0DECFC01, 8, true, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
// All implementations must embed UnimplementedAuthenticationServiceServer
// for forward compatibility
type AuthenticationServiceServer interface {
	Logon(context.Context, *LogonRequest) (*protocol.NoData, error)
	VerifyWebCredentials(context.Context, *VerifyWebCredentialsRequest) (*protocol.NoData, error)
	GenerateWebCredentials(context.Context, *GenerateWebCredentialsRequest) (*GenerateWebCredentialsResponse, error)
	mustEmbedUnimplementedAuthenticationServiceServer()
}

// UnimplementedAuthenticationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (UnimplementedAuthenticationServiceServer) Logon(context.Context, *LogonRequest) (*protocol.NoData, error) {
	return nil, status.Errorf(codes.ERROR_RPC_NOT_IMPLEMENTED, "method Logon not implemented")
}
func (UnimplementedAuthenticationServiceServer) VerifyWebCredentials(context.Context, *VerifyWebCredentialsRequest) (*protocol.NoData, error) {
	return nil, status.Errorf(codes.ERROR_RPC_NOT_IMPLEMENTED, "method VerifyWebCredentials not implemented")
}
func (UnimplementedAuthenticationServiceServer) GenerateWebCredentials(context.Context, *GenerateWebCredentialsRequest) (*GenerateWebCredentialsResponse, error) {
	return nil, status.Errorf(codes.ERROR_RPC_NOT_IMPLEMENTED, "method GenerateWebCredentials not implemented")
}
func (UnimplementedAuthenticationServiceServer) mustEmbedUnimplementedAuthenticationServiceServer() {}

// UnsafeAuthenticationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServiceServer will
// result in compilation errors.
type UnsafeAuthenticationServiceServer interface {
	mustEmbedUnimplementedAuthenticationServiceServer()
}

func RegisterAuthenticationServiceServer(s rpc.ServiceRegistrar, srv AuthenticationServiceServer) {
	s.RegisterService(&AuthenticationService_ServiceDesc, srv)
}

func _AuthenticationService_Logon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor rpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).Logon(ctx, in)
	}
	info := &rpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_Logon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).Logon(ctx, req.(*LogonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_VerifyWebCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor rpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyWebCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).VerifyWebCredentials(ctx, in)
	}
	info := &rpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_VerifyWebCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).VerifyWebCredentials(ctx, req.(*VerifyWebCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_GenerateWebCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor rpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateWebCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).GenerateWebCredentials(ctx, in)
	}
	info := &rpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_GenerateWebCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).GenerateWebCredentials(ctx, req.(*GenerateWebCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationService_ServiceDesc is the rpc.ServiceDesc for AuthenticationService service.
// It's only intended for direct use with rpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationService_ServiceDesc = rpc.ServiceDesc{
	ServiceHash: 0x0DECFC01,
	ServiceName: "bgs.protocol.authentication.v1.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []rpc.MethodDesc{
		{
			MethodName: "Logon",
			MethodId:   1,
			NoResponse: false,
			Handler:    _AuthenticationService_Logon_Handler,
		},
		{
			MethodName: "VerifyWebCredentials",
			MethodId:   7,
			NoResponse: false,
			Handler:    _AuthenticationService_VerifyWebCredentials_Handler,
		},
		{
			MethodName: "GenerateWebCredentials",
			MethodId:   8,
			NoResponse: false,
			Handler:    _AuthenticationService_GenerateWebCredentials_Handler,
		},
	},
	Metadata: "bgs/low/pb/client/authentication_service.proto",
}
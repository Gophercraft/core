// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: admin.proto

package admin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AdminService_BanAccount_FullMethodName           = "/admin.AdminService/BanAccount"
	AdminService_UnbanAccount_FullMethodName         = "/admin.AdminService/UnbanAccount"
	AdminService_SuspendAccount_FullMethodName       = "/admin.AdminService/SuspendAccount"
	AdminService_UnsuspendAccount_FullMethodName     = "/admin.AdminService/UnsuspendAccount"
	AdminService_LockAccount_FullMethodName          = "/admin.AdminService/LockAccount"
	AdminService_UnlockAccount_FullMethodName        = "/admin.AdminService/UnlockAccount"
	AdminService_BanGameAccount_FullMethodName       = "/admin.AdminService/BanGameAccount"
	AdminService_UnbanGameAccount_FullMethodName     = "/admin.AdminService/UnbanGameAccount"
	AdminService_SuspendGameAccount_FullMethodName   = "/admin.AdminService/SuspendGameAccount"
	AdminService_UnsuspendGameAccount_FullMethodName = "/admin.AdminService/UnsuspendGameAccount"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	BanAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*BanStatus, error)
	UnbanAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*BanStatus, error)
	SuspendAccount(ctx context.Context, in *SuspendAccountRequest, opts ...grpc.CallOption) (*BanStatus, error)
	UnsuspendAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*BanStatus, error)
	LockAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*LockStatus, error)
	UnlockAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*LockStatus, error)
	BanGameAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*BanStatus, error)
	UnbanGameAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*BanStatus, error)
	SuspendGameAccount(ctx context.Context, in *SuspendAccountRequest, opts ...grpc.CallOption) (*BanStatus, error)
	UnsuspendGameAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*BanStatus, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) BanAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*BanStatus, error) {
	out := new(BanStatus)
	err := c.cc.Invoke(ctx, AdminService_BanAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UnbanAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*BanStatus, error) {
	out := new(BanStatus)
	err := c.cc.Invoke(ctx, AdminService_UnbanAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) SuspendAccount(ctx context.Context, in *SuspendAccountRequest, opts ...grpc.CallOption) (*BanStatus, error) {
	out := new(BanStatus)
	err := c.cc.Invoke(ctx, AdminService_SuspendAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UnsuspendAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*BanStatus, error) {
	out := new(BanStatus)
	err := c.cc.Invoke(ctx, AdminService_UnsuspendAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) LockAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*LockStatus, error) {
	out := new(LockStatus)
	err := c.cc.Invoke(ctx, AdminService_LockAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UnlockAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*LockStatus, error) {
	out := new(LockStatus)
	err := c.cc.Invoke(ctx, AdminService_UnlockAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) BanGameAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*BanStatus, error) {
	out := new(BanStatus)
	err := c.cc.Invoke(ctx, AdminService_BanGameAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UnbanGameAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*BanStatus, error) {
	out := new(BanStatus)
	err := c.cc.Invoke(ctx, AdminService_UnbanGameAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) SuspendGameAccount(ctx context.Context, in *SuspendAccountRequest, opts ...grpc.CallOption) (*BanStatus, error) {
	out := new(BanStatus)
	err := c.cc.Invoke(ctx, AdminService_SuspendGameAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UnsuspendGameAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*BanStatus, error) {
	out := new(BanStatus)
	err := c.cc.Invoke(ctx, AdminService_UnsuspendGameAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	BanAccount(context.Context, *AccountRequest) (*BanStatus, error)
	UnbanAccount(context.Context, *AccountRequest) (*BanStatus, error)
	SuspendAccount(context.Context, *SuspendAccountRequest) (*BanStatus, error)
	UnsuspendAccount(context.Context, *AccountRequest) (*BanStatus, error)
	LockAccount(context.Context, *AccountRequest) (*LockStatus, error)
	UnlockAccount(context.Context, *AccountRequest) (*LockStatus, error)
	BanGameAccount(context.Context, *AccountRequest) (*BanStatus, error)
	UnbanGameAccount(context.Context, *AccountRequest) (*BanStatus, error)
	SuspendGameAccount(context.Context, *SuspendAccountRequest) (*BanStatus, error)
	UnsuspendGameAccount(context.Context, *AccountRequest) (*BanStatus, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) BanAccount(context.Context, *AccountRequest) (*BanStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BanAccount not implemented")
}
func (UnimplementedAdminServiceServer) UnbanAccount(context.Context, *AccountRequest) (*BanStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbanAccount not implemented")
}
func (UnimplementedAdminServiceServer) SuspendAccount(context.Context, *SuspendAccountRequest) (*BanStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuspendAccount not implemented")
}
func (UnimplementedAdminServiceServer) UnsuspendAccount(context.Context, *AccountRequest) (*BanStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsuspendAccount not implemented")
}
func (UnimplementedAdminServiceServer) LockAccount(context.Context, *AccountRequest) (*LockStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockAccount not implemented")
}
func (UnimplementedAdminServiceServer) UnlockAccount(context.Context, *AccountRequest) (*LockStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlockAccount not implemented")
}
func (UnimplementedAdminServiceServer) BanGameAccount(context.Context, *AccountRequest) (*BanStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BanGameAccount not implemented")
}
func (UnimplementedAdminServiceServer) UnbanGameAccount(context.Context, *AccountRequest) (*BanStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbanGameAccount not implemented")
}
func (UnimplementedAdminServiceServer) SuspendGameAccount(context.Context, *SuspendAccountRequest) (*BanStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuspendGameAccount not implemented")
}
func (UnimplementedAdminServiceServer) UnsuspendGameAccount(context.Context, *AccountRequest) (*BanStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsuspendGameAccount not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_BanAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).BanAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_BanAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).BanAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UnbanAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UnbanAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_UnbanAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UnbanAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_SuspendAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuspendAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).SuspendAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_SuspendAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).SuspendAccount(ctx, req.(*SuspendAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UnsuspendAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UnsuspendAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_UnsuspendAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UnsuspendAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_LockAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).LockAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_LockAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).LockAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UnlockAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UnlockAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_UnlockAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UnlockAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_BanGameAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).BanGameAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_BanGameAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).BanGameAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UnbanGameAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UnbanGameAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_UnbanGameAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UnbanGameAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_SuspendGameAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuspendAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).SuspendGameAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_SuspendGameAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).SuspendGameAccount(ctx, req.(*SuspendAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UnsuspendGameAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UnsuspendGameAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_UnsuspendGameAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UnsuspendGameAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BanAccount",
			Handler:    _AdminService_BanAccount_Handler,
		},
		{
			MethodName: "UnbanAccount",
			Handler:    _AdminService_UnbanAccount_Handler,
		},
		{
			MethodName: "SuspendAccount",
			Handler:    _AdminService_SuspendAccount_Handler,
		},
		{
			MethodName: "UnsuspendAccount",
			Handler:    _AdminService_UnsuspendAccount_Handler,
		},
		{
			MethodName: "LockAccount",
			Handler:    _AdminService_LockAccount_Handler,
		},
		{
			MethodName: "UnlockAccount",
			Handler:    _AdminService_UnlockAccount_Handler,
		},
		{
			MethodName: "BanGameAccount",
			Handler:    _AdminService_BanGameAccount_Handler,
		},
		{
			MethodName: "UnbanGameAccount",
			Handler:    _AdminService_UnbanGameAccount_Handler,
		},
		{
			MethodName: "SuspendGameAccount",
			Handler:    _AdminService_SuspendGameAccount_Handler,
		},
		{
			MethodName: "UnsuspendGameAccount",
			Handler:    _AdminService_UnsuspendGameAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
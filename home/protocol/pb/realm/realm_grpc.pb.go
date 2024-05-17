// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: realm.proto

package realm

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
	RealmService_Enlist_FullMethodName                    = "/realm.RealmService/Enlist"
	RealmService_Announce_FullMethodName                  = "/realm.RealmService/Announce"
	RealmService_SaveCharacterCount_FullMethodName        = "/realm.RealmService/SaveCharacterCount"
	RealmService_SaveLastCharacterLoggedIn_FullMethodName = "/realm.RealmService/SaveLastCharacterLoggedIn"
)

// RealmServiceClient is the client API for RealmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RealmServiceClient interface {
	Enlist(ctx context.Context, in *EnlistRequest, opts ...grpc.CallOption) (*EnlistResponse, error)
	// This call checks the ECDSA fingerprint of the caller to determine if valid
	Announce(ctx context.Context, in *AnnounceRequest, opts ...grpc.CallOption) (*AnnounceResponse, error)
	// Sent whenever the character list is changed
	SaveCharacterCount(ctx context.Context, in *CharacterCountData, opts ...grpc.CallOption) (*SaveResponse, error)
	// Sent whenever a character logs in
	SaveLastCharacterLoggedIn(ctx context.Context, in *CharacterLoggedInData, opts ...grpc.CallOption) (*SaveResponse, error)
}

type realmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRealmServiceClient(cc grpc.ClientConnInterface) RealmServiceClient {
	return &realmServiceClient{cc}
}

func (c *realmServiceClient) Enlist(ctx context.Context, in *EnlistRequest, opts ...grpc.CallOption) (*EnlistResponse, error) {
	out := new(EnlistResponse)
	err := c.cc.Invoke(ctx, RealmService_Enlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmServiceClient) Announce(ctx context.Context, in *AnnounceRequest, opts ...grpc.CallOption) (*AnnounceResponse, error) {
	out := new(AnnounceResponse)
	err := c.cc.Invoke(ctx, RealmService_Announce_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmServiceClient) SaveCharacterCount(ctx context.Context, in *CharacterCountData, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, RealmService_SaveCharacterCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmServiceClient) SaveLastCharacterLoggedIn(ctx context.Context, in *CharacterLoggedInData, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, RealmService_SaveLastCharacterLoggedIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RealmServiceServer is the server API for RealmService service.
// All implementations must embed UnimplementedRealmServiceServer
// for forward compatibility
type RealmServiceServer interface {
	Enlist(context.Context, *EnlistRequest) (*EnlistResponse, error)
	// This call checks the ECDSA fingerprint of the caller to determine if valid
	Announce(context.Context, *AnnounceRequest) (*AnnounceResponse, error)
	// Sent whenever the character list is changed
	SaveCharacterCount(context.Context, *CharacterCountData) (*SaveResponse, error)
	// Sent whenever a character logs in
	SaveLastCharacterLoggedIn(context.Context, *CharacterLoggedInData) (*SaveResponse, error)
	mustEmbedUnimplementedRealmServiceServer()
}

// UnimplementedRealmServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRealmServiceServer struct {
}

func (UnimplementedRealmServiceServer) Enlist(context.Context, *EnlistRequest) (*EnlistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enlist not implemented")
}
func (UnimplementedRealmServiceServer) Announce(context.Context, *AnnounceRequest) (*AnnounceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Announce not implemented")
}
func (UnimplementedRealmServiceServer) SaveCharacterCount(context.Context, *CharacterCountData) (*SaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveCharacterCount not implemented")
}
func (UnimplementedRealmServiceServer) SaveLastCharacterLoggedIn(context.Context, *CharacterLoggedInData) (*SaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveLastCharacterLoggedIn not implemented")
}
func (UnimplementedRealmServiceServer) mustEmbedUnimplementedRealmServiceServer() {}

// UnsafeRealmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RealmServiceServer will
// result in compilation errors.
type UnsafeRealmServiceServer interface {
	mustEmbedUnimplementedRealmServiceServer()
}

func RegisterRealmServiceServer(s grpc.ServiceRegistrar, srv RealmServiceServer) {
	s.RegisterService(&RealmService_ServiceDesc, srv)
}

func _RealmService_Enlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmServiceServer).Enlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmService_Enlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmServiceServer).Enlist(ctx, req.(*EnlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmService_Announce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnnounceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmServiceServer).Announce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmService_Announce_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmServiceServer).Announce(ctx, req.(*AnnounceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmService_SaveCharacterCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CharacterCountData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmServiceServer).SaveCharacterCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmService_SaveCharacterCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmServiceServer).SaveCharacterCount(ctx, req.(*CharacterCountData))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmService_SaveLastCharacterLoggedIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CharacterLoggedInData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmServiceServer).SaveLastCharacterLoggedIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmService_SaveLastCharacterLoggedIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmServiceServer).SaveLastCharacterLoggedIn(ctx, req.(*CharacterLoggedInData))
	}
	return interceptor(ctx, in, info, handler)
}

// RealmService_ServiceDesc is the grpc.ServiceDesc for RealmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RealmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "realm.RealmService",
	HandlerType: (*RealmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enlist",
			Handler:    _RealmService_Enlist_Handler,
		},
		{
			MethodName: "Announce",
			Handler:    _RealmService_Announce_Handler,
		},
		{
			MethodName: "SaveCharacterCount",
			Handler:    _RealmService_SaveCharacterCount_Handler,
		},
		{
			MethodName: "SaveLastCharacterLoggedIn",
			Handler:    _RealmService_SaveLastCharacterLoggedIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "realm.proto",
}
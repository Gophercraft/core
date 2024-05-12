package rpc

import (
	"context"

	"github.com/Gophercraft/core/bnet/util"
)

type ServiceRegistrar interface {
	RegisterService(descriptor *ServiceDesc, service any) error
}

// UnaryServerInfo consists of various information about a unary RPC on
// server side. All per-rpc information may be mutated by the interceptor.
type UnaryServerInfo struct {
	// Server is the service implementation the user provides. This is read-only.
	Server any
	// FullMethod is the full RPC method string, i.e., /package.service/method.
	FullMethod  string
	ServiceHash util.ServiceHash
	Method      Method
}

// UnaryInvoker is called by UnaryClientInterceptor to complete RPCs.
type UnaryInvoker func(ctx context.Context, service_hash util.ServiceHash, method Method, wait_for_response bool, req, reply any, cc ClientConnectionInterface, opts ...CallOption) error

// UnaryClientInterceptor intercepts the execution of a unary RPC on the client.
// Unary interceptors can be specified as a DialOption, using
// WithUnaryInterceptor() or WithChainUnaryInterceptor(), when creating a
// ClientConn. When a unary interceptor(s) is set on a ClientConn, gRPC
// delegates all unary RPC invocations to the interceptor, and it is the
// responsibility of the interceptor to call invoker to complete the processing
// of the RPC.
//
// method is the RPC name. req and reply are the corresponding request and
// response messages. cc is the ClientConn on which the RPC was invoked. invoker
// is the handler to complete the RPC and it is the responsibility of the
// interceptor to call it. opts contain all applicable call options, including
// defaults from the ClientConn as well as per-call options.
//
// The returned error must be compatible with the status package.
type UnaryClientInterceptor func(ctx context.Context, full_method string, service_hash util.ServiceHash, method Method, wait_for_response bool, req, reply any, cc ClientConnectionInterface, invoker UnaryInvoker, opts ...CallOption) error

// UnaryHandler defines the handler invoked by UnaryServerInterceptor to complete the normal
// execution of a unary RPC.
//
// If a UnaryHandler returns an error, it should either be produced by the
// status package, or be one of the context errors. Otherwise, gRPC will use
// codes.Unknown as the status code and err.Error() as the status message of the
// RPC.
type UnaryHandler func(ctx context.Context, req any) (any, error)

type UnaryServerInterceptor func(ctx context.Context, req any, info *UnaryServerInfo, handler UnaryHandler) (resp any, err error)

type methodHandler func(srv any, ctx context.Context, dec func(any) error, interceptor UnaryServerInterceptor) (any, error)

type MethodDesc struct {
	MethodName string
	MethodId   Method
	NoResponse bool
	Handler    methodHandler
}

type ServiceDesc struct {
	ServiceName string
	ServiceHash util.ServiceHash
	// The pointer to the service interface. Used to check whether the user
	// provided implementation satisfies the interface requirements.
	HandlerType any
	Methods     []MethodDesc
	Metadata    any
}

package bnet

import (
	"context"
	"errors"
	"net"

	"github.com/Gophercraft/core/bnet/rpc"
	"github.com/Gophercraft/core/bnet/util"
)

type registered_service struct {
	descriptor *rpc.ServiceDesc
	service    any
}

type Server struct {
	// root context
	ctx context.Context

	// TLS listener
	listener net.Listener

	// registered services
	services map[util.ServiceHash]*registered_service

	// default server interceptor
	unary_server_interceptor rpc.UnaryServerInterceptor

	// default client interceptor
	unary_client_interceptor rpc.UnaryClientInterceptor
}

type ServerOption func(server *Server)

// Intercept server's
func WithUnaryServerInterceptor(unary_server_interceptor rpc.UnaryServerInterceptor) ServerOption {
	return func(server *Server) {
		server.unary_server_interceptor = unary_server_interceptor
	}
}

// Intercept calls
func WithUnaryClientInterceptor(unary_client_interceptor rpc.UnaryClientInterceptor) ServerOption {
	return func(server *Server) {
		server.unary_client_interceptor = unary_client_interceptor
	}
}

func NewServer(options ...ServerOption) (server *Server) {
	server = new(Server)
	server.ctx = context.TODO()
	server.services = make(map[util.ServiceHash]*registered_service)

	for _, option := range options {
		option(server)
	}
	return
}

func (server *Server) SetListener(listener net.Listener) {
	server.listener = listener
}

func (server *Server) Serve() (err error) {
	for {
		var (
			tls_connection net.Conn
		)
		tls_connection, err = server.listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			return
		}

		server.dispatch_incoming_connection(tls_connection)
	}
}

func (server *Server) Stop() (err error) {
	// TODO: context cancel
	return server.listener.Close()
}

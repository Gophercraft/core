// Package bnet_rpc implements a BattleNet/BGS RPC server. In newer clients, this protocol is used instead of Grunt.
package bnet_rpc

import (
	"context"
	"crypto/tls"
	"net"

	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/core/bnet"
	"github.com/Gophercraft/core/bnet/rpc"
	"github.com/Gophercraft/core/bnet/util"
	"github.com/Gophercraft/log"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type ServiceConfig struct {
	Address string
	TLS     *tls.Config
}

type Service struct {
	config *ServiceConfig
	server *bnet.Server
}

func server_interceptor(ctx context.Context, req any, info *rpc.UnaryServerInfo, handler rpc.UnaryHandler) (resp any, err error) {
	log.Println("client requested", info.FullMethod)
	log.Println("client args = ", protojson.Format(req.(proto.Message)))
	resp, err = handler(ctx, req)
	if err != nil {
		log.Warn("rpc returned error:", err)
	}
	if resp != nil {
		log.Println("server replied", protojson.Format(resp.(proto.Message)))
	}
	return resp, err
}

func client_interceptor(ctx context.Context, full_method string, service_hash util.ServiceHash, method rpc.Method, wait_for_response bool, req, reply any, cc rpc.ClientConnectionInterface, invoker rpc.UnaryInvoker, opts ...rpc.CallOption) error {
	log.Println("calling", full_method, "from client")
	log.Println("with arguments =", protojson.Format(req.(proto.Message)))
	err := invoker(ctx, service_hash, method, wait_for_response, req, reply, cc, opts...)
	if err != nil {
		log.Warn("client returned error:", err)
	}
	log.Println("client replied =", protojson.Format(reply.(proto.Message)))
	return err
}

func New(service_config *ServiceConfig) (service *Service) {
	service = new(Service)
	service.server = bnet.NewServer(
		bnet.WithUnaryServerInterceptor(server_interceptor),
		bnet.WithUnaryClientInterceptor(client_interceptor))
	if service_config.TLS == nil {
		service_config.TLS = util.GetTrinityTLSConfig()
	}
	service.config = service_config
	return service
}

func (service *Service) RegisterService(descriptor *rpc.ServiceDesc, service_impl any) (err error) {
	return service.server.RegisterService(descriptor, service_impl)
}

func (service *Service) ID() config.HomeServiceID {
	return config.BNetRPCService
}

func (service *Service) Start() (err error) {
	var (
		listener net.Listener
	)
	listener, err = tls.Listen("tcp", service.config.Address, service.config.TLS)
	if err != nil {
		return
	}

	service.server.SetListener(listener)

	go func() {
		if err := service.server.Serve(); err != nil {
			log.Warn("home/service/bnet_rpc: error serving:", err)
		}
	}()

	return
}

func (service *Service) Stop() (err error) {
	err = service.server.Stop()
	return
}

// Package core provides a home.Service for a GPRC server
package core

import (
	"crypto/tls"
	"net"

	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Service struct {
	config      *ServiceConfig
	grpc_server *grpc.Server
	listener    net.Listener
}

type ServiceConfig struct {
	Address string
	TLS     *tls.Config
}

func New(service_config *ServiceConfig) (service *Service) {
	service = new(Service)
	service.config = service_config
	// use TLS config
	credential_option := grpc.Creds(credentials.NewTLS(service.config.TLS))
	service.grpc_server = grpc.NewServer(credential_option)
	return
}

var _ grpc.ServiceRegistrar = (*Service)(nil)

func (service *Service) RegisterService(service_desc *grpc.ServiceDesc, service_impl any) {
	service.grpc_server.RegisterService(service_desc, service_impl)
}

func (service *Service) Start() (err error) {
	service.listener, err = net.Listen("tcp", service.config.Address)
	if err != nil {
		return
	}

	go func() {
		if err = service.grpc_server.Serve(service.listener); err != nil {
			log.Warn(err)
		}
	}()

	return
}

func (service *Service) Stop() (err error) {
	service.grpc_server.GracefulStop()
	return nil
}

func (service *Service) ID() config.HomeServiceID {
	return config.CoreService
}

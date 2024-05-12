package bnet

import (
	"github.com/Gophercraft/core/bnet/rpc"
	"github.com/Gophercraft/core/bnet/util"
)

func (server *Server) RegisterService(descriptor *rpc.ServiceDesc, service any) (err error) {
	server.services[descriptor.ServiceHash] = &registered_service{
		descriptor: descriptor,
		service:    service,
	}
	return
}

func (server *Server) get_registered_service(hash util.ServiceHash) (service *registered_service) {
	service = server.services[hash]
	return
}

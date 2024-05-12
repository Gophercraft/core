package auth

import (
	"github.com/Gophercraft/core/game/network"
	"github.com/Gophercraft/core/game/protocol/message"
)

var (
	auth_service_id = network.RegisterServiceID("core/game/server/service/auth")
)

type ServiceConfig struct {
	RealmID uint64
}

type wait_queue struct {
}

type Service struct {
	config   *ServiceConfig
	provider ServiceProvider
	state    service_state
}

func (service *Service) ID() network.ServiceID {
	return auth_service_id
}

func (service *Service) Init(server *network.Server) error {
	// SMSG_AUTH_CHALLENGE should be sent upon the creation of a network session
	server.HandleNewSessionState(network.NewConnection, func(newstate network.SessionState, session *network.Session) {
		if err := start_auth_challenge(session); err != nil {
			session.Terminate(err)
		}
	})

	return nil
}

func (service *Service) Descriptor() (desc *network.ServiceDescriptor) {
	// Setup network handlers
	desc = network.NewServiceDescriptor()
	desc.Handle(message.CMSG_AUTH_SESSION, network.DecodeHandler(service.handle_client_auth_session))
	return
}

func New(config *ServiceConfig, provider ServiceProvider) (service *Service) {
	service = new(Service)
	service.config = config
	service.provider = provider
	return
}

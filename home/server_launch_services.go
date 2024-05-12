package home

import (
	"crypto/tls"
	"fmt"

	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/log"
)

func (server *Server) create_tls_config() (tls_config *tls.Config) {
	tls_config = &tls.Config{
		Certificates: []tls.Certificate{server.home_config.Certificate},
		MinVersion:   tls.VersionTLS12,
		ClientAuth:   tls.RequireAnyClientCert,
	}
	return
}

// Register a component service of the Home server according to a config declaration
func (server *Server) create_home_service(service_config *config.HomeServiceConfig) (err error) {
	// Create one multiple services listed in the config file

	// Services listen on a particular address
	// handling the specifications of how their particular protocol communicates
	// whereas the service's "provider" interface handles the internal application logic, saving and loading information into the database

	var service Service
	switch service_config.Service {
	case config.CoreService:
		service, err = server.create_core_service(service_config)
	case config.OldRealmlistService:
		service, err = server.create_old_realmlist_service(service_config)
	case config.GruntService:
		service, err = server.create_grunt_service(service_config)
	case config.BNetRESTService:
		service, err = server.create_bnet_rest_service(service_config)
	case config.BNetRPCService:
		service, err = server.create_bnet_rpc_service(service_config)
	case config.WebService:
		service, err = server.create_web_service(service_config)
	}

	if service == nil {
		err = fmt.Errorf("home: no service could be installed for %s", service_config.Service.String())
		return
	}

	server.services = append(server.services, service)

	return
}

// Start listening o
func (server *Server) launch_services() (err error) {
	for i := range server.home_config.File.HostServices {
		service_config := &server.home_config.File.HostServices[i]
		if err = server.create_home_service(service_config); err != nil {
			return
		}
		log.Println("launching", service_config.Service.String(), "service at address", service_config.Address)
	}

	for _, service := range server.services {
		if err = service.Start(); err != nil {
			err = fmt.Errorf("home: could not start service %s: %w", service.ID(), err)
			return
		}
	}

	return
}

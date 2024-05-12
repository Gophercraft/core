package home

import (
	"github.com/Gophercraft/core/app/config"
	bnet_rest_provider "github.com/Gophercraft/core/home/provider/bnet/rest"
	bnet_rest_service "github.com/Gophercraft/core/home/service/bnet_rest"
)

func (server *Server) create_bnet_rest_service(home_service_config *config.HomeServiceConfig) (service Service, err error) {
	// Setup service provider config
	var service_provider_config bnet_rest_provider.ProviderConfig
	service_provider_config.Endpoint = server.home_config.File.ServiceEndpoints[config.BNetRESTService]

	// Setup service config
	var service_config bnet_rest_service.ServiceConfig
	service_config.Address = home_service_config.Address

	provider := bnet_rest_provider.New(&service_provider_config, server.home_database)
	service = bnet_rest_service.New(&service_config, provider)
	return
}

package home

import (
	"github.com/Gophercraft/core/app/config"

	alphalist_provider "github.com/Gophercraft/core/home/provider/alphalist"
	alphalist_service "github.com/Gophercraft/core/home/service/alphalist"
)

func (server *Server) create_old_realmlist_service(service_config *config.HomeServiceConfig) (service Service, err error) {
	// Make service provider
	provider := alphalist_provider.New(&alphalist_provider.ServiceConfig{
		Build: service_config.Build,
	}, server.home_database)
	// Make service
	service = alphalist_service.New(service_config.Address, provider)
	return
}

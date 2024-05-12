package home

import (
	"github.com/Gophercraft/core/app/config"

	grunt "github.com/Gophercraft/core/grunt"
	grunt_provider "github.com/Gophercraft/core/home/provider/grunt"
)

func (server *Server) create_grunt_service(service_config *config.HomeServiceConfig) (service Service, err error) {
	var service_provider_config grunt_provider.ServiceProviderConfig
	service_provider_config.Mandatory2FA = server.home_config.File.TwoFactorAuthRequired

	provider := grunt_provider.New(&service_provider_config, server.home_database)
	service = grunt.NewServer(&grunt.ServerConfig{
		Address: service_config.Address,
	}, provider)
	return
}

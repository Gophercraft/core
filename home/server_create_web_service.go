package home

import (
	"github.com/Gophercraft/core/app/config"
	web_provider "github.com/Gophercraft/core/home/provider/web"
	web_service "github.com/Gophercraft/core/home/service/web"
	"github.com/Gophercraft/core/home/webapp"
)

func (server *Server) create_web_service(service_config *config.HomeServiceConfig) (service Service, err error) {
	// Create the web service provider (application logic)
	provider := web_provider.New(&web_provider.ProviderConfig{
		UseCaptchas:      true,
		EmailRequired:    false,
		WebRegistration:  server.home_config.File.OpenRegistration,
		ServiceAddresses: server.home_config.File.ServiceEndpoints,
	}, server.home_database)
	// Create a HTTP/JSON api based on the service provider
	// as well as a web app filesystem
	// TODO: allow the user to supply a custom directory to serve the webapp filesystem from
	// or disable entirely
	service = web_service.New(&web_service.ServiceConfig{
		Address: service_config.Address,
		WebApp:  webapp.FileServer(),
	}, provider)
	return
}

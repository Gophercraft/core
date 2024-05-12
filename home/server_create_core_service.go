package home

import (
	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/core/home/protocol/pb/account"
	"github.com/Gophercraft/core/home/protocol/pb/admin"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/core/home/protocol/pb/realm"

	core_account_provider "github.com/Gophercraft/core/home/provider/core/account"
	core_admin_provider "github.com/Gophercraft/core/home/provider/core/admin"
	core_auth_provider "github.com/Gophercraft/core/home/provider/core/auth"
	core_realm_provider "github.com/Gophercraft/core/home/provider/core/realm"
	core_service "github.com/Gophercraft/core/home/service/core"
)

func (server *Server) create_core_service(service_config *config.HomeServiceConfig) (service Service, err error) {
	// The core RPC service connects the Gophercraft home server
	// with other network members
	new_core_service := core_service.New(&core_service.ServiceConfig{
		Address: service_config.Address,
		TLS:     server.create_tls_config(),
	})

	// create GRPC core authentication service
	core_auth := core_auth_provider.New(server.home_database)
	auth.RegisterAuthServiceServer(new_core_service, core_auth)

	// create GRPC core realm registration service
	core_realm := core_realm_provider.New(server.home_database)
	realm.RegisterRealmServiceServer(new_core_service, core_realm)

	// create GRPC core realm administration service
	core_admin := core_admin_provider.New(server.home_database)
	admin.RegisterAdminServiceServer(new_core_service, core_admin)

	// create GRPC core account service
	core_account := core_account_provider.New(server.home_database)
	account.RegisterAccountServiceServer(new_core_service, core_account)

	service = new_core_service
	return
}

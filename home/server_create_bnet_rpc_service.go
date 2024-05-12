package home

import (
	"github.com/Gophercraft/core/app/config"
	account_v1 "github.com/Gophercraft/core/bnet/pb/bgs/protocol/account/v1"
	authentication_v1 "github.com/Gophercraft/core/bnet/pb/bgs/protocol/authentication/v1"
	connection_v1 "github.com/Gophercraft/core/bnet/pb/bgs/protocol/connection/v1"
	game_utilities_v1 "github.com/Gophercraft/core/bnet/pb/bgs/protocol/game_utilities/v1"
	bnet_rpc_account_provider "github.com/Gophercraft/core/home/provider/bnet/rpc/account"
	bnet_rpc_authentication_provider "github.com/Gophercraft/core/home/provider/bnet/rpc/authentication"
	bnet_rpc_connection_provider "github.com/Gophercraft/core/home/provider/bnet/rpc/connection"
	bnet_rpc_game_utilities_provider "github.com/Gophercraft/core/home/provider/bnet/rpc/game_utilities"
	bnet_rpc_service "github.com/Gophercraft/core/home/service/bnet_rpc"
)

func (server *Server) create_bnet_rpc_service(service_config *config.HomeServiceConfig) (service Service, err error) {
	// Make BattleNet service
	new_bnet_rpc_service := bnet_rpc_service.New(&bnet_rpc_service.ServiceConfig{
		Address: service_config.Address,
	})

	// create BNet connection service
	connection_service := bnet_rpc_connection_provider.New()
	connection_v1.RegisterConnectionServiceServer(new_bnet_rpc_service, connection_service)

	// create BNet authentication service
	authentication_service := bnet_rpc_authentication_provider.New(&bnet_rpc_authentication_provider.ServiceConfig{
		RESTLoginEndpoint: server.home_config.File.ServiceEndpoints[config.BNetRESTService],
	}, server.home_database)
	authentication_v1.RegisterAuthenticationServiceServer(new_bnet_rpc_service, authentication_service)

	// create BNet account service
	account_service := bnet_rpc_account_provider.New(server.home_database)
	account_v1.RegisterAccountServiceServer(new_bnet_rpc_service, account_service)

	// create BNet game utilities service
	game_utilities_service := bnet_rpc_game_utilities_provider.New(server.home_database)
	game_utilities_v1.RegisterGameUtilitiesServiceServer(new_bnet_rpc_service, game_utilities_service)

	service = new_bnet_rpc_service
	return
}

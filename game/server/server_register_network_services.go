package server

import (
	"github.com/Gophercraft/core/game/realm/service/auth"
)

func (server *Server) register_network_services() error {
	// Register auth service
	auth_service := auth.New(server.auth_provider)
	server.network_server.RegisterService(auth_service)

	// Register character service
	character_list_db_provider := character_db_provider.New(server.player_database)

	return nil
}

package server

import (
	"github.com/Gophercraft/core/game/network"
	"github.com/Gophercraft/core/game/server/database"
)

func (server *Server) make_network_config() (network_config *network.ServerConfig, err error) {
	network_config = new(network.ServerConfig)
	network_config.Protocol.Bind = server.config.BindGameAddress
	network_config.Protocol.Build = server.config.Build
	return
}

func (server *Server) make_database_config() (database_config *database.EngineConfig, err error) {
	database_config = new(database.EngineConfig)

	database_config.CacheDatabaseEngine = server.config.CacheDatabaseEngine
	database_config.CacheDatabaseLocation = server.config.CacheDatabasePath

	database_config.WorldDatabaseEngine = server.config.WorldDatabaseEngine
	database_config.WorldDatabaseLocation = server.config.WorldDatabasePath
	return
}

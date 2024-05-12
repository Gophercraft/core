package server

import (
	"fmt"

	"github.com/Gophercraft/core/app/config"
	"github.com/Gophercraft/core/datapack"
	"github.com/Gophercraft/core/game/game/service/auth"
	"github.com/Gophercraft/core/game/network"
	"github.com/Gophercraft/core/game/server/database"
)

// Server is the main implementation of a world server
type Server struct {
	config          *config.World
	database_engine *database.Engine
	network_server  *network.Server
	auth_provider   auth.ServiceProvider
	datapack_loader *datapack.Loader
	managers        subsystem_managers
	err             chan<- error
}

func Run(config *config.World) (err error) {
	server := new(Server)
	server.config = config

	err_channel := make(chan error)
	server.err = err_channel

	var (
		network_config  *network.ServerConfig
		database_config *database.EngineConfig
	)

	// Open database
	database_config, err = server.make_database_config()
	if err != nil {
		return
	}

	server.database_engine, err = database.Open(database_config)
	if err != nil {
		return
	}

	// Load datapacks
	if err = server.database_engine.CompileDatapacks(); err != nil {
		return
	}

	// Initialize scripting

	// Create managers for various subsystems
	if err = server.create_managers(); err != nil {
		return
	}

	// Create network server
	network_config, err = server.make_network_config()
	if err != nil {
		return
	}

	server.network_server, err = network.NewServer(network_config)
	if err != nil {
		err = fmt.Errorf("game/server: error initializing network server: %w", err)
		return
	}

	err = server.register_network_services()
	if err != nil {
		return
	}

	err = server.launch_network_server_listeners()
	if err != nil {
		return
	}

	err = <-err_channel
	return
}

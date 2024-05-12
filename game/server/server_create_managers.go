package server

import "github.com/Gophercraft/core/game/server/manager/core"

type subsystem_managers struct {
	core_manager *core.Manager
}

func (server *Server) create_managers() (err error) {
	m := &server.managers

	m.core_manager = core.New(&core.ManagerConfig{
		HomeServerAddress:     server.config.HomeServerAddress,
		HomeServerFingerprint: server.config.HomeServerAddress,
	})
}

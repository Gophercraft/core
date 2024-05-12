package home

func (server *Server) validate_config() (err error) {
	if server.home_config.File.DatabaseEngine == "" {
		server.home_config.File.DatabaseEngine = "leveldb_core"
	}

	if server.home_config.File.DatabasePath == "" {
		server.home_config.File.DatabasePath = "home_db"
	}

	return
}

package grunt

type ServerConfig struct {
	Address string
}

func NewServer(server_config *ServerConfig, service_provider ServiceProvider) (server *Server) {
	server = new(Server)
	server.config = server_config
	server.service_provider = service_provider
	return
}

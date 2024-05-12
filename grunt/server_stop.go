package grunt

func (server *Server) Stop() (err error) {
	return server.listener.Close()
}

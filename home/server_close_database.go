package home

func (server *Server) close_databases() (err error) {
	err = server.home_database.Close()
	return
}

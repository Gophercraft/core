package home

func (server *Server) stop_services() (err error) {
	for i := range server.services {
		service := server.services[i]
		if service != nil {
			err = service.Stop()
			if err != nil {
				return
			}
			server.services[i] = nil
		}
	}

	return
}

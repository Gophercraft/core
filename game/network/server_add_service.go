package network

// Register a service object with the network. May not Init the service until launch, however
func (server *Server) RegisterService(service Service) error {
	server.services[service.ID()] = service
	return nil
}

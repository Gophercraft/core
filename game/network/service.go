package network

// Service describes a certain functionality that the user wishes to expose on the server.
type Service interface {
	// Get the ID
	ID() ServiceID
	Descriptor() *ServiceDescriptor
	Init(server *Server) error
}

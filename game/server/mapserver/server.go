package mapserver

import "github.com/Gophercraft/core/tempest"

type UpdateMessage struct {
}

type ClientBridge interface {
	// Changes the location in the map bridge
	// Where the client receives updates
	UpdateObserverPosition(position tempest.C3Vector)
	// Returns
	ReadUpdateMessage() (message *UpdateMessage, err error)

	Disconnect()
}

type ConnectionRequest struct {
	MapID            uint32
	ObserverPosition tempest.C3Vector
}

// Server describes a mapserver.

type Server interface {
	//
	ConnectBridge(request *ConnectionRequest) (client_bridge ClientBridge, err error)
}

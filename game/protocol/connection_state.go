package protocol

type ConnectionState uint8

const (
	//
	Unknown = iota

	//
	New

	//
	Handshaking

	//
	Verified

	//
	Disconnected

	EndConnectionStates
)

const NumConnectionStates = int(EndConnectionStates)

type NewConnectionStateHandlerFunc func(st ConnectionState, conn *Connection, err error)

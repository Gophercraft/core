package sync

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet/update"
)

type Object interface {
	GUID() guid.GUID
	TypeID() guid.TypeID
	// Values() must return a pointer to a values block that can be modified by the server
	Values() *update.ValuesBlock
}

// Objects that have a presence in the world in a specific location (players, creatures)
type WorldObject interface {
	Object
	// movement data
	Movement() *update.MovementBlock
}

type Vision tempest.C4Sphere

// Used to interface with players
type Seer interface {
	WorldObject

	Vision() *Vision
}

package models

import (
	"time"

	"github.com/Gophercraft/core/tempest"
)

type ObjectPathStep struct {
	Position tempest.C4Vector
}

type ObjectPath struct {
	Steps []ObjectPathStep
}

type PossibleSpawnPosition struct {
	Weight   int
	Position tempest.C4Vector
}

type SpawnObject struct {
	Kind        SpawnKind
	RelPosition tempest.C4Vector
	Rotation    tempest.C4Quaternion
	Chance      float32
	ID          string
}

type SpawnGroup struct {
	ID      string
	Phase   []string
	Map     uint32
	Respawn time.Duration // if zero, treat as default

	Chance float32 // 0 - 100.0

	Positions []PossibleSpawnPosition

	Objects []SpawnObject

	Path ObjectPath

	// TODO: Quartz expressions for encoding seasonal NPCs, NPCs that only come out at night etc.
	Start, End string
}

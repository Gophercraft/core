package models

import "github.com/Gophercraft/core/tempest"

type PortLocation struct {
	ID       string `xorm:"'port_id' pk" csv:"name"`
	Location tempest.C4Vector
	Map      uint32 `xorm:"'map'" csv:"mapID"`
}

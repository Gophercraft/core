package models

import (
	"github.com/Gophercraft/core/guid"
)

// ObjectTemplateRegistry contains the resolved IDs of a custom object template (think custom items and creatures)
// It allows you to add and remove new content and (hopefully) avoid conflicts between the client's WDB cache and the server's in-memory database.
type ObjectTemplateRegistry struct {
	ID    string      `xorm:"'id' pk"`
	Type  guid.TypeID `xorm:"'type'"`
	Entry uint32
}

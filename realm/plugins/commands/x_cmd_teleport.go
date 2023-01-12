package commands

import (
	"github.com/Gophercraft/core/realm"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/tempest"
	"github.com/davecgh/go-spew/spew"
)

func cmdTele(s *realm.Session, portID string) {
	// port string
	pos := tempest.C4Vector{}

	var mapID uint32

	var port *models.PortLocation
	s.DB().Lookup(wdb.BucketKeyStringID, portID, &port)

	if port == nil {
		// var ports []models.PortLocation
		// if err := s.DB().QueryField(s.Locale, "ID", ".*", 1, &ports); err != nil {
		// 	s.Warnf("%s", err)
		// 	return
		// }
		// if len(ports) == 0 {
		// 	s.Warnf("could not find port location: '%s'", portID)
		// 	return
		// }
		// port = &ports[0]
		// s.Warnf("Could not find teleport location %s, sending you to %s.", portID, port.ID)
		s.Warnf("Port location %s not found.", portID)
		return
	}

	mapID = port.Map
	pos = port.Location
	s.Warnf("%s", spew.Sdump(port))
	// }
	// } else {
	// 	pos.X = c.Float32(0)
	// 	pos.Y = c.Float32(1)
	// 	pos.Z = c.Float32(2)
	// 	pos.O = c.Float32(3)
	// 	mapID = c.Uint32(4)
	// }

	s.TeleportTo(mapID, pos)
}

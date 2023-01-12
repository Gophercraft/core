package commands

import (
	"github.com/Gophercraft/core/realm"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
)

func cmdAddNPC(s *realm.Session, npcID string) {
	var cr *models.CreatureTemplate
	s.DB().Lookup(wdb.BucketKeyStringID, npcID, &cr)
	if cr == nil {
		s.Warnf("No CreatureTemplate could be found with the ID %s", npcID)
		return
	}

	creature := s.Server.NewCreature(cr, s.Position())
	s.Map().AddObject(creature)

	s.Warnf("Object spawned successfully: %s", creature.GUID())
}

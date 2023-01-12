package commands

import (
	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/format/terrain"
	"github.com/Gophercraft/core/realm"
	"github.com/Gophercraft/core/realm/wdb"
)

func cmdGPS(s *realm.Session) {
	pos := s.Position()

	var zoneName, areaName string
	var areaEntry *dbdefs.Ent_AreaTable
	// Query main zone
	s.DB().Lookup(wdb.BucketKeyUint32ID, s.ZoneID, &areaEntry)
	if areaEntry != nil {
		zoneName = areaEntry.AreaName_lang.GetLocalized(s.Locale)
	}
	areaEntry = nil
	// Query specific zone/subzone
	s.DB().Lookup(wdb.BucketKeyUint32ID, s.CurrentArea, &areaEntry)
	if areaEntry != nil {
		areaName = areaEntry.AreaName_lang.GetLocalized(s.Locale)
	}

	mp := s.Map()

	s.Kv(" X", "%f", pos.X)
	s.Kv(" Y", "%f", pos.Y)
	s.Kv(" Z", "%f", pos.Z)
	s.Kv(" Facing", "%f", pos.W)
	s.Kv("Map", "%d '%s'", mp.ID, mp.Definition().MapName_lang.GetLocalized(s.Locale))
	s.Kv("Zone", "%d '%s'", s.ZoneID, zoneName)
	s.Kv("Area", "%d '%s'", s.CurrentArea, areaName)
	s.Kv("Phase", "%s", s.CurrentPhase)

	blockIndex, err := terrain.CalcBlockIndex(&terrain.DefaultMap, pos.C2())
	if err == nil {
		s.Kv("Block", "%d:%d", blockIndex.X, blockIndex.Y)

		chunkIndex, err := terrain.CalcChunkIndex(&terrain.DefaultMap, pos.C2())
		if err == nil {
			s.Kv("Chunk", "%d:%d", chunkIndex.X, chunkIndex.Y)
		}

	} else {
		s.Warnf("Block calc fail (wtf) %s", err)
	}
}

package realm

import (
	"fmt"
	"time"

	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/packet/xp"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"
)

func (ws *Server) GetNextLevelXP(forLevel uint32) uint32 {
	xp, ok := ws.LevelExperience[forLevel]
	if !ok {
		return 0
	}

	return xp
}

func (s *Session) GetNextLevelXP() uint32 {
	lvl := uint32(s.GetLevel())
	return s.Server.GetNextLevelXP(lvl)
}

func (s *Session) LevelUp(to uint32) {
	// Todo: reset health and other attributes

	s.SetUint32("XP", 0)
	s.SetUint32("NextLevelXP", s.GetNextLevelXP())
	s.Char.XP = 0
	s.Char.Level = to
	s.Char.LevelPlayedTime = 0
	s.LevelJoinTime = time.Now()
	s.SetUint32("Level", s.Char.Level)
	s.DB().Where("id = ?", s.PlayerID()).Cols("xp", "level", "playtime_level").Update(s.Char)

	s.UpdatePlayer()
}

func (s *Session) SetCurrentXP(current uint32) {
	s.SetUint32("XP", current)
	s.Char.XP = current
	s.DB().Where("id = ?", s.PlayerID()).Cols("xp").Update(s.Char)
}

func (s *Session) GetMaxLevel() int {
	// return 70
	return 255
}

func (s *Session) AddExperience(newXP uint32) {
	curXP := s.Get("XP").Uint32()
	curLevel := s.GetLevel()
	var nxtLevelXP = s.Server.GetNextLevelXP(uint32(curLevel))

	for newXP > 0 {
		// Don't add XP past level ceiling
		if curLevel == s.GetMaxLevel() {
			break
		}

		nxtLevelXP = s.Server.GetNextLevelXP(uint32(curLevel))
		if curXP+newXP >= nxtLevelXP {
			newXP -= nxtLevelXP - curXP
			if int32(newXP) < 0 {
				panic("WTF")
			}
			curXP = 0
			curLevel++
		} else {
			curXP += newXP
			break
		}

	}

	s.SetUint32("XP", curXP)
	if curLevel != s.GetLevel() {
		s.LevelUp(uint32(curLevel))
	} else {
		s.UpdateSelf()
	}
}

func (ws *Server) ExploreXPRate() float32 {
	return ws.FloatVar("XP.Rate")
}

func (ws *Server) GetExploreXP(explorationLevel uint32) uint32 {
	return 35
}

func (s *Session) ZoneExplored(zoneID uint32) bool {
	ct, err := s.DB().Where("player = ?", s.PlayerID()).Where("zone_id = ?", zoneID).Count(new(models.ExploredZone))
	if err != nil {
		panic(err)
	}

	return ct > 0
}

func (s *Session) SetExplorationFlag(exploreFlag uint32) {
	blockOffset := int(exploreFlag / 32)
	bitOffset := int(exploreFlag % 32)

	sli := s.Get("ExploredZones")
	if blockOffset >= sli.Len() {
		return
	}

	mask := sli.Index(blockOffset).Uint32()

	mask |= (1 << bitOffset)

	sli.Index(blockOffset).SetUint32(mask)
}

func (s *Session) HandleZoneExperience(zoneID uint32) {
	if s.ZoneExplored(zoneID) {
		fmt.Println("Zone", zoneID, "explored already")
		return
	}

	var areaTable *dbdefs.Ent_AreaTable
	s.DB().Lookup(wdb.BucketKeyUint32ID, zoneID, &areaTable)

	log.Dump("areaTable", areaTable)

	if areaTable == nil {
		return
	}

	if s.Build().AddedIn(vsn.V1_12_1) {
		if areaTable.AreaBit != 0 {
			// Reveal location in map
			s.SetExplorationFlag(uint32(areaTable.AreaBit))
			s.UpdateSelf()
			s.DB().Insert(models.ExploredZone{
				Player: s.PlayerID(),
				ZoneID: zoneID,
			})
		}
	}

	if areaTable.ExplorationLevel != 0 {
		if s.GetLevel() >= s.GetMaxLevel() {
			s.Warnf("Exploring zone %d", zoneID)
			s.Send(&xp.Exploration{
				ZoneID: zoneID,
			})
		} else {
			var exp uint32
			diff := s.GetLevel() - int(areaTable.ExplorationLevel)
			if diff < -5 {
				exp = uint32(float32(s.Server.GetExploreXP(uint32(s.GetLevel()+5))) * s.Server.ExploreXPRate())
			} else if diff > 5 {
				explorationPercent := (100 - ((diff - 5) * 5))
				if explorationPercent > 100 {
					explorationPercent = 100
				}

				if explorationPercent < 0 {
					explorationPercent = 0
				}

				exp = uint32(
					float32(s.Server.GetExploreXP(uint32(areaTable.ExplorationLevel))) * (float32(explorationPercent) / 100.0) * s.Server.ExploreXPRate())
			} else {
				exp = uint32(float32(s.Server.GetExploreXP(uint32(areaTable.ExplorationLevel))) * s.Server.ExploreXPRate())
			}

			s.Send(&xp.Exploration{
				ZoneID:     zoneID,
				Experience: exp,
			})
			s.AddExperience(exp)
		}
	} else {
		if areaTable.AreaBit != 0 {
			s.Send(&xp.Exploration{
				ZoneID:     zoneID,
				Experience: 0,
			})
		}
	}
}

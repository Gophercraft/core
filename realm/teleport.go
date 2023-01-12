package realm

import (
	"math"
	"time"

	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/packet/area"
	"github.com/Gophercraft/core/packet/teleport"
	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"
	"github.com/davecgh/go-spew/spew"
)

func (s *Session) BindpointUpdate() {
	//goldshire
	bind := &teleport.BindpointUpdate{
		Position: tempest.C3Vector{
			X: -8949.95,
			Y: -132.493,
			Z: 83.5312,
		},

		MapID: 0,

		ZoneID: 12,
	}

	s.Send(bind)
}

func (s *Session) SendNewWorld(mapID uint32, pos tempest.C4Vector) {
	s.Send(&teleport.NewWorld{
		MapID:    mapID,
		Position: pos,
	})
}

func (s *Session) SendTeleportAck(g guid.GUID, mapID uint32, pos tempest.C4Vector) {
	if s.Build().RemovedIn(vsn.V1_12_1) {
		pkt := packet.NewWorldPacket(packet.SMSG_MOVE_WORLDPORT_ACK)
		mi := &update.MovementInfo{
			TransportPosition: pos,
			Position:          pos,
		}
		update.EncodeMovementInfo(s.Build(), pkt, mi)
		s.SendPacket(pkt)
		return
	}

	s.Send(&teleport.Ack{
		GUID: s.GUID(),
		Info: &update.MovementInfo{
			Flags:    0,
			Position: pos,
		},
	})
}

func (s *Session) SendTransferPending(mapID uint32) {
	s.Send(&teleport.TransferPending{
		MapID: mapID,
	})
}

// For some reason this is necessary
func (s *Session) HandleWorldportAck() {
	// Resend inventory
	for _, v := range s.Items {
		s.SendObjectCreate(v)
	}

	// Tell the client they've successfully teleported.
	// s.SendObjectCreate(s)
	s.Notify(UnitNotifyCreate, s)

	// Notify our client of existing objects in this map.
	for _, obj := range s.Map().VisibleObjects(s) {
		s.Notify(UnitNotifyCreate, obj)
		// s.SendObjectCreate(obj)
	}

	s.CurrentArea = 0

	s.SyncTime()
	// s.UpdateArea()
}

// TeleportTo teleports a player to a new location. This function should be called carefully.
func (s *Session) TeleportTo(mapID uint32, newPos tempest.C4Vector) {
	log.Println("Teleporting", s.PlayerName(), "to", mapID, spew.Sdump(newPos))

	if s.CurrentMap == nil {
		return
	}

	if mapID == s.CurrentMap.ID {
		s.SendTeleportAck(s.GUID(), mapID, newPos)

		// for _, sess := range s.Map().VisibleObjects(s).Sessions() {
		// 	// if our new location is too far from this player
		// 	if newPos.Distance(sess.Position()) > s.Map().VisibilityDistance() {
		// 		// make player s disappear in their client
		// 		sess.SendObjectDelete(s.GUID())
		// 	} else { // or else
		// 		// make player s jump to new position in nearby player's client
		// 		sess.SendTeleportAck(s.GUID(), mapID, newPos)
		// 	}
		// }

		// s.MovementInfo.Position = newPos

		s.SetPosition(newPos)

		// set position
		s.UpdatePosition()

		// // Now that our new position is set, notify nearby players that we are here.
		// for _, sess := range s.Map().VisibleObjects(s).Sessions() {
		// 	sess.SendObjectCreate(s)
		// }

		// s.CurrentChunkIndex = nil
		// s.CurrentArea = 0

		s.SyncTime()
		// s.UpdateArea()
	} else {
		// Open appropriate loading screen for mapID
		s.SendTransferPending(mapID)

		// Remove from old map and notify other clients of this player's disappearance
		s.Map().RemoveObject(s.GUID())

		// tell player they are deleted
		// s.SendObjectDelete(s.GUID())

		s.SetPosition(newPos)

		// s.MovementInfo.Position = newPos
		s.CurrentMap = s.Phase().Map(mapID)
		// s.UpdatePosition()

		// trigger worldport ack once client is finished loading into zone
		s.SendNewWorld(mapID, newPos)

		s.Map().AddObject(s)
	}
}

// important todo: verify zone ID against geometric bounds of area:
// Just a dummy function until ADT parsing is available
func (s *Session) ValidateZoneWithPosition(zoneID uint32) bool {
	return true
}

func (s *Session) HandleZoneUpdate(zu *area.ZoneUpdate) {
	zoneID := zu.ID

	if !s.ValidateZoneWithPosition(zoneID) {
		return
	}

	_, err := s.DB().Cols("zone").Where("id = ?", s.PlayerID()).Update(&models.Character{
		Zone: zoneID,
	})

	if err != nil {
		panic(err)
	}

	s.ZoneID = zoneID
}

func (s *Session) HandleAreaTrigger(trigger *area.Trigger) {
	var aTrigger *dbdefs.Ent_AreaTrigger
	s.DB().Lookup(wdb.BucketKeyUint32ID, trigger.ID, &aTrigger)
	if aTrigger == nil {
		log.Warn("Player tried to call non existent area trigger", trigger.ID)
		return
	}
	// delta is safe radius
	const delta float32 = 5.0

	pos := s.Position()

	if !isPointInAreaTriggerZone(aTrigger, s.CurrentMap.ID, pos.X, pos.Y, pos.Z, delta) {
		log.Warn("Player", s.PlayerName(), "tried to teleport to area trigger", trigger.ID, "without being in correct position")
		return
	}

	log.Println("Area trigger", trigger.ID)

	s.Server.Call(AreaTriggerEvent, trigger.ID, s)
}

func (s *Session) SendRequiredLevelZoneError(lvl int) {
	s.Alertf("You must be at least level %d to enter this zone.", lvl)
}

func (s *Session) SendRequiredItemZoneError(itemID string) {
	tpl := s.GetItemTemplate(models.Item{
		ItemID: itemID,
	})
	s.Alertf("You must have %s to enter this zone.", tpl.Name)
}

func (s *Session) SendRequiredQuestZoneError(questEntry uint32) {
	// todo: send quest name
	s.Alertf("You must solve a quest to enter this zone.")
}

func fabs(i float32) float32 {
	return float32(math.Abs(float64(i)))
}

func isPointInAreaTriggerZone(atEntry *dbdefs.Ent_AreaTrigger, mapID uint32, x, y, z, delta float32) bool {
	if mapID != uint32(atEntry.ContinentID) {
		log.Warn("incorrect mapID", mapID)
		return false
	}

	if atEntry.Radius > 0 {
		// if we have radius check it
		dist2 := (x-atEntry.Pos[0])*(x-atEntry.Pos[0]) + (y-atEntry.Pos[1])*(y-atEntry.Pos[1]) + (z-atEntry.Pos[2])*(z-atEntry.Pos[2])
		if dist2 > (atEntry.Radius+delta)*(atEntry.Radius+delta) {
			return false
		}
	} else {
		// we have only extent

		// rotate the players position instead of rotating the whole cube, that way we can make a simplified
		// is-in-cube check and we have to calculate only one point instead of 4

		// 2PI = 360, keep in mind that ingame orientation is counter-clockwise
		rotation := float64(2)*math.Pi - float64(atEntry.Box_yaw)
		sinVal := float32(math.Sin(rotation))
		cosVal := float32(math.Cos(rotation))

		playerBoxDistX := x - atEntry.Pos[0]
		playerBoxDistY := y - atEntry.Pos[1]

		rotPlayerX := float32(atEntry.Pos[0] + playerBoxDistX*cosVal - playerBoxDistY*sinVal)
		rotPlayerY := float32(atEntry.Pos[1] + playerBoxDistY*cosVal + playerBoxDistX*sinVal)

		// box edges are parallel to coordiante axis, so we can treat every dimension independently :D
		dz := z - atEntry.Pos[2]
		dx := rotPlayerX - atEntry.Pos[0]
		dy := rotPlayerY - atEntry.Pos[1]
		if (fabs(dx) > atEntry.Box_width/2+delta) ||
			(fabs(dy) > atEntry.Box_length/2+delta) ||
			(fabs(dz) > atEntry.Box_yaw/2+delta) {
			return false
		}
	}

	return true
}

func (c *Session) PhaseTeleportTo(phaseID string, mapID uint32, pos tempest.C4Vector) {
	if phaseID == c.CurrentPhase {
		c.TeleportTo(mapID, pos)
	} else {
		panic("nyi")
	}
}

func (c *Session) HandleSummonResponse() {
	if !c.IsAlive() {
		return
	}

	if c.TeleportSummons == nil {
		return
	}

	c.PhaseTeleportTo(c.TeleportSummons.Phase, c.TeleportSummons.Map, c.TeleportSummons.Pos)
	c.TeleportSummons = nil
}

type Summons struct {
	Phase string
	Map   uint32
	Pos   tempest.C4Vector
}

func (s *Session) SetSummonLocation(phase string, mapID uint32, pos tempest.C4Vector) {
	s.TeleportSummons = &Summons{
		phase,
		mapID,
		pos,
	}
}

func (s *Session) SendSummonRequest(summoner guid.GUID, zoneID uint32, timeout time.Duration) {
	s.Send(&teleport.SummonRequest{
		ID:      summoner,
		Zone:    zoneID,
		Timeout: timeout,
	})
}

func (s *Session) HandleWorldTeleport(wt *teleport.Worldport) {
	if !s.IsGM() {
		return
	}
	s.TeleportTo(wt.MapID, wt.Pos)
}

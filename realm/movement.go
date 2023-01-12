package realm

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"
)

func (s *Session) SetFly(on bool) {
	s.Flying = on
	if s.Build().RemovedIn(vsn.V2_4_3) {
		// Hacky bullshit.
		// Flying wasn't actually implemented until version 2.0.
		// You can only move laterally and it's buggy as hell
		mask := update.MoveFlagCanFly | update.MoveFlagFlying

		if on {
			s.MovementInfo.Flags |= mask
		} else {
			s.MovementInfo.Flags &= ^mask
		}

		s.Send(&update.MovementPacket{
			Type:   packet.MSG_MOVE_START_SWIM,
			Server: true,
			GUID:   s.GUID(),
			Info:   s.MovementInfo,
		})
	} else {
		s.Send(&update.FlyState{
			GUID: s.GUID(),
			On:   on,
		})
	}
}

// Important TODO: validate position
func (s *Session) HandleUpdateMovement(mpacket *update.MovementPacket) {
	if mpacket.GUID == guid.Nil {
		mpacket.GUID = s.GUID()
	}

	if s.Moves != nil {
		s.Moves <- mpacket
	}

	// log.Dump("movement update", mpacket)

	// TODO: validate flags
	// s.UpdatePosition(s.MovementInfo.Position)

	// s.Map().PlayersInSightOfObject(s, s.Map().VisibilityDistance()).Iter(func(s *Session) {
	// 	s.Send(&update.MovementPacket{
	// 		Type:   mpacket.Type,
	// 		Server: true,
	// 		// Todo: allow player to control other GUIDs
	// 		GUID: s.GUID(),
	// 		Info: mpacket.Info,
	// 	})
	// })
}

// func (s *Session) UpdateCameraPosition(syncSelf bool, pos tempest.C4Vector) {

// 	visRange := s.Server.FloatVar("Sync.VisibilityRange")

// 	s.GuardTrackedGUIDs.Lock()
// 	defer s.GuardTrackedGUIDs.Unlock()

// 	tGuids := s.TrackedGUIDs[:0]

// 	// TODO: time out object synchronization if the player only moved a little bit, or has moved very recently

// 	// Remove GUIDs if too far away
// 	for _, g := range s.TrackedGUIDs {
// 		active := s.Map().GetObject(g)
// 		if active == nil {
// 			s.SendObjectDelete(g)
// 		} else {
// 			if active.Movement().Position.C3().Distance(s.Position().C3()) > visRange {
// 				s.SendObjectDelete(g)

// 				if syncSelf {
// 					if activeSession, ok := (active).(*Session); ok {
// 						activeSession.RemoveTrackedGUID(s.GUID())
// 						activeSession.SendObjectDelete(s.GUID())
// 					}
// 				}
// 			} else {
// 				tGuids = append(tGuids, g)
// 			}
// 		}
// 	}

// 	s.TrackedGUIDs = tGuids

// 	// Add new GUIDs if not found yet
// 	for _, nearObject := range s.Map().VisibleObjectsInRange(s, visRange) {
// 		if !s.IsTrackedGUID(nearObject.GUID()) {
// 			s.TrackedGUIDs = append(s.TrackedGUIDs, nearObject.GUID())
// 			s.SendObjectCreate(nearObject)

// 			if syncSelf {
// 				// TODO: it may be more efficient to send these creates as multiple blocks in a SMSG_COMPRESSED_OBJECT_UPDATE
// 				// We may be appearing to a new player. Notify them of us.
// 				if nearSession, ok := (nearObject).(*Session); ok {
// 					nearSession.SendObjectCreate(s)
// 				}
// 			}
// 		}
// 	}
// }

func (s *Session) HandleMoves(moves *update.MovementPacket) {
	s.HandleUpdateMovement(moves)
}

func (s *Session) Movement() *update.MovementBlock {
	s.MovementInfo.Time = s.Server.UptimeMS()

	mData := &update.MovementBlock{
		Speeds:   s.MoveSpeeds,
		Position: s.MovementInfo.Position,
		Info:     s.MovementInfo,
	}

	mData.UpdateFlags |= update.UpdateFlagLiving
	mData.UpdateFlags |= update.UpdateFlagHasPosition

	mData.UpdateFlags |= update.UpdateFlagAll
	mData.UpdateFlags |= update.UpdateFlagHighGUID
	mData.All = 0x1 // 5875 only
	mData.HighGUID = 0x1
	return mData
}

func (s *Session) Position() tempest.C4Vector {
	return s.MovementInfo.Position
}

func (s *Session) Speeds() update.Speeds {
	return s.MoveSpeeds
}

func (m *Map) UpdateSpeed(g guid.GUID, st update.SpeedType) {
	obj := m.GetObject(g)

	if obj == nil {
		return
	}

	speeds := obj.Movement().Speeds

	if g.HighType() == guid.Player {
		if sess, ok := obj.(*Session); ok {
			sess.Send(&update.ForceSpeedChange{
				Type:     st,
				ID:       g,
				NewSpeed: speeds[st],
			})
		}
	}

	// pkt := packet.NewWorldPacket(pair.Spline)
	// g.EncodePacked(m.Phase.Server.Build(), pkt)
	// pkt.WriteFloat32(speeds[st])

	m.VisibleObjects(obj).Sessions().Iter(func(s *Session) {
		s.Send(&update.SplineSpeedChange{
			Type:     st,
			ID:       g,
			NewSpeed: speeds[st],
		})
	})
}

func (s *Session) SyncSpeeds() {

	list := update.GetSpeedList(s.Build())
	for _, speed := range list {
		s.Map().UpdateSpeed(s.GUID(), speed)
	}

	// var sl update.Speed
	// vsn.QueryDescriptors(s.Build())

	// for k, v := range update.SpeedLists {
	// 	if k.Contains(s.Build()) {
	// 		for _, speed := range v {
	// 			s.Map().UpdateSpeed(s.GUID(), speed)
	// 		}
	// 		break
	// 	}
	// }
	s.SyncTime()
}

func (s *Session) HandleSetActiveMover(sam *update.SetActiveMover) {
	log.Println(s, "set active mover", sam.ID)
}

func (s *Session) SetPosition(c4 tempest.C4Vector) {
	s.Camera.Position = c4
	s.MovementInfo.Position = c4
	s.Char.X = s.MovementInfo.Position.X
	s.Char.Y = s.MovementInfo.Position.Y
	s.Char.Z = s.MovementInfo.Position.Z
	s.Char.O = s.MovementInfo.Position.W
}

func (s *Session) UpdatePosition() {
	s.Map().UpdateMapPosition(s.GUID())
}

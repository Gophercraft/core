package realm

import (
	"math"

	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet/chat"
	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/tempest"
)

func (s *Session) IsAlive() bool {
	return true
}

func (s *Session) CanSpeak() bool {
	return true
}

func (s *Session) HandleStandStateChange(sss *update.SetStandState) {
	// Validation
	switch sss.State {
	case update.StateStand,
		update.StateKneel,
		update.StateSit,
		update.StateSleep:
		break
	default:
		return
	}

	// Broadcast new stand state to server
	s.SetStandState(sss.State)
}

func (s *Session) SetStandState(value update.StandState) {
	s.SetByte("StandState", uint8(value))
	s.UpdatePlayer()
}

func (s *Session) HandleTextEmote(textemote *chat.TextEmoteRequest) {
	if !s.IsAlive() {
		return
	}

	if !s.CanSpeak() {
		return
	}

	target := s.GetTarget()

	var emote *dbdefs.Ent_EmotesText
	s.DB().Lookup(wdb.BucketKeyUint32ID, textemote.Text, &emote)

	if emote == nil {
		s.Warnf("You appear to have sent an invalid emote command. Check to see if you have a base datapack installed.")
		return
	}

	switch emote.EmoteID {
	case 12, 13, 16, 0: //sleep, sit, kneel, none
	default:
		s.HandleEmoteCommand(uint32(emote.EmoteID))
	}

	var data string
	var err error
	if target != guid.Nil {
		data, err = s.Server.GetUnitNameByGUID(target)
		if err != nil {
			s.Warnf("%s guid=%s", err.Error(), target.Summary())
			return
		}
	}

	textEmote := &chat.TextEmote{
		GUID:  s.GUID(),
		Text:  textemote.Text,
		Emote: textemote.Emote,
		Name:  data,
	}

	s.Map().ViewersOf(s).Sessions().Iter(func(xs *Session) {
		xs.Send(textEmote)
	})
}

func (s *Session) HandleEmoteCommand(emoteID uint32) {
	s.Map().ViewersOf(s).Sessions().Iter(func(xs *Session) {
		xs.Send(&chat.Emote{
			Emote: emoteID,
			GUID:  s.GUID(),
		})
	})
}

func (s *Session) SitChair(chair *GameObject) {
	chairPos := chair.Position
	gobjt := s.GetGameObjectTemplateByEntry(chair.Entry())

	slots := gobjt.Data[0]
	height := update.StandState(gobjt.Data[1])

	if slots > 0 {
		lowestDist := s.Map().VisibilityDistance()

		xLowest := chairPos.X
		yLowest := chairPos.Y

		orthogOrientation := chairPos.W + float32(math.Pi)*0.5

		pos := s.Position()

		for i := uint32(0); i < slots; i++ {
			relDistance := (gobjt.Size*float32(i) - float32(gobjt.Size)*float32(slots-1)/2.0)

			xI := chairPos.X + relDistance*float32(math.Cos(float64(orthogOrientation)))
			yI := chairPos.X + relDistance*float32(math.Sin(float64(orthogOrientation)))

			thisDistance := pos.C3().Distance(tempest.C3Vector{
				X: xI,
				Y: yI,
			})

			if thisDistance < lowestDist {
				lowestDist = thisDistance
				xLowest = xI
				yLowest = yI
			}
		}

		s.TeleportTo(s.CurrentMap.ID, tempest.C4Vector{
			X: xLowest,
			Y: yLowest,
			Z: chairPos.Z,
			W: chairPos.W,
		})
	} else {
		s.TeleportTo(s.CurrentMap.ID, chairPos)
	}

	s.SetStandState(update.StateSitLowChair + height)
}

func (s *Session) HandleSheathe(sheathe *update.SetSheathe) {
	s.ValuesBlock.SetByte("SheathState", uint8(sheathe.Mode))
	s.UpdatePlayer()
}

func (s *Session) HandleSetWeaponMode(swm *update.SetWeaponMode) {
	if swm.Mode == 1 {
		s.SetBit("Sheathe", true)
	} else {
		s.SetBit("Sheathe", false)
	}
	s.UpdatePlayer()
}

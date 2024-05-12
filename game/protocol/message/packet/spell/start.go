package spell

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

type Start struct {
	Data

	Source guid.GUID

	ModifiedPowers []ModifiedPower
	RuneCooldown   uint32

	SpellRuneState  uint8
	PlayerRuneState uint8

	RuneCooldownPasses []byte

	Elevation float32
	DelayTime int32

	VisualChain []int32

	DestinationLocationCount uint8
}

func (s *Start) Encode(build version.Build, out *message.Packet) (err error) {
	out.Type = message.SMSG_SPELL_START

	s.Caster.EncodePacked(build, out)
	s.Source.EncodePacked(build, out)

	if build.AddedIn(version.V3_0_2) {
		out.WriteUint32(s.CastCount)
	}

	out.WriteUint32(s.Spell)

	if build.AddedIn(version.V3_0_2) {
		out.WriteUint32(uint32(s.Flags))
	} else {
		out.WriteUint16(uint16(s.Flags))
	}

	if build.AddedIn(15005) {
		out.WriteUint64(s.CastTime)
	} else {
		out.WriteUint32(uint32(s.CastTime))
	}

	out.WriteUint32(uint32(len(s.HitTargets)))

	for _, target := range s.HitTargets {
		target.EncodeUnpacked(build, out)
	}

	for _, miss := range s.MissedTargets {
		if err = miss.Encode(build, out); err != nil {
			return err
		}
	}

	if err = s.Data.EncodeCastTargets(build, out); err != nil {
		return err
	}

	if build.AddedIn(9056) {
		if s.Flags&PredictedPower != 0 {
			if build.AddedIn(16309) {
				out.WriteUint32(uint32(len(s.ModifiedPowers)))
				for _, power := range s.ModifiedPowers {
					if err = power.Encode(build, out); err != nil {
						return
					}
				}
			} else {
				out.WriteUint32(s.RuneCooldown)
			}
		}

		if s.Flags&RuneInfo != 0 {
			out.WriteUint8(s.SpellRuneState)
			out.WriteUint8(s.PlayerRuneState)

			runeCooldownPasses := make([]byte, 6)
			copy(runeCooldownPasses, s.RuneCooldownPasses)

			for b := uint8(0); b < uint8(len(s.RuneCooldownPasses)); b++ {
				if build.RemovedIn(4545) {
					var mask uint8 = 1 << b
					if (mask & s.SpellRuneState) == 0 {
						continue
					}

					if (mask & s.PlayerRuneState) != 0 {
						continue
					}
				}

				out.WriteUint8(s.RuneCooldownPasses[int(b)])
			}
		}

		if s.Flags&AdjustMissile != 0 {
			out.WriteFloat32(s.Elevation)
			out.WriteInt32(s.DelayTime)
		}

		if s.Flags&Projectile != 0 {
			out.WriteInt32(s.AmmoDisplayID)
			out.WriteUint32(s.AmmoInventoryType)
		}

		if build.AddedIn(9056) {
			if s.Flags&VisualChain != 0 {
				visualChain := make([]int32, 2)
				copy(visualChain, s.VisualChain)
				out.WriteInt32(int32(visualChain[0]))
				out.WriteInt32(int32(visualChain[1]))
			}

			if s.TargetFlags&DestinationLocation != 0 {
				out.WriteUint8(s.DestinationLocationCount)
			}

			if s.TargetFlags&ExtraTargets != 0 {
				out.WriteUint32(uint32(len(s.TargetLocations)))

				for _, location := range s.TargetLocations {
					if err = location.Encode(build, out); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

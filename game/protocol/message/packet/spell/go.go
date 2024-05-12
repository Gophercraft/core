package spell

import (
	"fmt"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

const (
	HasTargetMask TargetFlags = Unit | CorpseEnemy | GameObject | CorpseAlly | UnitMinipet
	HasItemTarget TargetFlags = Item | TradeItem
)

type Go struct {
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

func (g *Go) Encode(build version.Build, out *message.Packet) (err error) {
	out.Type = message.SMSG_SPELL_GO

	g.Caster.EncodePacked(build, out)
	g.Source.EncodePacked(build, out)

	if build.AddedIn(version.V3_0_2) {
		out.WriteUint32(g.CastCount)
	}

	out.WriteUint32(g.Spell)

	if build.AddedIn(version.V3_0_2) {
		out.WriteUint32(uint32(g.Flags))
	} else {
		out.WriteUint16(uint16(g.Flags))
	}

	if build.AddedIn(15005) {
		out.WriteUint64(g.CastTime)
	} else {
		out.WriteUint32(uint32(g.CastTime))
	}

	out.WriteUint32(uint32(len(g.HitTargets)))

	for _, target := range g.HitTargets {
		target.EncodeUnpacked(build, out)
	}

	for _, miss := range g.MissedTargets {
		if err = miss.Encode(build, out); err != nil {
			return err
		}
	}

	if err = g.Data.EncodeCastTargets(build, out); err != nil {
		return err
	}

	if build.AddedIn(9056) {
		if g.Flags&PredictedPower != 0 {
			if build.AddedIn(16309) {
				out.WriteUint32(uint32(len(g.ModifiedPowers)))
				for _, power := range g.ModifiedPowers {
					if err = power.Encode(build, out); err != nil {
						return
					}
				}
			} else {
				out.WriteUint32(g.RuneCooldown)
			}
		}

		if g.Flags&RuneInfo != 0 {
			out.WriteUint8(g.SpellRuneState)
			out.WriteUint8(g.PlayerRuneState)

			runeCooldownPasses := make([]byte, 6)
			copy(runeCooldownPasses, g.RuneCooldownPasses)

			for b := uint8(0); b < uint8(len(g.RuneCooldownPasses)); b++ {
				if build.RemovedIn(4545) {
					var mask uint8 = 1 << b
					if (mask & g.SpellRuneState) == 0 {
						continue
					}

					if (mask & g.PlayerRuneState) != 0 {
						continue
					}
				}

				out.WriteUint8(g.RuneCooldownPasses[int(b)])
			}
		}

		if g.Flags&AdjustMissile != 0 {
			out.WriteFloat32(g.Elevation)
			out.WriteInt32(g.DelayTime)
		}

		if g.Flags&Projectile != 0 {
			out.WriteInt32(g.AmmoDisplayID)
			out.WriteUint32(g.AmmoInventoryType)
		}

		if build.AddedIn(9056) {
			if g.Flags&VisualChain != 0 {
				visualChain := make([]int32, 2)
				copy(visualChain, g.VisualChain)
				out.WriteInt32(int32(visualChain[0]))
				out.WriteInt32(int32(visualChain[1]))
			}

			if g.TargetFlags&DestinationLocation != 0 {
				out.WriteUint8(g.DestinationLocationCount)
			}

			if g.TargetFlags&ExtraTargets != 0 {
				out.WriteUint32(uint32(len(g.TargetLocations)))

				for _, location := range g.TargetLocations {
					if err = location.Encode(build, out); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (g *Go) Decode(build version.Build, in *message.Packet) (err error) {
	// return g.CastData.Decode(build, in)
	g.Caster, err = guid.DecodePacked(build, in)
	if err != nil {
		return
	}

	if build.AddedIn(version.V3_0_2) {
		g.CastCount = uint32(in.ReadUint32())
	}

	g.Spell = in.ReadUint32()

	if build.AddedIn(version.V3_0_2) {
		g.Flags = CastFlags(in.ReadInt32())
	} else {
		g.Flags = CastFlags(in.ReadInt16())
	}

	if build.AddedIn(15005) {
		g.CastTime = in.ReadUint64()
	} else {
		g.CastTime = uint64(in.ReadUint32())
	}

	hitCount := int(in.ReadUint8())
	g.HitTargets = make([]guid.GUID, hitCount)

	for i := 0; i < hitCount; i++ {
		g.HitTargets[i], err = guid.DecodeUnpacked(build, in)
		if err != nil {
			return
		}
	}

	missCount := int(in.ReadUint8())
	g.MissedTargets = make([]MissStatus, missCount)

	for i := 0; i < missCount; i++ {
		st := &g.MissedTargets[i]
		if err = st.Decode(build, in); err != nil {
			return
		}
	}

	if err = g.Data.DecodeCastTargets(build, in); err != nil {
		return
	}

	if build.AddedIn(9056) {
		if g.Flags&PredictedPower != 0 {
			if build.AddedIn(16309) {
				modifiedPowerCount := int(in.ReadUint32())
				modifiedPowers := make([]ModifiedPower, modifiedPowerCount)
				for i := range modifiedPowers {
					if err = modifiedPowers[i].Decode(build, in); err != nil {
						return
					}
				}
				g.ModifiedPowers = modifiedPowers
			} else {
				g.RuneCooldown = in.ReadUint32()
			}
		}

		if g.Flags&RuneInfo != 0 {
			g.SpellRuneState = in.ReadUint8()
			g.PlayerRuneState = in.ReadUint8()

			g.RuneCooldownPasses = make([]byte, 6)

			for b := uint8(0); b < uint8(len(g.RuneCooldownPasses)); b++ {
				if build.RemovedIn(4545) {
					var mask uint8 = 1 << b
					if (mask & g.SpellRuneState) == 0 {
						continue
					}

					if (mask & g.PlayerRuneState) != 0 {
						continue
					}
				}

				g.RuneCooldownPasses[int(b)] = in.ReadUint8()
			}
		}

		if g.Flags&AdjustMissile != 0 {
			g.Elevation = in.ReadFloat32()
			g.DelayTime = in.ReadInt32()
		}

		if g.Flags&Projectile != 0 {
			g.AmmoDisplayID = in.ReadInt32()
			g.AmmoInventoryType = in.ReadUint32()
		}

		if build.AddedIn(9056) {
			if g.Flags&VisualChain != 0 {
				g.VisualChain = make([]int32, 2)
				g.VisualChain[0] = in.ReadInt32()
				g.VisualChain[1] = in.ReadInt32()
			}

			if g.TargetFlags&DestinationLocation != 0 {
				g.DestinationLocationCount = in.ReadUint8()
			}

			if g.TargetFlags&ExtraTargets != 0 {
				extraTargetCount := int(in.ReadUint32())

				const maxTargetCount int = (1 << 16)

				if extraTargetCount >= maxTargetCount {
					return fmt.Errorf("spell: malformed server packet: %d extra target counts?", extraTargetCount)
				}

				g.TargetLocations = make([]TargetLocation, extraTargetCount)

				for _, location := range g.TargetLocations {
					if err = location.Decode(build, in); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

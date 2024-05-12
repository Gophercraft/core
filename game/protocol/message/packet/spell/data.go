package spell

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/version"
)

type CastFlags uint64

const (
	None           CastFlags = 0x00000000
	PendingCast    CastFlags = 0x00000001 // 4.x NoCombatLog
	HasTrajectory  CastFlags = 0x00000002
	Unknown2       CastFlags = 0x00000004
	Unknown3       CastFlags = 0x00000008
	Unknown4       CastFlags = 0x00000010
	Projectile     CastFlags = 0x00000020
	Unknown5       CastFlags = 0x00000040
	Unknown6       CastFlags = 0x00000080
	Unknown7       CastFlags = 0x00000100
	Unknown8       CastFlags = 0x00000200
	Unknown9       CastFlags = 0x00000400
	PredictedPower CastFlags = 0x00000800
	Unknown10      CastFlags = 0x00001000
	Unknown11      CastFlags = 0x00002000
	Unknown12      CastFlags = 0x00004000
	Unknown13      CastFlags = 0x00008000
	Unknown14      CastFlags = 0x00010000
	AdjustMissile  CastFlags = 0x00020000 // 4.x
	NoGcd          CastFlags = 0x00040000
	VisualChain    CastFlags = 0x00080000 // 4.x
	Unknown18      CastFlags = 0x00100000
	RuneInfo       CastFlags = 0x00200000 // 4.x PredictedRunes
	Unknown19      CastFlags = 0x00400000
	Unknown20      CastFlags = 0x00800000
	Unknown21      CastFlags = 0x01000000
	Unknown22      CastFlags = 0x02000000
	Immunity       CastFlags = 0x04000000 // 4.x
	Unknown23      CastFlags = 0x08000000
	Unknown24      CastFlags = 0x10000000
	Unknown25      CastFlags = 0x20000000
	HealPrediction CastFlags = 0x40000000 // 4.x
	Unknown27      CastFlags = 0x80000000
)

type Data struct {
	GlyphIndex          int32
	Caster              guid.GUID
	Spell               uint32
	HitTargets          []guid.GUID
	TargetUnit          guid.GUID
	DestinationPosition tempest.C3Vector
	CastGUID            guid.GUID
	CastID              uint32
	Flags               CastFlags
	CastTime            uint64
	AmmoDisplayID       int32
	MissedTargets       []MissStatus
	TargetLocations     []TargetLocation
	AmmoInventoryType   uint32
	CastCount           uint32
	TargetFlags         TargetFlags
	Target              guid.GUID
	ItemTarget          guid.GUID

	SourceTransport      guid.GUID
	SourcePosition       tempest.C3Vector
	DestinationTransport guid.GUID
	TargetString         string
}

func (d *Data) DecodeCastTargets(build version.Build, in *message.Packet) (err error) {
	d.TargetFlags = TargetFlags(in.ReadUint32())

	if d.TargetFlags&HasTargetMask != 0 {
		d.Target, err = guid.DecodePacked(build, in)
		if err != nil {
			return err
		}
	}

	if d.TargetFlags&HasItemTarget != 0 {
		d.ItemTarget, err = guid.DecodePacked(build, in)
		if err != nil {
			return err
		}
	}

	if d.TargetFlags&SourceLocation != 0 {
		if build.AddedIn(10192) {
			d.SourceTransport, err = guid.DecodePacked(build, in)
			if err != nil {
				return err
			}

			d.SourcePosition, err = tempest.DecodeC3Vector(in)
			if err != nil {
				return err
			}
		}
	}

	if d.TargetFlags&DestinationLocation != 0 {
		if build.AddedIn(9464) {
			d.DestinationTransport, err = guid.DecodePacked(build, in)
			if err != nil {
				return err
			}
		}
		d.DestinationPosition, err = tempest.DecodeC3Vector(in)
		if err != nil {
			return err
		}
	}

	if d.TargetFlags&NameString != 0 {
		d.TargetString = in.ReadCString()
	}

	return nil
}

func (d *Data) EncodeCastTargets(build version.Build, out *message.Packet) error {
	out.WriteUint32(uint32(d.TargetFlags))

	if d.TargetFlags&HasTargetMask != 0 {
		d.Target.EncodePacked(build, out)
	}

	if d.TargetFlags&HasItemTarget != 0 {
		d.ItemTarget.EncodePacked(build, out)
	}

	if d.TargetFlags&SourceLocation != 0 {
		if build.AddedIn(10192) {
			d.SourceTransport.EncodePacked(build, out)
			d.SourcePosition.Encode(out)
		}
	}

	if d.TargetFlags&DestinationLocation != 0 {
		if build.AddedIn(9464) {
			d.DestinationTransport.EncodePacked(build, out)
		}
		d.DestinationPosition.Encode(out)
	}

	if d.TargetFlags&NameString != 0 {
		out.WriteCString(d.TargetString)
	}

	return nil
}

package spell

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/vsn"
)

type TargetFlags uint32

const (
	Self                TargetFlags = 0x00000000
	SpellDynamic1       TargetFlags = 0x00000001
	Unit                TargetFlags = 0x00000002
	UnitRaid            TargetFlags = 0x00000004
	UnitParty           TargetFlags = 0x00000008
	Item                TargetFlags = 0x00000010
	SourceLocation      TargetFlags = 0x00000020
	DestinationLocation TargetFlags = 0x00000040
	UnitEnemy           TargetFlags = 0x00000080
	UnitAlly            TargetFlags = 0x00000100
	CorpseEnemy         TargetFlags = 0x00000200
	UnitDead            TargetFlags = 0x00000400
	GameObject          TargetFlags = 0x00000800
	TradeItem           TargetFlags = 0x00001000
	NameString          TargetFlags = 0x00002000
	GameObjectItem      TargetFlags = 0x00004000
	CorpseAlly          TargetFlags = 0x00008000
	UnitMinipet         TargetFlags = 0x00010000
	Glyph               TargetFlags = 0x00020000
	DestinationTarget   TargetFlags = 0x00040000
	ExtraTargets        TargetFlags = 0x00080000 // 4.x VisualChain
	UnitPassenger       TargetFlags = 0x00100000
	Unk400000           TargetFlags = 0x00400000
	Unk1000000          TargetFlags = 0x01000000
	Unk4000000          TargetFlags = 0x04000000
	Unk10000000         TargetFlags = 0x10000000
	Unk40000000         TargetFlags = 0x40000000
)

type TargetLocation struct {
	Transport guid.GUID
	Location  tempest.C3Vector
}

func (t *TargetLocation) Encode(build vsn.Build, out *packet.WorldPacket) error {
	t.Transport.EncodePacked(build, out)
	return t.Location.Encode(out)
}

func (t *TargetLocation) Decode(build vsn.Build, in *packet.WorldPacket) error {
	var err error
	t.Transport, err = guid.DecodePacked(build, in)
	if err != nil {
		return err
	}
	t.Location, err = tempest.DecodeC3Vector(in)
	return err
}

type TargetData struct {
	Flags       TargetFlags
	Unit        guid.GUID
	Item        guid.GUID
	SrcLocation *TargetLocation
	DstLocation *TargetLocation
	Name        string
}

func (td *TargetData) Flag(mask TargetFlags) bool {
	return td.Flags&mask != 0
}

func (td *TargetData) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.WriteUint32(uint32(td.Flags))
	if td.Flag(Unit) {
		td.Unit.EncodePacked(build, out)
	}

	if td.Flag(Item) {
		td.Item.EncodePacked(build, out)
	}

	if td.Flag(SourceLocation) {
		if err := td.SrcLocation.Encode(build, out); err != nil {
			return err
		}
	}

	if td.Flag(DestinationLocation) {
		if err := td.DstLocation.Encode(build, out); err != nil {
			return err
		}
	}

	if td.Flag(NameString) {
		out.WriteCString(td.Name)
	}
	return nil
}

func (td *TargetData) Decode(build vsn.Build, in *packet.WorldPacket) error {
	td.Flags = TargetFlags(in.ReadUint32())

	var err error

	if td.Flag(Unit) {
		td.Unit, err = guid.DecodePacked(build, in)
		if err != nil {
			return err
		}
	}

	if td.Flag(Item) {
		td.Item, err = guid.DecodePacked(build, in)
		if err != nil {
			return err
		}
	}

	if td.Flag(SourceLocation) {
		td.SrcLocation = &TargetLocation{}
		if err := td.SrcLocation.Decode(build, in); err != nil {
			return err
		}
	}

	if td.Flag(DestinationLocation) {
		td.DstLocation = &TargetLocation{}
		if err := td.DstLocation.Decode(build, in); err != nil {
			return err
		}
	}

	if td.Flag(NameString) {
		td.Name = in.ReadCString()
	}
	return nil
}

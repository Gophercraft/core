package character

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type Char struct {
	GUID          guid.GUID
	Name          string
	Race          Race
	Class         Class
	BodyType      uint8
	Skin          uint8
	Face          uint8
	HairStyle     uint8
	HairColor     uint8
	FacialHair    uint8
	Level         uint8
	Zone          uint32
	Map           uint32
	X             float32
	Y             float32
	Z             float32
	Guild         uint32
	Flags         Flags
	Customization uint32
	FirstLogin    bool
	PetModel      uint32
	PetLevel      uint32
	PetFamily     uint32
	Equipment     []Equipment
}

type Equipment struct {
	Model       uint32
	Type        uint8
	Enchantment uint32
}

type List struct {
	Chars []Char
}

func (c *List) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_CHAR_ENUM
	out.WriteByte(uint8(len(c.Chars)))
	for _, char := range c.Chars {
		char.GUID.EncodeUnpacked(build, out)
		out.WriteCString(char.Name)
		out.WriteByte(uint8(char.Race))
		out.WriteByte(uint8(char.Class))
		out.WriteByte(char.BodyType)
		out.WriteByte(char.Skin)
		out.WriteByte(char.Face)
		out.WriteByte(char.HairStyle)
		out.WriteByte(char.HairColor)
		out.WriteByte(char.FacialHair)
		out.WriteByte(char.Level)
		out.WriteUint32(char.Zone)
		out.WriteUint32(char.Map)
		out.WriteFloat32(char.X)
		out.WriteFloat32(char.Y)
		out.WriteFloat32(char.Z)
		out.WriteUint32(char.Guild)

		if build.AddedIn(vsn.V1_12_1) {
			if err := char.Flags.Encode(build, out); err != nil {
				return err
			}
		}

		if build.AddedIn(vsn.V3_0_2) {
			out.WriteUint32(char.Customization)
		}

		if build.AddedIn(vsn.V1_12_1) {
			out.WriteBool(char.FirstLogin)
		}

		out.WriteUint32(char.PetModel)
		out.WriteUint32(char.PetLevel)
		out.WriteUint32(char.PetFamily)

		for i := 0; i < EquipLen(build); i++ {
			var eq Equipment

			if i < len(char.Equipment) {
				eq = char.Equipment[i]
			}

			out.WriteUint32(eq.Model)
			out.WriteByte(eq.Type)

			if build.AddedIn(vsn.V2_0_1) {
				out.WriteUint32(eq.Enchantment)
			}
		}

		if build.RemovedIn(vsn.V3_0_2) {
			// Bags
			out.WriteUint32(0)
			out.WriteByte(0)

			if build.AddedIn(vsn.V2_0_1) {
				out.WriteUint32(0)
			}
		}
	}
	return nil
}

func (chh *List) Decode(build vsn.Build, in *packet.WorldPacket) error {
	count := int(in.ReadByte())
	for x := 0; x < count; x++ {
		ch := Char{}
		var err error
		ch.GUID, err = guid.DecodeUnpacked(build, in)
		if err != nil {
			return err
		}
		ch.Name = in.ReadCString()
		ch.Race = Race(in.ReadByte())
		ch.Class = Class(in.ReadByte())
		ch.BodyType = in.ReadByte()
		ch.Skin = in.ReadByte()
		ch.Face = in.ReadByte()
		ch.HairStyle = in.ReadByte()
		ch.HairColor = in.ReadByte()
		ch.FacialHair = in.ReadByte()
		ch.Level = in.ReadByte()
		ch.Zone = in.ReadUint32()
		ch.Map = in.ReadUint32()
		ch.X = in.ReadFloat32()
		ch.Y = in.ReadFloat32()
		ch.Z = in.ReadFloat32()
		ch.Guild = in.ReadUint32()

		if err := ch.Flags.Decode(build, in); err != nil {
			return err
		}
		if build >= vsn.V3_0_2 {
			ch.Customization = in.ReadUint32()
		}
		if build.AddedIn(vsn.V1_12_1) {
			ch.FirstLogin = in.ReadBool()
		}
		ch.PetModel = in.ReadUint32()
		ch.PetLevel = in.ReadUint32()
		ch.PetFamily = in.ReadUint32()

		// Get equipment
		for j := 0; j < EquipLen(build); j++ {
			model := in.ReadUint32()
			typ := in.ReadByte()
			item := Equipment{
				Model: model,
				Type:  typ,
			}
			if build >= vsn.V3_0_2 {
				item.Enchantment = in.ReadUint32()
			}
			ch.Equipment = append(ch.Equipment, item)
		}

		if build.RemovedIn(vsn.V2_0_1) {
			//bags
			in.ReadUint32()
			in.ReadByte()
		}

		chh.Chars = append(chh.Chars, ch)
	}
	return nil
}

func EquipLen(build vsn.Build) int {
	switch {
	case vsn.Range(0, 8606).Contains(build):
		return 19
	default:
		return 23
	}
}

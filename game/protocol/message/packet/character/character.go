package character

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

type Character struct {
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
	Characters []Character
}

func (c *List) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_CHAR_ENUM
	out.WriteUint8(uint8(len(c.Characters)))
	for _, char := range c.Characters {
		char.GUID.EncodeUnpacked(build, out)
		out.WriteCString(char.Name)
		out.WriteUint8(uint8(char.Race))
		out.WriteUint8(uint8(char.Class))
		out.WriteUint8(char.BodyType)
		out.WriteUint8(char.Skin)
		out.WriteUint8(char.Face)
		out.WriteUint8(char.HairStyle)
		out.WriteUint8(char.HairColor)
		out.WriteUint8(char.FacialHair)
		out.WriteUint8(char.Level)
		out.WriteUint32(char.Zone)
		out.WriteUint32(char.Map)
		out.WriteFloat32(char.X)
		out.WriteFloat32(char.Y)
		out.WriteFloat32(char.Z)
		out.WriteUint32(char.Guild)

		if build.AddedIn(version.V1_12_1) {
			if err := char.Flags.Encode(build, out); err != nil {
				return err
			}
		}

		if build.AddedIn(version.V3_0_2) {
			out.WriteUint32(char.Customization)
		}

		if build.AddedIn(version.V1_12_1) {
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
			out.WriteUint8(eq.Type)

			if build.AddedIn(version.V2_0_1) {
				out.WriteUint32(eq.Enchantment)
			}
		}

		if build.RemovedIn(version.V3_0_2) {
			// Bags
			out.WriteUint32(0)
			out.WriteUint8(0)

			if build.AddedIn(version.V2_0_1) {
				out.WriteUint32(0)
			}
		}
	}
	return nil
}

func (chh *List) Decode(build version.Build, in *message.Packet) error {
	count := int(in.ReadUint8())
	for x := 0; x < count; x++ {
		ch := Character{}
		var err error
		ch.GUID, err = guid.DecodeUnpacked(build, in)
		if err != nil {
			return err
		}
		ch.Name = in.ReadCString()
		ch.Race = Race(in.ReadUint8())
		ch.Class = Class(in.ReadUint8())
		ch.BodyType = in.ReadUint8()
		ch.Skin = in.ReadUint8()
		ch.Face = in.ReadUint8()
		ch.HairStyle = in.ReadUint8()
		ch.HairColor = in.ReadUint8()
		ch.FacialHair = in.ReadUint8()
		ch.Level = in.ReadUint8()
		ch.Zone = in.ReadUint32()
		ch.Map = in.ReadUint32()
		ch.X = in.ReadFloat32()
		ch.Y = in.ReadFloat32()
		ch.Z = in.ReadFloat32()
		ch.Guild = in.ReadUint32()

		if err := ch.Flags.Decode(build, in); err != nil {
			return err
		}
		if build >= version.V3_0_2 {
			ch.Customization = in.ReadUint32()
		}
		if build.AddedIn(version.V1_12_1) {
			ch.FirstLogin = in.ReadBool()
		}
		ch.PetModel = in.ReadUint32()
		ch.PetLevel = in.ReadUint32()
		ch.PetFamily = in.ReadUint32()

		// Get equipment
		for j := 0; j < EquipLen(build); j++ {
			model := in.ReadUint32()
			typ := in.ReadUint8()
			item := Equipment{
				Model: model,
				Type:  typ,
			}
			if build >= version.V3_0_2 {
				item.Enchantment = in.ReadUint32()
			}
			ch.Equipment = append(ch.Equipment, item)
		}

		if build.RemovedIn(version.V2_0_1) {
			//bags
			in.ReadUint32()
			in.ReadUint8()
		}

		chh.Characters = append(chh.Characters, ch)
	}
	return nil
}

func EquipLen(build version.Build) int {
	switch {
	case version.Range(0, 8606).Contains(build):
		return 19
	default:
		return 23
	}
}

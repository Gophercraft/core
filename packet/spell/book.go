package spell

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type Page struct {
	Spell uint32
	Flag  uint16
	Slot  uint16
}

type Cooldown struct {
	Spell                uint32
	Item                 uint32
	Category             uint16
	RecoveryTime         uint32
	CategoryRecoveryTime uint32
}

func (c *Cooldown) Encode(build vsn.Build, out *packet.WorldPacket) error {
	switch {
	case build < 9767:
		out.WriteUint16(uint16(c.Spell))
	default:
		out.WriteUint32(uint32(c.Spell))
	}

	switch {
	case build < 14545:
		out.WriteUint16(uint16(c.Item))
	default:
		out.WriteUint32(uint32(c.Item))
	}

	out.WriteUint16(c.Category)
	out.WriteUint32(c.RecoveryTime)
	out.WriteUint32(c.CategoryRecoveryTime)

	return nil
}

type Book struct {
	TalentSpec uint8

	Spells   []Page
	Cooldown []Cooldown
}

func (p *Page) Encode(build vsn.Build, out *packet.WorldPacket) error {
	if build == vsn.Alpha {
		out.WriteUint16(uint16(p.Spell))
		out.WriteUint16(uint16(p.Slot))
		return nil
	}

	out.WriteUint32(p.Spell)
	if build.AddedIn(vsn.V3_0_2) {
		out.WriteUint16(p.Flag)
	}
	return nil
}

func (p *Page) Decode(build vsn.Build, in *packet.WorldPacket) error {
	if build == vsn.Alpha {
		p.Spell = uint32(in.ReadUint16())
		p.Slot = in.ReadUint16()
		return nil
	}

	p.Spell = in.ReadUint32()
	if build.AddedIn(vsn.V3_0_2) {
		p.Flag = in.ReadUint16()
	}
	return nil
}

func (l *Book) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_INITIAL_SPELLS
	out.WriteByte(l.TalentSpec)
	out.WriteUint16(uint16(len(l.Spells)))
	for _, spell := range l.Spells {
		if err := spell.Encode(build, out); err != nil {
			return err
		}
	}
	out.WriteUint16(uint16(len(l.Cooldown)))
	for _, cooldown := range l.Cooldown {
		if err := cooldown.Encode(build, out); err != nil {
			return err
		}
	}
	return nil
}

type Learned struct {
	Page
}

func (l *Learned) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_LEARNED_SPELL
	return l.Page.Encode(build, out)
}

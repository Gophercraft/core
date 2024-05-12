package spell

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

type AuraCastLog struct {
	// packet.ReadGuid("Caster GUID");
	// packet.ReadGuid("Target GUID");
	// packet.ReadUInt32<SpellId>("Spell ID");
	// packet.ReadSingle("Unk 1");
	// packet.ReadSingle("Unk 2");
	Caster guid.GUID
	Target guid.GUID
	Spell  uint32
	X      float32
	Y      float32
}

func (lg *AuraCastLog) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_AURACASTLOG
	lg.Caster.EncodeUnpacked(build, out)
	lg.Target.EncodeUnpacked(build, out)
	out.WriteUint32(lg.Spell)
	out.WriteFloat32(lg.X)
	out.WriteFloat32(lg.Y)
	return nil
}

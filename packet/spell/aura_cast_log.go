package spell

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
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

func (lg *AuraCastLog) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_AURACASTLOG
	lg.Caster.EncodeUnpacked(build, out)
	lg.Target.EncodeUnpacked(build, out)
	out.WriteUint32(lg.Spell)
	out.WriteFloat32(lg.X)
	out.WriteFloat32(lg.Y)
	return nil
}

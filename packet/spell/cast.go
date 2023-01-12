package spell

import (
	p "github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type Cast struct {
	Data
}

func (c *Cast) Decode(build vsn.Build, packet *p.WorldPacket) (err error) {
	// c.Data.CastCount = uint32(packet.ReadByte())

	c.Data.Spell = packet.ReadUint32()

	if build.AddedIn(vsn.V4_0_1) {
		c.Data.GlyphIndex = packet.ReadInt32()
	}

	c.Data.Flags = CastFlags(packet.ReadByte())

	if c.Flags&HasTrajectory != 0 {
		if err = c.Data.DecodeCastTargets(build, packet); err != nil {
			return
		}
	}

	return nil
}

func (c *Cast) Encode(build vsn.Build, packet p.WorldPacket) error {
	panic("cast encode nyi")
	return nil
}

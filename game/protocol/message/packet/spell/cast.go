package spell

import (
	p "github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/version"
)

type Cast struct {
	Data
}

func (c *Cast) Decode(build version.Build, packet *p.WorldPacket) (err error) {
	// c.Data.CastCount = uint32(packet.ReadUint8())

	c.Data.Spell = packet.ReadUint32()

	if build.AddedIn(version.V4_0_1) {
		c.Data.GlyphIndex = packet.ReadInt32()
	}

	c.Data.Flags = CastFlags(packet.ReadUint8())

	if c.Flags&HasTrajectory != 0 {
		if err = c.Data.DecodeCastTargets(build, packet); err != nil {
			return
		}
	}

	return nil
}

func (c *Cast) Encode(build version.Build, packet p.WorldPacket) error {
	panic("cast encode nyi")
	return nil
}

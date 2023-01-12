package spell

import (
	"fmt"

	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type RuneData struct {
	SpellState  uint8
	PlayerState uint8
	Cooldowns   []uint8
}

func (rd *RuneData) cooldownSize(b vsn.Build) int {
	return 6
}

func (rd *RuneData) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.WriteByte(rd.SpellState)
	out.WriteByte(rd.PlayerState)

	bild := vsn.Build(build)

	cooldownSize := rd.cooldownSize(build)

	if len(rd.Cooldowns) != cooldownSize {
		err = fmt.Errorf("packet/spell: must have %s cooldowns", bild)
		return err
	}

	for _, cool := range rd.Cooldowns {
		out.WriteByte(cool)
	}
	return nil
}

func (rd *RuneData) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	rd.SpellState = in.ReadByte()
	rd.PlayerState = in.ReadByte()

	rd.Cooldowns = make([]uint8, rd.cooldownSize(build))
	for i := range rd.Cooldowns {
		rd.Cooldowns[i] = in.ReadByte()
	}

	return nil
}

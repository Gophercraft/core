package spell

import (
	"fmt"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type RuneData struct {
	SpellState  uint8
	PlayerState uint8
	Cooldowns   []uint8
}

func (rd *RuneData) cooldownSize(b version.Build) int {
	return 6
}

func (rd *RuneData) Encode(build version.Build, out *message.Packet) (err error) {
	out.WriteUint8(rd.SpellState)
	out.WriteUint8(rd.PlayerState)

	bild := version.Build(build)

	cooldownSize := rd.cooldownSize(build)

	if len(rd.Cooldowns) != cooldownSize {
		err = fmt.Errorf("packet/spell: must have %s cooldowns", bild)
		return err
	}

	for _, cool := range rd.Cooldowns {
		out.WriteUint8(cool)
	}
	return nil
}

func (rd *RuneData) Decode(build version.Build, in *message.Packet) (err error) {
	rd.SpellState = in.ReadUint8()
	rd.PlayerState = in.ReadUint8()

	rd.Cooldowns = make([]uint8, rd.cooldownSize(build))
	for i := range rd.Cooldowns {
		rd.Cooldowns[i] = in.ReadUint8()
	}

	return nil
}

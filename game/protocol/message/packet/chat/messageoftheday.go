package chat

import (
	"fmt"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type MOTD struct {
	Lines []string
}

func (m *MOTD) Encode(build version.Build, out *message.Packet) error {
	if build.AddedIn(version.V2_0_1) {
		out.Type = message.SMSG_MOTD
		out.WriteUint32(uint32(len(m.Lines)))

		for _, el := range m.Lines {
			out.WriteCString(el)
		}

		return nil
	}
	return fmt.Errorf("chat: MOTD not implemented in %s", build)
}

func (m *MOTD) Decode(build version.Build, in *message.Packet) error {
	if build.AddedIn(version.V2_0_1) {
		count := in.ReadUint32()
		if count > 100 {
			count = 100
		}

		m.Lines = make([]string, count)
		for i := range m.Lines {
			m.Lines[i] = in.ReadCString()
		}

	}
	return nil
}

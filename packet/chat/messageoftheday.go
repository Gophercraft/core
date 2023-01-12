package chat

import (
	"fmt"

	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type MOTD struct {
	Lines []string
}

func (m *MOTD) Encode(build vsn.Build, out *packet.WorldPacket) error {
	if build.AddedIn(vsn.V2_0_1) {
		out.Type = packet.SMSG_MOTD
		out.WriteUint32(uint32(len(m.Lines)))

		for _, el := range m.Lines {
			out.WriteCString(el)
		}

		return nil
	}
	return fmt.Errorf("chat: MOTD not implemented in %s", build)
}

func (m *MOTD) Decode(build vsn.Build, in *packet.WorldPacket) error {
	if build.AddedIn(vsn.V2_0_1) {
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

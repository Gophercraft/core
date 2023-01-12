package realm

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

// Sends metadata important to handshake after session is initially confirmed
func (s *Session) SendSessionMetadata() {
	if s.Build().AddedIn(vsn.V3_0_2) {
		s.Send(&packet.ClientCacheVersion{
			Build: s.Build(),
		})
	}
}

func (s *Session) SendUnlearnSpell() {
	p := packet.NewWorldPacket(packet.SMSG_SEND_UNLEARN_SPELLS)
	p.WriteUint32(0)
	s.SendPacket(p)
}

// func (s *Session) SendWorldLoginMetadata() {

// }

// func (s *Session) SendMetadataAfterSpawn() {
// 	if s.Build().AddedIn(vsn.V2_4_3) {

// 	}
// }

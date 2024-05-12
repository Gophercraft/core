package teleport

import (
	"time"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

// pkt := packet.NewWorldPacket(packet.SMSG_SUMMON_REQUEST)
// summoner.EncodeUnpacked(s.Build(), pkt)
// pkt.WriteUint32(zoneID)
// pkt.WriteUint32(uint32(timeout / time.Millisecond))
// s.SendPacket(pkt)

// s.Send(&teleport.SummonRequest{
// 	ID: summoner,
// 	Zone: zoneID,
// 	Timeout: timeout,
// })

type SummonRequest struct {
	ID      guid.GUID
	Zone    uint32
	Timeout time.Duration
}

func (s *SummonRequest) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_SUMMON_REQUEST
	s.ID.EncodeUnpacked(build, out)
	out.WriteUint32(s.Zone)
	out.WriteUint32(uint32(s.Timeout / time.Millisecond))
	return nil
}

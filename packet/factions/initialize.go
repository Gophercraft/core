package factions

import (
	"encoding/binary"

	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
	"github.com/superp00t/etc"
)

type ReputationFlags uint8

const (
	ReputationVisible         ReputationFlags = 1 << iota // makes visible in client (set or can be set at interaction with target of this faction)
	ReputationAtWar                                       // enable AtWar-button in client. player controlled (except opposition team always war state), Flag only set on initial creation
	ReputationHidden                                      // hidden faction from reputation pane in client (player can gain reputation, but this update not sent to client)
	ReputationInvisibleForced                             // always overwrite FACTION_FLAG_VISIBLE and hide faction in rep.list, used for hide opposite team factions
	ReputationPeaceForced                                 // always overwrite FACTION_FLAG_AT_WAR, used for prevent war with own team factions
	ReputationInactive                                    // player controlled, state stored in characters.data ( CMSG_SET_FACTION_INACTIVE )
	ReputationRival                                       // flag for the two competing outland factions
)

type ReputationRank float32

type Info struct {
	Flags    ReputationFlags
	Standing ReputationRank
}

type Initialize struct {
	Flags    uint32
	Factions []Info
}

func (init *Initialize) InitDefault(build vsn.Build) {
	init.Flags = 0x00000040
	init.Factions = make([]Info, 64)
}

func (init *Initialize) Encode(build vsn.Build, out *packet.WorldPacket) error {
	var facSize = 64

	out.Type = packet.SMSG_INITIALIZE_FACTIONS

	data := make([]byte, 4+(facSize*5))

	binary.LittleEndian.PutUint32(data, uint32(init.Flags))

	for i, fac := range init.Factions {
		offset := 4 + (i * 5)
		if offset >= len(data) {
			break
		}
		rec := data[offset : offset+5]
		rec[0] = uint8(fac.Flags)
		binary.LittleEndian.PutUint32(rec[1:], uint32(fac.Standing))
	}

	out.Buffer = etc.OfBytes(data)
	return nil
}

func (init *Initialize) Decode(build vsn.Build, in *packet.WorldPacket) error {
	init.Flags = in.ReadUint32()

	init.Factions = make([]Info, in.Available()/5)

	i := 0

	for in.Available() > 0 {
		init.Factions[i].Flags = ReputationFlags(in.ReadByte())
		init.Factions[i].Standing = ReputationRank(in.ReadUint32())
		i++
	}

	return nil
}

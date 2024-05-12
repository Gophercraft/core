package query

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/version"
)

type Name struct {
	ID guid.GUID
}

func (nq *Name) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_NAME_QUERY
	return nq.ID.EncodeUnpacked(build, out)
}

func (nq *Name) Decode(build version.Build, in *message.Packet) (err error) {
	nq.ID, err = guid.DecodeUnpacked(build, in)
	return
}

// resp := packet.NewWorldPacket(packet.SMSG_NAME_QUERY_RESPONSE)
// id := guid.RealmSpecific(guid.Player, s.Server.RealmID(), char.ID)
// if s.Build().AddedIn(version.V3_3_5a) {
// 	id.EncodePacked(s.Build(), resp)
// } else {
// 	id.EncodeUnpacked(s.Build(), resp)
// }

// if s.Build().AddedIn(version.V3_3_5a) {
// 	if char == nil || char.Name == "" {
// 		resp.WriteUint8(1)
// 		s.SendPacket(resp)
// 		return
// 	} else {
// 		resp.WriteUint8(0)
// 	}
// }

// resp.WriteCString(char.Name)
// // resp.WriteCString(s.Config().RealmName)
// resp.WriteUint8(0)
// resp.WriteUint32(uint32(char.Race))
// resp.WriteUint32(uint32(char.BodyType))
// resp.WriteUint32(uint32(char.Class))

//	if s.Build().AddedIn(version.V2_4_3) {
//		resp.WriteUint8(0)
//	}
type NameResponse struct {
	ID        guid.GUID
	NotFound  bool
	Name      string
	RealmName string
	Race      models.Race
	Class     models.Class
	BodyType  models.BodyType

	DeclinedNames []string
}

func (nr *NameResponse) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_NAME_QUERY_RESPONSE

	if build.AddedIn(9767) {
		nr.ID.EncodePacked(build, out)
		out.WriteBool(nr.NotFound)
		if nr.NotFound {
			return nil
		}
	} else {
		nr.ID.EncodeUnpacked(build, out)
	}

	out.WriteCString(nr.Name)
	out.WriteCString(nr.RealmName)

	if build.AddedIn(9767) {
		out.WriteUint8(uint8(nr.Race))
		out.WriteUint8(uint8(nr.BodyType))
		out.WriteUint8(uint8(nr.Class))
	} else {
		out.WriteUint32(uint32(nr.Race))
		out.WriteUint32(uint32(nr.BodyType))
		out.WriteUint32(uint32(nr.Class))
	}

	hasDeclinedNames := len(nr.DeclinedNames) >= 5

	out.WriteBool(hasDeclinedNames)

	if hasDeclinedNames {
		for i := 0; i < 5; i++ {
			out.WriteCString(nr.DeclinedNames[i])
		}
	}

	return nil
}

func (nr *NameResponse) Decode(build version.Build, in *message.Packet) (err error) {
	if build.AddedIn(9767) {
		nr.ID, err = guid.DecodePacked(build, in)
		if err != nil {
			return
		}
		nr.NotFound = in.ReadBool()
		if nr.NotFound {
			return nil
		}
	} else {
		nr.ID, err = guid.DecodeUnpacked(build, in)
		if err != nil {
			return
		}
	}

	nr.Name = in.ReadCString()
	nr.RealmName = in.ReadCString()

	if build.AddedIn(9767) {
		nr.Race = models.Race(in.ReadUint8())
		nr.BodyType = models.BodyType(in.ReadUint8())
		nr.Class = models.Class(in.ReadUint8())
	} else {
		nr.Race = models.Race(in.ReadUint32())
		nr.BodyType = models.BodyType(in.ReadUint32())
		nr.Class = models.Class(in.ReadUint32())
	}

	hasDeclinedNames := in.ReadBool()

	if hasDeclinedNames {
		nr.DeclinedNames = make([]string, 5)
		for i := 0; i < len(nr.DeclinedNames); i++ {
			nr.DeclinedNames[i] = in.ReadCString()
		}
	}

	return
}

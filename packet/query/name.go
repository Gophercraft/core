package query

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
)

type Name struct {
	ID guid.GUID
}

func (nq *Name) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.CMSG_NAME_QUERY
	return nq.ID.EncodeUnpacked(build, out)
}

func (nq *Name) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	nq.ID, err = guid.DecodeUnpacked(build, in)
	return
}

// resp := packet.NewWorldPacket(packet.SMSG_NAME_QUERY_RESPONSE)
// id := guid.RealmSpecific(guid.Player, s.Server.RealmID(), char.ID)
// if s.Build().AddedIn(vsn.V3_3_5a) {
// 	id.EncodePacked(s.Build(), resp)
// } else {
// 	id.EncodeUnpacked(s.Build(), resp)
// }

// if s.Build().AddedIn(vsn.V3_3_5a) {
// 	if char == nil || char.Name == "" {
// 		resp.WriteByte(1)
// 		s.SendPacket(resp)
// 		return
// 	} else {
// 		resp.WriteByte(0)
// 	}
// }

// resp.WriteCString(char.Name)
// // resp.WriteCString(s.Config().RealmName)
// resp.WriteByte(0)
// resp.WriteUint32(uint32(char.Race))
// resp.WriteUint32(uint32(char.BodyType))
// resp.WriteUint32(uint32(char.Class))

//	if s.Build().AddedIn(vsn.V2_4_3) {
//		resp.WriteByte(0)
//	}
type NameResponse struct {
	ID        guid.GUID
	NotFound  bool
	Name      string
	RealmName string
	Race      models.Race
	Class     models.Class
	BodyType  uint8

	DeclinedNames []string
}

func (nr *NameResponse) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_NAME_QUERY_RESPONSE

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
		out.WriteByte(uint8(nr.Race))
		out.WriteByte(nr.BodyType)
		out.WriteByte(uint8(nr.Class))
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

func (nr *NameResponse) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
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
		nr.Race = models.Race(in.ReadByte())
		nr.BodyType = in.ReadByte()
		nr.Class = models.Class(in.ReadByte())
	} else {
		nr.Race = models.Race(in.ReadUint32())
		nr.BodyType = byte(in.ReadUint32())
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

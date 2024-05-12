package gameobject

import (
	"fmt"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/packet/query"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/version"
)

const (
	TypeDoor                 = 0
	TypeButton               = 1
	TypeQuestgiver           = 2
	TypeChest                = 3
	TypeBinder               = 4
	TypeGeneric              = 5
	TypeTrap                 = 6
	TypeChair                = 7
	TypeSpellFocus           = 8
	TypeText                 = 9
	TypeGoober               = 10
	TypeTransport            = 11
	TypeAreadamage           = 12
	TypeCamera               = 13
	TypeMapObject            = 14
	TypeMoTransport          = 15
	TypeDuelArbiter          = 16
	TypeFishingnode          = 17
	TypeRitual               = 18
	TypeMailbox              = 19
	TypeAuctionHouse         = 20
	TypeGuardpost            = 21
	TypeSpellcaster          = 22
	TypeMeetingstone         = 23
	TypeFlagstand            = 24
	TypeFishinghole          = 25
	TypeFlagdrop             = 26
	TypeMiniGame             = 27
	TypeLotteryKiosk         = 28
	TypeCapturePoint         = 29
	TypeAuraGenerator        = 30
	TypeDungeonDifficulty    = 31
	TypeBarberChair          = 32
	TypeDestructibleBuilding = 33
	TypeGuildBank            = 34
)

type Use struct {
	ID guid.GUID
}

func (gou *Use) Encode(build version.Build, out *message.Packet) error {
	gou.ID.EncodeUnpacked(build, out)
	return nil
}

func (gou *Use) Decode(build version.Build, in *message.Packet) error {
	var err error
	gou.ID, err = guid.DecodeUnpacked(build, in)
	return err
}

type Query struct {
	ID uint32
}

func (gq *Query) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_GAMEOBJECT_QUERY
	out.WriteUint32(gq.ID)
	return nil
}

func (gq *Query) Decode(build version.Build, in *message.Packet) error {
	gq.ID = in.ReadUint32()
	if gq.ID >= query.EntryNotFound {
		return fmt.Errorf("packet: client sent gameobject query ID larger than query.EntryNotFound flag")
	}
	return nil
}

type QueryResponse struct {
	ID uint32
	// Locale should be set by client and server. If not, English will be used as the default!
	Locale     i18n.Locale
	GameObject *models.GameObjectTemplate
}

func (gr *QueryResponse) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_GAMEOBJECT_QUERY_RESPONSE

	qID := gr.ID
	if qID >= query.EntryNotFound {
		return fmt.Errorf("packet: gameobject query response ID is too large. this should not happen without large amounts of tomfoolery")
	}

	if gr.GameObject == nil {
		qID |= query.EntryNotFound
	}

	out.WriteUint32(qID)

	if qID&query.EntryNotFound != 0 {
		return nil
	}

	out.WriteUint32(gr.GameObject.Type)
	out.WriteUint32(gr.GameObject.DisplayID)
	out.WriteCString(gr.GameObject.Name.GetLocalized(gr.Locale))
	for i := 0; i < 3; i++ {
		out.WriteCString("")
	}
	out.WriteCString(gr.GameObject.IconName)
	out.WriteCString(gr.GameObject.CastBarCaption)
	out.WriteCString("") // unk

	for x := 0; x < 24; x++ {
		if x < len(gr.GameObject.Data) {
			out.WriteUint32(gr.GameObject.Data[x])
		} else {
			out.WriteUint32(0)
		}
	}
	return nil
}

func (gr *QueryResponse) Decode(build version.Build, in *message.Packet) error {
	gr.ID = in.ReadUint32()

	if gr.ID&query.EntryNotFound != 0 {
		return nil
	}

	gr.GameObject.Type = in.ReadUint32()
	gr.GameObject.DisplayID = in.ReadUint32()
	gr.GameObject.Name = i18n.Text{
		gr.Locale: in.ReadCString(),
	}
	for i := 0; i < 3; i++ {
		in.ReadCString()
	}

	gr.GameObject.IconName = in.ReadCString()
	gr.GameObject.CastBarCaption = in.ReadCString()
	in.ReadCString()

	for i := 0; i < 24; i++ {
		gr.GameObject.Data = append(gr.GameObject.Data, in.ReadUint32())
	}
	return nil
}

package gossip

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

const (
	IconChat      = iota // White chat bubble
	IconVendor           // 1 Brown bag
	IconTaxi             // 2 Flight
	IconTrainer          // 3 Book
	IconInteract1        // 4	Interaction wheel
	IconInteract2        // 5	Interaction wheel
	IconGold             // 6 Brown bag with yellow dot (gold)
	IconTalk             // White chat bubble with black dots (...)
	IconTabard           // 8 Tabard
	IconBattle           // 9 Two swords
	IconDot              // 10 Yellow dot
	IconChat11           // 11	White chat bubble
	IconChat12           // 12	White chat bubble
	IconChat13           // 13	White chat bubble
	IconInvalid14        // 14	INVALID - DO NOT USE
	IconInvalid15        // 15	INVALID - DO NOT USE
	IconChat16           // 16	White chat bubble
	IconChat17           // 17	White chat bubble
	IconChat18           // 18	White chat bubble
	IconChat19           // 19	White chat bubble
	IconChat20           // 20	White chat bubble
	IconTransmog         // 21	Transmogrifier?
)

type Item struct {
	ItemID  uint32
	Icon    uint8
	Coded   bool
	Message string
}

type QuestItem struct {
	QuestID    uint32
	QuestIcon  uint32
	QuestLevel int32
	QuestTitle string
}

type Menu struct {
	Speaker    guid.GUID
	TextEntry  uint32
	Items      []Item
	QuestItems []QuestItem
}

func NewMenu(id guid.GUID, textID uint32) *Menu {
	return &Menu{Speaker: id, TextEntry: textID}
}

func (g *Menu) SetTextEntry(entry uint32) {
	g.TextEntry = entry
}

func (g *Menu) AddItem(i Item) {
	g.Items = append(g.Items, i)
}

func (g *Menu) Encode(build vsn.Build, p *packet.WorldPacket) error {
	p.Type = packet.SMSG_GOSSIP_MESSAGE
	g.Speaker.EncodeUnpacked(build, p)
	p.WriteUint32(g.TextEntry)
	p.WriteUint32(uint32(len(g.Items)))
	for _, item := range g.Items {
		p.WriteUint32(item.ItemID)
		p.WriteByte(item.Icon)
		p.WriteBool(item.Coded)
		p.WriteCString(item.Message)
	}

	p.WriteUint32(uint32(len(g.QuestItems)))
	for _, qItem := range g.QuestItems {
		p.WriteUint32(qItem.QuestID)
		p.WriteUint32(qItem.QuestIcon)
		p.WriteInt32(qItem.QuestLevel)
		p.WriteCString(qItem.QuestTitle)
	}

	return nil
}

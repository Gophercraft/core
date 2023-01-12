package realm

import (
	"fmt"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet/gossip"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
)

func (s *Session) SendGossip(m *gossip.Menu) {
	s.Send(m)
}

func (s *Session) GetValidGossipObject(id guid.GUID) (WorldObject, string) {
	fmt.Println("Asked to speak with ", id)

	object := s.Map().GetObject(id)
	if object == nil {
		return nil, ""
	}

	var menuID string

	switch object.TypeID() {
	case guid.TypeUnit:
		var creatureTemplate *models.CreatureTemplate
		s.DB().Lookup(wdb.BucketKeyEntry, object.Values().Get("Entry").Uint32(), &creatureTemplate)
		if creatureTemplate == nil {
			return nil, ""
		} else {
			menuID = creatureTemplate.GossipMenuId
		}
		if creatureTemplate.Gossip == false && creatureTemplate.Innkeeper == false {
			return nil, ""
		}
	case guid.TypeGameObject:
		var objectTemplate *models.GameObjectTemplate
		s.DB().Lookup(wdb.BucketKeyEntry, object.Values().Get("Entry").Uint32(), &objectTemplate)
		if objectTemplate != nil {

		}
	default:
		fmt.Println("Client tried to speak with ", object.TypeID())
		return nil, ""
	}

	// Todo: invalidate if too far away.
	return object, menuID
}

func (s *Session) HandleGossipHello(hello *gossip.Hello) {
	object, menuID := s.GetValidGossipObject(hello.ID)
	if object == nil {
		return
	}

	fmt.Println("found object", object)

	// GameObjects will be supported in the future.

	if menuID == "" {
		fmt.Println("No menu found")
		return
	}

	menu := &gossip.Menu{
		Speaker:   hello.ID,
		TextEntry: 0,
	}

	// Quests should be offered no matter what.
	s.Server.Call(GossipEvent, s, menuID, 0, menu)

	s.SendGossip(menu)
}

func (s *Session) HandleGossipSelectOption(so *gossip.SelectOption) {
	id := so.ID
	option := so.Option

	object, menuID := s.GetValidGossipObject(id)
	if object == nil {
		return
	}

	menu := &gossip.Menu{
		Speaker:   id,
		TextEntry: 0,
	}

	// Quests should be offered no matter what.
	s.Server.Call(GossipEvent, s, menuID, option, menu)
	s.SendGossip(menu)
}

func (s *Session) HandleGossipTextQuery(tq *gossip.TextQuery) {
	var nt *models.NPCText
	s.DB().Lookup(wdb.BucketKeyEntry, tq.Entry, &nt)

	var response gossip.TextUpdate
	response.Entry = tq.Entry
	response.Locale = s.Locale
	if nt != nil {
		response.Text = nt
	}

	s.Send(&response)
}

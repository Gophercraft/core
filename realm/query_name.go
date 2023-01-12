package realm

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet/query"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/log"
)

func (s *Session) HandleNameQuery(nq *query.Name) {
	var char models.Character

	found, err := s.Server.DB.Where("id = ?", nq.ID.Counter()).Get(&char)
	if err != nil {
		s.DB().Println(err)
		return
	}

	if !found {
		log.Warn("No such data exists for", nq.ID)
		return
	}

	s.SendNameQueryResponseFor(&char)
}

func (s *Session) SendNameQueryResponseFor(char *models.Character) {
	if char == nil {
		return
	}

	var resp query.NameResponse
	resp.ID = guid.RealmSpecific(guid.Player, s.Server.RealmID(), char.ID)
	resp.NotFound = char == nil || char.Name == ""

	if !resp.NotFound {
		resp.Name = char.Name
		resp.RealmName = ""
		resp.Race = char.Race
		resp.Class = char.Class
		resp.BodyType = char.BodyType
	}

	s.Send(&resp)
}

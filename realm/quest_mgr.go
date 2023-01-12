package realm

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet/quest"
)

func (s *Session) QuestDone(q uint32) bool {
	return true
}

func (s *Session) HandleQuestgiverStatusQuery(q *quest.GiverStatusQuery) {
	objectGUID := q.ID
	object := s.Map().GetObject(objectGUID)
	if object == nil {
		return
	}

	if object.TypeID() == guid.TypePlayer {
		s.SendQuestGiverStatus(objectGUID, quest.Unavailable)
		return
	}

	var dialogStatus quest.GiverStatus = quest.Unavailable

	switch object.TypeID() {
	case guid.TypeUnit:
	}

	s.SendQuestGiverStatus(objectGUID, dialogStatus)
}

func (s *Session) SendQuestGiverStatus(id guid.GUID, dialogStatus quest.GiverStatus) {
	var resp quest.GiverStatusResponse
	resp.ID = id
	resp.Status = dialogStatus
	s.Send(&resp)
}

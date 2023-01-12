package realm

import (
	"strings"

	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/text"
)

const ObjectDebug models.PropID = "gophercraft.realm.ObjectDebug"
const NullProp models.PropID = ""

var (
	DefaultProps = map[models.PropID]string{
		ObjectDebug: "true",
	}
)

func (s *Session) BoolProp(id models.PropID) bool {
	var truth bool
	s.GetProp(id, &truth)
	return truth
}

func unmarshalProp(str string, to interface{}) {
	dec := text.NewDecoder(strings.NewReader(str))
	if err := dec.Decode(to); err != nil {
		panic(err)
	}
}

func (s *Session) GetProp(id models.PropID, value interface{}) bool {
	if s.HasState(InWorld) == false {
		return false
	}

	idx := s.GetPropIndex(id)
	if idx < 0 {
		return false
	}

	s.GuardSession.Lock()
	unmarshalProp(s.CharacterProps[idx].Value, value)
	s.GuardSession.Unlock()
	return true
}

func (s *Session) GetPropIndex(id models.PropID) int {
	s.GuardSession.Lock()
	defer s.GuardSession.Unlock()
	for i, prop := range s.CharacterProps {
		if prop.PropID == id {
			return i
		}
	}
	return -1
}

func (s *Session) SetStringProp(id models.PropID, value string) {
	s.SetProp(id, value)
}

func (s *Session) SetBoolProp(id models.PropID, value bool) {
	s.SetProp(id, value)
}

func (s *Session) SetProp(id models.PropID, value interface{}) {
	if value == nil {
		value = true
	}
	if s.GetProp(id, value) {
		s.RemoveProp(id)
	}

	j, err := text.Marshal(value)
	if err != nil {
		panic(err)
	}
	s.GuardSession.Lock()
	defer s.GuardSession.Unlock()
	var newProp = models.CharacterProp{
		ID:     s.PlayerID(),
		PropID: id,
		Value:  string(j),
	}

	s.CharacterProps = append(s.CharacterProps, newProp)
	s.DB().Insert(&newProp)
}

func (s *Session) RemoveProp(id models.PropID) {
	if s.HasState(InWorld) == false {
		return
	}

	s.GuardSession.Lock()
	defer s.GuardSession.Unlock()
	var index int = -1
	for i, element := range s.CharacterProps {
		if element.PropID == id {
			index = i
		}
	}

	if index == -1 {
		return
	}

	// Remove the element at index  from s.CharacterProps.
	s.CharacterProps[index] = s.CharacterProps[len(s.CharacterProps)-1] // Copy last element to index.
	s.CharacterProps[len(s.CharacterProps)-1] = models.CharacterProp{}  // Erase last element (write zero value).
	s.CharacterProps = s.CharacterProps[:len(s.CharacterProps)-1]       // Truncate slice.

	s.DB().Where("id = ?", s.PlayerID()).Where("prop = ?", id).Delete(new(models.CharacterProp))
}

package realm

import (
	"github.com/Gophercraft/core/packet/spell"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/log"
)

func (s *Session) SendActionButtons() {
	var ab []models.ActionButton
	if err := s.DB().Where("player = ?", s.PlayerID()).Find(&ab); err != nil {
		panic(err)
	}

	nab := spell.NumActionButtons(s.Build())

	if len(ab) > nab {
		ab = ab[:nab]
	}

	data := &spell.ActionButtons{}
	data.Buttons = make([]spell.ActionButton, nab)

	for _, button := range ab {
		if int(button.Button) > len(data.Buttons) {
			continue
		}
		data.Buttons[int(button.Button)] = spell.ActionButton{
			Action: button.Action,
			Type:   button.Type,
		}
	}

	s.Send(data)
}

func (s *Session) HandleSetActionButton(sab *spell.SetActionButton) {
	log.Dump("Set action button", sab)

	s.DB().Where("player = ?", s.PlayerID()).Where("button = ?", sab.Slot).Delete(new(models.ActionButton))

	if sab.Action != 0 {
		s.DB().Insert(&models.ActionButton{
			Player: s.PlayerID(),
			Button: uint8(sab.Slot),
			Action: sab.Action,
			Type:   sab.Type,
		})
	}
}

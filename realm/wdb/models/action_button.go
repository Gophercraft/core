package models

import "fmt"

type ActionType uint8

const (
	ActionSpell ActionType = iota
	ActionMacro
	ActionItem
)

func (at ActionType) String() string {
	if str, err := at.EncodeWord(); err == nil {
		return str
	}

	return fmt.Sprintf("ActionType(%d)", at)
}

func (at *ActionType) EncodeWord() (string, error) {
	switch *at {
	case ActionSpell:
		return "Spell", nil
	case ActionMacro:
		return "Macro", nil
	case ActionItem:
		return "Item", nil
	default:
		return "", fmt.Errorf("spell: invalid ActionType in EncodeWord %d", *at)
	}
}

func (at *ActionType) DecodeWord(data string) error {
	switch data {
	case "Spell":
		*at = ActionSpell
	case "Macro":
		*at = ActionMacro
	case "Item":
		*at = ActionItem
	default:
		return fmt.Errorf("spell: invalid string in ActionType.DecodeWord %s", data)
	}

	return nil
}

// ActionButton stores all the buttons a player has in their action bars.
type ActionButton struct {
	Player uint64 `xorm:"'player' index"`
	Button uint8
	Action uint32
	Type   ActionType
	Misc   uint8
}

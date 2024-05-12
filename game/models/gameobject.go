package models

import (
	"github.com/Gophercraft/core/i18n"
)

type GameObjectTemplate struct {
	ID             string `xorm:"'id' index"`
	Entry          uint32 `csv:"-" xorm:"'entry' bigint pk"`
	Type           uint32
	DisplayID      uint32 `xorm:"'display_id'"`
	Name           i18n.Text
	IconName       string
	CastBarCaption string
	Faction        uint32
	Flags          GameObjectFlags
	HasCustomAnim  bool
	Size           float32
	Data           []uint32
	MinGold        Money
	MaxGold        Money
}

package main

import (
	"fmt"

	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/tempest"
)

type GameTele struct {
	ID          int64   `xorm:"'entry'"`
	PositionX   float32 `xorm:"'x'"`
	PositionY   float32 `xorm:"'y'"`
	PositionZ   float32 `xorm:"'z'"`
	Orientation float32 `xorm:"'o'"`
	Map         uint32  `xorm:"'map'"`
	Name        string  `xorm:"'name'"`
}

func (gt GameTele) TableName() string {
	return "worldports"
}

type AreatriggerTeleport struct {
	ID                uint32  `xorm:"'id'"`
	Name              string  `xorm:"'name'"`
	RequiredLevel     uint32  `xorm:"'required_level'"`
	RequiredItem      uint32  `xorm:"'required_item'"`
	RequiredItem2     uint32  `xorm:"'required_item2'"`
	RequiredQuestDone uint32  `xorm:"'required_quest_done'"`
	TargetMap         uint32  `xorm:"'target_map'"`
	TargetPositionX   float32 `xorm:"'target_position_x'"`
	TargetPositionY   float32 `xorm:"'target_position_y'"`
	TargetPositionZ   float32 `xorm:"'target_position_z'"`
	TargetPositionO   float32 `xorm:"'target_orientation'"`
}

func extractGameTeleports() {
	var gt []GameTele
	err := DB.Find(&gt)
	if err != nil {
		panic(err)
	}

	fl := openFile("DB/PortLocation.txt")
	printTimestamp(fl)
	wr := openTextWriter(fl)

	for _, pl := range gt {
		if err := wr.Encode(&models.PortLocation{
			ID: pl.Name,
			Location: tempest.C4Vector{
				X: pl.PositionX,
				Y: pl.PositionY,
				Z: pl.PositionZ,
				W: pl.Orientation,
			},
			Map: pl.Map,
		}); err != nil {
			panic(fmt.Sprintf("error in encode %s", err))
		}
	}

	fl.Close()
}

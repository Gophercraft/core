package main

import (
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/tempest"
)

type GameTele struct {
	ID          int64   `xorm:"'id'"`
	PositionX   float32 `xorm:"'position_x'"`
	PositionY   float32 `xorm:"'position_y'"`
	PositionZ   float32 `xorm:"'position_z'"`
	Orientation float32 `xorm:"'orientation'"`
	Map         uint32  `xorm:"'map'"`
	Name        string  `xorm:"'name'"`
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
		if err := wr.Encode(models.PortLocation{
			ID: pl.Name,
			Location: tempest.C4Vector{
				X: pl.PositionX,
				Y: pl.PositionY,
				Z: pl.PositionZ,
				W: pl.Orientation,
			},
			Map: pl.Map,
		}); err != nil {
			panic(err)
		}
	}

	fl.Close()
}

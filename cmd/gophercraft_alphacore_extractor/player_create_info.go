package main

import (
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/text"
)

type PlayerCreateInfo struct {
	Race  uint8   `xorm:"'race'"`
	Class uint8   `xorm:"'class'"`
	Map   uint32  `xorm:"'map'"`
	Zone  uint32  `xorm:"'zone'"`
	X     float32 `xorm:"'position_x'"`
	Y     float32 `xorm:"'position_y'"`
	Z     float32 `xorm:"'position_z'"`
	O     float32 `xorm:"'orientation'"`
}

func (pci PlayerCreateInfo) TableName() string {
	return "Playercreateinfo"
}

func extractPlayerCreateInfo() {
	fl := openFile("DB/PlayerCreateInfo.txt")
	wr := text.NewEncoder(fl)

	var pCreateInfo []PlayerCreateInfo

	err := DB.Find(&pCreateInfo)
	if err != nil {
		panic(err)
	}

	for _, pci := range pCreateInfo {
		mod := &models.PlayerCreateInfo{
			Race:  models.Race(pci.Race),
			Class: models.Class(pci.Class),
			Position: tempest.C4Vector{
				pci.X,
				pci.Y,
				pci.Z,
				pci.O,
			},
			Map:  pci.Map,
			Zone: pci.Zone,
		}

		if err := wr.Encode(mod); err != nil {
			panic(err)
		}
	}

	fl.Close()
}

package main

import (
	"fmt"
	"reflect"

	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/realm/wdb/models"
)

type GameobjectTemplate struct {
	Entry     uint32 `xorm:"'entry'"`
	Type      uint32 `xorm:"'type'"`
	DisplayID uint32 `xorm:"'displayId'"`
	Name      string `xorm:"'name'"`
	// IconName       string `xorm:"'IconName'"`
	// CastBarCaption string `xorm:"'castBarCaption'"`
	Faction uint32 `xorm:"'faction'"`
	Flags   uint32 `xorm:"'flags'"`
	// ExtraFlags  uint32  `xorm:"'ExtraFlags'"`
	Size  float32 `xorm:"'size'"`
	Data0 uint32  `xorm:"'data0'"`
	Data1 uint32  `xorm:"'data1'"`
	Data2 uint32  `xorm:"'data2'"`
	Data3 uint32  `xorm:"'data3'"`
	Data4 uint32  `xorm:"'data4'"`
	Data5 uint32  `xorm:"'data5'"`
	Data6 uint32  `xorm:"'data6'"`
	Data7 uint32  `xorm:"'data7'"`
	Data8 uint32  `xorm:"'data8'"`
	Data9 uint32  `xorm:"'data9'"`
	// CustomData1 uint32  `xorm:"'CustomData1'"`
	Mingold    int64  `xorm:"'mingold'"`
	Maxgold    int64  `xorm:"'maxgold'"`
	ScriptName string `xorm:"'script_name'"`
}

func (GameobjectTemplate) TableName() string {
	return "gameobject_template"
}

func extractGameObjects() {
	var gtt []GameobjectTemplate
	err := DB.Find(&gtt)
	if err != nil {
		panic(err)
	}

	gfl := openFile("DB/GameObjectTemplate.txt")
	printTimestamp(gfl)
	wr := openTextWriter(gfl)

	for _, v := range gtt {
		data := make([]uint32, 10)
		d := reflect.ValueOf(v)
		for x := 0; x < 10; x++ {
			data[x] = uint32(d.FieldByName(fmt.Sprintf("Data%d", x)).Uint())
		}

		if err := wr.Encode(models.GameObjectTemplate{
			ID:        fmt.Sprintf("go:%d", v.Entry),
			Type:      v.Type,
			DisplayID: v.DisplayID,
			Name:      i18n.GetEnglish(v.Name),
			// IconName:       v.IconName,
			// CastBarCaption: v.CastBarCaption,
			Faction: v.Faction,
			Flags:   models.GameObjectFlags(v.Flags),
			// HasCustomAnim: v.ExtraFlags == 1,
			Size:    v.Size,
			Data:    data,
			MinGold: models.Money(v.Mingold),
			MaxGold: models.Money(v.Maxgold),
		}); err != nil {
			panic(err)
		}
	}

	gfl.Close()

	gtt = nil
}

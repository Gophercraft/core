package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	DB *xorm.Engine
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(os.Args[0], "<xorm url>")
		return
	}

	var err error
	DB, err = xorm.NewEngine("mysql", os.Args[1])
	if err != nil {
		panic(err)
	}

	extractGameTeleports()
	extractItems()
	extractCreatures()
	extractNPCText()
	extractGameObjects()
	extractPlayerCreateInfo()
	extractPlayerClassLevelStats()
	extractPlayerLevelStats()
	extractExpForLevel()
	// var fl *os.File
	// var wr *text.Encoder

	// Disabled due to containing incorrect results.
	// var pca []PlayerCreateActionButton
	// err = DB.Find(&pca)
	// if err != nil {
	// 	panic(err)
	// }
	// fl := openFile("DB/PlayerCreateActionButton.txt")
	// printTimestamp(fl)
	// wr := openTextWriter(fl)
	// for _, pcab := range pca {
	// 	if err := wr.Encode(wdb.PlayerCreateActionButton{
	// 		Race:   int8(pcab.Race),
	// 		Class:  int8(pcab.Class),
	// 		Button: pcab.Button,
	// 		Action: pcab.Action,
	// 		Type:   pcab.Type,
	// 	}); err != nil {
	// 		panic(err)
	// 	}
	// }

	// fl.Close()

	// var att []AreatriggerTeleport

	// err = DB.Find(&att)
	// if err != nil {
	// 	panic(err)
	// }

	// fl = openFile("Scripts/AreaTriggers.lua")
	// for _, v := range att {
	// 	fmt.Fprintf(fl, "-- %s\n", v.Name)
	// 	fmt.Fprintf(fl, "OnAreaTrigger(%d, function(plyr)\n", v.ID)
	// 	if v.RequiredItem != 0 {
	// 		addItemReq(fl, v.RequiredItem)
	// 	}

	// 	if v.RequiredItem2 != 0 {
	// 		addItemReq(fl, v.RequiredItem2)
	// 	}

	// 	if v.RequiredLevel != 0 {
	// 		fmt.Fprintf(fl, "  if plyr:GetLevel() < %d then\n", v.RequiredLevel)
	// 		fmt.Fprintf(fl, "    plyr:SendRequiredLevelZoneError(%d)\n", v.RequiredLevel)
	// 		fmt.Fprintf(fl, "    return\n")
	// 		fmt.Fprintf(fl, "  end\n\n")
	// 	}

	// 	if v.RequiredQuestDone != 0 {
	// 		fmt.Fprintf(fl, "  if not plyr:QuestDone(%d) then\n", v.RequiredQuestDone)
	// 		fmt.Fprintf(fl, "    plyr:SendRequiredQuestZoneError(%d)\n", v.RequiredQuestDone)
	// 		fmt.Fprintf(fl, "    return\n")
	// 		fmt.Fprintf(fl, "  end\n\n")
	// 	}

	// 	fmt.Fprintf(fl, "  plyr:Teleport(%d, %f, %f, %f, %f)\n", v.TargetMap, v.TargetPositionX, v.TargetPositionY, v.TargetPositionZ, v.TargetPositionO)
	// 	fmt.Fprintf(fl, "end)\n\n")
	// }

	// fl.Close()

}

func addItemReq(fl *os.File, item uint32) {
	fmt.Fprintf(fl, "  if not plyr:HasItem(\"it:%d\") then\n", item)
	fmt.Fprintf(fl, "    plyr:SendRequiredItemZoneError(\"it:%d\")\n", item)
	fmt.Fprintf(fl, "    return\n")
	fmt.Fprintf(fl, "  end\n\n")
}

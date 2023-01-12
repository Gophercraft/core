package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
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

	DB.Logger().SetLevel(0xff)

	DB.SetLogLevel(log.LOG_ERR)
	DB.SetLogger(log.NewSimpleLogger(os.Stdout))

	extractGameTeleports()
	extractItems()
	extractCreatures()
	extractNPCText()
	extractGameObjects()
	extractPlayerCreateActionButtons()
	extractPlayerCreateInfo()
	extractPlayerCreateItems()
	extractPlayerCreateSpells()
	extractPlayerClassLevelStats()
	extractPlayerLevelStats()
}

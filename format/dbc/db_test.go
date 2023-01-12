package dbc

import (
	"bytes"
	"os"
	"testing"

	"github.com/Gophercraft/core/format/content"
	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/log"
)

func TestDB(t *testing.T) {
	vol, err := content.Open("E:\\Gaymes\\World of Warcraft 1.12")
	if err != nil {
		panic(err)
	}

	data, err := vol.ReadFile("DBFilesClient\\AreaTable.dbc")
	if err != nil {
		panic(err)
	}

	os.WriteFile("E:\\Gaymes\\World of Warcraft 1.12\\AreaTable.dbc", data, 0700)

	db := NewDB(vol.Build())
	table, err := db.Open("AreaTable", bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	if err := table.Range(func(cursor *dbdefs.Ent_AreaTable) bool {
		log.Dump("cursor", cursor)
		return true
	}); err != nil {
		panic(err)
	}
}

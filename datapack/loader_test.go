package datapack

import (
	"testing"

	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/superp00t/etc"
)

type celebRecord struct {
	Name     string
	KnownFor []string
}

func TestLoader(t *testing.T) {
	ld, err := Open(etc.Import("github.com/Gophercraft/core/datapack/testdata").Render())
	if err != nil {
		t.Fatal(err)
	}

	var pl []models.PortLocation
	ld.ReadAll("DB/PortLocation.csv", &pl)

	log.Dump("pl", pl)
	ld.Close()
}

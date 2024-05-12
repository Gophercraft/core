package dbc

import (
	"fmt"

	"github.com/Gophercraft/core/format/dbc/dbd"
	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/version"
)

type DB struct {
	Build version.Build
	Table []*Table
}

func NewDB(v version.Build) *DB {
	return &DB{
		Build: v,
	}
}

func (db *DB) lookupDef(name string) (*dbd.Definition, error) {
	return dbdefs.Lookup(name)
}

func (db *DB) detectLayout(table *Table) error {
	def, err := db.lookupDef(table.Name)
	if err != nil {
		return err
	}

	table.Definition = def

	for i := range def.Layouts {
		layout := &def.Layouts[i]
		// Search for exact builds
		for _, exact := range layout.VerifiedBuilds {
			if exact == db.Build {
				table.Layout = layout
				return nil
			}
		}
		// Failing that, see if there is a match within build ranges
		for _, rng := range layout.BuildRanges {
			if rng.Contains(db.Build) {
				table.Layout = layout
				return nil
			}
		}
	}

	return fmt.Errorf("dbc: table found, but could not find layouts matching build %s", db.Build)
}

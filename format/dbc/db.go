package dbc

import (
	"fmt"

	"github.com/Gophercraft/core/format/dbc/dbd"
	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/vsn"
)

type DB struct {
	Build vsn.Build
	Table []*Table
}

func NewDB(v vsn.Build) *DB {
	return &DB{
		Build: v,
	}
}

func (db *DB) lookupDef(name string) (*dbd.Definition, error) {
	alldefs := dbdefs.All

	// // Fast binary search
	// i := sort.Search(len(alldefs), func(i int) bool {
	// 	return alldefs[i].Name >= name
	// })

	// if i < len(alldefs) && alldefs[i].Name == name {
	// 	return &alldefs[i], nil
	// }

	// // could be that defs are not sorted, slooow search
	// for i := range alldefs {
	// 	if alldefs[i].Name == name {
	// 		return &alldefs[i], nil
	// 	}
	// }

	def, ok := alldefs[name]
	if !ok {
		return nil, fmt.Errorf("dbc: no such definition for table %s", name)
	}

	return def, nil
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

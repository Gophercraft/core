package wizard

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/Gophercraft/core/datapack"
	"github.com/Gophercraft/core/format/dbc"
	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/text"
)

const DBPackName = "!db.zip"

type ExDb struct {
	Name string
	Type reflect.Type
}

// Return the list of DB files required to produce a base datapack
func (ex *Extractor) neededDBs() []ExDb {
	var need []ExDb

	need = []ExDb{
		{"AreaTable", reflect.TypeFor[dbdefs.Ent_AreaTable]()},
		{"AreaTrigger", reflect.TypeFor[dbdefs.Ent_AreaTrigger]()},
		{"CharBaseInfo", reflect.TypeFor[dbdefs.Ent_CharBaseInfo]()},
		{"ChrRaces", reflect.TypeFor[dbdefs.Ent_ChrRaces]()},
		{"ChrClasses", reflect.TypeFor[dbdefs.Ent_ChrClasses]()},
		{"CharStartOutfit", reflect.TypeFor[dbdefs.Ent_CharStartOutfit]()},
		{"CreatureFamily", reflect.TypeFor[dbdefs.Ent_CreatureFamily]()},
		{"EmotesText", reflect.TypeFor[dbdefs.Ent_EmotesText]()},
		{"ItemClass", reflect.TypeFor[dbdefs.Ent_ItemClass]()},
		{"Map", reflect.TypeFor[dbdefs.Ent_Map]()},
		{"Spell", reflect.TypeFor[dbdefs.Ent_Spell]()},
		{"SpellCastTimes", reflect.TypeFor[dbdefs.Ent_SpellCastTimes]()},
		{"SpellDuration", reflect.TypeFor[dbdefs.Ent_SpellDuration]()},
		{"SpellVisual", reflect.TypeFor[dbdefs.Ent_SpellVisual]()},
	}

	return need
}

func get_extractor_author() string {
	exAuthor := fmt.Sprintf("Gophercraft core wizard v%s", version.GophercraftVersion)
	return exAuthor
}

func (ex *Extractor) ExtractDatabases() error {
	const tempPackDir = "x-db"

	if ex.packExists(DBPackName) {
		ex.removePack(DBPackName)
	}

	pack, err := ex.AuthorPack(tempPackDir, &datapack.PackInfo{
		Name: "Gophercraft Base Databases",
		Authors: []string{
			fmt.Sprintf("Generated by %s", get_extractor_author()),
		},
		Version: 0,
		Base:    true,
		Description: strings.Join([]string{
			"These are base databases required by Gophercraft core.",
			"These are essential for interfacing with the WoW client.",
			ex.generationNotice(),
		}, "\n"),
		MinimumCoreVersion: version.GophercraftVersion.String(),
		SupportedClients: []version.BuildRange{
			// Unary range
			version.Range(ex.Source.Build(), ex.Source.Build()),
		},
	})

	if err != nil {
		return err
	}

	need := ex.neededDBs()

	pb := log.NewIntProgressBar("Extracting databases...", 0, int64(len(need)))
	log.StartProgressBar(pb)

	for i := range need {
		err := ex.extractDB(pack, &need[i])
		if err != nil {
			return err
		}
		pb.SetInt(int64(i))
	}

	pb.Complete()

	if err := pack.Create(filepath.Join(ex.Dir, DBPackName)); err != nil {
		ex.removePack(tempPackDir)
		return err
	}

	return ex.removePack(tempPackDir)
}

func (ex *Extractor) extractDB(pack *datapack.Creator, exDB *ExDb) error {
	prefix := "DBFilesClient\\"
	suffix := ".dbc"

	path := prefix + exDB.Name + suffix

	dbfile, err := ex.Source.Open(path)
	if err != nil {
		return err
	}

	db := dbc.NewDB(ex.Source.Build())
	table, err := db.Open(exDB.Name, dbfile)
	if err != nil {
		return err
	}

	numRecords := int(table.Header.RecordCount)

	file, err := pack.Open(fmt.Sprintf("DB/Client/%s.txt", exDB.Name))
	if err != nil {
		return err
	}

	encoder := text.NewEncoder(file)
	encoder.Tabular = true
	encoder.Indent = " "

	zero := reflect.New(exDB.Type)
	cursor := reflect.New(exDB.Type)

	progress := log.NewIntProgressBar(
		fmt.Sprintf("Loading %s", exDB.Name),
		0, int64(numRecords),
	)

	log.StartProgressBar(progress)

	for i := 0; i < numRecords; i++ {
		// set cursor to zero value
		cursor.Elem().Set(zero.Elem())

		// read dbc file into struct
		if err := table.Index(i, cursor.Interface()); err != nil {
			return err
		}

		// encode struct to text file
		if err := encoder.Encode(cursor.Interface()); err != nil {
			return err
		}

		progress.SetInt(int64(i))
	}

	progress.Complete()

	return file.Close()
}

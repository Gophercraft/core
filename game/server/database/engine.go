package database

import (
	"errors"
	"io"
	"reflect"

	"github.com/Gophercraft/core/datapack"
	"github.com/Gophercraft/phylactery/database"
)

type Engine struct {
	// the long term world database
	world_db_container *database.Container
	//
	cache_db_container *database.Container
	//
	datapack_loader *datapack.Loader
}

type EngineConfig struct {
	//
	DatapacksDirectory string

	// the Phylactery database engine to use
	WorldDatabaseEngine   string
	WorldDatabaseLocation string

	// the Phylactery database engine to use
	CacheDatabaseEngine   string
	CacheDatabaseLocation string
}

func Open(engine_config *EngineConfig) (engine *Engine, err error) {
	world_db_engine := engine_config.WorldDatabaseEngine
	cache_db_engine := engine_config.CacheDatabaseEngine
	if engine_config.WorldDatabaseLocation == "" {
		engine_config.WorldDatabaseLocation = "world_db"
	}
	if engine_config.CacheDatabaseLocation == "" {
		engine_config.CacheDatabaseLocation = "cache_db"
	}

	engine = new(Engine)

	engine.world_db_container, err = database.Open(engine_config.WorldDatabaseLocation, database.WithEngine(world_db_engine))
	if err != nil {
		return
	}

	engine.cache_db_container, err = database.Open(engine_config.CacheDatabaseLocation, database.WithEngine(cache_db_engine))
	if err != nil {
		return
	}

	if engine.datapack_loader, err = datapack.Load(engine_config.DatapacksDirectory); err != nil {
		return
	}

	return
}

func (engine *Engine) compile_table(table_name string, entry *model_entry) (err error) {
	var text_db_loader *datapack.TextDatabaseLoader
	text_db_loader, err = datapack.NewTextDatabaseLoader(engine.datapack_loader, table_name)
	if err != nil {
		return
	}

	table := engine.cache_db_container.Table(entry.name)

	cursor_value := reflect.New(entry.schema)

	for {
		err = text_db_loader.Load(cursor_value.Interface())
		if errors.Is(err, io.EOF) {
			err = nil
			return
		}

		err = table.Insert(cursor_value.Interface())
		if err != nil {
			return
		}
	}
}

func (engine *Engine) CompileDatapacks() (err error) {
	// Compile client databases
	for _, dbc_table := range dbc_model_tables {
		err = engine.compile_table("Client/"+dbc_table.name, &dbc_table)
		if err != nil {
			return
		}
	}

	// Compile static information tables 

	for _, static_info_table := range static_info_tables {
		err = engine.compile_table("Static/" + static_info_table.name,
	}

	return
}

package database

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/Gophercraft/core/datapack"
)

func (engine *Engine) compile_table(table_path string, entry *model_entry) (err error) {
	var text_db_loader *datapack.TextDatabaseLoader
	text_db_loader, err = datapack.NewTextDatabaseLoader(engine.datapack_loader, table_path)
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

func find_template_ID(value reflect.Value) (template_ID string) {
	field := value.FieldByName("ID")
	if !field.IsValid() {
		panic(fmt.Errorf("TYPE %s HAS NO ID", value.Type()))
	}
	if field.Type() != reflect.TypeFor[string]() {
		panic(fmt.Errorf("TYPE %s ENTRY IS NOT STRING", value.Type()))
	}

	template_ID = field.String()
	return
}

func set_template_entry(value reflect.Value, entry int32) {
	field := value.FieldByName("Entry")
	if !field.IsValid() {
		panic(fmt.Errorf("TYPE %s HAS NO ENTRY", value.Type()))
	}
	if field.Type() != reflect.TypeFor[int32]() {
		panic(fmt.Errorf("TYPE %s ENTRY IS NOT INT32", value.Type()))
	}
	field.SetInt(int64(entry))
}

// Gophercraft datapacks should use strings when identifying object templates
func (engine *Engine) compile_object_template_table(table_path string, entry *model_template_entry) (err error) {
	var text_db_loader *datapack.TextDatabaseLoader
	text_db_loader, err = datapack.NewTextDatabaseLoader(engine.datapack_loader, table_path)
	if err != nil {
		return
	}

	//
	base_entry := int32(0x40000000)

	entry_counter := base_entry

	table := engine.cache_db_container.Table(entry.name)

	cursor_value := reflect.New(entry.schema)

	// game:item:1234
	// will result in the entry being 1234 instead of a random counter > 0x40000000
	game_prefix := fmt.Sprintf("game:%s:", entry.type_id)

	for {
		err = text_db_loader.Load(cursor_value.Interface())
		if errors.Is(err, io.EOF) {
			err = nil
			return
		}

		// An ID is either something like
		// custom_datapack:some_kind_of_item
		// A placeholder value, for the server to assign a code deterministically from
		// or it is a game:item:XXXX
		//

		template_ID := find_template_ID(cursor_value.Elem())
		var entry_number int32

		if strings.HasPrefix(template_ID, game_prefix) {
			//
			var entry_i64 int64
			entry_string := template_ID[strings.LastIndex(template_ID, ":")+1:]
			entry_i64, err = strconv.ParseInt(entry_string, 10, 32)
			if err != nil {
				panic(fmt.Errorf("cannot parse"))
			}
			entry_number = int32(entry_i64)
		} else {
			entry_counter++
			entry_number = entry_counter
		}

		set_template_entry(cursor_value.Elem(), entry_number)

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
		err = engine.compile_table("Static/"+static_info_table.name, &static_info_table)
		if err != nil {
			return
		}
	}

	// Compile object templates
	for _, object_template_table := range object_template_tables {
		err = engine.compile_object_template_table("Object/"+object_template_table.name, &object_template_table)
		if err != nil {
			return
		}
	}

	return
}

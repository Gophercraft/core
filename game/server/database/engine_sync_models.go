package database

import (
	"fmt"
	"reflect"

	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/game/models"
)

var (
	dbc_model_tables = []model_entry{
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
		{"SpellDuration", reflect.TypeFor[dbdefs.Ent_SpellDuration]()},
		{"SpellCastTimes", reflect.TypeFor[dbdefs.Ent_SpellCastTimes]()},
		{"SpellVisual", reflect.TypeFor[dbdefs.Ent_SpellVisual]()},
	}

	player_info_tables = []model_entry{
		{"Character", reflect.TypeFor[models.Character]()},
		{"LearnedAbility", reflect.TypeFor[models.LearnedAbility]()},
		{"ActionButton", reflect.TypeFor[models.ActionButton]()},
		{"Item", reflect.TypeFor[models.Item]()},
		{"Inventory", reflect.TypeFor[models.Inventory]()},
		{"Contact", reflect.TypeFor[models.Contact]()},
		{"ExploredZone", reflect.TypeFor[models.ExploredZone]()},
		{"FactionStanding", reflect.TypeFor[models.FactionStanding]()},
		{"CharacterAchievement", reflect.TypeFor[models.CharacterAchievement]()},
	}

	static_info_tables = []model_entry{
		{"PlayerCreateInfo", reflect.TypeFor[models.PlayerCreateInfo]()},
		{"PlayerCreateAbilities", reflect.TypeFor[models.PlayerCreateAbilities]()},
		{"PlayerCreateActionButtons", reflect.TypeFor[models.PlayerCreateActionButtons]()},
		{"PlayerCreateItem", reflect.TypeFor[models.PlayerCreateItem]()},
		{"LevelExperience", reflect.TypeFor[models.LevelExperience]()},
		{"PortLocation", reflect.TypeFor[models.PortLocation]()},
		{"LocString", reflect.TypeFor[models.LocString]()},
		{"BaseClassLevelStats", reflect.TypeFor[models.BaseClassLevelStats]()},
		{"RaceStats", reflect.TypeFor[models.RaceStats]()},
		{"ClassStats", reflect.TypeFor[models.ClassStats]()},
		{"SpawnGroup", reflect.TypeFor[models.SpawnGroup]()},
	}

	object_template_tables = []model_template_entry{
		{"GameObjectTemplate", "gameobject", reflect.TypeFor[models.GameObjectTemplate]()},
		{"ItemTemplate", "item", reflect.TypeFor[models.ItemTemplate]()},
		{"CreatureTemplate", "creature", reflect.TypeFor[models.CreatureTemplate]()},
		{"NPCText", "npc_text", reflect.TypeFor[models.NPCText]()},
	}
)

type model_entry struct {
	name   string
	schema reflect.Type
}

type model_template_entry struct {
	name    string
	type_id string
	schema  reflect.Type
}

func (engine *Engine) sync_world_models() (err error) {
	for _, player_info_table := range player_info_tables {
		if err = engine.world_db_container.Table(player_info_table.name).SyncType(player_info_table.schema); err != nil {
			err = fmt.Errorf("database: error syncing table %s: %w", player_info_table.name, err)
			return
		}
	}

	return
}

func (engine *Engine) sync_cache_models() (err error) {

	for _, dbc_model_table := range dbc_model_tables {
		if err = engine.world_db_container.Table(dbc_model_table.name).SyncType(dbc_model_table.schema); err != nil {
			err = fmt.Errorf("database: error syncing table %s: %w", dbc_model_table.name, err)
			return
		}
	}

	for _, static_info_table := range static_info_tables {
		if err = engine.world_db_container.Table(static_info_table.name).SyncType(static_info_table.schema); err != nil {
			err = fmt.Errorf("database: error syncing table %s: %w", static_info_table.name, err)
			return
		}
	}

	return
}

func (engine *Engine) sync_models() (err error) {
	if err = engine.sync_cache_models(); err != nil {
		return
	}

	if err = engine.sync_world_models(); err != nil {
		return
	}

	return
}

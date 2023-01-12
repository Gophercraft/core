package realm

import (
	"fmt"
	"io"
	"path/filepath"
	"reflect"
	"sort"

	"strconv"
	"strings"

	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/log"
)

// It seems unlikely that the files will ever reach this super-high threshold.
// Thus, it makes an ideal starting point for custom templates.
func getHighCounter() uint32 {
	// begins at 0x1000000 (16777216)
	return uint32(0) | (1 << 24)
}

func findKnownID(id string, registry []models.ObjectTemplateRegistry) (bool, uint32) {
	for _, reg := range registry {
		if reg.ID == id {
			return true, reg.Entry
		}
	}

	index := sort.Search(len(registry), func(i int) bool {
		return registry[i].ID >= id
	})

	if index == len(registry) {
		return false, 0
	}

	reg := &registry[index]
	if reg.ID == id {
		return true, reg.Entry
	}

	return false, 0
}

// LoadObjectTemplates (item/creature/gameobject templates)
// Most pertinent to object templates that are queried by the client, i.e. structs that have an Entry field
func (ws *Server) LoadObjectTemplates(name string, typeID guid.TypeID, typeOf reflect.Type) error {
	ws.DB.Println(name, typeID, typeOf)

	ws.DB.Cache.Clear(typeOf)

	var highCounter uint32
	highCounter = getHighCounter()
	var startCounter = highCounter

	// The server keeps a list of ID-entry associations (ObjectTemplateRegistry)
	// This allows you to repeatedly add and remove custom object templates without always needing to refresh the client cache
	var knownIDs []models.ObjectTemplateRegistry
	var newIDs []models.ObjectTemplateRegistry
	ws.DB.Where("type = ?", typeID).Find(&knownIDs)

	log.Dump("Known IDs - before sort", knownIDs)

	sort.Slice(knownIDs, func(i, j int) bool {
		return knownIDs[i].ID < knownIDs[j].ID
	})

	log.Dump("Known IDs - after sort", knownIDs)

	for _, id := range knownIDs {
		// offset counter to largest in registry. This prevents collisions from occuring.
		if id.Entry >= startCounter {
			startCounter = id.Entry + 1
		}
	}

	txtLoader, err := ws.PackLoader.NewTextLoader("DB", name)
	if err != nil {
		log.Fatal(err)
	}

	for {
		newObject := reflect.New(typeOf)
		object := newObject.Elem()
		err := txtLoader.Scan(newObject.Interface())
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		id := wdb.GetID(object)
		if id == "" {
			log.Println("No id in object", name, "of type", typeOf)
			continue
		}
		var entry uint32
		foundId, foundEntry := findKnownID(id, knownIDs)
		if foundId {
			entry = foundEntry
		} else {
			str := strings.Split(id, ":")
			if len(str) >= 2 {
				num, err := strconv.ParseUint(str[1], 10, 32)
				if err == nil {
					// If this ID contains a number, use this as the entry code.
					entry = uint32(num)
				} else {
					entry = startCounter
					startCounter++
					newIDs = append(newIDs, models.ObjectTemplateRegistry{
						ID:    id,
						Type:  typeID,
						Entry: entry,
					})
				}
			} else {
				entry = startCounter
				startCounter++
				newIDs = append(newIDs, models.ObjectTemplateRegistry{
					ID:    id,
					Type:  typeID,
					Entry: entry,
				})
			}
		}

		wdb.SetEntry(object, entry)
		ws.DB.Cache.Store(newObject)
	}

	if len(newIDs) == 0 {
		return nil
	}

	_, err = ws.DB.Insert(&newIDs)
	return err
}

// Read game data as a unified source. The BASE data can be supplied by the game itself.
// DBC edits can be supplied in text format
func (ws *Server) LoadServerDB(name string, typeOf reflect.Type) error {
	packName := "Ent_" + name
	ws.DB.Println(packName)
	// Wipe previous DB entries.
	ws.DB.Cache.Clear(typeOf)

	// Start reading from base files.
	// dbPath := ws.dbPath(name)

	// overrideBase := false

	// for _, pack := range ws.PackLoader.Volumes {
	// 	for _, ot := range pack.OverrideTables {
	// 		if ot == packName {
	// 			overrideBase = true
	// 			break
	// 		}
	// 	}
	// }

	// if ws.ContentSource != nil && !overrideBase {
	// 	baseContentFile, err := ws.ContentSource.ReadFile(dbPath)
	// 	if err == nil {
	// 		baseContentTable, err := ws.DBC.Open(name, bytes.NewReader(baseContentFile))
	// 		if err != nil {
	// 			panic(err)
	// 		}

	// 		for i := 0; i < baseContentTable.Len(); i++ {
	// 			rec := reflect.New(typeOf)
	// 			err := baseContentTable.Index(i, rec.Interface())
	// 			if err != nil {
	// 				return err
	// 			}
	// 			ws.DB.Cache.Store(rec.Elem())
	// 		}
	// 	}
	// }

	txtLoader, err := ws.PackLoader.NewTextLoader("DB", packName)
	if err != nil {
		return err
	}

	for {
		rec := reflect.New(typeOf)
		err := txtLoader.Scan(rec.Interface())
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		ws.DB.Cache.Store(rec)
	}
}

func (ws *Server) LoadPlayerCreateItems() error {
	log.Println("Loading player create items")

	overrideBase := false

	for _, pack := range ws.PackLoader.Volumes {
		for _, ot := range pack.OverrideTables {
			if ot == "Ent_CharStartOutfit" {
				overrideBase = true
				break
			}
		}
	}

	if !overrideBase {
		// These files can be converted and recombined into the more generic PlayerCreateItem
		dbcLoader, err := ws.PackLoader.NewTextLoader("DB", "Ent_CharStartOutfit")
		if err == nil {
			for {
				var cso dbdefs.Ent_CharStartOutfit
				if err = dbcLoader.Scan(&cso); err != nil {
					return err
				}

				// wpn := 0

				for i, itemID := range cso.ItemID {
					if itemID > 0 {
						pci := &models.PlayerCreateItem{}
						pci.Equip = models.EquipPaperDoll
						pci.Item = fmt.Sprintf("it:%d", itemID)
						pci.Class.Set(models.Class(cso.ClassID), true)
						pci.Race.Set(models.Race(cso.RaceID), true)

						iType := models.InventoryType(cso.InventoryType[i])
						if iType != models.IT_Unequippable {
							ok := false

							// if iType == models.IT_Weapon {
							// 	pci.Slot = models.PaperDoll_MainHand + models.ItemSlot(wpn)
							// 	wpn++
							// 	ok = true
							// } else {
							// 	pci.Slot, ok = models.ItemDisplaySlots[iType]
							// }

							if ok {
								ws.DB.Cache.Store(reflect.ValueOf(pci))
							} else {
								panic(cso.InventoryType[i])
							}
						}
					}
				}
			}
		}
	}

	txtLoader, err := ws.PackLoader.NewTextLoader("DB", "PlayerCreateItem")
	if err != nil {
		return err
	}

	for {
		pci := new(models.PlayerCreateItem)
		err := txtLoader.Scan(pci)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ws.DB.Cache.Store(reflect.ValueOf(pci))
	}
	return nil
}

func (ws *Server) LoadStaticInfo(name string, typeOf reflect.Type) error {
	ws.DB.Println("Loading static fields from", name)

	txtLoader, err := ws.PackLoader.NewTextLoader("DB", name)
	if err != nil {
		return err
	}

	for {
		staticInfo := reflect.New(typeOf)
		err := txtLoader.Scan(staticInfo.Interface())
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ws.DB.Cache.Store(staticInfo)
	}
	return nil
}

// LoadDatapacks loops through all the text ZIP files owned by this Worldserver, refreshing the server database cache
func (ws *Server) LoadDatapacks() error {
	if err := ws.DB.InitCache(filepath.Join(ws.Config.Dir, "Cache")); err != nil {
		return err
	}

	var serverDBs = []struct {
		Name string
		Type reflect.Type
	}{
		{"AreaTable", reflect.TypeOf(dbdefs.Ent_AreaTable{})},
		{"AreaTrigger", reflect.TypeOf(dbdefs.Ent_AreaTrigger{})},
		{"ChrRaces", reflect.TypeOf(dbdefs.Ent_ChrRaces{})},
		{"ChrClasses", reflect.TypeOf(dbdefs.Ent_ChrClasses{})},
		{"CharStartOutfit", reflect.TypeOf(dbdefs.Ent_CharStartOutfit{})},
		{"CreatureFamily", reflect.TypeOf(dbdefs.Ent_CreatureFamily{})},
		{"EmotesText", reflect.TypeOf(dbdefs.Ent_EmotesText{})},
		{"Map", reflect.TypeOf(dbdefs.Ent_Map{})},
		{"Spell", reflect.TypeOf(dbdefs.Ent_Spell{})},
		{"SpellDuration", reflect.TypeOf(dbdefs.Ent_SpellDuration{})},
		{"SpellCastTimes", reflect.TypeOf(dbdefs.Ent_SpellCastTimes{})},
		{"SpellVisual", reflect.TypeOf(dbdefs.Ent_SpellVisual{})},
	}

	var objectTemplates = []struct {
		Table  string
		TypeID guid.TypeID
		Type   reflect.Type
	}{
		{"GameObjectTemplate", guid.TypeGameObject, reflect.TypeOf(models.GameObjectTemplate{})},
		{"ItemTemplate", guid.TypeItem, reflect.TypeOf(models.ItemTemplate{})},
		{"CreatureTemplate", guid.TypeUnit, reflect.TypeOf(models.CreatureTemplate{})},
		{"NPCText", guid.TypeNPCText, reflect.TypeOf(models.NPCText{})},
	}

	var staticInfo = []struct {
		Name string
		Type reflect.Type
	}{
		{"PlayerCreateInfo", reflect.TypeOf(models.PlayerCreateInfo{})},
		{"PlayerCreateAbility", reflect.TypeOf(models.PlayerCreateAbility{})},
		{"PlayerCreateActionButton", reflect.TypeOf(models.PlayerCreateActionButton{})},
		{"LevelExperience", reflect.TypeOf(models.LevelExperience{})},
		{"PortLocation", reflect.TypeOf(models.PortLocation{})},
		{"LocString", reflect.TypeOf(models.LocString{})},
		{"ClassLevelStats", reflect.TypeOf(models.ClassLevelStats{})},
		{"RaceClassLevelStats", reflect.TypeOf(models.RaceClassLevelStats{})},
		{"SpawnGroup", reflect.TypeOf(models.SpawnGroup{})},
	}

	numDBs := len(serverDBs) + len(objectTemplates) + len(staticInfo) + 1
	prog := log.NewIntProgressBar("Loading databases", 0, int64(numDBs))
	log.StartProgressBar(prog)
	var c int64

	for _, dbd := range serverDBs {
		if err := ws.LoadServerDB(dbd.Name, dbd.Type); err != nil {
			log.Warn("Error reading server DB", err)
			return err
		}
		c++
		prog.SetInt(c)
	}

	// todo: replace guid typeIDs with wdb type IDS
	for _, obj := range objectTemplates {
		if err := ws.LoadObjectTemplates(obj.Table, obj.TypeID, obj.Type); err != nil {
			return err
		}
		c++
		prog.SetInt(c)
	}

	for _, stat := range staticInfo {
		if err := ws.LoadStaticInfo(stat.Name, stat.Type); err != nil {
			return err
		}
		c++
		prog.SetInt(c)
	}

	ws.LoadPlayerCreateItems()
	c++
	prog.SetInt(c)
	prog.Complete()

	return nil
}

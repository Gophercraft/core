// Package wdb contains the realm databases. For server content, in-memory databases are utilized to maximize performance.
// For character-generated data, a SQL backend is used.
package wdb

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	_ "github.com/Gophercraft/core/home/dbsupport"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/log"
	"xorm.io/xorm"
)

// var DataStores = map[reflect.Type]*Store{}

type Store struct {
	store sync.Map
}

func (s *Store) Range(fn func(k, v interface{}) bool) {
	s.store.Range(fn)
}

type Core struct {
	*xorm.Engine
	Cache
}

type Tx struct {
	*xorm.Session
}

func (c *Core) AsyncTx(fn func(*Tx)) {
	go func() {
		session := c.Engine.NewSession()
		fn(&Tx{session})
		if err := session.Commit(); err != nil {
			log.Warn("Error in async transaction: ", err)
		}
	}()
}

// func (c *Core) StoreData(value reflect.Value) {
// 	if value.Kind() != reflect.Ptr {
// 		panic("not a pointer " + value.Type().String())
// 	}
// 	base := value.Elem()
// 	store := c.DataStores[base.Type()]

// 	if store == nil {
// 		store = &Store{}
// 		c.DataStores[base.Type()] = store
// 	}

// 	idField := base.FieldByName("ID")
// 	if idField.Kind() == reflect.Uint32 {
// 		store.store.Store(uint32(idField.Uint()), value.Interface())
// 	} else if idField.Kind() == reflect.String {
// 		store.store.Store(idField.String(), value.Interface())
// 	} else if idField.MethodByName("ID").IsValid() {
// 		store.store.Store(idField.MethodByName("ID").Call(nil)[0].Interface(), value.Interface())
// 	} else {
// 		panic(idField.Kind())
// 	}

// 	// Some types should be accessible by name.
// 	if idField.Kind() == reflect.Uint32 {
// 		nameField := base.FieldByName("Name")
// 		if nameField.IsValid() {
// 			store.store.Store(nameField.String(), value.Interface())
// 		}
// 	}

// 	entryField := base.FieldByName("Entry")
// 	if entryField.IsValid() {
// 		entry := uint32(base.FieldByName("Entry").Uint())
// 		store.store.Store(entry, value.Interface())
// 	}
// }

func NewCore(driver, source string) (*Core, error) {
	var err error
	cn := new(Core)
	cn.Engine, err = xorm.NewEngine(driver, source)
	if err != nil {
		return nil, err
	}

	err = cn.Engine.Sync2(
		new(models.Character),
		new(models.LearnedAbility),
		new(models.ActionButton),
		new(models.Item),
		new(models.Inventory),
		new(models.ObjectTemplateRegistry),
		new(models.Contact),
		new(models.ExploredZone),
		new(models.CharacterProp),
		new(models.FactionStanding),
		new(models.CharacterAchievement),
	)

	if err != nil {
		return nil, err
	}

	_, err = cn.Engine.Count(new(models.Character))
	if err != nil {
		return nil, err
	}

	// cn.SetLogger()

	return cn, nil
}

func (c *Core) Println(args ...any) {
	data := fmt.Sprintln(args...)
	data = strings.TrimRight(data, "\r\n")

	log.DefaultLogger.LogLine(&log.Line{
		Category: "db",
		Time:     time.Now(),
		Text:     data,
	})
}

func (c *Core) InitCache(dir string) error {
	c.Cache.Buckets = make(map[reflect.Type]*Bucket)
	c.Dir = dir
	c.Println("initing cache", dir)
	if err := os.RemoveAll(c.Dir); err != nil {
		return err
	}
	return os.MkdirAll(c.Dir, 0700)
}

func GetID(value reflect.Value) string {
	_, ok := value.Type().FieldByName("ID")
	if !ok {
		panic(value.Type().String())
	}
	text := value.FieldByName("ID").String()
	// if text == "" {
	// 	panic("ID cannot be empty + " + spew.Sdump(value.Interface()))
	// }
	return text
}

func SetEntry(value reflect.Value, entry uint32) {
	value.FieldByName("Entry").SetUint(uint64(entry))
}

func (c *Core) GetGameObjectTemplate(id string) (*models.GameObjectTemplate, error) {
	var gobj *models.GameObjectTemplate
	c.Lookup(BucketKeyStringID, id, &gobj)
	if gobj == nil {
		return nil, fmt.Errorf("no game object template by the ID %s", id)
	}
	return gobj, nil
}

func (c *Core) GetGameObjectTemplateByEntry(entry uint32) (*models.GameObjectTemplate, error) {
	var gobj *models.GameObjectTemplate
	c.Lookup(BucketKeyEntry, entry, &gobj)
	if gobj == nil {
		return nil, fmt.Errorf("no game object template by the entry %d", entry)
	}
	return gobj, nil
}

func (c *Core) GetItemTemplate(id string) (*models.ItemTemplate, error) {
	var item *models.ItemTemplate
	c.Lookup(BucketKeyStringID, id, &item)
	if item == nil {
		return nil, fmt.Errorf("no ItemTemplate by the ID %s", id)
	}
	return item, nil
}

func (c *Core) GetItemTemplateByEntry(entry uint32) (*models.ItemTemplate, error) {
	var item *models.ItemTemplate
	c.Lookup(BucketKeyEntry, entry, &item)
	if item == nil {
		return nil, fmt.Errorf("no ItemTemplate by the entry %d", entry)
	}
	return item, nil
}

// func (c *Core) GetStore(in interface{}) *Store {
// 	val := reflect.TypeOf(in)
// 	if val.Kind() == reflect.Ptr {
// 		val = val.Elem()
// 	}

// 	if val.Kind() == reflect.Slice {
// 		val = val.Elem()
// 	}

// 	if val.Kind() == reflect.Ptr {
// 		val = val.Elem()
// 	}

// 	st := c.DataStores[val]
// 	if st == nil {
// 		panic("no datastore for type " + val.String())
// 	}

// 	return st
// }

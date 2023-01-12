package wdb

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/cybriq/gotiny"

	"github.com/Gophercraft/log"
	"github.com/syndtr/goleveldb/leveldb"
)

type Bucket struct {
	Type     reflect.Type
	Storage  *leveldb.DB
	Index    *leveldb.DB
	Sequence uint64
	lrusize  int
	lruHead  *lruanswer
	lruTail  *lruanswer
}

type BucketKeyType uint8

const (
	BucketKeyStringID BucketKeyType = iota
	BucketKeyUint32ID
	BucketKeySequence
	BucketKeyEntry
	BucketIndexSequence
)

type BucketKey struct {
	Type BucketKeyType
	Data []byte
}

type Cache struct {
	Dir     string
	Buckets map[reflect.Type]*Bucket
}

func (c *Cache) Bucket(typeOf reflect.Type) (*Bucket, error) {
	if c.Buckets[typeOf] != nil {
		return c.Buckets[typeOf], nil
	}

	bckt := new(Bucket)
	bckt.Type = typeOf
	var err error
	bckt.Storage, err = leveldb.OpenFile(c.BucketPath(typeOf), nil)
	if err != nil {
		return nil, err
	}
	bckt.Index, err = leveldb.OpenFile(c.IndexPath(typeOf), nil)
	if err != nil {
		return nil, err
	}
	c.Buckets[typeOf] = bckt
	return bckt, nil
}

func (c *Cache) join(typeOf reflect.Type, kind string) string {
	return filepath.Join(c.Dir, fmt.Sprintf("%s.%s", typeOf.Name(), kind))
}

func (c *Cache) BucketPath(typeOf reflect.Type) string {
	return c.join(typeOf, "cache")
}

func (c *Cache) IndexPath(typeOf reflect.Type) string {
	return c.join(typeOf, "index")
}

func (c *Cache) Clear(typeOf reflect.Type) {
	bucket, ok := c.Buckets[typeOf]
	if ok {
		bucket.Storage.Close()
	}
	os.RemoveAll(c.BucketPath(typeOf))
	os.RemoveAll(c.IndexPath(typeOf))
}

func (c *Cache) Store(value reflect.Value) {
	if value.Kind() != reflect.Ptr {
		panic("value needs to be pointer " + value.Type().String())
	}

	base := value.Elem()
	bucket, err := c.Bucket(base.Type())
	if err != nil {
		panic(err)
	}

	var sequence [8]byte
	binary.LittleEndian.PutUint64(sequence[:], bucket.Sequence)
	bucket.Sequence++

	bucket.Index.Put(
		[]byte{
			byte(BucketIndexSequence),
		},

		sequence[:],

		nil,
	)

	data := gotiny.Marshal(value.Interface())

	if err := bucket.Storage.Put(sequence[:], data, nil); err != nil {
		panic(err)
		return
	}

	if base.Kind() == reflect.Struct {
		// Allow for fast lookup by entry using different BucketKeyType's
		entryField := base.FieldByName("Entry")
		if entryField.IsValid() {
			entry := uint32(base.FieldByName("Entry").Uint())
			var indexKey [5]byte
			indexKey[0] = byte(BucketKeyEntry)
			binary.LittleEndian.PutUint32(indexKey[1:], entry)
			if err := bucket.Index.Put(indexKey[:], sequence[:], nil); err != nil {
				panic(err)
			}
		}

		idField := base.FieldByName("ID")
		if idField.IsValid() {
			switch idField.Kind() {
			// Integral id
			case reflect.Uint32, reflect.Int32:
				var indexKey [5]byte
				indexKey[0] = byte(BucketKeyUint32ID)
				if idField.Kind() == reflect.Int32 {
					binary.LittleEndian.PutUint32(indexKey[1:], uint32(idField.Int()))
				} else {
					binary.LittleEndian.PutUint32(indexKey[1:], uint32(idField.Uint()))
				}
				if err := bucket.Index.Put(indexKey[:], sequence[:], nil); err != nil {
					panic(err)
				}
			case reflect.String:
				// Flattening ID
				var indexKey = append([]byte{byte(BucketKeyStringID)}, []byte(idField.String())...)
				if err := bucket.Index.Put(indexKey[:], sequence[:], nil); err != nil {
					panic(err)
				}
			}
		}
	}
}

func (bucket *Bucket) getSequence(sequence []byte, storage reflect.Value) {
	baseType := storage.Type().Elem().Elem()

	objectData, err := bucket.Storage.Get(sequence, nil)
	if err != nil {
		panic(err)
	}

	newPtr := reflect.New(baseType)
	gotiny.Unmarshal(objectData, newPtr.Interface())

	if storage.Elem().Kind() != reflect.Ptr {
		storage.Elem().Set(newPtr.Elem())
	} else {
		storage.Elem().Set(newPtr)
	}
}

func (c *Cache) Lookup(indexType BucketKeyType, key, value any) {
	// found := false

	storage := reflect.ValueOf(value)

	if storage.Kind() != reflect.Ptr {
		panic(storage.Type().String() + ": Lookup needs a pointer to a pointer to record data to")
	}

	if storage.Elem().Kind() != reflect.Map && storage.Elem().Kind() != reflect.Ptr {
		panic(storage.Type().String() + ": Storage type should be a reference")
	}

	baseType := storage.Type().Elem().Elem()

	bucket, err := c.Bucket(baseType)
	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if found {
	// 		bucket.cacheAnswer(indexType, key, storage)
	// 	}
	// }()

	// if bucket.findInMemoryAnswer(indexType, key, storage) {
	// 	return
	// }

	var sequence []byte

	// No need to find the sequence, we already have it
	if indexType == BucketKeySequence {
		binary.LittleEndian.PutUint64(sequence[:], key.(uint64))
	} else {
		// Get the sequence from the index
		var index []byte

		switch indexType {
		case BucketKeyUint32ID, BucketKeyEntry:
			index = make([]byte, 4)
			binary.LittleEndian.PutUint32(index, key.(uint32))
		case BucketKeyStringID:
			index = []byte(key.(string))
		default:
			panic(indexType)
		}

		var indexBytes = append([]byte{byte(indexType)}, index...)

		sequence, err = bucket.Index.Get(indexBytes, nil)
		if err != nil {
			return
		}
	}

	bucket.getSequence(sequence, storage)
	// found = true
}

func (bucket *Bucket) rangeSequence(callback reflect.Value) {
	record := reflect.New(bucket.Type)

	iter := bucket.Storage.NewIterator(nil, nil)

	if callback.Type().NumOut() == 1 && callback.Type().Out(0).Kind() == reflect.Bool {
		// Return false to break
		for iter.Next() {
			value := iter.Value()
			gotiny.Unmarshal(value, record.Interface())
			out := callback.Call([]reflect.Value{record})
			if !out[0].Bool() {
				return
			}
		}
		return
	} else {
		for iter.Next() {
			value := iter.Value()
			gotiny.Unmarshal(value, record.Interface())
			callback.Call([]reflect.Value{record})
		}
	}

	iter.Release()
	err := iter.Error()
	log.Warn(err)
	return
}

func (c *Cache) Range(fn interface{}) {
	callback := reflect.ValueOf(fn)

	recordType := callback.Type().In(0)
	if recordType.Kind() == reflect.Ptr {
		recordType = recordType.Elem()
	}

	bucket, err := c.Bucket(recordType)
	if err != nil {
		panic(err)
	}

	bucket.rangeSequence(callback)
}

func (c *Cache) Size() uint64 {
	var dirSize uint64 = 0

	readSize := func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			dirSize += uint64(file.Size())
		}

		return nil
	}

	filepath.Walk(c.Dir, readSize)

	return dirSize
}

// type Query struct {
// 	Cache  *Cache
// 	Checks []QueryFunc
// }

// type QueryFunc func(cursor reflect.Value) bool

// func (c *Cache) Query() *Query {
// 	return &Query{Cache: c}
// }

// func (q *Query) Q(fn QueryFunc) *Query {
// 	q.Checks = append(q.Checks, fn)
// 	return q
// }

// func (q *Query) Range(func (any) bool )

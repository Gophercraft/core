package wdb

import (
	"reflect"

	"github.com/cybriq/gotiny"
)

type QueryFlags uint8

const (
	QueryIs QueryFlags = 1 << iota
	QueryNot
	QueryExact
	QueryGt
	QueryLt
	QueryEq
)

type QueryCondition struct {
	Field    string
	Flags    QueryFlags
	Param1   reflect.Value
	Likeness float64
}

type Query struct {
	Core         *Core
	Conditions   []*QueryCondition
	OrderByField int
	ResultLimit  int
	Asc          bool
}

func (db *Core) Query() *Query {
	return &Query{
		Core: db,
	}
}

func (q *Query) Limit(limit int) *Query {
	q.ResultLimit = limit
	return q
}

func (q *Query) WhereFieldIs(fieldName string, value any) *Query {
	q.Conditions = append(q.Conditions, &QueryCondition{
		Field:  fieldName,
		Flags:  QueryExact | QueryIs,
		Param1: reflect.ValueOf(value),
	})
	return q
}

func getRecordField(record reflect.Value, field string) reflect.Value {
	t := record.Type()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == field {
			return record.Field(i)
		}
	}
	panic(field)
}

// func fieldMatchGtLtEq(cond *QueryCondition, field reflect.Value, param reflect.Value) bool {
// 	if cond.Flags&QueryEq != 0 {
// 		isEq := reflect.DeepEqual(field, param)
// 		if isEq {
// 			return true
// 		}
// 	}

// 	if cond.Flags&QueryLt != 0 {
// 		switch field.Kind() {
// 			case reflect.Int
// 		}
// 	}
// }

func (q *Query) recordIsMatch(record reflect.Value) bool {
	for _, cond := range q.Conditions {
		recordResult := false

		switch {
		case cond.Flags&(QueryIs|QueryExact) != 0:
			recordResult = reflect.DeepEqual(getRecordField(record, cond.Field), cond.Param1)
		default:

		}

		// Invert result if the flags contain Not bit.
		if recordResult && cond.Flags&QueryNot != 0 {
			return false
		} else {
			if !recordResult && cond.Flags&QueryNot == 0 {
				return false
			}
		}
	}

	return true
}

func (q *Query) Get(sliceptr any) error {
	slice := reflect.ValueOf(sliceptr).Elem()

	sliceElementType := slice.Type().Elem()

	bucket, err := q.Core.Bucket(sliceElementType)
	if err != nil {
		return err
	}

	zero := reflect.New(bucket.Type)
	record := reflect.New(bucket.Type)

	iter := bucket.Storage.NewIterator(nil, nil)

	for iter.Next() {
		record.Elem().Set(zero.Elem())
		value := iter.Value()
		gotiny.Unmarshal(value, record.Interface())

		if q.recordIsMatch(record.Elem()) {
			slice.Set(reflect.Append(slice, record.Elem()))
			continue
		}
	}

	iter.Release()
	return iter.Error()
}

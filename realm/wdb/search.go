package wdb

import (
	"fmt"
	"reflect"
	"regexp"
	"sort"

	"github.com/Gophercraft/core/i18n"
	"github.com/cybriq/gotiny"
)

var (
	stringType   = reflect.TypeOf(string(""))
	i18nTextType = reflect.TypeOf(i18n.Text{})
)

type sortableTemplateSlice struct {
	Locale i18n.Locale
	reflect.Value
}

func (v sortableTemplateSlice) Swap(i, j int) {
	x, y := v.Index(i).Interface(), v.Index(j).Interface()
	v.Index(i).Set(reflect.ValueOf(y))
	v.Index(j).Set(reflect.ValueOf(x))
}

func (a sortableTemplateSlice) Less(i, j int) bool {
	return GetName(a.Index(i).Elem(), a.Locale) < GetName(a.Index(j).Elem(), a.Locale)
}

func SortNamedTemplates(locale i18n.Locale, value reflect.Value) {
	sort.Sort(sortableTemplateSlice{locale, value})
}

func GetField(fieldName string, value reflect.Value, locale i18n.Locale) (text string) {
	field := value.FieldByName(fieldName)
	if !field.IsValid() {
		text = ""
		return
	}

	switch field.Type() {
	case stringType:
		text = field.Interface().(string)
	case i18nTextType:
		itext := field.Interface().(i18n.Text)
		text = itext.GetLocalized(locale)
	}
	return
}

func GetName(value reflect.Value, locale i18n.Locale) (text string) {
	text = GetField("Name", value, locale)
	return
}

// type CacheQuery struct {
// 	Fields
// }

// Queries data of the type that is the element of "slice" by its name
// Slice will hold the result if a name of the locale "locale" is matched by the regex.

func (c *Cache) QueryField(locale i18n.Locale, field string, regex string, limit int64, sliceptr any) error {
	storage := reflect.ValueOf(sliceptr)
	if storage.Kind() != reflect.Ptr {
		panic(storage.Type().String() + ": QueryNames needs a pointer to a slice to record data to")
	}

	slice := storage.Elem()

	elementType := storage.Type().Elem().Elem()

	nameField, ok := elementType.FieldByName(field)
	if !ok {
		return fmt.Errorf("wdb: You attempted to run QueryNames on a struct (%s) that has no %s field.", elementType, field)
	}

	if nameField.Type != stringType && nameField.Type != i18nTextType {
		return fmt.Errorf("wdb: Name field exists in a struct %s but that Name isn't of a queryable type: %s, (must be either string or i18n.Text)", elementType, nameField.Type)
	}

	rgx, err := regexp.Compile(regex)
	if err != nil {
		return err
	}

	bukket := c.Buckets[elementType]
	if bukket == nil {
		return fmt.Errorf("Cannot query names %s", elementType)
	}

	iter := bukket.Storage.NewIterator(nil, nil)
	defer iter.Release()

	var matches int64

	cursor := reflect.New(elementType)

	for iter.Next() {
		cursor.Elem().Set(reflect.Zero(elementType))
		value := iter.Value()

		gotiny.Unmarshal(value, cursor.Interface())

		var text string = GetField(field, cursor.Elem(), locale)

		if text != "" {
			if rgx.MatchString(text) {
				if matches >= limit {
					return nil
				}
				matches++
				slice.Set(reflect.Append(slice, cursor.Elem()))
			}
		}
	}

	// SortNamedTemplates(locale, slice)

	return nil
}

func (c *Cache) QueryNames(locale i18n.Locale, regex string, limit int64, sliceptr any) error {
	return c.QueryField(locale, "Name", regex, limit, sliceptr)
}

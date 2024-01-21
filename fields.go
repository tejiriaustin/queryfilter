package main

import (
	"reflect"
)

type (
	Struct struct {
		raw     interface{}
		value   reflect.Value
		TagName string
		Key     string
	}
)

var (
	DefaultTagName = "structs"
)

func New(s interface{}, key string) *Struct {
	return &Struct{
		raw:     s,
		value:   strctVal(s),
		TagName: DefaultTagName,
		Key:     key,
	}
}

func strctVal(s interface{}) reflect.Value {
	v := reflect.ValueOf(s)

	// if pointer get the underlying element≤
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("not struct")
	}

	return v
}

func (s *Struct) Fields(fields ...string) []string {
	return getStructFields(s.value, "", s.Key, s.TagName, fields...)
}

func Fields(s interface{}, key string) []string {
	var (
		fields []string
	)
	return New(s, key).Fields(fields...)
}

func getStructFields(v reflect.Value, parentField string, key, tagName string, fields ...string) []string {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		fieldValue := v.Field(i)
		structField := t.Field(i)

		switch fieldValue.Kind() {
		case reflect.Struct:
			fields = append(fields, getStructFields(fieldValue, joinFields(parentField, structField.Tag.Get(key)), key, tagName)...)
		case reflect.Map:
			fields = append(fields, getMapFields(fieldValue, joinFields(parentField, structField.Tag.Get(key)))...)
		case reflect.String:
			if tag := structField.Tag.Get(key); tag == "-" { // ignore empty tags
				continue
			}
			fields = append(fields, joinFields(parentField, structField.Tag.Get(key)))
		default:
			continue
		}
	}

	return fields
}

func joinFields(parentField string, field string) string {
	if parentField != "" {
		field = parentField + "." + field
	}
	return field
}

func getMapFields(v reflect.Value, parentField string) []string {
	var (
		fields []string
	)
	keys := v.MapKeys()
	for _, key := range keys {
		fields = append(fields, joinFields(parentField, key.String()))
	}
	return fields
}

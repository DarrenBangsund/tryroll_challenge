package db

import (
	"reflect"
)

type field struct {
	Name      string
	Value     interface{}
	Type      string
	Queryable bool
	Tag       string
}

type document struct {
	value  interface{}
	fields map[string]field
}

func handleMap(value map[string]interface{}, v reflect.Value) *document {
	f := map[string]field{}

	for mapKey, mapValue := range value {
		t := reflect.TypeOf(mapValue)

		f[mapKey] = field{
			Name:  mapKey,
			Value: mapValue,
			Type:  t.Kind().String(),
			Tag:   mapKey,
		}
	}

	return &document{
		value:  value,
		fields: f,
	}
}

func newDocument(value interface{}) *document {
	var v reflect.Value

	switch reflect.TypeOf(value).Kind().String() {
	// if the kind of the value is a pointer then we extract the value from the memory location
	case "ptr":
		v = reflect.ValueOf(value).Elem()
	// otherwise just pull the value
	default:
		v = reflect.ValueOf(value)
	}

	switch value.(type) {
	// maps are supported, but need to be handled different than structs
	case map[string]interface{}:
		return handleMap(value.(map[string]interface{}), v)
	// handle the struct
	default:
		fields := map[string]field{}

		for i := 0; i < v.NumField(); i++ {
			// extract metadata from the fields. This is used for filtering
			fieldName := v.Type().Field(i).Name
			fieldType := v.Type().Field(i).Type.String()
			fieldValue := v.Field(i).Interface()
			comparable := v.Field(i).Comparable()
			tag, ok := v.Type().Field(i).Tag.Lookup("db")

			if !ok {
				tag = fieldName
			}

			fields[tag] = field{
				Name:      fieldName,
				Value:     fieldValue,
				Type:      fieldType,
				Queryable: comparable,
				Tag:       tag,
			}
		}

		return &document{
			value:  value,
			fields: fields,
		}
	}
}

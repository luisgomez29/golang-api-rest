package utils

import (
	"reflect"
)

func Fields(model interface{}, args ...string) []string {
	e := reflect.ValueOf(model).Elem()
	var fields []string
	for i := 0; i < e.NumField(); i++ {
		field := e.Type().Field(i)
		if field.Tag.Get("json") == "-" || field.Type.Kind() == reflect.Struct || Includes(args, field.Name) {
			continue
		}
		fields = append(fields, field.Name)
	}
	return fields
}

func Includes(s interface{}, e interface{}) bool {
	listVal := reflect.ValueOf(s)
	for i := 0; i < listVal.Len(); i++ {
		if listVal.Index(i).Interface() == e {
			return true
		}
	}
	return false
}

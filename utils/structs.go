package utils

import (
	"reflect"
)

// Fields obtiene los campos de las estructuras. Puede pasar los campos a excluir.
func Fields(model interface{}, args ...string) []string {
	e := reflect.ValueOf(model).Elem()
	var fields []string
	for i := 0; i < e.NumField(); i++ {
		field := e.Type().Field(i)
		kind := field.Type.Kind()
		if field.Tag.Get("gorm") == "-" || kind == reflect.Struct || kind == reflect.Slice || Includes(args, field.Name) {
			continue
		}
		fields = append(fields, field.Name)
	}
	return fields
}

// Includes verifica si un slice incluye un elemento determinado.
func Includes(s interface{}, e interface{}) bool {
	listVal := reflect.ValueOf(s)
	for i := 0; i < listVal.Len(); i++ {
		if listVal.Index(i).Interface() == e {
			return true
		}
	}
	return false
}

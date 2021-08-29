package suct

import (
	"reflect"
)

// GetStructByName get struct value by fieldName
func GetStructByName(su interface{}, structName string) interface{} {
	if reflect.TypeOf(su).Kind() != reflect.Struct {
		return nil
	}
	return reflect.ValueOf(su).FieldByName(structName).Interface()
}

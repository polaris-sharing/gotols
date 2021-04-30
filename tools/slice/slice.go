package slice

import (
	"reflect"
)

// Int64InSlice arg int64 is in list
func Int64InSlice(arg int64, list []int64) bool {
	for _, v := range list {
		if arg == v {
			return true
		}
	}
	return false
}

// Contain Contain if in the list
func Contain(target interface{}, list interface{}) bool {
	if list == nil {
		return false
	}
	vl := reflect.ValueOf(list)
	switch reflect.TypeOf(list).Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < vl.Len(); i++ {
			if reflect.DeepEqual(target, vl.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		if vl.MapIndex(reflect.ValueOf(target)).IsValid() {
			return true
		}
	}
	return false
}

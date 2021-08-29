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

func AppendPos(src []interface{}, item interface{}, pos int) []interface{} {
	if len(src) == 0 {
		return nil
	}
	target := make([]interface{}, len(src)+1)
	if pos >= len(src) {
		target[cap(target)] = item
		return target
	}
	copy(target[:pos], src[:pos])
	target[pos] = item
	copy(target[pos+1:], src[pos:])
	return target
}

func MoveTail(src []interface{}, index int) []interface{} {
	if len(src) == 0 {
		return nil
	}
	target := make([]interface{}, len(src))
	if index >= len(src) {
		copy(target, src)
		return target
	}
	temp := src[index]
	copy(target[:index], src[:index])
	copy(target[index:], src[index+1:])
	target[len(target)-1] = temp
	return target
}

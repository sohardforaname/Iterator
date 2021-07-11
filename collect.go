package Iterator

import (
	"reflect"
)

func Collect(iterator Iterator) interface{} {
	slice := reflect.New(reflect.SliceOf(iterator.Type())).Elem()

	for iterator.HasNext() {
		slice = reflect.Append(slice, reflect.ValueOf(iterator.Next()))
	}
	return slice.Interface()
}
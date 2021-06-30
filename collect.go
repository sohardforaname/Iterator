package Iterator

func Collect(iterator Iterator) interface{} {
	slice := make([]interface{}, 0)
	for iterator.HasNext() {
		slice = append(slice, iterator.Next())
	}
	return slice
}
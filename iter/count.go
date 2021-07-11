package iter

func Count(iterator Iterator) int64 {
	cnt := int64(0)
	for iterator.HasNext() {
		iterator.Next()
		cnt++
	}
	return cnt
}
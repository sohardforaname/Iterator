package Iterator

type RepeatIterator struct {
	times int64
	first Iterator
	Iterator
}

func NewRepeatIterator(iterator Iterator, times int64) *RepeatIterator {
	return &RepeatIterator{times, iterator.Copy(), iterator}
}
package Iterator

type RepeatStruct struct {
	times int64
	first Iterator
	Iterator
}

func NewRepeatStruct(times int64, iterator Iterator) *RepeatStruct {
	return &RepeatStruct{times, iterator.Copy(), iterator}
}
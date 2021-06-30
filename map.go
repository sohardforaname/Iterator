package Iterator

type MapStruct struct {
	Iterator
	Mapper
}

func NewMapStruct(iterator Iterator, mapper Mapper) *MapStruct {
	return &MapStruct{iterator, mapper}
}

func (m *MapStruct) HasNext() bool {
	return m.Iterator.HasNext()
}

func (m *MapStruct) Next() interface{} {
	return m.Mapper(m.Iterator.Next())
}

func (m *MapStruct) Collect() interface{} {
	return Collect(m)
}

package Iterator

import "reflect"

type MapIterator struct {
	Iterator
	mapper interface{}
}

func NewMapIterator(iterator Iterator, mapper interface{}) *MapIterator {
	return &MapIterator{iterator, mapper}
}

func (m *MapIterator) HasNext() bool {
	return m.Iterator.HasNext()
}

func (m *MapIterator) Next() interface{} {
	return reflect.ValueOf(m.mapper).Call([]reflect.Value{reflect.ValueOf(m.Iterator.Next())})[0].Interface()
}

func (m *MapIterator) Type() reflect.Type {
	return reflect.TypeOf(m.mapper).Out(0)
}

func (m *MapIterator) Map(mapper interface{}) *MapIterator {
	return NewMapIterator(m, mapper)
}

func (m *MapIterator) Filter(predictor interface{}) *FilterIterator {
	return NewFilterIterator(m, predictor)
}

func (m *MapIterator) Collect() interface{} {
	return Collect(m)
}

func (m *MapIterator) Count() int64 {
	return Count(m)
}

func (m *MapIterator) Repeat(times int64) *RepeatIterator {
	return NewRepeatIterator(m, times)
}

func (m *MapIterator) Copy() Iterator {
	return &MapIterator{m.Iterator.Copy(), m.mapper}
}
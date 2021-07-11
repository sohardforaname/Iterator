package Iterator

import "reflect"

type FilterIterator struct {
	buffer interface{}
	Iterator
	predictor interface{}
}

func NewFilterIterator(iterator Iterator, predictor interface{}) *FilterIterator {
	return &FilterIterator{nil, iterator, predictor}
}

func (f *FilterIterator) HasNext() bool {
	if f.buffer != nil {
		return true
	}
	for f.Iterator.HasNext() {
		val := f.Iterator.Next()
		if reflect.ValueOf(f.predictor).Call([]reflect.Value{reflect.ValueOf(val)})[0].Bool() {
			f.buffer = val
			return true
		}
	}
	return false
}

func (f *FilterIterator) Next() interface{} {
	if f.buffer != nil || f.HasNext() {
		val := f.buffer
		f.buffer = nil
		return val
	}
	panic("try get the element after the end of a iterator")
}

func (f *FilterIterator) Type() reflect.Type {
	return f.Iterator.Type()
}

func (f *FilterIterator) Map(mapper interface{}) *MapIterator {
	return NewMapIterator(f, mapper)
}

func (f *FilterIterator) Filter(predictor interface{}) *FilterIterator {
	return NewFilterIterator(f, predictor)
}

func (f *FilterIterator) Collect() interface{} {
	return Collect(f)
}

func (f *FilterIterator) Count() int64 {
	return Count(f)
}

func (f *FilterIterator) Repeat(times int64) *RepeatIterator {
	return NewRepeatIterator(f, times)
}

func (f *FilterIterator) Copy() Iterator {
	return &FilterIterator{f.buffer, f.Iterator.Copy(), f.predictor}
}

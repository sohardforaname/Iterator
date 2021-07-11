package iter

import (
	"reflect"
)

type base struct {
	size, index int
}

type SliceIterator struct {
	*base
	buf interface{}
}

func OfSlice(s interface{}) *SliceIterator {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Slice {
		panic("Not a slice")
	}
	c := reflect.ValueOf(s).Len()
	return &SliceIterator{
		buf: s,
		base: &base{
			index: 0,
			size:  c,
		},
	}
}

func (s *SliceIterator) HasNext() bool {
	return s.index < s.size
}

func (s *SliceIterator) Next() interface{} {
	cur := reflect.ValueOf(s.buf).Index(s.index).Interface()
	s.index++
	return cur
}

func (s *SliceIterator) Type() reflect.Type {
	return reflect.TypeOf(s.buf).Elem()
}

func (s *SliceIterator) Map(mapper interface{}) *MapIterator {
	return NewMapIterator(s, mapper)
}

func (s *SliceIterator) Filter(predictor interface{}) *FilterIterator {
	return NewFilterIterator(s, predictor)
}

func (s *SliceIterator) Collect() interface{} {
	return Collect(s)
}

func (s *SliceIterator) Count() int64 {
	return Count(s)
}

func (s *SliceIterator) Repeat(times int64) *RepeatIterator {
	return NewRepeatIterator(s, times)
}

func (s *SliceIterator) Copy() Iterator {
	return &SliceIterator{
		&base{
			index: s.index,
			size:  s.size,
		},
		s.buf,
	}
}

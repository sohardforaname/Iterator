package Iterator

import "reflect"

type base struct {
	size, index int
}

type SliceIterator struct {
	*base
	buf interface{}
}

func OfSlice(s interface{}) *SliceIterator {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		panic("Not a slice")
	}
	v := reflect.ValueOf(s)
	c := v.Len()
	return &SliceIterator{
		buf: v.Interface(),
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

func (s *SliceIterator) Map(mapper Mapper) *MapStruct {
	return NewMapStruct(s, mapper)
}

func (s *SliceIterator) Filter(predictor Predictor) *FilterStruct {
	return NewFilterStruct(s, predictor)
}

func (s *SliceIterator) Collect() interface{} {
	return Collect(s)
}

func (s *SliceIterator) Count() int64 {
	return Count(s)
}

func (s *SliceIterator) Repeat(times int64) *RepeatStruct {
	return NewRepeatStruct(times, s)
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

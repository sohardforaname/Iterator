package iter

import "reflect"

type RepeatIterator struct {
	times int64
	first Iterator
	Iterator
}

func NewRepeatIterator(iterator Iterator, times int64) *RepeatIterator {
	return &RepeatIterator{times, iterator.Copy(), iterator}
}

func (r *RepeatIterator) HasNext() bool {
	return r.times > 0 || r.Iterator.HasNext()
}

func (r *RepeatIterator) Next() interface{} {
	if r.Iterator.HasNext() {
		return r.Iterator.Next()
	}
	r.Iterator = r.first.Copy()
	r.times--
	return r.Iterator.Next()
}

func (r *RepeatIterator) Type() reflect.Type {
	return r.Iterator.Type()
}

func (r *RepeatIterator) Map(mapper interface{}) *MapIterator {
	return NewMapIterator(r, mapper)
}

func (r *RepeatIterator) Filter(predictor interface{}) *FilterIterator {
	return NewFilterIterator(r, predictor)
}

func (r *RepeatIterator) Collect() interface{} {
	return Collect(r)
}

func (r *RepeatIterator) Count() int64 {
	return Count(r)
}

func (r *RepeatIterator) Repeat(times int64) *RepeatIterator {
	return NewRepeatIterator(r, times)
}

func (r *RepeatIterator) Copy() Iterator {
	return &RepeatIterator{
		r.times,
		r.first,
		r.Iterator.Copy(),
	}
}

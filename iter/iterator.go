package iter

import "reflect"

type Iterator interface {
	HasNext() bool
	Next() interface{}
	Type() reflect.Type
	Map(mapper interface{}) *MapIterator
	Filter(predictor interface{}) *FilterIterator
	Collect() interface{}
	Count() int64
	Repeat(times int64) *RepeatIterator
	Copy() Iterator
}

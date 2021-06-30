package Iterator

type Iterator interface {
	HasNext() bool
	Next() interface{}
	Map(mapper Mapper) *MapStruct
	Filter(predictor Predictor) *FilterStruct
	Collect() interface{}
	Count() int64
	Repeat(times int64) *RepeatStruct
	Copy() Iterator

}

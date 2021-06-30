package Iterator

type FilterStruct struct {
	buffer interface{}
	Iterator
	Predictor
}

func NewFilterStruct(iterator Iterator, predictor Predictor) *FilterStruct {
	return &FilterStruct{nil, iterator, predictor}
}
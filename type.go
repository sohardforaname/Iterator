package Iterator

type Mapper func(a interface{}) interface{}

type Predictor func(a interface{}) bool

type Equal func(a, b interface{}) bool

type BinaryOp func(a, b interface{}) interface{}

type Comparator func(o1, o2 interface{}) int

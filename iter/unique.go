package iter

import "reflect"

type UniqueIterator struct {
	buffer interface{}
	equal interface{}
	Iterator
}

func NewUniqueIterator(iterator Iterator, equal interface{}) *UniqueIterator {
	return &UniqueIterator{nil,equal, iterator}
}

func (u *UniqueIterator) HasNext() bool {
	if u.buffer == nil {
		return u.Iterator.HasNext()
	}
	for u.Iterator.HasNext() {
		val := u.Iterator.Next()
		if reflect.ValueOf(u.equal).Call([]reflect.Value{reflect.ValueOf(u.buffer), reflect.ValueOf(val)})[0].Bool() {
			u.buffer = val
			return true
		}
	}
	return false
}
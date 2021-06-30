package Iterator

import (
	"fmt"
	"testing"
)

func TestSliceIterator_Map(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	slice1 := OfSlice(slice).Map(func(a interface{}) interface{} {
		return a.(int) + 1
	}).Collect()

	fmt.Println(slice1)
}

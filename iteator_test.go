package Iterator

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	slice1 := OfSlice(slice).Map(func (a int) int64 {
		return int64(a) + 1
	}).Filter(func (a int64) bool {
		return a % 2 == 0
	}).Collect()

	slice2 := OfSlice(slice1).Map(func (a int64) int {
		return int(a) + 1
	}).Filter(func (a int) bool {
		return a != 5
	}).Collect()

	fmt.Println(slice2)
	fmt.Println(slice1)
}

package test

import (
	"Iterator/iter"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	slice1 := iter.OfSlice(slice).Map(func(a int) int64 {
		return int64(a) + 1
	}).Filter(func(a int64) bool {
		return a%2 == 0
	}).Collect()

	slice2 := iter.OfSlice(slice1).Map(func (a int64) int {
		return int(a) + 1
	}).Filter(func (a int) bool {
		return a != 5
	}).Repeat(3).Collect()

	fmt.Println(slice1)
	fmt.Println(slice2)
}

func Benchmark1(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5}
	_ = iter.OfSlice(slice).Map(func(a int) int64 {
		return int64(a) + 1
	}).Filter(func(a int64) bool {
		return a%2 == 0
	}).Collect()
}
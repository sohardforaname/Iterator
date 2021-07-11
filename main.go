package main

import (
	"Iterator/iter"
	"fmt"
	"testing"
)

var length = 1000000

func main() {

	fn := func(b *testing.B) {
		b.StopTimer()
		slice := make([]int, length)
		for i := 0; i < length; i++ {
			slice[i] = i
		}
		b.StartTimer()
		l := iter.OfSlice(slice).Map(func(a int) int64 {
			return int64(a) + 1
		}).Filter(func(a int64) bool {
			return a%2 == 0
		}).Collect()
		b.StopTimer()

		fmt.Sprintln(l)
	}
	r := testing.Benchmark(fn)

	fmt.Println(r)
}

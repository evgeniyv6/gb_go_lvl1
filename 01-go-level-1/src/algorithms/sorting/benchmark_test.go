package sorting

import (
	"algorithms/various"
	"testing"
)

var (
	GlobalRes    []int
	sliceElemNum = 10
)

func BenchmarkInsertationSort(b *testing.B) {
	var res []int
	mySl := various.GenerSlice(sliceElemNum)
	for i := 0; i < b.N; i++ {
		res = InsertationSort(mySl)
	}
	GlobalRes = res
}

func BenchmarkBubleSort(b *testing.B) {
	var res []int
	mySl := various.GenerSlice(sliceElemNum)
	for i := 0; i < b.N; i++ {
		res = BubleSort(mySl)
	}
	GlobalRes = res
}

func BenchmarkQuickSort(b *testing.B) {
	mySl := various.GenerSlice(sliceElemNum)
	for i := 0; i < b.N; i++ {
		QuickSort(mySl, 0, len(mySl)-1)
	}
}

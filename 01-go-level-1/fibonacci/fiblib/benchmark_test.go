package fiblib

import "testing"

var (
	GlobalRes uint64
	GlobalN   uint64 = 10
)

func BenchmarkFibonacciRecur(b *testing.B) {
	var res uint64
	for i := 0; i < b.N; i++ {
		res = FibonacciRecur(GlobalN)
	}
	GlobalRes = res
}

func BenchmarkFibonacciGoldenRatio(b *testing.B) {
	var res uint64
	for i := 0; i < b.N; i++ {
		res = FibonacciRecur(GlobalN)
	}
	GlobalRes = res
}

func BenchmarkFibonacciMapUse(b *testing.B) {
	var res uint64
	for i := 0; i < b.N; i++ {
		res = FibonacciRecur(GlobalN)
	}
	GlobalRes = res
}

func BenchmarkFibonacciAnonymousFunc(b *testing.B) {
	var res uint64
	for i := 0; i < b.N; i++ {
		res = FibonacciRecur(GlobalN)
	}
	GlobalRes = res
}

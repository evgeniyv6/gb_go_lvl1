package fiblib

import (
	"math"
)

// FibonacciMapUse descr (bif o: O(n))
func FibonacciMapUse(n uint64) uint64 {
	fibMap := make(map[uint64]uint64)
	var num, i uint64
	for ; i <= n; i++ {
		switch i {
		case 0:
			num = 0
		case 1:
			num = 1
		default:
			num = fibMap[i-1] + fibMap[i-2]
		}
		fibMap[i] = num
	}
	return fibMap[n]
}

// FibonacciAnonymousFunc descr (bif o: O(n))
func FibonacciAnonymousFunc() func() uint64 {
	var n1, n2 uint64 = 0, 1
	return func() uint64 {
		n1, n2 = n2, n1+n2
		return n1
	}
}

// FibonacciGoldenRatio descr
// Thanks wiki (Binet's Formula):
// 		phi = (1+sqrt(5))/2 =~ 1.618033988749895
//      formula: n = (pow(phi,n) - pow((1 - phi), n)) / sqrt(5)
func FibonacciGoldenRatio(n uint64) uint64 {
	// const phi = 1.618033988749895
	if n < 0 {
		return 0
	}
	var nf = float64(n)
	phi := (1 + math.Pow(5, .5)) / 2.0
	return uint64(math.Round((math.Pow(phi, nf) - math.Pow((1-phi), nf)) / math.Sqrt(5.0)))
}

// FibonacciRecur descr
func FibonacciRecur(n uint64) uint64 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}
	return FibonacciRecur(n-1) + FibonacciRecur(n-2)
}

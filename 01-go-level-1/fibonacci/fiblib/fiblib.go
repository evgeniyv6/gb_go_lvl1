package fiblib

import (
	"fmt"
	"math"
)

// FibonacciMapUse descr (bif o: O(n))
func FibonacciMapUse(n int) (fibMap map[int]int) {
	fmt.Println("--== Fibonacci map using ==--")
	fibMap = make(map[int]int)
	if n < 0 {
		return
	}
	var num int
	for i := 0; i <= n; i++ {
		if i <= 1 {
			num = i
		} else {
			num = fibMap[i-1] + fibMap[i-2]
		}
		fibMap[i] = num
	}
	return
}

// FibonacciAnonymousFunc descr (bif o: O(n))
func FibonacciAnonymousFunc() func() int {
	fmt.Println("--== Fibonacci anonymous func using ==--")
	n1, n2 := 0, 1
	return func() int {
		n1, n2 = n2, n1+n2
		return n1
	}
}

// FibonacciGoldenRatio descr
// Thanks wiki (Binet's Formula):
// 		phi = (1+sqrt(5))/2 =~ 1.618033988749895
//      formula: n = (pow(phi,n) - pow((1 - phi), n)) / sqrt(5)
func FibonacciGoldenRatio(n int) int64 {
	// const phi = 1.618033988749895
	fmt.Println("--== Fibonacci golden ratio ==--")
	if n < 0 {
		return 0
	}
	var nf = float64(n)
	phi := (1 + math.Pow(5, .5)) / 2.0
	return int64(math.Round((math.Pow(phi, nf) - math.Pow((1-phi), nf)) / math.Sqrt(5.0)))
}

// FibonacciRecur descr
func FibonacciRecur(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecur(n-1) + FibonacciRecur(n-2)
}

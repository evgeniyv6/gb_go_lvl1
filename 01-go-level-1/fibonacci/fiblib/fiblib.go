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
// negative numbers will be catched by uint
func FibonacciGoldenRatio(n uint64) float64 {
	fmt.Println("--== Fibonacci golden ratio ==--")
	// const phi = 1.618033988749895
	if n <= 1 {
		return float64(n)
	}
	n = n - 1
	phi := (1 + math.Pow(5, .5)) / 2.0
	return math.Round((math.Pow(phi, float64(n)) - math.Pow((1-phi), float64(n))) / math.Sqrt(float64(5)))

}

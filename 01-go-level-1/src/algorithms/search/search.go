package search

import (
	"math"
	"strconv"
)

// FizzBuzz exercise
// на вход сделал указатель, чтобы не было лишнего копирования :-)
func FizzBuzz(num *int) string {
	switch true {
	case *num == 0:
		return ""
	case *num%3 == 0 && *num%5 == 0:
		return "FizzBuzz"
	case *num%3 == 0:
		return "Fizz"
	case *num%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(*num)
	}
}

// IsPrime will find prime numbers
// на вход указатель, чтобы не делать лишнего копирования :-)
func IsPrime(num *int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(*num)))); i++ {
		if *num%i == 0 {
			return false
		}
	}
	return *num > 1
}

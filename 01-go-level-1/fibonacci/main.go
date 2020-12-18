package main

import (
	"bufio"
	"fmt"
	"github.com/evgeniyv6/gb_go_lvl1/01-go-level-1/fibonacci/fiblib"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

type intMap map[int]int

func (m intMap) mapSorter() (sorted []int) {
	for k, _ := range m {
		sorted = append(sorted, k)
	}
	sort.Ints(sorted)
	return
}

// Fibonacci printer
// u need to point non-negative number, then one of the Fibonacci algorithms will be randomly selected:
//   - using map
//   - using golden ratio
//   - anonymous func
func main() {
	const helper = "to use: print not negative number"
	var num string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите число > ")
	for scanner.Scan() {
		num = scanner.Text()
		if scanner.Err() != nil {
			fmt.Println("Error while stdin read.")
		}
		intNum, err := strconv.Atoi(num)
		if err != nil || intNum < 0 {
			fmt.Println("Cannot convert to int or number is negative.")
			fmt.Printf("num %v\n", intNum)
			fmt.Print("> ")
		} else {
			randChoiceNum := rand.Intn(4)
			switch {
			case randChoiceNum == 0:
				nn := fiblib.FibonacciMapUse(intNum)
				fmt.Print("Fibonacci numbers from algorithm (countdown index starts from 0): ")
				for _, k := range intMap(nn).mapSorter() {
					fmt.Printf("%d ", nn[k])
				}
				fmt.Printf("\nFibonacci number for %d is: %d", intNum, nn[intNum-1])
				fmt.Println()
			case randChoiceNum == 1:
				fib := fiblib.FibonacciAnonymousFunc()
				fmt.Print("Fibonacci numbers from algorithm: ")
				if intNum == 1 {
					fmt.Print("0")
				}
				var res int
				for i := 1; i < intNum; i++ {
					res = fib()
					fmt.Printf("%d ", res)
				}
				fmt.Printf("\nFibonacci number for %d is: %d", intNum, res)
				fmt.Println()
			case randChoiceNum == 2:
				var fib int64
				if intNum == 1 {
					fmt.Println("--== Fibonacci golden ratio ==--")
					fmt.Printf("Fibonacci number for %d is: %v", intNum, 0)
					fmt.Println()
				} else {
					fib = fiblib.FibonacciGoldenRatio(intNum - 1)
					fmt.Printf("Fibonacci number for %d is: %v", intNum, fib)
					fmt.Println()
				}
			case randChoiceNum == 3:
				if intNum == 1 {
					fmt.Println("--== Fibonacci recursion ==--")
					fmt.Printf("Fibonacci number for %d is: %v", intNum, 0)
					fmt.Println()
				} else {
					fmt.Println("--== Fibonacci recursion ==--")
					fmt.Printf("Fibonacci number for %d is: %v", intNum, fiblib.FibonacciRecur((intNum)-1))
					fmt.Println()
				}

			default:
				fmt.Println(helper)
			}
			fmt.Print("> ")
		}
	}
}

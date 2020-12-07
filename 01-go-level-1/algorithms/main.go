package main

import (
	"algorithms/search"
	"algorithms/sorting"
	"algorithms/various"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// to start: go run ./main.go -action=<available action>
// Available actions: primes fizzbuzz len qsort brackets
//		primes: returns prime numbers from a list with the selected capacity. big O: O(n^2)
//  	fizzbuzz: returns fizzbuzz numbers from a list with the selected capacity
//		len: returns bytes and chars len of "Привет go!" string
//		qsort: tell the <number>, then will be created slice with <number> capacity,
//								then the quick sort algorithm will be applied. big O: O(sqrt (n) / log n)
//      brackets: tell the <number>, see the brackets combinations
//      buble: tell the <number>, then will be created slice with <number> capacity,
//     							then the bubble sort algorithm will be applied. big O(n^2) at worst
//      insert: tell the <number>, then will be created slice with <number> capacity,
//     							then the insert sort algorithm will be applied.big O(n^2) at worst
func main() {
	chooseAction := flag.String("action", "", "primes fizzbuzz len qsort brackets bubble insert")
	flag.Parse()

	switch true {
	case *chooseAction == "primes" || *chooseAction == "fizzbuzz":
		var a int
		m := map[string][]string{
			"prime":    {},
			"fizzbuzz": {},
		}
		fmt.Print("Введите число: ")
		fmt.Scanln(&a)

		for i := 0; i <= a; i++ {
			if search.IsPrime(i) {
				m["prime"] = append(m["prime"], strconv.Itoa(i))
			}
			m["fizzbuzz"] = append(m["fizzbuzz"], search.FizzBuzz(i))
		}
		fmt.Printf("Простые числа: %s\nFizzBuzz: %s", strings.Join(m["prime"], " "),
			strings.Join(m["fizzbuzz"], " "))
		fmt.Println()
	case *chooseAction == "len":
		strBytesCount, strCharsCount, strCharsUtf8Count := various.FindStrLen("Привет go!")
		fmt.Printf("'Hello go!' -> bytes: %d, chars1: %d, chars2: %d", strBytesCount,
			strCharsCount, strCharsUtf8Count)
		fmt.Println()
	case *chooseAction == "qsort":
		var a int
		fmt.Print("Введите число: ")
		fmt.Scanln(&a)
		nums := various.GenerSlice(a)
		fmt.Println("Исходный список: ", nums)
		sorting.QuickSort(nums, 0, len(nums)-1)
		fmt.Println("Отсортированный список: ", nums)
	case *chooseAction == "brackets":
		var a int8
		fmt.Print("Введите число: ")
		fmt.Scanln(&a)
		various.PrintBrackets(a, 0, 0, "")
	case *chooseAction == "bubble":
		var a int
		fmt.Print("Введите число: ")
		fmt.Scanln(&a)
		sor := various.GenerSlice(a)
		fmt.Printf("Оригинальный список: %v\n", sor)
		fmt.Printf("Сортировка пузырьком: %v\n", sorting.BubleSort(sor))
		fmt.Printf("Оригинальный список не изменился: %v\n", sor)
	case *chooseAction == "insert":
		var a int
		fmt.Print("Введите число: ")
		fmt.Scanln(&a)
		sor := various.GenerSlice(a)
		fmt.Printf("Оригинальный список: %v\n", sor)
		fmt.Printf("Сортировка вставками: %v\n", sorting.InsertationSort(sor))
		fmt.Printf("Оригинальный список не изменился: %v\n", sor)
	default:

		fmt.Println("Неверно выбран флаг, возможные значения: primes fizzbuzz len qsort brackets")
	}
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type calcLog func(string, float64, ...float64) float64

func chkElem(opers []string, elem string) bool {
	for _, op := range opers {
		if op == elem {
			return true
		}
	}
	return false
}

func addLog(f func(o string, a float64, b ...float64) float64) func(o string, a float64, b ...float64) float64 {
	return func(o string, a float64, b ...float64) float64 {
		fmt.Print("Пучатаем результат: ")
		res := f(o, a, b...)
		return res
	}
}

// decorator with type calcLog func(string, float64, ...float64) float64
func addLog2(f calcLog) calcLog {
	return func(o string, a float64, b ...float64) float64 {
		fmt.Println("Пучатаем результат... ")
		res := f(o, a, b...)
		fmt.Println("Результат: ")
		return res
	}
}

func simpleCalculator(op string, a float64, b ...float64) float64 {
	switch op {
	case "+":
		return a + b[0]
	case "-":
		return a - b[0]
	case "*":
		return a * b[0]
	case "/":
		return a / b[0]
	case "pow":
		return math.Pow(a, b[0])
	case "sqrt":
		return math.Sqrt(a)
	case "log":
		return math.Log(a)
	case "exp":
		return math.Exp(a)
	default:
		fmt.Print("Неверная операция")
		return 0
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var fa, fb float64
	oneOp := []string{"sqrt", "log", "exp"}
	twoOp := []string{"+", "-", "*", "/", "pow"}

	fmt.Print("Введите арифметическую операцию (+, -, *, /, pow) для 2х чисел, (sqrt, log, exp) для одного: ")
	op, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	op = strings.TrimSuffix(op, "\n")

	if chkElem(oneOp, op) || chkElem(twoOp, op) {
		fmt.Print("Введите число: ")
		a, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fa, err = strconv.ParseFloat(strings.TrimSuffix(a, "\n"), 64)
		if err != nil {
			fmt.Println("Вы ввели не число")
			os.Exit(1)
		}
		if chkElem(twoOp, op) {
			fmt.Print("Введите 2е число: ")
			b, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}
			fb, err = strconv.ParseFloat(strings.TrimSuffix(b, "\n"), 64)
			if err != nil {
				fmt.Println("Вы ввели не число")
				os.Exit(1)
			}
			xCalc := addLog2(simpleCalculator)
			fmt.Printf("%.2f\n", xCalc(op, fa, fb))
		} else {
			xCalc := addLog(simpleCalculator)
			fmt.Printf("%.2f\n", xCalc(op, fa))
		}
	} else {
		fmt.Println("Неверная операция ")
		os.Exit(1)
	}
}

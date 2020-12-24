package main

import "fmt"

func main() {
	fmt.Println("Hi from strongs")
	var s string = "тест this is my string"
	fmt.Println(s[:4])
	runes := []rune(s)
	fmt.Println(string(runes[:4]))
}

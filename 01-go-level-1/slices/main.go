package main

import (
	"fmt"
	"github.com/evgeniyv6/gb_go_lvl1/01-go-level-1/slices/slactn"
)

func main() {
	fmt.Println("Hello from slices dir! Add exclamation mark.")
	xsOrign := []int{9, 9, 8, 8, 7, 7, 6, 5, 4, 3, 2, 1}
	fmt.Printf("Origin slice - %d", xsOrign)
	slactn.ReverseSlice(xsOrign)
	fmt.Printf("\nReversed slice - %d", xsOrign)
	fmt.Printf("\nDelete dublicates %d", slactn.DelDublicates(xsOrign))

}

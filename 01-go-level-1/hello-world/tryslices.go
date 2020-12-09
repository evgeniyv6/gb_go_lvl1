package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	var buf [256]byte
	fmt.Println(buf[255])

	var slice []byte = buf[:]
	fmt.Println(slice)
	slashPos := bytes.IndexRune(slice, 0)
	fmt.Println(slashPos)

	fmt.Println(len(slice), slice[255])

	xs := []int{1, 2, 3}
	fmt.Println(xs)
	for i, e := range xs {
		fmt.Println(i, e, xs[2], len(xs), cap(xs))
	}

	xsnew := xs[0:1]
	fmt.Println(cap(xsnew), len(xsnew))

	fmt.Println("--------------")

	// revers slice
	xsOrig := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, j := 0, len(xsOrig)-1; i < j; i, j = i+1, j-1 {
		xsOrig[i], xsOrig[j] = xsOrig[j], xsOrig[i]
	}
	fmt.Printf("xsOrig - %d", xsOrig)

	// dublicate slice
	xsDouble := append([]int(nil), xsOrig...)
	fmt.Printf("\nDublicated csOrig - %d", xsDouble)

	// copy try
	xsbCopy := make([]int, len(xsOrig)+2)
	copy(xsbCopy, xsOrig)
	fmt.Printf("\n\txsbCopy - %d", xsbCopy)

	// cut
	xsCutted := append(xsOrig[:3], xsOrig[4:]...)
	fmt.Printf("\nxsCutted %d", xsCutted)

	// dedublicates
	xsWithDuble := []int{1, 1, 2, 3, 4, 4, 9, 9, 0, 0, 3, 3, 2, 2, 8, 8, 7}
	removeDublicatesFromSlice(xsWithDuble)
	fmt.Printf("\nOrig dublicates %d", xsWithDuble)
	fmt.Printf("\ndel dublicates %d", removeDublicatesFromSlice(xsWithDuble))

	fmt.Printf("\nRemove dublicates 2 %d", removeDublicatesFromSlice2(xsWithDuble))

	fmt.Printf("\nOrig slice %v", xsOrig)
	reversSlice(xsOrig)
	fmt.Printf("\nReversted slice %v", xsOrig)

}

func removeDublicatesFromSlice(s []int) []int {
	keys := make(map[int]bool)
	newS := []int{}
	for _, el := range s {
		if _, val := keys[el]; !val {
			keys[el] = true
			newS = append(newS, el)
		}

	}
	sort.Ints(newS)
	return newS
}

func removeDublicatesFromSlice2(s []int) []int {
	sort.Ints(s)
	j := 0
	for i := 1; i < len(s); i++ {
		if s[j] == s[i] {
			continue
		}
		j++
		s[j] = s[i]
	}
	return s[:j+1]
}

func reversSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

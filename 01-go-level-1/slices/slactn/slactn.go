package slactn

import "fmt"

// ReverseSlice descr
func ReverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// DelDublicates descr
func DelDublicates(s []int) []int {
	keys := make(map[int]bool)
	newS := []int{}
	for _, el := range s {
		if i, v := keys[el]; !v {
			fmt.Printf("i = %v", i)
			keys[el] = true
			newS = append(newS, el)
		}
	}
	return newS
}

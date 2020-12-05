package various

import (
	"fmt"
	"math/rand"
	"unicode/utf8"
)

// PrintBrackets description here
func PrintBrackets(num int8, open int8, closed int8, bracket string) {
	if open+closed == 2*num {
		fmt.Println(bracket)
	}
	if open < num {
		PrintBrackets(num, open+1, closed, bracket+"{")
	}
	if closed < open {
		PrintBrackets(num, open, closed+1, bracket+"}")
	}
}

// FindStrLen returns bytes len, chars len
func FindStrLen(s string) (int, int, int) {
	return len(s), len([]rune(s)), utf8.RuneCountInString(s)
}

// GenerSlice generate slice
func GenerSlice(num int) []int {
	randNum := make([]int, num)
	for i := 0; i < num; i++ {
		randNum[i] = rand.Intn(100)
	}
	return randNum
}

package main

import (
	"fmt"
	"practice/collections"
	"practice/linked"
	"practice/queue"
	"practice/stack"
	"unicode/utf8"
)

var (
	customStr = "foo barпривет"
)

func main() {
	fmt.Println(len([]rune(customStr)))
	fmt.Println(utf8.RuneCountInString(customStr))
	fmt.Println(len(customStr))

	var (
		matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		m1     = map[int]int{
			1: 2,
			3: 4,
		}
		m2 = map[int]int{
			1: 2,
			3: 5,
		}
		unitsSlice  = []int{1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1}
		unitsMatrix = [][]int{{1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1}}

		primeCollection = map[string]collections.Collections{
			"stack": &collections.Stack{},
			"queue": &collections.Queue{},
		}

		ll = &linked.LinkedList{1, &linked.LinkedList{2, &linked.LinkedList{Item: 4}}}
	)

	fmt.Printf("matrix summ - %d\n", matrixSumDiag(matrix))
	fmt.Printf("units count - %d\n", units(unitsSlice))
	sliceWorker()
	fmt.Println(Filter([]int{1, 2, 3, 5, 6, 7, 8, 9}, filterBool))
	fmt.Println(matrixLinesUnits(unitsMatrix))
	fmt.Println(changeDiag(matrix))
	fmt.Println(mapEqual(m1, m2))

	// TODO @yevgeniy test stack
	var st stack.Stack
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Pop()
	fmt.Printf("Stack - %v\n", st.Items)

	// Queue
	var q queue.Queue
	q.Enqueu(1)
	q.Enqueu(2)
	q.Enqueu(3)
	q.Deque()
	fmt.Printf("Queue - %v\n", q.Items)

	// Collections
	flag := "stack"

	op, ok := primeCollection[flag]
	if !ok {
		fmt.Println("unknown collection")
		return
	}

	op.Add(1)
	op.Add(2)
	op.Add(3)
	op.Delete()
	fmt.Printf("Collection - %v\n", op.Dump())

	// LinkedList
	fmt.Println("----")
	fmt.Printf("LL print: %v\n", ll)
	fmt.Printf("LL copy rec: %v\n", ll.CopyLLRec())
	fmt.Printf("LL copy manual: %v\n", ll.CopyLLManual())
	ll.NthFromLastRec(2)
	fmt.Printf("(manual) %v ll nth from last: %v\n", 2, ll.NthFromLastManual(2))
	fmt.Printf("Reverse %v\n", ll.Reverse())

	// other
	fmt.Println("----")
	sl := make([]int, 0, 1000)
	sl = append(sl, 1, 2, 3)
	fmt.Println(sl)
	process(sl)
	fmt.Println(sl)

	fmt.Printf("Merge sort: %v", MergeSort([]int{2, 4, 8, 15}, []int{1, 5, 11, 14, 28}))

	///////

}

func process(sl []int) {
	fmt.Printf("-->> %v\n", sl)
	sl = append(sl, 1)
	sl = append(sl, 2)
	sl = append(sl, 3)
	sl[0] = 0
}

func matrixSumDiag(matrix [][]int) int {
	matrixLen := len(matrix)
	var sum int
	for i := 0; i < matrixLen; i++ {
		sum += matrix[i][i] + matrix[i][matrixLen-i-1]
	}
	if matrixLen%2 == 1 {
		sum -= matrix[matrixLen/2][matrixLen/2]
	}
	return sum
}

func units(list []int) int {
	var (
		maxv, unitsCount int
	)
	for i := range list {
		if list[i] == 1 {
			unitsCount += 1
			maxv = max(maxv, unitsCount)
		} else {
			unitsCount = 0
		}
	}
	return maxv
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func sliceWorker() {
	var (
		a = []string{"hello"}
		b = []string{"world", "foo", "bar"}
		_ []int // nil
	)
	a = append(a, b...)
	fmt.Println(a)
}

func Filter(s []int, fn func(int) bool) []int {
	var p []int
	for _, e := range s {
		if fn(e) {
			p = append(p, e)
		}
	}
	return p
}

func filterBool(e int) bool {
	if e%2 == 1 {
		return true
	}
	return false
}

func matrixLinesUnits(matrix [][]int) []int {
	var unitsLines []int
	var sum int
	for i, mEl := range matrix {
		for _, lEl := range mEl {
			if lEl == 1 {
				sum++
			}
		}
		if sum >= 3 {
			unitsLines = append(unitsLines, i)
		}
		sum = 0
	}
	return unitsLines
}

func changeDiag(matrix [][]int) [][]int {
	l := len(matrix)
	for i := range matrix {
		matrix[i][i], matrix[i][l-i-1] = matrix[i][l-i-1], matrix[i][i]
	}
	return matrix
}

func mapEqual(x, y map[int]int) bool {
	if len(x) != len(y) {
		return false
	}
	for i, mx := range x {
		if my, ok := y[i]; !ok || mx != my {
			return false
		}
	}
	return true
}

func MergeSort(a, b []int) (res []int) {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		res = append(res, a[i])
	}

	for ; j < len(b); j++ {
		res = append(res, b[j])
	}
	return
}

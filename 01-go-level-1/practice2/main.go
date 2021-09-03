package main

import (
	"fmt"
	c "practice2/collections"
	ll "practice2/linked"
)

func main() {
	var (
		curColl    = "queue"
		collChoose = map[string]c.Collections{
			"stack": &c.Stack{},
			"queue": &c.Queue{},
		}

		curLL = &ll.LinkedList{1, &ll.LinkedList{3, &ll.LinkedList{6, &ll.LinkedList{Item: 2}}}}

		sl1 = []int{1, 2, 3, 6, 7}
		sl2 = []int{4, 8, 9, 12}
		sl  = []int{-1, 2, 3, -3, 22}

		m1 = map[rune]int{'1': 1, '2': 2}
		m2 = map[rune]int{'1': 1, '2': 2}
	)
	// Coll
	coll, ok := collChoose[curColl]
	if !ok {
		fmt.Println("unknown collection")
		return
	}
	coll.Add(1)
	coll.Add(3)
	coll.Add(6)
	coll.Delete()
	fmt.Printf("Coll %v: %v\n", curColl, coll.Dump())

	// LinkedList
	fmt.Printf("LL print: %v\n", curLL)
	fmt.Printf("CopyLL Rec: %v\n", curLL.CopyLLRec())
	fmt.Printf("CopyLL Manual: %v\n", curLL.CopyLLManual())
	fmt.Printf("Nth Manual: %v\n", curLL.NthFromLastManual(2))
	curLL.NthFromLastRec(2)
	fmt.Println()
	fmt.Printf("LL Reverse: %v\n", curLL.Reverse())

	// Merge
	fmt.Printf("Merge: %v\n", MergeSlice(sl1, sl2))

	// MapEq

	fmt.Printf("Map eq: %v\n", MapEq(m1, m2))

	// Filter

	fmt.Printf("isPos: %v", Filter(sl, isPos))
}

func MergeSlice(a, b []int) (res []int) {
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

func MapEq(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for ak, av := range a {
		if bv, ok := b[ak]; !ok || av != bv {
			return false
		}
	}
	return true
}

func Filter(sl []int, fn func(e int) bool) (res []int) {
	for _, el := range sl {
		if fn(el) {
			res = append(res, el)
		}
	}
	return
}

func isPos(x int) bool {
	return x >= 0
}

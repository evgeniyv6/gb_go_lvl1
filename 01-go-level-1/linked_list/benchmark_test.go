package main

import "testing"

var (
	myLL = &ListNode{"lean1", &ListNode{2,&ListNode{"three", &ListNode{4, nil}}}}
	GlobalRes *ListNode
)

func BenchmarkReverseList(b *testing.B) {
	var res *ListNode
	for i:=0; i < b.N; i++ {
		res = ReverseList(myLL)
	}
	GlobalRes = res
}

func BenchmarkReverseList2(b *testing.B) {
	var res *ListNode
	for i:=0; i < b.N; i++ {
		res = ReverseList2(myLL)
	}
	GlobalRes = res
}

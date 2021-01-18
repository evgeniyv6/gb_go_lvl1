package main

import "fmt"

type ListNode struct {
	Val interface{}
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return "nil"
	}
	if l.Next == nil {
		return fmt.Sprintf("%v -> nil", l.Val)
	}
	return fmt.Sprintf("%v -> %s", l.Val, l.Next)
}

func ReverseList(head *ListNode) *ListNode {
	var l *ListNode
	for head != nil {
		l = &ListNode{head.Val, l}
		head = head.Next
	}
	return l
}

func ReverseList2(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}

func main() {
	var myLL = &ListNode{"1", &ListNode{2,&ListNode{"3", &ListNode{4, nil}}}}
	fmt.Println(myLL)
	fmt.Println(ReverseList(myLL))

	fmt.Println("------")
	fmt.Println(myLL)
	fmt.Println(ReverseList2(myLL))
}

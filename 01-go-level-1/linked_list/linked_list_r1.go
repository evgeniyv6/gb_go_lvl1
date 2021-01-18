package main

import (
	"fmt"
)

type LL struct {
	Val interface {}
	Next *LL
}

func (l *LL) String() string {
	if l == nil {
		return "Empty"
	}
	if l.Next == nil {
		return fmt.Sprintf("%v -> nil", l.Val)
	}
	return fmt.Sprintf("%v -> %s", l.Val, l.Next)
}

func Rev1(head *LL) *LL {
	var l *LL
	for head !=nil {
		l = &LL{head.Val,  l}
		head = head.Next
	}
	return l
}

func Rev2(head *LL) *LL {
	var prev, next *LL
	for head != nil {
		next = head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}

func main() {
	myLL := &LL{1,&LL{"2", &LL{3, &LL{4,nil}}}}
	fmt.Printf("Init linked: %s", myLL)
	fmt.Printf("\nRev1: %s", Rev1(myLL))
	fmt.Printf("\nRev2: %s", Rev2(myLL))
}

package main

import (
	"fmt"
	"sync"
)

type elem interface {}

type StackR1 struct {
	Items []elem
	Lock sync.Mutex
}


func (st *StackR1) Push(el elem) {
	st.Lock.Lock()
	defer st.Lock.Unlock()

	st.Items = append(st.Items, el)
}

func (st *StackR1) Pop() (elem, bool) {
	st.Lock.Lock()
	defer st.Lock.Unlock()

	if len(st.Items) == 0 {
		return 0, false
	}
	l := len(st.Items)
	old := st.Items[l-1]
	st.Items = st.Items[:l-1]
	return old, true

}

func (st *StackR1) Dump() []elem {
	st.Lock.Lock()
	defer st.Lock.Unlock()

	tmpSt := make([]elem, len(st.Items))
	copy(tmpSt, st.Items)
	return tmpSt

}

func (st *StackR1) isEmpty() bool {
	st.Lock.Lock()
	defer st.Lock.Unlock()
	return len(st.Items) == 0
}

func (st *StackR1) Clear() {
	st.Lock.Lock()
	defer st.Lock.Unlock()

	st.Items = nil
}

func (st *StackR1) Peek() elem {
	st.Lock.Lock()
	defer st.Lock.Unlock()

	if len(st.Items) == 0 {
		return nil
	}
	return st.Items[len(st.Items)-1]
}

func main() {
	var mySt StackR1
	mySt.Push(1); mySt.Push("2"); mySt.Push(3)
	fmt.Println(mySt.Dump())
	fmt.Println(mySt.Pop())
	fmt.Println(mySt.Dump())
	fmt.Printf("is empty: %v", mySt.isEmpty())
	fmt.Printf("\nPeek: %v", mySt.Peek())
}
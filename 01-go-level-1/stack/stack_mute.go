package main

import (
	"log"
	"sync"
)

type Item interface {}

type Stackm struct {
	items []Item
	mute sync.Mutex
}

func (st *Stackm) Push(el Item) {
	st.mute.Lock()
	defer st.mute.Unlock()
	st.items = append(st.items, el)
}

func (st *Stackm) isEmpty() bool {
	st.mute.Lock()
	defer st.mute.Unlock()
	return len(st.items) == 0
}

func (st *Stackm) Pop() Item {
	st.mute.Lock()
	defer st.mute.Unlock()

	l := len(st.items)

	if len(st.items) == 0 {
		return nil
	}
	oldItem := st.items[l-1]
	st.items = st.items[:l - 1]
	return oldItem
}

func (st *Stackm) Clear() {
	st.mute.Lock()
	defer st.mute.Unlock()

	st.items = nil
}


func (st *Stackm) Dump() []Item {
	st.mute.Lock()
	defer st.mute.Unlock()

	var tmpStack = make([]Item,len(st.items))
	copy(tmpStack, st.items)
	return tmpStack

}

func (st *Stackm) Peek() Item {
	st.mute.Lock()
	defer st.mute.Unlock()

	if len(st.items) == 0 {
		return nil
	}
	return st.items[len(st.items)-1]
}

func main() {
	var myStack Stackm
	myStack.Push(1); myStack.Push(2); myStack.Push("three")
	log.Println(myStack.Dump())
	log.Println(myStack.Pop())
	log.Println(myStack.Dump())
	log.Println(myStack.Peek())
	myStack.Clear()
	log.Println(myStack.Dump())


}
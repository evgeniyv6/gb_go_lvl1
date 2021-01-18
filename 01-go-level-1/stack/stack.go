package main

import "log"

type Stack []int

func StackIsEmpty(st *Stack) bool {
	return len(*st) == 0
}

func (st *Stack) StackPush(el int) {
	*st = append(*st, el)
}

func (st *Stack) StackPop() (int, bool) {
	if StackIsEmpty(st) {
		return 0, false
	} else {
		l := len(*st)
		old := (*st)[l-1]
		*st = (*st)[:l - 1]
		return old, true
	}
}

func main() {
	var myStack = make(Stack,0)
	log.Println(myStack)
	myStack.StackPush(1); myStack.StackPush(2); myStack.StackPush(3)
	log.Println(myStack)
	if old,ok := myStack.StackPop(); ok {
		log.Println(old)
	}
	log.Println(myStack)
}
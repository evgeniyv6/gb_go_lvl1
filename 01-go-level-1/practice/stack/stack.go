package stack

import "sync"

type Item interface{}

type Stack struct {
	Items []Item
	sync.Mutex
}

func (st *Stack) Push(elem Item) {
	st.Lock()
	defer st.Unlock()

	st.Items = append(st.Items, elem)
}

func (st *Stack) Pop() Item {
	st.Lock()
	defer st.Unlock()

	l := len(st.Items)

	if len(st.Items) == 0 {
		return nil
	}

	oldVal := st.Items[l-1]
	st.Items = st.Items[:l-1]
	return oldVal
}

func (st *Stack) isEmpty() bool {
	st.Lock()
	defer st.Unlock()

	return len(st.Items) == 0
}

func (st *Stack) Clear() {
	st.Lock()
	defer st.Unlock()

	st.Items = nil
}

func (st *Stack) Dump() []Item {
	st.Lock()
	defer st.Unlock()

	tmpStack := make([]Item, len(st.Items))
	copy(tmpStack, st.Items)

	return tmpStack
}

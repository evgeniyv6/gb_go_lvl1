package collections

import "sync"

type Stack struct {
	sync.Mutex
	Items []Item
}

func (st *Stack) Add(el Item) {
	st.Lock()
	defer st.Unlock()

	st.Items = append(st.Items, el)
}

func (st *Stack) Delete() Item {
	st.Lock()
	defer st.Unlock()
	l := len(st.Items)
	old := st.Items[l-1]
	st.Items = st.Items[:l-1]
	return old
}

func (st *Stack) Dump() []Item {
	st.Lock()
	defer st.Unlock()

	tmp := make([]Item, len(st.Items))
	copy(tmp, st.Items)
	return tmp
}

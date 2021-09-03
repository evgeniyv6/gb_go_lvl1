package collections

import "sync"

type Stack struct {
	Items []Item
	sync.Mutex
}

func (st *Stack) Add(it Item) {
	st.Lock()
	defer st.Unlock()

	st.Items = append(st.Items, it)
}

func (st *Stack) Delete() Item {
	st.Lock()
	defer st.Unlock()

	l := len(st.Items)

	oldIt := st.Items[l-1]
	st.Items = st.Items[:l-1]

	return oldIt
}

func (st *Stack) Dump() []Item {
	st.Lock()
	defer st.Unlock()

	tmpSl := make([]Item, len(st.Items))
	copy(tmpSl, st.Items)
	return tmpSl
}

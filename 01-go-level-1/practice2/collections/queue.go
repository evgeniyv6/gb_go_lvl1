package collections

import "sync"

type Queue struct {
	sync.Mutex
	Items []Item
}

func (q *Queue) Add(el Item) {
	q.Lock()
	defer q.Unlock()

	q.Items = append(q.Items, el)
}

func (q *Queue) Delete() Item {
	q.Lock()
	defer q.Unlock()

	l := len(q.Items)
	old := q.Items[0]
	q.Items = q.Items[1:l]
	return old
}

func (q *Queue) Dump() []Item {
	q.Lock()
	defer q.Unlock()

	tmp := make([]Item, len(q.Items))
	copy(tmp, q.Items)
	return tmp
}

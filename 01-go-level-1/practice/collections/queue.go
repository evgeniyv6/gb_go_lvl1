package collections

import "sync"

type Queue struct {
	Items []Item
	sync.Mutex
}

func (q *Queue) Add(it Item) {
	q.Lock()
	defer q.Unlock()

	q.Items = append(q.Items, it)
}

func (q *Queue) Delete() Item {
	q.Lock()
	defer q.Unlock()

	oldIt := q.Items[0]
	q.Items = q.Items[1:]

	return oldIt
}

func (q *Queue) Dump() []Item {
	q.Lock()
	defer q.Unlock()

	tmpSl := make([]Item, len(q.Items))
	copy(tmpSl, q.Items)

	return q.Items
}

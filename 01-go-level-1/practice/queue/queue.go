package queue

import "sync"

type Item interface {
}

type Queue struct {
	Items []Item
	sync.Mutex
}

func (q *Queue) Enqueu(e Item) {
	q.Lock()
	defer q.Unlock()

	q.Items = append(q.Items, e)
}

func (q *Queue) Deque() Item {
	q.Lock()
	defer q.Unlock()

	oldVal := q.Items[0]
	q.Items = q.Items[1:]
	return oldVal
}

func (q *Queue) Dump() []Item {
	q.Lock()
	defer q.Unlock()

	tmpQ := make([]Item, len(q.Items))

	copy(tmpQ, q.Items)

	return tmpQ
}

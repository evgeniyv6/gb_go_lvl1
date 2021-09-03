package linked

import "fmt"

type Item interface {
}

type LinkedList struct {
	Item Item
	Next *LinkedList
}

func (ll *LinkedList) String() string {
	if ll == nil {
		return "nil"
	}
	if ll.Next == nil {
		return fmt.Sprintf("%v -> nil", ll.Item)
	}

	return fmt.Sprintf("%v -> %v ", ll.Item, ll.Next)
}

func (head *LinkedList) CopyLLRec() *LinkedList {
	if head == nil {
		return head
	}
	newList := &LinkedList{Item: head.Item}
	newList.Next = head.Next.CopyLLRec()
	return newList
}

func (head *LinkedList) CopyLLManual() *LinkedList {
	cur := head
	res := new(LinkedList)
	tmp := res
	for cur != nil {
		tmp.Next = &LinkedList{cur.Item, tmp.Next}
		cur = cur.Next
		tmp = tmp.Next
	}
	return res.Next
}

func (head *LinkedList) NthFromLastRec(k Item) Item {
	if head == nil {
		return 0
	}
	i := head.Next.NthFromLastRec(k).(int) + 1
	if i == k {
		fmt.Printf("(rec) %v elem from last is %v\n", k, head.Item)
	}
	return i
}

func (head *LinkedList) NthFromLastManual(k Item) Item {
	p1, p2 := head, head
	for i := 0; i < k.(int)-1; i++ {
		p1 = p1.Next
	}
	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2.Item
}

func (head *LinkedList) Reverse() *LinkedList {
	cur := head
	prev := new(LinkedList)
	for {
		if cur == nil {
			break
		}
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

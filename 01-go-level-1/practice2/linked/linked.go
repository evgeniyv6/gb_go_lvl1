package linked

import "fmt"

type Item int

type LinkedList struct {
	Item Item
	Next *LinkedList
}

func (head *LinkedList) String() string {
	if head == nil {
		return "nil"
	}
	if head.Next == nil {
		return fmt.Sprintf("%v -> nil", head.Item)
	}
	return fmt.Sprintf("%v -> %v", head.Item, head.Next)
}

func (head *LinkedList) CopyLLRec() *LinkedList {
	if head == nil {
		return nil
	}
	newLL := &LinkedList{Item: head.Item}
	newLL.Next = head.Next.CopyLLRec()
	return newLL
}

func (head *LinkedList) CopyLLManual() *LinkedList {
	cur := head
	res := new(LinkedList)
	tmp := res
	for cur != nil {
		tmp.Next = &LinkedList{cur.Item, tmp.Next}
		tmp = tmp.Next
		cur = cur.Next
	}
	return res.Next
}

func (head *LinkedList) NthFromLastManual(k int) Item {
	p1, p2 := head, head
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2.Item
}

func (head *LinkedList) NthFromLastRec(k int) int {
	if head == nil {
		return 0
	}
	i := head.Next.NthFromLastRec(k) + 1
	if i == k {
		fmt.Printf("nth from last rec: %v", head.Item)
	}
	return i
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

package link

import "errors"

/*
双向链表（带哨兵）
*/
type DoublyLinkedListSentinel struct {
	head *pNode
	tail *pNode
}

func NewDoublyLinkedListSentinel() *DoublyLinkedListSentinel {
	head := &pNode{value: 666}
	tail := &pNode{value: 888}
	head.next = tail
	tail.prev = head
	return &DoublyLinkedListSentinel{head: head, tail: tail}
}

func (d *DoublyLinkedListSentinel) findNode(index int) *pNode {
	i := -1
	for p := d.head; p != d.tail; p = p.next {
		if i == index {
			return p
		}
		i++
	}
	return nil
}

func (d *DoublyLinkedListSentinel) AddFirst(val int) {
	d.Insert(0, val)
}

func (d *DoublyLinkedListSentinel) Insert(index, val int) error {
	prev := d.findNode(index - 1)
	if prev == nil {
		return errors.New("index 不合法")
	}
	next := prev.next
	inserted := &pNode{prev: prev, value: val, next: next}
	prev.next = inserted
	next.prev = inserted
	return nil
}

func (d *DoublyLinkedListSentinel) AddLast(val int) {
	prev := d.tail.prev
	added := &pNode{prev: prev, value: val, next: d.tail}
	prev.next = added
	d.tail.prev = added
}

func (d *DoublyLinkedListSentinel) RemoveFirst() {
	if d.head.next != d.tail {
		removed := d.head.next
		next := removed.next
		d.head.next = next
		next.prev = d.head
	}
}

func (d *DoublyLinkedListSentinel) RemoveLast() {
	if d.tail.prev != d.head {
		removed := d.tail.prev
		prev := removed.prev
		prev.next = d.tail
		d.tail.prev = prev
	}
}

func (d *DoublyLinkedListSentinel) Remove(index int) error {
	prev := d.findNode(index - 1)
	if prev == nil {
		return errors.New("index不合法")
	}
	removed := prev.next
	if removed == d.tail {
		return errors.New("index不合法")
	}
	next := removed.next

	prev.next = next
	next.prev = prev
	return nil
}

func (d *DoublyLinkedListSentinel) Each(f func(i, v int)) {
	for i, p := 0, d.head.next; p != d.tail; p = p.next {
		f(i, p.value)
	}
}

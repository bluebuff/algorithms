package link

/*
双向环形链表（带哨兵）
*/
type DoublyRingLinkedListSentinel struct {
	sentinel *pNode
}

func NewDoublyRingLinkedListSentinel() *DoublyRingLinkedListSentinel {
	sentinel := &pNode{value: 666}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return &DoublyRingLinkedListSentinel{sentinel: sentinel}
}

/*
头部添加元素
*/
func (d *DoublyRingLinkedListSentinel) AddFirst(val int) {
	prev := d.sentinel
	next := d.sentinel.next
	added := &pNode{prev: prev, value: val, next: next}
	prev.next = added
	next.prev = added
}

/*
尾部添加元素
*/
func (d *DoublyRingLinkedListSentinel) AddLast(val int) {
	prev := d.sentinel.prev
	next := d.sentinel
	added := &pNode{prev: prev, value: val, next: next}
	prev.next = added
	next.prev = added
}

/*
头部移除元素
*/
func (d *DoublyRingLinkedListSentinel) RemoveFirst() {
	removed := d.sentinel.next
	if removed == d.sentinel {
		return
	}
	next := removed.next
	d.sentinel.next = next
	next.prev = d.sentinel
}

/*
尾部移除元素
*/
func (d *DoublyRingLinkedListSentinel) RemoveLast() {
	removed := d.sentinel.prev
	if removed == d.sentinel {
		return
	}
	prev := removed.prev
	d.sentinel.prev = prev
	prev.next = d.sentinel
}

/*
根据值删除元素
*/
func (d *DoublyRingLinkedListSentinel) RemoveByValue(val int) {
	removed := d.findNodeByValue(val)
	if removed == nil {
		return
	}
	prev := removed.prev
	next := removed.next
	prev.next = next
	next.prev = prev
}

func (d *DoublyRingLinkedListSentinel) findNodeByValue(val int) *pNode {
	for p := d.sentinel.next; p != d.sentinel; p = p.next {
		if p.value == val {
			return p
		}
	}
	return nil
}

func (d *DoublyRingLinkedListSentinel) Each(f func(i, v int)) {
	for i, p := 0, d.sentinel.next; p != d.sentinel; p = p.next {
		f(i, p.value)
		i++
	}
}

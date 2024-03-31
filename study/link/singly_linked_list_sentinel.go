package link

import "errors"

/*
单向链表（带哨兵）
*/
type SinglyLinkedListSentinel struct {
	head *node // 头指针
}

func NewSinglyLinkedListSentinel() *SinglyLinkedListSentinel {
	return &SinglyLinkedListSentinel{head: &node{value: 666}} // 哨兵
}

/*
向链表头部添加
*/
func (s *SinglyLinkedListSentinel) AddFirst(val int) {
	s.head.next = &node{value: val, next: s.head.next}
}

/*
向链表尾部添加
*/
func (s *SinglyLinkedListSentinel) AddLast(val int) {
	last := s.findLast()
	last.next = &node{value: val}
}

/*
查找尾节点
*/
func (s *SinglyLinkedListSentinel) findLast() *node {
	p := s.head
	for ; p.next != nil; p = p.next {

	}
	return p
}

/*
根据索引查询索引值
*/
func (s *SinglyLinkedListSentinel) Get(index int) (int, bool) {
	node := s.findNodeByIndex(index)
	if node == nil {
		return 0, false
	}
	return node.value, true
}

/*
根据索引查询结点
*/
func (s *SinglyLinkedListSentinel) findNodeByIndex(index int) *node {
	for i, p := 0, s.head.next; p != nil; p = p.next {
		if i == index {
			return p
		}
		i++
	}
	return nil
}

/*
向索引位置插入
*/
func (s *SinglyLinkedListSentinel) Insert(index int, value int) error {
	if index == 0 {
		s.AddFirst(value)
		return nil
	}
	pre := s.findNodeByIndex(index - 1)
	if pre == nil {
		return errors.New("index 不合法")
	}
	pre.next = &node{value: value, next: pre.next}
	return nil
}

/*
删除头部元素
*/
func (s *SinglyLinkedListSentinel) RemoveFirst() {
	if s.head.next == nil {
		return
	}
	s.head.next = s.head.next.next
}

/*
删除指定索引元素
*/
func (s *SinglyLinkedListSentinel) Remove(index int) error {
	if index == 0 {
		s.RemoveFirst()
		return nil
	}
	pre := s.findNodeByIndex(index - 1)
	if pre == nil {
		return errors.New("index 不合法")
	}
	removed := pre.next
	if removed == nil {
		return errors.New("index 不合法")
	}
	pre.next = removed.next
	return nil
}

/*
循环遍历
*/
func (s *SinglyLinkedListSentinel) Each(f func(i, v int)) {
	i := 0
	for p := s.head.next; p != nil; p = p.next {
		f(i, p.value)
		i++
	}
}

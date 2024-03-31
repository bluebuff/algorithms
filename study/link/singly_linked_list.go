package link

import "errors"

/*
单向链表
*/
type SinglyLinkedList struct {
	head *node // 头指针
}

/*
向链表头部添加
*/
func (s *SinglyLinkedList) AddFirst(val int) {
	s.head = &node{value: val, next: s.head}
}

/*
向链表尾部添加
*/
func (s *SinglyLinkedList) AddLast(val int) {
	last := s.findLast()
	if last == nil {
		s.AddFirst(val)
		return
	}
	last.next = &node{value: val}
}

/*
查找尾节点
*/
func (s *SinglyLinkedList) findLast() *node {
	if s.head == nil {
		return nil
	}
	p := s.head
	for ; p.next != nil; p = p.next {

	}
	return p
}

/*
根据索引查询索引值
*/
func (s *SinglyLinkedList) Get(index int) (int, bool) {
	node := s.findNodeByIndex(index)
	if node == nil {
		return 0, false
	}
	return node.value, true
}

/*
根据索引查询结点
*/
func (s *SinglyLinkedList) findNodeByIndex(index int) *node {
	for i, p := 0, s.head; p != nil; p = p.next {
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
func (s *SinglyLinkedList) Insert(index int, value int) error {
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
func (s *SinglyLinkedList) RemoveFirst() {
	if s.head == nil {
		return
	}
	s.head = s.head.next
}

/*
删除指定索引元素
*/
func (s *SinglyLinkedList) Remove(index int) error {
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
func (s *SinglyLinkedList) Each(f func(i, v int)) {
	i := 0
	for p := s.head; p != nil; p = p.next {
		f(i, p.value)
		i++
	}
}

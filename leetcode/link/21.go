package link

import (
	"algorithms/study/entity"
)

func mergeTwoList(node1, node2 *entity.Node) *entity.Node {
	s := &entity.Node{Value: -1, Next: nil}
	p := s
	p1, p2 := node1, node2
	for ; p1 != nil && p2 != nil; p = p.Next {
		if p1.Value < p2.Value {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
	}
	if p1 == nil {
		p.Next = p2
	}
	if p2 == nil {
		p.Next = p1
	}
	return s.Next
}

func mergeTwoList2(node1, node2 *entity.Node) *entity.Node {
	if node2 == nil {
		return node1
	}
	if node1 == nil {
		return node2
	}
	if node1.Value < node2.Value {
		node1.Next = mergeTwoList2(node1.Next, node2)
		return node1
	} else {
		node2.Next = mergeTwoList2(node1, node2.Next)
		return node2
	}
}

func mergeTwoList3(node ...*entity.Node) *entity.Node {
	if len(node) == 0 {
		return nil
	}
	return split(0, len(node)-1, node...)
}

func split(i, j int, node ...*entity.Node) *entity.Node {
	if i == j {
		return node[i]
	}
	m := (i + j) >> 1
	left := split(i, m, node...)
	right := split(m+1, j, node...)
	return mergeTwoList(left, right)
}

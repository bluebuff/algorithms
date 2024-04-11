package link

import "algorithms/study/entity"

func removeByValue(node *entity.Node, val int) *entity.Node {
	s := &entity.Node{Value: -1, Next: node}
	p1 := s
	p2 := s.Next
	for p2 != nil {
		if p2.Value == val {
			// 删除
			p1.Next = p2.Next
			p2 = p2.Next
		} else {
			// 移动
			p1 = p1.Next
			p2 = p2.Next
		}
	}
	return s.Next
}

func removeRecursion(node *entity.Node, val int) *entity.Node {
	if node == nil {
		return nil
	}
	if node.Value == val {
		return removeRecursion(node.Next, val)
	} else {
		node.Next = removeRecursion(node.Next, val)
		return node
	}
}

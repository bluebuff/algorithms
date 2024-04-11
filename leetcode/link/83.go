package link

import (
	"algorithms/study/entity"
)

func deleteDuplicatges(node *entity.Node) {
	if node == nil || node.Next == nil {
		return
	}
	n1 := node
	for n2 := n1.Next; n2 != nil; n2 = n2.Next {
		if n1.Value == n2.Value {
			n1.Next = n2.Next
		} else {
			n1 = n1.Next
		}
	}
}

func deleteDuplicatgesRecursion(node *entity.Node) *entity.Node {
	if node == nil || node.Next == nil {
		return node
	}
	if node.Value == node.Next.Value {
		return deleteDuplicatgesRecursion(node.Next)
	} else {
		node.Next = deleteDuplicatgesRecursion(node.Next)
		return node
	}
}

func deleteDuplicates(node *entity.Node) *entity.Node {
	if node == nil || node.Next == nil {
		return node
	}
	if node.Value == node.Next.Value {
		x := node.Next.Next
		for x != nil && node.Value == x.Value {
			x = x.Next
		}
		return deleteDuplicates(x)
	} else {
		node.Next = deleteDuplicates(node.Next)
		return node
	}
}

func deleteDuplicates2(node *entity.Node) *entity.Node {
	if node == nil || node.Next == nil {
		return node
	}
	s := &entity.Node{Value: -1, Next: node}
	p1 := s
	var p2, p3 *entity.Node
	for {
		if p2 = p1.Next; p2 == nil {
			break
		}
		if p3 = p2.Next; p3 == nil {
			break
		}
		if p2.Value == p3.Value {
			for p3 = p3.Next; p3 != nil && p3.Value == p2.Value; p3 = p3.Next {
			}
			p1.Next = p3
		} else {
			p1 = p1.Next
		}
	}
	return s.Next
}

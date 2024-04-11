package link

import "algorithms/study/entity"

func middleNode(node *entity.Node) *entity.Node {
	p1, p2 := node, node
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1
}

package link

import (
	"algorithms/study/entity"
)

func isPalindrome(node *entity.Node) bool {
	middle := middleNode(node)
	newNode := reverse(middle)
	for newNode != nil {
		if newNode.Value != node.Value {
			return false
		}
		newNode = newNode.Next
		node = node.Next
	}
	return true
}

func reverse(node *entity.Node) *entity.Node {
	var p *entity.Node
	for node != nil {
		o2 := node.Next
		node.Next = p
		p = node
		node = o2
	}
	return p
}

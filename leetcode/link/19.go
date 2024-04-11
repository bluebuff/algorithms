package link

import (
	"algorithms/study/entity"
	"fmt"
)

func removeNthFromEnd(node *entity.Node, n int) *entity.Node {
	recurisionNth(node, n)
	return node
}

func recurisionNth(node *entity.Node, n int) int {
	if node == nil {
		return 0
	}
	nth := recurisionNth(node.Next, n)
	fmt.Println(node.Value, nth)
	if nth == n {
		// 3 -> 4 -> 5
		// p  p.Next p.Next.Next
		node.Next = node.Next.Next
	}
	return nth + 1
}

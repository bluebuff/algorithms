package searchbinary

import (
	"algorithms/study/entity"
)

func reverseLink1(node *entity.Node) *entity.Node {
	var n *entity.Node
	for p := node; p != nil; p = p.Next {
		n = &entity.Node{Value: p.Value, Next: n}
	}
	return n
}

func reverseLink2(node *entity.Node) *entity.Node {
	list1 := &List{node: node}
	list2 := &List{node: nil}
	for {
		first := list1.RemoveFirst()
		if first == nil {
			break
		}
		list2.AddFirst(first)
	}
	return list2.node
}

func reverseLink3(node *entity.Node) *entity.Node {
	if node == nil || node.Next == nil {
		return node
	}
	last := reverseLink3(node.Next)
	// node node.Next
	//  4 -> 5
	//  --------
	//  5 -> 4 -> nil
	//  node.Next.Next  = node
	//  node.Next = nil
	node.Next.Next = node
	node.Next = nil
	return last
}

/*
	 n1
	 o1
	 1 -> 2 -> 3 -> 4 -> 5 -> nil

	 n1
	      o1
	 2 -> 1 -> 3 -> 4 -> 5 -> nil

	 n1
	           o1
	 3 -> 2 -> 1 -> 4 -> 5 -> nil

	 n1
			        o1
	 4 -> 3 -> 2 -> 1 -> 5 -> nil

	 n1
	                     o1
	 5 -> 4 -> 3 -> 2 -> 1 -> nil
*/
func reverseLink4(node *entity.Node) *entity.Node {
	n1, o1 := node, node
	for o1.Next != nil {
		removed := o1.Next
		o1.Next = removed.Next
		removed.Next = n1
		n1 = removed
	}
	return n1
}

/*
n1    o1      o2
---   ---  ->  2 -> 3 -> 4 -> 5 -> nil
nil    1

o1      n1      o2
--- -> ---  ->  2 -> 3 -> 4 -> 5 -> nil
1      nil

n1o1     o2
---  ->  2 -> 3 -> 4 -> 5 -> nil
1

n1       o1o2
---  ->  2 -> 3 -> 4 -> 5 -> nil
1
*/
func reverseLink5(node *entity.Node) *entity.Node {
	if node == nil || node.Next == nil {
		return node
	}
	var n1 *entity.Node
	o1 := node
	for o1 != nil {
		o2 := o1.Next
		o1.Next = n1
		n1 = o1
		o1 = o2
	}
	return n1
}

type List struct {
	node *entity.Node
}

func (l *List) AddFirst(node *entity.Node) {
	node.Next = l.node
	l.node = node
}

func (l *List) RemoveFirst() *entity.Node {
	if l.node != nil {
		first := l.node
		l.node = first.Next
		return first
	}
	return nil
}

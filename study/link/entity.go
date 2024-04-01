package link

type node struct {
	value int
	next  *node
}

type pNode struct {
	prev  *pNode
	value int
	next  *pNode
}

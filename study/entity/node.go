package entity

import (
	"bytes"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := 0, n; p != nil; p = p.Next {
		if i > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprint(p.Value))
		i++
	}
	buf.WriteString("]")
	return buf.String()
}

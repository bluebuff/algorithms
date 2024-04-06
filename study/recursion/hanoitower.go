package recursion

import (
	"algorithms/study/link"
	"fmt"
)

type Hanoitower struct {
	a *link.DoublyLinkedListSentinel
	b *link.DoublyLinkedListSentinel
	c *link.DoublyLinkedListSentinel
}

func NewHanoitower() *Hanoitower {
	a := link.NewDoublyLinkedListSentinel()
	b := link.NewDoublyLinkedListSentinel()
	c := link.NewDoublyLinkedListSentinel()
	return &Hanoitower{
		a: a,
		b: b,
		c: c,
	}
}

func (h *Hanoitower) Move(n int) {
	for i := n; i > 0; i-- {
		h.a.AddLast(i)
	}
	print(h)
	move(n, h.a, h.b, h.c, h)
}

func move(n int, a, b, c *link.DoublyLinkedListSentinel, h *Hanoitower) {
	if n == 0 {
		return
	}
	move(n-1, a, c, b, h)     // 把n-1个盘子，由a借c移动到b
	c.AddLast(a.RemoveLast()) // 把最后的盘子，由a移动到c
	move(n-1, b, a, c, h)     // 把n-1个盘子，由b借a移动到c
}

func (h *Hanoitower) Print() {
	print(h)
}

func print(h *Hanoitower) {
	fmt.Println(h.a.String())
	fmt.Println(h.b.String())
	fmt.Println(h.c.String())
	fmt.Println("---")
}

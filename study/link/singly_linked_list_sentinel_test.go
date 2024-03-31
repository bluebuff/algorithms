package link

import (
	"fmt"
	"testing"
)

func TestSinglyLinkedListSentinel_AddFirst(t *testing.T) {
	link := NewSinglyLinkedListSentinel()

	link.AddFirst(1)
	link.AddFirst(2)
	link.AddFirst(3)
	link.AddFirst(4)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

}

func TestSinglyLinkedListSentinel_AddLast(t *testing.T) {
	link := NewSinglyLinkedListSentinel()

	link.AddLast(5)
	link.AddLast(6)
	link.AddLast(7)
	link.AddLast(8)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})
}

func TestSinglyLinkedListSentinel_Get(t *testing.T) {
	link := NewSinglyLinkedListSentinel()

	link.AddLast(5)
	link.AddLast(6)
	link.AddLast(7)
	link.AddLast(8)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

	fmt.Println(link.Get(0))
	fmt.Println(link.Get(1))
	fmt.Println(link.Get(2))
	fmt.Println(link.Get(3))
	fmt.Println(link.Get(5))
}

func TestSinglyLinkedListSentinel_Insert(t *testing.T) {
	link := NewSinglyLinkedListSentinel()

	link.AddLast(1)
	link.AddLast(2)
	link.AddLast(3)
	link.AddLast(4)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

	fmt.Println("---")

	link.Insert(4, 5)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})
}

func TestSinglyLinkedListSentinel_RemoveFirst(t *testing.T) {
	link := NewSinglyLinkedListSentinel()

	link.AddLast(1)
	link.AddLast(2)
	link.AddLast(3)
	link.AddLast(4)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

	fmt.Println("---")

	link.RemoveFirst()

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

}

func TestSinglyLinkedListSentinel_Remove(t *testing.T) {
	link := NewSinglyLinkedListSentinel()

	link.AddLast(1)
	link.AddLast(2)
	link.AddLast(3)
	link.AddLast(4)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

	fmt.Println("---")

	link.Remove(0)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

	fmt.Println("---")

	link.Remove(1)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

}

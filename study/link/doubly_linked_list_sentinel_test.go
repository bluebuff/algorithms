package link

import (
	"fmt"
	"testing"
)

func TestDoublyLinkedListSentinel_AddFirst(t *testing.T) {
	link := NewDoublyLinkedListSentinel()

	link.AddFirst(1)
	link.AddFirst(2)
	link.AddFirst(3)
	link.AddFirst(4)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})
}

func TestDoublyLinkedListSentinel_Insert(t *testing.T) {
	link := NewDoublyLinkedListSentinel()

	link.AddFirst(1)
	link.AddFirst(2)
	link.AddFirst(3)
	link.AddFirst(4)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

	fmt.Println("---")

	link.Insert(2, 5)
	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

}

func TestDoublyLinkedListSentinel_AddLast(t *testing.T) {
	link := NewDoublyLinkedListSentinel()

	link.AddFirst(1)
	link.AddFirst(2)
	link.AddFirst(3)
	link.AddFirst(4)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

	fmt.Println("---")

	link.AddLast(5)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

}

func TestDoublyLinkedListSentinel_RemoveFirst(t *testing.T) {
	link := NewDoublyLinkedListSentinel()

	link.AddFirst(1)
	link.AddFirst(2)
	link.AddFirst(3)
	link.AddFirst(4)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

	fmt.Println("---")

	link.RemoveFirst()

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

}

func TestDoublyLinkedListSentinel_RemoveLast(t *testing.T) {
	link := NewDoublyLinkedListSentinel()

	link.AddFirst(1)
	link.AddFirst(2)
	link.AddFirst(3)
	link.AddFirst(4)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

	fmt.Println("---")

	link.RemoveLast()

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

}

func TestDoublyLinkedListSentinel_Remove(t *testing.T) {
	link := NewDoublyLinkedListSentinel()

	link.AddFirst(1)
	link.AddFirst(2)
	link.AddFirst(3)
	link.AddFirst(4)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

	fmt.Println("---")

	link.Remove(2)

	link.Remove(0)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

}

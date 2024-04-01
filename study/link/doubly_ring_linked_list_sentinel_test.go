package link

import (
	"fmt"
	"testing"
)

func TestDoublyRingLinkedListSentinel_AddFirst(t *testing.T) {
	link := NewDoublyRingLinkedListSentinel()

	link.AddFirst(1)
	link.AddFirst(2)
	link.AddFirst(3)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})
}

func TestDoublyRingLinkedListSentinel_AddLast(t *testing.T) {
	link := NewDoublyRingLinkedListSentinel()

	link.AddLast(1)
	link.AddLast(2)
	link.AddLast(3)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})
}

func TestDoublyRingLinkedListSentinel_RemoveFirst(t *testing.T) {
	link := NewDoublyRingLinkedListSentinel()

	link.AddLast(1)
	link.AddLast(2)
	link.AddLast(3)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

	link.RemoveFirst()

	fmt.Println("---")

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})
}

func TestDoublyRingLinkedListSentinel_RemoveLast(t *testing.T) {
	link := NewDoublyRingLinkedListSentinel()

	link.AddLast(1)
	link.AddLast(2)
	link.AddLast(3)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

	link.RemoveLast()

	fmt.Println("---")

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})
}

func TestDoublyRingLinkedListSentinel_RemoveByValue(t *testing.T) {
	link := NewDoublyRingLinkedListSentinel()

	link.AddLast(1)
	link.AddLast(2)
	link.AddLast(3)

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})

	link.RemoveByValue(2)

	fmt.Println("---")

	link.Each(func(i, v int) {
		fmt.Println(i, v)
	})
}

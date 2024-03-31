package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	d := NewDynamicArray()
	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)
	d.AddLast(4)
	d.AddLast(5)
	d.Add(2, 2)
	d.Each(func(v int) {
		fmt.Println(v)
	})

	d.Remove(2)

	fmt.Println("-----")

	d.Each(func(v int) {
		fmt.Println(v)
	})

	d.AddLast(6)

	d.AddLast(7)

	fmt.Println("-----")

	d.Each(func(v int) {
		fmt.Println(v)
	})

	d.Add(0, 0)
	d.Add(0, 0)
	d.AddLast(8)
	d.Remove(0)
	fmt.Println("-----")

	d.Each(func(v int) {
		fmt.Println(v)
	})

	d.Add(10, 9)

	fmt.Println("-----")

	d.Each(func(v int) {
		fmt.Println(v)
	})
}

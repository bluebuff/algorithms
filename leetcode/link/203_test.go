package link

import (
	"algorithms/study/entity"
	"fmt"
	"testing"
)

func TestRemoveByValue(t *testing.T) {
	o1 := &entity.Node{Value: 5, Next: nil}
	o2 := &entity.Node{Value: 4, Next: o1}
	o3 := &entity.Node{Value: 5, Next: o2}
	o4 := &entity.Node{Value: 2, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}

	n5 := removeByValue(o5, 5)

	fmt.Println(n5)
}

func TestRemoveRecursion(t *testing.T) {
	o1 := &entity.Node{Value: 5, Next: nil}
	o2 := &entity.Node{Value: 4, Next: o1}
	o3 := &entity.Node{Value: 3, Next: o2}
	o4 := &entity.Node{Value: 2, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}

	n5 := removeRecursion(o5, 2)

	fmt.Println(n5)
}

package link

import (
	"algorithms/study/entity"
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	o1 := &entity.Node{Value: 3, Next: nil}
	o2 := &entity.Node{Value: 2, Next: o1}
	o3 := &entity.Node{Value: 2, Next: o2}
	o4 := &entity.Node{Value: 1, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}
	deleteDuplicatges(o5)
	fmt.Println(o5)
}

func TestDeleyeDuplicateRecursion(t *testing.T) {
	o1 := &entity.Node{Value: 3, Next: nil}
	o2 := &entity.Node{Value: 2, Next: o1}
	o3 := &entity.Node{Value: 2, Next: o2}
	o4 := &entity.Node{Value: 1, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}
	n5 := deleteDuplicatgesRecursion(o5)
	fmt.Println(n5)
}

func TestDeleyeDuplicate(t *testing.T) {
	o1 := &entity.Node{Value: 3, Next: nil}
	o2 := &entity.Node{Value: 2, Next: o1}
	o3 := &entity.Node{Value: 1, Next: o2}
	o4 := &entity.Node{Value: 1, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}
	n5 := deleteDuplicates(o5)
	fmt.Println(n5)
}

func TestDeleyeDuplicate2(t *testing.T) {
	o1 := &entity.Node{Value: 3, Next: nil}
	o2 := &entity.Node{Value: 2, Next: o1}
	o3 := &entity.Node{Value: 1, Next: o2}
	o4 := &entity.Node{Value: 1, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}
	n5 := deleteDuplicates2(o5)
	fmt.Println(n5)
}

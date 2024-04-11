package link

import (
	"algorithms/study/entity"
	"fmt"
	"testing"
)

func TestMergeTwoList(t *testing.T) {
	o1 := &entity.Node{Value: 9, Next: nil}
	o2 := &entity.Node{Value: 7, Next: o1}
	o3 := &entity.Node{Value: 5, Next: o2}
	o4 := &entity.Node{Value: 3, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}

	o6 := &entity.Node{Value: 8, Next: nil}
	o7 := &entity.Node{Value: 6, Next: o6}
	o8 := &entity.Node{Value: 4, Next: o7}
	o9 := &entity.Node{Value: 2, Next: o8}

	list := mergeTwoList(o5, o9)

	fmt.Println(list)
}

func TestMergeTwoList2(t *testing.T) {
	o1 := &entity.Node{Value: 9, Next: nil}
	o2 := &entity.Node{Value: 7, Next: o1}
	o3 := &entity.Node{Value: 5, Next: o2}
	o4 := &entity.Node{Value: 3, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}

	o6 := &entity.Node{Value: 8, Next: nil}
	o7 := &entity.Node{Value: 6, Next: o6}
	o8 := &entity.Node{Value: 4, Next: o7}
	o9 := &entity.Node{Value: 2, Next: o8}

	list := mergeTwoList2(o5, o9)

	fmt.Println(list)
}

func TestMergeTwoList3(t *testing.T) {
	o1 := &entity.Node{Value: 9, Next: nil}
	o2 := &entity.Node{Value: 7, Next: o1}
	o3 := &entity.Node{Value: 5, Next: o2}
	o4 := &entity.Node{Value: 3, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}

	o6 := &entity.Node{Value: 8, Next: nil}
	o7 := &entity.Node{Value: 6, Next: o6}
	o8 := &entity.Node{Value: 4, Next: o7}
	o9 := &entity.Node{Value: 2, Next: o8}

	o10 := &entity.Node{Value: 15, Next: nil}
	o11 := &entity.Node{Value: 13, Next: o10}
	o12 := &entity.Node{Value: 12, Next: o11}
	o13 := &entity.Node{Value: 11, Next: o12}

	list := mergeTwoList3(o5, o9, o13)
	fmt.Println(list)
}

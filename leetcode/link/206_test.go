package searchbinary

import (
	"algorithms/study/entity"
	"fmt"
	"testing"
)

func TestReverseLink1(t *testing.T) {
	o1 := &entity.Node{Value: 5, Next: nil}
	o2 := &entity.Node{Value: 4, Next: o1}
	o3 := &entity.Node{Value: 3, Next: o2}
	o4 := &entity.Node{Value: 2, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}
	fmt.Println(o5)
	n5 := reverseLink1(o5)
	fmt.Println(n5)
}

func TestReverseLink2(t *testing.T) {
	o1 := &entity.Node{Value: 5, Next: nil}
	o2 := &entity.Node{Value: 4, Next: o1}
	o3 := &entity.Node{Value: 3, Next: o2}
	o4 := &entity.Node{Value: 2, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}
	fmt.Println(o5)
	n5 := reverseLink2(o5)
	fmt.Println(n5)
}

func TestReverseLink3(t *testing.T) {
	o1 := &entity.Node{Value: 5, Next: nil}
	o2 := &entity.Node{Value: 4, Next: o1}
	o3 := &entity.Node{Value: 3, Next: o2}
	o4 := &entity.Node{Value: 2, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}
	fmt.Println(o5)
	n5 := reverseLink3(o5)
	fmt.Println(n5)
}

func TestReverseLink4(t *testing.T) {
	o1 := &entity.Node{Value: 5, Next: nil}
	o2 := &entity.Node{Value: 4, Next: o1}
	o3 := &entity.Node{Value: 3, Next: o2}
	o4 := &entity.Node{Value: 2, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}
	fmt.Println(o5)
	n5 := reverseLink4(o5)
	fmt.Println(n5)
}

func TestReverseLink5(t *testing.T) {
	o1 := &entity.Node{Value: 5, Next: nil}
	o2 := &entity.Node{Value: 4, Next: o1}
	o3 := &entity.Node{Value: 3, Next: o2}
	o4 := &entity.Node{Value: 2, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}
	fmt.Println(o5)
	n5 := reverseLink5(o5)
	fmt.Println(n5)
}

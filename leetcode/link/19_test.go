package link

import (
	"algorithms/study/entity"
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	o1 := &entity.Node{Value: 5, Next: nil}
	o2 := &entity.Node{Value: 4, Next: o1}
	o3 := &entity.Node{Value: 3, Next: o2}
	o4 := &entity.Node{Value: 2, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}
	fmt.Println(o5)
	n5 := removeNthFromEnd(o5, 2)
	fmt.Println(n5)
}

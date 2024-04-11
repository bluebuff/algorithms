package link

import (
	"algorithms/study/entity"
	"fmt"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	o1 := &entity.Node{Value: 9, Next: nil}
	o2 := &entity.Node{Value: 7, Next: o1}
	o3 := &entity.Node{Value: 5, Next: o2}
	o4 := &entity.Node{Value: 3, Next: o3}
	o5 := &entity.Node{Value: 1, Next: o4}

	fmt.Println(middleNode(o5))
}

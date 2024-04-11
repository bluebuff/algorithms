package link

import (
	"algorithms/study/entity"
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	{
		o1 := &entity.Node{Value: 5, Next: nil}
		o2 := &entity.Node{Value: 4, Next: o1}
		o3 := &entity.Node{Value: 3, Next: o2}
		o4 := &entity.Node{Value: 2, Next: o3}
		o5 := &entity.Node{Value: 1, Next: o4}
		fmt.Println(isPalindrome(o5))
	}

	{
		o1 := &entity.Node{Value: 1, Next: nil}
		o2 := &entity.Node{Value: 2, Next: o1}
		o3 := &entity.Node{Value: 3, Next: o2}
		o4 := &entity.Node{Value: 2, Next: o3}
		o5 := &entity.Node{Value: 1, Next: o4}
		fmt.Println(isPalindrome(o5))
	}

	{
		o1 := &entity.Node{Value: 1, Next: nil}
		o2 := &entity.Node{Value: 2, Next: o1}
		o3 := &entity.Node{Value: 2, Next: o2}
		o4 := &entity.Node{Value: 1, Next: o3}
		fmt.Println(isPalindrome(o4))
	}
}

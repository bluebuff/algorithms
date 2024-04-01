package recursion

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{6, 5, 4, 3, 2, 1}
	bubbleSort(a)
	fmt.Println(a)
}

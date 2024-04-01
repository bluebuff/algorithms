package recursion

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(a, 5))
	fmt.Println(binarySearch(a, 10))
}

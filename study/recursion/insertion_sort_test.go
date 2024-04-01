package recursion

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	a := []int{6, 5, 4, 3, 2, 1}
	insertionSort(a)
	fmt.Println(a)
}

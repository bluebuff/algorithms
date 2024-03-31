package searchbinary

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchBinary(t *testing.T) {
	a := []int{2, 5, 8}
	target := 4
	i := BinarySearch(a, target)
	fmt.Println(i)

	if i < 0 {
		b := make([]int, len(a)+1)
		// 插入点索引
		insert := int(math.Abs(float64(i) + 1))
		copy(b[:insert], a[:insert])
		b[insert] = target
		copy(b[insert+1:], a[insert:])
		a = b
	}
	fmt.Println(a)
}

func TestSearchBinaryBasic(t *testing.T) {
	a := []int{7, 13, 21, 30, 38, 44, 52, 53}
	fmt.Println(BinarySearchBasic(a, 7))
	fmt.Println(BinarySearchBasic(a, 13))
	fmt.Println(BinarySearchBasic(a, 21))
	fmt.Println(BinarySearchBasic(a, 30))
	fmt.Println(BinarySearchBasic(a, 38))
	fmt.Println(BinarySearchBasic(a, 44))
	fmt.Println(BinarySearchBasic(a, 52))
	fmt.Println(BinarySearchBasic(a, 53))
	fmt.Println(BinarySearchBasic(a, 55))
}

func TestSearchBinaryAlternative(t *testing.T) {
	a := []int{7, 13, 21, 30, 38, 44, 52, 53}
	fmt.Println(BinarySearchAlternative(a, 7))
	fmt.Println(BinarySearchAlternative(a, 13))
	fmt.Println(BinarySearchAlternative(a, 21))
	fmt.Println(BinarySearchAlternative(a, 30))
	fmt.Println(BinarySearchAlternative(a, 38))
	fmt.Println(BinarySearchAlternative(a, 44))
	fmt.Println(BinarySearchAlternative(a, 52))
	fmt.Println(BinarySearchAlternative(a, 53))
	fmt.Println(BinarySearchAlternative(a, 55))
}

func TestSearchBinaryBalance(t *testing.T) {
	a := []int{7, 13, 21, 30, 38, 44, 52, 53}
	fmt.Println(BinarySearchBalance(a, 7))
	fmt.Println(BinarySearchBalance(a, 13))
	fmt.Println(BinarySearchBalance(a, 21))
	fmt.Println(BinarySearchBalance(a, 30))
	fmt.Println(BinarySearchBalance(a, 38))
	fmt.Println(BinarySearchBalance(a, 44))
	fmt.Println(BinarySearchBalance(a, 52))
	fmt.Println(BinarySearchBalance(a, 53))
	fmt.Println(BinarySearchBalance(a, 55))
}

func TestBinarySearchLeftMost(t *testing.T) {
	a := []int{1, 2, 4, 4, 4, 5, 6, 6, 7}
	assert.Equal(t, 0, BinarySearchLeftMost(a, 1))
	assert.Equal(t, 1, BinarySearchLeftMost(a, 2))
	assert.Equal(t, 2, BinarySearchLeftMost(a, 4))
	assert.Equal(t, 5, BinarySearchLeftMost(a, 5))
	assert.Equal(t, 6, BinarySearchLeftMost(a, 6))
	assert.Equal(t, 8, BinarySearchLeftMost(a, 7))

	assert.Equal(t, -1, BinarySearchLeftMost(a, 0))
	assert.Equal(t, -1, BinarySearchLeftMost(a, 9))
}

func TestBinarySearchRightMost(t *testing.T) {
	a := []int{1, 2, 4, 4, 4, 5, 6, 6, 7}
	assert.Equal(t, 0, BinarySearchRightMost(a, 1))
	assert.Equal(t, 1, BinarySearchRightMost(a, 2))
	assert.Equal(t, 4, BinarySearchRightMost(a, 4))
	assert.Equal(t, 5, BinarySearchRightMost(a, 5))
	assert.Equal(t, 7, BinarySearchRightMost(a, 6))
	assert.Equal(t, 8, BinarySearchRightMost(a, 7))

	assert.Equal(t, -1, BinarySearchRightMost(a, 0))
	if a[BinarySearchRightMost(a, 9)] != 9 {
		t.Fatal("the right moust not eq 9")
	}
}

package searchbinary

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchBinary(t *testing.T) {
	a := []int{1, 2, 3}
	assert.Equal(t, 0, searchBasic(a, 1))
	assert.Equal(t, 1, searchBasic(a, 2))
	assert.Equal(t, 2, searchBasic(a, 3))
	assert.Equal(t, -1, searchBasic(a, 5))

	assert.Equal(t, 0, search2(a, 1))
	assert.Equal(t, 1, search2(a, 2))
	assert.Equal(t, 2, search2(a, 3))
	assert.Equal(t, -1, search2(a, 5))

	assert.Equal(t, 0, search3(a, 1))
	assert.Equal(t, 1, search3(a, 2))
	assert.Equal(t, 2, search3(a, 3))
	assert.Equal(t, -1, search3(a, 5))

	assert.Equal(t, 0, searchInsert(a, 1))
	assert.Equal(t, 1, searchInsert(a, 2))
	assert.Equal(t, 2, searchInsert(a, 3))
	assert.Equal(t, 3, searchInsert(a, 5))

}

func TestSearchRange(t *testing.T) {
	{
		a := []int{5, 7, 7, 8, 8, 10}
		assert.Equal(t, []int{3, 4}, searchRange(a, 8))

		assert.Equal(t, []int{-1, -1}, searchRange(a, 6))

		assert.Equal(t, []int{-1, -1}, searchRange([]int{}, 0))
	}
	{
		a := []int{1, 2, 3, 4, 5}

		b := make([]int, len(a)+1) // [1,2,2,3,4,5]

		copy(b[:2], a[:2])
		b[2] = 2
		copy(b[3:], a[2:])

		fmt.Println(b)
	}
}

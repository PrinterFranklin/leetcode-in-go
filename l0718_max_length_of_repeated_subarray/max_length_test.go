package leetcode0718

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLength(t *testing.T) {
	type testArr struct {
		nums1  []int
		nums2  []int
		result int
	}
	a := [4]testArr{
		{nums1: []int{1, 2, 3, 2, 1}, nums2: []int{3, 2, 1, 4, 7}, result: 3},
		{nums1: []int{0, 0, 0, 0, 0}, nums2: []int{0, 0, 0, 0, 0}, result: 5},
		{nums1: []int{1, 2, 3}, nums2: []int{3, 2, 1}, result: 1},
		{nums1: []int{1, 0, 0, 0, 1}, nums2: []int{1, 0, 0, 1, 1}, result: 3},
	}
	for _, v := range a {
		re := findLength(v.nums1, v.nums2)
		assert.Equal(t, v.result, re)
	}
}

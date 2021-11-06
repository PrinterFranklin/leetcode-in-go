package leetcode0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	type testArr struct {
		nums   []int
		target int
		result []int
	}
	a := []testArr{
		{nums: []int{2, 7, 11, 15}, target: 9, result: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, result: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, result: []int{0, 1}},
	}
	for _, v := range a {
		re := twoSum(v.nums, v.target)
		assert.Equal(t, v.result, re)
	}
}

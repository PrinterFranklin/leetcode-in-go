package leetcode0575

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributeCandies(t *testing.T) {
	type testArr struct {
		candyType []int
		result    int
	}
	a := [3]testArr{
		{candyType: []int{1, 1, 2, 2, 3, 3}, result: 3},
		{candyType: []int{1, 1, 2, 3}, result: 2},
		{candyType: []int{6, 6, 6, 6}, result: 1},
	}
	for _, v := range a {
		re := distributeCandies(v.candyType)
		assert.Equal(t, v.result, re)
	}
}

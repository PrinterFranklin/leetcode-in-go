package leetcode0322

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChange(t *testing.T) {
	type testArr struct {
		coins  []int
		amount int
		result int
	}
	a := [4]testArr{
		{coins: []int{1, 2, 5}, amount: 11, result: 3},
		{coins: []int{2}, amount: 3, result: -1},
		{coins: []int{1}, amount: 0, result: 0},
		{coins: []int{1}, amount: 2, result: 2},
	}
	for _, v := range a {
		re := coinChange(v.coins, v.amount)
		assert.Equal(t, v.result, re)
	}
}

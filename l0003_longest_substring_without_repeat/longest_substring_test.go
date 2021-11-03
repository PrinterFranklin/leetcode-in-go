package leetcode0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	type testArr struct {
		s      string
		result int
	}
	a := [4]testArr{
		{s: "abcabcbb", result: 3},
		{s: "bbbbb", result: 1},
		{s: "pwwkew", result: 3},
		{s: "", result: 0},
	}
	for _, v := range a {
		re := lengthOfLongestSubstring(v.s)
		assert.Equal(t, v.result, re)
	}
}

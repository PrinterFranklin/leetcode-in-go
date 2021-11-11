package leetcode0076

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	type testArr struct {
		s      string
		t      string
		result string
	}
	a := []testArr{
		{s: "ADOBECODEBANC", t: "ABC", result: "BANC"},
		{s: "a", t: "a", result: "a"},
		{s: "a", t: "aa", result: ""},
	}
	for _, v := range a {
		re := minWindow(v.s, v.t)
		assert.Equal(t, v.result, re)
	}
}

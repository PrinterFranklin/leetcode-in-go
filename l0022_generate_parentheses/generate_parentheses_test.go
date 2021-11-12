package leetcode0022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParentheses(t *testing.T) {
	type testArr struct {
		n      int
		result []string
	}
	a := []testArr{
		{n: 1, result: []string{"()"}},
		{n: 2, result: []string{"(())", "()()"}},
		{n: 3, result: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, v := range a {
		re := generateParenthesis(v.n)
		assert.Equal(t, v.result, re)
	}
}

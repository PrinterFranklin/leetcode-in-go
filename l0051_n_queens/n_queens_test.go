package leetcode0051

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveNQueens(t *testing.T) {
	type testArr struct {
		n      int
		result [][]string
	}
	a := []testArr{
		{n: 1, result: [][]string{{"Q"}}},
		{n: 2, result: [][]string{}},
		{n: 3, result: [][]string{}},
		{n: 4, result: [][]string{
			{".Q..", "...Q", "Q...", "..Q."},
			{"..Q.", "Q...", "...Q", ".Q.."},
		}},
	}
	for _, v := range a {
		re := solveNQueens(v.n)
		assert.Equal(t, v.result, re)
	}
}

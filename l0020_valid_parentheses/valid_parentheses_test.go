package leetcode0020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	type testArr struct {
		s      string
		result bool
	}
	a := []testArr{
		{s: "([)]", result: false},
		{s: "(", result: false},
		{s: "}", result: false},
		{s: "()[]{}", result: true},
		{s: "[]({()()})", result: true},
	}
	for _, v := range a {
		re := isValid(v.s)
		assert.Equal(t, v.result, re)
	}
}

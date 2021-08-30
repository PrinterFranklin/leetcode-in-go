package l0001_add_two_numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCase_1(t *Testing.T) {
	nums := [...]int{2, 7, 11, 15}
	target := 9
	assert.Equal(t, 0, twoSum(&nums, target)[0])
	assert.Equal(t, 1, twoSum(&nums, target)[1])
}
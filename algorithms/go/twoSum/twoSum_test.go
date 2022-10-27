package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	nums := []int{
		2, 7, 11, 15,
	}
	expected := []int{1, 0}
	assert.Equal(t, expected, twoSum(nums, 9))
}

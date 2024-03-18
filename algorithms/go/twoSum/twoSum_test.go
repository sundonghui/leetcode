package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	nums := []int{
		2, 7, 11, 15,
	}
	expected := []int{1, 2}
	assert.Equal(t, expected, twoSum(nums, 9))
}

func Test_twoSumTwoPointer(t *testing.T) {
	nums := []int{
		2, 7, 11, 15,
	}
	expected := []int{1, 2}
	assert.Equal(t, expected, twoSumTwoPointer(nums, 9))
}

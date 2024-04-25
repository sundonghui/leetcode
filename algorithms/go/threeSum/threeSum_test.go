package threesum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	expected := [][]int{
		{
			-1, -1, 2,
		},
		{
			-1, 0, 1,
		},
	}
	assert.Equal(t, expected, threeSum_self(nums))
}

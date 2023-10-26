package searchrange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	assert.Equal(t, []int{3, 4}, searchRange(nums, target))
}

func Test_searchRange1(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	assert.Equal(t, []int{-1, -1}, searchRange(nums, target))
}

func Test_searchRange2(t *testing.T) {
	nums := []int{}
	target := 0
	assert.Equal(t, []int{-1, -1}, searchRange(nums, target))
}

func Test_searchRange3(t *testing.T) {
	nums := []int{2, 2}
	target := 2
	assert.Equal(t, []int{0, 1}, searchRange(nums, target))
}

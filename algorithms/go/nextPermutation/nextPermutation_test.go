package nextpermutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	expected := []int{1, 3, 2}
	assert.Equal(t, nums, expected)
}

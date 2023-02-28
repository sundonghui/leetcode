package threesumclosest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSumClosest(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	target := 1
	expected := 2
	assert.Equal(t, expected, threeSumClosest(nums, target))
}

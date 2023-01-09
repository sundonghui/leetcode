package maxslidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	expected := []int{3, 3, 5, 5, 6, 7}
	assert.Equal(t, expected, maxSlidingWindow(nums, k))
}

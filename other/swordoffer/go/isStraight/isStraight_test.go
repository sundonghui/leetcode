package isstraight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStraight(t *testing.T) {
	nums := []int{0, 0, 8, 5, 4}
	assert.Equal(t, true, isStraight(nums))

	nums1 := []int{1, 2, 12, 0, 3}
	assert.Equal(t, false, isStraight(nums1))
}

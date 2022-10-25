package findrepeatnumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	nums := []int{3, 2, 4, 3}
	nums1 := []int{3, 2, 3, 4, 5, 6, 8, 9, 0}
	expected := 3
	assert.Equal(t, expected, findRepeatNumber(nums))
	assert.Equal(t, expected, findRepeatNumber(nums1))
}

package findpeakelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPeakElement(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	assert.Equal(t, 2, findPeakElement(nums))
}

func Test_findPeakElement1(t *testing.T) {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	assert.Equal(t, 5, findPeakElement(nums))
}

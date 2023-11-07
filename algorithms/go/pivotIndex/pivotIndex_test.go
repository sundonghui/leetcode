package pivotindex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pivotIndex(t *testing.T) {
	nums := []int{2, 1, -1}
	assert.Equal(t, 0, pivotIndex(nums))
}

func Test_pivotIndex1(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	assert.Equal(t, 3, pivotIndex(nums))
}

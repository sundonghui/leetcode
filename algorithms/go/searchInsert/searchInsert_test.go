package searchinsert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	assert.Equal(t, 2, searchInsert(nums, target))
}

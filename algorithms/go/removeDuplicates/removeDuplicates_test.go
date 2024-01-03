package removeduplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 4, 5}
	assert.Equal(t, 5, removeDuplicates(nums))
}

func Test_removeDuplicates2(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	assert.Equal(t, removeDuplicates2(nums), 5)
}

package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 4, 4, 5, 5, 6}
	assert.Equal(t, 3, search(nums, 4))
}

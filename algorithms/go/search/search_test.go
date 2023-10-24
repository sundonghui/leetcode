package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	assert.Equal(t, 4, search(nums, target))
}

func Test_search1(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	assert.Equal(t, -1, search(nums, target))
}

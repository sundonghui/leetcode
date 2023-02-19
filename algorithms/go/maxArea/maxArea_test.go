package maxarea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	assert.Equal(t, 49, maxArea(height))
}

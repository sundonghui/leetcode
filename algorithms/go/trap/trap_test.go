package trap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tarp(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	assert.Equal(t, 6, trap(height))
}

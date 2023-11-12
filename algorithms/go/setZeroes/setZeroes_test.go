package setzeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setZeroes(t *testing.T) {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	expected := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}
	setZeroes(matrix)
	assert.Equal(t, expected, matrix)
}

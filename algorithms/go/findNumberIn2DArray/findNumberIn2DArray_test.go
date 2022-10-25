package findnumberin2darray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNumberIn2DArray(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
	}
	expected := true
	assert.Equal(t, expected, findNumberIn2DArray(matrix, 5))
}

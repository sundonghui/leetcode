package findnumberin2darray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindNumberIn2DArray(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
	}
	expected := true
	assert.Equal(t, expected, findNumberIn2DArray(matrix, 5))
}

func TestFindNumberIn2DArray2(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
	}
	matrix1 := [][]int{
		{-5},
	}
	matrix2 := [][]int{
		{1, 1},
	}
	expected := true
	assert.Equal(t, expected, findNumberIn2DArray2(matrix, 5))
	assert.Equal(t, expected, findNumberIn2DArray2(matrix1, -5))
	assert.Equal(t, false, findNumberIn2DArray2(matrix2, 2))
}

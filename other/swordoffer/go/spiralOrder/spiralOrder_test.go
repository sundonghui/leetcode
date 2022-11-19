package spiralorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSpiralOrder(t *testing.T) {
	matrix1 := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	expected1 := []int{1,2,3,6,9,8,7,4,5}
	assert.Equal(t, expected1, spiralOrder(matrix1))

	matrix2 := [][]int{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12},
	}
	expected2 := []int{1,2,3,4,8,12,11,10,9,5,6,7}
	assert.Equal(t, expected2, spiralOrder(matrix2))
	
}
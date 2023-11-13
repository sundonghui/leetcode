package finddiagonalorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDiagonalOrder(t *testing.T) {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := []int{1, 2, 4, 7, 5, 3, 6, 8, 9}
	assert.Equal(t, expected, findDiagonalOrder(mat))
}

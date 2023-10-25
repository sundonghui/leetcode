package findmin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMin(t *testing.T) {
	nums := []int{4, 5, 6, 7, 8, 1, 2, 3, 4}
	assert.Equal(t, 1, findMin(nums))
}

func Test_findMin1(t *testing.T) {
	nums := []int{11, 13, 15, 17}
	assert.Equal(t, 11, findMin(nums))
}

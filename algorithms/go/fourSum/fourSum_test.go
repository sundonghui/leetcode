package foursum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSum(t *testing.T) {
	expected := [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}
	arr := []int{1, 0, -1, 0, -2, 2}
	assert.Equal(t, expected, fourSum(arr, 0))
}

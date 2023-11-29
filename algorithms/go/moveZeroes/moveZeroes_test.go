package movezeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 1}
	expected := []int{1, 3, 1, 0, 0}
	moveZeroes(nums)
	assert.Equal(t, expected, nums)
}

func Test_moveZeroes_1(t *testing.T) {
	nums := []int{0, 0, 1}
	expected := []int{1, 0, 0}
	moveZeroes(nums)
	assert.Equal(t, expected, nums)
}

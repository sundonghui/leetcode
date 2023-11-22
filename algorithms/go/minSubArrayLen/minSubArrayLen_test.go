package minsubarraylen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	assert.Equal(t, 2, minSubArrayLen(target, nums))
}

func Test_minSubArrayLen1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	target := 11
	assert.Equal(t, 3, minSubArrayLen(target, nums))
}

func Test_minSubArrayLen2(t *testing.T) {
	target := 213
	nums := []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}
	assert.Equal(t, 8, minSubArrayLen(target, nums))
}

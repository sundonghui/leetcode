package singlenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	nums := []int{1, 2, 2, 3, 3}
	assert.Equal(t, 1, singleNumber(nums))
}

func TestSingleNumber2(t *testing.T) {
	nums := []int{1, 2, 2, 3, 3}
	assert.Equal(t, 1, singleNumber2(nums))
}

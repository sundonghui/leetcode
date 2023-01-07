package singlenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	nums := []int{9, 1, 7, 9, 7, 9, 7}
	assert.Equal(t, 1, singleNumber(nums))
}

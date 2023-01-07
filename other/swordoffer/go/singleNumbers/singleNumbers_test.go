package singlenumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumbers(t *testing.T) {
	nums := []int{4, 1, 4, 6}
	expected := []int{1, 6}
	assert.Equal(t, expected, singleNumbers(nums))
}

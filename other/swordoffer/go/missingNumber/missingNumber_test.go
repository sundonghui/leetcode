package missingnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 8, 9}
	expected := 7
	assert.Equal(t, expected, missingNumber(nums))
}

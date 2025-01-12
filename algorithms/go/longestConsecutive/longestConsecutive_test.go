package longestconsecutive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	numbers := []int{100, 4, 200, 1, 3, 2}
	assert.Equal(t, longestConsecutive(numbers), 4) // [1, 2, 3, 4]
}
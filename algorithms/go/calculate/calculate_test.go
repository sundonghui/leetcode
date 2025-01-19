package calculate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculate_example1(t *testing.T) {
	// Example 1:
	// Input: s = "3+2+2"
	// Output: 7
	// Explanation: (3 + 2) * 2 = 7.
	assert.Equal(t, 7, calculate("3+2+2"))
}

func Test_calculate_example2(t *testing.T) {
	// Example 2:
	// Input: s = "(1+(4+5+2)-3)+(6+8)"
	// Output: 23
	assert.Equal(t, 23, calculate("(1+(4+5+2)-3)+(6+8)"))
}

func Test_calculate_example3(t *testing.T) {
    // Example 3:
	// Input: s = ""
	// Output: 0
	assert.Equal(t, 0, calculate(""))
}

func Test_calculate_example4(t *testing.T) {
	// Example 4:
	// Input: s = " 2-1 + 2 "
	// Output: 3
	assert.Equal(t, 3, calculate(" 2-1 + 2 "))
}
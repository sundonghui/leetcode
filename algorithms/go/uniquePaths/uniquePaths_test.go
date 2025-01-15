package uniquepaths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	// Test case 1
	m1 := 3
	n1 := 2
	expected1 := 3
	result1 := uniquePaths(m1, n1)
	assert.Equal(t, expected1, result1)
}

func TestUniquePaths2(t *testing.T) {
	// Test case 2
	m2 := 7
	n2 := 3
	expected2 := 28
	result2 := uniquePaths(m2, n2)
	assert.Equal(t, expected2, result2)
}

func TestUniquePaths3(t *testing.T) {
    m3 := 3
	n3 := 3
	expected3 := 6
	result3 := uniquePaths(m3, n3)
	assert.Equal(t, expected3, result3)
}
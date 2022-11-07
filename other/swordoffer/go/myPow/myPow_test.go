package mypow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyPow(t *testing.T) {
	expected1,expected2, expected3 := 1024.00000, 10.648000000000003, 0.25000
	assert.Equal(t, expected1, myPow(2.0, 10))
	assert.Equal(t, expected2, myPow(2.2, 3))
	assert.Equal(t, expected3, myPow(2.0, -2))
}
package fib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	n := 5
	expected := 5
	assert.Equal(t, expected, fib(n))
	assert.Equal(t, expected, fib2(n))
	assert.Equal(t, expected, fib3(n))
}

package minarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinArray(t *testing.T) {
	expected := 1
	array := []int{
		3, 4, 5, 1, 2,
	}
	assert.Equal(t, expected, minArray(array))
}

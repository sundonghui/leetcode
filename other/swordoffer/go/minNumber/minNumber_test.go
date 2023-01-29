package minnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinNumber(t *testing.T) {
	//assert.Equal(t, "102", minNumber([]int{10, 2}))
	assert.Equal(t, "3033459", minNumber([]int{3, 30, 34, 5, 9}))
	assert.Equal(t, "3033459", MinNumber([]int{3, 30, 34, 5, 9}))
}

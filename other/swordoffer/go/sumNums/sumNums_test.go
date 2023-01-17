package sumnums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumNums(t *testing.T) {
	assert.Equal(t, 6, sumNums(3))
	assert.Equal(t, 45, sumNums(9))
}

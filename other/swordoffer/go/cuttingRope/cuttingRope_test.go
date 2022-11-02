package cuttingrope

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCuttingRope(t *testing.T) {
	assert.Equal(t, 36, cuttingRope(10))
	assert.Equal(t, 36, cuttingRope2(10))
}

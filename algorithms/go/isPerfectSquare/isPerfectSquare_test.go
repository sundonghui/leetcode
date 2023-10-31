package isperfectsquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPerfectSquare(t *testing.T) {
	assert.Equal(t, true, isPerfectSquare(16))
}

func Test_isPerfectSquare1(t *testing.T) {
	assert.Equal(t, false, isPerfectSquare(15))
}

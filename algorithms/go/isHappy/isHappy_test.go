package ishappy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHappy_19(t *testing.T) {
	assert.Equal(t, true, isHappy(19))
}

func TestIsHappy_2(t *testing.T) {
	assert.Equal(t, false, isHappy(2))
}

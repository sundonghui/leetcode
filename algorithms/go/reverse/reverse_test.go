package reverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	x := 123
	assert.Equal(t, 321, reverse(x))
	y := -123
	assert.Equal(t, -321, reverse(y))
	z := 120
	assert.Equal(t, 21, reverse(z))
	assert.Equal(t, 0, reverse(0))
	w := 1563847412
	assert.Equal(t, 0, reverse(w))
}

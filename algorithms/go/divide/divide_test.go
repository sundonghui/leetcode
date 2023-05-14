package divide

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	dividend := 10
	divisor := 3
	assert.Equal(t, 3, divide(dividend, divisor))
}
func TestDivide2(t *testing.T) {
	dividend := -10
	divisor := -3
	assert.Equal(t, 3, divide(dividend, divisor))
}
func TestDivide3(t *testing.T) {
	dividend := 10
	divisor := -3
	assert.Equal(t, -3, divide(dividend, divisor))
}

func TestDivide4(t *testing.T) {
	dividend := -2147483648
	divisor := -1
	assert.Equal(t, 2147483647, divide(dividend, divisor))
}

func TestDivide5(t *testing.T) {
	dividend := -2147483648
	divisor := 1
	assert.Equal(t, -2147483648, divide(dividend, divisor))
}

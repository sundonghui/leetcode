package inttoroman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	assert.Equal(t, "LVIII", intToRoman(58))
	assert.Equal(t, "MCMXCIV", intToRoman(1994))
}

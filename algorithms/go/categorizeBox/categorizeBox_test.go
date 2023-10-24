package categorizebox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_categorizeBox(t *testing.T) {
	length := 1000
	width := 35
	height := 700
	mass := 300
	assert.Equal(t, "Heavy", categorizeBox(length, width, height, mass))
}

func Test_categorizeBox1(t *testing.T) {
	length := 200
	width := 50
	height := 800
	mass := 50
	assert.Equal(t, "Neither", categorizeBox(length, width, height, mass))
}

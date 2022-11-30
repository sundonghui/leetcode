package countdigitone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDigitOne(t *testing.T) {
	assert.Equal(t, 5, countDigitOne(12))
}

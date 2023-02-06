package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 3
	expected := "PAHNAPLSIIGYIR"
	assert.Equal(t, expected, convert(s, numRows))
}

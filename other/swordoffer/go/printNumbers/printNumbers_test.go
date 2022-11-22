package printnumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintNumbers(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expected, printNumbers(1))
}

package findnthdigit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindNthDigit(t *testing.T) {
	assert.Equal(t, 1, findNthDigit(12))
}

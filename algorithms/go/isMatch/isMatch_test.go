package ismatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	a := "pp"
	b := "aa"
	c := "a*"
	d := ".*"
	assert.Equal(t, false, isMatch(a, b))
	assert.Equal(t, true, isMatch(b, c))
	assert.Equal(t, true, isMatch(b, d))
}

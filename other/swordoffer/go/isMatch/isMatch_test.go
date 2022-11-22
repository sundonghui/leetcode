package ismatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	s1, p1 := "aa", "a"
	assert.Equal(t, false, isMatch(s1, p1))
	p2 := "a*"
	assert.Equal(t, true, isMatch(s1, p2))
	p3 := ".*"
	assert.Equal(t, true, isMatch(s1, p3))
}

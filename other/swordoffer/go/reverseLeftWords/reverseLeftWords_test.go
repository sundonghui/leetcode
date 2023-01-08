package reverseleftwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseLeftWords(t *testing.T) {
	s := "abcdefg"
	k := 2
	expected := "cdefgab"
	assert.Equal(t, expected, reverseLeftWords(s, k))
}

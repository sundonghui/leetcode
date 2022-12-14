package firstuniqchar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUniqChar(t *testing.T) {
	expected := []byte("l")
	assert.Equal(t, expected[0], firstUniqChar("leetcode"))
}

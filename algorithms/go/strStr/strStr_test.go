package strstr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	haystack := "sadbutsad"
	needle := "sad"
	assert.Equal(t, 0, strStr(haystack, needle))
}

func TestStrStr2(t *testing.T) {
	haystack := "aaa"
	needle := "aaaa"
	assert.Equal(t, -1, strStr(haystack, needle))
}

func TestStrStr3(t *testing.T) {
	haystack := "a"
	needle := "a"
	assert.Equal(t, 0, strStr(haystack, needle))
}

func TestStrStr4(t *testing.T) {
	haystack := "abc"
	needle := "c"
	assert.Equal(t, 2, strStr(haystack, needle))
}

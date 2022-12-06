package lengthoflongestsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s1 := "pwwkew"
	s2 := "abcabcbb"
	s3 := "bbbbb"
	assert.Equal(t, 3, lengthOfLongestSubstring(s1))
	assert.Equal(t, 3, lengthOfLongestSubstring(s2))
	assert.Equal(t, 1, lengthOfLongestSubstring(s3))
}

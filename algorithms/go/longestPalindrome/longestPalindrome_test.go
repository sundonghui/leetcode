package longestpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	expected := "bab"
	assert.Equal(t, expected, longestPalindrome(s))
}

func TestLongestPalindrome1(t *testing.T) {
	s := "cbbd"
	expected := "bb"
	assert.Equal(t, expected, longestPalindrome(s))
}

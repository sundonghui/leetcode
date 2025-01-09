package wordpattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordPattern_abba(t *testing.T) {
	pattern := "abba"
	s := "dog cat cat dog"
	assert.Equal(t, true, wordPattern(pattern, s))
}

func TestWordPattern_abab(t *testing.T) {
	pattern := "abab"
	s := "dog cat dog cat"
	assert.Equal(t, true, wordPattern(pattern, s))
}

func TestWordPattern_aaaa(t *testing.T) {
	pattern := "aaaa"
	s := "dog cat dog cat"
	assert.Equal(t, false, wordPattern(pattern, s))
}

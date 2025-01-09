package isanagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram_true(t *testing.T) {
	s := "anagram"
	t1 := "nagaram"
	assert.Equal(t, true, isAnagram(s, t1))
}

func TestIsAnagram_false(t *testing.T) {
	s := "rat"
	t1 := "car"
	assert.Equal(t, false, isAnagram(s, t1))
}

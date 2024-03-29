package reversewords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseWords(t *testing.T) {
	s := "the sky is blue"
	expected := "blue is sky the"

	assert.Equal(t, expected, reverseWords(s))
}

func Test_reverseWords1(t *testing.T) {
	s := " hello world "
	expected := "world hello"

	assert.Equal(t, expected, reverseWords(s))
}

func Test_ReverseWords(t *testing.T) {
	s := "Let's take LeetCode contest"
	expected := "s'teL ekat edoCteeL tsetnoc"
	assert.Equal(t, expected, ReverseWords(s))
}

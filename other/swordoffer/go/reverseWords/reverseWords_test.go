package reversewords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	str := "the sky is blue"
	revStr := "blue is sky the"
	str1 := "a good   example"
	revStr1 := "example good a"
	assert.Equal(t, revStr, reverseWords(str))
	assert.Equal(t, revStr1, reverseWords(str1))
}

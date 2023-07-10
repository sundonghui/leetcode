package longestvalidparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestValidParentheses(t *testing.T) {
	str := "(((())))"
	assert.Equal(t, 8, longestValidParentheses(str))
}

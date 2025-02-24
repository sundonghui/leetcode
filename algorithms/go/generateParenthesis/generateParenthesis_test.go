package generateparenthesis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	assert.Equal(t, expected, generateParenthesis(3))
}

func TestGenerateParenthesis2(t *testing.T) {
	expected := []string{"()"}
	assert.Equal(t, expected, generateParenthesis(1))
}

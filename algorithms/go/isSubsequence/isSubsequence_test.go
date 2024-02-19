package issubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSubsequence(t *testing.T) {
	s := "abc"
	p := "ahbgdc"
	assert.Equal(t, true, isSubsequence(s, p))
}

func Test_isSubsequence_1(t *testing.T) {
	s := "axc"
	p := "ahbgdc"
	assert.Equal(t, false, isSubsequence(s, p))
}

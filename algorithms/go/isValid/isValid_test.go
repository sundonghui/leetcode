package isvalid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	str := "([])"
	str1 := "[(])"
	str2 := "(){}}{"
	str3 := "]"
	assert.Equal(t, true, isValid(str))
	assert.Equal(t, false, isValid(str1))
	assert.Equal(t, false, isValid(str2))
	assert.Equal(t, false, isValid(str3))
}

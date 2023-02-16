package myatoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	str := "     42"
	str1 := "   -65 sshhshs"
	str2 := "+1"
	str3 := "-+1"
	str4 := "+-1"
	assert.Equal(t, 42, myAtoi(str))
	assert.Equal(t, -65, myAtoi(str1))
	assert.Equal(t, 1, myAtoi(str2))
	assert.Equal(t, 0, myAtoi(str3))
	assert.Equal(t, 0, myAtoi(str4))
}

package isnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber(t *testing.T) {
	intListStr := []string{"+100", "5e2", "-123", "3.1416", "-1E-16", "0123"}
	unIntListStr := []string{"12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"}
	for _, v := range intListStr {
		assert.Equal(t, true, isNumber(v))
	}
	for _, v := range unIntListStr {
		assert.Equal(t, false, isNumber(v))
	}
}

package canconstruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canConstruct(t *testing.T) {
	a := "aa"
	b := "aab"
	assert.Equal(t, true, canConstruct(a, b))
}

func Test_canConstruct_false(t *testing.T) {
	a := "a"
	b := "b"
	assert.Equal(t, false, canConstruct(a, b))
}

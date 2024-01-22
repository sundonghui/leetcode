package lengthoflastword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLastWord(t *testing.T) {
	s := "Hello World"
	assert.Equal(t, 5, lengthOfLastWord(s))
}

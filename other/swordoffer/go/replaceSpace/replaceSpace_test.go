package replacespace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceString(t *testing.T) {
	str := "We are happy."
	expected := "We%20are%20happy."
	assert.Equal(t, expected, replaceSpace(str))
}

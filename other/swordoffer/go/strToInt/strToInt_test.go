package strtoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrToInt(t *testing.T) {
	assert.Equal(t, -345, strToInt("       -345"))
}

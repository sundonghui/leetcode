package getrow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getRow(t *testing.T) {
	expected := []int{1, 3, 3, 1}
	assert.Equal(t, expected, getRow(3))
}

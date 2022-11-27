package permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutation(t *testing.T) {
	expected := []string{"abc", "acb", "bac", "bca", "cba", "cab"}
	assert.Equal(t, expected, permutation("abc"))
}

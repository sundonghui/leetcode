package hammingweight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingWeight(t *testing.T) {
	expected1, expected2, expected3 := 3, 1, 31
	assert.Equal(t, expected1, hammingWeight(11))
	assert.Equal(t, expected2, hammingWeight(128))
	assert.Equal(t, expected3, hammingWeight(4294967293))
	assert.Equal(t, expected1, hammingWeight1(11))
	assert.Equal(t, expected2, hammingWeight1(128))
	assert.Equal(t, expected3, hammingWeight1(4294967293))
}

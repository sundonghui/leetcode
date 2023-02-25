package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{
		"flower", "flow", "flight",
	}
	assert.Equal(t, "fl", longestCommonPrefix(strs))
	strs1 := []string{
		"dog", "racecar", "car",
	}
	assert.Equal(t, "", longestCommonPrefix(strs1))
}

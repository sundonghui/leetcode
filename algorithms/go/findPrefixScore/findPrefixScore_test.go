package findprefixscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPrefixScore(t *testing.T) {
	nums := []int{2,3,7,5,10}
	expected := []int64{4,10,24,36,56}
	assert.Equal(t, expected, findPrefixScore(nums))
}
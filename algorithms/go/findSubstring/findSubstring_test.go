package findsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubstring(t *testing.T) {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	expected := []int{0, 9}
	assert.Equal(t, expected, findSubstring(s, words))
}

func TestFindSubstring2(t *testing.T) {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "word"}
	expected := []int{}
	assert.Equal(t, expected, findSubstring(s, words))
}

func TestFindSubstring3(t *testing.T) {
	s := "barfoofoobarthefoobarman"
	words := []string{"bar", "foo", "the"}
	expected := []int{6, 9, 12}
	assert.Equal(t, expected, findSubstring(s, words))
}

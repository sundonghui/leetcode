package groupanagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams_example1(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{
		{"eat", "tea", "ate"},
		{"tan", "nat"},
		{"bat"},
	}
	result := groupAnagrams(strs)
	if len(result) != len(expected) {
		t.Errorf("Expected %d groups, got %d", len(expected), len(result))
	}
	assert.Equal(t, expected, result)
}

func TestGroupAnagrams_example2(t *testing.T) {
	strs := []string{}
	expected := [][]string{}
	result := groupAnagrams(strs)
	if len(result) != len(expected) {
		t.Errorf("Expected %d groups, got %d", len(expected), len(result))
	}
	assert.Equal(t, expected, result)
}

func TestGroupAnagrams_example3(t *testing.T) {
	strs := []string{"a"}
	expected := [][]string{{"a"}}
	result := groupAnagrams(strs)
	if len(result) != len(expected) {
		t.Errorf("Expected %d groups, got %d", len(expected), len(result))
	}
}

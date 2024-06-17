package isisomorphic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isIsomorphic(t *testing.T) {
	ss := "egg"
	ts := "add"
	assert.Equal(t, true, isIsomorphic(ss, ts))
}

func Test_isIsomorphic_false(t *testing.T) {
	ss := "foo"
	ts := "bar"
	assert.Equal(t, false, isIsomorphic(ss, ts))
}

func Test_isIsomorphic_false_fix(t *testing.T) {
	ss := "badc"
	ts := "baba"
	assert.Equal(t, false, isIsomorphic(ss, ts))
}

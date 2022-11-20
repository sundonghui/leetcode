package validatestacksequences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateStackSequences(t *testing.T) {
	pushed := []int{1,2,3,4,5}
	poped := []int{4,5,3,2,1}
	assert.Equal(t, true, validateStackSequences(pushed, poped))

	pushedAgain :=  []int{1,2,3,4,5}
	popedAgain := []int{4,3,5,1,2}
	assert.Equal(t, false, validateStackSequences(pushedAgain, popedAgain))
}
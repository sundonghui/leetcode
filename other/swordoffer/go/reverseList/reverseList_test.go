package reverselist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T)  {
	head := &ListNode{
		Val: 2,	
		Next: &ListNode{
			Val: 1,
		},	
	}	
	expected := &ListNode{
		Val: 1,	
		Next: &ListNode{
			Val: 2,
		},	
	}	
	assert.Equal(t, expected, reverseList(head))
}
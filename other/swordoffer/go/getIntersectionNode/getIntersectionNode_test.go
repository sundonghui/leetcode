package getintersectionnode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIntersectionNode(t *testing.T) {
	head := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	headA := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: head,
		},	
	}
	headB := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: head,
			},
		},	
	}		
	assert.Equal(t, head, getIntersectionNode(headA, headB))
}
package swappairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	expected := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	assert.Equal(t, &expected, swapPairs(&head))
}

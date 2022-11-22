package getkthfromend

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetKthFromEnd(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	expected := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
		},
	}
	assert.Equal(t, expected, getKthFromEnd(head, 2))
}

package deletenode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteNode(t *testing.T) {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
					Next: nil,
				},
			},
		},
	}
	expected := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 9,
				Next: nil,
			},
		},
	}
	assert.Equal(t, expected, deleteNode(head, 1))
	assert.Equal(t, expected, deleteNode1(head, 1))
}
package reversekgroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseKGroup(t *testing.T) {
	head := ListNode{
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
	expected := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	assert.Equal(t, &expected, reverseKGroup(&head, 2))
}

func TestReverseKGroup2(t *testing.T) {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	expected := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
		},
	}
	assert.Equal(t, &expected, reverseKGroup(&head, 2))
}

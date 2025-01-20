package hascycle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasCycle_example1(t *testing.T) {
	// Input: head = [3,2,0,-4], pos = 1
	// Output: true
	// Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
	head := &ListNode{
		Val:  3,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  0,
				Next: &ListNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}	
	head.Next.Next.Next.Next = head.Next
	assert.Equal(t, true, hasCycle(head))			
}

func Test_hasCycle_example2(t *testing.T) {
	// Input: head = [1,2], pos = 0
	// Output: true	
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	head.Next.Next = head
	assert.Equal(t, true, hasCycle(head))
}

func Test_hasCycle_example3(t *testing.T) {
	// Input: head = [], pos = -1
	// Output: false
	head := &ListNode{}
	assert.Equal(t, false, hasCycle(head))	
}
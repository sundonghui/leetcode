package reverseprint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversePrint(t *testing.T) {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
			},
		},
	}
	expected := []int{
		2, 3, 1,
	}
	assert.Equal(t, expected, reversePrint(&head))
	assert.Equal(t, expected, reversePrint2(&head))

}

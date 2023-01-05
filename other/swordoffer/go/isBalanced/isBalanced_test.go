package isbalanced

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBalanced(t *testing.T) {
	n1 := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	n2 := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}
	assert.Equal(t, true, isBalanced(&n1))
	assert.Equal(t, false, isBalanced(&n2))
}
